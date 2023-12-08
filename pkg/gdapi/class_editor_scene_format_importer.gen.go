// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorSceneFormatImporter struct {
  obj gdc.ObjectPtr
}

func createEditorSceneFormatImporter(obj gdc.ObjectPtr) *EditorSceneFormatImporter {
  return &EditorSceneFormatImporter{
    obj: obj,
  }
}

func (me *EditorSceneFormatImporter) BaseClass() string {
  return "EditorSceneFormatImporter"
}
