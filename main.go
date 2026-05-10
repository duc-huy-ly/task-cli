package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dly/task-cli/core"
)

type Hero struct {
	Id int
	Name string
	Beers int
}

func main() {
	fmt.Println("Hello app")
	t := core.NewTask("first task")
	println(t.ToString())
	data, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	himmer := Hero{
		Id : 1000,
		Name: "I love frieren",
		Beers: 677,
	}
	himData, _ := json.Marshal(himmer)
	fmt.Println(string(himData))

}



