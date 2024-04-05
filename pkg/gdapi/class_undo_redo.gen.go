// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForUndoRedoList struct {
  fnCreateAction gdc.MethodBindPtr
  fnCommitAction gdc.MethodBindPtr
  fnIsCommittingAction gdc.MethodBindPtr
  fnAddDoMethod gdc.MethodBindPtr
  fnAddUndoMethod gdc.MethodBindPtr
  fnAddDoProperty gdc.MethodBindPtr
  fnAddUndoProperty gdc.MethodBindPtr
  fnAddDoReference gdc.MethodBindPtr
  fnAddUndoReference gdc.MethodBindPtr
  fnStartForceKeepInMergeEnds gdc.MethodBindPtr
  fnEndForceKeepInMergeEnds gdc.MethodBindPtr
  fnGetHistoryCount gdc.MethodBindPtr
  fnGetCurrentAction gdc.MethodBindPtr
  fnGetActionName gdc.MethodBindPtr
  fnClearHistory gdc.MethodBindPtr
  fnGetCurrentActionName gdc.MethodBindPtr
  fnHasUndo gdc.MethodBindPtr
  fnHasRedo gdc.MethodBindPtr
  fnGetVersion gdc.MethodBindPtr
  fnRedo gdc.MethodBindPtr
  fnUndo gdc.MethodBindPtr
}

var ptrsForUndoRedo ptrsForUndoRedoList

func initUndoRedoPtrs(iface gdc.Interface) {

  className := StringNameFromStr("UndoRedo")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("create_action")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnCreateAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3171901514))
  }
  {
    methodName := StringNameFromStr("commit_action")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnCommitAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3216645846))
  }
  {
    methodName := StringNameFromStr("is_committing_action")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnIsCommittingAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("add_do_method")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnAddDoMethod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
  }
  {
    methodName := StringNameFromStr("add_undo_method")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnAddUndoMethod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
  }
  {
    methodName := StringNameFromStr("add_do_property")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnAddDoProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1017172818))
  }
  {
    methodName := StringNameFromStr("add_undo_property")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnAddUndoProperty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1017172818))
  }
  {
    methodName := StringNameFromStr("add_do_reference")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnAddDoReference = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3975164845))
  }
  {
    methodName := StringNameFromStr("add_undo_reference")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnAddUndoReference = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3975164845))
  }
  {
    methodName := StringNameFromStr("start_force_keep_in_merge_ends")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnStartForceKeepInMergeEnds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("end_force_keep_in_merge_ends")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnEndForceKeepInMergeEnds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_history_count")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnGetHistoryCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_current_action")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnGetCurrentAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_action_name")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnGetActionName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 990163283))
  }
  {
    methodName := StringNameFromStr("clear_history")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnClearHistory = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3216645846))
  }
  {
    methodName := StringNameFromStr("get_current_action_name")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnGetCurrentActionName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("has_undo")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnHasUndo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("has_redo")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnHasRedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_version")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnGetVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("redo")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnRedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("undo")
    defer methodName.Destroy()
    ptrsForUndoRedo.fnUndo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
}

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
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&merge_mode) , gdc.ConstTypePtr(&backward_undo_ops) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnCreateAction), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) CommitAction(execute bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&execute) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnCommitAction), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) IsCommittingAction() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnIsCommittingAction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) AddDoMethod(callable Callable, )  {
  cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnAddDoMethod), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddUndoMethod(callable Callable, )  {
  cargs := []gdc.ConstTypePtr{callable.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnAddUndoMethod), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddDoProperty(object Object, property StringName, value Variant, )  {
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnAddDoProperty), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddUndoProperty(object Object, property StringName, value Variant, )  {
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnAddUndoProperty), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddDoReference(object Object, )  {
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnAddDoReference), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) AddUndoReference(object Object, )  {
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnAddUndoReference), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) StartForceKeepInMergeEnds()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnStartForceKeepInMergeEnds), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) EndForceKeepInMergeEnds()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnEndForceKeepInMergeEnds), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) GetHistoryCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnGetHistoryCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) GetCurrentAction() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnGetCurrentAction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) GetActionName(id int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnGetActionName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UndoRedo) ClearHistory(increase_version bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&increase_version) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnClearHistory), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *UndoRedo) GetCurrentActionName() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnGetCurrentActionName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *UndoRedo) HasUndo() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnHasUndo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) HasRedo() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnHasRedo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) GetVersion() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnGetVersion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) Redo() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnRedo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *UndoRedo) Undo() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForUndoRedo.fnUndo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
