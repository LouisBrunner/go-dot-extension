// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  EditorUndoRedoManagerSpecialHistoryGlobalHistory EditorUndoRedoManagerSpecialHistory = 0
  EditorUndoRedoManagerSpecialHistoryRemoteHistory EditorUndoRedoManagerSpecialHistory = -9
  EditorUndoRedoManagerSpecialHistoryInvalidHistory EditorUndoRedoManagerSpecialHistory = -99
)

func  (me *EditorUndoRedoManager) CreateAction(name String, merge_mode UndoRedoMergeMode, custom_context Object, backward_undo_ops bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) CommitAction(execute bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) IsCommittingAction() { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) AddDoMethod(object Object, method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) AddUndoMethod(object Object, method StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) AddDoProperty(object Object, property StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) AddUndoProperty(object Object, property StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) AddDoReference(object Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) AddUndoReference(object Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) GetObjectHistoryId(object Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *EditorUndoRedoManager) GetHistoryUndoRedo(id int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
