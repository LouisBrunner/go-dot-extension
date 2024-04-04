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

type Texture2DArrayRD struct {
  TextureLayeredRD
}

func (me *Texture2DArrayRD) BaseClass() string {
  return "Texture2DArrayRD"
}

func NewTexture2DArrayRD() *Texture2DArrayRD {
  str := StringNameFromStr("Texture2DArrayRD") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Texture2DArrayRD{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Texture2DArrayRD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Texture2DArrayRD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture2DArrayRD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
