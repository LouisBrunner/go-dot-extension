// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeBlendSpace2D struct {
  AnimationRootNode
}

func (me *AnimationNodeBlendSpace2D) BaseClass() string {
  return "AnimationNodeBlendSpace2D"
}



// Enums

type AnimationNodeBlendSpace2DBlendMode int
const (
  AnimationNodeBlendSpace2DBlendModeBlendModeInterpolated AnimationNodeBlendSpace2DBlendMode = 0
  AnimationNodeBlendSpace2DBlendModeBlendModeDiscrete AnimationNodeBlendSpace2DBlendMode = 1
  AnimationNodeBlendSpace2DBlendModeBlendModeDiscreteCarry AnimationNodeBlendSpace2DBlendMode = 2
)

func (me *AnimationNodeBlendSpace2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeBlendSpace2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlendSpace2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeBlendSpace2D) AddBlendPoint(node AnimationRootNode, pos Vector2, at_index int, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_blend_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 402261981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&at_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointPosition(point int, pos Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), gdc.ConstTypePtr(pos.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointPosition(point int, ) Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointNode(point int, node AnimationRootNode, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_point_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4240341528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointNode(point int, ) AnimationRootNode {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 665599029) // FIXME: should cache?
  var ret AnimationRootNode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) RemoveBlendPoint(point int, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_blend_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointCount() int {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) AddTriangle(x int, y int, z int, at_index int, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_triangle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 753017335) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&at_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetTrianglePoint(triangle int, point int, ) int {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_triangle_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 50157827) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&triangle), gdc.ConstTypePtr(&point), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) RemoveTriangle(triangle int, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_triangle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&triangle), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetTriangleCount() int {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_triangle_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetMinSpace(min_space Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(min_space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetMinSpace() Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetMaxSpace(max_space Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(max_space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetMaxSpace() Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetSnap(snap Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(snap.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetSnap() Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetXLabel(text String, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_x_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetXLabel() String {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_x_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetYLabel(text String, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_y_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetYLabel() String {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_y_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetAutoTriangles(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_triangles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetAutoTriangles() bool {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_triangles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetBlendMode(mode AnimationNodeBlendSpace2DBlendMode, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 81193520) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) GetBlendMode() AnimationNodeBlendSpace2DBlendMode {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1398433632) // FIXME: should cache?
  var ret AnimationNodeBlendSpace2DBlendMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetUseSync(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeBlendSpace2D) IsUsingSync() bool {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
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

type AnimationNodeBlendSpace2DTrianglesUpdatedSignalFn func()

func (me *AnimationNodeBlendSpace2D) ConnectTrianglesUpdated(subs SignalSubscribers, fn AnimationNodeBlendSpace2DTrianglesUpdatedSignalFn) {
  sig := StringNameFromStr("triangles_updated")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationNodeBlendSpace2D) DisconnectTrianglesUpdated(subs SignalSubscribers, fn AnimationNodeBlendSpace2DTrianglesUpdatedSignalFn) {
  sig := StringNameFromStr("triangles_updated")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
