// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3577985681) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(&merge_mode), gdc.ConstTypePtr(custom_context.AsCTypePtr()), gdc.ConstTypePtr(&backward_undo_ops), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorUndoRedoManager) CommitAction(execute bool, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("commit_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3216645846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&execute), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorUndoRedoManager) IsCommittingAction() bool {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_committing_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorUndoRedoManager) AddDoMethod(object Object, method StringName, varargs ...Variant)  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_do_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1517810467) // FIXME: should cache?
  cargs := []gdc.ConstVariantPtr{gdc.ConstVariantPtr(object.AsCTypePtr()), gdc.ConstVariantPtr(method.AsCTypePtr()), }
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

}

func  (me *EditorUndoRedoManager) AddUndoMethod(object Object, method StringName, varargs ...Variant)  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_method")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1517810467) // FIXME: should cache?
  cargs := []gdc.ConstVariantPtr{gdc.ConstVariantPtr(object.AsCTypePtr()), gdc.ConstVariantPtr(method.AsCTypePtr()), }
  for _, v := range varargs {
    cargs = append(cargs, v.AsCPtr())
  }
  err := &gdc.CallError{}
  giface.ObjectMethodBindCall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, err)
  if err.Error != gdc.CallOk {
    panic(err) // TODO: return `err`?
  }

}

func  (me *EditorUndoRedoManager) AddDoProperty(object Object, property StringName, value Variant, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_do_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1017172818) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorUndoRedoManager) AddUndoProperty(object Object, property StringName, value Variant, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_property")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1017172818) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), gdc.ConstTypePtr(property.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorUndoRedoManager) AddDoReference(object Object, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_do_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3975164845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorUndoRedoManager) AddUndoReference(object Object, )  {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_undo_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3975164845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *EditorUndoRedoManager) GetObjectHistoryId(object Object, ) int {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_object_history_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1107568780) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(object.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *EditorUndoRedoManager) GetHistoryUndoRedo(id int, ) UndoRedo {
  classNameV := StringNameFromStr("EditorUndoRedoManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_history_undo_redo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2417974513) // FIXME: should cache?
  var ret UndoRedo
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals

type EditorUndoRedoManagerHistoryChangedSignalFn func()

func (me *EditorUndoRedoManager) ConnectHistoryChanged(subs SignalSubscribers, fn EditorUndoRedoManagerHistoryChangedSignalFn) {
  sig := StringNameFromStr("history_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *EditorUndoRedoManager) DisconnectHistoryChanged(subs SignalSubscribers, fn EditorUndoRedoManagerHistoryChangedSignalFn) {
  sig := StringNameFromStr("history_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type EditorUndoRedoManagerVersionChangedSignalFn func()

func (me *EditorUndoRedoManager) ConnectVersionChanged(subs SignalSubscribers, fn EditorUndoRedoManagerVersionChangedSignalFn) {
  sig := StringNameFromStr("version_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *EditorUndoRedoManager) DisconnectVersionChanged(subs SignalSubscribers, fn EditorUndoRedoManagerVersionChangedSignalFn) {
  sig := StringNameFromStr("version_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
