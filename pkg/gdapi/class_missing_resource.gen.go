// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MissingResource struct {
  obj gdc.ObjectPtr
}

func createMissingResource(obj gdc.ObjectPtr) *MissingResource {
  return &MissingResource{
    obj: obj,
  }
}

func (me *MissingResource) BaseClass() string {
  return "MissingResource"
}
