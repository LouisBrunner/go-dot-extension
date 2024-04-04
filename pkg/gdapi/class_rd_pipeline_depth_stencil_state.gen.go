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

type RDPipelineDepthStencilState struct {
  RefCounted
}

func (me *RDPipelineDepthStencilState) BaseClass() string {
  return "RDPipelineDepthStencilState"
}

func NewRDPipelineDepthStencilState() *RDPipelineDepthStencilState {
  str := StringNameFromStr("RDPipelineDepthStencilState") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDPipelineDepthStencilState{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RDPipelineDepthStencilState) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDPipelineDepthStencilState) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDPipelineDepthStencilState) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDPipelineDepthStencilState) SetEnableDepthTest(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_depth_test")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetEnableDepthTest() bool {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_depth_test")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetEnableDepthWrite(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_depth_write")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetEnableDepthWrite() bool {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_depth_write")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetDepthCompareOperator(p_member RenderingDeviceCompareOperator, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_compare_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2573711505) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetDepthCompareOperator() RenderingDeviceCompareOperator {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_compare_operator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 269730778) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceCompareOperator

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineDepthStencilState) SetEnableDepthRange(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_depth_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetEnableDepthRange() bool {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_depth_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetDepthRangeMin(p_member float64, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_range_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetDepthRangeMin() float64 {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_range_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetDepthRangeMax(p_member float64, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth_range_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetDepthRangeMax() float64 {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth_range_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetEnableStencil(p_member bool, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_stencil")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetEnableStencil() bool {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_stencil")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetFrontOpFail(p_member RenderingDeviceStencilOperation, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_front_op_fail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2092799566) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetFrontOpFail() RenderingDeviceStencilOperation {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_front_op_fail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1714732389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceStencilOperation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineDepthStencilState) SetFrontOpPass(p_member RenderingDeviceStencilOperation, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_front_op_pass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2092799566) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetFrontOpPass() RenderingDeviceStencilOperation {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_front_op_pass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1714732389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceStencilOperation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineDepthStencilState) SetFrontOpDepthFail(p_member RenderingDeviceStencilOperation, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_front_op_depth_fail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2092799566) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetFrontOpDepthFail() RenderingDeviceStencilOperation {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_front_op_depth_fail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1714732389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceStencilOperation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineDepthStencilState) SetFrontOpCompare(p_member RenderingDeviceCompareOperator, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_front_op_compare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2573711505) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetFrontOpCompare() RenderingDeviceCompareOperator {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_front_op_compare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 269730778) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceCompareOperator

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineDepthStencilState) SetFrontOpCompareMask(p_member int64, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_front_op_compare_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetFrontOpCompareMask() int64 {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_front_op_compare_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetFrontOpWriteMask(p_member int64, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_front_op_write_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetFrontOpWriteMask() int64 {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_front_op_write_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetFrontOpReference(p_member int64, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_front_op_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetFrontOpReference() int64 {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_front_op_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetBackOpFail(p_member RenderingDeviceStencilOperation, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_back_op_fail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2092799566) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetBackOpFail() RenderingDeviceStencilOperation {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_back_op_fail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1714732389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceStencilOperation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineDepthStencilState) SetBackOpPass(p_member RenderingDeviceStencilOperation, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_back_op_pass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2092799566) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetBackOpPass() RenderingDeviceStencilOperation {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_back_op_pass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1714732389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceStencilOperation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineDepthStencilState) SetBackOpDepthFail(p_member RenderingDeviceStencilOperation, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_back_op_depth_fail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2092799566) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetBackOpDepthFail() RenderingDeviceStencilOperation {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_back_op_depth_fail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1714732389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceStencilOperation

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineDepthStencilState) SetBackOpCompare(p_member RenderingDeviceCompareOperator, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_back_op_compare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2573711505) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetBackOpCompare() RenderingDeviceCompareOperator {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_back_op_compare")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 269730778) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret RenderingDeviceCompareOperator

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *RDPipelineDepthStencilState) SetBackOpCompareMask(p_member int64, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_back_op_compare_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetBackOpCompareMask() int64 {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_back_op_compare_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetBackOpWriteMask(p_member int64, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_back_op_write_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetBackOpWriteMask() int64 {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_back_op_write_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *RDPipelineDepthStencilState) SetBackOpReference(p_member int64, )  {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_back_op_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDPipelineDepthStencilState) GetBackOpReference() int64 {
  classNameV := StringNameFromStr("RDPipelineDepthStencilState")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_back_op_reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
