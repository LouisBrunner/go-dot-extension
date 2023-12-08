// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorUndoRedoManager struct {
  obj gdc.ObjectPtr
}

func (me *EditorUndoRedoManager) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorUndoRedoManager) BaseClass() string {
  return "EditorUndoRedoManager"
}

type EditorUndoRedoManagerSpecialHistory int
const (
  EditorUndoRedoManagerGlobalHistory EditorUndoRedoManagerSpecialHistory = 0
  EditorUndoRedoManagerRemoteHistory EditorUndoRedoManagerSpecialHistory = -9
  EditorUndoRedoManagerInvalidHistory EditorUndoRedoManagerSpecialHistory = -99
)
