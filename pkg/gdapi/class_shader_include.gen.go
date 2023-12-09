// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *ShaderInclude) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ShaderInclude) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *ShaderInclude) SetCode(code String, )  {
  panic("TODO: implement")
}

func  (me *ShaderInclude) GetCode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
