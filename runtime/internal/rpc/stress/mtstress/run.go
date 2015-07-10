// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"v.io/v23"
	"v.io/v23/context"

	"v.io/x/ref/lib/stats/histogram"
)

// params encapsulates "input" information to the loadtester.
type params struct {
	NetworkDistance time.Duration // Distance (in time) of the server from this driver
	Rate            float64       // Desired rate of sending RPCs (RPC/sec)
	Duration        time.Duration // Duration over which loadtest traffic should be sent
	Context         *context.T
	Reauthenticate  bool // If true, each RPC should establish a network connection (and authenticate)
}

// report is generated by a run of the loadtest.
type report struct {
	Count      int64         // Count of RPCs sent
	Elapsed    time.Duration // Time period over which Count RPCs were sent
	AvgLatency time.Duration
	HistMS     *histogram.Histogram // Histogram of latencies in milliseconds
}

func (r *report) Print(params params) error {
	if r.HistMS != nil {
		fmt.Println("RPC latency histogram (in ms):")
		fmt.Println(r.HistMS.Value())
	}
	actualRate := float64(r.Count*int64(time.Second)) / float64(r.Elapsed)
	fmt.Printf("Network Distance: %v\n", params.NetworkDistance)
	fmt.Printf("#RPCs sent:       %v\n", r.Count)
	fmt.Printf("RPCs/sec sent:    %.2f (%.2f%% of the desired rate of %v)\n", actualRate, actualRate*100/params.Rate, params.Rate)
	fmt.Printf("Avg. Latency:     %v\n", r.AvgLatency)
	// Mark the results are tainted if the deviation from the desired rate is too large
	if 0.9*params.Rate > actualRate {
		return fmt.Errorf("TAINTED RESULTS: drove less traffic than desired: either server or loadtester had a bottleneck")
	}
	return nil
}

func run(f func(*context.T) (time.Duration, error), p params) error {
	var (
		ticker    = time.NewTicker(time.Duration(float64(time.Second) / p.Rate))
		latency   = make(chan time.Duration, 1000)
		started   = time.Now()
		stop      = time.After(p.Duration)
		interrupt = make(chan os.Signal)
		ret       report
	)
	defer ticker.Stop()
	warmup(p.Context, f)
	signal.Notify(interrupt, os.Interrupt)
	defer signal.Stop(interrupt)

	stopped := false
	var sumMS int64
	var lastInterrupt time.Time
	for !stopped {
		select {
		case <-ticker.C:
			go call(p.Context, f, p.Reauthenticate, latency)
		case d := <-latency:
			if ret.HistMS != nil {
				ret.HistMS.Add(int64(d / time.Millisecond))
			}
			ret.Count++
			sumMS += int64(d / time.Millisecond)
			// Use 10 samples to determine how the histogram should be setup
			if ret.Count == 10 {
				avgms := sumMS / ret.Count
				opts := histogram.Options{
					NumBuckets: 32,
					// Mostly interested in tail latencies,
					// so have the histogram start close to
					// the current average.
					MinValue:     int64(float64(avgms) * 0.95),
					GrowthFactor: 0.20,
				}
				p.Context.Infof("Creating histogram after %d samples (%vms avg latency): %+v", ret.Count, avgms, opts)
				ret.HistMS = histogram.New(opts)
			}
		case sig := <-interrupt:
			if time.Since(lastInterrupt) < time.Second {
				p.Context.Infof("Multiple %v signals received, aborting test", sig)
				stopped = true
				break
			}
			lastInterrupt = time.Now()
			// Print a temporary report
			fmt.Println("INTERMEDIATE REPORT:")
			ret.Elapsed = time.Since(started)
			ret.AvgLatency = time.Duration(float64(sumMS)/float64(ret.Count)) * time.Millisecond
			if err := ret.Print(p); err != nil {
				fmt.Println(err)
			}
			fmt.Println("--------------------------------------------------------------------------------")
		case <-stop:
			stopped = true
		}
	}
	ret.Elapsed = time.Since(started)
	ret.AvgLatency = time.Duration(float64(sumMS)/float64(ret.Count)) * time.Millisecond
	return ret.Print(p)
}

func warmup(ctx *context.T, f func(*context.T) (time.Duration, error)) {
	const nWarmup = 10
	ctx.Infof("Sending %d requests as warmup", nWarmup)
	var wg sync.WaitGroup
	for i := 0; i < nWarmup; i++ {
		wg.Add(1)
		go func() {
			f(ctx)
			wg.Done()
		}()
	}
	wg.Wait()
	ctx.Infof("Done with warmup")
}

func call(ctx *context.T, f func(*context.T) (time.Duration, error), reauth bool, d chan<- time.Duration) {
	client := v23.GetClient(ctx)
	if reauth {
		// HACK ALERT: At the time the line below was written, it was
		// known that the implementation would cause 'ctx' to be setup
		// such that any subsequent RPC will establish a network
		// connection from scratch (a new VIF, new VC etc.) If that
		// implementation changes, then this line below will have to
		// change!
		var err error
		if ctx, err = v23.WithPrincipal(ctx, v23.GetPrincipal(ctx)); err != nil {
			ctx.Infof("%v", err)
			return
		}
		client = v23.GetClient(ctx)
		defer client.Close()
	}
	sample, err := f(ctx)
	if err != nil {
		ctx.Infof("%v", err)
		return
	}
	d <- sample
}
