// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGShape3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGShape3D) BaseClass() string {
  return "CSGShape3D"
}

type CSGShape3DOperation int
const (
  CSGShape3DOperationUnion CSGShape3DOperation = 0
  CSGShape3DOperationIntersection CSGShape3DOperation = 1
  CSGShape3DOperationSubtraction CSGShape3DOperation = 2
)
