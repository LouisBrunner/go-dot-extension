// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VBoxContainer struct {
  obj gdc.ObjectPtr
}

func createVBoxContainer(obj gdc.ObjectPtr) *VBoxContainer {
  return &VBoxContainer{
    obj: obj,
  }
}

func (me *VBoxContainer) BaseClass() string {
  return "VBoxContainer"
}
