// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorScript struct {
  obj gdc.ObjectPtr
}

func createEditorScript(obj gdc.ObjectPtr) *EditorScript {
  return &EditorScript{
    obj: obj,
  }
}

func (me *EditorScript) BaseClass() string {
  return "EditorScript"
}
