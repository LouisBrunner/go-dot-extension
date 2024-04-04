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

type ShaderInclude struct {
  Resource
}

func (me *ShaderInclude) BaseClass() string {
  return "ShaderInclude"
}

func NewShaderInclude() *ShaderInclude {
  str := StringNameFromStr("ShaderInclude") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ShaderInclude{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{code.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShaderInclude) GetCode() String {
  classNameV := StringNameFromStr("ShaderInclude")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_code")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
