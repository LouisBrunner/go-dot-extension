// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextServerManager struct {
  obj gdc.ObjectPtr
}

func (me *TextServerManager) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextServerManager) BaseClass() string {
  return "TextServerManager"
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
  classNameV := StringNameFromStr("TextServerManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1799689403) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interface_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServerManager) GetInterfaceCount() int {
  classNameV := StringNameFromStr("TextServerManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interface_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServerManager) RemoveInterface(interface_ TextServer, )  {
  classNameV := StringNameFromStr("TextServerManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1799689403) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(interface_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServerManager) GetInterface(idx int, ) TextServer {
  classNameV := StringNameFromStr("TextServerManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1672475555) // FIXME: should cache?
  var ret TextServer
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServerManager) GetInterfaces() Dictionary {
  classNameV := StringNameFromStr("TextServerManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_interfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServerManager) FindInterface(name String, ) TextServer {
  classNameV := StringNameFromStr("TextServerManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240905781) // FIXME: should cache?
  var ret TextServer
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServerManager) SetPrimaryInterface(index TextServer, )  {
  classNameV := StringNameFromStr("TextServerManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_primary_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1799689403) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(index.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServerManager) GetPrimaryInterface() TextServer {
  classNameV := StringNameFromStr("TextServerManager")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_primary_interface")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 905850878) // FIXME: should cache?
  var ret TextServer
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API