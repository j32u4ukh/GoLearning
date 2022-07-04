package structs

import (
	"fmt"
)

func (person *Person) Speak(content string) {
	fmt.Println(person.Name + ": " + content)
}

func (ojisan *Ojisan) Speak(content string) {
	fmt.Println(ojisan.Name + ": " + content)
}

func (man *Man) Speak(content string) {
	fmt.Println(man.Ojisan.Name + ": " + content)
}
