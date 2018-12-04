package myApis

import (
	"fmt"
	"log"
	"models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works!")
}

//渲染html页面
func ShowHtmlPage(c *gin.Context) {
	c.HTML(http.StatusOK, "list.html", gin.H{
		"title": "GIN: HTML页面",
	})
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	p := models.Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.HTML(http.StatusOK, "add.html", gin.H{
		"data": true,
		"msg":  msg,
	})
}

func ModPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	p := models.Person{Id: int64(id)}

	p.GetPerson()
	c.Header("Content-Type", "text/html; charset=utf-8")
	if p.FirstName != "" {
		p.FirstName = firstName
		p.LastName = lastName
		ra, err := p.ModPerson()
		if err != nil {
			log.Fatalln(err)
		}
		msg := fmt.Sprintf("update successful %d", ra)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
			"msg":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}
}

func DelPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	p := models.Person{Id: int64(id)}

	ra, err := p.DelPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("delete successful %d", ra)
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.JSON(http.StatusOK, gin.H{
		"data": true,
		"msg":  msg,
	})
}

func GetPersonApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	person := models.Person{Id: int64(id)}

	person.GetPerson()
	c.Header("Content-Type", "text/html; charset=utf-8")
	if person.FirstName != "" {
		c.JSON(http.StatusOK, gin.H{
			"data": person,
			"msg":  "success",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": nil,
			"msg":  "Person not found",
		})
	}

}

func GetPersonsApi(c *gin.Context) {
	var p models.Person
	persons, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": persons,
		"msg":  "success",
	})
}
