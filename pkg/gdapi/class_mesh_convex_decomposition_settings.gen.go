// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshConvexDecompositionSettings struct {
  obj gdc.ObjectPtr
}

func createMeshConvexDecompositionSettings(obj gdc.ObjectPtr) *MeshConvexDecompositionSettings {
  return &MeshConvexDecompositionSettings{
    obj: obj,
  }
}

func (me *MeshConvexDecompositionSettings) BaseClass() string {
  return "MeshConvexDecompositionSettings"
}
