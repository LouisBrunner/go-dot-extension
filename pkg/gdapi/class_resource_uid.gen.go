// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceUID struct {
  obj gdc.ObjectPtr
}

func createResourceUID(obj gdc.ObjectPtr) *ResourceUID {
  return &ResourceUID{
    obj: obj,
  }
}

func (me *ResourceUID) BaseClass() string {
  return "ResourceUID"
}
