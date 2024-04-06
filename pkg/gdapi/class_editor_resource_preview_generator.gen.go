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

type ptrsForEditorResourcePreviewGeneratorList struct {
	fnXHandles                           gdc.MethodBindPtr
	fnXGenerate                          gdc.MethodBindPtr
	fnXGenerateFromPath                  gdc.MethodBindPtr
	fnXGenerateSmallPreviewAutomatically gdc.MethodBindPtr
	fnXCanGenerateSmallPreview           gdc.MethodBindPtr
}

var ptrsForEditorResourcePreviewGenerator ptrsForEditorResourcePreviewGeneratorList

func initEditorResourcePreviewGeneratorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorResourcePreviewGenerator")
	defer className.Destroy()

}

type EditorResourcePreviewGenerator struct {
	RefCounted
}

func (me *EditorResourcePreviewGenerator) BaseClass() string {
	return "EditorResourcePreviewGenerator"
}

func NewEditorResourcePreviewGenerator() *EditorResourcePreviewGenerator {
	str := StringNameFromStr("EditorResourcePreviewGenerator") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorResourcePreviewGenerator{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorResourcePreviewGenerator) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorResourcePreviewGenerator) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorResourcePreviewGenerator) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
