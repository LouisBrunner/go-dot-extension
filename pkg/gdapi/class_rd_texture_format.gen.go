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

type RDTextureFormat struct {
  RefCounted
}

func (me *RDTextureFormat) BaseClass() string {
  return "RDTextureFormat"
}

func NewRDTextureFormat() *RDTextureFormat {
  str := StringNameFromStr("RDTextureFormat") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDTextureFormat{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RDTextureFormat) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDTextureFormat) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDTextureFormat) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDTextureFormat) SetFormat(p_member RenderingDeviceDataFormat, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 565531219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) GetFormat() RenderingDeviceDataFormat {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2235804183) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceDataFormat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureFormat) SetWidth(p_member int64, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) GetWidth() int64 {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDTextureFormat) SetHeight(p_member int64, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) GetHeight() int64 {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDTextureFormat) SetDepth(p_member int64, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) GetDepth() int64 {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDTextureFormat) SetArrayLayers(p_member int64, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_array_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) GetArrayLayers() int64 {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_array_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDTextureFormat) SetMipmaps(p_member int64, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) GetMipmaps() int64 {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDTextureFormat) SetTextureType(p_member RenderingDeviceTextureType, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 652343381) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) GetTextureType() RenderingDeviceTextureType {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4036357416) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureFormat) SetSamples(p_member RenderingDeviceTextureSamples, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_samples")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3774171498) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) GetSamples() RenderingDeviceTextureSamples {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_samples")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 407791724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureSamples

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureFormat) SetUsageBits(p_member RenderingDeviceTextureUsageBits, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_usage_bits")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 245642367) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) GetUsageBits() RenderingDeviceTextureUsageBits {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_usage_bits")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1313398998) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceTextureUsageBits

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDTextureFormat) AddShareableFormat(format RenderingDeviceDataFormat, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_shareable_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 565531219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDTextureFormat) RemoveShareableFormat(format RenderingDeviceDataFormat, )  {
  classNameV := StringNameFromStr("RDTextureFormat")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_shareable_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 565531219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
