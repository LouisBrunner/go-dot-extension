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

type ptrsForAnimationPlayerList struct {
	fnAnimationSetNext            gdc.MethodBindPtr
	fnAnimationGetNext            gdc.MethodBindPtr
	fnSetBlendTime                gdc.MethodBindPtr
	fnGetBlendTime                gdc.MethodBindPtr
	fnSetDefaultBlendTime         gdc.MethodBindPtr
	fnGetDefaultBlendTime         gdc.MethodBindPtr
	fnPlay                        gdc.MethodBindPtr
	fnPlayBackwards               gdc.MethodBindPtr
	fnPlayWithCapture             gdc.MethodBindPtr
	fnPause                       gdc.MethodBindPtr
	fnStop                        gdc.MethodBindPtr
	fnIsPlaying                   gdc.MethodBindPtr
	fnSetCurrentAnimation         gdc.MethodBindPtr
	fnGetCurrentAnimation         gdc.MethodBindPtr
	fnSetAssignedAnimation        gdc.MethodBindPtr
	fnGetAssignedAnimation        gdc.MethodBindPtr
	fnQueue                       gdc.MethodBindPtr
	fnGetQueue                    gdc.MethodBindPtr
	fnClearQueue                  gdc.MethodBindPtr
	fnSetSpeedScale               gdc.MethodBindPtr
	fnGetSpeedScale               gdc.MethodBindPtr
	fnGetPlayingSpeed             gdc.MethodBindPtr
	fnSetAutoplay                 gdc.MethodBindPtr
	fnGetAutoplay                 gdc.MethodBindPtr
	fnSetMovieQuitOnFinishEnabled gdc.MethodBindPtr
	fnIsMovieQuitOnFinishEnabled  gdc.MethodBindPtr
	fnGetCurrentAnimationPosition gdc.MethodBindPtr
	fnGetCurrentAnimationLength   gdc.MethodBindPtr
	fnSeek                        gdc.MethodBindPtr
	fnSetProcessCallback          gdc.MethodBindPtr
	fnGetProcessCallback          gdc.MethodBindPtr
	fnSetMethodCallMode           gdc.MethodBindPtr
	fnGetMethodCallMode           gdc.MethodBindPtr
	fnSetRoot                     gdc.MethodBindPtr
	fnGetRoot                     gdc.MethodBindPtr
}

var ptrsForAnimationPlayer ptrsForAnimationPlayerList

func initAnimationPlayerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationPlayer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("animation_set_next")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnAnimationSetNext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("animation_get_next")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnAnimationGetNext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965194235))
	}
	{
		methodName := StringNameFromStr("set_blend_time")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetBlendTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3231131886))
	}
	{
		methodName := StringNameFromStr("get_blend_time")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetBlendTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1958752504))
	}
	{
		methodName := StringNameFromStr("set_default_blend_time")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetDefaultBlendTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_default_blend_time")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetDefaultBlendTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("play")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnPlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3697947785))
	}
	{
		methodName := StringNameFromStr("play_backwards")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnPlayBackwards = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3890664824))
	}
	{
		methodName := StringNameFromStr("play_with_capture")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnPlayWithCapture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3265781472))
	}
	{
		methodName := StringNameFromStr("pause")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnPause = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("stop")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 107499316))
	}
	{
		methodName := StringNameFromStr("is_playing")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnIsPlaying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_current_animation")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetCurrentAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_current_animation")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetCurrentAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_assigned_animation")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetAssignedAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_assigned_animation")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetAssignedAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("queue")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnQueue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_queue")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetQueue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
	}
	{
		methodName := StringNameFromStr("clear_queue")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnClearQueue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_speed_scale")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_speed_scale")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_playing_speed")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetPlayingSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_autoplay")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_autoplay")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_movie_quit_on_finish_enabled")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetMovieQuitOnFinishEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_movie_quit_on_finish_enabled")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnIsMovieQuitOnFinishEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_current_animation_position")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetCurrentAnimationPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_current_animation_length")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetCurrentAnimationLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("seek")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSeek = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1807872683))
	}
	{
		methodName := StringNameFromStr("set_process_callback")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetProcessCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1663839457))
	}
	{
		methodName := StringNameFromStr("get_process_callback")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetProcessCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4207496604))
	}
	{
		methodName := StringNameFromStr("set_method_call_mode")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetMethodCallMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3413514846))
	}
	{
		methodName := StringNameFromStr("get_method_call_mode")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetMethodCallMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3583380054))
	}
	{
		methodName := StringNameFromStr("set_root")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnSetRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_root")
		defer methodName.Destroy()
		ptrsForAnimationPlayer.fnGetRoot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}

}

type AnimationPlayer struct {
	AnimationMixer
}

func (me *AnimationPlayer) BaseClass() string {
	return "AnimationPlayer"
}

func NewAnimationPlayer() *AnimationPlayer {
	str := StringNameFromStr("AnimationPlayer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationPlayer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AnimationPlayerAnimationProcessCallback int

const (
	AnimationPlayerAnimationProcessCallbackAnimationProcessPhysics AnimationPlayerAnimationProcessCallback = 0
	AnimationPlayerAnimationProcessCallbackAnimationProcessIdle    AnimationPlayerAnimationProcessCallback = 1
	AnimationPlayerAnimationProcessCallbackAnimationProcessManual  AnimationPlayerAnimationProcessCallback = 2
)

type AnimationPlayerAnimationMethodCallMode int

const (
	AnimationPlayerAnimationMethodCallModeAnimationMethodCallDeferred  AnimationPlayerAnimationMethodCallMode = 0
	AnimationPlayerAnimationMethodCallModeAnimationMethodCallImmediate AnimationPlayerAnimationMethodCallMode = 1
)

func (me *AnimationPlayer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationPlayer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationPlayer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimationPlayer) AnimationSetNext(animation_from StringName, animation_to StringName) {
	cargs := []gdc.ConstTypePtr{animation_from.AsCTypePtr(), animation_to.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnAnimationSetNext), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) AnimationGetNext(animation_from StringName) StringName {
	cargs := []gdc.ConstTypePtr{animation_from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnAnimationGetNext), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationPlayer) SetBlendTime(animation_from StringName, animation_to StringName, sec float64) {
	cargs := []gdc.ConstTypePtr{animation_from.AsCTypePtr(), animation_to.AsCTypePtr(), gdc.ConstTypePtr(&sec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetBlendTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetBlendTime(animation_from StringName, animation_to StringName) float64 {
	cargs := []gdc.ConstTypePtr{animation_from.AsCTypePtr(), animation_to.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetBlendTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationPlayer) SetDefaultBlendTime(sec float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetDefaultBlendTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetDefaultBlendTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetDefaultBlendTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationPlayer) Play(name StringName, custom_blend float64, custom_speed float64, from_end bool) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&custom_blend), gdc.ConstTypePtr(&custom_speed), gdc.ConstTypePtr(&from_end)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnPlay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) PlayBackwards(name StringName, custom_blend float64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&custom_blend)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnPlayBackwards), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) PlayWithCapture(name StringName, duration float64, custom_blend float64, custom_speed float64, from_end bool, trans_type TweenTransitionType, ease_type TweenEaseType) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&duration), gdc.ConstTypePtr(&custom_blend), gdc.ConstTypePtr(&custom_speed), gdc.ConstTypePtr(&from_end), gdc.ConstTypePtr(&trans_type), gdc.ConstTypePtr(&ease_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnPlayWithCapture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) Pause() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnPause), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) Stop(keep_state bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keep_state)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) IsPlaying() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnIsPlaying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationPlayer) SetCurrentAnimation(animation String) {
	cargs := []gdc.ConstTypePtr{animation.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetCurrentAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetCurrentAnimation() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetCurrentAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationPlayer) SetAssignedAnimation(animation String) {
	cargs := []gdc.ConstTypePtr{animation.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetAssignedAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetAssignedAnimation() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetAssignedAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationPlayer) Queue(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnQueue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetQueue() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetQueue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationPlayer) ClearQueue() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnClearQueue), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) SetSpeedScale(speed float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetSpeedScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationPlayer) GetPlayingSpeed() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetPlayingSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationPlayer) SetAutoplay(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetAutoplay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetAutoplay() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetAutoplay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationPlayer) SetMovieQuitOnFinishEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetMovieQuitOnFinishEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) IsMovieQuitOnFinishEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnIsMovieQuitOnFinishEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationPlayer) GetCurrentAnimationPosition() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetCurrentAnimationPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationPlayer) GetCurrentAnimationLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetCurrentAnimationLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationPlayer) Seek(seconds float64, update bool, update_only bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds), gdc.ConstTypePtr(&update), gdc.ConstTypePtr(&update_only)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSeek), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) SetProcessCallback(mode AnimationPlayerAnimationProcessCallback) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetProcessCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetProcessCallback() AnimationPlayerAnimationProcessCallback {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationPlayerAnimationProcessCallback

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetProcessCallback), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationPlayer) SetMethodCallMode(mode AnimationPlayerAnimationMethodCallMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetMethodCallMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetMethodCallMode() AnimationPlayerAnimationMethodCallMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationPlayerAnimationMethodCallMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetMethodCallMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationPlayer) SetRoot(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnSetRoot), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationPlayer) GetRoot() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationPlayer.fnGetRoot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationPlayerCurrentAnimationChangedSignalFn func(name String)

func (me *AnimationPlayer) ConnectCurrentAnimationChanged(subs SignalSubscribers, fn AnimationPlayerCurrentAnimationChangedSignalFn) {
	sig := StringNameFromStr("current_animation_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectCurrentAnimationChanged(subs SignalSubscribers, fn AnimationPlayerCurrentAnimationChangedSignalFn) {
	sig := StringNameFromStr("current_animation_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationPlayerAnimationChangedSignalFn func(old_name StringName, new_name StringName)

func (me *AnimationPlayer) ConnectAnimationChanged(subs SignalSubscribers, fn AnimationPlayerAnimationChangedSignalFn) {
	sig := StringNameFromStr("animation_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationPlayer) DisconnectAnimationChanged(subs SignalSubscribers, fn AnimationPlayerAnimationChangedSignalFn) {
	sig := StringNameFromStr("animation_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
