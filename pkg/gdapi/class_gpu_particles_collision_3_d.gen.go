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

type ptrsForGPUParticlesCollision3DList struct {
	fnSetCullMask gdc.MethodBindPtr
	fnGetCullMask gdc.MethodBindPtr
}

var ptrsForGPUParticlesCollision3D ptrsForGPUParticlesCollision3DList

func initGPUParticlesCollision3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GPUParticlesCollision3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_cull_mask")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollision3D.fnSetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_cull_mask")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollision3D.fnGetCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type GPUParticlesCollision3D struct {
	VisualInstance3D
}

func (me *GPUParticlesCollision3D) BaseClass() string {
	return "GPUParticlesCollision3D"
}

func NewGPUParticlesCollision3D() *GPUParticlesCollision3D {
	str := StringNameFromStr("GPUParticlesCollision3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GPUParticlesCollision3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *GPUParticlesCollision3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GPUParticlesCollision3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollision3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GPUParticlesCollision3D) SetCullMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollision3D.fnSetCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollision3D) GetCullMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollision3D.fnGetCullMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
