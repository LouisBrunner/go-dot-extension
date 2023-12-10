// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GradientTexture1D struct {
  obj gdc.ObjectPtr
}

func (me *GradientTexture1D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GradientTexture1D) BaseClass() string {
  return "GradientTexture1D"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(gradient.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture1D) GetGradient() Gradient {
  classNameV := StringNameFromStr("GradientTexture1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gradient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  var ret Gradient
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GradientTexture1D) SetWidth(width int, )  {
  classNameV := StringNameFromStr("GradientTexture1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture1D) SetUseHdr(enabled bool, )  {
  classNameV := StringNameFromStr("GradientTexture1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_hdr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture1D) IsUsingHdr() bool {
  classNameV := StringNameFromStr("GradientTexture1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_hdr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *GradientTexture1D) GetPropGradient() Gradient {
  panic("TODO: implement")
}

func (me *GradientTexture1D) SetPropGradient(value Gradient) {
  panic("TODO: implement")
}

func (me *GradientTexture1D) GetPropWidth() int {
  panic("TODO: implement")
}

func (me *GradientTexture1D) SetPropWidth(value int) {
  panic("TODO: implement")
}

func (me *GradientTexture1D) GetPropUseHdr() bool {
  panic("TODO: implement")
}

func (me *GradientTexture1D) SetPropUseHdr(value bool) {
  panic("TODO: implement")
}