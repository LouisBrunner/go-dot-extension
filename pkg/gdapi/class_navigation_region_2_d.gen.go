// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationRegion2D struct {
  obj gdc.ObjectPtr
}

func createNavigationRegion2D(obj gdc.ObjectPtr) *NavigationRegion2D {
  return &NavigationRegion2D{
    obj: obj,
  }
}

func (me *NavigationRegion2D) BaseClass() string {
  return "NavigationRegion2D"
}
