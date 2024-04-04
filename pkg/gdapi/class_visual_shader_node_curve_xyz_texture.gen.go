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
  classNameV := StringNameFromStr("VisualShaderNodeCurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 8031783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeCurveXYZTexture) GetTexture() CurveXYZTexture {
  classNameV := StringNameFromStr("VisualShaderNodeCurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1950275015) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurveXYZTexture()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
