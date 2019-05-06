package main

import (
	"fmt"

	"github.com/vugu/vugu"
)

type RootData struct {
	Show   bool
	Colors []string
	i      int
}

func (c *Root) NewData(props vugu.Props) (interface{}, error) {
	fmt.Println("in NewData")
	return &RootData{
		Colors: []string{"red", "blue", "green"},
	}, nil
}

func (d *RootData) Toggle() {
	d.Show = !d.Show
}

func (d *RootData) ToggleColor() {
	d.i++
	d.i = d.i % 3 //len(d.Colors)
}
