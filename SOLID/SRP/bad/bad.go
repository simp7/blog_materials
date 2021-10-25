package main

import (
	"fmt"
	"log"
	"os"
)

type Item struct {
	name  string
	cost  int
	stock int
}

func (i Item) Info() {
	fmt.Printf("----------\nname: %s\ncost: %d\nstock: %d\n----------", i.name, i.cost, i.stock)
}

func (i Item) Save() error {
	f, err := os.Create(i.name + ".txt")
	if err != nil {
		return err
	}
	_, err = f.WriteString(fmt.Sprintf("cost: %d\nstock: %d\n", i.cost, i.stock))
	return err
}

func main() {
	i := Item{name: "Mug", cost: 8000, stock: 17}
	i.Info()
	if err := i.Save(); err != nil {
		log.Fatal(err)
	}
}
