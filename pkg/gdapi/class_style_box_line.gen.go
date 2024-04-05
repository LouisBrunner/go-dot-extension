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

type ptrsForStyleBoxLineList struct {
  fnSetColor gdc.MethodBindPtr
  fnGetColor gdc.MethodBindPtr
  fnSetThickness gdc.MethodBindPtr
  fnGetThickness gdc.MethodBindPtr
  fnSetGrowBegin gdc.MethodBindPtr
  fnGetGrowBegin gdc.MethodBindPtr
  fnSetGrowEnd gdc.MethodBindPtr
  fnGetGrowEnd gdc.MethodBindPtr
  fnSetVertical gdc.MethodBindPtr
  fnIsVertical gdc.MethodBindPtr
}

var ptrsForStyleBoxLine ptrsForStyleBoxLineList

func initStyleBoxLinePtrs(iface gdc.Interface) {

  className := StringNameFromStr("StyleBoxLine")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_color")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_color")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
  {
    methodName := StringNameFromStr("set_thickness")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnSetThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_thickness")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnGetThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_grow_begin")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnSetGrowBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_grow_begin")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnGetGrowBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_grow_end")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnSetGrowEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_grow_end")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnGetGrowEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_vertical")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnSetVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_vertical")
    defer methodName.Destroy()
    ptrsForStyleBoxLine.fnIsVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) GetColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StyleBoxLine) SetThickness(thickness int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&thickness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnSetThickness), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) GetThickness() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnGetThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxLine) SetGrowBegin(offset float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnSetGrowBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) GetGrowBegin() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnGetGrowBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxLine) SetGrowEnd(offset float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnSetGrowEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) GetGrowEnd() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnGetGrowEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StyleBoxLine) SetVertical(vertical bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnSetVertical), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StyleBoxLine) IsVertical() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBoxLine.fnIsVertical), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
