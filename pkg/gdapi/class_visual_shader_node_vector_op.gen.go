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

type VisualShaderNodeVectorOp struct {
  VisualShaderNodeVectorBase
}

func (me *VisualShaderNodeVectorOp) BaseClass() string {
  return "VisualShaderNodeVectorOp"
}

func NewVisualShaderNodeVectorOp() *VisualShaderNodeVectorOp {
  str := StringNameFromStr("VisualShaderNodeVectorOp") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeVectorOp{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeVectorOpOperator int
const (
  VisualShaderNodeVectorOpOperatorOpAdd VisualShaderNodeVectorOpOperator = 0
  VisualShaderNodeVectorOpOperatorOpSub VisualShaderNodeVectorOpOperator = 1
  VisualShaderNodeVectorOpOperatorOpMul VisualShaderNodeVectorOpOperator = 2
  VisualShaderNodeVectorOpOperatorOpDiv VisualShaderNodeVectorOpOperator = 3
  VisualShaderNodeVectorOpOperatorOpMod VisualShaderNodeVectorOpOperator = 4
  VisualShaderNodeVectorOpOperatorOpPow VisualShaderNodeVectorOpOperator = 5
  VisualShaderNodeVectorOpOperatorOpMax VisualShaderNodeVectorOpOperator = 6
  VisualShaderNodeVectorOpOperatorOpMin VisualShaderNodeVectorOpOperator = 7
  VisualShaderNodeVectorOpOperatorOpCross VisualShaderNodeVectorOpOperator = 8
  VisualShaderNodeVectorOpOperatorOpAtan2 VisualShaderNodeVectorOpOperator = 9
  VisualShaderNodeVectorOpOperatorOpReflect VisualShaderNodeVectorOpOperator = 10
  VisualShaderNodeVectorOpOperatorOpStep VisualShaderNodeVectorOpOperator = 11
  VisualShaderNodeVectorOpOperatorOpEnumSize VisualShaderNodeVectorOpOperator = 12
)

func (me *VisualShaderNodeVectorOp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeVectorOp) SetOperator(op VisualShaderNodeVectorOpOperator, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVectorOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3371507302) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeVectorOp) GetOperator() VisualShaderNodeVectorOpOperator {
  classNameV := StringNameFromStr("VisualShaderNodeVectorOp")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 11793929) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeVectorOpOperator

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
