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

type StandardMaterial3D struct {
  BaseMaterial3D
}

func (me *StandardMaterial3D) BaseClass() string {
  return "StandardMaterial3D"
}

func NewStandardMaterial3D() *StandardMaterial3D {
  str := StringNameFromStr("StandardMaterial3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StandardMaterial3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *StandardMaterial3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StandardMaterial3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StandardMaterial3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
