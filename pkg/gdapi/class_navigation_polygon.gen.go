// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationPolygon struct {
  obj gdc.ObjectPtr
}

func createNavigationPolygon(obj gdc.ObjectPtr) *NavigationPolygon {
  return &NavigationPolygon{
    obj: obj,
  }
}

func (me *NavigationPolygon) BaseClass() string {
  return "NavigationPolygon"
}
