// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XRCamera3D struct {
  obj gdc.ObjectPtr
}

func (me *XRCamera3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRCamera3D) BaseClass() string {
  return "XRCamera3D"
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
