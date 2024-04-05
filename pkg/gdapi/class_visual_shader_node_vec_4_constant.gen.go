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

type ptrsForVisualShaderNodeVec4ConstantList struct {
  fnSetConstant gdc.MethodBindPtr
  fnGetConstant gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeVec4Constant ptrsForVisualShaderNodeVec4ConstantList

func initVisualShaderNodeVec4ConstantPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeVec4Constant")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeVec4Constant.fnSetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1727505552))
  }
  {
    methodName := StringNameFromStr("get_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeVec4Constant.fnGetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1222331677))
  }
}

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
  cargs := []gdc.ConstTypePtr{constant.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec4Constant.fnSetConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeVec4Constant) GetConstant() Quaternion {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewQuaternion()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeVec4Constant.fnGetConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
