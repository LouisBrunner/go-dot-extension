// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *RDFramebufferPass) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDFramebufferPass) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RDFramebufferPass) SetColorAttachments(p_member PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *RDFramebufferPass) GetColorAttachments()  {
  panic("TODO: implement")
}

func  (me *RDFramebufferPass) SetInputAttachments(p_member PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *RDFramebufferPass) GetInputAttachments()  {
  panic("TODO: implement")
}

func  (me *RDFramebufferPass) SetResolveAttachments(p_member PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *RDFramebufferPass) GetResolveAttachments()  {
  panic("TODO: implement")
}

func  (me *RDFramebufferPass) SetPreserveAttachments(p_member PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *RDFramebufferPass) GetPreserveAttachments()  {
  panic("TODO: implement")
}

func  (me *RDFramebufferPass) SetDepthAttachment(p_member int, )  {
  panic("TODO: implement")
}

func  (me *RDFramebufferPass) GetDepthAttachment()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
