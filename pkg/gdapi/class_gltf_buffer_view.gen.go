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

type ptrsForGLTFBufferViewList struct {
  fnGetBuffer gdc.MethodBindPtr
  fnSetBuffer gdc.MethodBindPtr
  fnGetByteOffset gdc.MethodBindPtr
  fnSetByteOffset gdc.MethodBindPtr
  fnGetByteLength gdc.MethodBindPtr
  fnSetByteLength gdc.MethodBindPtr
  fnGetByteStride gdc.MethodBindPtr
  fnSetByteStride gdc.MethodBindPtr
  fnGetIndices gdc.MethodBindPtr
  fnSetIndices gdc.MethodBindPtr
}

var ptrsForGLTFBufferView ptrsForGLTFBufferViewList

func initGLTFBufferViewPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GLTFBufferView")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_buffer")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnGetBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_buffer")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnSetBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_byte_offset")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnGetByteOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_byte_offset")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnSetByteOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_byte_length")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnGetByteLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_byte_length")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnSetByteLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_byte_stride")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnGetByteStride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_byte_stride")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnSetByteStride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_indices")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnGetIndices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_indices")
    defer methodName.Destroy()
    ptrsForGLTFBufferView.fnSetIndices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
}

type GLTFBufferView struct {
  Resource
}

func (me *GLTFBufferView) BaseClass() string {
  return "GLTFBufferView"
}

func NewGLTFBufferView() *GLTFBufferView {
  str := StringNameFromStr("GLTFBufferView") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFBufferView{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GLTFBufferView) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFBufferView) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFBufferView) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFBufferView) GetBuffer() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnGetBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFBufferView) SetBuffer(buffer int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnSetBuffer), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFBufferView) GetByteOffset() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnGetByteOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFBufferView) SetByteOffset(byte_offset int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnSetByteOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFBufferView) GetByteLength() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnGetByteLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFBufferView) SetByteLength(byte_length int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnSetByteLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFBufferView) GetByteStride() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnGetByteStride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFBufferView) SetByteStride(byte_stride int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_stride) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnSetByteStride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFBufferView) GetIndices() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnGetIndices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFBufferView) SetIndices(indices bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&indices) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFBufferView.fnSetIndices), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
