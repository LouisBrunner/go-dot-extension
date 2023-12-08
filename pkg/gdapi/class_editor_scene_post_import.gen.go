// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorScenePostImport struct {
  obj gdc.ObjectPtr
}

func createEditorScenePostImport(obj gdc.ObjectPtr) *EditorScenePostImport {
  return &EditorScenePostImport{
    obj: obj,
  }
}

func (me *EditorScenePostImport) BaseClass() string {
  return "EditorScenePostImport"
}
