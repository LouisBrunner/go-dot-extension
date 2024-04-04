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

type AnimationNodeBlendSpace2D struct {
  AnimationRootNode
}

func (me *AnimationNodeBlendSpace2D) BaseClass() string {
  return "AnimationNodeBlendSpace2D"
}

func NewAnimationNodeBlendSpace2D() *AnimationNodeBlendSpace2D {
  str := StringNameFromStr("AnimationNodeBlendSpace2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeBlendSpace2D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *AnimationNodeBlendSpace2D) AddBlendPoint(node AnimationRootNode, pos Vector2, at_index int64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_blend_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 402261981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&at_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointPosition(point int64, pos Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , pos.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointPosition(point int64, ) Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_point_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&point)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointNode(point int64, node AnimationRootNode, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_point_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4240341528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointNode(point int64, ) AnimationRootNode {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
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

func  (me *AnimationNodeBlendSpace2D) RemoveBlendPoint(point int64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_blend_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointCount() int64 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
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

func  (me *AnimationNodeBlendSpace2D) AddTriangle(x int64, y int64, z int64, at_index int64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_triangle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 753017335) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x) , gdc.ConstTypePtr(&y) , gdc.ConstTypePtr(&z) , gdc.ConstTypePtr(&at_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetTrianglePoint(triangle int64, point int64, ) int64 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_triangle_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 50157827) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&triangle) , gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&triangle)
  pinner.Pin(&point)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace2D) RemoveTriangle(triangle int64, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_triangle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&triangle) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetTriangleCount() int64 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_triangle_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace2D) SetMinSpace(min_space Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_min_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{min_space.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetMinSpace() Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_min_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetMaxSpace(max_space Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{max_space.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetMaxSpace() Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetSnap(snap Vector2, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{snap.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetSnap() Vector2 {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_snap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetXLabel(text String, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_x_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetXLabel() String {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_x_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetYLabel(text String, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_y_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetYLabel() String {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_y_label")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetAutoTriangles(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_triangles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetAutoTriangles() bool {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_triangles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace2D) SetBlendMode(mode AnimationNodeBlendSpace2DBlendMode, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 81193520) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetBlendMode() AnimationNodeBlendSpace2DBlendMode {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1398433632) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeBlendSpace2DBlendMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetUseSync(enable bool, )  {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_sync")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) IsUsingSync() bool {
  classNameV := StringNameFromStr("AnimationNodeBlendSpace2D")
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
