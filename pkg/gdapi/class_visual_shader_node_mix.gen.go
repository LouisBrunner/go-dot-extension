// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeMix struct {
  VisualShaderNode
}

func (me *VisualShaderNodeMix) BaseClass() string {
  return "VisualShaderNodeMix"
}

func NewVisualShaderNodeMix() *VisualShaderNodeMix {
  str := StringNameFromStr("VisualShaderNodeMix") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeMix{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeMixOpType int
const (
  VisualShaderNodeMixOpTypeOpTypeScalar VisualShaderNodeMixOpType = 0
  VisualShaderNodeMixOpTypeOpTypeVector2D VisualShaderNodeMixOpType = 1
  VisualShaderNodeMixOpTypeOpTypeVector2DScalar VisualShaderNodeMixOpType = 2
  VisualShaderNodeMixOpTypeOpTypeVector3D VisualShaderNodeMixOpType = 3
  VisualShaderNodeMixOpTypeOpTypeVector3DScalar VisualShaderNodeMixOpType = 4
  VisualShaderNodeMixOpTypeOpTypeVector4D VisualShaderNodeMixOpType = 5
  VisualShaderNodeMixOpTypeOpTypeVector4DScalar VisualShaderNodeMixOpType = 6
  VisualShaderNodeMixOpTypeOpTypeMax VisualShaderNodeMixOpType = 7
)

func (me *VisualShaderNodeMix) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeMix) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeMix) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeMix) SetOpType(op_type VisualShaderNodeMixOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeMix")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3397501671) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op_type), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeMix) GetOpType() VisualShaderNodeMixOpType {
  classNameV := StringNameFromStr("VisualShaderNodeMix")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4013957297) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret VisualShaderNodeMixOpType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
