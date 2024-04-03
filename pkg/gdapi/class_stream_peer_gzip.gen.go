// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StreamPeerGZIP struct {
  StreamPeer
}

func (me *StreamPeerGZIP) BaseClass() string {
  return "StreamPeerGZIP"
}

func NewStreamPeerGZIP() *StreamPeerGZIP {
  str := StringNameFromStr("StreamPeerGZIP") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StreamPeerGZIP{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *StreamPeerGZIP) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StreamPeerGZIP) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StreamPeerGZIP) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *StreamPeerGZIP) StartCompression(use_deflate bool, buffer_size int64, ) Error {
  classNameV := StringNameFromStr("StreamPeerGZIP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_compression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 781582770) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_deflate), gdc.ConstTypePtr(&buffer_size), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerGZIP) StartDecompression(use_deflate bool, buffer_size int64, ) Error {
  classNameV := StringNameFromStr("StreamPeerGZIP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_decompression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 781582770) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_deflate), gdc.ConstTypePtr(&buffer_size), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerGZIP) Finish() Error {
  classNameV := StringNameFromStr("StreamPeerGZIP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("finish")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166280745) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeerGZIP) Clear()  {
  classNameV := StringNameFromStr("StreamPeerGZIP")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
