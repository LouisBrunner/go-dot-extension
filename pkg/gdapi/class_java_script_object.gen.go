// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type JavaScriptObject struct {
  obj gdc.ObjectPtr
}

func (me *JavaScriptObject) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JavaScriptObject) BaseClass() string {
  return "JavaScriptObject"
}



// Enums

func (me *JavaScriptObject) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JavaScriptObject) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JavaScriptObject) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
