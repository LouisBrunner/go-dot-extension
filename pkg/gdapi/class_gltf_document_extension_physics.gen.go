// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFDocumentExtensionPhysics struct {
  obj gdc.ObjectPtr
}

func (me *GLTFDocumentExtensionPhysics) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFDocumentExtensionPhysics) BaseClass() string {
  return "GLTFDocumentExtensionPhysics"
}



// Enums

func (me *GLTFDocumentExtensionPhysics) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFDocumentExtensionPhysics) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFDocumentExtensionPhysics) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
