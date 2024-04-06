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

type ptrsForPath3DList struct {
	fnSetCurve gdc.MethodBindPtr
	fnGetCurve gdc.MethodBindPtr
}

var ptrsForPath3D ptrsForPath3DList

func initPath3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Path3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_curve")
		defer methodName.Destroy()
		ptrsForPath3D.fnSetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408955118))
	}
	{
		methodName := StringNameFromStr("get_curve")
		defer methodName.Destroy()
		ptrsForPath3D.fnGetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4244715212))
	}
}

type Path3D struct {
	Node3D
}

func (me *Path3D) BaseClass() string {
	return "Path3D"
}

func NewPath3D() *Path3D {
	str := StringNameFromStr("Path3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Path3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Path3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Path3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Path3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Path3D) SetCurve(curve Curve3D) {
	cargs := []gdc.ConstTypePtr{curve.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPath3D.fnSetCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Path3D) GetCurve() Curve3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCurve3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPath3D.fnGetCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type Path3DCurveChangedSignalFn func()

func (me *Path3D) ConnectCurveChanged(subs SignalSubscribers, fn Path3DCurveChangedSignalFn) {
	sig := StringNameFromStr("curve_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Path3D) DisconnectCurveChanged(subs SignalSubscribers, fn Path3DCurveChangedSignalFn) {
	sig := StringNameFromStr("curve_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
