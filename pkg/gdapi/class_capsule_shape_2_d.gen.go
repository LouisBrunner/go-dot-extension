// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CapsuleShape2D struct {
  obj gdc.ObjectPtr
}

func (me *CapsuleShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CapsuleShape2D) BaseClass() string {
  return "CapsuleShape2D"
}
