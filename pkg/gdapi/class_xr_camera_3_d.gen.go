// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XRCamera3D struct {
  Camera3D
}

func (me *XRCamera3D) BaseClass() string {
  return "XRCamera3D"
}

func NewXRCamera3D() *XRCamera3D {
  str := StringNameFromStr("XRCamera3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &XRCamera3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *XRCamera3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XRCamera3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRCamera3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
