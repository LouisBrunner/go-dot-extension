// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeSample3D struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeSample3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeSample3D) BaseClass() string {
  return "VisualShaderNodeSample3D"
}



// Enums

type VisualShaderNodeSample3DSource int
const (
  VisualShaderNodeSample3DSourceSourceTexture VisualShaderNodeSample3DSource = 0
  VisualShaderNodeSample3DSourceSourcePort VisualShaderNodeSample3DSource = 1
  VisualShaderNodeSample3DSourceSourceMax VisualShaderNodeSample3DSource = 2
)

func (me *VisualShaderNodeSample3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSample3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSample3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeSample3D) SetSource(value VisualShaderNodeSample3DSource, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeSample3D) GetSource()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
