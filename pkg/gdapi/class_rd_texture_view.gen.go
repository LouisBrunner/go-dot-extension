// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDTextureView struct {
  obj gdc.ObjectPtr
}

func (me *RDTextureView) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDTextureView) BaseClass() string {
  return "RDTextureView"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDTextureView) GetFormatOverride() RenderingDeviceDataFormat {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2235804183) // FIXME: should cache?
  var ret RenderingDeviceDataFormat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDTextureView) SetSwizzleR(p_member RenderingDeviceTextureSwizzle, )  {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_swizzle_r")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3833362581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDTextureView) GetSwizzleR() RenderingDeviceTextureSwizzle {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swizzle_r")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4150792614) // FIXME: should cache?
  var ret RenderingDeviceTextureSwizzle
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDTextureView) SetSwizzleG(p_member RenderingDeviceTextureSwizzle, )  {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_swizzle_g")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3833362581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDTextureView) GetSwizzleG() RenderingDeviceTextureSwizzle {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swizzle_g")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4150792614) // FIXME: should cache?
  var ret RenderingDeviceTextureSwizzle
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDTextureView) SetSwizzleB(p_member RenderingDeviceTextureSwizzle, )  {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_swizzle_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3833362581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDTextureView) GetSwizzleB() RenderingDeviceTextureSwizzle {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swizzle_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4150792614) // FIXME: should cache?
  var ret RenderingDeviceTextureSwizzle
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDTextureView) SetSwizzleA(p_member RenderingDeviceTextureSwizzle, )  {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_swizzle_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3833362581) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDTextureView) GetSwizzleA() RenderingDeviceTextureSwizzle {
  classNameV := StringNameFromStr("RDTextureView")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swizzle_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4150792614) // FIXME: should cache?
  var ret RenderingDeviceTextureSwizzle
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *RDTextureView) GetPropFormatOverride() int {
  panic("TODO: implement")
}

func (me *RDTextureView) SetPropFormatOverride(value int) {
  panic("TODO: implement")
}

func (me *RDTextureView) GetPropSwizzleR() int {
  panic("TODO: implement")
}

func (me *RDTextureView) SetPropSwizzleR(value int) {
  panic("TODO: implement")
}

func (me *RDTextureView) GetPropSwizzleG() int {
  panic("TODO: implement")
}

func (me *RDTextureView) SetPropSwizzleG(value int) {
  panic("TODO: implement")
}

func (me *RDTextureView) GetPropSwizzleB() int {
  panic("TODO: implement")
}

func (me *RDTextureView) SetPropSwizzleB(value int) {
  panic("TODO: implement")
}

func (me *RDTextureView) GetPropSwizzleA() int {
  panic("TODO: implement")
}

func (me *RDTextureView) SetPropSwizzleA(value int) {
  panic("TODO: implement")
}