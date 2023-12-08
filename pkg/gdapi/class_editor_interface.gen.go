// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorInterface struct {
  obj gdc.ObjectPtr
}

func createEditorInterface(obj gdc.ObjectPtr) *EditorInterface {
  return &EditorInterface{
    obj: obj,
  }
}

func (me *EditorInterface) BaseClass() string {
  return "EditorInterface"
}
