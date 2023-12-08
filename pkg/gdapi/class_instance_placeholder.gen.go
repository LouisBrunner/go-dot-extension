// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InstancePlaceholder struct {
  obj gdc.ObjectPtr
}

func createInstancePlaceholder(obj gdc.ObjectPtr) *InstancePlaceholder {
  return &InstancePlaceholder{
    obj: obj,
  }
}

func (me *InstancePlaceholder) BaseClass() string {
  return "InstancePlaceholder"
}
