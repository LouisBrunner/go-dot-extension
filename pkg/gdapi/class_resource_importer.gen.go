// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ResourceImporter struct {
  obj gdc.ObjectPtr
}

func createResourceImporter(obj gdc.ObjectPtr) *ResourceImporter {
  return &ResourceImporter{
    obj: obj,
  }
}

func (me *ResourceImporter) BaseClass() string {
  return "ResourceImporter"
}
