// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDPipelineMultisampleState struct {
  obj gdc.ObjectPtr
}

func (me *RDPipelineMultisampleState) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDPipelineMultisampleState) BaseClass() string {
  return "RDPipelineMultisampleState"
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
  var ret RenderingDeviceTextureSamples
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineMultisampleState) SetMinSampleShading(p_member float32, )  {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_sample_shading")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineMultisampleState) GetMinSampleShading() float32 {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_sample_shading")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineMultisampleState) SetSampleMasks(masks int, )  {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sample_masks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&masks), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineMultisampleState) GetSampleMasks() int {
  classNameV := StringNameFromStr("RDPipelineMultisampleState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sample_masks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *RDPipelineMultisampleState) GetPropSampleCount() int {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) SetPropSampleCount(value int) {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) GetPropEnableSampleShading() bool {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) SetPropEnableSampleShading(value bool) {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) GetPropMinSampleShading() float32 {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) SetPropMinSampleShading(value float32) {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) GetPropEnableAlphaToCoverage() bool {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) SetPropEnableAlphaToCoverage(value bool) {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) GetPropEnableAlphaToOne() bool {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) SetPropEnableAlphaToOne(value bool) {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) GetPropSampleMasks() int {
  panic("TODO: implement")
}

func (me *RDPipelineMultisampleState) SetPropSampleMasks(value int) {
  panic("TODO: implement")
}