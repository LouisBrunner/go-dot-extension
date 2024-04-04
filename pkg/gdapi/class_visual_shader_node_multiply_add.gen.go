// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeMultiplyAdd struct {
  VisualShaderNode
}

func (me *VisualShaderNodeMultiplyAdd) BaseClass() string {
  return "VisualShaderNodeMultiplyAdd"
}

func NewVisualShaderNodeMultiplyAdd() *VisualShaderNodeMultiplyAdd {
  str := StringNameFromStr("VisualShaderNodeMultiplyAdd") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeMultiplyAdd{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeMultiplyAddOpType int
const (
  VisualShaderNodeMultiplyAddOpTypeOpTypeScalar VisualShaderNodeMultiplyAddOpType = 0
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector2D VisualShaderNodeMultiplyAddOpType = 1
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector3D VisualShaderNodeMultiplyAddOpType = 2
  VisualShaderNodeMultiplyAddOpTypeOpTypeVector4D VisualShaderNodeMultiplyAddOpType = 3
  VisualShaderNodeMultiplyAddOpTypeOpTypeMax VisualShaderNodeMultiplyAddOpType = 4
)

func (me *VisualShaderNodeMultiplyAdd) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeMultiplyAdd) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeMultiplyAdd) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeMultiplyAdd) SetOpType(type_ VisualShaderNodeMultiplyAddOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeMultiplyAdd")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1409862380) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeMultiplyAdd) GetOpType() VisualShaderNodeMultiplyAddOpType {
  classNameV := StringNameFromStr("VisualShaderNodeMultiplyAdd")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2823201991) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeMultiplyAddOpType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
