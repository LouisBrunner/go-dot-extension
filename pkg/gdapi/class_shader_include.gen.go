// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ShaderInclude struct {
  obj gdc.ObjectPtr
}

func (me *ShaderInclude) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ShaderInclude) BaseClass() string {
  return "ShaderInclude"
}



// Enums

func (me *ShaderInclude) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ShaderInclude) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ShaderInclude) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ShaderInclude) SetCode(code String, )  {
  classNameV := StringNameFromStr("ShaderInclude")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_code")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(code.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShaderInclude) GetCode() String {
  classNameV := StringNameFromStr("ShaderInclude")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_code")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
