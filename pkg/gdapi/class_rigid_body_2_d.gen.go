// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RigidBody2D struct {
  obj gdc.ObjectPtr
}

func (me *RigidBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RigidBody2D) BaseClass() string {
  return "RigidBody2D"
}

type RigidBody2DFreezeMode int
const (
  RigidBody2DFreezeModeStatic RigidBody2DFreezeMode = 0
  RigidBody2DFreezeModeKinematic RigidBody2DFreezeMode = 1
)

type RigidBody2DCenterOfMassMode int
const (
  RigidBody2DCenterOfMassModeAuto RigidBody2DCenterOfMassMode = 0
  RigidBody2DCenterOfMassModeCustom RigidBody2DCenterOfMassMode = 1
)

type RigidBody2DDampMode int
const (
  RigidBody2DDampModeCombine RigidBody2DDampMode = 0
  RigidBody2DDampModeReplace RigidBody2DDampMode = 1
)

type RigidBody2DCCDMode int
const (
  RigidBody2DCcdModeDisabled RigidBody2DCCDMode = 0
  RigidBody2DCcdModeCastRay RigidBody2DCCDMode = 1
  RigidBody2DCcdModeCastShape RigidBody2DCCDMode = 2
)
