// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerExtension struct {
  obj gdc.ObjectPtr
}

func (me *StreamPeerExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StreamPeerExtension) BaseClass() string {
  return "StreamPeerExtension"
}



// Enums

func (me *StreamPeerExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StreamPeerExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *StreamPeerExtension) XGetData(r_buffer *uint8, r_bytes int, r_received *int32, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerExtension) XGetPartialData(r_buffer *uint8, r_bytes int, r_received *int32, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerExtension) XPutData(p_data *uint8, p_bytes int, r_sent *int32, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerExtension) XPutPartialData(p_data *uint8, p_bytes int, r_sent *int32, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerExtension) XGetAvailableBytes()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
