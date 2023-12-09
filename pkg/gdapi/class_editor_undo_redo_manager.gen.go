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



// Enums

type EditorUndoRedoManagerSpecialHistory int
const (
  EditorUndoRedoManagerSpecialHistoryGlobalHistory EditorUndoRedoManagerSpecialHistory = 0
  EditorUndoRedoManagerSpecialHistoryRemoteHistory EditorUndoRedoManagerSpecialHistory = -9
  EditorUndoRedoManagerSpecialHistoryInvalidHistory EditorUndoRedoManagerSpecialHistory = -99
)

func (me *EditorUndoRedoManager) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorUndoRedoManager) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorUndoRedoManager) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorUndoRedoManager) CreateAction(name String, merge_mode UndoRedoMergeMode, custom_context Object, backward_undo_ops bool, )  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) CommitAction(execute bool, )  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) IsCommittingAction()  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) AddDoMethod(object Object, method StringName, )  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) AddUndoMethod(object Object, method StringName, )  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) AddDoProperty(object Object, property StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) AddUndoProperty(object Object, property StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) AddDoReference(object Object, )  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) AddUndoReference(object Object, )  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) GetObjectHistoryId(object Object, )  {
  panic("TODO: implement")
}

func  (me *EditorUndoRedoManager) GetHistoryUndoRedo(id int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
