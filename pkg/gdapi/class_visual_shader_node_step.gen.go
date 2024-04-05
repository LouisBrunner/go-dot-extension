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

type ptrsForVisualShaderNodeStepList struct {
  fnSetOpType gdc.MethodBindPtr
  fnGetOpType gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeStep ptrsForVisualShaderNodeStepList

func initVisualShaderNodeStepPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeStep")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_op_type")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeStep.fnSetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 715172489))
  }
  {
    methodName := StringNameFromStr("get_op_type")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeStep.fnGetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3274022781))
  }
}

type VisualShaderNodeStep struct {
  VisualShaderNode
}

func (me *VisualShaderNodeStep) BaseClass() string {
  return "VisualShaderNodeStep"
}

func NewVisualShaderNodeStep() *VisualShaderNodeStep {
  str := StringNameFromStr("VisualShaderNodeStep") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeStep{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeStepOpType int
const (
  VisualShaderNodeStepOpTypeOpTypeScalar VisualShaderNodeStepOpType = 0
  VisualShaderNodeStepOpTypeOpTypeVector2D VisualShaderNodeStepOpType = 1
  VisualShaderNodeStepOpTypeOpTypeVector2DScalar VisualShaderNodeStepOpType = 2
  VisualShaderNodeStepOpTypeOpTypeVector3D VisualShaderNodeStepOpType = 3
  VisualShaderNodeStepOpTypeOpTypeVector3DScalar VisualShaderNodeStepOpType = 4
  VisualShaderNodeStepOpTypeOpTypeVector4D VisualShaderNodeStepOpType = 5
  VisualShaderNodeStepOpTypeOpTypeVector4DScalar VisualShaderNodeStepOpType = 6
  VisualShaderNodeStepOpTypeOpTypeMax VisualShaderNodeStepOpType = 7
)

func (me *VisualShaderNodeStep) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeStep) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeStep) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeStep) SetOpType(op_type VisualShaderNodeStepOpType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeStep.fnSetOpType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeStep) GetOpType() VisualShaderNodeStepOpType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeStepOpType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeStep.fnGetOpType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
