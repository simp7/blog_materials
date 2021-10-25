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

func (i Item) State() string {
	return fmt.Sprintf("cost: %d\nstock: %d\n", i.cost, i.stock)
}

func (i Item) String() string {
	return fmt.Sprintf("----------\nname: %s\n%s----------", i.name, i.State())
}

type fileDB struct {
}

func (db *fileDB) Save(item Item) error {
	f, err := os.Create(item.name + ".txt")
	if err != nil {
		return err
	}
	_, err = f.WriteString(item.State())
	return err
}

func main() {
	i := Item{name: "Mug", cost: 8000, stock: 17}
	db := new(fileDB)
	fmt.Println(i)
	if err := db.Save(i); err != nil {
		log.Fatal(err)
	}
}
