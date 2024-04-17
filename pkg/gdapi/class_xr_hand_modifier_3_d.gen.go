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

type ptrsForXRHandModifier3DList struct {
	fnSetHandTracker gdc.MethodBindPtr
	fnGetHandTracker gdc.MethodBindPtr
	fnSetBoneUpdate  gdc.MethodBindPtr
	fnGetBoneUpdate  gdc.MethodBindPtr
}

var ptrsForXRHandModifier3D ptrsForXRHandModifier3DList

func initXRHandModifier3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRHandModifier3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_hand_tracker")
		defer methodName.Destroy()
		ptrsForXRHandModifier3D.fnSetHandTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_hand_tracker")
		defer methodName.Destroy()
		ptrsForXRHandModifier3D.fnGetHandTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_bone_update")
		defer methodName.Destroy()
		ptrsForXRHandModifier3D.fnSetBoneUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635701455))
	}
	{
		methodName := StringNameFromStr("get_bone_update")
		defer methodName.Destroy()
		ptrsForXRHandModifier3D.fnGetBoneUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2873665691))
	}

}

type XRHandModifier3D struct {
	SkeletonModifier3D
}

func (me *XRHandModifier3D) BaseClass() string {
	return "XRHandModifier3D"
}

func NewXRHandModifier3D() *XRHandModifier3D {
	str := StringNameFromStr("XRHandModifier3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRHandModifier3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type XRHandModifier3DBoneUpdate int

const (
	XRHandModifier3DBoneUpdateBoneUpdateFull         XRHandModifier3DBoneUpdate = 0
	XRHandModifier3DBoneUpdateBoneUpdateRotationOnly XRHandModifier3DBoneUpdate = 1
	XRHandModifier3DBoneUpdateBoneUpdateMax          XRHandModifier3DBoneUpdate = 2
)

func (me *XRHandModifier3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRHandModifier3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRHandModifier3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRHandModifier3D) SetHandTracker(tracker_name StringName) {
	cargs := []gdc.ConstTypePtr{tracker_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandModifier3D.fnSetHandTracker), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandModifier3D) GetHandTracker() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandModifier3D.fnGetHandTracker), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRHandModifier3D) SetBoneUpdate(bone_update XRHandModifier3DBoneUpdate) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_update)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandModifier3D.fnSetBoneUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRHandModifier3D) GetBoneUpdate() XRHandModifier3DBoneUpdate {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret XRHandModifier3DBoneUpdate

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRHandModifier3D.fnGetBoneUpdate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
