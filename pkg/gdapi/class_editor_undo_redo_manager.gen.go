// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorUndoRedoManager struct {
  obj gdc.ObjectPtr
}

func createEditorUndoRedoManager(obj gdc.ObjectPtr) *EditorUndoRedoManager {
  return &EditorUndoRedoManager{
    obj: obj,
  }
}

func (me *EditorUndoRedoManager) BaseClass() string {
  return "EditorUndoRedoManager"
}
