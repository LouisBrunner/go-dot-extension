// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JavaScriptBridge struct {
  obj gdc.ObjectPtr
}

func (me *JavaScriptBridge) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JavaScriptBridge) BaseClass() string {
  return "JavaScriptBridge"
}



// Enums

func (me *JavaScriptBridge) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JavaScriptBridge) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JavaScriptBridge) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *JavaScriptBridge) Eval(code String, use_global_execution_context bool, )  {
  panic("TODO: implement")
}

func  (me *JavaScriptBridge) GetInterface(interface_ String, )  {
  panic("TODO: implement")
}

func  (me *JavaScriptBridge) CreateCallback(callable Callable, )  {
  panic("TODO: implement")
}

func  (me *JavaScriptBridge) CreateObject(object String, )  {
  panic("TODO: implement")
}

func  (me *JavaScriptBridge) DownloadBuffer(buffer PackedByteArray, name String, mime String, )  {
  panic("TODO: implement")
}

func  (me *JavaScriptBridge) PwaNeedsUpdate()  {
  panic("TODO: implement")
}

func  (me *JavaScriptBridge) PwaUpdate()  {
  panic("TODO: implement")
}

func  (me *JavaScriptBridge) ForceFsSync()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
