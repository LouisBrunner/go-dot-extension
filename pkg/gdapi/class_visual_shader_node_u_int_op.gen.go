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

type ptrsForVisualShaderNodeUIntOpList struct {
  fnSetOperator gdc.MethodBindPtr
  fnGetOperator gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeUIntOp ptrsForVisualShaderNodeUIntOpList

func initVisualShaderNodeUIntOpPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeUIntOp")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_operator")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeUIntOp.fnSetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3463048345))
  }
  {
    methodName := StringNameFromStr("get_operator")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeUIntOp.fnGetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 256631461))
  }
}

type VisualShaderNodeUIntOp struct {
  VisualShaderNode
}

func (me *VisualShaderNodeUIntOp) BaseClass() string {
  return "VisualShaderNodeUIntOp"
}

func NewVisualShaderNodeUIntOp() *VisualShaderNodeUIntOp {
  str := StringNameFromStr("VisualShaderNodeUIntOp") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeUIntOp{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeUIntOpOperator int
const (
  VisualShaderNodeUIntOpOperatorOpAdd VisualShaderNodeUIntOpOperator = 0
  VisualShaderNodeUIntOpOperatorOpSub VisualShaderNodeUIntOpOperator = 1
  VisualShaderNodeUIntOpOperatorOpMul VisualShaderNodeUIntOpOperator = 2
  VisualShaderNodeUIntOpOperatorOpDiv VisualShaderNodeUIntOpOperator = 3
  VisualShaderNodeUIntOpOperatorOpMod VisualShaderNodeUIntOpOperator = 4
  VisualShaderNodeUIntOpOperatorOpMax VisualShaderNodeUIntOpOperator = 5
  VisualShaderNodeUIntOpOperatorOpMin VisualShaderNodeUIntOpOperator = 6
  VisualShaderNodeUIntOpOperatorOpBitwiseAnd VisualShaderNodeUIntOpOperator = 7
  VisualShaderNodeUIntOpOperatorOpBitwiseOr VisualShaderNodeUIntOpOperator = 8
  VisualShaderNodeUIntOpOperatorOpBitwiseXor VisualShaderNodeUIntOpOperator = 9
  VisualShaderNodeUIntOpOperatorOpBitwiseLeftShift VisualShaderNodeUIntOpOperator = 10
  VisualShaderNodeUIntOpOperatorOpBitwiseRightShift VisualShaderNodeUIntOpOperator = 11
  VisualShaderNodeUIntOpOperatorOpEnumSize VisualShaderNodeUIntOpOperator = 12
)

func (me *VisualShaderNodeUIntOp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeUIntOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeUIntOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeUIntOp) SetOperator(op VisualShaderNodeUIntOpOperator, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeUIntOp.fnSetOperator), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeUIntOp) GetOperator() VisualShaderNodeUIntOpOperator {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeUIntOpOperator

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeUIntOp.fnGetOperator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
