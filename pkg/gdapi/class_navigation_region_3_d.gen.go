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

type NavigationRegion3D struct {
  Node3D
}

func (me *NavigationRegion3D) BaseClass() string {
  return "NavigationRegion3D"
}

func NewNavigationRegion3D() *NavigationRegion3D {
  str := StringNameFromStr("NavigationRegion3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &NavigationRegion3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *NavigationRegion3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *NavigationRegion3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *NavigationRegion3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *NavigationRegion3D) SetNavigationMesh(navigation_mesh NavigationMesh, )  {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2923361153) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationRegion3D) GetNavigationMesh() NavigationMesh {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_navigation_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1468720886) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNavigationMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationRegion3D) SetEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationRegion3D) IsEnabled() bool {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationRegion3D) SetNavigationMap(navigation_map RID, )  {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{navigation_map.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationRegion3D) GetNavigationMap() RID {
  classNameV := StringNameFromStr("NavigationRegion3D")
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

func  (me *NavigationRegion3D) SetUseEdgeConnections(enabled bool, )  {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationRegion3D) GetUseEdgeConnections() bool {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_edge_connections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationRegion3D) SetNavigationLayers(navigation_layers int64, )  {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layers")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationRegion3D) GetNavigationLayers() int64 {
  classNameV := StringNameFromStr("NavigationRegion3D")
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

func  (me *NavigationRegion3D) SetNavigationLayerValue(layer_number int64, value bool, )  {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_navigation_layer_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationRegion3D) GetNavigationLayerValue(layer_number int64, ) bool {
  classNameV := StringNameFromStr("NavigationRegion3D")
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

func  (me *NavigationRegion3D) GetRegionRid() RID {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_region_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *NavigationRegion3D) SetEnterCost(enter_cost float64, )  {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationRegion3D) GetEnterCost() float64 {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enter_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationRegion3D) SetTravelCost(travel_cost float64, )  {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&travel_cost) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *NavigationRegion3D) GetTravelCost() float64 {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_travel_cost")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *NavigationRegion3D) BakeNavigationMesh(on_thread bool, )  {
  classNameV := StringNameFromStr("NavigationRegion3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bake_navigation_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3216645846) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&on_thread) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type NavigationRegion3DNavigationMeshChangedSignalFn func()

func (me *NavigationRegion3D) ConnectNavigationMeshChanged(subs SignalSubscribers, fn NavigationRegion3DNavigationMeshChangedSignalFn) {
  sig := StringNameFromStr("navigation_mesh_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationRegion3D) DisconnectNavigationMeshChanged(subs SignalSubscribers, fn NavigationRegion3DNavigationMeshChangedSignalFn) {
  sig := StringNameFromStr("navigation_mesh_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationRegion3DBakeFinishedSignalFn func()

func (me *NavigationRegion3D) ConnectBakeFinished(subs SignalSubscribers, fn NavigationRegion3DBakeFinishedSignalFn) {
  sig := StringNameFromStr("bake_finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationRegion3D) DisconnectBakeFinished(subs SignalSubscribers, fn NavigationRegion3DBakeFinishedSignalFn) {
  sig := StringNameFromStr("bake_finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
