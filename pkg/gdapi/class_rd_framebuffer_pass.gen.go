// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  RDFramebufferPassAttachmentUnused = "-1" // TODO: construct correctly
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
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{p_member.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetColorAttachments() PackedInt32Array {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDFramebufferPass) SetInputAttachments(p_member PackedInt32Array, )  {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{p_member.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetInputAttachments() PackedInt32Array {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDFramebufferPass) SetResolveAttachments(p_member PackedInt32Array, )  {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resolve_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{p_member.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetResolveAttachments() PackedInt32Array {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolve_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDFramebufferPass) SetPreserveAttachments(p_member PackedInt32Array, )  {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_preserve_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{p_member.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetPreserveAttachments() PackedInt32Array {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_preserve_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDFramebufferPass) SetDepthAttachment(p_member int64, )  {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_attachment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDFramebufferPass) GetDepthAttachment() int64 {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_attachment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
