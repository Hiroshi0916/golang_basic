package main

import (
	"fmt"
	"testing/todo_app/app/controllers"
	"testing/todo_app/app/models"
)

func TestConnection() {

}

func main() {

	fmt.Println(models.Db)

	controllers.StartMainServer()

}
