// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeSwitch struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeSwitch) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeSwitch) BaseClass() string {
  return "VisualShaderNodeSwitch"
}



// Enums

type VisualShaderNodeSwitchOpType int
const (
  VisualShaderNodeSwitchOpTypeOpTypeFloat VisualShaderNodeSwitchOpType = 0
  VisualShaderNodeSwitchOpTypeOpTypeInt VisualShaderNodeSwitchOpType = 1
  VisualShaderNodeSwitchOpTypeOpTypeUint VisualShaderNodeSwitchOpType = 2
  VisualShaderNodeSwitchOpTypeOpTypeVector2D VisualShaderNodeSwitchOpType = 3
  VisualShaderNodeSwitchOpTypeOpTypeVector3D VisualShaderNodeSwitchOpType = 4
  VisualShaderNodeSwitchOpTypeOpTypeVector4D VisualShaderNodeSwitchOpType = 5
  VisualShaderNodeSwitchOpTypeOpTypeBoolean VisualShaderNodeSwitchOpType = 6
  VisualShaderNodeSwitchOpTypeOpTypeTransform VisualShaderNodeSwitchOpType = 7
  VisualShaderNodeSwitchOpTypeOpTypeMax VisualShaderNodeSwitchOpType = 8
)

func (me *VisualShaderNodeSwitch) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeSwitch) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeSwitch) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeSwitch) SetOpType(type_ VisualShaderNodeSwitchOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeSwitch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 510471861) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeSwitch) GetOpType() VisualShaderNodeSwitchOpType {
  classNameV := StringNameFromStr("VisualShaderNodeSwitch")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2517845071) // FIXME: should cache?
  var ret VisualShaderNodeSwitchOpType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
