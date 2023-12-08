// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorScriptPicker struct {
  obj gdc.ObjectPtr
}

func createEditorScriptPicker(obj gdc.ObjectPtr) *EditorScriptPicker {
  return &EditorScriptPicker{
    obj: obj,
  }
}

func (me *EditorScriptPicker) BaseClass() string {
  return "EditorScriptPicker"
}
