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
	ok := IsValid("password_text", "aes256v1$MwvsyhRqCLStfOa/$74e7a82f69ccb88df6b6169d8d310c84bb7d049dd013f4d6336c41752f972b8b3dfee4c3ee2bb5bcb91d7abfca756200c2195bdd70b4b7089c539158675757d884114074c949f132c7bd2a35", "AES256Key-32Characters1234567890")
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
