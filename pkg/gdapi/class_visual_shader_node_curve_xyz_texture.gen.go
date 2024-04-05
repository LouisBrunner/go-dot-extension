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

type ptrsForVisualShaderNodeCurveXYZTextureList struct {
  fnSetTexture gdc.MethodBindPtr
  fnGetTexture gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeCurveXYZTexture ptrsForVisualShaderNodeCurveXYZTextureList

func initVisualShaderNodeCurveXYZTexturePtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeCurveXYZTexture")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_texture")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCurveXYZTexture.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 8031783))
  }
  {
    methodName := StringNameFromStr("get_texture")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCurveXYZTexture.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1950275015))
  }
}

type VisualShaderNodeCurveXYZTexture struct {
  VisualShaderNodeResizableBase
}

func (me *VisualShaderNodeCurveXYZTexture) BaseClass() string {
  return "VisualShaderNodeCurveXYZTexture"
}

func NewVisualShaderNodeCurveXYZTexture() *VisualShaderNodeCurveXYZTexture {
  str := StringNameFromStr("VisualShaderNodeCurveXYZTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeCurveXYZTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeCurveXYZTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCurveXYZTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCurveXYZTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeCurveXYZTexture) SetTexture(texture CurveXYZTexture, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCurveXYZTexture.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeCurveXYZTexture) GetTexture() CurveXYZTexture {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurveXYZTexture()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCurveXYZTexture.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
