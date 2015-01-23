package impl_test

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"testing"

	"v.io/core/veyron2"
	"v.io/core/veyron2/naming"
	"v.io/core/veyron2/services/mgmt/repository"

	"v.io/core/veyron/lib/testutil"
	"v.io/core/veyron/services/mgmt/binary/impl"
)

// TestHTTP checks that HTTP download works.
func TestHTTP(t *testing.T) {
	ctx, shutdown := veyron2.Init()
	defer shutdown()
	veyron2.GetNamespace(ctx).CacheCtl(naming.DisableCache(true))

	// TODO(caprita): This is based on TestMultiPart (impl_test.go).  Share
	// the code where possible.
	for length := 2; length < 5; length++ {
		binary, _, url, cleanup := startServer(t, ctx, 2)
		defer cleanup()
		// Create <length> chunks of up to 4MB of random bytes.
		data := make([][]byte, length)
		for i := 0; i < length; i++ {
			// Random size, but at least 1 (avoid empty parts).
			size := testutil.Rand.Intn(1000*impl.BufferLength) + 1
			data[i] = testutil.RandomBytes(size)
		}
		mediaInfo := repository.MediaInfo{Type: "application/octet-stream"}
		if err := binary.Create(ctx, int32(length), mediaInfo); err != nil {
			t.Fatalf("Create() failed: %v", err)
		}
		for i := 0; i < length; i++ {
			if streamErr, err := invokeUpload(t, ctx, binary, data[i], int32(i)); streamErr != nil || err != nil {
				t.FailNow()
			}
		}
		parts, _, err := binary.Stat(ctx)
		if err != nil {
			t.Fatalf("Stat() failed: %v", err)
		}
		response, err := http.Get(url)
		if err != nil {
			t.Fatal(err)
		}
		downloaded, err := ioutil.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}
		from, to := 0, 0
		for i := 0; i < length; i++ {
			hpart := md5.New()
			to += len(data[i])
			if ld := len(downloaded); to > ld {
				t.Fatalf("Download falls short: len(downloaded):%d, need:%d (i:%d, length:%d)", ld, to, i, length)
			}
			output := downloaded[from:to]
			from = to
			if bytes.Compare(output, data[i]) != 0 {
				t.Fatalf("Unexpected output: expected %v, got %v", data[i], output)
			}
			hpart.Write(data[i])
			checksum := hex.EncodeToString(hpart.Sum(nil))
			if expected, got := checksum, parts[i].Checksum; expected != got {
				t.Fatalf("Unexpected checksum: expected %v, got %v", expected, got)
			}
			if expected, got := len(data[i]), int(parts[i].Size); expected != got {
				t.Fatalf("Unexpected size: expected %v, got %v", expected, got)
			}
		}
		if err := binary.Delete(ctx); err != nil {
			t.Fatalf("Delete() failed: %v", err)
		}
	}
}
