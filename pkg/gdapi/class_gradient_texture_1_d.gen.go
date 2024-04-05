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

type ptrsForGradientTexture1DList struct {
  fnSetGradient gdc.MethodBindPtr
  fnGetGradient gdc.MethodBindPtr
  fnSetWidth gdc.MethodBindPtr
  fnSetUseHdr gdc.MethodBindPtr
  fnIsUsingHdr gdc.MethodBindPtr
}

var ptrsForGradientTexture1D ptrsForGradientTexture1DList

func initGradientTexture1DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GradientTexture1D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_gradient")
    defer methodName.Destroy()
    ptrsForGradientTexture1D.fnSetGradient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2756054477))
  }
  {
    methodName := StringNameFromStr("get_gradient")
    defer methodName.Destroy()
    ptrsForGradientTexture1D.fnGetGradient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 132272999))
  }
  {
    methodName := StringNameFromStr("set_width")
    defer methodName.Destroy()
    ptrsForGradientTexture1D.fnSetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("set_use_hdr")
    defer methodName.Destroy()
    ptrsForGradientTexture1D.fnSetUseHdr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_using_hdr")
    defer methodName.Destroy()
    ptrsForGradientTexture1D.fnIsUsingHdr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{gradient.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture1D.fnSetGradient), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GradientTexture1D) GetGradient() Gradient {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewGradient()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture1D.fnGetGradient), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GradientTexture1D) SetWidth(width int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture1D.fnSetWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GradientTexture1D) SetUseHdr(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture1D.fnSetUseHdr), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GradientTexture1D) IsUsingHdr() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGradientTexture1D.fnIsUsingHdr), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
