// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CapsuleShape2D struct {
  obj gdc.ObjectPtr
}

func createCapsuleShape2D(obj gdc.ObjectPtr) *CapsuleShape2D {
  return &CapsuleShape2D{
    obj: obj,
  }
}

func (me *CapsuleShape2D) BaseClass() string {
  return "CapsuleShape2D"
}
