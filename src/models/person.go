package models

import (
	"fmt"
	db "myDatabase"
)

type Person struct {
	Id        int64  `json:"id" form:"id"` // 只要保证名称是Id,且类型为int64, 则xorm会自动认为主键并自增。
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	id, err = db.Engine.Insert(p)
	fmt.Println(p)
	return
}

func (p *Person) ModPerson() (id int64, err error) {
	id, err = db.Engine.Where("id = ?", p.Id).Update(p)
	fmt.Println(p)
	return
}

func (p *Person) DelPerson() (id int64, err error) {
	id, err = db.Engine.Where("id = ?", p.Id).Delete(p)
	fmt.Println(p)
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
	fmt.Println(p)
	return
}
