// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFDocumentExtension struct {
  obj gdc.ObjectPtr
}

func (me *GLTFDocumentExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFDocumentExtension) BaseClass() string {
  return "GLTFDocumentExtension"
}



// Enums

func (me *GLTFDocumentExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFDocumentExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFDocumentExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
