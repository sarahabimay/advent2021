package main
import "testing"

func TestMain(t *testing.T) {
	act, exp := countIncreases("123\n344"), 1
	if act != exp {
		t.Fatalf("exp: %d, got %d", exp, act)
	}
}