// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WeakRef struct {
  obj gdc.ObjectPtr
}

func createWeakRef(obj gdc.ObjectPtr) *WeakRef {
  return &WeakRef{
    obj: obj,
  }
}

func (me *WeakRef) BaseClass() string {
  return "WeakRef"
}
