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

type ptrsForSpriteFramesList struct {
	fnAddAnimation      gdc.MethodBindPtr
	fnHasAnimation      gdc.MethodBindPtr
	fnRemoveAnimation   gdc.MethodBindPtr
	fnRenameAnimation   gdc.MethodBindPtr
	fnGetAnimationNames gdc.MethodBindPtr
	fnSetAnimationSpeed gdc.MethodBindPtr
	fnGetAnimationSpeed gdc.MethodBindPtr
	fnSetAnimationLoop  gdc.MethodBindPtr
	fnGetAnimationLoop  gdc.MethodBindPtr
	fnAddFrame          gdc.MethodBindPtr
	fnSetFrame          gdc.MethodBindPtr
	fnRemoveFrame       gdc.MethodBindPtr
	fnGetFrameCount     gdc.MethodBindPtr
	fnGetFrameTexture   gdc.MethodBindPtr
	fnGetFrameDuration  gdc.MethodBindPtr
	fnClear             gdc.MethodBindPtr
	fnClearAll          gdc.MethodBindPtr
}

var ptrsForSpriteFrames ptrsForSpriteFramesList

func initSpriteFramesPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SpriteFrames")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_animation")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnAddAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("has_animation")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnHasAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("remove_animation")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnRemoveAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("rename_animation")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnRenameAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("get_animation_names")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnGetAnimationNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_animation_speed")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnSetAnimationSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4135858297))
	}
	{
		methodName := StringNameFromStr("get_animation_speed")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnGetAnimationSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2349060816))
	}
	{
		methodName := StringNameFromStr("set_animation_loop")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnSetAnimationLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2524380260))
	}
	{
		methodName := StringNameFromStr("get_animation_loop")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnGetAnimationLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("add_frame")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnAddFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1351332740))
	}
	{
		methodName := StringNameFromStr("set_frame")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnSetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 56804795))
	}
	{
		methodName := StringNameFromStr("remove_frame")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnRemoveFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2415702435))
	}
	{
		methodName := StringNameFromStr("get_frame_count")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnGetFrameCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2458036349))
	}
	{
		methodName := StringNameFromStr("get_frame_texture")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnGetFrameTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2900517879))
	}
	{
		methodName := StringNameFromStr("get_frame_duration")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnGetFrameDuration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1129309260))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("clear_all")
		defer methodName.Destroy()
		ptrsForSpriteFrames.fnClearAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type SpriteFrames struct {
	Resource
}

func (me *SpriteFrames) BaseClass() string {
	return "SpriteFrames"
}

func NewSpriteFrames() *SpriteFrames {
	str := StringNameFromStr("SpriteFrames") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SpriteFrames{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SpriteFrames) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SpriteFrames) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SpriteFrames) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SpriteFrames) AddAnimation(anim StringName) {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnAddAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpriteFrames) HasAnimation(anim StringName) bool {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnHasAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SpriteFrames) RemoveAnimation(anim StringName) {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnRemoveAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpriteFrames) RenameAnimation(anim StringName, newname StringName) {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), newname.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnRenameAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpriteFrames) GetAnimationNames() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnGetAnimationNames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SpriteFrames) SetAnimationSpeed(anim StringName, fps float64) {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&fps)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnSetAnimationSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpriteFrames) GetAnimationSpeed(anim StringName) float64 {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnGetAnimationSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SpriteFrames) SetAnimationLoop(anim StringName, loop bool) {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&loop)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnSetAnimationLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpriteFrames) GetAnimationLoop(anim StringName) bool {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnGetAnimationLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SpriteFrames) AddFrame(anim StringName, texture Texture2D, duration float64, at_position int64) {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), texture.AsCTypePtr(), gdc.ConstTypePtr(&duration), gdc.ConstTypePtr(&at_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnAddFrame), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpriteFrames) SetFrame(anim StringName, idx int64, texture Texture2D, duration float64) {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&idx), texture.AsCTypePtr(), gdc.ConstTypePtr(&duration)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnSetFrame), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpriteFrames) RemoveFrame(anim StringName, idx int64) {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnRemoveFrame), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpriteFrames) GetFrameCount(anim StringName) int64 {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnGetFrameCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SpriteFrames) GetFrameTexture(anim StringName, idx int64) Texture2D {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnGetFrameTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SpriteFrames) GetFrameDuration(anim StringName, idx int64) float64 {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnGetFrameDuration), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SpriteFrames) Clear(anim StringName) {
	cargs := []gdc.ConstTypePtr{anim.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SpriteFrames) ClearAll() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpriteFrames.fnClearAll), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
