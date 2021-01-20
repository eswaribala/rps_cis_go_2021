package main

import (
	"fmt"
	"sort"
)
type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	s := []int{400, 20, 3000, 1}
	sort.Ints(s)
	fmt.Println(s) // [1 2 3
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Emman", 2},
		{"Banu", 25},
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Name < family[j].Name
	})
	fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]


	personFamily := []Person{
		{"Alice", 23},
		{"Eve", 2},
		{"Bob", 25},
	}
	sort.Sort(ByAge(personFamily))
	fmt.Println(personFamily) // [{Eve 2} {Alice 23} {Bob 25}]


	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

}
