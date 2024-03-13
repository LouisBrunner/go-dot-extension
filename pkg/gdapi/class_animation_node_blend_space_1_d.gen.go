// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeBlendSpace1D struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeBlendSpace1D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeBlendSpace1D) BaseClass() string {
  return "AnimationNodeBlendSpace1D"
}



// Enums

type AnimationNodeBlendSpace1DBlendMode int
const (
  AnimationNodeBlendSpace1DBlendModeBlendModeInterpolated AnimationNodeBlendSpace1DBlendMode = 0
  AnimationNodeBlendSpace1DBlendModeBlendModeDiscrete AnimationNodeBlendSpace1DBlendMode = 1
  AnimationNodeBlendSpace1DBlendModeBlendModeDiscreteCarry AnimationNodeBlendSpace1DBlendMode = 2
)

func (me *AnimationNodeBlendSpace1D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeBlendSpace1D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlendSpace1D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeBlendSpace1D) AddBlendPoint(node AnimationRootNode, pos float32, at_index int, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_blend_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 285050433) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(&pos), gdc.ConstTypePtr(&at_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) SetBlendPointPosition(point int, pos float32, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), gdc.ConstTypePtr(&pos), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointPosition(point int, ) float32 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace1D) SetBlendPointNode(point int, node AnimationRootNode, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_point_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4240341528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointNode(point int, ) AnimationRootNode {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 665599029) // FIXME: should cache?
  var ret AnimationRootNode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace1D) RemoveBlendPoint(point int, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_blend_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointCount() int {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace1D) SetMinSpace(min_space float32, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_space), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) GetMinSpace() float32 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace1D) SetMaxSpace(max_space float32, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_space), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) GetMaxSpace() float32 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace1D) SetSnap(snap float32, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&snap), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) GetSnap() float32 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace1D) SetValueLabel(text String, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_value_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) GetValueLabel() String {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_value_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace1D) SetBlendMode(mode AnimationNodeBlendSpace1DBlendMode, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2600869457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) GetBlendMode() AnimationNodeBlendSpace1DBlendMode {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1547667849) // FIXME: should cache?
  var ret AnimationNodeBlendSpace1DBlendMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace1D) SetUseSync(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace1D) IsUsingSync() bool {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
