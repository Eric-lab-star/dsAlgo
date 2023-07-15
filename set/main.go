package main

import "fmt"

type Set struct {
	intMap map[int]bool
}

func (set *Set) Contain(element int) bool {
	_, exists := set.intMap[element]
	return exists
}

func (set *Set) Add(element int) {
	if !set.Contain(element) {
		set.intMap[element] = true
	}
}

func (set *Set) Delete(element int) {
	delete(set.intMap, element)
}

func (set *Set) Intersect(anotherSet *Set) *Set {
	intersectSet := &Set{intMap: map[int]bool{}}
	for value := range set.intMap {
		if anotherSet.Contain(value) {
			intersectSet.Add(value)
		}
	}
	return intersectSet
}

func (set *Set) Union(anotherSet *Set) *Set {
	unionSet := &Set{intMap: map[int]bool{}}
	for value := range set.intMap {
		unionSet.Add(value)
	}

	for value := range anotherSet.intMap {
		unionSet.Add(value)
	}
	return unionSet
}

func (set *Set) AddAll(elements ...int) {
	for _, element := range elements {
		set.Add(element)
	}
}

func main() {
	set1 := &Set{intMap: make(map[int]bool)}
	set1.AddAll(2, 4, 6, 8)

	set2 := &Set{intMap: make(map[int]bool)}
	set2.AddAll(3, 6, 9, 12)

	intersect := set1.Intersect(set2)
	fmt.Println(intersect)

}
