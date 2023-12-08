// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorFileDialog struct {
  obj gdc.ObjectPtr
}

func createEditorFileDialog(obj gdc.ObjectPtr) *EditorFileDialog {
  return &EditorFileDialog{
    obj: obj,
  }
}

func (me *EditorFileDialog) BaseClass() string {
  return "EditorFileDialog"
}
