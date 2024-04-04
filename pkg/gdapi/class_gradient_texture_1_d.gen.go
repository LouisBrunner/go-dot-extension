// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type GradientTexture1D struct {
  Texture2D
}

func (me *GradientTexture1D) BaseClass() string {
  return "GradientTexture1D"
}

func NewGradientTexture1D() *GradientTexture1D {
  str := StringNameFromStr("GradientTexture1D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GradientTexture1D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GradientTexture1D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GradientTexture1D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GradientTexture1D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GradientTexture1D) SetGradient(gradient Gradient, )  {
  classNameV := StringNameFromStr("GradientTexture1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gradient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gradient.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GradientTexture1D) GetGradient() Gradient {
  classNameV := StringNameFromStr("GradientTexture1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gradient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGradient()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GradientTexture1D) SetWidth(width int64, )  {
  classNameV := StringNameFromStr("GradientTexture1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GradientTexture1D) SetUseHdr(enabled bool, )  {
  classNameV := StringNameFromStr("GradientTexture1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_hdr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GradientTexture1D) IsUsingHdr() bool {
  classNameV := StringNameFromStr("GradientTexture1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_hdr")
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
