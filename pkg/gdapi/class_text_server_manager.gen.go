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

type ptrsForTextServerManagerList struct {
  fnAddInterface gdc.MethodBindPtr
  fnGetInterfaceCount gdc.MethodBindPtr
  fnRemoveInterface gdc.MethodBindPtr
  fnGetInterface gdc.MethodBindPtr
  fnGetInterfaces gdc.MethodBindPtr
  fnFindInterface gdc.MethodBindPtr
  fnSetPrimaryInterface gdc.MethodBindPtr
  fnGetPrimaryInterface gdc.MethodBindPtr
}

var ptrsForTextServerManager ptrsForTextServerManagerList

func initTextServerManagerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("TextServerManager")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_interface")
    defer methodName.Destroy()
    ptrsForTextServerManager.fnAddInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1799689403))
  }
  {
    methodName := StringNameFromStr("get_interface_count")
    defer methodName.Destroy()
    ptrsForTextServerManager.fnGetInterfaceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("remove_interface")
    defer methodName.Destroy()
    ptrsForTextServerManager.fnRemoveInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1799689403))
  }
  {
    methodName := StringNameFromStr("get_interface")
    defer methodName.Destroy()
    ptrsForTextServerManager.fnGetInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1672475555))
  }
  {
    methodName := StringNameFromStr("get_interfaces")
    defer methodName.Destroy()
    ptrsForTextServerManager.fnGetInterfaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("find_interface")
    defer methodName.Destroy()
    ptrsForTextServerManager.fnFindInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240905781))
  }
  {
    methodName := StringNameFromStr("set_primary_interface")
    defer methodName.Destroy()
    ptrsForTextServerManager.fnSetPrimaryInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1799689403))
  }
  {
    methodName := StringNameFromStr("get_primary_interface")
    defer methodName.Destroy()
    ptrsForTextServerManager.fnGetPrimaryInterface = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 905850878))
  }
}

type TextServerManager struct {
  Object
}

func (me *TextServerManager) BaseClass() string {
  return "TextServerManager"
}

func NewTextServerManager() *TextServerManager {
  str := StringNameFromStr("TextServerManager") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TextServerManager{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *TextServerManager) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextServerManager) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextServerManager) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TextServerManager) AddInterface(interface_ TextServer, )  {
  cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServerManager.fnAddInterface), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextServerManager) GetInterfaceCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServerManager.fnGetInterfaceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TextServerManager) RemoveInterface(interface_ TextServer, )  {
  cargs := []gdc.ConstTypePtr{interface_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServerManager.fnRemoveInterface), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextServerManager) GetInterface(idx int64, ) TextServer {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTextServer()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServerManager.fnGetInterface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextServerManager) GetInterfaces() []Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServerManager.fnGetInterfaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *TextServerManager) FindInterface(name String, ) TextServer {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTextServer()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServerManager.fnFindInterface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TextServerManager) SetPrimaryInterface(index TextServer, )  {
  cargs := []gdc.ConstTypePtr{index.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServerManager.fnSetPrimaryInterface), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TextServerManager) GetPrimaryInterface() TextServer {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTextServer()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextServerManager.fnGetPrimaryInterface), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

type TextServerManagerInterfaceAddedSignalFn func(interface_name StringName, )

func (me *TextServerManager) ConnectInterfaceAdded(subs SignalSubscribers, fn TextServerManagerInterfaceAddedSignalFn) {
  sig := StringNameFromStr("interface_added")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextServerManager) DisconnectInterfaceAdded(subs SignalSubscribers, fn TextServerManagerInterfaceAddedSignalFn) {
  sig := StringNameFromStr("interface_added")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type TextServerManagerInterfaceRemovedSignalFn func(interface_name StringName, )

func (me *TextServerManager) ConnectInterfaceRemoved(subs SignalSubscribers, fn TextServerManagerInterfaceRemovedSignalFn) {
  sig := StringNameFromStr("interface_removed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextServerManager) DisconnectInterfaceRemoved(subs SignalSubscribers, fn TextServerManagerInterfaceRemovedSignalFn) {
  sig := StringNameFromStr("interface_removed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
