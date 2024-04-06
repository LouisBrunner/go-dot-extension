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

type ptrsForAnimatedSprite3DList struct {
	fnSetSpriteFrames     gdc.MethodBindPtr
	fnGetSpriteFrames     gdc.MethodBindPtr
	fnSetAnimation        gdc.MethodBindPtr
	fnGetAnimation        gdc.MethodBindPtr
	fnSetAutoplay         gdc.MethodBindPtr
	fnGetAutoplay         gdc.MethodBindPtr
	fnIsPlaying           gdc.MethodBindPtr
	fnPlay                gdc.MethodBindPtr
	fnPlayBackwards       gdc.MethodBindPtr
	fnPause               gdc.MethodBindPtr
	fnStop                gdc.MethodBindPtr
	fnSetFrame            gdc.MethodBindPtr
	fnGetFrame            gdc.MethodBindPtr
	fnSetFrameProgress    gdc.MethodBindPtr
	fnGetFrameProgress    gdc.MethodBindPtr
	fnSetFrameAndProgress gdc.MethodBindPtr
	fnSetSpeedScale       gdc.MethodBindPtr
	fnGetSpeedScale       gdc.MethodBindPtr
	fnGetPlayingSpeed     gdc.MethodBindPtr
}

var ptrsForAnimatedSprite3D ptrsForAnimatedSprite3DList

func initAnimatedSprite3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimatedSprite3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_sprite_frames")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnSetSpriteFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 905781144))
	}
	{
		methodName := StringNameFromStr("get_sprite_frames")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnGetSpriteFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3804851214))
	}
	{
		methodName := StringNameFromStr("set_animation")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnSetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_animation")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnGetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_autoplay")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnSetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_autoplay")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnGetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_playing")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnIsPlaying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("play")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnPlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2372066587))
	}
	{
		methodName := StringNameFromStr("play_backwards")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnPlayBackwards = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1421762485))
	}
	{
		methodName := StringNameFromStr("pause")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnPause = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("stop")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_frame")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnSetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_frame")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnGetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_frame_progress")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnSetFrameProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_frame_progress")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnGetFrameProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_frame_and_progress")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnSetFrameAndProgress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("set_speed_scale")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_speed_scale")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnGetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_playing_speed")
		defer methodName.Destroy()
		ptrsForAnimatedSprite3D.fnGetPlayingSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type AnimatedSprite3D struct {
	SpriteBase3D
}

func (me *AnimatedSprite3D) BaseClass() string {
	return "AnimatedSprite3D"
}

func NewAnimatedSprite3D() *AnimatedSprite3D {
	str := StringNameFromStr("AnimatedSprite3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimatedSprite3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimatedSprite3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimatedSprite3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimatedSprite3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimatedSprite3D) SetSpriteFrames(sprite_frames SpriteFrames) {
	cargs := []gdc.ConstTypePtr{sprite_frames.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnSetSpriteFrames), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) GetSpriteFrames() SpriteFrames {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSpriteFrames()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnGetSpriteFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimatedSprite3D) SetAnimation(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnSetAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) GetAnimation() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnGetAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimatedSprite3D) SetAutoplay(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnSetAutoplay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) GetAutoplay() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnGetAutoplay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimatedSprite3D) IsPlaying() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnIsPlaying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimatedSprite3D) Play(name StringName, custom_speed float64, from_end bool) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&custom_speed), gdc.ConstTypePtr(&from_end)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnPlay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) PlayBackwards(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnPlayBackwards), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) Pause() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnPause), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) Stop() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) SetFrame(frame int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnSetFrame), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) GetFrame() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnGetFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimatedSprite3D) SetFrameProgress(progress float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&progress)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnSetFrameProgress), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) GetFrameProgress() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnGetFrameProgress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimatedSprite3D) SetFrameAndProgress(frame int64, progress float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame), gdc.ConstTypePtr(&progress)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnSetFrameAndProgress), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) SetSpeedScale(speed_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnSetSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimatedSprite3D) GetSpeedScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnGetSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimatedSprite3D) GetPlayingSpeed() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedSprite3D.fnGetPlayingSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimatedSprite3DSpriteFramesChangedSignalFn func()

func (me *AnimatedSprite3D) ConnectSpriteFramesChanged(subs SignalSubscribers, fn AnimatedSprite3DSpriteFramesChangedSignalFn) {
	sig := StringNameFromStr("sprite_frames_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite3D) DisconnectSpriteFramesChanged(subs SignalSubscribers, fn AnimatedSprite3DSpriteFramesChangedSignalFn) {
	sig := StringNameFromStr("sprite_frames_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimatedSprite3DAnimationChangedSignalFn func()

func (me *AnimatedSprite3D) ConnectAnimationChanged(subs SignalSubscribers, fn AnimatedSprite3DAnimationChangedSignalFn) {
	sig := StringNameFromStr("animation_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite3D) DisconnectAnimationChanged(subs SignalSubscribers, fn AnimatedSprite3DAnimationChangedSignalFn) {
	sig := StringNameFromStr("animation_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimatedSprite3DFrameChangedSignalFn func()

func (me *AnimatedSprite3D) ConnectFrameChanged(subs SignalSubscribers, fn AnimatedSprite3DFrameChangedSignalFn) {
	sig := StringNameFromStr("frame_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite3D) DisconnectFrameChanged(subs SignalSubscribers, fn AnimatedSprite3DFrameChangedSignalFn) {
	sig := StringNameFromStr("frame_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimatedSprite3DAnimationLoopedSignalFn func()

func (me *AnimatedSprite3D) ConnectAnimationLooped(subs SignalSubscribers, fn AnimatedSprite3DAnimationLoopedSignalFn) {
	sig := StringNameFromStr("animation_looped")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite3D) DisconnectAnimationLooped(subs SignalSubscribers, fn AnimatedSprite3DAnimationLoopedSignalFn) {
	sig := StringNameFromStr("animation_looped")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimatedSprite3DAnimationFinishedSignalFn func()

func (me *AnimatedSprite3D) ConnectAnimationFinished(subs SignalSubscribers, fn AnimatedSprite3DAnimationFinishedSignalFn) {
	sig := StringNameFromStr("animation_finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimatedSprite3D) DisconnectAnimationFinished(subs SignalSubscribers, fn AnimatedSprite3DAnimationFinishedSignalFn) {
	sig := StringNameFromStr("animation_finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
