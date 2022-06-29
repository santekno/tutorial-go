package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

	name, ok := elements["Ca"]
	if ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println("Tidak terdapat dalam kamus data")
	}

	// 	if name, ok := elements["Un"]; ok {
	// 		fmt.Println(name, ok)
	// }
}
