// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDPipelineSpecializationConstant struct {
  obj gdc.ObjectPtr
}

func (me *RDPipelineSpecializationConstant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDPipelineSpecializationConstant) BaseClass() string {
  return "RDPipelineSpecializationConstant"
}



// Enums

func (me *RDPipelineSpecializationConstant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDPipelineSpecializationConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDPipelineSpecializationConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDPipelineSpecializationConstant) SetValue(value Variant, )  {
  classNameV := StringNameFromStr("RDPipelineSpecializationConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1114965689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineSpecializationConstant) GetValue() Variant {
  classNameV := StringNameFromStr("RDPipelineSpecializationConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1214101251) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineSpecializationConstant) SetConstantId(constant_id int, )  {
  classNameV := StringNameFromStr("RDPipelineSpecializationConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&constant_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineSpecializationConstant) GetConstantId() int {
  classNameV := StringNameFromStr("RDPipelineSpecializationConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
