// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XMLParser struct {
  obj gdc.ObjectPtr
}

func (me *XMLParser) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XMLParser) BaseClass() string {
  return "XMLParser"
}



// Enums

type XMLParserNodeType int
const (
  XMLParserNodeTypeNodeNone XMLParserNodeType = 0
  XMLParserNodeTypeNodeElement XMLParserNodeType = 1
  XMLParserNodeTypeNodeElementEnd XMLParserNodeType = 2
  XMLParserNodeTypeNodeText XMLParserNodeType = 3
  XMLParserNodeTypeNodeComment XMLParserNodeType = 4
  XMLParserNodeTypeNodeCdata XMLParserNodeType = 5
  XMLParserNodeTypeNodeUnknown XMLParserNodeType = 6
)

func (me *XMLParser) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XMLParser) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *XMLParser) Read()  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetNodeType()  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetNodeName()  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetNodeData()  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetNodeOffset()  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetAttributeCount()  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetAttributeName(idx int, )  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetAttributeValue(idx int, )  {
  panic("TODO: implement")
}

func  (me *XMLParser) HasAttribute(name String, )  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetNamedAttributeValue(name String, )  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetNamedAttributeValueSafe(name String, )  {
  panic("TODO: implement")
}

func  (me *XMLParser) IsEmpty()  {
  panic("TODO: implement")
}

func  (me *XMLParser) GetCurrentLine()  {
  panic("TODO: implement")
}

func  (me *XMLParser) SkipSection()  {
  panic("TODO: implement")
}

func  (me *XMLParser) Seek(position int, )  {
  panic("TODO: implement")
}

func  (me *XMLParser) Open(file String, )  {
  panic("TODO: implement")
}

func  (me *XMLParser) OpenBuffer(buffer PackedByteArray, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
