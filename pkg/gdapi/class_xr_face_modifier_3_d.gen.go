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

type ptrsForXRFaceModifier3DList struct {
	fnSetFaceTracker gdc.MethodBindPtr
	fnGetFaceTracker gdc.MethodBindPtr
	fnSetTarget      gdc.MethodBindPtr
	fnGetTarget      gdc.MethodBindPtr
}

var ptrsForXRFaceModifier3D ptrsForXRFaceModifier3DList

func initXRFaceModifier3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRFaceModifier3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_face_tracker")
		defer methodName.Destroy()
		ptrsForXRFaceModifier3D.fnSetFaceTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_face_tracker")
		defer methodName.Destroy()
		ptrsForXRFaceModifier3D.fnGetFaceTracker = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_target")
		defer methodName.Destroy()
		ptrsForXRFaceModifier3D.fnSetTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_target")
		defer methodName.Destroy()
		ptrsForXRFaceModifier3D.fnGetTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}

}

type XRFaceModifier3D struct {
	Node3D
}

func (me *XRFaceModifier3D) BaseClass() string {
	return "XRFaceModifier3D"
}

func NewXRFaceModifier3D() *XRFaceModifier3D {
	str := StringNameFromStr("XRFaceModifier3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRFaceModifier3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *XRFaceModifier3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRFaceModifier3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRFaceModifier3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRFaceModifier3D) SetFaceTracker(tracker_name StringName) {
	cargs := []gdc.ConstTypePtr{tracker_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRFaceModifier3D.fnSetFaceTracker), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRFaceModifier3D) GetFaceTracker() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRFaceModifier3D.fnGetFaceTracker), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRFaceModifier3D) SetTarget(target NodePath) {
	cargs := []gdc.ConstTypePtr{target.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRFaceModifier3D.fnSetTarget), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRFaceModifier3D) GetTarget() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRFaceModifier3D.fnGetTarget), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
