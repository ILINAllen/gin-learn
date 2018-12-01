package models

import (
	db "mydatabase"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	id, err = db.Engine.Insert(p)
	return
}

func (p *Person) ModPerson() (id int64, err error) {
	id, err = db.Engine.Where("id = ?", p.Id).Update(p)
	return
}

func (p *Person) DelPerson() (id int64, err error) {
	id, err = db.Engine.Where("id = ?", p.Id).Delete(p)
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	person := new(Person)
	rows, err := db.Engine.Rows(person)
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(person)
		if err == nil {
			persons = append(persons, *person)
		}
	}
	return
}

// get person
func (p *Person) GetPerson() (err error) {
	_, err = db.Engine.Where("id = ?", p.Id).Get(p)
	return
}
