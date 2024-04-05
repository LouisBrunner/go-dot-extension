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

type ptrsForPlaceholderTextureLayeredList struct {
  fnSetSize gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnSetLayers gdc.MethodBindPtr
}

var ptrsForPlaceholderTextureLayered ptrsForPlaceholderTextureLayeredList

func initPlaceholderTextureLayeredPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PlaceholderTextureLayered")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_size")
    defer methodName.Destroy()
    ptrsForPlaceholderTextureLayered.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForPlaceholderTextureLayered.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_layers")
    defer methodName.Destroy()
    ptrsForPlaceholderTextureLayered.fnSetLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
}

type PlaceholderTextureLayered struct {
  TextureLayered
}

func (me *PlaceholderTextureLayered) BaseClass() string {
  return "PlaceholderTextureLayered"
}

func NewPlaceholderTextureLayered() *PlaceholderTextureLayered {
  str := StringNameFromStr("PlaceholderTextureLayered") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PlaceholderTextureLayered{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaceholderTextureLayered.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PlaceholderTextureLayered) GetSize() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaceholderTextureLayered.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PlaceholderTextureLayered) SetLayers(layers int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPlaceholderTextureLayered.fnSetLayers), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
