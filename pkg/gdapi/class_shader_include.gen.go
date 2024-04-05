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

type ptrsForShaderIncludeList struct {
  fnSetCode gdc.MethodBindPtr
  fnGetCode gdc.MethodBindPtr
}

var ptrsForShaderInclude ptrsForShaderIncludeList

func initShaderIncludePtrs(iface gdc.Interface) {

  className := StringNameFromStr("ShaderInclude")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_code")
    defer methodName.Destroy()
    ptrsForShaderInclude.fnSetCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_code")
    defer methodName.Destroy()
    ptrsForShaderInclude.fnGetCode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
}

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
  cargs := []gdc.ConstTypePtr{code.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShaderInclude.fnSetCode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShaderInclude) GetCode() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForShaderInclude.fnGetCode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
