// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RigidBody3D struct {
  obj gdc.ObjectPtr
}

func (me *RigidBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RigidBody3D) BaseClass() string {
  return "RigidBody3D"
}

type RigidBody3DFreezeMode int
const (
  RigidBody3DFreezeModeStatic RigidBody3DFreezeMode = 0
  RigidBody3DFreezeModeKinematic RigidBody3DFreezeMode = 1
)

type RigidBody3DCenterOfMassMode int
const (
  RigidBody3DCenterOfMassModeAuto RigidBody3DCenterOfMassMode = 0
  RigidBody3DCenterOfMassModeCustom RigidBody3DCenterOfMassMode = 1
)

type RigidBody3DDampMode int
const (
  RigidBody3DDampModeCombine RigidBody3DDampMode = 0
  RigidBody3DDampModeReplace RigidBody3DDampMode = 1
)
