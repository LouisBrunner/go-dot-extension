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

type ptrsForGPUParticlesCollisionSDF3DList struct {
	fnSetSize          gdc.MethodBindPtr
	fnGetSize          gdc.MethodBindPtr
	fnSetResolution    gdc.MethodBindPtr
	fnGetResolution    gdc.MethodBindPtr
	fnSetTexture       gdc.MethodBindPtr
	fnGetTexture       gdc.MethodBindPtr
	fnSetThickness     gdc.MethodBindPtr
	fnGetThickness     gdc.MethodBindPtr
	fnSetBakeMask      gdc.MethodBindPtr
	fnGetBakeMask      gdc.MethodBindPtr
	fnSetBakeMaskValue gdc.MethodBindPtr
	fnGetBakeMaskValue gdc.MethodBindPtr
}

var ptrsForGPUParticlesCollisionSDF3D ptrsForGPUParticlesCollisionSDF3DList

func initGPUParticlesCollisionSDF3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GPUParticlesCollisionSDF3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_resolution")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnSetResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1155629297))
	}
	{
		methodName := StringNameFromStr("get_resolution")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnGetResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2919555867))
	}
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1188404210))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373985333))
	}
	{
		methodName := StringNameFromStr("set_thickness")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnSetThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_thickness")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnGetThickness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_bake_mask")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnSetBakeMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_bake_mask")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnGetBakeMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_bake_mask_value")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnSetBakeMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_bake_mask_value")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionSDF3D.fnGetBakeMaskValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
}

type GPUParticlesCollisionSDF3D struct {
	GPUParticlesCollision3D
}

func (me *GPUParticlesCollisionSDF3D) BaseClass() string {
	return "GPUParticlesCollisionSDF3D"
}

func NewGPUParticlesCollisionSDF3D() *GPUParticlesCollisionSDF3D {
	str := StringNameFromStr("GPUParticlesCollisionSDF3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GPUParticlesCollisionSDF3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type GPUParticlesCollisionSDF3DResolution int

const (
	GPUParticlesCollisionSDF3DResolutionResolution16  GPUParticlesCollisionSDF3DResolution = 0
	GPUParticlesCollisionSDF3DResolutionResolution32  GPUParticlesCollisionSDF3DResolution = 1
	GPUParticlesCollisionSDF3DResolutionResolution64  GPUParticlesCollisionSDF3DResolution = 2
	GPUParticlesCollisionSDF3DResolutionResolution128 GPUParticlesCollisionSDF3DResolution = 3
	GPUParticlesCollisionSDF3DResolutionResolution256 GPUParticlesCollisionSDF3DResolution = 4
	GPUParticlesCollisionSDF3DResolutionResolution512 GPUParticlesCollisionSDF3DResolution = 5
	GPUParticlesCollisionSDF3DResolutionResolutionMax GPUParticlesCollisionSDF3DResolution = 6
)

func (me *GPUParticlesCollisionSDF3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GPUParticlesCollisionSDF3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollisionSDF3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GPUParticlesCollisionSDF3D) SetSize(size Vector3) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionSDF3D) GetSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GPUParticlesCollisionSDF3D) SetResolution(resolution GPUParticlesCollisionSDF3DResolution) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resolution)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnSetResolution), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionSDF3D) GetResolution() GPUParticlesCollisionSDF3DResolution {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GPUParticlesCollisionSDF3DResolution

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnGetResolution), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GPUParticlesCollisionSDF3D) SetTexture(texture Texture3D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionSDF3D) GetTexture() Texture3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GPUParticlesCollisionSDF3D) SetThickness(thickness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&thickness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnSetThickness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionSDF3D) GetThickness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnGetThickness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticlesCollisionSDF3D) SetBakeMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnSetBakeMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionSDF3D) GetBakeMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnGetBakeMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticlesCollisionSDF3D) SetBakeMaskValue(layer_number int64, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnSetBakeMaskValue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionSDF3D) GetBakeMaskValue(layer_number int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer_number)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer_number)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionSDF3D.fnGetBakeMaskValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
