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

type ptrsForVisualInstance3DList struct {
	fnXGetAabb                gdc.MethodBindPtr
	fnSetBase                 gdc.MethodBindPtr
	fnGetBase                 gdc.MethodBindPtr
	fnGetInstance             gdc.MethodBindPtr
	fnSetLayerMask            gdc.MethodBindPtr
	fnGetLayerMask            gdc.MethodBindPtr
	fnSetLayerMaskValue       gdc.MethodBindPtr
	fnGetLayerMaskValue       gdc.MethodBindPtr
	fnSetSortingOffset        gdc.MethodBindPtr
	fnGetSortingOffset        gdc.MethodBindPtr
	fnSetSortingUseAabbCenter gdc.MethodBindPtr
	fnIsSortingUseAabbCenter  gdc.MethodBindPtr
	fnGetAabb                 gdc.MethodBindPtr
}

var ptrsForVisualInstance3D ptrsForVisualInstance3DList

func initVisualInstance3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VisualInstance3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_base")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnSetBase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("get_base")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnGetBase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_instance")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnGetInstance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_layer_mask")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnSetLayerMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_layer_mask")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnGetLayerMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_layer_mask_value")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnSetLayerMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_layer_mask_value")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnGetLayerMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_sorting_offset")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnSetSortingOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_sorting_offset")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnGetSortingOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_sorting_use_aabb_center")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnSetSortingUseAabbCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_sorting_use_aabb_center")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnIsSortingUseAabbCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_aabb")
		defer methodName.Destroy()
		ptrsForVisualInstance3D.fnGetAabb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
	}
}

type VisualInstance3D struct {
	Node3D
}

func (me *VisualInstance3D) BaseClass() string {
	return "VisualInstance3D"
}

func NewVisualInstance3D() *VisualInstance3D {
	str := StringNameFromStr("VisualInstance3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VisualInstance3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VisualInstance3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VisualInstance3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VisualInstance3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VisualInstance3D) SetBase(base RID) {
	cargs := []gdc.ConstTypePtr{base.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnSetBase), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualInstance3D) GetBase() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnGetBase), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualInstance3D) GetInstance() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnGetInstance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VisualInstance3D) SetLayerMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnSetLayerMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualInstance3D) GetLayerMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnGetLayerMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualInstance3D) SetLayerMaskValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnSetLayerMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualInstance3D) GetLayerMaskValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnGetLayerMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualInstance3D) SetSortingOffset(offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnSetSortingOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualInstance3D) GetSortingOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnGetSortingOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualInstance3D) SetSortingUseAabbCenter(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnSetSortingUseAabbCenter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VisualInstance3D) IsSortingUseAabbCenter() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnIsSortingUseAabbCenter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VisualInstance3D) GetAabb() AABB {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAABB()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualInstance3D.fnGetAabb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
