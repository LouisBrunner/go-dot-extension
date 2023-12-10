// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("VisualShaderNodeSample3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3315130991) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeSample3D) GetSource() VisualShaderNodeSample3DSource {
  classNameV := StringNameFromStr("VisualShaderNodeSample3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1079494121) // FIXME: should cache?
  var ret VisualShaderNodeSample3DSource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeSample3D) GetPropSource() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeSample3D) SetPropSource(value int) {
  panic("TODO: implement")
}