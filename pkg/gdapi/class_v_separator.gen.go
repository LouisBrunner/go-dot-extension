// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VSeparator struct {
  obj gdc.ObjectPtr
}

func (me *VSeparator) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VSeparator) BaseClass() string {
  return "VSeparator"
}



// Enums

func (me *VSeparator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VSeparator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VSeparator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties