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

type ptrsForColorRectList struct {
  fnSetColor gdc.MethodBindPtr
  fnGetColor gdc.MethodBindPtr
}

var ptrsForColorRect ptrsForColorRectList

func initColorRectPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ColorRect")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_color")
    defer methodName.Destroy()
    ptrsForColorRect.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_color")
    defer methodName.Destroy()
    ptrsForColorRect.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
}

type ColorRect struct {
  Control
}

func (me *ColorRect) BaseClass() string {
  return "ColorRect"
}

func NewColorRect() *ColorRect {
  str := StringNameFromStr("ColorRect") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ColorRect{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ColorRect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ColorRect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ColorRect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ColorRect) SetColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorRect.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ColorRect) GetColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForColorRect.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
