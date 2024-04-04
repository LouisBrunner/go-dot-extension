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

type VisualShaderNodeVec4Constant struct {
  VisualShaderNodeConstant
}

func (me *VisualShaderNodeVec4Constant) BaseClass() string {
  return "VisualShaderNodeVec4Constant"
}

func NewVisualShaderNodeVec4Constant() *VisualShaderNodeVec4Constant {
  str := StringNameFromStr("VisualShaderNodeVec4Constant") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeVec4Constant{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeVec4Constant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVec4Constant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec4Constant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeVec4Constant) SetConstant(constant Quaternion, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVec4Constant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1727505552) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{constant.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeVec4Constant) GetConstant() Quaternion {
  classNameV := StringNameFromStr("VisualShaderNodeVec4Constant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1222331677) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewQuaternion()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
