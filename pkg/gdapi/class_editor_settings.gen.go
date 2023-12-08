// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorSettings struct {
  obj gdc.ObjectPtr
}

func createEditorSettings(obj gdc.ObjectPtr) *EditorSettings {
  return &EditorSettings{
    obj: obj,
  }
}

func (me *EditorSettings) BaseClass() string {
  return "EditorSettings"
}
