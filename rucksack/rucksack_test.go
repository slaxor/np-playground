package rucksack

import (
	"log"
	"testing"
)

// func createLoot(n int) Loot {
//     rng := rand.New(rand.NewSource(time.Now().UnixNano()))
//     l := make(Loot, n)

//     for i := 0; i < n; i++ {
//         l[i] = Item{rng.Float64(), rng.Float64()}
//     }
//     return l
// }
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
	log.Printf("%s", r)
}

func TestCollectQuickAndDirty(t *testing.T) {
	/*
		expect to find:
		 W: 1.994050, V: 5.293350, e: %!s(<nil>), ni: 8
		[
		{W: 0.174530, V: 0.368270, (relV: 2.110067)},
		{W: 0.297420, V: 0.652520, (relV: 2.193935)},
		{W: 0.161680, V: 0.822290, (relV: 5.085910)},
		{W: 0.130070, V: 0.371710, (relV: 2.857769)},
		{W: 0.281050, V: 0.919250, (relV: 3.270770)},
		{W: 0.417220, V: 0.793590, (relV: 1.902090)},
		{W: 0.001760, V: 0.412800, (relV: 234.545455)},
		{W: 0.530320, V: 0.952920, (relV: 1.796877)},]
	*/
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
	log.Printf("%s", r)
}
