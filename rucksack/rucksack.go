package rucksack

import (
	"fmt"
	"log"
	"sort"
)

type Item struct {
	Weight, Value float64
}

func (i Item) String() string {
	return fmt.Sprintf("\n{W: %f, V: %f, (relV: %f)},",
		i.Weight, i.Value, i.Rel())
}

func (i Item) Rel() float64 {
	return i.Value / i.Weight
}

type Loot []Item

func (l Loot) Weight() float64 {
	w := 0.0
	for _, item := range l {
		w += item.Weight
	}
	return w
}

func (l Loot) Value() float64 {
	v := 0.0
	for _, item := range l {
		v += item.Value
	}
	return v
}

type Rucksack struct {
	Loot      Loot
	maxWeight float64
}

func (r Rucksack) String() string {
	w, err := r.Weight()
	v := r.Value()
	l := len(r.Loot)
	return fmt.Sprintf("W: %f, V: %f, e: %s, ni: %d\n%s",
		w, v, err, l, r.Loot)
}

func (r *Rucksack) Weight() (float64, error) {
	w := 0.0
	for _, item := range r.Loot {
		w += item.Weight
	}
	if w > r.maxWeight {
		return w, fmt.Errorf(
			"Your rucksack will break because %f is too "+
				"much weight for it. It can only handle %f",
			w, r.maxWeight)
	}
	return w, nil
}

func (r *Rucksack) Value() float64 {
	v := 0.0
	for _, item := range r.Loot {
		v += item.Value
	}
	return v
}

// Collect takes a Loot and a load maximum and tries to pack the
// greatest value equal or lower than the max load.
// This is a naïve approach.
func Collect(l Loot, c float64) Rucksack {
	r := Rucksack{maxWeight: c}
	e := 1 << uint(len(l))
	for i := 0; i < e; i++ {
		tr := Rucksack{maxWeight: c}
		for j, item := range l {
			tf := i >> uint(j) % 2
			if tf > 0 {
				tr.Loot = append(tr.Loot, item)
			}
		}
		_, err := tr.Weight()
		if err != nil {
			// log.Print("too heavy")
			continue
		}
		v := tr.Value()
		if v > r.Value() {
			r.Loot = tr.Loot
			// log.Printf("Found a better one W: %v, V: %v", w, v)
		}
	}
	return r
}

// CollectRealworld takes a Loot and a load maximum and tries to pack
// the greatest value equal or lower than the max load.
// It'll not be mathematical correct but should render a good result
// in a pretty short time
func CollectQuickAndDirty(l Loot, c float64) Rucksack {
	r := Rucksack{maxWeight: c}
	sort.Slice(l, func(i, j int) bool { return l[i].Rel() > l[j].Rel() })
	log.Print(l)
	for i := 0; i < len(l); i++ {
		if l[:i].Weight() > c {
			break
		}
		r.Loot = l[:i]
	}
	return r
}
