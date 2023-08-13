package functions

import (
	"encoding/xml"
)

// 定义 XML 数据的结构体
type StringTable struct {
	XMLName   xml.Name `xml:"StringTable"`
	Name      string   `xml:"Name,attr"`
	Namespace string   `xml:"Namespace,attr"`
	DataList  []Data   `xml:"Data"`
}

type Data struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value"`
}

type Translation struct {
	XMLName      xml.Name      `xml:"Translation"`
	StringTables []StringTable `xml:"StringTable"`
}
