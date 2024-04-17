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

type ptrsForNavigationRegion2DList struct {
	fnGetRid                  gdc.MethodBindPtr
	fnSetNavigationPolygon    gdc.MethodBindPtr
	fnGetNavigationPolygon    gdc.MethodBindPtr
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
	fnBakeNavigationPolygon   gdc.MethodBindPtr
	fnIsBaking                gdc.MethodBindPtr
}

var ptrsForNavigationRegion2D ptrsForNavigationRegion2DList

func initNavigationRegion2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationRegion2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_navigation_polygon")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnSetNavigationPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1515040758))
	}
	{
		methodName := StringNameFromStr("get_navigation_polygon")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnGetNavigationPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1046532237))
	}
	{
		methodName := StringNameFromStr("set_enabled")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_enabled")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_navigation_map")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnSetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_navigation_map")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_use_edge_connections")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnSetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_use_edge_connections")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnGetUseEdgeConnections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_navigation_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnSetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_navigation_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnGetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_region_rid")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnGetRegionRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnSetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnGetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnSetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnGetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("bake_navigation_polygon")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnBakeNavigationPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3216645846))
	}
	{
		methodName := StringNameFromStr("is_baking")
		defer methodName.Destroy()
		ptrsForNavigationRegion2D.fnIsBaking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type NavigationRegion2D struct {
	Node2D
}

func (me *NavigationRegion2D) BaseClass() string {
	return "NavigationRegion2D"
}

func NewNavigationRegion2D() *NavigationRegion2D {
	str := StringNameFromStr("NavigationRegion2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationRegion2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *NavigationRegion2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationRegion2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationRegion2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationRegion2D) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationRegion2D) SetNavigationPolygon(navigation_polygon NavigationPolygon) {
	cargs := []gdc.ConstTypePtr{navigation_polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnSetNavigationPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion2D) GetNavigationPolygon() NavigationPolygon {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNavigationPolygon()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnGetNavigationPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationRegion2D) SetEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion2D) IsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion2D) SetNavigationMap(navigation_map RID) {
	cargs := []gdc.ConstTypePtr{navigation_map.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnSetNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion2D) GetNavigationMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationRegion2D) SetUseEdgeConnections(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnSetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion2D) GetUseEdgeConnections() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnGetUseEdgeConnections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion2D) SetNavigationLayers(navigation_layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion2D) GetNavigationLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion2D) SetNavigationLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnSetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion2D) GetNavigationLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnGetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion2D) GetRegionRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnGetRegionRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationRegion2D) SetEnterCost(enter_cost float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnSetEnterCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion2D) GetEnterCost() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnGetEnterCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion2D) SetTravelCost(travel_cost float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&travel_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnSetTravelCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion2D) GetTravelCost() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnGetTravelCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationRegion2D) BakeNavigationPolygon(on_thread bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&on_thread)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnBakeNavigationPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationRegion2D) IsBaking() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationRegion2D.fnIsBaking), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type NavigationRegion2DNavigationPolygonChangedSignalFn func()

func (me *NavigationRegion2D) ConnectNavigationPolygonChanged(subs SignalSubscribers, fn NavigationRegion2DNavigationPolygonChangedSignalFn) {
	sig := StringNameFromStr("navigation_polygon_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationRegion2D) DisconnectNavigationPolygonChanged(subs SignalSubscribers, fn NavigationRegion2DNavigationPolygonChangedSignalFn) {
	sig := StringNameFromStr("navigation_polygon_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NavigationRegion2DBakeFinishedSignalFn func()

func (me *NavigationRegion2D) ConnectBakeFinished(subs SignalSubscribers, fn NavigationRegion2DBakeFinishedSignalFn) {
	sig := StringNameFromStr("bake_finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *NavigationRegion2D) DisconnectBakeFinished(subs SignalSubscribers, fn NavigationRegion2DBakeFinishedSignalFn) {
	sig := StringNameFromStr("bake_finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
