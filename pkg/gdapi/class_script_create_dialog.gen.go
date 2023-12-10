// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("ScriptCreateDialog")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("config")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4210001628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(inherits.AsCTypePtr()), gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&built_in_enabled), gdc.ConstTypePtr(&load_enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties
// Signals
// FIXME: can't seem to be able to connect them from this side of the API