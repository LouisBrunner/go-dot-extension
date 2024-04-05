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

type ptrsForVisualShaderNodeVec3ConstantList struct {
  fnSetConstant gdc.MethodBindPtr
  fnGetConstant gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeVec3Constant ptrsForVisualShaderNodeVec3ConstantList

func initVisualShaderNodeVec3ConstantPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeVec3Constant")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeVec3Constant.fnSetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeVec3Constant.fnGetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
}

type VisualShaderNodeVec3Constant struct {
  VisualShaderNodeConstant
}

func (me *VisualShaderNodeVec3Constant) BaseClass() string {
  return "VisualShaderNodeVec3Constant"
}

func NewVisualShaderNodeVec3Constant() *VisualShaderNodeVec3Constant {
  str := StringNameFromStr("VisualShaderNodeVec3Constant") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeVec3Constant{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeVec3Constant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVec3Constant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec3Constant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeVec3Constant) SetConstant(constant Vector3, )  {
  cargs := []gdc.ConstTypePtr{constant.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec3Constant.fnSetConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeVec3Constant) GetConstant() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec3Constant.fnGetConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
