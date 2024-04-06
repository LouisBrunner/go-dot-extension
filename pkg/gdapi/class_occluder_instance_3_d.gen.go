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

type ptrsForOccluderInstance3DList struct {
	fnSetBakeMask                   gdc.MethodBindPtr
	fnGetBakeMask                   gdc.MethodBindPtr
	fnSetBakeMaskValue              gdc.MethodBindPtr
	fnGetBakeMaskValue              gdc.MethodBindPtr
	fnSetBakeSimplificationDistance gdc.MethodBindPtr
	fnGetBakeSimplificationDistance gdc.MethodBindPtr
	fnSetOccluder                   gdc.MethodBindPtr
	fnGetOccluder                   gdc.MethodBindPtr
}

var ptrsForOccluderInstance3D ptrsForOccluderInstance3DList

func initOccluderInstance3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OccluderInstance3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_bake_mask")
		defer methodName.Destroy()
		ptrsForOccluderInstance3D.fnSetBakeMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_bake_mask")
		defer methodName.Destroy()
		ptrsForOccluderInstance3D.fnGetBakeMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_bake_mask_value")
		defer methodName.Destroy()
		ptrsForOccluderInstance3D.fnSetBakeMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_bake_mask_value")
		defer methodName.Destroy()
		ptrsForOccluderInstance3D.fnGetBakeMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_bake_simplification_distance")
		defer methodName.Destroy()
		ptrsForOccluderInstance3D.fnSetBakeSimplificationDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bake_simplification_distance")
		defer methodName.Destroy()
		ptrsForOccluderInstance3D.fnGetBakeSimplificationDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_occluder")
		defer methodName.Destroy()
		ptrsForOccluderInstance3D.fnSetOccluder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1664878165))
	}
	{
		methodName := StringNameFromStr("get_occluder")
		defer methodName.Destroy()
		ptrsForOccluderInstance3D.fnGetOccluder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1696836198))
	}

}

type OccluderInstance3D struct {
	Node3D
}

func (me *OccluderInstance3D) BaseClass() string {
	return "OccluderInstance3D"
}

func NewOccluderInstance3D() *OccluderInstance3D {
	str := StringNameFromStr("OccluderInstance3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OccluderInstance3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OccluderInstance3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OccluderInstance3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OccluderInstance3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OccluderInstance3D) SetBakeMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderInstance3D.fnSetBakeMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OccluderInstance3D) GetBakeMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderInstance3D.fnGetBakeMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OccluderInstance3D) SetBakeMaskValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderInstance3D.fnSetBakeMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OccluderInstance3D) GetBakeMaskValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderInstance3D.fnGetBakeMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OccluderInstance3D) SetBakeSimplificationDistance(simplification_distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&simplification_distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderInstance3D.fnSetBakeSimplificationDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OccluderInstance3D) GetBakeSimplificationDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderInstance3D.fnGetBakeSimplificationDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *OccluderInstance3D) SetOccluder(occluder Occluder3D) {
	cargs := []gdc.ConstTypePtr{occluder.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderInstance3D.fnSetOccluder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OccluderInstance3D) GetOccluder() Occluder3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewOccluder3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOccluderInstance3D.fnGetOccluder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
