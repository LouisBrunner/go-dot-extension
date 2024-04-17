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

type ptrsForNavigationRegion3DList struct {
	fnGetRid                  gdc.MethodBindPtr
	fnSetNavigationMesh       gdc.MethodBindPtr
	fnGetNavigationMesh       gdc.MethodBindPtr
	fnSetEnabled              gdc.MethodBindPtr
	fnIsEnabled               gdc.MethodBindPtr
	fnSetNavigationMap        gdc.MethodBindPtr
	fnGetNavigationMap        gdc.MethodBindPtr
	fnSetUseEdgeConnections   gdc.MethodBindPtr
	fnGetUseEdgeConnections   gdc.MethodBindPtr
	fnSetNavigationLayers     gdc.MethodBindPtr
	fnGetNavigationLayers     gdc.MethodBindPtr
	fnSetNavigationLayerValue gdc.MethodBindPtr
	fnGetNavigationLayerValue gdc.MethodBindPtr
	fnGetRegionRid            gdc.MethodBindPtr
	fnSetEnterCost            gdc.MethodBindPtr
	fnGetEnterCost            gdc.MethodBindPtr
	fnSetTravelCost           gdc.MethodBindPtr
	fnGetTravelCost           gdc.MethodBindPtr
	fnBakeNavigationMesh      gdc.MethodBindPtr
	fnIsBaking                gdc.MethodBindPtr
}

var ptrsForNavigationRegion3D ptrsForNavigationRegion3DList

func initNavigationRegion3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationRegion3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_navigation_mesh")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnSetNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2923361153))
	}
	{
		methodName := StringNameFromStr("get_navigation_mesh")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnGetNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1468720886))
	}
	{
		methodName := StringNameFromStr("set_enabled")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_enabled")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_navigation_map")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnSetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_navigation_map")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_use_edge_connections")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnSetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_use_edge_connections")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnGetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_navigation_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnSetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_navigation_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnGetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_region_rid")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnGetRegionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnSetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnGetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnSetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnGetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("bake_navigation_mesh")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnBakeNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3216645846))
	}
	{
		methodName := StringNameFromStr("is_baking")
		defer methodName.Destroy()
		ptrsForNavigationRegion3D.fnIsBaking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

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

func (me *NavigationRegion3D) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationRegion3D) SetNavigationMesh(navigation_mesh NavigationMesh) {
	cargs := []gdc.ConstTypePtr{navigation_mesh.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnSetNavigationMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion3D) GetNavigationMesh() NavigationMesh {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNavigationMesh()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnGetNavigationMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationRegion3D) SetEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion3D) IsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion3D) SetNavigationMap(navigation_map RID) {
	cargs := []gdc.ConstTypePtr{navigation_map.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnSetNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion3D) GetNavigationMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationRegion3D) SetUseEdgeConnections(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnSetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion3D) GetUseEdgeConnections() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnGetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion3D) SetNavigationLayers(navigation_layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion3D) GetNavigationLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion3D) SetNavigationLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnSetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion3D) GetNavigationLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnGetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion3D) GetRegionRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnGetRegionRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationRegion3D) SetEnterCost(enter_cost float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnSetEnterCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion3D) GetEnterCost() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnGetEnterCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion3D) SetTravelCost(travel_cost float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&travel_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnSetTravelCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion3D) GetTravelCost() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnGetTravelCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion3D) BakeNavigationMesh(on_thread bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&on_thread)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnBakeNavigationMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion3D) IsBaking() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion3D.fnIsBaking), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
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
