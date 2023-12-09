// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ShaderGlobalsOverride struct {
  obj gdc.ObjectPtr
}

func (me *ShaderGlobalsOverride) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ShaderGlobalsOverride) BaseClass() string {
  return "ShaderGlobalsOverride"
}



// Enums

func (me *ShaderGlobalsOverride) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ShaderGlobalsOverride) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
