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

type ptrsForMobileVRInterfaceList struct {
	fnSetEyeHeight     gdc.MethodBindPtr
	fnGetEyeHeight     gdc.MethodBindPtr
	fnSetIod           gdc.MethodBindPtr
	fnGetIod           gdc.MethodBindPtr
	fnSetDisplayWidth  gdc.MethodBindPtr
	fnGetDisplayWidth  gdc.MethodBindPtr
	fnSetDisplayToLens gdc.MethodBindPtr
	fnGetDisplayToLens gdc.MethodBindPtr
	fnSetOversample    gdc.MethodBindPtr
	fnGetOversample    gdc.MethodBindPtr
	fnSetK1            gdc.MethodBindPtr
	fnGetK1            gdc.MethodBindPtr
	fnSetK2            gdc.MethodBindPtr
	fnGetK2            gdc.MethodBindPtr
}

var ptrsForMobileVRInterface ptrsForMobileVRInterfaceList

func initMobileVRInterfacePtrs(iface gdc.Interface) {

	className := StringNameFromStr("MobileVRInterface")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_eye_height")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnSetEyeHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_eye_height")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnGetEyeHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_iod")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnSetIod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_iod")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnGetIod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_display_width")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnSetDisplayWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_display_width")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnGetDisplayWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_display_to_lens")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnSetDisplayToLens = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_display_to_lens")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnGetDisplayToLens = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_oversample")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnSetOversample = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_oversample")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnGetOversample = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_k1")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnSetK1 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_k1")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnGetK1 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_k2")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnSetK2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_k2")
		defer methodName.Destroy()
		ptrsForMobileVRInterface.fnGetK2 = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type MobileVRInterface struct {
	XRInterface
}

func (me *MobileVRInterface) BaseClass() string {
	return "MobileVRInterface"
}

func NewMobileVRInterface() *MobileVRInterface {
	str := StringNameFromStr("MobileVRInterface") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &MobileVRInterface{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *MobileVRInterface) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *MobileVRInterface) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *MobileVRInterface) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *MobileVRInterface) SetEyeHeight(eye_height float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&eye_height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnSetEyeHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MobileVRInterface) GetEyeHeight() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnGetEyeHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MobileVRInterface) SetIod(iod float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&iod)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnSetIod), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MobileVRInterface) GetIod() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnGetIod), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MobileVRInterface) SetDisplayWidth(display_width float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&display_width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnSetDisplayWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MobileVRInterface) GetDisplayWidth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnGetDisplayWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MobileVRInterface) SetDisplayToLens(display_to_lens float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&display_to_lens)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnSetDisplayToLens), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MobileVRInterface) GetDisplayToLens() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnGetDisplayToLens), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MobileVRInterface) SetOversample(oversample float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversample)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnSetOversample), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MobileVRInterface) GetOversample() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnGetOversample), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MobileVRInterface) SetK1(k float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&k)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnSetK1), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MobileVRInterface) GetK1() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnGetK1), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *MobileVRInterface) SetK2(k float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&k)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnSetK2), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *MobileVRInterface) GetK2() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForMobileVRInterface.fnGetK2), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
