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

type ptrsForDirectionalLight3DList struct {
	fnSetShadowMode        gdc.MethodBindPtr
	fnGetShadowMode        gdc.MethodBindPtr
	fnSetBlendSplits       gdc.MethodBindPtr
	fnIsBlendSplitsEnabled gdc.MethodBindPtr
	fnSetSkyMode           gdc.MethodBindPtr
	fnGetSkyMode           gdc.MethodBindPtr
}

var ptrsForDirectionalLight3D ptrsForDirectionalLight3DList

func initDirectionalLight3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("DirectionalLight3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_shadow_mode")
		defer methodName.Destroy()
		ptrsForDirectionalLight3D.fnSetShadowMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1261211726))
	}
	{
		methodName := StringNameFromStr("get_shadow_mode")
		defer methodName.Destroy()
		ptrsForDirectionalLight3D.fnGetShadowMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2765228544))
	}
	{
		methodName := StringNameFromStr("set_blend_splits")
		defer methodName.Destroy()
		ptrsForDirectionalLight3D.fnSetBlendSplits = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_blend_splits_enabled")
		defer methodName.Destroy()
		ptrsForDirectionalLight3D.fnIsBlendSplitsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_sky_mode")
		defer methodName.Destroy()
		ptrsForDirectionalLight3D.fnSetSkyMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2691194817))
	}
	{
		methodName := StringNameFromStr("get_sky_mode")
		defer methodName.Destroy()
		ptrsForDirectionalLight3D.fnGetSkyMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3819982774))
	}
}

type DirectionalLight3D struct {
	Light3D
}

func (me *DirectionalLight3D) BaseClass() string {
	return "DirectionalLight3D"
}

func NewDirectionalLight3D() *DirectionalLight3D {
	str := StringNameFromStr("DirectionalLight3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &DirectionalLight3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type DirectionalLight3DShadowMode int

const (
	DirectionalLight3DShadowModeShadowOrthogonal      DirectionalLight3DShadowMode = 0
	DirectionalLight3DShadowModeShadowParallel2Splits DirectionalLight3DShadowMode = 1
	DirectionalLight3DShadowModeShadowParallel4Splits DirectionalLight3DShadowMode = 2
)

type DirectionalLight3DSkyMode int

const (
	DirectionalLight3DSkyModeSkyModeLightAndSky DirectionalLight3DSkyMode = 0
	DirectionalLight3DSkyModeSkyModeLightOnly   DirectionalLight3DSkyMode = 1
	DirectionalLight3DSkyModeSkyModeSkyOnly     DirectionalLight3DSkyMode = 2
)

func (me *DirectionalLight3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *DirectionalLight3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *DirectionalLight3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *DirectionalLight3D) SetShadowMode(mode DirectionalLight3DShadowMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirectionalLight3D.fnSetShadowMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DirectionalLight3D) GetShadowMode() DirectionalLight3DShadowMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret DirectionalLight3DShadowMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirectionalLight3D.fnGetShadowMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DirectionalLight3D) SetBlendSplits(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirectionalLight3D.fnSetBlendSplits), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DirectionalLight3D) IsBlendSplitsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirectionalLight3D.fnIsBlendSplitsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DirectionalLight3D) SetSkyMode(mode DirectionalLight3DSkyMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirectionalLight3D.fnSetSkyMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DirectionalLight3D) GetSkyMode() DirectionalLight3DSkyMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret DirectionalLight3DSkyMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirectionalLight3D.fnGetSkyMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
