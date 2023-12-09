// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorImportPlugin struct {
  obj gdc.ObjectPtr
}

func (me *EditorImportPlugin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorImportPlugin) BaseClass() string {
  return "EditorImportPlugin"
}



// Enums

func (me *EditorImportPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorImportPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorImportPlugin) XGetImporterName()  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetVisibleName()  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetPresetCount()  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetPresetName(preset_index int, )  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetRecognizedExtensions()  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetImportOptions(path String, preset_index int, )  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetSaveExtension()  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetResourceType()  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetPriority()  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetImportOrder()  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XGetOptionVisibility(path String, option_name StringName, options Dictionary, )  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) XImport(source_file String, save_path String, options Dictionary, platform_variants String, gen_files String, )  {
  panic("TODO: implement")
}

func  (me *EditorImportPlugin) AppendImportExternalResource(path String, custom_options Dictionary, custom_importer String, generator_parameters Variant, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
