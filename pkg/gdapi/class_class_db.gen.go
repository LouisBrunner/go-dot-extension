// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ClassDB struct {
  obj gdc.ObjectPtr
}

func (me *ClassDB) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ClassDB) BaseClass() string {
  return "ClassDB"
}
