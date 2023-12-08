// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CapsuleMesh struct {
  obj gdc.ObjectPtr
}

func (me *CapsuleMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CapsuleMesh) BaseClass() string {
  return "CapsuleMesh"
}
