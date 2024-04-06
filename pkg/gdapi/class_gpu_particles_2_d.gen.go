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

type ptrsForGPUParticles2DList struct {
	fnSetEmitting                 gdc.MethodBindPtr
	fnSetAmount                   gdc.MethodBindPtr
	fnSetLifetime                 gdc.MethodBindPtr
	fnSetOneShot                  gdc.MethodBindPtr
	fnSetPreProcessTime           gdc.MethodBindPtr
	fnSetExplosivenessRatio       gdc.MethodBindPtr
	fnSetRandomnessRatio          gdc.MethodBindPtr
	fnSetVisibilityRect           gdc.MethodBindPtr
	fnSetUseLocalCoordinates      gdc.MethodBindPtr
	fnSetFixedFps                 gdc.MethodBindPtr
	fnSetFractionalDelta          gdc.MethodBindPtr
	fnSetInterpolate              gdc.MethodBindPtr
	fnSetProcessMaterial          gdc.MethodBindPtr
	fnSetSpeedScale               gdc.MethodBindPtr
	fnSetCollisionBaseSize        gdc.MethodBindPtr
	fnSetInterpToEnd              gdc.MethodBindPtr
	fnIsEmitting                  gdc.MethodBindPtr
	fnGetAmount                   gdc.MethodBindPtr
	fnGetLifetime                 gdc.MethodBindPtr
	fnGetOneShot                  gdc.MethodBindPtr
	fnGetPreProcessTime           gdc.MethodBindPtr
	fnGetExplosivenessRatio       gdc.MethodBindPtr
	fnGetRandomnessRatio          gdc.MethodBindPtr
	fnGetVisibilityRect           gdc.MethodBindPtr
	fnGetUseLocalCoordinates      gdc.MethodBindPtr
	fnGetFixedFps                 gdc.MethodBindPtr
	fnGetFractionalDelta          gdc.MethodBindPtr
	fnGetInterpolate              gdc.MethodBindPtr
	fnGetProcessMaterial          gdc.MethodBindPtr
	fnGetSpeedScale               gdc.MethodBindPtr
	fnGetCollisionBaseSize        gdc.MethodBindPtr
	fnGetInterpToEnd              gdc.MethodBindPtr
	fnSetDrawOrder                gdc.MethodBindPtr
	fnGetDrawOrder                gdc.MethodBindPtr
	fnSetTexture                  gdc.MethodBindPtr
	fnGetTexture                  gdc.MethodBindPtr
	fnCaptureRect                 gdc.MethodBindPtr
	fnRestart                     gdc.MethodBindPtr
	fnSetSubEmitter               gdc.MethodBindPtr
	fnGetSubEmitter               gdc.MethodBindPtr
	fnEmitParticle                gdc.MethodBindPtr
	fnSetTrailEnabled             gdc.MethodBindPtr
	fnSetTrailLifetime            gdc.MethodBindPtr
	fnIsTrailEnabled              gdc.MethodBindPtr
	fnGetTrailLifetime            gdc.MethodBindPtr
	fnSetTrailSections            gdc.MethodBindPtr
	fnGetTrailSections            gdc.MethodBindPtr
	fnSetTrailSectionSubdivisions gdc.MethodBindPtr
	fnGetTrailSectionSubdivisions gdc.MethodBindPtr
	fnConvertFromParticles        gdc.MethodBindPtr
	fnSetAmountRatio              gdc.MethodBindPtr
	fnGetAmountRatio              gdc.MethodBindPtr
}

var ptrsForGPUParticles2D ptrsForGPUParticles2DList

func initGPUParticles2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("GPUParticles2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_emitting")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_amount")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_lifetime")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_one_shot")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_pre_process_time")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetPreProcessTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_explosiveness_ratio")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetExplosivenessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_randomness_ratio")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetRandomnessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_visibility_rect")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetVisibilityRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
	}
	{
		methodName := StringNameFromStr("set_use_local_coordinates")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetUseLocalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_fixed_fps")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetFixedFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_fractional_delta")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetFractionalDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_interpolate")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_process_material")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetProcessMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("set_speed_scale")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_collision_base_size")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetCollisionBaseSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_interp_to_end")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetInterpToEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("is_emitting")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnIsEmitting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_amount")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetAmount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_lifetime")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_one_shot")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_pre_process_time")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetPreProcessTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_explosiveness_ratio")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetExplosivenessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_randomness_ratio")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetRandomnessRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_visibility_rect")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetVisibilityRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("get_use_local_coordinates")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetUseLocalCoordinates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_fixed_fps")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetFixedFps = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_fractional_delta")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetFractionalDelta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_interpolate")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_process_material")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetProcessMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
	{
		methodName := StringNameFromStr("get_speed_scale")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_collision_base_size")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetCollisionBaseSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_interp_to_end")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetInterpToEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_draw_order")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetDrawOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1939677959))
	}
	{
		methodName := StringNameFromStr("get_draw_order")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetDrawOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 941479095))
	}
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("capture_rect")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnCaptureRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("restart")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnRestart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_sub_emitter")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetSubEmitter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_sub_emitter")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetSubEmitter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("emit_particle")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnEmitParticle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2179202058))
	}
	{
		methodName := StringNameFromStr("set_trail_enabled")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetTrailEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_trail_lifetime")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetTrailLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("is_trail_enabled")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnIsTrailEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_trail_lifetime")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetTrailLifetime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_trail_sections")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetTrailSections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_trail_sections")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetTrailSections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_trail_section_subdivisions")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetTrailSectionSubdivisions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_trail_section_subdivisions")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetTrailSectionSubdivisions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("convert_from_particles")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnConvertFromParticles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("set_amount_ratio")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnSetAmountRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_amount_ratio")
		defer methodName.Destroy()
		ptrsForGPUParticles2D.fnGetAmountRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type GPUParticles2D struct {
	Node2D
}

func (me *GPUParticles2D) BaseClass() string {
	return "GPUParticles2D"
}

func NewGPUParticles2D() *GPUParticles2D {
	str := StringNameFromStr("GPUParticles2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &GPUParticles2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type GPUParticles2DDrawOrder int

const (
	GPUParticles2DDrawOrderDrawOrderIndex           GPUParticles2DDrawOrder = 0
	GPUParticles2DDrawOrderDrawOrderLifetime        GPUParticles2DDrawOrder = 1
	GPUParticles2DDrawOrderDrawOrderReverseLifetime GPUParticles2DDrawOrder = 2
)

type GPUParticles2DEmitFlags int

const (
	GPUParticles2DEmitFlagsEmitFlagPosition      GPUParticles2DEmitFlags = 1
	GPUParticles2DEmitFlagsEmitFlagRotationScale GPUParticles2DEmitFlags = 2
	GPUParticles2DEmitFlagsEmitFlagVelocity      GPUParticles2DEmitFlags = 4
	GPUParticles2DEmitFlagsEmitFlagColor         GPUParticles2DEmitFlags = 8
	GPUParticles2DEmitFlagsEmitFlagCustom        GPUParticles2DEmitFlags = 16
)

func (me *GPUParticles2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *GPUParticles2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *GPUParticles2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *GPUParticles2D) SetEmitting(emitting bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&emitting)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetEmitting), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetAmount(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetAmount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetLifetime(secs float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetLifetime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetOneShot(secs bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetOneShot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetPreProcessTime(secs float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetPreProcessTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetExplosivenessRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetExplosivenessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetRandomnessRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetRandomnessRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetVisibilityRect(visibility_rect Rect2) {
	cargs := []gdc.ConstTypePtr{visibility_rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetVisibilityRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetUseLocalCoordinates(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetUseLocalCoordinates), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetFixedFps(fps int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetFixedFps), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetFractionalDelta(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetFractionalDelta), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetInterpolate(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetInterpolate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetProcessMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetProcessMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetSpeedScale(scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetCollisionBaseSize(size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetCollisionBaseSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetInterpToEnd(interp float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interp)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetInterpToEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) IsEmitting() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnIsEmitting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetAmount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetAmount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetLifetime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetLifetime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetOneShot() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetOneShot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetPreProcessTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetPreProcessTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetExplosivenessRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetExplosivenessRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetRandomnessRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetRandomnessRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetVisibilityRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetVisibilityRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GPUParticles2D) GetUseLocalCoordinates() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetUseLocalCoordinates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetFixedFps() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetFixedFps), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetFractionalDelta() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetFractionalDelta), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetInterpolate() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetInterpolate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetProcessMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetProcessMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GPUParticles2D) GetSpeedScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetCollisionBaseSize() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetCollisionBaseSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetInterpToEnd() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetInterpToEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) SetDrawOrder(order GPUParticles2DDrawOrder) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetDrawOrder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) GetDrawOrder() GPUParticles2DDrawOrder {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret GPUParticles2DDrawOrder

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetDrawOrder), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *GPUParticles2D) SetTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) GetTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GPUParticles2D) CaptureRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnCaptureRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GPUParticles2D) Restart() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnRestart), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetSubEmitter(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetSubEmitter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) GetSubEmitter() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetSubEmitter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *GPUParticles2D) EmitParticle(xform Transform2D, velocity Vector2, color Color, custom Color, flags int64) {
	cargs := []gdc.ConstTypePtr{xform.AsCTypePtr(), velocity.AsCTypePtr(), color.AsCTypePtr(), custom.AsCTypePtr(), gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnEmitParticle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetTrailEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetTrailEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetTrailLifetime(secs float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetTrailLifetime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) IsTrailEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnIsTrailEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) GetTrailLifetime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetTrailLifetime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) SetTrailSections(sections int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sections)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetTrailSections), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) GetTrailSections() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetTrailSections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) SetTrailSectionSubdivisions(subdivisions int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdivisions)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetTrailSectionSubdivisions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) GetTrailSectionSubdivisions() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetTrailSectionSubdivisions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *GPUParticles2D) ConvertFromParticles(particles Node) {
	cargs := []gdc.ConstTypePtr{particles.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnConvertFromParticles), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) SetAmountRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnSetAmountRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *GPUParticles2D) GetAmountRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGPUParticles2D.fnGetAmountRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GPUParticles2DFinishedSignalFn func()

func (me *GPUParticles2D) ConnectFinished(subs SignalSubscribers, fn GPUParticles2DFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *GPUParticles2D) DisconnectFinished(subs SignalSubscribers, fn GPUParticles2DFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
