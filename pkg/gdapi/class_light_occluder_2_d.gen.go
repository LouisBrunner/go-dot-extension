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

type ptrsForLightOccluder2DList struct {
	fnSetOccluderPolygon   gdc.MethodBindPtr
	fnGetOccluderPolygon   gdc.MethodBindPtr
	fnSetOccluderLightMask gdc.MethodBindPtr
	fnGetOccluderLightMask gdc.MethodBindPtr
	fnSetAsSdfCollision    gdc.MethodBindPtr
	fnIsSetAsSdfCollision  gdc.MethodBindPtr
}

var ptrsForLightOccluder2D ptrsForLightOccluder2DList

func initLightOccluder2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("LightOccluder2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_occluder_polygon")
		defer methodName.Destroy()
		ptrsForLightOccluder2D.fnSetOccluderPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3258315893))
	}
	{
		methodName := StringNameFromStr("get_occluder_polygon")
		defer methodName.Destroy()
		ptrsForLightOccluder2D.fnGetOccluderPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3962317075))
	}
	{
		methodName := StringNameFromStr("set_occluder_light_mask")
		defer methodName.Destroy()
		ptrsForLightOccluder2D.fnSetOccluderLightMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_occluder_light_mask")
		defer methodName.Destroy()
		ptrsForLightOccluder2D.fnGetOccluderLightMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_as_sdf_collision")
		defer methodName.Destroy()
		ptrsForLightOccluder2D.fnSetAsSdfCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_set_as_sdf_collision")
		defer methodName.Destroy()
		ptrsForLightOccluder2D.fnIsSetAsSdfCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type LightOccluder2D struct {
	Node2D
}

func (me *LightOccluder2D) BaseClass() string {
	return "LightOccluder2D"
}

func NewLightOccluder2D() *LightOccluder2D {
	str := StringNameFromStr("LightOccluder2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &LightOccluder2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *LightOccluder2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *LightOccluder2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *LightOccluder2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *LightOccluder2D) SetOccluderPolygon(polygon OccluderPolygon2D) {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightOccluder2D.fnSetOccluderPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightOccluder2D) GetOccluderPolygon() OccluderPolygon2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewOccluderPolygon2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightOccluder2D.fnGetOccluderPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *LightOccluder2D) SetOccluderLightMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightOccluder2D.fnSetOccluderLightMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightOccluder2D) GetOccluderLightMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightOccluder2D.fnGetOccluderLightMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *LightOccluder2D) SetAsSdfCollision(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightOccluder2D.fnSetAsSdfCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *LightOccluder2D) IsSetAsSdfCollision() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLightOccluder2D.fnIsSetAsSdfCollision), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
