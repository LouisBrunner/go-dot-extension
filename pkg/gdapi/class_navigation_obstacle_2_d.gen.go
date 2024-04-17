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

type ptrsForNavigationObstacle2DList struct {
	fnGetRid                  gdc.MethodBindPtr
	fnSetAvoidanceEnabled     gdc.MethodBindPtr
	fnGetAvoidanceEnabled     gdc.MethodBindPtr
	fnSetNavigationMap        gdc.MethodBindPtr
	fnGetNavigationMap        gdc.MethodBindPtr
	fnSetRadius               gdc.MethodBindPtr
	fnGetRadius               gdc.MethodBindPtr
	fnSetVelocity             gdc.MethodBindPtr
	fnGetVelocity             gdc.MethodBindPtr
	fnSetVertices             gdc.MethodBindPtr
	fnGetVertices             gdc.MethodBindPtr
	fnSetAvoidanceLayers      gdc.MethodBindPtr
	fnGetAvoidanceLayers      gdc.MethodBindPtr
	fnSetAvoidanceLayerValue  gdc.MethodBindPtr
	fnGetAvoidanceLayerValue  gdc.MethodBindPtr
	fnSetAffectNavigationMesh gdc.MethodBindPtr
	fnGetAffectNavigationMesh gdc.MethodBindPtr
	fnSetCarveNavigationMesh  gdc.MethodBindPtr
	fnGetCarveNavigationMesh  gdc.MethodBindPtr
}

var ptrsForNavigationObstacle2D ptrsForNavigationObstacle2DList

func initNavigationObstacle2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationObstacle2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_avoidance_enabled")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnSetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_avoidance_enabled")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetAvoidanceEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_navigation_map")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnSetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_navigation_map")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetNavigationMap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_velocity")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnSetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_velocity")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_vertices")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnSetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_vertices")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("set_avoidance_layers")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnSetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_avoidance_layers")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetAvoidanceLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_avoidance_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnSetAvoidanceLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_avoidance_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetAvoidanceLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_affect_navigation_mesh")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnSetAffectNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_affect_navigation_mesh")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetAffectNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_carve_navigation_mesh")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnSetCarveNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_carve_navigation_mesh")
		defer methodName.Destroy()
		ptrsForNavigationObstacle2D.fnGetCarveNavigationMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type NavigationObstacle2D struct {
	Node2D
}

func (me *NavigationObstacle2D) BaseClass() string {
	return "NavigationObstacle2D"
}

func NewNavigationObstacle2D() *NavigationObstacle2D {
	str := StringNameFromStr("NavigationObstacle2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationObstacle2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *NavigationObstacle2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationObstacle2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationObstacle2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationObstacle2D) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationObstacle2D) SetAvoidanceEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnSetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationObstacle2D) GetAvoidanceEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetAvoidanceEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationObstacle2D) SetNavigationMap(navigation_map RID) {
	cargs := []gdc.ConstTypePtr{navigation_map.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnSetNavigationMap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationObstacle2D) GetNavigationMap() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetNavigationMap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationObstacle2D) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationObstacle2D) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationObstacle2D) SetVelocity(velocity Vector2) {
	cargs := []gdc.ConstTypePtr{velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnSetVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationObstacle2D) GetVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationObstacle2D) SetVertices(vertices PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnSetVertices), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationObstacle2D) GetVertices() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetVertices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationObstacle2D) SetAvoidanceLayers(layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnSetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationObstacle2D) GetAvoidanceLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetAvoidanceLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationObstacle2D) SetAvoidanceLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnSetAvoidanceLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationObstacle2D) GetAvoidanceLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetAvoidanceLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationObstacle2D) SetAffectNavigationMesh(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnSetAffectNavigationMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationObstacle2D) GetAffectNavigationMesh() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetAffectNavigationMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationObstacle2D) SetCarveNavigationMesh(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnSetCarveNavigationMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationObstacle2D) GetCarveNavigationMesh() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationObstacle2D.fnGetCarveNavigationMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
