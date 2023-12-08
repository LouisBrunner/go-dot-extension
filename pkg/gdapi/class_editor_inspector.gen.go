// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorInspector struct {
  obj gdc.ObjectPtr
}

func createEditorInspector(obj gdc.ObjectPtr) *EditorInspector {
  return &EditorInspector{
    obj: obj,
  }
}

func (me *EditorInspector) BaseClass() string {
  return "EditorInspector"
}
