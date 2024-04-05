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

type ptrsForAnimationNodeBlendSpace2DList struct {
  fnAddBlendPoint gdc.MethodBindPtr
  fnSetBlendPointPosition gdc.MethodBindPtr
  fnGetBlendPointPosition gdc.MethodBindPtr
  fnSetBlendPointNode gdc.MethodBindPtr
  fnGetBlendPointNode gdc.MethodBindPtr
  fnRemoveBlendPoint gdc.MethodBindPtr
  fnGetBlendPointCount gdc.MethodBindPtr
  fnAddTriangle gdc.MethodBindPtr
  fnGetTrianglePoint gdc.MethodBindPtr
  fnRemoveTriangle gdc.MethodBindPtr
  fnGetTriangleCount gdc.MethodBindPtr
  fnSetMinSpace gdc.MethodBindPtr
  fnGetMinSpace gdc.MethodBindPtr
  fnSetMaxSpace gdc.MethodBindPtr
  fnGetMaxSpace gdc.MethodBindPtr
  fnSetSnap gdc.MethodBindPtr
  fnGetSnap gdc.MethodBindPtr
  fnSetXLabel gdc.MethodBindPtr
  fnGetXLabel gdc.MethodBindPtr
  fnSetYLabel gdc.MethodBindPtr
  fnGetYLabel gdc.MethodBindPtr
  fnSetAutoTriangles gdc.MethodBindPtr
  fnGetAutoTriangles gdc.MethodBindPtr
  fnSetBlendMode gdc.MethodBindPtr
  fnGetBlendMode gdc.MethodBindPtr
  fnSetUseSync gdc.MethodBindPtr
  fnIsUsingSync gdc.MethodBindPtr
}

var ptrsForAnimationNodeBlendSpace2D ptrsForAnimationNodeBlendSpace2DList

func initAnimationNodeBlendSpace2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AnimationNodeBlendSpace2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_blend_point")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnAddBlendPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 402261981))
  }
  {
    methodName := StringNameFromStr("set_blend_point_position")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetBlendPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
  }
  {
    methodName := StringNameFromStr("get_blend_point_position")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetBlendPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
  }
  {
    methodName := StringNameFromStr("set_blend_point_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetBlendPointNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4240341528))
  }
  {
    methodName := StringNameFromStr("get_blend_point_node")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetBlendPointNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 665599029))
  }
  {
    methodName := StringNameFromStr("remove_blend_point")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnRemoveBlendPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_blend_point_count")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetBlendPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("add_triangle")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnAddTriangle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 753017335))
  }
  {
    methodName := StringNameFromStr("get_triangle_point")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetTrianglePoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 50157827))
  }
  {
    methodName := StringNameFromStr("remove_triangle")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnRemoveTriangle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_triangle_count")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetTriangleCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_min_space")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetMinSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_min_space")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetMinSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_max_space")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetMaxSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_max_space")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetMaxSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_snap")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetSnap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_snap")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetSnap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_x_label")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetXLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_x_label")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetXLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_y_label")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetYLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_y_label")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetYLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_auto_triangles")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetAutoTriangles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_auto_triangles")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetAutoTriangles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_blend_mode")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 81193520))
  }
  {
    methodName := StringNameFromStr("get_blend_mode")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnGetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1398433632))
  }
  {
    methodName := StringNameFromStr("set_use_sync")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnSetUseSync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_using_sync")
    defer methodName.Destroy()
    ptrsForAnimationNodeBlendSpace2D.fnIsUsingSync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), pos.AsCTypePtr(), gdc.ConstTypePtr(&at_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnAddBlendPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointPosition(point int64, pos Vector2, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , pos.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetBlendPointPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointPosition(point int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&point)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetBlendPointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetBlendPointNode(point int64, node AnimationRootNode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetBlendPointNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointNode(point int64, ) AnimationRootNode {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationRootNode()
  pinner.Pin(&point)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetBlendPointNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) RemoveBlendPoint(point int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnRemoveBlendPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetBlendPointCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetBlendPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace2D) AddTriangle(x int64, y int64, z int64, at_index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x) , gdc.ConstTypePtr(&y) , gdc.ConstTypePtr(&z) , gdc.ConstTypePtr(&at_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnAddTriangle), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetTrianglePoint(triangle int64, point int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&triangle) , gdc.ConstTypePtr(&point) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&triangle)
  pinner.Pin(&point)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetTrianglePoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace2D) RemoveTriangle(triangle int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&triangle) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnRemoveTriangle), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetTriangleCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetTriangleCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace2D) SetMinSpace(min_space Vector2, )  {
  cargs := []gdc.ConstTypePtr{min_space.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetMinSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetMinSpace() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetMinSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetMaxSpace(max_space Vector2, )  {
  cargs := []gdc.ConstTypePtr{max_space.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetMaxSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetMaxSpace() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetMaxSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetSnap(snap Vector2, )  {
  cargs := []gdc.ConstTypePtr{snap.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetSnap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetSnap() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetSnap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetXLabel(text String, )  {
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetXLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetXLabel() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetXLabel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetYLabel(text String, )  {
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetYLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetYLabel() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetYLabel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeBlendSpace2D) SetAutoTriangles(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetAutoTriangles), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetAutoTriangles() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetAutoTriangles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeBlendSpace2D) SetBlendMode(mode AnimationNodeBlendSpace2DBlendMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetBlendMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) GetBlendMode() AnimationNodeBlendSpace2DBlendMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeBlendSpace2DBlendMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnGetBlendMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationNodeBlendSpace2D) SetUseSync(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnSetUseSync), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeBlendSpace2D) IsUsingSync() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace2D.fnIsUsingSync), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
