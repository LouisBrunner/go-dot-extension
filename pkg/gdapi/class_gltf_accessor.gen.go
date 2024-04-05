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

type ptrsForGLTFAccessorList struct {
  fnGetBufferView gdc.MethodBindPtr
  fnSetBufferView gdc.MethodBindPtr
  fnGetByteOffset gdc.MethodBindPtr
  fnSetByteOffset gdc.MethodBindPtr
  fnGetComponentType gdc.MethodBindPtr
  fnSetComponentType gdc.MethodBindPtr
  fnGetNormalized gdc.MethodBindPtr
  fnSetNormalized gdc.MethodBindPtr
  fnGetCount gdc.MethodBindPtr
  fnSetCount gdc.MethodBindPtr
  fnGetType gdc.MethodBindPtr
  fnSetType gdc.MethodBindPtr
  fnGetMin gdc.MethodBindPtr
  fnSetMin gdc.MethodBindPtr
  fnGetMax gdc.MethodBindPtr
  fnSetMax gdc.MethodBindPtr
  fnGetSparseCount gdc.MethodBindPtr
  fnSetSparseCount gdc.MethodBindPtr
  fnGetSparseIndicesBufferView gdc.MethodBindPtr
  fnSetSparseIndicesBufferView gdc.MethodBindPtr
  fnGetSparseIndicesByteOffset gdc.MethodBindPtr
  fnSetSparseIndicesByteOffset gdc.MethodBindPtr
  fnGetSparseIndicesComponentType gdc.MethodBindPtr
  fnSetSparseIndicesComponentType gdc.MethodBindPtr
  fnGetSparseValuesBufferView gdc.MethodBindPtr
  fnSetSparseValuesBufferView gdc.MethodBindPtr
  fnGetSparseValuesByteOffset gdc.MethodBindPtr
  fnSetSparseValuesByteOffset gdc.MethodBindPtr
}

var ptrsForGLTFAccessor ptrsForGLTFAccessorList

func initGLTFAccessorPtrs(iface gdc.Interface) {

  className := StringNameFromStr("GLTFAccessor")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_buffer_view")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetBufferView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_buffer_view")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetBufferView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_byte_offset")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetByteOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_byte_offset")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetByteOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_component_type")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetComponentType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_component_type")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetComponentType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_normalized")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetNormalized = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_normalized")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetNormalized = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_count")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_count")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_type")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_type")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_min")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 148677866))
  }
  {
    methodName := StringNameFromStr("set_min")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2576592201))
  }
  {
    methodName := StringNameFromStr("get_max")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 148677866))
  }
  {
    methodName := StringNameFromStr("set_max")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2576592201))
  }
  {
    methodName := StringNameFromStr("get_sparse_count")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetSparseCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_sparse_count")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetSparseCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_sparse_indices_buffer_view")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetSparseIndicesBufferView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_sparse_indices_buffer_view")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetSparseIndicesBufferView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_sparse_indices_byte_offset")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetSparseIndicesByteOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_sparse_indices_byte_offset")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetSparseIndicesByteOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_sparse_indices_component_type")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetSparseIndicesComponentType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_sparse_indices_component_type")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetSparseIndicesComponentType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_sparse_values_buffer_view")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetSparseValuesBufferView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_sparse_values_buffer_view")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetSparseValuesBufferView = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_sparse_values_byte_offset")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnGetSparseValuesByteOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_sparse_values_byte_offset")
    defer methodName.Destroy()
    ptrsForGLTFAccessor.fnSetSparseValuesByteOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
}

type GLTFAccessor struct {
  Resource
}

func (me *GLTFAccessor) BaseClass() string {
  return "GLTFAccessor"
}

func NewGLTFAccessor() *GLTFAccessor {
  str := StringNameFromStr("GLTFAccessor") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFAccessor{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GLTFAccessor) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFAccessor) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFAccessor) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFAccessor) GetBufferView() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetBufferView), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetBufferView(buffer_view int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_view) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetBufferView), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetByteOffset() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetByteOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetByteOffset(byte_offset int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetByteOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetComponentType() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetComponentType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetComponentType(component_type int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&component_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetComponentType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetNormalized() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetNormalized), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetNormalized(normalized bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalized) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetNormalized), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetCount(count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetType() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetType(type_ int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetMin() PackedFloat64Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat64Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFAccessor) SetMin(min PackedFloat64Array, )  {
  cargs := []gdc.ConstTypePtr{min.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetMin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetMax() PackedFloat64Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat64Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFAccessor) SetMax(max PackedFloat64Array, )  {
  cargs := []gdc.ConstTypePtr{max.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetMax), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetSparseCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetSparseCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetSparseCount(sparse_count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetSparseCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetSparseIndicesBufferView() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetSparseIndicesBufferView), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetSparseIndicesBufferView(sparse_indices_buffer_view int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_indices_buffer_view) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetSparseIndicesBufferView), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetSparseIndicesByteOffset() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetSparseIndicesByteOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetSparseIndicesByteOffset(sparse_indices_byte_offset int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_indices_byte_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetSparseIndicesByteOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetSparseIndicesComponentType() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetSparseIndicesComponentType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetSparseIndicesComponentType(sparse_indices_component_type int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_indices_component_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetSparseIndicesComponentType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetSparseValuesBufferView() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetSparseValuesBufferView), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetSparseValuesBufferView(sparse_values_buffer_view int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_values_buffer_view) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetSparseValuesBufferView), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFAccessor) GetSparseValuesByteOffset() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnGetSparseValuesByteOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFAccessor) SetSparseValuesByteOffset(sparse_values_byte_offset int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_values_byte_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFAccessor.fnSetSparseValuesByteOffset), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
