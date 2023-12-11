// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimatableBody3D struct {
  obj gdc.ObjectPtr
}

func (me *AnimatableBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatableBody3D) BaseClass() string {
  return "AnimatableBody3D"
}



// Enums

func (me *AnimatableBody3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimatableBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimatableBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimatableBody3D) SetSyncToPhysics(enable bool, )  {
  classNameV := StringNameFromStr("AnimatableBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sync_to_physics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimatableBody3D) IsSyncToPhysicsEnabled() bool {
  classNameV := StringNameFromStr("AnimatableBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_sync_to_physics_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
