// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDPipelineColorBlendState struct {
  RefCounted
}

func (me *RDPipelineColorBlendState) BaseClass() string {
  return "RDPipelineColorBlendState"
}

func NewRDPipelineColorBlendState() *RDPipelineColorBlendState {
  str := StringNameFromStr("RDPipelineColorBlendState") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDPipelineColorBlendState{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RDPipelineColorBlendState) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDPipelineColorBlendState) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDPipelineColorBlendState) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDPipelineColorBlendState) SetEnableLogicOp(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_logic_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendState) GetEnableLogicOp() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_logic_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineColorBlendState) SetLogicOp(p_member RenderingDeviceLogicOperation, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_logic_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3610841058) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendState) GetLogicOp() RenderingDeviceLogicOperation {
  classNameV := StringNameFromStr("RDPipelineColorBlendState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_logic_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 988254690) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret RenderingDeviceLogicOperation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendState) SetBlendConstant(p_member Color, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(p_member.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendState) GetBlendConstant() Color {
  classNameV := StringNameFromStr("RDPipelineColorBlendState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDPipelineColorBlendState) SetAttachments(attachments []RDPipelineColorBlendStateAttachment, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&attachments), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendState) GetAttachments() []RDPipelineColorBlendStateAttachment {
  classNameV := StringNameFromStr("RDPipelineColorBlendState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[RDPipelineColorBlendStateAttachment](ret)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
