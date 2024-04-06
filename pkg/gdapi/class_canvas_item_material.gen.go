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

type ptrsForCanvasItemMaterialList struct {
	fnSetBlendMode            gdc.MethodBindPtr
	fnGetBlendMode            gdc.MethodBindPtr
	fnSetLightMode            gdc.MethodBindPtr
	fnGetLightMode            gdc.MethodBindPtr
	fnSetParticlesAnimation   gdc.MethodBindPtr
	fnGetParticlesAnimation   gdc.MethodBindPtr
	fnSetParticlesAnimHFrames gdc.MethodBindPtr
	fnGetParticlesAnimHFrames gdc.MethodBindPtr
	fnSetParticlesAnimVFrames gdc.MethodBindPtr
	fnGetParticlesAnimVFrames gdc.MethodBindPtr
	fnSetParticlesAnimLoop    gdc.MethodBindPtr
	fnGetParticlesAnimLoop    gdc.MethodBindPtr
}

var ptrsForCanvasItemMaterial ptrsForCanvasItemMaterialList

func initCanvasItemMaterialPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CanvasItemMaterial")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_blend_mode")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnSetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1786054936))
	}
	{
		methodName := StringNameFromStr("get_blend_mode")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnGetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3318684035))
	}
	{
		methodName := StringNameFromStr("set_light_mode")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnSetLightMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 628074070))
	}
	{
		methodName := StringNameFromStr("get_light_mode")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnGetLightMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3863292382))
	}
	{
		methodName := StringNameFromStr("set_particles_animation")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnSetParticlesAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_particles_animation")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnGetParticlesAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_particles_anim_h_frames")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnSetParticlesAnimHFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_particles_anim_h_frames")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnGetParticlesAnimHFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_particles_anim_v_frames")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnSetParticlesAnimVFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_particles_anim_v_frames")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnGetParticlesAnimVFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_particles_anim_loop")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnSetParticlesAnimLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_particles_anim_loop")
		defer methodName.Destroy()
		ptrsForCanvasItemMaterial.fnGetParticlesAnimLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type CanvasItemMaterial struct {
	Material
}

func (me *CanvasItemMaterial) BaseClass() string {
	return "CanvasItemMaterial"
}

func NewCanvasItemMaterial() *CanvasItemMaterial {
	str := StringNameFromStr("CanvasItemMaterial") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &CanvasItemMaterial{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type CanvasItemMaterialBlendMode int

const (
	CanvasItemMaterialBlendModeBlendModeMix          CanvasItemMaterialBlendMode = 0
	CanvasItemMaterialBlendModeBlendModeAdd          CanvasItemMaterialBlendMode = 1
	CanvasItemMaterialBlendModeBlendModeSub          CanvasItemMaterialBlendMode = 2
	CanvasItemMaterialBlendModeBlendModeMul          CanvasItemMaterialBlendMode = 3
	CanvasItemMaterialBlendModeBlendModePremultAlpha CanvasItemMaterialBlendMode = 4
)

type CanvasItemMaterialLightMode int

const (
	CanvasItemMaterialLightModeLightModeNormal    CanvasItemMaterialLightMode = 0
	CanvasItemMaterialLightModeLightModeUnshaded  CanvasItemMaterialLightMode = 1
	CanvasItemMaterialLightModeLightModeLightOnly CanvasItemMaterialLightMode = 2
)

func (me *CanvasItemMaterial) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *CanvasItemMaterial) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *CanvasItemMaterial) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *CanvasItemMaterial) SetBlendMode(blend_mode CanvasItemMaterialBlendMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnSetBlendMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItemMaterial) GetBlendMode() CanvasItemMaterialBlendMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CanvasItemMaterialBlendMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnGetBlendMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CanvasItemMaterial) SetLightMode(light_mode CanvasItemMaterialLightMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnSetLightMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItemMaterial) GetLightMode() CanvasItemMaterialLightMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CanvasItemMaterialLightMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnGetLightMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CanvasItemMaterial) SetParticlesAnimation(particles_anim bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&particles_anim)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnSetParticlesAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItemMaterial) GetParticlesAnimation() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnGetParticlesAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItemMaterial) SetParticlesAnimHFrames(frames int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnSetParticlesAnimHFrames), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItemMaterial) GetParticlesAnimHFrames() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnGetParticlesAnimHFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItemMaterial) SetParticlesAnimVFrames(frames int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnSetParticlesAnimVFrames), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItemMaterial) GetParticlesAnimVFrames() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnGetParticlesAnimVFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItemMaterial) SetParticlesAnimLoop(loop bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnSetParticlesAnimLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItemMaterial) GetParticlesAnimLoop() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItemMaterial.fnGetParticlesAnimLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
