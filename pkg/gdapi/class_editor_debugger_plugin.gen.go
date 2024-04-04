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

type EditorDebuggerPlugin struct {
  RefCounted
}

func (me *EditorDebuggerPlugin) BaseClass() string {
  return "EditorDebuggerPlugin"
}

func NewEditorDebuggerPlugin() *EditorDebuggerPlugin {
  str := StringNameFromStr("EditorDebuggerPlugin") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &EditorDebuggerPlugin{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *EditorDebuggerPlugin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *EditorDebuggerPlugin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorDebuggerPlugin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *EditorDebuggerPlugin) GetSession(id int64, ) EditorDebuggerSession {
  classNameV := StringNameFromStr("EditorDebuggerPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_session")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3061968499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorDebuggerSession()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorDebuggerPlugin) GetSessions() Array {
  classNameV := StringNameFromStr("EditorDebuggerPlugin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sessions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
