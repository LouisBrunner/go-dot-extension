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

type ptrsForEditorSceneFormatImporterFBXList struct {
}

var ptrsForEditorSceneFormatImporterFBX ptrsForEditorSceneFormatImporterFBXList

func initEditorSceneFormatImporterFBXPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorSceneFormatImporterFBX")
	defer className.Destroy()

}

type EditorSceneFormatImporterFBX struct {
	EditorSceneFormatImporter
}

func (me *EditorSceneFormatImporterFBX) BaseClass() string {
	return "EditorSceneFormatImporterFBX"
}

func NewEditorSceneFormatImporterFBX() *EditorSceneFormatImporterFBX {
	str := StringNameFromStr("EditorSceneFormatImporterFBX") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorSceneFormatImporterFBX{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorSceneFormatImporterFBX) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorSceneFormatImporterFBX) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporterFBX) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
