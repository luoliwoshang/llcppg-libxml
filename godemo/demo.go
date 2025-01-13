package main

import (
	"libxml_2_0"
	"unsafe"

	"github.com/goplus/llgo/c"
)

func main() {
	libxml_2_0.XmlInitParser()
	xml := "<?xml version='1.0'?><root><person><name>Alice</name><age>25</age></person></root>"
	doc := libxml_2_0.XmlReadMemory((*int8)(unsafe.Pointer(unsafe.StringData(xml))), c.Int(len(xml)), nil, nil, 0)
	if doc == nil {
		panic("Failed to parse XML")
	}
	docPtr := (*libxml_2_0.XmlDoc)(unsafe.Pointer(doc))
	root := docPtr.XmlDocGetRootElement()
	c.Printf(c.Str("Root element: %s\n"), root.Name)
	libxml_2_0.XmlFreeDoc(doc)
	libxml_2_0.XmlCleanupParser()
}
