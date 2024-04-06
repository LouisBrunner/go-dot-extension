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

type ptrsForGLTFDocumentExtensionConvertImporterMeshList struct {
}

var ptrsForGLTFDocumentExtensionConvertImporterMesh ptrsForGLTFDocumentExtensionConvertImporterMeshList

func initGLTFDocumentExtensionConvertImporterMeshPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GLTFDocumentExtensionConvertImporterMesh")
	defer className.Destroy()
}

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
