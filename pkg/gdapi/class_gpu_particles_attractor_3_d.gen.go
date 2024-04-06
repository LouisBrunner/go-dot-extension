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

type ptrsForGPUParticlesAttractor3DList struct {
	fnSetCullMask       gdc.MethodBindPtr
	fnGetCullMask       gdc.MethodBindPtr
	fnSetStrength       gdc.MethodBindPtr
	fnGetStrength       gdc.MethodBindPtr
	fnSetAttenuation    gdc.MethodBindPtr
	fnGetAttenuation    gdc.MethodBindPtr
	fnSetDirectionality gdc.MethodBindPtr
	fnGetDirectionality gdc.MethodBindPtr
}

var ptrsForGPUParticlesAttractor3D ptrsForGPUParticlesAttractor3DList

func initGPUParticlesAttractor3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GPUParticlesAttractor3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_cull_mask")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractor3D.fnSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_cull_mask")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractor3D.fnGetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_strength")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractor3D.fnSetStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_strength")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractor3D.fnGetStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_attenuation")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractor3D.fnSetAttenuation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_attenuation")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractor3D.fnGetAttenuation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_directionality")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractor3D.fnSetDirectionality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_directionality")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractor3D.fnGetDirectionality = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type GPUParticlesAttractor3D struct {
	VisualInstance3D
}

func (me *GPUParticlesAttractor3D) BaseClass() string {
	return "GPUParticlesAttractor3D"
}

func NewGPUParticlesAttractor3D() *GPUParticlesAttractor3D {
	str := StringNameFromStr("GPUParticlesAttractor3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GPUParticlesAttractor3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GPUParticlesAttractor3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GPUParticlesAttractor3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesAttractor3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GPUParticlesAttractor3D) SetCullMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractor3D.fnSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesAttractor3D) GetCullMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractor3D.fnGetCullMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticlesAttractor3D) SetStrength(strength float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractor3D.fnSetStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesAttractor3D) GetStrength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractor3D.fnGetStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticlesAttractor3D) SetAttenuation(attenuation float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&attenuation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractor3D.fnSetAttenuation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesAttractor3D) GetAttenuation() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractor3D.fnGetAttenuation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticlesAttractor3D) SetDirectionality(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractor3D.fnSetDirectionality), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesAttractor3D) GetDirectionality() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractor3D.fnGetDirectionality), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
