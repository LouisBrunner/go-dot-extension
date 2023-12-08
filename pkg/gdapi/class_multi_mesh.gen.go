// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiMesh struct {
  obj gdc.ObjectPtr
}

func (me *MultiMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiMesh) BaseClass() string {
  return "MultiMesh"
}

type MultiMeshTransformFormat int
const (
  MultiMeshTransform2D MultiMeshTransformFormat = 0
  MultiMeshTransform3D MultiMeshTransformFormat = 1
)
