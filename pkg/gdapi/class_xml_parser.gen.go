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

func  (me *XMLParser) Read() { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetNodeType() { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetNodeName() { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetNodeData() { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetNodeOffset() { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetAttributeCount() { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetAttributeName(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetAttributeValue(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) HasAttribute(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetNamedAttributeValue(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetNamedAttributeValueSafe(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) IsEmpty() { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) GetCurrentLine() { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) SkipSection() { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) Seek(position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) Open(file String, ) { // TODO: return value
  // TODO: implement
}

func  (me *XMLParser) OpenBuffer(buffer PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
