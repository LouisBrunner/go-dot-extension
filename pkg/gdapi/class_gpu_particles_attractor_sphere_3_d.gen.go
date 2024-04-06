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

type ptrsForGPUParticlesAttractorSphere3DList struct {
	fnSetRadius gdc.MethodBindPtr
	fnGetRadius gdc.MethodBindPtr
}

var ptrsForGPUParticlesAttractorSphere3D ptrsForGPUParticlesAttractorSphere3DList

func initGPUParticlesAttractorSphere3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GPUParticlesAttractorSphere3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractorSphere3D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractorSphere3D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type GPUParticlesAttractorSphere3D struct {
	GPUParticlesAttractor3D
}

func (me *GPUParticlesAttractorSphere3D) BaseClass() string {
	return "GPUParticlesAttractorSphere3D"
}

func NewGPUParticlesAttractorSphere3D() *GPUParticlesAttractorSphere3D {
	str := StringNameFromStr("GPUParticlesAttractorSphere3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GPUParticlesAttractorSphere3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GPUParticlesAttractorSphere3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GPUParticlesAttractorSphere3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesAttractorSphere3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GPUParticlesAttractorSphere3D) SetRadius(radius float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractorSphere3D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesAttractorSphere3D) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractorSphere3D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
