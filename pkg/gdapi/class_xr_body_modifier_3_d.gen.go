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

type ptrsForXRBodyModifier3DList struct {
	fnSetBodyTracker     gdc.MethodBindPtr
	fnGetBodyTracker     gdc.MethodBindPtr
	fnSetBodyUpdate      gdc.MethodBindPtr
	fnGetBodyUpdate      gdc.MethodBindPtr
	fnSetBoneUpdate      gdc.MethodBindPtr
	fnGetBoneUpdate      gdc.MethodBindPtr
	fnSetShowWhenTracked gdc.MethodBindPtr
	fnGetShowWhenTracked gdc.MethodBindPtr
}

var ptrsForXRBodyModifier3D ptrsForXRBodyModifier3DList

func initXRBodyModifier3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRBodyModifier3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_body_tracker")
		defer methodName.Destroy()
		ptrsForXRBodyModifier3D.fnSetBodyTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_body_tracker")
		defer methodName.Destroy()
		ptrsForXRBodyModifier3D.fnGetBodyTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_body_update")
		defer methodName.Destroy()
		ptrsForXRBodyModifier3D.fnSetBodyUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2211199417))
	}
	{
		methodName := StringNameFromStr("get_body_update")
		defer methodName.Destroy()
		ptrsForXRBodyModifier3D.fnGetBodyUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2642335328))
	}
	{
		methodName := StringNameFromStr("set_bone_update")
		defer methodName.Destroy()
		ptrsForXRBodyModifier3D.fnSetBoneUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3356796943))
	}
	{
		methodName := StringNameFromStr("get_bone_update")
		defer methodName.Destroy()
		ptrsForXRBodyModifier3D.fnGetBoneUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1309305964))
	}
	{
		methodName := StringNameFromStr("set_show_when_tracked")
		defer methodName.Destroy()
		ptrsForXRBodyModifier3D.fnSetShowWhenTracked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_show_when_tracked")
		defer methodName.Destroy()
		ptrsForXRBodyModifier3D.fnGetShowWhenTracked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type XRBodyModifier3D struct {
	SkeletonModifier3D
}

func (me *XRBodyModifier3D) BaseClass() string {
	return "XRBodyModifier3D"
}

func NewXRBodyModifier3D() *XRBodyModifier3D {
	str := StringNameFromStr("XRBodyModifier3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRBodyModifier3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type XRBodyModifier3DBodyUpdate int

const (
	XRBodyModifier3DBodyUpdateBodyUpdateUpperBody XRBodyModifier3DBodyUpdate = 1
	XRBodyModifier3DBodyUpdateBodyUpdateLowerBody XRBodyModifier3DBodyUpdate = 2
	XRBodyModifier3DBodyUpdateBodyUpdateHands     XRBodyModifier3DBodyUpdate = 4
)

type XRBodyModifier3DBoneUpdate int

const (
	XRBodyModifier3DBoneUpdateBoneUpdateFull         XRBodyModifier3DBoneUpdate = 0
	XRBodyModifier3DBoneUpdateBoneUpdateRotationOnly XRBodyModifier3DBoneUpdate = 1
	XRBodyModifier3DBoneUpdateBoneUpdateMax          XRBodyModifier3DBoneUpdate = 2
)

func (me *XRBodyModifier3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRBodyModifier3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRBodyModifier3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRBodyModifier3D) SetBodyTracker(tracker_name StringName) {
	cargs := []gdc.ConstTypePtr{tracker_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyModifier3D.fnSetBodyTracker), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRBodyModifier3D) GetBodyTracker() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyModifier3D.fnGetBodyTracker), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRBodyModifier3D) SetBodyUpdate(body_update XRBodyModifier3DBodyUpdate) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&body_update)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyModifier3D.fnSetBodyUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRBodyModifier3D) GetBodyUpdate() XRBodyModifier3DBodyUpdate {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRBodyModifier3DBodyUpdate

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyModifier3D.fnGetBodyUpdate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRBodyModifier3D) SetBoneUpdate(bone_update XRBodyModifier3DBoneUpdate) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_update)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyModifier3D.fnSetBoneUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRBodyModifier3D) GetBoneUpdate() XRBodyModifier3DBoneUpdate {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRBodyModifier3DBoneUpdate

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyModifier3D.fnGetBoneUpdate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *XRBodyModifier3D) SetShowWhenTracked(show bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyModifier3D.fnSetShowWhenTracked), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRBodyModifier3D) GetShowWhenTracked() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRBodyModifier3D.fnGetShowWhenTracked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
