// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RefCounted struct {
  obj gdc.ObjectPtr
}

func createRefCounted(obj gdc.ObjectPtr) *RefCounted {
  return &RefCounted{
    obj: obj,
  }
}

func (me *RefCounted) BaseClass() string {
  return "RefCounted"
}
