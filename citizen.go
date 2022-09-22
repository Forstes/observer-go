package main

import "fmt"

type Citizen struct {
	id   string
	name string
}

func newCitizen(id string, name string) *Citizen {
	return &Citizen{
		id:   id,
		name: name,
	}
}

func (c *Citizen) update(details string) {
	fmt.Printf("Отправляем повестку гражданину %s %s. Адрес и время прибытия: %s\n", c.id, c.name, details)
}

func (c *Citizen) getID() string {
	return c.id
}
