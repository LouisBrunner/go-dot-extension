// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorProperty struct {
  obj gdc.ObjectPtr
}

func createEditorProperty(obj gdc.ObjectPtr) *EditorProperty {
  return &EditorProperty{
    obj: obj,
  }
}

func (me *EditorProperty) BaseClass() string {
  return "EditorProperty"
}
