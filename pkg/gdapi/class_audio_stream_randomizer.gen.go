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

type ptrsForAudioStreamRandomizerList struct {
	fnAddStream                  gdc.MethodBindPtr
	fnMoveStream                 gdc.MethodBindPtr
	fnRemoveStream               gdc.MethodBindPtr
	fnSetStream                  gdc.MethodBindPtr
	fnGetStream                  gdc.MethodBindPtr
	fnSetStreamProbabilityWeight gdc.MethodBindPtr
	fnGetStreamProbabilityWeight gdc.MethodBindPtr
	fnSetStreamsCount            gdc.MethodBindPtr
	fnGetStreamsCount            gdc.MethodBindPtr
	fnSetRandomPitch             gdc.MethodBindPtr
	fnGetRandomPitch             gdc.MethodBindPtr
	fnSetRandomVolumeOffsetDb    gdc.MethodBindPtr
	fnGetRandomVolumeOffsetDb    gdc.MethodBindPtr
	fnSetPlaybackMode            gdc.MethodBindPtr
	fnGetPlaybackMode            gdc.MethodBindPtr
}

var ptrsForAudioStreamRandomizer ptrsForAudioStreamRandomizerList

func initAudioStreamRandomizerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamRandomizer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnAddStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1892018854))
	}
	{
		methodName := StringNameFromStr("move_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnMoveStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("remove_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnRemoveStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnSetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 111075094))
	}
	{
		methodName := StringNameFromStr("get_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnGetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2739380747))
	}
	{
		methodName := StringNameFromStr("set_stream_probability_weight")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnSetStreamProbabilityWeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_stream_probability_weight")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnGetStreamProbabilityWeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_streams_count")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnSetStreamsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_streams_count")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnGetStreamsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_random_pitch")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnSetRandomPitch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_random_pitch")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnGetRandomPitch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_random_volume_offset_db")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnSetRandomVolumeOffsetDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_random_volume_offset_db")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnGetRandomVolumeOffsetDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_playback_mode")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnSetPlaybackMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3950967023))
	}
	{
		methodName := StringNameFromStr("get_playback_mode")
		defer methodName.Destroy()
		ptrsForAudioStreamRandomizer.fnGetPlaybackMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3943055077))
	}

}

type AudioStreamRandomizer struct {
	AudioStream
}

func (me *AudioStreamRandomizer) BaseClass() string {
	return "AudioStreamRandomizer"
}

func NewAudioStreamRandomizer() *AudioStreamRandomizer {
	str := StringNameFromStr("AudioStreamRandomizer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamRandomizer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AudioStreamRandomizerPlaybackMode int

const (
	AudioStreamRandomizerPlaybackModePlaybackRandomNoRepeats AudioStreamRandomizerPlaybackMode = 0
	AudioStreamRandomizerPlaybackModePlaybackRandom          AudioStreamRandomizerPlaybackMode = 1
	AudioStreamRandomizerPlaybackModePlaybackSequential      AudioStreamRandomizerPlaybackMode = 2
)

func (me *AudioStreamRandomizer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamRandomizer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamRandomizer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamRandomizer) AddStream(index int64, stream AudioStream, weight float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), stream.AsCTypePtr(), gdc.ConstTypePtr(&weight)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnAddStream), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamRandomizer) MoveStream(index_from int64, index_to int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index_from), gdc.ConstTypePtr(&index_to)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnMoveStream), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamRandomizer) RemoveStream(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnRemoveStream), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamRandomizer) SetStream(index int64, stream AudioStream) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), stream.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnSetStream), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamRandomizer) GetStream(index int64) AudioStream {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAudioStream()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnGetStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AudioStreamRandomizer) SetStreamProbabilityWeight(index int64, weight float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&weight)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnSetStreamProbabilityWeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamRandomizer) GetStreamProbabilityWeight(index int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnGetStreamProbabilityWeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamRandomizer) SetStreamsCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnSetStreamsCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamRandomizer) GetStreamsCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnGetStreamsCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamRandomizer) SetRandomPitch(scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnSetRandomPitch), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamRandomizer) GetRandomPitch() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnGetRandomPitch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamRandomizer) SetRandomVolumeOffsetDb(db_offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&db_offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnSetRandomVolumeOffsetDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamRandomizer) GetRandomVolumeOffsetDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnGetRandomVolumeOffsetDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamRandomizer) SetPlaybackMode(mode AudioStreamRandomizerPlaybackMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnSetPlaybackMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamRandomizer) GetPlaybackMode() AudioStreamRandomizerPlaybackMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioStreamRandomizerPlaybackMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamRandomizer.fnGetPlaybackMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
