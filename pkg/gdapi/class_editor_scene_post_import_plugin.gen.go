// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorScenePostImportPlugin struct {
  obj gdc.ObjectPtr
}

func createEditorScenePostImportPlugin(obj gdc.ObjectPtr) *EditorScenePostImportPlugin {
  return &EditorScenePostImportPlugin{
    obj: obj,
  }
}

func (me *EditorScenePostImportPlugin) BaseClass() string {
  return "EditorScenePostImportPlugin"
}
