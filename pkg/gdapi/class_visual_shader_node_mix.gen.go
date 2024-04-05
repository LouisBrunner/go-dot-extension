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

type ptrsForVisualShaderNodeMixList struct {
  fnSetOpType gdc.MethodBindPtr
  fnGetOpType gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeMix ptrsForVisualShaderNodeMixList

func initVisualShaderNodeMixPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeMix")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_op_type")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeMix.fnSetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3397501671))
  }
  {
    methodName := StringNameFromStr("get_op_type")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeMix.fnGetOpType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4013957297))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&op_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeMix.fnSetOpType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeMix) GetOpType() VisualShaderNodeMixOpType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeMixOpType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeMix.fnGetOpType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
