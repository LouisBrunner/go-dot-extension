// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Shape3D struct {
  obj gdc.ObjectPtr
}

func (me *Shape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Shape3D) BaseClass() string {
  return "Shape3D"
}



// Enums

func (me *Shape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Shape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Shape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Shape3D) SetCustomSolverBias(bias float32, )  {
  classNameV := StringNameFromStr("Shape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_solver_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Shape3D) GetCustomSolverBias() float32 {
  classNameV := StringNameFromStr("Shape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_solver_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shape3D) SetMargin(margin float32, )  {
  classNameV := StringNameFromStr("Shape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Shape3D) GetMargin() float32 {
  classNameV := StringNameFromStr("Shape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shape3D) GetDebugMesh() ArrayMesh {
  classNameV := StringNameFromStr("Shape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1605880883) // FIXME: should cache?
  var ret ArrayMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Shape3D) GetPropCustomSolverBias() float32 {
  panic("TODO: implement")
}

func (me *Shape3D) SetPropCustomSolverBias(value float32) {
  panic("TODO: implement")
}

func (me *Shape3D) GetPropMargin() float32 {
  panic("TODO: implement")
}

func (me *Shape3D) SetPropMargin(value float32) {
  panic("TODO: implement")
}