// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorInspectorPlugin struct {
  obj gdc.ObjectPtr
}

func createEditorInspectorPlugin(obj gdc.ObjectPtr) *EditorInspectorPlugin {
  return &EditorInspectorPlugin{
    obj: obj,
  }
}

func (me *EditorInspectorPlugin) BaseClass() string {
  return "EditorInspectorPlugin"
}
