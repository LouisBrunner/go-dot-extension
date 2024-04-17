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

type ptrsForEditorSceneFormatImporterFBX2GLTFList struct {
}

var ptrsForEditorSceneFormatImporterFBX2GLTF ptrsForEditorSceneFormatImporterFBX2GLTFList

func initEditorSceneFormatImporterFBX2GLTFPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorSceneFormatImporterFBX2GLTF")
	defer className.Destroy()

}

type EditorSceneFormatImporterFBX2GLTF struct {
	EditorSceneFormatImporter
}

func (me *EditorSceneFormatImporterFBX2GLTF) BaseClass() string {
	return "EditorSceneFormatImporterFBX2GLTF"
}

func NewEditorSceneFormatImporterFBX2GLTF() *EditorSceneFormatImporterFBX2GLTF {
	str := StringNameFromStr("EditorSceneFormatImporterFBX2GLTF") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorSceneFormatImporterFBX2GLTF{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorSceneFormatImporterFBX2GLTF) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorSceneFormatImporterFBX2GLTF) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporterFBX2GLTF) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
