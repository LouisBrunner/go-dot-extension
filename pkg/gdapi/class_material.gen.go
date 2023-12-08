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

// TODO: needed?
// const (
// // )

var (
  MaterialRenderPriorityMax = "127" // TODO: construct correctly
  MaterialRenderPriorityMin = "-128" // TODO: construct correctly
)

func  (me *Material) XGetShaderRid() { // TODO: return value
  // TODO: implement
}

func  (me *Material) XGetShaderMode() { // TODO: return value
  // TODO: implement
}

func  (me *Material) XCanDoNextPass() { // TODO: return value
  // TODO: implement
}

func  (me *Material) XCanUseRenderPriority() { // TODO: return value
  // TODO: implement
}

func  (me *Material) SetNextPass(next_pass Material, ) { // TODO: return value
  // TODO: implement
}

func  (me *Material) GetNextPass() { // TODO: return value
  // TODO: implement
}

func  (me *Material) SetRenderPriority(priority int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Material) GetRenderPriority() { // TODO: return value
  // TODO: implement
}

func  (me *Material) InspectNativeShaderCode() { // TODO: return value
  // TODO: implement
}

func  (me *Material) CreatePlaceholder() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
