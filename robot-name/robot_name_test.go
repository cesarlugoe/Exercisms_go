// +build !bonus

package robotname

import (
	"testing"
)

func TestNameValid(t *testing.T) {
	n := New().getName(t, false)
	if !namePat.MatchString(n) {
		t.Errorf(`Invalid robot name %q, want form "AA###".`, n)
	}
}

func TestGetRandomNumber(t *testing.T) {
	random := getRandomNumber(10)
	if random > 10 {
		t.Fatalf("Random number %d out of bounce", random)
	}
}

func TestNameSticks(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	n2 := r.getName(t, true)
	if n2 != n1 {
		t.Errorf(`Robot name changed.  Now %s, was %s.`, n2, n1)
	}
}

func TestSuccessiveRobotsHaveDifferentNames(t *testing.T) {
	n1 := New().getName(t, false)
	n2 := New().getName(t, false)
	if n1 == n2 {
		t.Errorf(`Robots with same name.  Two %s's.`, n1)
	}
}

func TestResetName(t *testing.T) {
	r := New()
	n1 := r.getName(t, false)
	r.Reset()
	if r.getName(t, false) == n1 {
		t.Errorf(`Robot name not cleared on reset.  Still %s.`, n1)
	}
}

func BenchmarkName(b *testing.B) {
	// Benchmark combined time to create robot and name.
	for i := 0; i < b.N; i++ {
		New().getName(b, false)
	}
}
