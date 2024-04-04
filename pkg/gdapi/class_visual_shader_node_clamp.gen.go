// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeClamp struct {
  VisualShaderNode
}

func (me *VisualShaderNodeClamp) BaseClass() string {
  return "VisualShaderNodeClamp"
}

func NewVisualShaderNodeClamp() *VisualShaderNodeClamp {
  str := StringNameFromStr("VisualShaderNodeClamp") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeClamp{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeClampOpType int
const (
  VisualShaderNodeClampOpTypeOpTypeFloat VisualShaderNodeClampOpType = 0
  VisualShaderNodeClampOpTypeOpTypeInt VisualShaderNodeClampOpType = 1
  VisualShaderNodeClampOpTypeOpTypeUint VisualShaderNodeClampOpType = 2
  VisualShaderNodeClampOpTypeOpTypeVector2D VisualShaderNodeClampOpType = 3
  VisualShaderNodeClampOpTypeOpTypeVector3D VisualShaderNodeClampOpType = 4
  VisualShaderNodeClampOpTypeOpTypeVector4D VisualShaderNodeClampOpType = 5
  VisualShaderNodeClampOpTypeOpTypeMax VisualShaderNodeClampOpType = 6
)

func (me *VisualShaderNodeClamp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeClamp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeClamp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeClamp) SetOpType(op_type VisualShaderNodeClampOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeClamp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 405010749) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeClamp) GetOpType() VisualShaderNodeClampOpType {
  classNameV := StringNameFromStr("VisualShaderNodeClamp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 233276050) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeClampOpType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
