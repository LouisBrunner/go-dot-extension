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

type ptrsForAudioStreamInteractiveList struct {
	fnSetClipCount                gdc.MethodBindPtr
	fnGetClipCount                gdc.MethodBindPtr
	fnSetInitialClip              gdc.MethodBindPtr
	fnGetInitialClip              gdc.MethodBindPtr
	fnSetClipName                 gdc.MethodBindPtr
	fnGetClipName                 gdc.MethodBindPtr
	fnSetClipStream               gdc.MethodBindPtr
	fnGetClipStream               gdc.MethodBindPtr
	fnSetClipAutoAdvance          gdc.MethodBindPtr
	fnGetClipAutoAdvance          gdc.MethodBindPtr
	fnSetClipAutoAdvanceNextClip  gdc.MethodBindPtr
	fnGetClipAutoAdvanceNextClip  gdc.MethodBindPtr
	fnAddTransition               gdc.MethodBindPtr
	fnHasTransition               gdc.MethodBindPtr
	fnEraseTransition             gdc.MethodBindPtr
	fnGetTransitionList           gdc.MethodBindPtr
	fnGetTransitionFromTime       gdc.MethodBindPtr
	fnGetTransitionToTime         gdc.MethodBindPtr
	fnGetTransitionFadeMode       gdc.MethodBindPtr
	fnGetTransitionFadeBeats      gdc.MethodBindPtr
	fnIsTransitionUsingFillerClip gdc.MethodBindPtr
	fnGetTransitionFillerClip     gdc.MethodBindPtr
	fnIsTransitionHoldingPrevious gdc.MethodBindPtr
}

var ptrsForAudioStreamInteractive ptrsForAudioStreamInteractiveList

func initAudioStreamInteractivePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamInteractive")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_clip_count")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnSetClipCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_clip_count")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetClipCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_initial_clip")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnSetInitialClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_initial_clip")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetInitialClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_clip_name")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnSetClipName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780747571))
	}
	{
		methodName := StringNameFromStr("get_clip_name")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetClipName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("set_clip_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnSetClipStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 111075094))
	}
	{
		methodName := StringNameFromStr("get_clip_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetClipStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2739380747))
	}
	{
		methodName := StringNameFromStr("set_clip_auto_advance")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnSetClipAutoAdvance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 57217598))
	}
	{
		methodName := StringNameFromStr("get_clip_auto_advance")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetClipAutoAdvance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1778634807))
	}
	{
		methodName := StringNameFromStr("set_clip_auto_advance_next_clip")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnSetClipAutoAdvanceNextClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_clip_auto_advance_next_clip")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetClipAutoAdvanceNextClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("add_transition")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnAddTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1630280552))
	}
	{
		methodName := StringNameFromStr("has_transition")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnHasTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}
	{
		methodName := StringNameFromStr("erase_transition")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnEraseTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_transition_list")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetTransitionList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
	}
	{
		methodName := StringNameFromStr("get_transition_from_time")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetTransitionFromTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3453338158))
	}
	{
		methodName := StringNameFromStr("get_transition_to_time")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetTransitionToTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1369651373))
	}
	{
		methodName := StringNameFromStr("get_transition_fade_mode")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetTransitionFadeMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4065396087))
	}
	{
		methodName := StringNameFromStr("get_transition_fade_beats")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetTransitionFadeBeats = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
	}
	{
		methodName := StringNameFromStr("is_transition_using_filler_clip")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnIsTransitionUsingFillerClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}
	{
		methodName := StringNameFromStr("get_transition_filler_clip")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnGetTransitionFillerClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("is_transition_holding_previous")
		defer methodName.Destroy()
		ptrsForAudioStreamInteractive.fnIsTransitionHoldingPrevious = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}

}

type AudioStreamInteractive struct {
	AudioStream
}

func (me *AudioStreamInteractive) BaseClass() string {
	return "AudioStreamInteractive"
}

func NewAudioStreamInteractive() *AudioStreamInteractive {
	str := StringNameFromStr("AudioStreamInteractive") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamInteractive{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	AudioStreamInteractiveClipAny = -1
)

// Enums

type AudioStreamInteractiveTransitionFromTime int

const (
	AudioStreamInteractiveTransitionFromTimeTransitionFromTimeImmediate AudioStreamInteractiveTransitionFromTime = 0
	AudioStreamInteractiveTransitionFromTimeTransitionFromTimeNextBeat  AudioStreamInteractiveTransitionFromTime = 1
	AudioStreamInteractiveTransitionFromTimeTransitionFromTimeNextBar   AudioStreamInteractiveTransitionFromTime = 2
	AudioStreamInteractiveTransitionFromTimeTransitionFromTimeEnd       AudioStreamInteractiveTransitionFromTime = 3
)

type AudioStreamInteractiveTransitionToTime int

const (
	AudioStreamInteractiveTransitionToTimeTransitionToTimeSamePosition AudioStreamInteractiveTransitionToTime = 0
	AudioStreamInteractiveTransitionToTimeTransitionToTimeStart        AudioStreamInteractiveTransitionToTime = 1
)

type AudioStreamInteractiveFadeMode int

const (
	AudioStreamInteractiveFadeModeFadeDisabled  AudioStreamInteractiveFadeMode = 0
	AudioStreamInteractiveFadeModeFadeIn        AudioStreamInteractiveFadeMode = 1
	AudioStreamInteractiveFadeModeFadeOut       AudioStreamInteractiveFadeMode = 2
	AudioStreamInteractiveFadeModeFadeCross     AudioStreamInteractiveFadeMode = 3
	AudioStreamInteractiveFadeModeFadeAutomatic AudioStreamInteractiveFadeMode = 4
)

type AudioStreamInteractiveAutoAdvanceMode int

const (
	AudioStreamInteractiveAutoAdvanceModeAutoAdvanceDisabled     AudioStreamInteractiveAutoAdvanceMode = 0
	AudioStreamInteractiveAutoAdvanceModeAutoAdvanceEnabled      AudioStreamInteractiveAutoAdvanceMode = 1
	AudioStreamInteractiveAutoAdvanceModeAutoAdvanceReturnToHold AudioStreamInteractiveAutoAdvanceMode = 2
)

func (me *AudioStreamInteractive) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamInteractive) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamInteractive) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamInteractive) SetClipCount(clip_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnSetClipCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamInteractive) GetClipCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetClipCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamInteractive) SetInitialClip(clip_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnSetInitialClip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamInteractive) GetInitialClip() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetInitialClip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamInteractive) SetClipName(clip_index int64, name StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnSetClipName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamInteractive) GetClipName(clip_index int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&clip_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetClipName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AudioStreamInteractive) SetClipStream(clip_index int64, stream AudioStream) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index), stream.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnSetClipStream), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamInteractive) GetClipStream(clip_index int64) AudioStream {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAudioStream()
	pinner.Pin(&clip_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetClipStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AudioStreamInteractive) SetClipAutoAdvance(clip_index int64, mode AudioStreamInteractiveAutoAdvanceMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index), gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnSetClipAutoAdvance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamInteractive) GetClipAutoAdvance(clip_index int64) AudioStreamInteractiveAutoAdvanceMode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioStreamInteractiveAutoAdvanceMode
	pinner.Pin(&clip_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetClipAutoAdvance), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AudioStreamInteractive) SetClipAutoAdvanceNextClip(clip_index int64, auto_advance_next_clip int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index), gdc.ConstTypePtr(&auto_advance_next_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnSetClipAutoAdvanceNextClip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamInteractive) GetClipAutoAdvanceNextClip(clip_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&clip_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetClipAutoAdvanceNextClip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamInteractive) AddTransition(from_clip int64, to_clip int64, from_time AudioStreamInteractiveTransitionFromTime, to_time AudioStreamInteractiveTransitionToTime, fade_mode AudioStreamInteractiveFadeMode, fade_beats float64, use_filler_clip bool, filler_clip int64, hold_previous bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip), gdc.ConstTypePtr(&from_time), gdc.ConstTypePtr(&to_time), gdc.ConstTypePtr(&fade_mode), gdc.ConstTypePtr(&fade_beats), gdc.ConstTypePtr(&use_filler_clip), gdc.ConstTypePtr(&filler_clip), gdc.ConstTypePtr(&hold_previous)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnAddTransition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamInteractive) HasTransition(from_clip int64, to_clip int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&from_clip)
	pinner.Pin(&to_clip)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnHasTransition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamInteractive) EraseTransition(from_clip int64, to_clip int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnEraseTransition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamInteractive) GetTransitionList() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetTransitionList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AudioStreamInteractive) GetTransitionFromTime(from_clip int64, to_clip int64) AudioStreamInteractiveTransitionFromTime {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioStreamInteractiveTransitionFromTime
	pinner.Pin(&from_clip)
	pinner.Pin(&to_clip)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetTransitionFromTime), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AudioStreamInteractive) GetTransitionToTime(from_clip int64, to_clip int64) AudioStreamInteractiveTransitionToTime {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioStreamInteractiveTransitionToTime
	pinner.Pin(&from_clip)
	pinner.Pin(&to_clip)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetTransitionToTime), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AudioStreamInteractive) GetTransitionFadeMode(from_clip int64, to_clip int64) AudioStreamInteractiveFadeMode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioStreamInteractiveFadeMode
	pinner.Pin(&from_clip)
	pinner.Pin(&to_clip)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetTransitionFadeMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AudioStreamInteractive) GetTransitionFadeBeats(from_clip int64, to_clip int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&from_clip)
	pinner.Pin(&to_clip)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetTransitionFadeBeats), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamInteractive) IsTransitionUsingFillerClip(from_clip int64, to_clip int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&from_clip)
	pinner.Pin(&to_clip)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnIsTransitionUsingFillerClip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamInteractive) GetTransitionFillerClip(from_clip int64, to_clip int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&from_clip)
	pinner.Pin(&to_clip)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnGetTransitionFillerClip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamInteractive) IsTransitionHoldingPrevious(from_clip int64, to_clip int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_clip), gdc.ConstTypePtr(&to_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&from_clip)
	pinner.Pin(&to_clip)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamInteractive.fnIsTransitionHoldingPrevious), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
