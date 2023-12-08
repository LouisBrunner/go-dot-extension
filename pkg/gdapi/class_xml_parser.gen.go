// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  XMLParserNodeNone XMLParserNodeType = 0
  XMLParserNodeElement XMLParserNodeType = 1
  XMLParserNodeElementEnd XMLParserNodeType = 2
  XMLParserNodeText XMLParserNodeType = 3
  XMLParserNodeComment XMLParserNodeType = 4
  XMLParserNodeCdata XMLParserNodeType = 5
  XMLParserNodeUnknown XMLParserNodeType = 6
)
