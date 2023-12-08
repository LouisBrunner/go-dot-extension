// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Resource struct {
  obj gdc.ObjectPtr
}

func createResource(obj gdc.ObjectPtr) *Resource {
  return &Resource{
    obj: obj,
  }
}

func (me *Resource) BaseClass() string {
  return "Resource"
}
