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

type ptrsForGLTFDocumentExtensionList struct {
  fnXImportPreflight gdc.MethodBindPtr
  fnXGetSupportedExtensions gdc.MethodBindPtr
  fnXParseNodeExtensions gdc.MethodBindPtr
  fnXParseImageData gdc.MethodBindPtr
  fnXGetImageFileExtension gdc.MethodBindPtr
  fnXParseTextureJson gdc.MethodBindPtr
  fnXGenerateSceneNode gdc.MethodBindPtr
  fnXImportPostParse gdc.MethodBindPtr
  fnXImportNode gdc.MethodBindPtr
  fnXImportPost gdc.MethodBindPtr
  fnXExportPreflight gdc.MethodBindPtr
  fnXConvertSceneNode gdc.MethodBindPtr
  fnXExportPreserialize gdc.MethodBindPtr
  fnXGetSaveableImageFormats gdc.MethodBindPtr
  fnXSerializeImageToBytes gdc.MethodBindPtr
  fnXSaveImageAtPath gdc.MethodBindPtr
  fnXSerializeTextureJson gdc.MethodBindPtr
  fnXExportNode gdc.MethodBindPtr
  fnXExportPost gdc.MethodBindPtr
}

var ptrsForGLTFDocumentExtension ptrsForGLTFDocumentExtensionList

func initGLTFDocumentExtensionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GLTFDocumentExtension")
  defer className.Destroy()
}

type GLTFDocumentExtension struct {
  Resource
}

func (me *GLTFDocumentExtension) BaseClass() string {
  return "GLTFDocumentExtension"
}

func NewGLTFDocumentExtension() *GLTFDocumentExtension {
  str := StringNameFromStr("GLTFDocumentExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFDocumentExtension{}
  obj.SetBaseObject(objPtr)
  return obj
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
