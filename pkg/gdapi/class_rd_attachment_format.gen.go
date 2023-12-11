// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDAttachmentFormat struct {
  obj gdc.ObjectPtr
}

func (me *RDAttachmentFormat) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDAttachmentFormat) BaseClass() string {
  return "RDAttachmentFormat"
}



// Enums

func (me *RDAttachmentFormat) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDAttachmentFormat) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDAttachmentFormat) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDAttachmentFormat) SetFormat(p_member RenderingDeviceDataFormat, )  {
  classNameV := StringNameFromStr("RDAttachmentFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 565531219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDAttachmentFormat) GetFormat() RenderingDeviceDataFormat {
  classNameV := StringNameFromStr("RDAttachmentFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2235804183) // FIXME: should cache?
  var ret RenderingDeviceDataFormat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDAttachmentFormat) SetSamples(p_member RenderingDeviceTextureSamples, )  {
  classNameV := StringNameFromStr("RDAttachmentFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_samples")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3774171498) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDAttachmentFormat) GetSamples() RenderingDeviceTextureSamples {
  classNameV := StringNameFromStr("RDAttachmentFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_samples")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 407791724) // FIXME: should cache?
  var ret RenderingDeviceTextureSamples
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDAttachmentFormat) SetUsageFlags(p_member int, )  {
  classNameV := StringNameFromStr("RDAttachmentFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_usage_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDAttachmentFormat) GetUsageFlags() int {
  classNameV := StringNameFromStr("RDAttachmentFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_usage_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
