// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ClassDB struct {
  obj gdc.ObjectPtr
}

func createClassDB(obj gdc.ObjectPtr) *ClassDB {
  return &ClassDB{
    obj: obj,
  }
}

func (me *ClassDB) BaseClass() string {
  return "ClassDB"
}
