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

type AnimationNodeBlendSpace1D struct {
  AnimationRootNode
}

func (me *AnimationNodeBlendSpace1D) BaseClass() string {
  return "AnimationNodeBlendSpace1D"
}

func NewAnimationNodeBlendSpace1D() *AnimationNodeBlendSpace1D {
  str := StringNameFromStr("AnimationNodeBlendSpace1D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeBlendSpace1D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *AnimationNodeBlendSpace1D) AddBlendPoint(node AnimationRootNode, pos float64, at_index int64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_blend_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 285050433) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), gdc.ConstTypePtr(&pos) , gdc.ConstTypePtr(&at_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) SetBlendPointPosition(point int64, pos float64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , gdc.ConstTypePtr(&pos) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointPosition(point int64, ) float64 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&point)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace1D) SetBlendPointNode(point int64, node AnimationRootNode, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_point_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4240341528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointNode(point int64, ) AnimationRootNode {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 665599029) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationRootNode()
  pinner.Pin(&point)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace1D) RemoveBlendPoint(point int64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_blend_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) GetBlendPointCount() int64 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace1D) SetMinSpace(min_space float64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_space) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) GetMinSpace() float64 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace1D) SetMaxSpace(max_space float64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_space) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) GetMaxSpace() float64 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace1D) SetSnap(snap float64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&snap) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) GetSnap() float64 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace1D) SetValueLabel(text String, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_value_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) GetValueLabel() String {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_value_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace1D) SetBlendMode(mode AnimationNodeBlendSpace1DBlendMode, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2600869457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) GetBlendMode() AnimationNodeBlendSpace1DBlendMode {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1547667849) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeBlendSpace1DBlendMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationNodeBlendSpace1D) SetUseSync(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace1D) IsUsingSync() bool {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace1D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
