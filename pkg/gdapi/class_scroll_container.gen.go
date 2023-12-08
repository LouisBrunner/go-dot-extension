// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScrollContainer struct {
  obj gdc.ObjectPtr
}

func createScrollContainer(obj gdc.ObjectPtr) *ScrollContainer {
  return &ScrollContainer{
    obj: obj,
  }
}

func (me *ScrollContainer) BaseClass() string {
  return "ScrollContainer"
}
