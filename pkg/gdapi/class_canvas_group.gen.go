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

type CanvasGroup struct {
  Node2D
}

func (me *CanvasGroup) BaseClass() string {
  return "CanvasGroup"
}

func NewCanvasGroup() *CanvasGroup {
  str := StringNameFromStr("CanvasGroup") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CanvasGroup{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CanvasGroup) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CanvasGroup) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CanvasGroup) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CanvasGroup) SetFitMargin(fit_margin float64, )  {
  classNameV := StringNameFromStr("CanvasGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fit_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fit_margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasGroup) GetFitMargin() float64 {
  classNameV := StringNameFromStr("CanvasGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fit_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasGroup) SetClearMargin(clear_margin float64, )  {
  classNameV := StringNameFromStr("CanvasGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clear_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clear_margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasGroup) GetClearMargin() float64 {
  classNameV := StringNameFromStr("CanvasGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clear_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasGroup) SetUseMipmaps(use_mipmaps bool, )  {
  classNameV := StringNameFromStr("CanvasGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_mipmaps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasGroup) IsUsingMipmaps() bool {
  classNameV := StringNameFromStr("CanvasGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_mipmaps")
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
