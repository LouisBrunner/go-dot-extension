// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

const (
  EditorSceneFormatImporterIMPORT_SCENE = 1
  EditorSceneFormatImporterIMPORT_ANIMATION = 2
  EditorSceneFormatImporterIMPORT_FAIL_ON_MISSING_DEPENDENCIES = 4
  EditorSceneFormatImporterIMPORT_GENERATE_TANGENT_ARRAYS = 8
  EditorSceneFormatImporterIMPORT_USE_NAMED_SKIN_BINDS = 16
  EditorSceneFormatImporterIMPORT_DISCARD_MESHES_AND_MATERIALS = 32
)
