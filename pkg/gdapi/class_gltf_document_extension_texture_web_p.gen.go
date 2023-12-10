// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFDocumentExtensionTextureWebP struct {
  obj gdc.ObjectPtr
}

func (me *GLTFDocumentExtensionTextureWebP) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFDocumentExtensionTextureWebP) BaseClass() string {
  return "GLTFDocumentExtensionTextureWebP"
}



// Enums

func (me *GLTFDocumentExtensionTextureWebP) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFDocumentExtensionTextureWebP) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFDocumentExtensionTextureWebP) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties