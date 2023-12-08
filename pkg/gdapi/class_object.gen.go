// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Object struct {
  obj gdc.ObjectPtr
}

func (me *Object) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Object) BaseClass() string {
  return "Object"
}
