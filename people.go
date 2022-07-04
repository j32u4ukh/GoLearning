package main

import "GoLearning/structs"

type People struct {
	person1 structs.Person
	person2 structs.Person
}

func (people People) Speak(content string) {
	people.person1.Speak(content)
	people.person2.Speak(content)
}
