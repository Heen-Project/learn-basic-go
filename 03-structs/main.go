package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// value type: int, float, string, bool, structs
// referenes type: slices, maps, channels, pointer, function
func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Rothschild",
		contact: contactInfo{
			email:   "alex@example.com",
			zipCode: 012011,
		},
	}
	// -- 1. old method
	// alexPointer := &alex
	// alexPointer.updateName("Alexander")
	// -- 2. go shortcut (automatically simplify the process)
	alex.updateName("Alexander")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Pointer
// var x, y int
// fmt.Println(x == 0, x == y)  // "true true"
// fmt.Println(&x == &y)        // "false"
// fmt.Printf("%T %T\n", x, &x) // "int *int"

// var p *int
// fmt.Println(p == nil) // "true"

// p = &x
// fmt.Println(p == &x, x == *p) // "true true"
// fmt.Printf("%T %T\n", p, &p)  // "*int **int"

// vertex := struct{ X, Y float64 }{1, 2}
// vptr := &vertex
// fmt.Println(vptr.X == (*vptr).X)       // "true"
// fmt.Println(&(vptr.X) == &((*vptr).X)) // "true"
