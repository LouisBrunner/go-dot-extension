// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PanelContainer struct {
  obj gdc.ObjectPtr
}

func createPanelContainer(obj gdc.ObjectPtr) *PanelContainer {
  return &PanelContainer{
    obj: obj,
  }
}

func (me *PanelContainer) BaseClass() string {
  return "PanelContainer"
}
