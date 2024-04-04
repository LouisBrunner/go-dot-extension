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

type EditorUndoRedoManager struct {
  Object
}

func (me *EditorUndoRedoManager) BaseClass() string {
  return "EditorUndoRedoManager"
}

func NewEditorUndoRedoManager() *EditorUndoRedoManager {
  str := StringNameFromStr("EditorUndoRedoManager") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorUndoRedoManager{}
  obj.SetBaseObject(objPtr)
  return obj
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
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2107025470) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&merge_mode) , custom_context.AsCTypePtr(), gdc.ConstTypePtr(&backward_undo_ops) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorUndoRedoManager) CommitAction(execute bool, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("commit_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3216645846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&execute) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorUndoRedoManager) IsCommittingAction() bool {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
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

func  (me *EditorUndoRedoManager) AddDoMethod(object Object, method StringName, varargs ...Variant)  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_do_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1517810467) // FIXME: should cache?
  cargs := make([]gdc.ConstVariantPtr, 0, 2 + len(varargs))
  var0 := object.AsVariant()
  defer var0.Destroy()
  cargs = append(cargs, var0.AsCPtr())
  var1 := method.AsVariant()
  defer var1.Destroy()
  cargs = append(cargs, var1.AsCPtr())
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }

  cerr := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return
  }

}

func  (me *EditorUndoRedoManager) AddUndoMethod(object Object, method StringName, varargs ...Variant)  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1517810467) // FIXME: should cache?
  cargs := make([]gdc.ConstVariantPtr, 0, 2 + len(varargs))
  var0 := object.AsVariant()
  defer var0.Destroy()
  cargs = append(cargs, var0.AsCPtr())
  var1 := method.AsVariant()
  defer var1.Destroy()
  cargs = append(cargs, var1.AsCPtr())
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }

  cerr := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, cerr)
  if cerr.Error != gdc.CallOk {
    log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
    return
  }

}

func  (me *EditorUndoRedoManager) AddDoProperty(object Object, property StringName, value Variant, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_do_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1017172818) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorUndoRedoManager) AddUndoProperty(object Object, property StringName, value Variant, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1017172818) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), property.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorUndoRedoManager) AddDoReference(object Object, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_do_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3975164845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorUndoRedoManager) AddUndoReference(object Object, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3975164845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *EditorUndoRedoManager) GetObjectHistoryId(object Object, ) int64 {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_object_history_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1107568780) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{object.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *EditorUndoRedoManager) GetHistoryUndoRedo(id int64, ) UndoRedo {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_history_undo_redo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2417974513) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewUndoRedo()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

type EditorUndoRedoManagerHistoryChangedSignalFn func()

func (me *EditorUndoRedoManager) ConnectHistoryChanged(subs SignalSubscribers, fn EditorUndoRedoManagerHistoryChangedSignalFn) {
  sig := StringNameFromStr("history_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorUndoRedoManager) DisconnectHistoryChanged(subs SignalSubscribers, fn EditorUndoRedoManagerHistoryChangedSignalFn) {
  sig := StringNameFromStr("history_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type EditorUndoRedoManagerVersionChangedSignalFn func()

func (me *EditorUndoRedoManager) ConnectVersionChanged(subs SignalSubscribers, fn EditorUndoRedoManagerVersionChangedSignalFn) {
  sig := StringNameFromStr("version_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorUndoRedoManager) DisconnectVersionChanged(subs SignalSubscribers, fn EditorUndoRedoManagerVersionChangedSignalFn) {
  sig := StringNameFromStr("version_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
