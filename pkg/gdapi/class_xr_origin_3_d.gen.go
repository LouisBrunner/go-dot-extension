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

type ptrsForXROrigin3DList struct {
	fnSetWorldScale gdc.MethodBindPtr
	fnGetWorldScale gdc.MethodBindPtr
	fnSetCurrent    gdc.MethodBindPtr
	fnIsCurrent     gdc.MethodBindPtr
}

var ptrsForXROrigin3D ptrsForXROrigin3DList

func initXROrigin3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XROrigin3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_world_scale")
		defer methodName.Destroy()
		ptrsForXROrigin3D.fnSetWorldScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_world_scale")
		defer methodName.Destroy()
		ptrsForXROrigin3D.fnGetWorldScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_current")
		defer methodName.Destroy()
		ptrsForXROrigin3D.fnSetCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_current")
		defer methodName.Destroy()
		ptrsForXROrigin3D.fnIsCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type XROrigin3D struct {
	Node3D
}

func (me *XROrigin3D) BaseClass() string {
	return "XROrigin3D"
}

func NewXROrigin3D() *XROrigin3D {
	str := StringNameFromStr("XROrigin3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XROrigin3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *XROrigin3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XROrigin3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XROrigin3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XROrigin3D) SetWorldScale(world_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&world_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXROrigin3D.fnSetWorldScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XROrigin3D) GetWorldScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXROrigin3D.fnGetWorldScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XROrigin3D) SetCurrent(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXROrigin3D.fnSetCurrent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XROrigin3D) IsCurrent() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXROrigin3D.fnIsCurrent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
