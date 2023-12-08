// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Object struct {
  obj gdc.ObjectPtr
}

func createObject(obj gdc.ObjectPtr) *Object {
  return &Object{
    obj: obj,
  }
}

func (me *Object) BaseClass() string {
  return "Object"
}
