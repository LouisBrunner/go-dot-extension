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

type ptrsForNavigationLink2DList struct {
	fnGetRid                  gdc.MethodBindPtr
	fnSetEnabled              gdc.MethodBindPtr
	fnIsEnabled               gdc.MethodBindPtr
	fnSetBidirectional        gdc.MethodBindPtr
	fnIsBidirectional         gdc.MethodBindPtr
	fnSetNavigationLayers     gdc.MethodBindPtr
	fnGetNavigationLayers     gdc.MethodBindPtr
	fnSetNavigationLayerValue gdc.MethodBindPtr
	fnGetNavigationLayerValue gdc.MethodBindPtr
	fnSetStartPosition        gdc.MethodBindPtr
	fnGetStartPosition        gdc.MethodBindPtr
	fnSetEndPosition          gdc.MethodBindPtr
	fnGetEndPosition          gdc.MethodBindPtr
	fnSetGlobalStartPosition  gdc.MethodBindPtr
	fnGetGlobalStartPosition  gdc.MethodBindPtr
	fnSetGlobalEndPosition    gdc.MethodBindPtr
	fnGetGlobalEndPosition    gdc.MethodBindPtr
	fnSetEnterCost            gdc.MethodBindPtr
	fnGetEnterCost            gdc.MethodBindPtr
	fnSetTravelCost           gdc.MethodBindPtr
	fnGetTravelCost           gdc.MethodBindPtr
}

var ptrsForNavigationLink2D ptrsForNavigationLink2DList

func initNavigationLink2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("NavigationLink2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_rid")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnGetRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_enabled")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_enabled")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_bidirectional")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetBidirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_bidirectional")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnIsBidirectional = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_navigation_layers")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnGetNavigationLayers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_navigation_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_navigation_layer_value")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnGetNavigationLayerValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_start_position")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_start_position")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnGetStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_end_position")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_end_position")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnGetEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_global_start_position")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetGlobalStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_global_start_position")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnGetGlobalStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_global_end_position")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetGlobalEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_global_end_position")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnGetGlobalEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_enter_cost")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnGetEnterCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnSetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_travel_cost")
		defer methodName.Destroy()
		ptrsForNavigationLink2D.fnGetTravelCost = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type NavigationLink2D struct {
	Node2D
}

func (me *NavigationLink2D) BaseClass() string {
	return "NavigationLink2D"
}

func NewNavigationLink2D() *NavigationLink2D {
	str := StringNameFromStr("NavigationLink2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &NavigationLink2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *NavigationLink2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *NavigationLink2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *NavigationLink2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *NavigationLink2D) GetRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnGetRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationLink2D) SetEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) IsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationLink2D) SetBidirectional(bidirectional bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bidirectional)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetBidirectional), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) IsBidirectional() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnIsBidirectional), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationLink2D) SetNavigationLayers(navigation_layers int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&navigation_layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetNavigationLayers), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) GetNavigationLayers() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnGetNavigationLayers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationLink2D) SetNavigationLayerValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) GetNavigationLayerValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnGetNavigationLayerValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationLink2D) SetStartPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetStartPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) GetStartPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnGetStartPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationLink2D) SetEndPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetEndPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) GetEndPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnGetEndPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationLink2D) SetGlobalStartPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetGlobalStartPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) GetGlobalStartPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnGetGlobalStartPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationLink2D) SetGlobalEndPosition(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetGlobalEndPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) GetGlobalEndPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnGetGlobalEndPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *NavigationLink2D) SetEnterCost(enter_cost float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enter_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetEnterCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) GetEnterCost() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnGetEnterCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *NavigationLink2D) SetTravelCost(travel_cost float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&travel_cost)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnSetTravelCost), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *NavigationLink2D) GetTravelCost() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNavigationLink2D.fnGetTravelCost), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
