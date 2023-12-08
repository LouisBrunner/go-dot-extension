// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationServer2D struct {
  obj gdc.ObjectPtr
}

func createNavigationServer2D(obj gdc.ObjectPtr) *NavigationServer2D {
  return &NavigationServer2D{
    obj: obj,
  }
}

func (me *NavigationServer2D) BaseClass() string {
  return "NavigationServer2D"
}
