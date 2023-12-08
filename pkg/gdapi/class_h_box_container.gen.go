// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HBoxContainer struct {
  obj gdc.ObjectPtr
}

func createHBoxContainer(obj gdc.ObjectPtr) *HBoxContainer {
  return &HBoxContainer{
    obj: obj,
  }
}

func (me *HBoxContainer) BaseClass() string {
  return "HBoxContainer"
}
