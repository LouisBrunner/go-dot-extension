// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeParticleMultiplyByAxisAngle struct {
  VisualShaderNode
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) BaseClass() string {
  return "VisualShaderNodeParticleMultiplyByAxisAngle"
}

func NewVisualShaderNodeParticleMultiplyByAxisAngle() *VisualShaderNodeParticleMultiplyByAxisAngle {
  str := StringNameFromStr("VisualShaderNodeParticleMultiplyByAxisAngle") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeParticleMultiplyByAxisAngle{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParticleMultiplyByAxisAngle) SetDegreesMode(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMultiplyByAxisAngle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_degrees_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeParticleMultiplyByAxisAngle) IsDegreesMode() bool {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMultiplyByAxisAngle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_degrees_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
