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

type ptrsForRDPipelineColorBlendStateAttachmentList struct {
  fnSetAsMix gdc.MethodBindPtr
  fnSetEnableBlend gdc.MethodBindPtr
  fnGetEnableBlend gdc.MethodBindPtr
  fnSetSrcColorBlendFactor gdc.MethodBindPtr
  fnGetSrcColorBlendFactor gdc.MethodBindPtr
  fnSetDstColorBlendFactor gdc.MethodBindPtr
  fnGetDstColorBlendFactor gdc.MethodBindPtr
  fnSetColorBlendOp gdc.MethodBindPtr
  fnGetColorBlendOp gdc.MethodBindPtr
  fnSetSrcAlphaBlendFactor gdc.MethodBindPtr
  fnGetSrcAlphaBlendFactor gdc.MethodBindPtr
  fnSetDstAlphaBlendFactor gdc.MethodBindPtr
  fnGetDstAlphaBlendFactor gdc.MethodBindPtr
  fnSetAlphaBlendOp gdc.MethodBindPtr
  fnGetAlphaBlendOp gdc.MethodBindPtr
  fnSetWriteR gdc.MethodBindPtr
  fnGetWriteR gdc.MethodBindPtr
  fnSetWriteG gdc.MethodBindPtr
  fnGetWriteG gdc.MethodBindPtr
  fnSetWriteB gdc.MethodBindPtr
  fnGetWriteB gdc.MethodBindPtr
  fnSetWriteA gdc.MethodBindPtr
  fnGetWriteA gdc.MethodBindPtr
}

var ptrsForRDPipelineColorBlendStateAttachment ptrsForRDPipelineColorBlendStateAttachmentList

func initRDPipelineColorBlendStateAttachmentPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RDPipelineColorBlendStateAttachment")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_as_mix")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetAsMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_enable_blend")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetEnableBlend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_enable_blend")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetEnableBlend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_src_color_blend_factor")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetSrcColorBlendFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2251019273))
  }
  {
    methodName := StringNameFromStr("get_src_color_blend_factor")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetSrcColorBlendFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3691288359))
  }
  {
    methodName := StringNameFromStr("set_dst_color_blend_factor")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetDstColorBlendFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2251019273))
  }
  {
    methodName := StringNameFromStr("get_dst_color_blend_factor")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetDstColorBlendFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3691288359))
  }
  {
    methodName := StringNameFromStr("set_color_blend_op")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetColorBlendOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3073022720))
  }
  {
    methodName := StringNameFromStr("get_color_blend_op")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetColorBlendOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1385093561))
  }
  {
    methodName := StringNameFromStr("set_src_alpha_blend_factor")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetSrcAlphaBlendFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2251019273))
  }
  {
    methodName := StringNameFromStr("get_src_alpha_blend_factor")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetSrcAlphaBlendFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3691288359))
  }
  {
    methodName := StringNameFromStr("set_dst_alpha_blend_factor")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetDstAlphaBlendFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2251019273))
  }
  {
    methodName := StringNameFromStr("get_dst_alpha_blend_factor")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetDstAlphaBlendFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3691288359))
  }
  {
    methodName := StringNameFromStr("set_alpha_blend_op")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetAlphaBlendOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3073022720))
  }
  {
    methodName := StringNameFromStr("get_alpha_blend_op")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetAlphaBlendOp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1385093561))
  }
  {
    methodName := StringNameFromStr("set_write_r")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetWriteR = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_write_r")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetWriteR = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_write_g")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetWriteG = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_write_g")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetWriteG = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_write_b")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetWriteB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_write_b")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetWriteB = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_write_a")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnSetWriteA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_write_a")
    defer methodName.Destroy()
    ptrsForRDPipelineColorBlendStateAttachment.fnGetWriteA = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetAsMix), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) SetEnableBlend(p_member bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetEnableBlend), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetEnableBlend() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetEnableBlend), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineColorBlendStateAttachment) SetSrcColorBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetSrcColorBlendFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetSrcColorBlendFactor() RenderingDeviceBlendFactor {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendFactor

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetSrcColorBlendFactor), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetDstColorBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetDstColorBlendFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetDstColorBlendFactor() RenderingDeviceBlendFactor {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendFactor

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetDstColorBlendFactor), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetColorBlendOp(p_member RenderingDeviceBlendOperation, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetColorBlendOp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetColorBlendOp() RenderingDeviceBlendOperation {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendOperation

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetColorBlendOp), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetSrcAlphaBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetSrcAlphaBlendFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetSrcAlphaBlendFactor() RenderingDeviceBlendFactor {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendFactor

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetSrcAlphaBlendFactor), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetDstAlphaBlendFactor(p_member RenderingDeviceBlendFactor, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetDstAlphaBlendFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetDstAlphaBlendFactor() RenderingDeviceBlendFactor {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendFactor

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetDstAlphaBlendFactor), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetAlphaBlendOp(p_member RenderingDeviceBlendOperation, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetAlphaBlendOp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetAlphaBlendOp() RenderingDeviceBlendOperation {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceBlendOperation

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetAlphaBlendOp), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteR(p_member bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetWriteR), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteR() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetWriteR), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteG(p_member bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetWriteG), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteG() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetWriteG), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteB(p_member bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetWriteB), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteB() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetWriteB), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineColorBlendStateAttachment) SetWriteA(p_member bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnSetWriteA), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineColorBlendStateAttachment) GetWriteA() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDPipelineColorBlendStateAttachment.fnGetWriteA), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
