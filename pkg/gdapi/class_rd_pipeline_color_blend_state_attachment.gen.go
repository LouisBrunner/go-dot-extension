// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDPipelineColorBlendStateAttachment struct {
  obj gdc.ObjectPtr
}

func (me *RDPipelineColorBlendStateAttachment) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDPipelineColorBlendStateAttachment) BaseClass() string {
  return "RDPipelineColorBlendStateAttachment"
}



// Enums

func (me *RDPipelineColorBlendStateAttachment) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDPipelineColorBlendStateAttachment) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDPipelineColorBlendStateAttachment) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDPipelineColorBlendStateAttachment) SetAsMix()  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) SetEnableBlend(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_blend")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetEnableBlend() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_blend")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetSrcColorBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_src_color_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2251019273) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetSrcColorBlendFactor() RenderingDeviceBlendFactor {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_src_color_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3691288359) // FIXME: should cache?
  var ret RenderingDeviceBlendFactor
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetDstColorBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dst_color_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2251019273) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetDstColorBlendFactor() RenderingDeviceBlendFactor {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dst_color_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3691288359) // FIXME: should cache?
  var ret RenderingDeviceBlendFactor
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetColorBlendOp(p_member RenderingDeviceBlendOperation, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_blend_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3073022720) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetColorBlendOp() RenderingDeviceBlendOperation {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_blend_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1385093561) // FIXME: should cache?
  var ret RenderingDeviceBlendOperation
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetSrcAlphaBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_src_alpha_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2251019273) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetSrcAlphaBlendFactor() RenderingDeviceBlendFactor {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_src_alpha_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3691288359) // FIXME: should cache?
  var ret RenderingDeviceBlendFactor
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetDstAlphaBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dst_alpha_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2251019273) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetDstAlphaBlendFactor() RenderingDeviceBlendFactor {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dst_alpha_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3691288359) // FIXME: should cache?
  var ret RenderingDeviceBlendFactor
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetAlphaBlendOp(p_member RenderingDeviceBlendOperation, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_blend_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3073022720) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetAlphaBlendOp() RenderingDeviceBlendOperation {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_blend_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1385093561) // FIXME: should cache?
  var ret RenderingDeviceBlendOperation
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteR(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_write_r")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteR() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_write_r")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteG(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_write_g")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteG() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_write_g")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteB(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_write_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteB() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_write_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteA(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_write_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteA() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_write_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *RDPipelineColorBlendStateAttachment) GetPropEnableBlend() bool {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropEnableBlend(value bool) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropSrcColorBlendFactor() int {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropSrcColorBlendFactor(value int) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropDstColorBlendFactor() int {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropDstColorBlendFactor(value int) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropColorBlendOp() int {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropColorBlendOp(value int) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropSrcAlphaBlendFactor() int {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropSrcAlphaBlendFactor(value int) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropDstAlphaBlendFactor() int {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropDstAlphaBlendFactor(value int) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropAlphaBlendOp() int {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropAlphaBlendOp(value int) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropWriteR() bool {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropWriteR(value bool) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropWriteG() bool {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropWriteG(value bool) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropWriteB() bool {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropWriteB(value bool) {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) GetPropWriteA() bool {
  panic("TODO: implement")
}

func (me *RDPipelineColorBlendStateAttachment) SetPropWriteA(value bool) {
  panic("TODO: implement")
}