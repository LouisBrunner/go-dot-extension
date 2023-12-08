// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationMeshSourceGeometryData3D struct {
  obj gdc.ObjectPtr
}

func createNavigationMeshSourceGeometryData3D(obj gdc.ObjectPtr) *NavigationMeshSourceGeometryData3D {
  return &NavigationMeshSourceGeometryData3D{
    obj: obj,
  }
}

func (me *NavigationMeshSourceGeometryData3D) BaseClass() string {
  return "NavigationMeshSourceGeometryData3D"
}
