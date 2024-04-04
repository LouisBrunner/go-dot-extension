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

type Cubemap struct {
  ImageTextureLayered
}

func (me *Cubemap) BaseClass() string {
  return "Cubemap"
}

func NewCubemap() *Cubemap {
  str := StringNameFromStr("Cubemap") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Cubemap{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Cubemap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Cubemap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Cubemap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Cubemap) CreatePlaceholder() Resource {
  classNameV := StringNameFromStr("Cubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121922552) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
