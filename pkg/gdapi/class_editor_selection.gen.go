// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorSelection struct {
  obj gdc.ObjectPtr
}

func createEditorSelection(obj gdc.ObjectPtr) *EditorSelection {
  return &EditorSelection{
    obj: obj,
  }
}

func (me *EditorSelection) BaseClass() string {
  return "EditorSelection"
}
