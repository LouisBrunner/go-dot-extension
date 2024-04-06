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

type ptrsForGPUParticlesCollisionHeightField3DList struct {
	fnSetSize                gdc.MethodBindPtr
	fnGetSize                gdc.MethodBindPtr
	fnSetResolution          gdc.MethodBindPtr
	fnGetResolution          gdc.MethodBindPtr
	fnSetUpdateMode          gdc.MethodBindPtr
	fnGetUpdateMode          gdc.MethodBindPtr
	fnSetFollowCameraEnabled gdc.MethodBindPtr
	fnIsFollowCameraEnabled  gdc.MethodBindPtr
}

var ptrsForGPUParticlesCollisionHeightField3D ptrsForGPUParticlesCollisionHeightField3DList

func initGPUParticlesCollisionHeightField3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GPUParticlesCollisionHeightField3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionHeightField3D.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionHeightField3D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_resolution")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionHeightField3D.fnSetResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1009996517))
	}
	{
		methodName := StringNameFromStr("get_resolution")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionHeightField3D.fnGetResolution = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1156065644))
	}
	{
		methodName := StringNameFromStr("set_update_mode")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionHeightField3D.fnSetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 673680859))
	}
	{
		methodName := StringNameFromStr("get_update_mode")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionHeightField3D.fnGetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1998141380))
	}
	{
		methodName := StringNameFromStr("set_follow_camera_enabled")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionHeightField3D.fnSetFollowCameraEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_follow_camera_enabled")
		defer methodName.Destroy()
		ptrsForGPUParticlesCollisionHeightField3D.fnIsFollowCameraEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type GPUParticlesCollisionHeightField3D struct {
	GPUParticlesCollision3D
}

func (me *GPUParticlesCollisionHeightField3D) BaseClass() string {
	return "GPUParticlesCollisionHeightField3D"
}

func NewGPUParticlesCollisionHeightField3D() *GPUParticlesCollisionHeightField3D {
	str := StringNameFromStr("GPUParticlesCollisionHeightField3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GPUParticlesCollisionHeightField3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type GPUParticlesCollisionHeightField3DResolution int

const (
	GPUParticlesCollisionHeightField3DResolutionResolution256  GPUParticlesCollisionHeightField3DResolution = 0
	GPUParticlesCollisionHeightField3DResolutionResolution512  GPUParticlesCollisionHeightField3DResolution = 1
	GPUParticlesCollisionHeightField3DResolutionResolution1024 GPUParticlesCollisionHeightField3DResolution = 2
	GPUParticlesCollisionHeightField3DResolutionResolution2048 GPUParticlesCollisionHeightField3DResolution = 3
	GPUParticlesCollisionHeightField3DResolutionResolution4096 GPUParticlesCollisionHeightField3DResolution = 4
	GPUParticlesCollisionHeightField3DResolutionResolution8192 GPUParticlesCollisionHeightField3DResolution = 5
	GPUParticlesCollisionHeightField3DResolutionResolutionMax  GPUParticlesCollisionHeightField3DResolution = 6
)

type GPUParticlesCollisionHeightField3DUpdateMode int

const (
	GPUParticlesCollisionHeightField3DUpdateModeUpdateModeWhenMoved GPUParticlesCollisionHeightField3DUpdateMode = 0
	GPUParticlesCollisionHeightField3DUpdateModeUpdateModeAlways    GPUParticlesCollisionHeightField3DUpdateMode = 1
)

func (me *GPUParticlesCollisionHeightField3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GPUParticlesCollisionHeightField3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GPUParticlesCollisionHeightField3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GPUParticlesCollisionHeightField3D) SetSize(size Vector3) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionHeightField3D.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionHeightField3D) GetSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionHeightField3D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GPUParticlesCollisionHeightField3D) SetResolution(resolution GPUParticlesCollisionHeightField3DResolution) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resolution)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionHeightField3D.fnSetResolution), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionHeightField3D) GetResolution() GPUParticlesCollisionHeightField3DResolution {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GPUParticlesCollisionHeightField3DResolution

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionHeightField3D.fnGetResolution), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GPUParticlesCollisionHeightField3D) SetUpdateMode(update_mode GPUParticlesCollisionHeightField3DUpdateMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&update_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionHeightField3D.fnSetUpdateMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionHeightField3D) GetUpdateMode() GPUParticlesCollisionHeightField3DUpdateMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GPUParticlesCollisionHeightField3DUpdateMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionHeightField3D.fnGetUpdateMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GPUParticlesCollisionHeightField3D) SetFollowCameraEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionHeightField3D.fnSetFollowCameraEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticlesCollisionHeightField3D) IsFollowCameraEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticlesCollisionHeightField3D.fnIsFollowCameraEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
