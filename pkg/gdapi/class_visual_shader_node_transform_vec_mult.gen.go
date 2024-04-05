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

type ptrsForVisualShaderNodeTransformVecMultList struct {
  fnSetOperator gdc.MethodBindPtr
  fnGetOperator gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeTransformVecMult ptrsForVisualShaderNodeTransformVecMultList

func initVisualShaderNodeTransformVecMultPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeTransformVecMult")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_operator")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeTransformVecMult.fnSetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1785665912))
  }
  {
    methodName := StringNameFromStr("get_operator")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeTransformVecMult.fnGetOperator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1622088722))
  }
}

type VisualShaderNodeTransformVecMult struct {
  VisualShaderNode
}

func (me *VisualShaderNodeTransformVecMult) BaseClass() string {
  return "VisualShaderNodeTransformVecMult"
}

func NewVisualShaderNodeTransformVecMult() *VisualShaderNodeTransformVecMult {
  str := StringNameFromStr("VisualShaderNodeTransformVecMult") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeTransformVecMult{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeTransformVecMultOperator int
const (
  VisualShaderNodeTransformVecMultOperatorOpAxb VisualShaderNodeTransformVecMultOperator = 0
  VisualShaderNodeTransformVecMultOperatorOpBxa VisualShaderNodeTransformVecMultOperator = 1
  VisualShaderNodeTransformVecMultOperatorOp3X3Axb VisualShaderNodeTransformVecMultOperator = 2
  VisualShaderNodeTransformVecMultOperatorOp3X3Bxa VisualShaderNodeTransformVecMultOperator = 3
  VisualShaderNodeTransformVecMultOperatorOpMax VisualShaderNodeTransformVecMultOperator = 4
)

func (me *VisualShaderNodeTransformVecMult) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTransformVecMult) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformVecMult) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeTransformVecMult) SetOperator(op VisualShaderNodeTransformVecMultOperator, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformVecMult.fnSetOperator), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeTransformVecMult) GetOperator() VisualShaderNodeTransformVecMultOperator {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeTransformVecMultOperator

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformVecMult.fnGetOperator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
