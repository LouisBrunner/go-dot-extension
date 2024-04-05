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

type ptrsForStreamPeerList struct {
  fnPutData gdc.MethodBindPtr
  fnPutPartialData gdc.MethodBindPtr
  fnGetData gdc.MethodBindPtr
  fnGetPartialData gdc.MethodBindPtr
  fnGetAvailableBytes gdc.MethodBindPtr
  fnSetBigEndian gdc.MethodBindPtr
  fnIsBigEndianEnabled gdc.MethodBindPtr
  fnPut8 gdc.MethodBindPtr
  fnPutU8 gdc.MethodBindPtr
  fnPut16 gdc.MethodBindPtr
  fnPutU16 gdc.MethodBindPtr
  fnPut32 gdc.MethodBindPtr
  fnPutU32 gdc.MethodBindPtr
  fnPut64 gdc.MethodBindPtr
  fnPutU64 gdc.MethodBindPtr
  fnPutFloat gdc.MethodBindPtr
  fnPutDouble gdc.MethodBindPtr
  fnPutString gdc.MethodBindPtr
  fnPutUtf8String gdc.MethodBindPtr
  fnPutVar gdc.MethodBindPtr
  fnGet8 gdc.MethodBindPtr
  fnGetU8 gdc.MethodBindPtr
  fnGet16 gdc.MethodBindPtr
  fnGetU16 gdc.MethodBindPtr
  fnGet32 gdc.MethodBindPtr
  fnGetU32 gdc.MethodBindPtr
  fnGet64 gdc.MethodBindPtr
  fnGetU64 gdc.MethodBindPtr
  fnGetFloat gdc.MethodBindPtr
  fnGetDouble gdc.MethodBindPtr
  fnGetString gdc.MethodBindPtr
  fnGetUtf8String gdc.MethodBindPtr
  fnGetVar gdc.MethodBindPtr
}

var ptrsForStreamPeer ptrsForStreamPeerList

func initStreamPeerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("StreamPeer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("put_data")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 680677267))
  }
  {
    methodName := StringNameFromStr("put_partial_data")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutPartialData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2934048347))
  }
  {
    methodName := StringNameFromStr("get_data")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1171824711))
  }
  {
    methodName := StringNameFromStr("get_partial_data")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetPartialData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1171824711))
  }
  {
    methodName := StringNameFromStr("get_available_bytes")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetAvailableBytes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_big_endian")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnSetBigEndian = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_big_endian_enabled")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnIsBigEndianEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("put_8")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPut8 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("put_u8")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutU8 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("put_16")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPut16 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("put_u16")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutU16 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("put_32")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPut32 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("put_u32")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutU32 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("put_64")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPut64 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("put_u64")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutU64 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("put_float")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutFloat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("put_double")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutDouble = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("put_string")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("put_utf8_string")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutUtf8String = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("put_var")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnPutVar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 738511890))
  }
  {
    methodName := StringNameFromStr("get_8")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGet8 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_u8")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetU8 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_16")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGet16 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_u16")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetU16 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_32")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGet32 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_u32")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetU32 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_64")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGet64 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_u64")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetU64 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("get_float")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetFloat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("get_double")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetDouble = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("get_string")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2309358862))
  }
  {
    methodName := StringNameFromStr("get_utf8_string")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetUtf8String = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2309358862))
  }
  {
    methodName := StringNameFromStr("get_var")
    defer methodName.Destroy()
    ptrsForStreamPeer.fnGetVar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3442865206))
  }
}

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
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutData), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *StreamPeer) PutPartialData(data PackedByteArray, ) Array {
  cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutPartialData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetData(bytes int64, ) Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&bytes)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetPartialData(bytes int64, ) Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&bytes)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetPartialData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetAvailableBytes() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetAvailableBytes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) SetBigEndian(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnSetBigEndian), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) IsBigEndianEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnIsBigEndianEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) Put8(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPut8), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutU8(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutU8), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) Put16(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPut16), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutU16(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutU16), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) Put32(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPut32), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutU32(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutU32), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) Put64(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPut64), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutU64(value int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutU64), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutFloat(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutFloat), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutDouble(value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutDouble), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutString(value String, )  {
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutString), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutUtf8String(value String, )  {
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutUtf8String), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) PutVar(value Variant, full_objects bool, )  {
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(&full_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnPutVar), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *StreamPeer) Get8() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGet8), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetU8() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetU8), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) Get16() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGet16), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetU16() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetU16), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) Get32() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGet32), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetU32() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetU32), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) Get64() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGet64), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetU64() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetU64), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetFloat() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetFloat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetDouble() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetDouble), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *StreamPeer) GetString(bytes int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&bytes)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetUtf8String(bytes int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bytes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&bytes)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetUtf8String), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *StreamPeer) GetVar(allow_objects bool, ) Variant {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow_objects) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&allow_objects)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStreamPeer.fnGetVar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
