// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorPlugin struct {
  obj gdc.ObjectPtr
}

func createEditorPlugin(obj gdc.ObjectPtr) *EditorPlugin {
  return &EditorPlugin{
    obj: obj,
  }
}

func (me *EditorPlugin) BaseClass() string {
  return "EditorPlugin"
}
