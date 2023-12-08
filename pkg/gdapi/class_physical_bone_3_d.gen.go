// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicalBone3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicalBone3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicalBone3D) BaseClass() string {
  return "PhysicalBone3D"
}

type PhysicalBone3DDampMode int
const (
  PhysicalBone3DDampModeCombine PhysicalBone3DDampMode = 0
  PhysicalBone3DDampModeReplace PhysicalBone3DDampMode = 1
)

type PhysicalBone3DJointType int
const (
  PhysicalBone3DJointTypeNone PhysicalBone3DJointType = 0
  PhysicalBone3DJointTypePin PhysicalBone3DJointType = 1
  PhysicalBone3DJointTypeCone PhysicalBone3DJointType = 2
  PhysicalBone3DJointTypeHinge PhysicalBone3DJointType = 3
  PhysicalBone3DJointTypeSlider PhysicalBone3DJointType = 4
  PhysicalBone3DJointType6Dof PhysicalBone3DJointType = 5
)
