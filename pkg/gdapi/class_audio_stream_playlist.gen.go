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

type ptrsForAudioStreamPlaylistList struct {
	fnSetStreamCount gdc.MethodBindPtr
	fnGetStreamCount gdc.MethodBindPtr
	fnGetBpm         gdc.MethodBindPtr
	fnSetListStream  gdc.MethodBindPtr
	fnGetListStream  gdc.MethodBindPtr
	fnSetShuffle     gdc.MethodBindPtr
	fnGetShuffle     gdc.MethodBindPtr
	fnSetFadeTime    gdc.MethodBindPtr
	fnGetFadeTime    gdc.MethodBindPtr
	fnSetLoop        gdc.MethodBindPtr
	fnHasLoop        gdc.MethodBindPtr
}

var ptrsForAudioStreamPlaylist ptrsForAudioStreamPlaylistList

func initAudioStreamPlaylistPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamPlaylist")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_stream_count")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnSetStreamCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_stream_count")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnGetStreamCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_bpm")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnGetBpm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_list_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnSetListStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 111075094))
	}
	{
		methodName := StringNameFromStr("get_list_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnGetListStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2739380747))
	}
	{
		methodName := StringNameFromStr("set_shuffle")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnSetShuffle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_shuffle")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnGetShuffle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_fade_time")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnSetFadeTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fade_time")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnGetFadeTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_loop")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnSetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_loop")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaylist.fnHasLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type AudioStreamPlaylist struct {
	AudioStream
}

func (me *AudioStreamPlaylist) BaseClass() string {
	return "AudioStreamPlaylist"
}

func NewAudioStreamPlaylist() *AudioStreamPlaylist {
	str := StringNameFromStr("AudioStreamPlaylist") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamPlaylist{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	AudioStreamPlaylistMaxStreams = 64
)

// Enums

func (me *AudioStreamPlaylist) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamPlaylist) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaylist) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamPlaylist) SetStreamCount(stream_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnSetStreamCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamPlaylist) GetStreamCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnGetStreamCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamPlaylist) GetBpm() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnGetBpm), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamPlaylist) SetListStream(stream_index int64, audio_stream AudioStream) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream_index), audio_stream.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnSetListStream), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamPlaylist) GetListStream(stream_index int64) AudioStream {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAudioStream()
	pinner.Pin(&stream_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnGetListStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AudioStreamPlaylist) SetShuffle(shuffle bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shuffle)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnSetShuffle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamPlaylist) GetShuffle() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnGetShuffle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamPlaylist) SetFadeTime(dec float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&dec)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnSetFadeTime), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamPlaylist) GetFadeTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnGetFadeTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamPlaylist) SetLoop(loop bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnSetLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamPlaylist) HasLoop() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaylist.fnHasLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
