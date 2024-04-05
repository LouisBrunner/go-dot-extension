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

type ptrsForVisualShaderNodeParticleEmitterList struct {
  fnSetMode2D gdc.MethodBindPtr
  fnIsMode2D gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeParticleEmitter ptrsForVisualShaderNodeParticleEmitterList

func initVisualShaderNodeParticleEmitterPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeParticleEmitter")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_mode_2d")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeParticleEmitter.fnSetMode2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_mode_2d")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeParticleEmitter.fnIsMode2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type VisualShaderNodeParticleEmitter struct {
  VisualShaderNode
}

func (me *VisualShaderNodeParticleEmitter) BaseClass() string {
  return "VisualShaderNodeParticleEmitter"
}

func NewVisualShaderNodeParticleEmitter() *VisualShaderNodeParticleEmitter {
  str := StringNameFromStr("VisualShaderNodeParticleEmitter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeParticleEmitter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeParticleEmitter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleEmitter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleEmitter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParticleEmitter) SetMode2D(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleEmitter.fnSetMode2D), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeParticleEmitter) IsMode2D() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeParticleEmitter.fnIsMode2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
