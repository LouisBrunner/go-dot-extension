// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorExportPlugin struct {
  obj gdc.ObjectPtr
}

func createEditorExportPlugin(obj gdc.ObjectPtr) *EditorExportPlugin {
  return &EditorExportPlugin{
    obj: obj,
  }
}

func (me *EditorExportPlugin) BaseClass() string {
  return "EditorExportPlugin"
}
