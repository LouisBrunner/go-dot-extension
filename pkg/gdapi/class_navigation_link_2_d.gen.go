// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationLink2D struct {
  obj gdc.ObjectPtr
}

func createNavigationLink2D(obj gdc.ObjectPtr) *NavigationLink2D {
  return &NavigationLink2D{
    obj: obj,
  }
}

func (me *NavigationLink2D) BaseClass() string {
  return "NavigationLink2D"
}
