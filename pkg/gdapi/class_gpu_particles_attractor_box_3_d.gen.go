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

type ptrsForGPUParticlesAttractorBox3DList struct {
	fnSetSize gdc.MethodBindPtr
	fnGetSize gdc.MethodBindPtr
}

var ptrsForGPUParticlesAttractorBox3D ptrsForGPUParticlesAttractorBox3DList

func initGPUParticlesAttractorBox3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GPUParticlesAttractorBox3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractorBox3D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForGPUParticlesAttractorBox3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}

}

type GPUParticlesAttractorBox3D struct {
	GPUParticlesAttractor3D
}

func (me *GPUParticlesAttractorBox3D) BaseClass() string {
	return "GPUParticlesAttractorBox3D"
}

func NewGPUParticlesAttractorBox3D() *GPUParticlesAttractorBox3D {
	str := StringNameFromStr("GPUParticlesAttractorBox3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GPUParticlesAttractorBox3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GPUParticlesAttractorBox3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GPUParticlesAttractorBox3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesAttractorBox3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GPUParticlesAttractorBox3D) SetSize(size Vector3) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractorBox3D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesAttractorBox3D) GetSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesAttractorBox3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
