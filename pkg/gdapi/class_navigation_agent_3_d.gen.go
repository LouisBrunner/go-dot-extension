// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type NavigationAgent3D struct {
  Node
}

func (me *NavigationAgent3D) BaseClass() string {
  return "NavigationAgent3D"
}

func NewNavigationAgent3D() *NavigationAgent3D {
  str := StringNameFromStr("NavigationAgent3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationAgent3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *NavigationAgent3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationAgent3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationAgent3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationAgent3D) GetRid() RID {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetAvoidanceEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceEnabled() bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetPathDesiredDistance(desired_distance float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_desired_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&desired_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathDesiredDistance() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_desired_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetTargetDesiredDistance(desired_distance float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_desired_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&desired_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetTargetDesiredDistance() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_desired_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetRadius(radius float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetRadius() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_radius")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetHeight(height float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetHeight() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetPathHeightOffset(path_height_offset float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_height_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_height_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathHeightOffset() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_height_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetUse3DAvoidance(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_3d_avoidance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetUse3DAvoidance() bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_3d_avoidance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetKeepYVelocity(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_y_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetKeepYVelocity() bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keep_y_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetNeighborDistance(neighbor_distance float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_neighbor_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&neighbor_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetNeighborDistance() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_neighbor_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetMaxNeighbors(max_neighbors int64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_neighbors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_neighbors) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetMaxNeighbors() int64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_neighbors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetTimeHorizonAgents(time_horizon float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_time_horizon_agents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetTimeHorizonAgents() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_horizon_agents")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetTimeHorizonObstacles(time_horizon float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_time_horizon_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_horizon) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetTimeHorizonObstacles() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_horizon_obstacles")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetMaxSpeed(max_speed float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetMaxSpeed() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetPathMaxDistance(max_speed float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_speed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathMaxDistance() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetNavigationLayers(navigation_layers int64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetNavigationLayers() int64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetNavigationLayerValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetNavigationLayerValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetPathfindingAlgorithm(pathfinding_algorithm NavigationPathQueryParameters3DPathfindingAlgorithm, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pathfinding_algorithm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 394560454) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pathfinding_algorithm) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathfindingAlgorithm() NavigationPathQueryParameters3DPathfindingAlgorithm {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pathfinding_algorithm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3398491350) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPathQueryParameters3DPathfindingAlgorithm

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationAgent3D) SetPathPostprocessing(path_postprocessing NavigationPathQueryParameters3DPathPostProcessing, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2267362344) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&path_postprocessing) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathPostprocessing() NavigationPathQueryParameters3DPathPostProcessing {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_postprocessing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3883858360) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPathQueryParameters3DPathPostProcessing

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationAgent3D) SetPathMetadataFlags(flags NavigationPathQueryParameters3DPathMetadataFlags, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_path_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2713846708) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetPathMetadataFlags() NavigationPathQueryParameters3DPathMetadataFlags {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_path_metadata_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1582332802) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret NavigationPathQueryParameters3DPathMetadataFlags

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *NavigationAgent3D) SetNavigationMap(navigation_map RID, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_map.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetNavigationMap() RID {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetTargetPosition(position Vector3, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetTargetPosition() Vector3 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) GetNextPathPosition() Vector3 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_path_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3783033775) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetVelocityForced(velocity Vector3, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity_forced")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) SetVelocity(velocity Vector3, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetVelocity() Vector3 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3783033775) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) DistanceToTarget() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("distance_to_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) GetCurrentNavigationResult() NavigationPathQueryResult3D {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_navigation_result")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 728825684) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNavigationPathQueryResult3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) GetCurrentNavigationPath() PackedVector3Array {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_navigation_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) GetCurrentNavigationPathIndex() int64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_navigation_path_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) IsTargetReached() bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_target_reached")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) IsTargetReachable() bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_target_reachable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) IsNavigationFinished() bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_navigation_finished")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) GetFinalPosition() Vector3 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_final_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3783033775) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetAvoidanceLayers(layers int64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceLayers() int64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetAvoidanceMask(mask int64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceMask() int64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetAvoidanceLayerValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceLayerValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer_number)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetAvoidanceMaskValue(mask_number int64, value bool, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidanceMaskValue(mask_number int64, ) bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_mask_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask_number) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&mask_number)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetAvoidancePriority(priority float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_avoidance_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetAvoidancePriority() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_avoidance_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetDebugEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetDebugEnabled() bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetDebugUseCustom(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_use_custom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetDebugUseCustom() bool {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_use_custom")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationAgent3D) SetDebugPathCustomColor(color Color, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_path_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetDebugPathCustomColor() Color {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_path_custom_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationAgent3D) SetDebugPathCustomPointSize(point_size float64, )  {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_path_custom_point_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationAgent3D) GetDebugPathCustomPointSize() float64 {
  classNameV := StringNameFromStr("NavigationAgent3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_path_custom_point_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type NavigationAgent3DPathChangedSignalFn func()

func (me *NavigationAgent3D) ConnectPathChanged(subs SignalSubscribers, fn NavigationAgent3DPathChangedSignalFn) {
  sig := StringNameFromStr("path_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectPathChanged(subs SignalSubscribers, fn NavigationAgent3DPathChangedSignalFn) {
  sig := StringNameFromStr("path_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DTargetReachedSignalFn func()

func (me *NavigationAgent3D) ConnectTargetReached(subs SignalSubscribers, fn NavigationAgent3DTargetReachedSignalFn) {
  sig := StringNameFromStr("target_reached")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectTargetReached(subs SignalSubscribers, fn NavigationAgent3DTargetReachedSignalFn) {
  sig := StringNameFromStr("target_reached")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DWaypointReachedSignalFn func(details Dictionary, )

func (me *NavigationAgent3D) ConnectWaypointReached(subs SignalSubscribers, fn NavigationAgent3DWaypointReachedSignalFn) {
  sig := StringNameFromStr("waypoint_reached")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectWaypointReached(subs SignalSubscribers, fn NavigationAgent3DWaypointReachedSignalFn) {
  sig := StringNameFromStr("waypoint_reached")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DLinkReachedSignalFn func(details Dictionary, )

func (me *NavigationAgent3D) ConnectLinkReached(subs SignalSubscribers, fn NavigationAgent3DLinkReachedSignalFn) {
  sig := StringNameFromStr("link_reached")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectLinkReached(subs SignalSubscribers, fn NavigationAgent3DLinkReachedSignalFn) {
  sig := StringNameFromStr("link_reached")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DNavigationFinishedSignalFn func()

func (me *NavigationAgent3D) ConnectNavigationFinished(subs SignalSubscribers, fn NavigationAgent3DNavigationFinishedSignalFn) {
  sig := StringNameFromStr("navigation_finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectNavigationFinished(subs SignalSubscribers, fn NavigationAgent3DNavigationFinishedSignalFn) {
  sig := StringNameFromStr("navigation_finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationAgent3DVelocityComputedSignalFn func(safe_velocity Vector3, )

func (me *NavigationAgent3D) ConnectVelocityComputed(subs SignalSubscribers, fn NavigationAgent3DVelocityComputedSignalFn) {
  sig := StringNameFromStr("velocity_computed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationAgent3D) DisconnectVelocityComputed(subs SignalSubscribers, fn NavigationAgent3DVelocityComputedSignalFn) {
  sig := StringNameFromStr("velocity_computed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
