// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeComment struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeComment) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeComment) BaseClass() string {
  return "VisualShaderNodeComment"
}



// Enums

func (me *VisualShaderNodeComment) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeComment) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeComment) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeComment) SetTitle(title String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeComment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(title.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeComment) GetTitle() String {
  classNameV := StringNameFromStr("VisualShaderNodeComment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeComment) SetDescription(description String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeComment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_description")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(description.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeComment) GetDescription() String {
  classNameV := StringNameFromStr("VisualShaderNodeComment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_description")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeComment) GetPropTitle() String {
  panic("TODO: implement")
}

func (me *VisualShaderNodeComment) SetPropTitle(value String) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeComment) GetPropDescription() String {
  panic("TODO: implement")
}

func (me *VisualShaderNodeComment) SetPropDescription(value String) {
  panic("TODO: implement")
}