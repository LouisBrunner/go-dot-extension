// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeExpression struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeExpression) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeExpression) BaseClass() string {
  return "VisualShaderNodeExpression"
}



// Enums

func (me *VisualShaderNodeExpression) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeExpression) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeExpression) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeExpression) SetExpression(expression String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeExpression")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(expression.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeExpression) GetExpression() String {
  classNameV := StringNameFromStr("VisualShaderNodeExpression")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_expression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeExpression) GetPropExpression() String {
  panic("TODO: implement")
}

func (me *VisualShaderNodeExpression) SetPropExpression(value String) {
  panic("TODO: implement")
}