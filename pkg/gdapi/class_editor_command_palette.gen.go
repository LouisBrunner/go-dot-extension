// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorCommandPalette struct {
  obj gdc.ObjectPtr
}

func createEditorCommandPalette(obj gdc.ObjectPtr) *EditorCommandPalette {
  return &EditorCommandPalette{
    obj: obj,
  }
}

func (me *EditorCommandPalette) BaseClass() string {
  return "EditorCommandPalette"
}
