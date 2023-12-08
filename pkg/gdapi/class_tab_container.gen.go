// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TabContainer struct {
  obj gdc.ObjectPtr
}

func createTabContainer(obj gdc.ObjectPtr) *TabContainer {
  return &TabContainer{
    obj: obj,
  }
}

func (me *TabContainer) BaseClass() string {
  return "TabContainer"
}
