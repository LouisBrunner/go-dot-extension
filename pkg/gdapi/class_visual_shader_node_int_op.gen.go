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

type ptrsForVisualShaderNodeIntOpList struct {
  fnSetOperator gdc.MethodBindPtr
  fnGetOperator gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeIntOp ptrsForVisualShaderNodeIntOpList

func initVisualShaderNodeIntOpPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeIntOp")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_operator")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeIntOp.fnSetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1677909323))
  }
  {
    methodName := StringNameFromStr("get_operator")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeIntOp.fnGetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1236987913))
  }
}

type VisualShaderNodeIntOp struct {
  VisualShaderNode
}

func (me *VisualShaderNodeIntOp) BaseClass() string {
  return "VisualShaderNodeIntOp"
}

func NewVisualShaderNodeIntOp() *VisualShaderNodeIntOp {
  str := StringNameFromStr("VisualShaderNodeIntOp") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeIntOp{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeIntOpOperator int
const (
  VisualShaderNodeIntOpOperatorOpAdd VisualShaderNodeIntOpOperator = 0
  VisualShaderNodeIntOpOperatorOpSub VisualShaderNodeIntOpOperator = 1
  VisualShaderNodeIntOpOperatorOpMul VisualShaderNodeIntOpOperator = 2
  VisualShaderNodeIntOpOperatorOpDiv VisualShaderNodeIntOpOperator = 3
  VisualShaderNodeIntOpOperatorOpMod VisualShaderNodeIntOpOperator = 4
  VisualShaderNodeIntOpOperatorOpMax VisualShaderNodeIntOpOperator = 5
  VisualShaderNodeIntOpOperatorOpMin VisualShaderNodeIntOpOperator = 6
  VisualShaderNodeIntOpOperatorOpBitwiseAnd VisualShaderNodeIntOpOperator = 7
  VisualShaderNodeIntOpOperatorOpBitwiseOr VisualShaderNodeIntOpOperator = 8
  VisualShaderNodeIntOpOperatorOpBitwiseXor VisualShaderNodeIntOpOperator = 9
  VisualShaderNodeIntOpOperatorOpBitwiseLeftShift VisualShaderNodeIntOpOperator = 10
  VisualShaderNodeIntOpOperatorOpBitwiseRightShift VisualShaderNodeIntOpOperator = 11
  VisualShaderNodeIntOpOperatorOpEnumSize VisualShaderNodeIntOpOperator = 12
)

func (me *VisualShaderNodeIntOp) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeIntOp) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeIntOp) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeIntOp) SetOperator(op VisualShaderNodeIntOpOperator, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntOp.fnSetOperator), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeIntOp) GetOperator() VisualShaderNodeIntOpOperator {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeIntOpOperator

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeIntOp.fnGetOperator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
