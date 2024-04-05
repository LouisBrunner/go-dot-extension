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

type ptrsForRDFramebufferPassList struct {
  fnSetColorAttachments gdc.MethodBindPtr
  fnGetColorAttachments gdc.MethodBindPtr
  fnSetInputAttachments gdc.MethodBindPtr
  fnGetInputAttachments gdc.MethodBindPtr
  fnSetResolveAttachments gdc.MethodBindPtr
  fnGetResolveAttachments gdc.MethodBindPtr
  fnSetPreserveAttachments gdc.MethodBindPtr
  fnGetPreserveAttachments gdc.MethodBindPtr
  fnSetDepthAttachment gdc.MethodBindPtr
  fnGetDepthAttachment gdc.MethodBindPtr
}

var ptrsForRDFramebufferPass ptrsForRDFramebufferPassList

func initRDFramebufferPassPtrs(iface gdc.Interface) {

  className := StringNameFromStr("RDFramebufferPass")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_color_attachments")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnSetColorAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
  {
    methodName := StringNameFromStr("get_color_attachments")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnGetColorAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
  }
  {
    methodName := StringNameFromStr("set_input_attachments")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnSetInputAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
  {
    methodName := StringNameFromStr("get_input_attachments")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnGetInputAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
  }
  {
    methodName := StringNameFromStr("set_resolve_attachments")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnSetResolveAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
  {
    methodName := StringNameFromStr("get_resolve_attachments")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnGetResolveAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
  }
  {
    methodName := StringNameFromStr("set_preserve_attachments")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnSetPreserveAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
  {
    methodName := StringNameFromStr("get_preserve_attachments")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnGetPreserveAttachments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
  }
  {
    methodName := StringNameFromStr("set_depth_attachment")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnSetDepthAttachment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_depth_attachment")
    defer methodName.Destroy()
    ptrsForRDFramebufferPass.fnGetDepthAttachment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type RDFramebufferPass struct {
  RefCounted
}

func (me *RDFramebufferPass) BaseClass() string {
  return "RDFramebufferPass"
}

func NewRDFramebufferPass() *RDFramebufferPass {
  str := StringNameFromStr("RDFramebufferPass") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDFramebufferPass{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  RDFramebufferPassAttachmentUnused = -1
)

// Enums

func (me *RDFramebufferPass) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDFramebufferPass) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDFramebufferPass) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDFramebufferPass) SetColorAttachments(p_member PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{p_member.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnSetColorAttachments), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetColorAttachments() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnGetColorAttachments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDFramebufferPass) SetInputAttachments(p_member PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{p_member.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnSetInputAttachments), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetInputAttachments() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnGetInputAttachments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDFramebufferPass) SetResolveAttachments(p_member PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{p_member.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnSetResolveAttachments), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetResolveAttachments() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnGetResolveAttachments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDFramebufferPass) SetPreserveAttachments(p_member PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{p_member.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnSetPreserveAttachments), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetPreserveAttachments() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnGetPreserveAttachments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDFramebufferPass) SetDepthAttachment(p_member int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnSetDepthAttachment), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetDepthAttachment() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRDFramebufferPass.fnGetDepthAttachment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
