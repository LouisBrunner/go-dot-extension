// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CapsuleShape3D struct {
  obj gdc.ObjectPtr
}

func (me *CapsuleShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CapsuleShape3D) BaseClass() string {
  return "CapsuleShape3D"
}
