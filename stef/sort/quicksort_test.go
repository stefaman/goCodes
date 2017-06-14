package sort

import (
	"math/rand"
	"testing"
	"time"
)

func genRand(rng *rand.Rand, n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rng.Int()
	}
	return s
}

//模拟重复元素很多的情况
func genRept(n int) []int {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = int(rng.Int31n(20))
	}
	return s
}

func genSeq(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}
func genInv(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = n - i
	}
	return s
}
func genSame(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = 333
	}
	return s
}

func benchmarkSortRand(b *testing.B, n int) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	s := genRand(rng, n)
	a := make([]int, n)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(a, s)
		b.StartTimer()
		QuickSort(a)
	}
}
func benchmarkSortRept(b *testing.B, n int) {
	s := genRept(n)
	a := make([]int, n)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(a, s)
		b.StartTimer()
		QuickSort(a)
	}
}
func benchmarkSortSeq(b *testing.B, n int) {
	s := genSeq(n)
	a := make([]int, n)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(a, s)
		b.StartTimer()
		QuickSort(a)
	}
}
func benchmarkSortInv(b *testing.B, n int) {
	s := genInv(n)
	a := make([]int, n)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(a, s)
		b.StartTimer()
		QuickSort(a)
	}
}
func benchmarkSortSame(b *testing.B, n int) {
	s := genSame(n)
	for i := 0; i < b.N; i++ {
		QuickSort(s)
	}
}
func BenchmarkSortSeqE6(b *testing.B) {
	benchmarkSortSeq(b, 1e6)
}
func BenchmarkSortInvE6(b *testing.B) {
	benchmarkSortInv(b, 1e6)
}
func BenchmarkSortSameE6(b *testing.B) {
	benchmarkSortSame(b, 1e6)
}
func BenchmarkSortReptE6(b *testing.B) {
	benchmarkSortRept(b, 1e6)
}
func BenchmarkSortRandE4(b *testing.B) {
	benchmarkSortRand(b, 1e4)
}
func BenchmarkSortRandE5(b *testing.B) {
	benchmarkSortRand(b, 1e5)
}
func BenchmarkSortRandE6(b *testing.B) {
	benchmarkSortRand(b, 1e6)
}

func TestSortRand(t *testing.T) {
	const n = 100000
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	s := genRand(rng, n)
	QuickSort(s)
	if !IsSorted(s) {
		t.Error("for seed", seed, "IsSorted false")
	}
}
func testSorted(t *testing.T, s []int) {
	QuickSort(s)
	if !IsSorted(s) {
		t.Error("IsSorted false")
	}
}
func TestSortSpecial(t *testing.T) {
	const n = 100000
	testSorted(t, genSeq(n))
	testSorted(t, genInv(n))
	testSorted(t, genSame(n))
	testSorted(t, []int{})
	testSorted(t, []int{1})
	testSorted(t, []int{1, 2})
	testSorted(t, []int{2, 1})
}
