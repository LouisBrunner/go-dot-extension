// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationAgent2D struct {
  obj gdc.ObjectPtr
}

func createNavigationAgent2D(obj gdc.ObjectPtr) *NavigationAgent2D {
  return &NavigationAgent2D{
    obj: obj,
  }
}

func (me *NavigationAgent2D) BaseClass() string {
  return "NavigationAgent2D"
}
