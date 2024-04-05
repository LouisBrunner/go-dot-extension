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

type ptrsForEditorDebuggerPluginList struct {
  fnXSetupSession gdc.MethodBindPtr
  fnXHasCapture gdc.MethodBindPtr
  fnXCapture gdc.MethodBindPtr
  fnGetSession gdc.MethodBindPtr
  fnGetSessions gdc.MethodBindPtr
}

var ptrsForEditorDebuggerPlugin ptrsForEditorDebuggerPluginList

func initEditorDebuggerPluginPtrs(iface gdc.Interface) {

  className := StringNameFromStr("EditorDebuggerPlugin")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_session")
    defer methodName.Destroy()
    ptrsForEditorDebuggerPlugin.fnGetSession = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3061968499))
  }
  {
    methodName := StringNameFromStr("get_sessions")
    defer methodName.Destroy()
    ptrsForEditorDebuggerPlugin.fnGetSessions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewEditorDebuggerSession()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorDebuggerPlugin.fnGetSession), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *EditorDebuggerPlugin) GetSessions() Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorDebuggerPlugin.fnGetSessions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
