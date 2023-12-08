// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshConvexDecompositionSettings struct {
  obj gdc.ObjectPtr
}

func (me *MeshConvexDecompositionSettings) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MeshConvexDecompositionSettings) BaseClass() string {
  return "MeshConvexDecompositionSettings"
}

type MeshConvexDecompositionSettingsMode int
const (
  MeshConvexDecompositionSettingsConvexDecompositionModeVoxel MeshConvexDecompositionSettingsMode = 0
  MeshConvexDecompositionSettingsConvexDecompositionModeTetrahedron MeshConvexDecompositionSettingsMode = 1
)
