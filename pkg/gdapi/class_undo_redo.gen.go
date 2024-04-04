// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type UndoRedo struct {
  Object
}

func (me *UndoRedo) BaseClass() string {
  return "UndoRedo"
}

func NewUndoRedo() *UndoRedo {
  str := StringNameFromStr("UndoRedo") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &UndoRedo{}
  obj.SetBaseObject(objPtr)
  return obj
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
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3171901514) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&merge_mode) , gdc.ConstTypePtr(&backward_undo_ops) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) CommitAction(execute bool, )  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("commit_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3216645846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&execute) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) IsCommittingAction() bool {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_committing_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) AddDoMethod(callable Callable, )  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_do_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddUndoMethod(callable Callable, )  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1611583062) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddDoProperty(object Object, property StringName, value Variant, )  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_do_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1017172818) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddUndoProperty(object Object, property StringName, value Variant, )  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1017172818) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddDoReference(object Object, )  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_do_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3975164845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddUndoReference(object Object, )  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3975164845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) StartForceKeepInMergeEnds()  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_force_keep_in_merge_ends")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) EndForceKeepInMergeEnds()  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("end_force_keep_in_merge_ends")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) GetHistoryCount() int64 {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_history_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) GetCurrentAction() int64 {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) GetActionName(id int64, ) String {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_action_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 990163283) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UndoRedo) ClearHistory(increase_version bool, )  {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_history")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3216645846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&increase_version) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) GetCurrentActionName() String {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_action_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UndoRedo) HasUndo() bool {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_undo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) HasRedo() bool {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_redo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) GetVersion() int64 {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) Redo() bool {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("redo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) Undo() bool {
  classNameV := StringNameFromStr("UndoRedo")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("undo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals

type UndoRedoVersionChangedSignalFn func()

func (me *UndoRedo) ConnectVersionChanged(subs SignalSubscribers, fn UndoRedoVersionChangedSignalFn) {
  sig := StringNameFromStr("version_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *UndoRedo) DisconnectVersionChanged(subs SignalSubscribers, fn UndoRedoVersionChangedSignalFn) {
  sig := StringNameFromStr("version_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
