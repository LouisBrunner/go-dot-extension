// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type RDPipelineColorBlendStateAttachment struct {
  RefCounted
}

func (me *RDPipelineColorBlendStateAttachment) BaseClass() string {
  return "RDPipelineColorBlendStateAttachment"
}

func NewRDPipelineColorBlendStateAttachment() *RDPipelineColorBlendStateAttachment {
  str := StringNameFromStr("RDPipelineColorBlendStateAttachment") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDPipelineColorBlendStateAttachment{}
  obj.SetBaseObject(objPtr)
  return obj
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
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) SetEnableBlend(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_blend")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetEnableBlend() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_blend")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineColorBlendStateAttachment) SetSrcColorBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_src_color_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2251019273) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetSrcColorBlendFactor() RenderingDeviceBlendFactor {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_src_color_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3691288359) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendFactor

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetDstColorBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dst_color_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2251019273) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetDstColorBlendFactor() RenderingDeviceBlendFactor {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dst_color_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3691288359) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendFactor

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetColorBlendOp(p_member RenderingDeviceBlendOperation, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_blend_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3073022720) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetColorBlendOp() RenderingDeviceBlendOperation {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_blend_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1385093561) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendOperation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetSrcAlphaBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_src_alpha_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2251019273) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetSrcAlphaBlendFactor() RenderingDeviceBlendFactor {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_src_alpha_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3691288359) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendFactor

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetDstAlphaBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dst_alpha_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2251019273) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetDstAlphaBlendFactor() RenderingDeviceBlendFactor {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dst_alpha_blend_factor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3691288359) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendFactor

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetAlphaBlendOp(p_member RenderingDeviceBlendOperation, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_blend_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3073022720) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetAlphaBlendOp() RenderingDeviceBlendOperation {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_blend_op")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1385093561) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendOperation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteR(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_write_r")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteR() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_write_r")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteG(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_write_g")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteG() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_write_g")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteB(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_write_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteB() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_write_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteA(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_write_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteA() bool {
  classNameV := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_write_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
