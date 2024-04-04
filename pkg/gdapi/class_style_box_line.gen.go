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

type StyleBoxLine struct {
  StyleBox
}

func (me *StyleBoxLine) BaseClass() string {
  return "StyleBoxLine"
}

func NewStyleBoxLine() *StyleBoxLine {
  str := StringNameFromStr("StyleBoxLine") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StyleBoxLine{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *StyleBoxLine) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StyleBoxLine) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StyleBoxLine) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *StyleBoxLine) SetColor(color Color, )  {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) GetColor() Color {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxLine) SetThickness(thickness int64, )  {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&thickness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) GetThickness() int64 {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxLine) SetGrowBegin(offset float64, )  {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_grow_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) GetGrowBegin() float64 {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_grow_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxLine) SetGrowEnd(offset float64, )  {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_grow_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) GetGrowEnd() float64 {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_grow_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxLine) SetVertical(vertical bool, )  {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) IsVertical() bool {
  classNameV := StringNameFromStr("StyleBoxLine")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_vertical")
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
