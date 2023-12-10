// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeResizableBase struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeResizableBase) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeResizableBase) BaseClass() string {
  return "VisualShaderNodeResizableBase"
}



// Enums

func (me *VisualShaderNodeResizableBase) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeResizableBase) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeResizableBase) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeResizableBase) SetSize(size Vector2, )  {
  classNameV := StringNameFromStr("VisualShaderNodeResizableBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeResizableBase) GetSize() Vector2 {
  classNameV := StringNameFromStr("VisualShaderNodeResizableBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeResizableBase) GetPropSize() Vector2 {
  panic("TODO: implement")
}

func (me *VisualShaderNodeResizableBase) SetPropSize(value Vector2) {
  panic("TODO: implement")
}