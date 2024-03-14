// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorImportPlugin struct {
  ResourceImporter
}

func (me *EditorImportPlugin) BaseClass() string {
  return "EditorImportPlugin"
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
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(custom_options.AsCTypePtr()), gdc.ConstTypePtr(custom_importer.AsCTypePtr()), gdc.ConstTypePtr(generator_parameters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
