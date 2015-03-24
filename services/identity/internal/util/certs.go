package util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// WriteCertAndKey creates a certificate and private key for a given host and
// duration and writes them to cert.pem and key.pem in tmpdir.  It returns the
// locations of the files, or an error if one is encountered.
func WriteCertAndKey(host string, duration time.Duration) (string, string, error) {
	listCmd := exec.Command("go", "list", "-f", "{{.Dir}}", "crypto/tls")
	output, err := listCmd.Output()
	if err != nil {
		return "", "", fmt.Errorf("%s failed: %v", strings.Join(listCmd.Args, " "), err)
	}
	tmpDir := os.TempDir()
	generateCertFile := filepath.Join(strings.TrimSpace(string(output)), "generate_cert.go")
	generateCertCmd := exec.Command("go", "run", generateCertFile, "--host", host, "--duration", duration.String())
	generateCertCmd.Dir = tmpDir
	if err := generateCertCmd.Run(); err != nil {
		return "", "", fmt.Errorf("Could not generate key and cert: %v", err)
	}
	return filepath.Join(tmpDir, "cert.pem"), filepath.Join(tmpDir, "key.pem"), nil
}