// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationMeshGenerator struct {
  obj gdc.ObjectPtr
}

func createNavigationMeshGenerator(obj gdc.ObjectPtr) *NavigationMeshGenerator {
  return &NavigationMeshGenerator{
    obj: obj,
  }
}

func (me *NavigationMeshGenerator) BaseClass() string {
  return "NavigationMeshGenerator"
}
