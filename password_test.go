package password

import (
	"fmt"
	"testing"
)

// Benchmark with "go test -bench=."
func TestHash(t *testing.T) {
	hash, err := Hash("password_text", "AES256Key-32Characters1234567890")
	if err != nil {
		t.Errorf("err should be nil, got %v", err)
	}

	fmt.Printf("%s\n", hash)
}

func TestIsValid(t *testing.T) {
	ok := IsValid("password_text", "aes256v1$yu448C5P0t6WcFSg$4d4022aa5c231f14abcbacd8705314fd41d157e0d3461ee3b7c8dd0ca5a51e107b29659d4b360f98ed8375d4f1a5ba0fa05127e3879e5790eca07e597fc803ae7d450d73e2e01adb3f7d20fe", "AES256Key-32Characters1234567890")
	if !ok {
		t.Errorf("should be true got false")
	}
}

func BenchmarkHash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := Hash("password_text", "AES256Key-32Characters1234567890")
		if err != nil {
			b.Errorf("err should be nil, got %v", err)
		}
	}
}
