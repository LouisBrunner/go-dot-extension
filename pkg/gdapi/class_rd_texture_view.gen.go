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

type ptrsForRDTextureViewList struct {
  fnSetFormatOverride gdc.MethodBindPtr
  fnGetFormatOverride gdc.MethodBindPtr
  fnSetSwizzleR gdc.MethodBindPtr
  fnGetSwizzleR gdc.MethodBindPtr
  fnSetSwizzleG gdc.MethodBindPtr
  fnGetSwizzleG gdc.MethodBindPtr
  fnSetSwizzleB gdc.MethodBindPtr
  fnGetSwizzleB gdc.MethodBindPtr
  fnSetSwizzleA gdc.MethodBindPtr
  fnGetSwizzleA gdc.MethodBindPtr
}

var ptrsForRDTextureView ptrsForRDTextureViewList

func initRDTextureViewPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RDTextureView")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_format_override")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnSetFormatOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 565531219))
  }
  {
    methodName := StringNameFromStr("get_format_override")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnGetFormatOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2235804183))
  }
  {
    methodName := StringNameFromStr("set_swizzle_r")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnSetSwizzleR = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3833362581))
  }
  {
    methodName := StringNameFromStr("get_swizzle_r")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnGetSwizzleR = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4150792614))
  }
  {
    methodName := StringNameFromStr("set_swizzle_g")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnSetSwizzleG = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3833362581))
  }
  {
    methodName := StringNameFromStr("get_swizzle_g")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnGetSwizzleG = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4150792614))
  }
  {
    methodName := StringNameFromStr("set_swizzle_b")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnSetSwizzleB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3833362581))
  }
  {
    methodName := StringNameFromStr("get_swizzle_b")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnGetSwizzleB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4150792614))
  }
  {
    methodName := StringNameFromStr("set_swizzle_a")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnSetSwizzleA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3833362581))
  }
  {
    methodName := StringNameFromStr("get_swizzle_a")
    defer methodName.Destroy()
    ptrsForRDTextureView.fnGetSwizzleA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4150792614))
  }
}

type RDTextureView struct {
  RefCounted
}

func (me *RDTextureView) BaseClass() string {
  return "RDTextureView"
}

func NewRDTextureView() *RDTextureView {
  str := StringNameFromStr("RDTextureView") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDTextureView{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RDTextureView) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDTextureView) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDTextureView) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDTextureView) SetFormatOverride(p_member RenderingDeviceDataFormat, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnSetFormatOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetFormatOverride() RenderingDeviceDataFormat {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceDataFormat

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnGetFormatOverride), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureView) SetSwizzleR(p_member RenderingDeviceTextureSwizzle, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnSetSwizzleR), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetSwizzleR() RenderingDeviceTextureSwizzle {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureSwizzle

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnGetSwizzleR), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureView) SetSwizzleG(p_member RenderingDeviceTextureSwizzle, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnSetSwizzleG), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetSwizzleG() RenderingDeviceTextureSwizzle {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureSwizzle

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnGetSwizzleG), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureView) SetSwizzleB(p_member RenderingDeviceTextureSwizzle, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnSetSwizzleB), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetSwizzleB() RenderingDeviceTextureSwizzle {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureSwizzle

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnGetSwizzleB), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureView) SetSwizzleA(p_member RenderingDeviceTextureSwizzle, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnSetSwizzleA), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetSwizzleA() RenderingDeviceTextureSwizzle {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureSwizzle

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDTextureView.fnGetSwizzleA), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
