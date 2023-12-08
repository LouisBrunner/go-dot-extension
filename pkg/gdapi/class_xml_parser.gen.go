// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XMLParser struct {
  obj gdc.ObjectPtr
}

func createXMLParser(obj gdc.ObjectPtr) *XMLParser {
  return &XMLParser{
    obj: obj,
  }
}

func (me *XMLParser) BaseClass() string {
  return "XMLParser"
}
