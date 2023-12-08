// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceSaver struct {
  obj gdc.ObjectPtr
}

func createResourceSaver(obj gdc.ObjectPtr) *ResourceSaver {
  return &ResourceSaver{
    obj: obj,
  }
}

func (me *ResourceSaver) BaseClass() string {
  return "ResourceSaver"
}
