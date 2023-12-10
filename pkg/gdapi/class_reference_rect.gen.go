// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ReferenceRect struct {
  obj gdc.ObjectPtr
}

func (me *ReferenceRect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ReferenceRect) BaseClass() string {
  return "ReferenceRect"
}



// Enums

func (me *ReferenceRect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ReferenceRect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ReferenceRect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ReferenceRect) GetBorderColor() Color {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReferenceRect) SetBorderColor(color Color, )  {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReferenceRect) GetBorderWidth() float32 {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_border_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReferenceRect) SetBorderWidth(width float32, )  {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ReferenceRect) GetEditorOnly() bool {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ReferenceRect) SetEditorOnly(enabled bool, )  {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *ReferenceRect) GetPropBorderColor() Color {
  panic("TODO: implement")
}

func (me *ReferenceRect) SetPropBorderColor(value Color) {
  panic("TODO: implement")
}

func (me *ReferenceRect) GetPropBorderWidth() float32 {
  panic("TODO: implement")
}

func (me *ReferenceRect) SetPropBorderWidth(value float32) {
  panic("TODO: implement")
}

func (me *ReferenceRect) GetPropEditorOnly() bool {
  panic("TODO: implement")
}

func (me *ReferenceRect) SetPropEditorOnly(value bool) {
  panic("TODO: implement")
}