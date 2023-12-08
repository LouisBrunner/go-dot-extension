// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GDExtension struct {
  obj gdc.ObjectPtr
}

func createGDExtension(obj gdc.ObjectPtr) *GDExtension {
  return &GDExtension{
    obj: obj,
  }
}

func (me *GDExtension) BaseClass() string {
  return "GDExtension"
}
