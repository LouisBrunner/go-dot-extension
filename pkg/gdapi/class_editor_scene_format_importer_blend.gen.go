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

type ptrsForEditorSceneFormatImporterBlendList struct {
}

var ptrsForEditorSceneFormatImporterBlend ptrsForEditorSceneFormatImporterBlendList

func initEditorSceneFormatImporterBlendPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorSceneFormatImporterBlend")
	defer className.Destroy()

}

type EditorSceneFormatImporterBlend struct {
	EditorSceneFormatImporter
}

func (me *EditorSceneFormatImporterBlend) BaseClass() string {
	return "EditorSceneFormatImporterBlend"
}

func NewEditorSceneFormatImporterBlend() *EditorSceneFormatImporterBlend {
	str := StringNameFromStr("EditorSceneFormatImporterBlend") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorSceneFormatImporterBlend{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorSceneFormatImporterBlend) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorSceneFormatImporterBlend) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporterBlend) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
