package main

import (
	"fmt"
	"sort"
)

type Orang struct {
	Nama string
	Umur int
}
type ByNama []Orang

func (this ByNama) Len() int {
	return len(this)
}
func (this ByNama) Less(i, j int) bool {
	return this[i].Nama < this[j].Nama
}
func (this ByNama) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByUmur []Orang

func (this ByUmur) Len() int {
	return len(this)
}
func (this ByUmur) Less(i, j int) bool {
	return this[i].Umur < this[j].Umur
}
func (this ByUmur) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Orang{
		{"Jill", 9},
		{"Jack", 10},
	}
	sort.Sort(ByNama(kids))
	fmt.Println(kids)

	sort.Sort(ByUmur(kids))
	fmt.Println(kids)
}
