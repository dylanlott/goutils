package main

import (
	"fmt"

	"github.com/dylanlott/goutils/Pluralsite/webservice/models"
)

func main() {
	u := models.Users{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
