package password_test

import (
	"fmt"
	"hetz-client/internal/auth/password"
	"os"
	"testing"
)

func TestHash(t *testing.T) {
	os.Setenv("SALT", "saltItUpFam")
	hash, err := password.Hash("examplePassword", 12)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hash)
}

func BenchmarkHash(b *testing.B) {
	os.Setenv("SALT", "saltItUpFam")
	for i := 0; i < b.N; i++ {
		_, _ = password.Hash("examplePassword", 11)
	}
}

func benchmarkHashCostRange(b *testing.B) {
	os.Setenv("SALT", "saltItUpFam")
	for _, cost := range []int{4, 8, 10, 12, 13, 14, 15} {
		b.Run(fmt.Sprintf("cost=%d", cost), func(b *testing.B) {
			benchmarkHashWithCost(b, cost)
		})
	}
}

func benchmarkHashWithCost(b *testing.B, cost int) {
	pass := "examplePassword"
	for i := 0; i < b.N; i++ {
		_, err := password.Hash(pass, cost)
		if err != nil {
			b.Fatal(err)
		}
	}
}
