// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFDocumentExtensionConvertImporterMesh struct {
  GLTFDocumentExtension
}

func (me *GLTFDocumentExtensionConvertImporterMesh) BaseClass() string {
  return "GLTFDocumentExtensionConvertImporterMesh"
}

func NewGLTFDocumentExtensionConvertImporterMesh() *GLTFDocumentExtensionConvertImporterMesh {
  str := StringNameFromStr("GLTFDocumentExtensionConvertImporterMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFDocumentExtensionConvertImporterMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GLTFDocumentExtensionConvertImporterMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFDocumentExtensionConvertImporterMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFDocumentExtensionConvertImporterMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
