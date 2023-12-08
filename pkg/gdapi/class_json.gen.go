// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JSON struct {
  obj gdc.ObjectPtr
}

func createJSON(obj gdc.ObjectPtr) *JSON {
  return &JSON{
    obj: obj,
  }
}

func (me *JSON) BaseClass() string {
  return "JSON"
}
