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

type ptrsForVideoStreamPlayerList struct {
	fnSetStream         gdc.MethodBindPtr
	fnGetStream         gdc.MethodBindPtr
	fnPlay              gdc.MethodBindPtr
	fnStop              gdc.MethodBindPtr
	fnIsPlaying         gdc.MethodBindPtr
	fnSetPaused         gdc.MethodBindPtr
	fnIsPaused          gdc.MethodBindPtr
	fnSetLoop           gdc.MethodBindPtr
	fnHasLoop           gdc.MethodBindPtr
	fnSetVolume         gdc.MethodBindPtr
	fnGetVolume         gdc.MethodBindPtr
	fnSetVolumeDb       gdc.MethodBindPtr
	fnGetVolumeDb       gdc.MethodBindPtr
	fnSetAudioTrack     gdc.MethodBindPtr
	fnGetAudioTrack     gdc.MethodBindPtr
	fnGetStreamName     gdc.MethodBindPtr
	fnGetStreamLength   gdc.MethodBindPtr
	fnSetStreamPosition gdc.MethodBindPtr
	fnGetStreamPosition gdc.MethodBindPtr
	fnSetAutoplay       gdc.MethodBindPtr
	fnHasAutoplay       gdc.MethodBindPtr
	fnSetExpand         gdc.MethodBindPtr
	fnHasExpand         gdc.MethodBindPtr
	fnSetBufferingMsec  gdc.MethodBindPtr
	fnGetBufferingMsec  gdc.MethodBindPtr
	fnSetBus            gdc.MethodBindPtr
	fnGetBus            gdc.MethodBindPtr
	fnGetVideoTexture   gdc.MethodBindPtr
}

var ptrsForVideoStreamPlayer ptrsForVideoStreamPlayerList

func initVideoStreamPlayerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VideoStreamPlayer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_stream")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2317102564))
	}
	{
		methodName := StringNameFromStr("get_stream")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 438621487))
	}
	{
		methodName := StringNameFromStr("play")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnPlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("stop")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_playing")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnIsPlaying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_paused")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_paused")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnIsPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_loop")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_loop")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnHasLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_volume")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetVolume = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_volume")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetVolume = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_volume_db")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_volume_db")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_audio_track")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetAudioTrack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_audio_track")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetAudioTrack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_stream_name")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetStreamName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_stream_length")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetStreamLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_stream_position")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetStreamPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_stream_position")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetStreamPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_autoplay")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_autoplay")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnHasAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_expand")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetExpand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_expand")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnHasExpand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_buffering_msec")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetBufferingMsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_buffering_msec")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetBufferingMsec = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_bus")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnSetBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_bus")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("get_video_texture")
		defer methodName.Destroy()
		ptrsForVideoStreamPlayer.fnGetVideoTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
}

type VideoStreamPlayer struct {
	Control
}

func (me *VideoStreamPlayer) BaseClass() string {
	return "VideoStreamPlayer"
}

func NewVideoStreamPlayer() *VideoStreamPlayer {
	str := StringNameFromStr("VideoStreamPlayer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VideoStreamPlayer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VideoStreamPlayer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VideoStreamPlayer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VideoStreamPlayer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VideoStreamPlayer) SetStream(stream VideoStream) {
	cargs := []gdc.ConstTypePtr{stream.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetStream), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) GetStream() VideoStream {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVideoStream()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VideoStreamPlayer) Play() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnPlay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) Stop() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) IsPlaying() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnIsPlaying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetPaused(paused bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paused)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetPaused), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) IsPaused() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnIsPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetLoop(loop bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) HasLoop() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnHasLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetVolume(volume float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&volume)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetVolume), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) GetVolume() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetVolume), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetVolumeDb(db float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&db)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetVolumeDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) GetVolumeDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetVolumeDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetAudioTrack(track int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetAudioTrack), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) GetAudioTrack() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetAudioTrack), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) GetStreamName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetStreamName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VideoStreamPlayer) GetStreamLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetStreamLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetStreamPosition(position float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetStreamPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) GetStreamPosition() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetStreamPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetAutoplay(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetAutoplay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) HasAutoplay() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnHasAutoplay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetExpand(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetExpand), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) HasExpand() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnHasExpand), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetBufferingMsec(msec int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetBufferingMsec), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) GetBufferingMsec() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetBufferingMsec), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VideoStreamPlayer) SetBus(bus StringName) {
	cargs := []gdc.ConstTypePtr{bus.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnSetBus), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VideoStreamPlayer) GetBus() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetBus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VideoStreamPlayer) GetVideoTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayer.fnGetVideoTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type VideoStreamPlayerFinishedSignalFn func()

func (me *VideoStreamPlayer) ConnectFinished(subs SignalSubscribers, fn VideoStreamPlayerFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *VideoStreamPlayer) DisconnectFinished(subs SignalSubscribers, fn VideoStreamPlayerFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
