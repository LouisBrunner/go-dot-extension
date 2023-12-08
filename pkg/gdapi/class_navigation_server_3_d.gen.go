// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationServer3D struct {
  obj gdc.ObjectPtr
}

func createNavigationServer3D(obj gdc.ObjectPtr) *NavigationServer3D {
  return &NavigationServer3D{
    obj: obj,
  }
}

func (me *NavigationServer3D) BaseClass() string {
  return "NavigationServer3D"
}
