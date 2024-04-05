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

type ptrsForCubemapArrayList struct {
  fnCreatePlaceholder gdc.MethodBindPtr
}

var ptrsForCubemapArray ptrsForCubemapArrayList

func initCubemapArrayPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CubemapArray")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("create_placeholder")
    defer methodName.Destroy()
    ptrsForCubemapArray.fnCreatePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121922552))
  }
}

type CubemapArray struct {
  ImageTextureLayered
}

func (me *CubemapArray) BaseClass() string {
  return "CubemapArray"
}

func NewCubemapArray() *CubemapArray {
  str := StringNameFromStr("CubemapArray") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CubemapArray{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CubemapArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CubemapArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CubemapArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CubemapArray) CreatePlaceholder() Resource {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCubemapArray.fnCreatePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
