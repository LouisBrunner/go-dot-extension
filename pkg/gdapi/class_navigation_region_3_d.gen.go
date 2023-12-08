// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationRegion3D struct {
  obj gdc.ObjectPtr
}

func createNavigationRegion3D(obj gdc.ObjectPtr) *NavigationRegion3D {
  return &NavigationRegion3D{
    obj: obj,
  }
}

func (me *NavigationRegion3D) BaseClass() string {
  return "NavigationRegion3D"
}
