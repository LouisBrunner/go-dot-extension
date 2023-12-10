// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PlaceholderTextureLayered struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderTextureLayered) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderTextureLayered) BaseClass() string {
  return "PlaceholderTextureLayered"
}



// Enums

func (me *PlaceholderTextureLayered) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderTextureLayered) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTextureLayered) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PlaceholderTextureLayered) SetSize(size Vector2i, )  {
  classNameV := StringNameFromStr("PlaceholderTextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PlaceholderTextureLayered) GetSize() Vector2i {
  classNameV := StringNameFromStr("PlaceholderTextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PlaceholderTextureLayered) SetLayers(layers int, )  {
  classNameV := StringNameFromStr("PlaceholderTextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *PlaceholderTextureLayered) GetPropSize() Vector2i {
  panic("TODO: implement")
}

func (me *PlaceholderTextureLayered) SetPropSize(value Vector2i) {
  panic("TODO: implement")
}

func (me *PlaceholderTextureLayered) GetPropLayers() int {
  panic("TODO: implement")
}

func (me *PlaceholderTextureLayered) SetPropLayers(value int) {
  panic("TODO: implement")
}