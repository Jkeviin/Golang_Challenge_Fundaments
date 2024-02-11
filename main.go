package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func printPersonInfo(person Person) {
	fmt.Printf("Nombre: %s, Edad: %d, Email: %s\n", person.Name, person.Age, person.Email)
}

func validatePerson(person Person) bool {
	return person.Name != "" && person.Age > 0 && person.Email != ""
}

func main() {
	people := []Person{
		{Name: "Alice", Age: 30, Email: "alice@example.com"},
		{Name: "Bob", Age: 25, Email: "bob@example.com"},
		{Name: "Charlie", Age: 22, Email: "charlie@example.com"},
	}

	peopleMap := make(map[string]*Person)
	for i := range people {
		peopleMap[people[i].Name] = &people[i]
	}

	var wg sync.WaitGroup
	wg.Add(len(people))

	for _, p := range people {
		go func(person Person) {
			defer wg.Done()
			if validatePerson(person) {
				printPersonInfo(person)
			} else {
				fmt.Printf("Persona inválida: %s\n", person.Name)
			}
		}(p)
	}
	wg.Wait()
}
