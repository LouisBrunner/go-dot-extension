// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorVCSInterface struct {
  obj gdc.ObjectPtr
}

func createEditorVCSInterface(obj gdc.ObjectPtr) *EditorVCSInterface {
  return &EditorVCSInterface{
    obj: obj,
  }
}

func (me *EditorVCSInterface) BaseClass() string {
  return "EditorVCSInterface"
}
