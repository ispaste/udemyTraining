package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleZero struct {
	person
	licenseToKill bool
}

func (p person) fullName() string {
	return p.first + p.last
}

// Example of method overriding
// fullName is defined for both Person and DoubleZero
// If the method is called from DoubleZero, it will invoke that method
// If it is called from DoubleZero.Person, then it will invode the Person method

func (dz doubleZero) fullName() string {
	if dz.person.last == "Bond" {
		return "Name's Bond, James Bond"
	} else if dz.first == "Miss" {
		return "Hi! M is waiting for you inside"
	} else {
		return "These are not the droids you're looking for"
	}
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}

	jb := doubleZero{p1, true}
	mp := doubleZero{p2, false}

	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())

	fmt.Println(jb.person)
	fmt.Println(mp.age)

	fmt.Println(jb.licenseToKill)
	fmt.Println(mp.licenseToKill)

	fmt.Println(jb.fullName())
	fmt.Println(mp.fullName())

	fmt.Println(jb)
	fmt.Println(p2)

}
