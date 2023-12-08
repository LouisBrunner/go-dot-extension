// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ORMMaterial3D struct {
  obj gdc.ObjectPtr
}

func (me *ORMMaterial3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ORMMaterial3D) BaseClass() string {
  return "ORMMaterial3D"
}
