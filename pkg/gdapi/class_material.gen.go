// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Material struct {
  obj gdc.ObjectPtr
}

func (me *Material) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Material) BaseClass() string {
  return "Material"
}



// Constants

var (
  MaterialRenderPriorityMax = "127" // TODO: construct correctly
  MaterialRenderPriorityMin = "-128" // TODO: construct correctly
)

// Enums

func (me *Material) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Material) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Material) XGetShaderRid()  {
  panic("TODO: implement")
}

func  (me *Material) XGetShaderMode()  {
  panic("TODO: implement")
}

func  (me *Material) XCanDoNextPass()  {
  panic("TODO: implement")
}

func  (me *Material) XCanUseRenderPriority()  {
  panic("TODO: implement")
}

func  (me *Material) SetNextPass(next_pass Material, )  {
  panic("TODO: implement")
}

func  (me *Material) GetNextPass()  {
  panic("TODO: implement")
}

func  (me *Material) SetRenderPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *Material) GetRenderPriority()  {
  panic("TODO: implement")
}

func  (me *Material) InspectNativeShaderCode()  {
  panic("TODO: implement")
}

func  (me *Material) CreatePlaceholder()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
