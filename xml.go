package main

import (
	"encoding/xml"
	"fmt"
)

// Plant 结构体
type Plant struct {
	// 规定了该结构的XML元素名称
	XMLName xml.Name `xml:"plant"`

	Id     int      `xml:"id,attr"` //id,attrr 表示ID是一个XML属性，不是嵌套元素
	Name   string   `xml:"name"`
	Origin []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v,name=%v,origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ") // 生成XML。第二、三个参数是缩进
	fmt.Println(string(out))
	fmt.Println("-----------------\n")
	fmt.Println(xml.Header + string(out)) // 加上 XML 格式头

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil { // 解析 XML 格式字节流
		panic(err)
	}
	fmt.Println(p)
	fmt.Println("------------------\n")

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "CaliFornia"}

	type Nseting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
		// `xml:"parent>child>plant"` 的 Tab 将 Plant 元素嵌套在<parent><child>里面
	}

	nesting := &Nseting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
