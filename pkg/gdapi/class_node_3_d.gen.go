// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Node3D struct {
  obj gdc.ObjectPtr
}

func (me *Node3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Node3D) BaseClass() string {
  return "Node3D"
}

const (
  Node3DNOTIFICATION_TRANSFORM_CHANGED = 2000
  Node3DNOTIFICATION_ENTER_WORLD = 41
  Node3DNOTIFICATION_EXIT_WORLD = 42
  Node3DNOTIFICATION_VISIBILITY_CHANGED = 43
  Node3DNOTIFICATION_LOCAL_TRANSFORM_CHANGED = 44
)

type Node3DRotationEditMode int
const (
  Node3DRotationEditModeEuler Node3DRotationEditMode = 0
  Node3DRotationEditModeQuaternion Node3DRotationEditMode = 1
  Node3DRotationEditModeBasis Node3DRotationEditMode = 2
)
