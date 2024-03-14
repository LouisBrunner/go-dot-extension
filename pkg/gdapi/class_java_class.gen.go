// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type JavaClass struct {
  RefCounted
}

func (me *JavaClass) BaseClass() string {
  return "JavaClass"
}



// Enums

func (me *JavaClass) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JavaClass) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JavaClass) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
