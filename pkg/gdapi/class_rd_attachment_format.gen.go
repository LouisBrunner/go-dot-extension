// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDAttachmentFormat struct {
  obj gdc.ObjectPtr
}

func (me *RDAttachmentFormat) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDAttachmentFormat) BaseClass() string {
  return "RDAttachmentFormat"
}



// Enums

func (me *RDAttachmentFormat) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDAttachmentFormat) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDAttachmentFormat) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RDAttachmentFormat) SetFormat(p_member RenderingDeviceDataFormat, )  {
  panic("TODO: implement")
}

func  (me *RDAttachmentFormat) GetFormat()  {
  panic("TODO: implement")
}

func  (me *RDAttachmentFormat) SetSamples(p_member RenderingDeviceTextureSamples, )  {
  panic("TODO: implement")
}

func  (me *RDAttachmentFormat) GetSamples()  {
  panic("TODO: implement")
}

func  (me *RDAttachmentFormat) SetUsageFlags(p_member int, )  {
  panic("TODO: implement")
}

func  (me *RDAttachmentFormat) GetUsageFlags()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
