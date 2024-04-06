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

type ptrsForEditorSceneFormatImporterList struct {
	fnXGetImportFlags      gdc.MethodBindPtr
	fnXGetExtensions       gdc.MethodBindPtr
	fnXImportScene         gdc.MethodBindPtr
	fnXGetImportOptions    gdc.MethodBindPtr
	fnXGetOptionVisibility gdc.MethodBindPtr
}

var ptrsForEditorSceneFormatImporter ptrsForEditorSceneFormatImporterList

func initEditorSceneFormatImporterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorSceneFormatImporter")
	defer className.Destroy()

}

type EditorSceneFormatImporter struct {
	RefCounted
}

func (me *EditorSceneFormatImporter) BaseClass() string {
	return "EditorSceneFormatImporter"
}

func NewEditorSceneFormatImporter() *EditorSceneFormatImporter {
	str := StringNameFromStr("EditorSceneFormatImporter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorSceneFormatImporter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	EditorSceneFormatImporterImportScene                       = 1
	EditorSceneFormatImporterImportAnimation                   = 2
	EditorSceneFormatImporterImportFailOnMissingDependencies   = 4
	EditorSceneFormatImporterImportGenerateTangentArrays       = 8
	EditorSceneFormatImporterImportUseNamedSkinBinds           = 16
	EditorSceneFormatImporterImportDiscardMeshesAndMaterials   = 32
	EditorSceneFormatImporterImportForceDisableMeshCompression = 64
)

// Enums

func (me *EditorSceneFormatImporter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorSceneFormatImporter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
