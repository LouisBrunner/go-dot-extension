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

type VisualShaderNodeColorConstant struct {
  VisualShaderNodeConstant
}

func (me *VisualShaderNodeColorConstant) BaseClass() string {
  return "VisualShaderNodeColorConstant"
}

func NewVisualShaderNodeColorConstant() *VisualShaderNodeColorConstant {
  str := StringNameFromStr("VisualShaderNodeColorConstant") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeColorConstant{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeColorConstant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeColorConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeColorConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeColorConstant) SetConstant(constant Color, )  {
  classNameV := StringNameFromStr("VisualShaderNodeColorConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{constant.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeColorConstant) GetConstant() Color {
  classNameV := StringNameFromStr("VisualShaderNodeColorConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
