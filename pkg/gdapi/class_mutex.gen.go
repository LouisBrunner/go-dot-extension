// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Mutex struct {
  obj gdc.ObjectPtr
}

func (me *Mutex) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Mutex) BaseClass() string {
  return "Mutex"
}
