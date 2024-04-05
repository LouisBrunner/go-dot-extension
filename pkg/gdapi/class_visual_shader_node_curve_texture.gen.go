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

type ptrsForVisualShaderNodeCurveTextureList struct {
  fnSetTexture gdc.MethodBindPtr
  fnGetTexture gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeCurveTexture ptrsForVisualShaderNodeCurveTextureList

func initVisualShaderNodeCurveTexturePtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeCurveTexture")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_texture")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCurveTexture.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 181872837))
  }
  {
    methodName := StringNameFromStr("get_texture")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCurveTexture.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2800800579))
  }
}

type VisualShaderNodeCurveTexture struct {
  VisualShaderNodeResizableBase
}

func (me *VisualShaderNodeCurveTexture) BaseClass() string {
  return "VisualShaderNodeCurveTexture"
}

func NewVisualShaderNodeCurveTexture() *VisualShaderNodeCurveTexture {
  str := StringNameFromStr("VisualShaderNodeCurveTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeCurveTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeCurveTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCurveTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCurveTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeCurveTexture) SetTexture(texture CurveTexture, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCurveTexture.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeCurveTexture) GetTexture() CurveTexture {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurveTexture()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCurveTexture.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
