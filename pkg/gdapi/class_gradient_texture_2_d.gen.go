// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GradientTexture2D struct {
  Texture2D
}

func (me *GradientTexture2D) BaseClass() string {
  return "GradientTexture2D"
}



// Enums

type GradientTexture2DFill int
const (
  GradientTexture2DFillFillLinear GradientTexture2DFill = 0
  GradientTexture2DFillFillRadial GradientTexture2DFill = 1
  GradientTexture2DFillFillSquare GradientTexture2DFill = 2
)

type GradientTexture2DRepeat int
const (
  GradientTexture2DRepeatRepeatNone GradientTexture2DRepeat = 0
  GradientTexture2DRepeatRepeat GradientTexture2DRepeat = 1
  GradientTexture2DRepeatRepeatMirror GradientTexture2DRepeat = 2
)

func (me *GradientTexture2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GradientTexture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GradientTexture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GradientTexture2D) SetGradient(gradient Gradient, )  {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gradient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2756054477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(gradient.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture2D) GetGradient() Gradient {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gradient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 132272999) // FIXME: should cache?
  var ret Gradient
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GradientTexture2D) SetWidth(width int, )  {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture2D) SetHeight(height int, )  {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture2D) SetUseHdr(enabled bool, )  {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_hdr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture2D) IsUsingHdr() bool {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_hdr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GradientTexture2D) SetFill(fill GradientTexture2DFill, )  {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fill")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3623927636) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fill), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture2D) GetFill() GradientTexture2DFill {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fill")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1876227217) // FIXME: should cache?
  var ret GradientTexture2DFill
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GradientTexture2D) SetFillFrom(fill_from Vector2, )  {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fill_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(fill_from.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture2D) GetFillFrom() Vector2 {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fill_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GradientTexture2D) SetFillTo(fill_to Vector2, )  {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fill_to")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(fill_to.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture2D) GetFillTo() Vector2 {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fill_to")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GradientTexture2D) SetRepeat(repeat GradientTexture2DRepeat, )  {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1357597002) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&repeat), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GradientTexture2D) GetRepeat() GradientTexture2DRepeat {
  classNameV := StringNameFromStr("GradientTexture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3351758665) // FIXME: should cache?
  var ret GradientTexture2DRepeat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
