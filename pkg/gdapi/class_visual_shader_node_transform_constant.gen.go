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

type ptrsForVisualShaderNodeTransformConstantList struct {
  fnSetConstant gdc.MethodBindPtr
  fnGetConstant gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeTransformConstant ptrsForVisualShaderNodeTransformConstantList

func initVisualShaderNodeTransformConstantPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeTransformConstant")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeTransformConstant.fnSetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
  }
  {
    methodName := StringNameFromStr("get_constant")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeTransformConstant.fnGetConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
  }
}

type VisualShaderNodeTransformConstant struct {
  VisualShaderNodeConstant
}

func (me *VisualShaderNodeTransformConstant) BaseClass() string {
  return "VisualShaderNodeTransformConstant"
}

func NewVisualShaderNodeTransformConstant() *VisualShaderNodeTransformConstant {
  str := StringNameFromStr("VisualShaderNodeTransformConstant") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeTransformConstant{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeTransformConstant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTransformConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeTransformConstant) SetConstant(constant Transform3D, )  {
  cargs := []gdc.ConstTypePtr{constant.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformConstant.fnSetConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeTransformConstant) GetConstant() Transform3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeTransformConstant.fnGetConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
