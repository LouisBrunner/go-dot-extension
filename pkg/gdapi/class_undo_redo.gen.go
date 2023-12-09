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



// Enums

type UndoRedoMergeMode int
const (
  UndoRedoMergeModeMergeDisable UndoRedoMergeMode = 0
  UndoRedoMergeModeMergeEnds UndoRedoMergeMode = 1
  UndoRedoMergeModeMergeAll UndoRedoMergeMode = 2
)

func (me *UndoRedo) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *UndoRedo) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *UndoRedo) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *UndoRedo) CreateAction(name String, merge_mode UndoRedoMergeMode, backward_undo_ops bool, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) CommitAction(execute bool, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) IsCommittingAction()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) AddDoMethod(callable Callable, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) AddUndoMethod(callable Callable, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) AddDoProperty(object Object, property StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) AddUndoProperty(object Object, property StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) AddDoReference(object Object, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) AddUndoReference(object Object, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) StartForceKeepInMergeEnds()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) EndForceKeepInMergeEnds()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) GetHistoryCount()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) GetCurrentAction()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) GetActionName(id int, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) ClearHistory(increase_version bool, )  {
  panic("TODO: implement")
}

func  (me *UndoRedo) GetCurrentActionName()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) HasUndo()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) HasRedo()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) GetVersion()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) Redo()  {
  panic("TODO: implement")
}

func  (me *UndoRedo) Undo()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
