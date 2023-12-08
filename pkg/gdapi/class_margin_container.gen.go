// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MarginContainer struct {
  obj gdc.ObjectPtr
}

func createMarginContainer(obj gdc.ObjectPtr) *MarginContainer {
  return &MarginContainer{
    obj: obj,
  }
}

func (me *MarginContainer) BaseClass() string {
  return "MarginContainer"
}
