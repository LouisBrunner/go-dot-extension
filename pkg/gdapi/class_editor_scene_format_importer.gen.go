// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type EditorSceneFormatImporter struct {
  RefCounted
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
  EditorSceneFormatImporterImportForceDisableMeshCompression = "64" // TODO: construct correctly
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
