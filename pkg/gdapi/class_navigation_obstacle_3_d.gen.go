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

type ptrsForNavigationObstacle3DList struct {
  fnGetRid gdc.MethodBindPtr
  fnSetAvoidanceEnabled gdc.MethodBindPtr
  fnGetAvoidanceEnabled gdc.MethodBindPtr
  fnSetNavigationMap gdc.MethodBindPtr
  fnGetNavigationMap gdc.MethodBindPtr
  fnSetRadius gdc.MethodBindPtr
  fnGetRadius gdc.MethodBindPtr
  fnSetHeight gdc.MethodBindPtr
  fnGetHeight gdc.MethodBindPtr
  fnSetVelocity gdc.MethodBindPtr
  fnGetVelocity gdc.MethodBindPtr
  fnSetVertices gdc.MethodBindPtr
  fnGetVertices gdc.MethodBindPtr
  fnSetAvoidanceLayers gdc.MethodBindPtr
  fnGetAvoidanceLayers gdc.MethodBindPtr
  fnSetAvoidanceLayerValue gdc.MethodBindPtr
  fnGetAvoidanceLayerValue gdc.MethodBindPtr
  fnSetUse3DAvoidance gdc.MethodBindPtr
  fnGetUse3DAvoidance gdc.MethodBindPtr
}

var ptrsForNavigationObstacle3D ptrsForNavigationObstacle3DList

func initNavigationObstacle3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("NavigationObstacle3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_rid")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
  {
    methodName := StringNameFromStr("set_avoidance_enabled")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnSetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_avoidance_enabled")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_navigation_map")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnSetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("get_navigation_map")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
  {
    methodName := StringNameFromStr("set_radius")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_radius")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_height")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_height")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_velocity")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_velocity")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_vertices")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnSetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334873810))
  }
  {
    methodName := StringNameFromStr("get_vertices")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 497664490))
  }
  {
    methodName := StringNameFromStr("set_avoidance_layers")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnSetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_avoidance_layers")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_avoidance_layer_value")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnSetAvoidanceLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_avoidance_layer_value")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetAvoidanceLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_use_3d_avoidance")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnSetUse3DAvoidance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_use_3d_avoidance")
    defer methodName.Destroy()
    ptrsForNavigationObstacle3D.fnGetUse3DAvoidance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type NavigationObstacle3D struct {
  Node3D
}

func (me *NavigationObstacle3D) BaseClass() string {
  return "NavigationObstacle3D"
}

func NewNavigationObstacle3D() *NavigationObstacle3D {
  str := StringNameFromStr("NavigationObstacle3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationObstacle3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *NavigationObstacle3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationObstacle3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationObstacle3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationObstacle3D) GetRid() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationObstacle3D) SetAvoidanceEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnSetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationObstacle3D) GetAvoidanceEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationObstacle3D) SetNavigationMap(navigation_map RID, )  {
  cargs := []gdc.ConstTypePtr{navigation_map.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnSetNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationObstacle3D) GetNavigationMap() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationObstacle3D) SetRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationObstacle3D) GetRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationObstacle3D) SetHeight(height float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationObstacle3D) GetHeight() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationObstacle3D) SetVelocity(velocity Vector3, )  {
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationObstacle3D) GetVelocity() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationObstacle3D) SetVertices(vertices PackedVector3Array, )  {
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnSetVertices), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationObstacle3D) GetVertices() PackedVector3Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetVertices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationObstacle3D) SetAvoidanceLayers(layers int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnSetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationObstacle3D) GetAvoidanceLayers() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationObstacle3D) SetAvoidanceLayerValue(layer_number int64, value bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnSetAvoidanceLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationObstacle3D) GetAvoidanceLayerValue(layer_number int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetAvoidanceLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationObstacle3D) SetUse3DAvoidance(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnSetUse3DAvoidance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationObstacle3D) GetUse3DAvoidance() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle3D.fnGetUse3DAvoidance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
