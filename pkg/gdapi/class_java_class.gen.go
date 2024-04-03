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

func NewJavaClass() *JavaClass {
  str := StringNameFromStr("JavaClass") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &JavaClass{}
  obj.SetBaseObject(objPtr)
  return obj
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
