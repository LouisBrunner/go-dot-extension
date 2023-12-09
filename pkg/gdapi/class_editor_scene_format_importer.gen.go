// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorSceneFormatImporter struct {
  obj gdc.ObjectPtr
}

func (me *EditorSceneFormatImporter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorSceneFormatImporter) BaseClass() string {
  return "EditorSceneFormatImporter"
}



// Constants

var (
  EditorSceneFormatImporterImportScene = "1" // TODO: construct correctly
  EditorSceneFormatImporterImportAnimation = "2" // TODO: construct correctly
  EditorSceneFormatImporterImportFailOnMissingDependencies = "4" // TODO: construct correctly
  EditorSceneFormatImporterImportGenerateTangentArrays = "8" // TODO: construct correctly
  EditorSceneFormatImporterImportUseNamedSkinBinds = "16" // TODO: construct correctly
  EditorSceneFormatImporterImportDiscardMeshesAndMaterials = "32" // TODO: construct correctly
)

// Enums

func (me *EditorSceneFormatImporter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorSceneFormatImporter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorSceneFormatImporter) XGetImportFlags()  {
  panic("TODO: implement")
}

func  (me *EditorSceneFormatImporter) XGetExtensions()  {
  panic("TODO: implement")
}

func  (me *EditorSceneFormatImporter) XImportScene(path String, flags int, options Dictionary, )  {
  panic("TODO: implement")
}

func  (me *EditorSceneFormatImporter) XGetImportOptions(path String, )  {
  panic("TODO: implement")
}

func  (me *EditorSceneFormatImporter) XGetOptionVisibility(path String, for_animation bool, option String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
