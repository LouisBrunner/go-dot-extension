// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CenterContainer struct {
  obj gdc.ObjectPtr
}

func createCenterContainer(obj gdc.ObjectPtr) *CenterContainer {
  return &CenterContainer{
    obj: obj,
  }
}

func (me *CenterContainer) BaseClass() string {
  return "CenterContainer"
}
