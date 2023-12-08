// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PathFollow3D struct {
  obj gdc.ObjectPtr
}

func (me *PathFollow3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PathFollow3D) BaseClass() string {
  return "PathFollow3D"
}

type PathFollow3DRotationMode int
const (
  PathFollow3DRotationNone PathFollow3DRotationMode = 0
  PathFollow3DRotationY PathFollow3DRotationMode = 1
  PathFollow3DRotationXy PathFollow3DRotationMode = 2
  PathFollow3DRotationXyz PathFollow3DRotationMode = 3
  PathFollow3DRotationOriented PathFollow3DRotationMode = 4
)
