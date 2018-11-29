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
	rs, err := db.Engine.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) ModPerson() (ra int64, err error) {
	rs, err := db.Engine.Exec("UPDATE person SET first_name = ?, last_name = ? WHERE id = ?", p.FirstName, p.LastName, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) DelPerson() (ra int64, err error) {
	rs, err := db.Engine.Exec("DELETE FROM person WHERE id = ?", p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
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
