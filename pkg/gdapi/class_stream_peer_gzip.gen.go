// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerGZIP struct {
  obj gdc.ObjectPtr
}

func (me *StreamPeerGZIP) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StreamPeerGZIP) BaseClass() string {
  return "StreamPeerGZIP"
}



// Enums

func (me *StreamPeerGZIP) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StreamPeerGZIP) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *StreamPeerGZIP) StartCompression(use_deflate bool, buffer_size int, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerGZIP) StartDecompression(use_deflate bool, buffer_size int, )  {
  panic("TODO: implement")
}

func  (me *StreamPeerGZIP) Finish()  {
  panic("TODO: implement")
}

func  (me *StreamPeerGZIP) Clear()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
