package rucksack

import (
	"math/rand"
	"testing"
	"time"
)

func createLoot(n int) Loot {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	l := make(Loot, n)

	for i := 0; i < n; i++ {
		l[i] = Item{rng.Float64(), rng.Float64()}
	}
	return l
}

var testLoot = Loot{
	{0.17453, 0.36827}, {0.37638, 0.06312}, {0.73349, 0.87194},
	{0.91805, 0.13192}, {0.75290, 0.06987}, {0.85971, 0.80406},
	{0.62478, 0.47000}, {0.29742, 0.65252}, {0.96285, 0.39023},
	{0.16168, 0.82229}, {0.72147, 0.73047}, {0.13007, 0.37171},
	{0.89845, 0.26445}, {0.28105, 0.91925}, {0.36442, 0.59823},
	{0.84517, 0.69147}, {0.56973, 0.42526}, {0.41722, 0.79359},
	{0.22531, 0.40366}, {0.52396, 0.48262}, {0.00176, 0.41280},
	{0.74864, 0.03195}, {0.53032, 0.95292}, {0.63777, 0.64209},
	{0.79919, 0.04608},
}

func TestRucksackWeight(t *testing.T) {
	tt := []struct {
		r Rucksack
		v float64
		e bool //expect error
	}{
		{Rucksack{Loot: testLoot, maxWeight: 99}, 13.556320000000001, false},
		{Rucksack{Loot: testLoot, maxWeight: 9}, 13.556320000000001, true},
	}
	for _, tCase := range tt {
		v, err := tCase.r.Weight()
		if (err != nil) != tCase.e {
			t.Fatalf("%s is not expected", err)
		}
		if v != tCase.v {
			t.Fatalf("%v is not the expected weight (%v)", v, tCase.v)
		}
	}
}

func TestRucksackValue(t *testing.T) {
	tt := []struct {
		r Rucksack
		v float64
	}{
		{Rucksack{Loot: testLoot, maxWeight: 99}, 12.410770000000001},
	}
	for _, tCase := range tt {
		v := tCase.r.Value()
		if v != tCase.v {
			t.Fatalf("%v is not the expected value (%v)", v, tCase.v)
		}
	}
}

func TestCollect(t *testing.T) {
	l := testLoot[:10]
	expect := Rucksack{
		Loot: Loot{
			l[0],
			l[2],
			l[6],
			l[7],
			l[9],
		},
		maxWeight: 2.0,
	}
	ew, err := expect.Weight()
	if err != nil {
		t.Fatalf("%s", err)
	}
	ev := expect.Value() // 3.185020

	r := Collect(l, 2.0)
	w, err := r.Weight()
	if err != nil {
		t.Fatalf("%s", err)
	}
	v := r.Value()
	if v < ev {
		t.Fatalf("%f is below %f, (W: %f EW: %f)", v, ev, w, ew)
	}
}

func TestCollectQuickAndDirty(t *testing.T) {
	l := testLoot
	expect := Rucksack{
		Loot: Loot{
			l[0],
			l[2],
			l[6],
			l[7],
			l[9],
		},
		maxWeight: 2.0,
	}
	ew, err := expect.Weight()
	if err != nil {
		t.Fatalf("%s", err)
	}
	ev := expect.Value() // 3.185020

	r := CollectQuickAndDirty(l, 2.0)
	w, err := r.Weight()
	if err != nil {
		t.Fatalf("%s", err)
	}
	v := r.Value()
	if v < ev {
		t.Fatalf("%f is below %f, (W: %f EW: %f)", v, ev, w, ew)
	}
}

func BenchmarkCollectQuickAndDirty4(b *testing.B) {
	l := createLoot(4)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CollectQuickAndDirty(l, 2)
	}
}

func BenchmarkCollectQuickAndDirty8(b *testing.B) {
	l := createLoot(8)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CollectQuickAndDirty(l, 2)
	}
}

func BenchmarkCollectQuickAndDirty16(b *testing.B) {
	l := createLoot(16)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CollectQuickAndDirty(l, 2)
	}
}

func BenchmarkCollect4(b *testing.B) {
	l := createLoot(4)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Collect(l, 2)
	}
}

func BenchmarkCollect8(b *testing.B) {
	l := createLoot(8)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Collect(l, 2)
	}
}

func BenchmarkCollect16(b *testing.B) {
	l := createLoot(16)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Collect(l, 2)
	}
}
