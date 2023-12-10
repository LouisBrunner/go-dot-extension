// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StreamPeerBuffer struct {
  obj gdc.ObjectPtr
}

func (me *StreamPeerBuffer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StreamPeerBuffer) BaseClass() string {
  return "StreamPeerBuffer"
}



// Enums

func (me *StreamPeerBuffer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StreamPeerBuffer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StreamPeerBuffer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *StreamPeerBuffer) Seek(position int, )  {
  classNameV := StringNameFromStr("StreamPeerBuffer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("seek")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StreamPeerBuffer) GetSize() int {
  classNameV := StringNameFromStr("StreamPeerBuffer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerBuffer) GetPosition() int {
  classNameV := StringNameFromStr("StreamPeerBuffer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerBuffer) Resize(size int, )  {
  classNameV := StringNameFromStr("StreamPeerBuffer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resize")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StreamPeerBuffer) SetDataArray(data PackedByteArray, )  {
  classNameV := StringNameFromStr("StreamPeerBuffer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_data_array")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2971499966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StreamPeerBuffer) GetDataArray() PackedByteArray {
  classNameV := StringNameFromStr("StreamPeerBuffer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_data_array")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2362200018) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StreamPeerBuffer) Clear()  {
  classNameV := StringNameFromStr("StreamPeerBuffer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StreamPeerBuffer) Duplicate() StreamPeerBuffer {
  classNameV := StringNameFromStr("StreamPeerBuffer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("duplicate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2474064677) // FIXME: should cache?
  var ret StreamPeerBuffer
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *StreamPeerBuffer) GetPropDataArray() PackedByteArray {
  panic("TODO: implement")
}

func (me *StreamPeerBuffer) SetPropDataArray(value PackedByteArray) {
  panic("TODO: implement")
}