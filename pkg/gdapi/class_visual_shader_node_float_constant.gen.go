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

type ptrsForVisualShaderNodeFloatConstantList struct {
  fnSetConstant gdc.MethodBindPtr
  fnGetConstant gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeFloatConstant ptrsForVisualShaderNodeFloatConstantList

func initVisualShaderNodeFloatConstantPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeFloatConstant")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatConstant.fnSetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeFloatConstant.fnGetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

type VisualShaderNodeFloatConstant struct {
  VisualShaderNodeConstant
}

func (me *VisualShaderNodeFloatConstant) BaseClass() string {
  return "VisualShaderNodeFloatConstant"
}

func NewVisualShaderNodeFloatConstant() *VisualShaderNodeFloatConstant {
  str := StringNameFromStr("VisualShaderNodeFloatConstant") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeFloatConstant{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeFloatConstant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFloatConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFloatConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeFloatConstant) SetConstant(constant float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&constant) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatConstant.fnSetConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeFloatConstant) GetConstant() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeFloatConstant.fnGetConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
