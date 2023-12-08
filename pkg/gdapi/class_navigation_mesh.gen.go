// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationMesh struct {
  obj gdc.ObjectPtr
}

func createNavigationMesh(obj gdc.ObjectPtr) *NavigationMesh {
  return &NavigationMesh{
    obj: obj,
  }
}

func (me *NavigationMesh) BaseClass() string {
  return "NavigationMesh"
}
