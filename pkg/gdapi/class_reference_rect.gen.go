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

type ReferenceRect struct {
  Control
}

func (me *ReferenceRect) BaseClass() string {
  return "ReferenceRect"
}

func NewReferenceRect() *ReferenceRect {
  str := StringNameFromStr("ReferenceRect") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ReferenceRect{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ReferenceRect) SetBorderColor(color Color, )  {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReferenceRect) GetBorderWidth() float64 {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_border_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReferenceRect) SetBorderWidth(width float64, )  {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_border_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ReferenceRect) GetEditorOnly() bool {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ReferenceRect) SetEditorOnly(enabled bool, )  {
  classNameV := StringNameFromStr("ReferenceRect")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editor_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
