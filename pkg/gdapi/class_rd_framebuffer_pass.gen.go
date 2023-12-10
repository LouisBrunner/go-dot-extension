// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDFramebufferPass struct {
  obj gdc.ObjectPtr
}

func (me *RDFramebufferPass) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDFramebufferPass) BaseClass() string {
  return "RDFramebufferPass"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(p_member.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDFramebufferPass) GetColorAttachments() PackedInt32Array {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDFramebufferPass) SetInputAttachments(p_member PackedInt32Array, )  {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(p_member.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDFramebufferPass) GetInputAttachments() PackedInt32Array {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDFramebufferPass) SetResolveAttachments(p_member PackedInt32Array, )  {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resolve_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(p_member.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDFramebufferPass) GetResolveAttachments() PackedInt32Array {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resolve_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDFramebufferPass) SetPreserveAttachments(p_member PackedInt32Array, )  {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_preserve_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(p_member.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDFramebufferPass) GetPreserveAttachments() PackedInt32Array {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_preserve_attachments")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDFramebufferPass) SetDepthAttachment(p_member int, )  {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_attachment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDFramebufferPass) GetDepthAttachment() int {
  classNameV := StringNameFromStr("RDFramebufferPass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_attachment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *RDFramebufferPass) GetPropColorAttachments() PackedInt32Array {
  panic("TODO: implement")
}

func (me *RDFramebufferPass) SetPropColorAttachments(value PackedInt32Array) {
  panic("TODO: implement")
}

func (me *RDFramebufferPass) GetPropInputAttachments() PackedInt32Array {
  panic("TODO: implement")
}

func (me *RDFramebufferPass) SetPropInputAttachments(value PackedInt32Array) {
  panic("TODO: implement")
}

func (me *RDFramebufferPass) GetPropResolveAttachments() PackedInt32Array {
  panic("TODO: implement")
}

func (me *RDFramebufferPass) SetPropResolveAttachments(value PackedInt32Array) {
  panic("TODO: implement")
}

func (me *RDFramebufferPass) GetPropPreserveAttachments() PackedInt32Array {
  panic("TODO: implement")
}

func (me *RDFramebufferPass) SetPropPreserveAttachments(value PackedInt32Array) {
  panic("TODO: implement")
}

func (me *RDFramebufferPass) GetPropDepthAttachment() int {
  panic("TODO: implement")
}

func (me *RDFramebufferPass) SetPropDepthAttachment(value int) {
  panic("TODO: implement")
}