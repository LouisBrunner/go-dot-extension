// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIs struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeIs) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeIs) BaseClass() string {
  return "VisualShaderNodeIs"
}

type VisualShaderNodeIsFunction int
const (
  VisualShaderNodeIsFunctionFuncIsInf VisualShaderNodeIsFunction = 0
  VisualShaderNodeIsFunctionFuncIsNan VisualShaderNodeIsFunction = 1
  VisualShaderNodeIsFunctionFuncMax VisualShaderNodeIsFunction = 2
)

func  (me *VisualShaderNodeIs) SetFunction(func_ VisualShaderNodeIsFunction, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeIs) GetFunction() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
