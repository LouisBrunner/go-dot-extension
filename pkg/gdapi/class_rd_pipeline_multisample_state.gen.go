// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDPipelineMultisampleState struct {
  RefCounted
}

func (me *RDPipelineMultisampleState) BaseClass() string {
  return "RDPipelineMultisampleState"
}

func NewRDPipelineMultisampleState() *RDPipelineMultisampleState {
  str := StringNameFromStr("RDPipelineMultisampleState") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDPipelineMultisampleState{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RDPipelineMultisampleState) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDPipelineMultisampleState) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDPipelineMultisampleState) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDPipelineMultisampleState) SetSampleCount(p_member RenderingDeviceTextureSamples, )  {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sample_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3774171498) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineMultisampleState) GetSampleCount() RenderingDeviceTextureSamples {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sample_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 407791724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret RenderingDeviceTextureSamples

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineMultisampleState) SetEnableSampleShading(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_sample_shading")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineMultisampleState) GetEnableSampleShading() bool {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_sample_shading")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineMultisampleState) SetMinSampleShading(p_member float64, )  {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_sample_shading")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineMultisampleState) GetMinSampleShading() float64 {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_sample_shading")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineMultisampleState) SetEnableAlphaToCoverage(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_alpha_to_coverage")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineMultisampleState) GetEnableAlphaToCoverage() bool {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_alpha_to_coverage")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineMultisampleState) SetEnableAlphaToOne(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_alpha_to_one")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineMultisampleState) GetEnableAlphaToOne() bool {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_alpha_to_one")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineMultisampleState) SetSampleMasks(masks []int, )  {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sample_masks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&masks), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineMultisampleState) GetSampleMasks() []int {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sample_masks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[int](ret)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
