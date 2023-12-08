// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorImportPlugin struct {
  obj gdc.ObjectPtr
}

func createEditorImportPlugin(obj gdc.ObjectPtr) *EditorImportPlugin {
  return &EditorImportPlugin{
    obj: obj,
  }
}

func (me *EditorImportPlugin) BaseClass() string {
  return "EditorImportPlugin"
}
