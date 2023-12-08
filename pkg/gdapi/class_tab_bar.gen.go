// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TabBar struct {
  obj gdc.ObjectPtr
}

func createTabBar(obj gdc.ObjectPtr) *TabBar {
  return &TabBar{
    obj: obj,
  }
}

func (me *TabBar) BaseClass() string {
  return "TabBar"
}
