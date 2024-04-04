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

func  (me *EditorImportPlugin) AppendImportExternalResource(path String, custom_options Dictionary, custom_importer String, generator_parameters Variant, ) Error {
  classNameV := StringNameFromStr("EditorImportPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("append_import_external_resource")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 320493106) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), custom_options.AsCTypePtr(), custom_importer.AsCTypePtr(), generator_parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

// Signals
