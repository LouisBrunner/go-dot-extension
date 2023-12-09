// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptCreateDialog struct {
  obj gdc.ObjectPtr
}

func (me *ScriptCreateDialog) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScriptCreateDialog) BaseClass() string {
  return "ScriptCreateDialog"
}



// Enums

func (me *ScriptCreateDialog) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScriptCreateDialog) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScriptCreateDialog) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ScriptCreateDialog) Config(inherits String, path String, built_in_enabled bool, load_enabled bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
