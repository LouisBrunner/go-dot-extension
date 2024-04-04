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

type VisualShaderNodeVectorBase struct {
  VisualShaderNode
}

func (me *VisualShaderNodeVectorBase) BaseClass() string {
  return "VisualShaderNodeVectorBase"
}

func NewVisualShaderNodeVectorBase() *VisualShaderNodeVectorBase {
  str := StringNameFromStr("VisualShaderNodeVectorBase") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeVectorBase{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeVectorBaseOpType int
const (
  VisualShaderNodeVectorBaseOpTypeOpTypeVector2D VisualShaderNodeVectorBaseOpType = 0
  VisualShaderNodeVectorBaseOpTypeOpTypeVector3D VisualShaderNodeVectorBaseOpType = 1
  VisualShaderNodeVectorBaseOpTypeOpTypeVector4D VisualShaderNodeVectorBaseOpType = 2
  VisualShaderNodeVectorBaseOpTypeOpTypeMax VisualShaderNodeVectorBaseOpType = 3
)

func (me *VisualShaderNodeVectorBase) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVectorBase) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVectorBase) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeVectorBase) SetOpType(type_ VisualShaderNodeVectorBaseOpType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVectorBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1692596998) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeVectorBase) GetOpType() VisualShaderNodeVectorBaseOpType {
  classNameV := StringNameFromStr("VisualShaderNodeVectorBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_op_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2568738462) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeVectorBaseOpType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
