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
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_format_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 565531219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetFormatOverride() RenderingDeviceDataFormat {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2235804183) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceDataFormat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureView) SetSwizzleR(p_member RenderingDeviceTextureSwizzle, )  {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_swizzle_r")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3833362581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetSwizzleR() RenderingDeviceTextureSwizzle {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swizzle_r")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4150792614) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureSwizzle

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureView) SetSwizzleG(p_member RenderingDeviceTextureSwizzle, )  {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_swizzle_g")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3833362581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetSwizzleG() RenderingDeviceTextureSwizzle {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swizzle_g")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4150792614) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureSwizzle

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureView) SetSwizzleB(p_member RenderingDeviceTextureSwizzle, )  {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_swizzle_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3833362581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetSwizzleB() RenderingDeviceTextureSwizzle {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swizzle_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4150792614) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureSwizzle

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureView) SetSwizzleA(p_member RenderingDeviceTextureSwizzle, )  {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_swizzle_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3833362581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureView) GetSwizzleA() RenderingDeviceTextureSwizzle {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swizzle_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4150792614) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureSwizzle

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
