// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextureLayered struct {
  obj gdc.ObjectPtr
}

func (me *TextureLayered) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureLayered) BaseClass() string {
  return "TextureLayered"
}



// Enums

type TextureLayeredLayeredType int
const (
  TextureLayeredLayeredTypeLayeredType2DArray TextureLayeredLayeredType = 0
  TextureLayeredLayeredTypeLayeredTypeCubemap TextureLayeredLayeredType = 1
  TextureLayeredLayeredTypeLayeredTypeCubemapArray TextureLayeredLayeredType = 2
)

func (me *TextureLayered) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureLayered) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureLayered) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TextureLayered) GetFormat() ImageFormat {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3847873762) // FIXME: should cache?
  var ret ImageFormat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureLayered) GetLayeredType() TextureLayeredLayeredType {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layered_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 518123893) // FIXME: should cache?
  var ret TextureLayeredLayeredType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureLayered) GetWidth() int {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureLayered) GetHeight() int {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureLayered) GetLayers() int {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureLayered) HasMipmaps() bool {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextureLayered) GetLayerData(layer int, ) Image {
  classNameV := StringNameFromStr("TextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_layer_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3655284255) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
