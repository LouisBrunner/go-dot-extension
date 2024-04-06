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

type ptrsForEditorImportPluginList struct {
	fnXGetImporterName             gdc.MethodBindPtr
	fnXGetVisibleName              gdc.MethodBindPtr
	fnXGetPresetCount              gdc.MethodBindPtr
	fnXGetPresetName               gdc.MethodBindPtr
	fnXGetRecognizedExtensions     gdc.MethodBindPtr
	fnXGetImportOptions            gdc.MethodBindPtr
	fnXGetSaveExtension            gdc.MethodBindPtr
	fnXGetResourceType             gdc.MethodBindPtr
	fnXGetPriority                 gdc.MethodBindPtr
	fnXGetImportOrder              gdc.MethodBindPtr
	fnXGetOptionVisibility         gdc.MethodBindPtr
	fnXImport                      gdc.MethodBindPtr
	fnAppendImportExternalResource gdc.MethodBindPtr
}

var ptrsForEditorImportPlugin ptrsForEditorImportPluginList

func initEditorImportPluginPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorImportPlugin")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("append_import_external_resource")
		defer methodName.Destroy()
		ptrsForEditorImportPlugin.fnAppendImportExternalResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 320493106))
	}
}

type EditorImportPlugin struct {
	ResourceImporter
}

func (me *EditorImportPlugin) BaseClass() string {
	return "EditorImportPlugin"
}

func NewEditorImportPlugin() *EditorImportPlugin {
	str := StringNameFromStr("EditorImportPlugin") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorImportPlugin{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorImportPlugin) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorImportPlugin) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorImportPlugin) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorImportPlugin) AppendImportExternalResource(path String, custom_options Dictionary, custom_importer String, generator_parameters Variant) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), custom_options.AsCTypePtr(), custom_importer.AsCTypePtr(), generator_parameters.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorImportPlugin.fnAppendImportExternalResource), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Signals
