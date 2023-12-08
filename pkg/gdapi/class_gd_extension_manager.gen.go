// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GDExtensionManager struct {
  obj gdc.ObjectPtr
}

func createGDExtensionManager(obj gdc.ObjectPtr) *GDExtensionManager {
  return &GDExtensionManager{
    obj: obj,
  }
}

func (me *GDExtensionManager) BaseClass() string {
  return "GDExtensionManager"
}
