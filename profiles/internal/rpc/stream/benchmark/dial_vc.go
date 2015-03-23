package benchmark

import (
	"testing"
	"time"

	"v.io/x/ref/profiles/internal/rpc/stream/manager"
	_ "v.io/x/ref/profiles/static"
	"v.io/x/ref/test/benchmark"
	"v.io/x/ref/test/testutil"

	"v.io/v23/naming"
	"v.io/v23/options"
	"v.io/v23/security"
)

// benchmarkDialVC measures VC creation time over the underlying VIF.
func benchmarkDialVC(b *testing.B, mode options.SecurityLevel) {
	stats := benchmark.AddStats(b, 16)

	server := manager.InternalNew(naming.FixedRoutingID(0x5))
	client := manager.InternalNew(naming.FixedRoutingID(0xc))
	var (
		principal security.Principal
		blessings security.Blessings
	)
	if mode == securityTLS {
		principal = testutil.NewPrincipal("test")
		blessings = principal.BlessingStore().Default()
	}

	_, ep, err := server.Listen("tcp", "127.0.0.1:0", principal, blessings)

	if err != nil {
		b.Fatal(err)
	}

	// Warmup to create the underlying VIF.
	_, err = client.Dial(ep, principal)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer() // Exclude setup time from measurement.

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		start := time.Now()

		_, err := client.Dial(ep, principal)
		if err != nil {
			b.Fatal(err)
		}

		duration := time.Since(start)
		b.StopTimer()

		stats.Add(duration)

		client.ShutdownEndpoint(ep)
	}

	client.Shutdown()
	server.Shutdown()
}
