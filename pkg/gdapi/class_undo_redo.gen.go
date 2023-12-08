// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type UndoRedo struct {
  obj gdc.ObjectPtr
}

func (me *UndoRedo) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *UndoRedo) BaseClass() string {
  return "UndoRedo"
}

type UndoRedoMergeMode int
const (
  UndoRedoMergeModeMergeDisable UndoRedoMergeMode = 0
  UndoRedoMergeModeMergeEnds UndoRedoMergeMode = 1
  UndoRedoMergeModeMergeAll UndoRedoMergeMode = 2
)

func  (me *UndoRedo) CreateAction(name String, merge_mode UndoRedoMergeMode, backward_undo_ops bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) CommitAction(execute bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) IsCommittingAction() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) AddDoMethod(callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) AddUndoMethod(callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) AddDoProperty(object Object, property StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) AddUndoProperty(object Object, property StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) AddDoReference(object Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) AddUndoReference(object Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) StartForceKeepInMergeEnds() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) EndForceKeepInMergeEnds() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) GetHistoryCount() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) GetCurrentAction() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) GetActionName(id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) ClearHistory(increase_version bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) GetCurrentActionName() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) HasUndo() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) HasRedo() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) GetVersion() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) Redo() { // TODO: return value
  // TODO: implement
}

func  (me *UndoRedo) Undo() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
