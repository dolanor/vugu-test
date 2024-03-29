package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "github.com/vugu/vugu"

type MyLineData struct {
	FileName   string
	LineNumber int
}

func (*MyLine) NewData(props vugu.Props) (interface{}, error) {
	return &MyLineData{
		FileName:   props["file-name"].(string),
		LineNumber: props["line-number"].(int),
	}, nil
}

var _ vugu.ComponentType = (*MyLine)(nil)

func (comp *MyLine) BuildVDOM(dataI interface{}) (vdom *vugu.VGNode, css *vugu.VGNode, reterr error) {
	data := dataI.(*MyLineData)
	_ = data
	_ = fmt.Sprint
	_ = reflect.Value{}
	event := vugu.DOMEventStub
	_ = event
	var n *vugu.VGNode
	n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "li", DataAtom: vugu.VGAtom(45570), Namespace: "", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "my-line"}}}
	vdom = n
	{
		parent := n
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    ", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "strong", DataAtom: vugu.VGAtom(449798), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n.InnerHTML = fmt.Sprint(data.FileName)
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: ":", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "span", DataAtom: vugu.VGAtom(40708), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
		n.InnerHTML = fmt.Sprint(data.LineNumber)
		n = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n", DataAtom: vugu.VGAtom(0), Namespace: "", Attr: []vugu.VGAttribute(nil)}
		parent.AppendChild(n)
	}
	return
}

type MyLine struct {}

func init() { vugu.RegisterComponentType("my-line", &MyLine{}) }
