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

type StreamPeer struct {
  RefCounted
}

func (me *StreamPeer) BaseClass() string {
  return "StreamPeer"
}

func NewStreamPeer() *StreamPeer {
  str := StringNameFromStr("StreamPeer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &StreamPeer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *StreamPeer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StreamPeer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StreamPeer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *StreamPeer) PutData(data PackedByteArray, ) Error {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeer) PutPartialData(data PackedByteArray, ) Array {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_partial_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2934048347) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetData(bytes int64, ) Array {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1171824711) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&bytes)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetPartialData(bytes int64, ) Array {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_partial_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1171824711) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&bytes)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetAvailableBytes() int64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_available_bytes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) SetBigEndian(enable bool, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_big_endian")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) IsBigEndianEnabled() bool {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_big_endian_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) Put8(value int64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_8")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutU8(value int64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_u8")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) Put16(value int64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_16")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutU16(value int64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_u16")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) Put32(value int64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_32")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutU32(value int64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_u32")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) Put64(value int64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutU64(value int64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_u64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutFloat(value float64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_float")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutDouble(value float64, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_double")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutString(value String, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutUtf8String(value String, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_utf8_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutVar(value Variant, full_objects bool, )  {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("put_var")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 738511890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(&full_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) Get8() int64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_8")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetU8() int64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_u8")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) Get16() int64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_16")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetU16() int64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_u16")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) Get32() int64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_32")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetU32() int64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_u32")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) Get64() int64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetU64() int64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_u64")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetFloat() float64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_float")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetDouble() float64 {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_double")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetString(bytes int64, ) String {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2309358862) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&bytes)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetUtf8String(bytes int64, ) String {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_utf8_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2309358862) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&bytes)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetVar(allow_objects bool, ) Variant {
  classNameV := StringNameFromStr("StreamPeer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_var")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3442865206) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&allow_objects)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
