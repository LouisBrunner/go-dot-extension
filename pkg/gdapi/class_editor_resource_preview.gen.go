// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorResourcePreview struct {
  obj gdc.ObjectPtr
}

func createEditorResourcePreview(obj gdc.ObjectPtr) *EditorResourcePreview {
  return &EditorResourcePreview{
    obj: obj,
  }
}

func (me *EditorResourcePreview) BaseClass() string {
  return "EditorResourcePreview"
}
