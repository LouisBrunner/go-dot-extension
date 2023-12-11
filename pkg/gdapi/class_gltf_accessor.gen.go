// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFAccessor struct {
  obj gdc.ObjectPtr
}

func (me *GLTFAccessor) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFAccessor) BaseClass() string {
  return "GLTFAccessor"
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

func  (me *GLTFAccessor) GetBufferView() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetBufferView(buffer_view int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffer_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_view), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetByteOffset() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_byte_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetByteOffset(byte_offset int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_byte_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&byte_offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetComponentType() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_component_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetComponentType(component_type int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_component_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&component_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetNormalized() bool {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normalized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetNormalized(normalized bool, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normalized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&normalized), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetCount() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetCount(count int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetType() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetType(type_ int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetMin() PackedFloat64Array {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 148677866) // FIXME: should cache?
  var ret PackedFloat64Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetMin(min PackedFloat64Array, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2576592201) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(min.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetMax() PackedFloat64Array {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 148677866) // FIXME: should cache?
  var ret PackedFloat64Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetMax(max PackedFloat64Array, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2576592201) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(max.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetSparseCount() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sparse_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetSparseCount(sparse_count int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sparse_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetSparseIndicesBufferView() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sparse_indices_buffer_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetSparseIndicesBufferView(sparse_indices_buffer_view int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sparse_indices_buffer_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_indices_buffer_view), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetSparseIndicesByteOffset() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sparse_indices_byte_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetSparseIndicesByteOffset(sparse_indices_byte_offset int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sparse_indices_byte_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_indices_byte_offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetSparseIndicesComponentType() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sparse_indices_component_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetSparseIndicesComponentType(sparse_indices_component_type int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sparse_indices_component_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_indices_component_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetSparseValuesBufferView() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sparse_values_buffer_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetSparseValuesBufferView(sparse_values_buffer_view int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sparse_values_buffer_view")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_values_buffer_view), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFAccessor) GetSparseValuesByteOffset() int {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sparse_values_byte_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFAccessor) SetSparseValuesByteOffset(sparse_values_byte_offset int, )  {
  classNameV := StringNameFromStr("GLTFAccessor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sparse_values_byte_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sparse_values_byte_offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
