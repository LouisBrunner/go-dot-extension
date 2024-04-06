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

type ptrsForEditorSceneFormatImporterGLTFList struct {
}

var ptrsForEditorSceneFormatImporterGLTF ptrsForEditorSceneFormatImporterGLTFList

func initEditorSceneFormatImporterGLTFPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorSceneFormatImporterGLTF")
	defer className.Destroy()

}

type EditorSceneFormatImporterGLTF struct {
	EditorSceneFormatImporter
}

func (me *EditorSceneFormatImporterGLTF) BaseClass() string {
	return "EditorSceneFormatImporterGLTF"
}

func NewEditorSceneFormatImporterGLTF() *EditorSceneFormatImporterGLTF {
	str := StringNameFromStr("EditorSceneFormatImporterGLTF") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorSceneFormatImporterGLTF{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorSceneFormatImporterGLTF) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorSceneFormatImporterGLTF) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporterGLTF) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
