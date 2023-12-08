// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoxContainer struct {
  obj gdc.ObjectPtr
}

func createBoxContainer(obj gdc.ObjectPtr) *BoxContainer {
  return &BoxContainer{
    obj: obj,
  }
}

func (me *BoxContainer) BaseClass() string {
  return "BoxContainer"
}
