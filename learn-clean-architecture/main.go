package main

import (
	"log"

	"github.com/irayspace/go-legos/learn-clean-architecture/app/employee"
)

func main() {
	h := employee.InitHandler()

	e, err := h.Create(employee.Employee{
		FirstName: "Ivan",
		LastName:  "Ivanov",
	})
	if err != nil {
		log.Println("unable to create Employee")
	}
	log.Println(e)

	employees := h.GetAll()
	log.Println(employees)

	if err := h.DeleteByID(e.ID); err != nil {
		log.Println("unable to remove Employee")
	}

	employees = h.GetAll()
	log.Println(employees)
}
