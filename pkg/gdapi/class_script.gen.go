// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Script struct {
  obj gdc.ObjectPtr
}

func (me *Script) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Script) BaseClass() string {
  return "Script"
}



// Enums

func (me *Script) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Script) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Script) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Script) CanInstantiate()  {
  panic("TODO: implement")
}

func  (me *Script) InstanceHas(base_object Object, )  {
  panic("TODO: implement")
}

func  (me *Script) HasSourceCode()  {
  panic("TODO: implement")
}

func  (me *Script) GetSourceCode()  {
  panic("TODO: implement")
}

func  (me *Script) SetSourceCode(source String, )  {
  panic("TODO: implement")
}

func  (me *Script) Reload(keep_state bool, )  {
  panic("TODO: implement")
}

func  (me *Script) GetBaseScript()  {
  panic("TODO: implement")
}

func  (me *Script) GetInstanceBaseType()  {
  panic("TODO: implement")
}

func  (me *Script) HasScriptSignal(signal_name StringName, )  {
  panic("TODO: implement")
}

func  (me *Script) GetScriptPropertyList()  {
  panic("TODO: implement")
}

func  (me *Script) GetScriptMethodList()  {
  panic("TODO: implement")
}

func  (me *Script) GetScriptSignalList()  {
  panic("TODO: implement")
}

func  (me *Script) GetScriptConstantMap()  {
  panic("TODO: implement")
}

func  (me *Script) GetPropertyDefaultValue(property StringName, )  {
  panic("TODO: implement")
}

func  (me *Script) IsTool()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
