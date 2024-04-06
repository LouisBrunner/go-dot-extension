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

type ptrsForCameraAttributesList struct {
	fnSetExposureMultiplier  gdc.MethodBindPtr
	fnGetExposureMultiplier  gdc.MethodBindPtr
	fnSetExposureSensitivity gdc.MethodBindPtr
	fnGetExposureSensitivity gdc.MethodBindPtr
	fnSetAutoExposureEnabled gdc.MethodBindPtr
	fnIsAutoExposureEnabled  gdc.MethodBindPtr
	fnSetAutoExposureSpeed   gdc.MethodBindPtr
	fnGetAutoExposureSpeed   gdc.MethodBindPtr
	fnSetAutoExposureScale   gdc.MethodBindPtr
	fnGetAutoExposureScale   gdc.MethodBindPtr
}

var ptrsForCameraAttributes ptrsForCameraAttributesList

func initCameraAttributesPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CameraAttributes")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_exposure_multiplier")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnSetExposureMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_exposure_multiplier")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnGetExposureMultiplier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_exposure_sensitivity")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnSetExposureSensitivity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_exposure_sensitivity")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnGetExposureSensitivity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_auto_exposure_enabled")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnSetAutoExposureEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_auto_exposure_enabled")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnIsAutoExposureEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_auto_exposure_speed")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnSetAutoExposureSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_auto_exposure_speed")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnGetAutoExposureSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_auto_exposure_scale")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnSetAutoExposureScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_auto_exposure_scale")
		defer methodName.Destroy()
		ptrsForCameraAttributes.fnGetAutoExposureScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type CameraAttributes struct {
	Resource
}

func (me *CameraAttributes) BaseClass() string {
	return "CameraAttributes"
}

func NewCameraAttributes() *CameraAttributes {
	str := StringNameFromStr("CameraAttributes") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CameraAttributes{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *CameraAttributes) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CameraAttributes) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CameraAttributes) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CameraAttributes) SetExposureMultiplier(multiplier float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&multiplier)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnSetExposureMultiplier), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CameraAttributes) GetExposureMultiplier() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnGetExposureMultiplier), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CameraAttributes) SetExposureSensitivity(sensitivity float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sensitivity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnSetExposureSensitivity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CameraAttributes) GetExposureSensitivity() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnGetExposureSensitivity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CameraAttributes) SetAutoExposureEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnSetAutoExposureEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CameraAttributes) IsAutoExposureEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnIsAutoExposureEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CameraAttributes) SetAutoExposureSpeed(exposure_speed float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnSetAutoExposureSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CameraAttributes) GetAutoExposureSpeed() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnGetAutoExposureSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CameraAttributes) SetAutoExposureScale(exposure_grey float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exposure_grey)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnSetAutoExposureScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CameraAttributes) GetAutoExposureScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCameraAttributes.fnGetAutoExposureScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
