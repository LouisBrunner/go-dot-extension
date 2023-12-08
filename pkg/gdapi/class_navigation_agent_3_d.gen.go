// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationAgent3D struct {
  obj gdc.ObjectPtr
}

func createNavigationAgent3D(obj gdc.ObjectPtr) *NavigationAgent3D {
  return &NavigationAgent3D{
    obj: obj,
  }
}

func (me *NavigationAgent3D) BaseClass() string {
  return "NavigationAgent3D"
}
