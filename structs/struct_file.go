package structs

type Person struct {
	Name string
}

type Ojisan struct {
	Name string
}

type Man struct {
	Person
	Ojisan
}
