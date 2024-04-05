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

type ptrsForVisualShaderNodeUIntConstantList struct {
  fnSetConstant gdc.MethodBindPtr
  fnGetConstant gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeUIntConstant ptrsForVisualShaderNodeUIntConstantList

func initVisualShaderNodeUIntConstantPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeUIntConstant")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeUIntConstant.fnSetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeUIntConstant.fnGetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type VisualShaderNodeUIntConstant struct {
  VisualShaderNodeConstant
}

func (me *VisualShaderNodeUIntConstant) BaseClass() string {
  return "VisualShaderNodeUIntConstant"
}

func NewVisualShaderNodeUIntConstant() *VisualShaderNodeUIntConstant {
  str := StringNameFromStr("VisualShaderNodeUIntConstant") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeUIntConstant{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeUIntConstant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeUIntConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeUIntConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeUIntConstant) SetConstant(constant int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&constant) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeUIntConstant.fnSetConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeUIntConstant) GetConstant() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeUIntConstant.fnGetConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
