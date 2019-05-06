package main

import "github.com/vugu/vugu"

type MyLineData struct {
	FileName   string
	LineNumber int
}

func (d *MyLine) NewData(props vugu.Props) (interface{}, error) {
	return &MyLineData{
		FileName:   props["file-name"].(string),
		LineNumber: props["line-number"].(int),
	}, nil
}
