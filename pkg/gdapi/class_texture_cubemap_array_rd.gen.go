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

type TextureCubemapArrayRD struct {
  TextureLayeredRD
}

func (me *TextureCubemapArrayRD) BaseClass() string {
  return "TextureCubemapArrayRD"
}

func NewTextureCubemapArrayRD() *TextureCubemapArrayRD {
  str := StringNameFromStr("TextureCubemapArrayRD") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TextureCubemapArrayRD{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *TextureCubemapArrayRD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureCubemapArrayRD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureCubemapArrayRD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
