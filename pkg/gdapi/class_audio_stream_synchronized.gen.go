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

type ptrsForAudioStreamSynchronizedList struct {
	fnSetStreamCount      gdc.MethodBindPtr
	fnGetStreamCount      gdc.MethodBindPtr
	fnSetSyncStream       gdc.MethodBindPtr
	fnGetSyncStream       gdc.MethodBindPtr
	fnSetSyncStreamVolume gdc.MethodBindPtr
	fnGetSyncStreamVolume gdc.MethodBindPtr
}

var ptrsForAudioStreamSynchronized ptrsForAudioStreamSynchronizedList

func initAudioStreamSynchronizedPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamSynchronized")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_stream_count")
		defer methodName.Destroy()
		ptrsForAudioStreamSynchronized.fnSetStreamCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_stream_count")
		defer methodName.Destroy()
		ptrsForAudioStreamSynchronized.fnGetStreamCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_sync_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamSynchronized.fnSetSyncStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 111075094))
	}
	{
		methodName := StringNameFromStr("get_sync_stream")
		defer methodName.Destroy()
		ptrsForAudioStreamSynchronized.fnGetSyncStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2739380747))
	}
	{
		methodName := StringNameFromStr("set_sync_stream_volume")
		defer methodName.Destroy()
		ptrsForAudioStreamSynchronized.fnSetSyncStreamVolume = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_sync_stream_volume")
		defer methodName.Destroy()
		ptrsForAudioStreamSynchronized.fnGetSyncStreamVolume = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}

}

type AudioStreamSynchronized struct {
	AudioStream
}

func (me *AudioStreamSynchronized) BaseClass() string {
	return "AudioStreamSynchronized"
}

func NewAudioStreamSynchronized() *AudioStreamSynchronized {
	str := StringNameFromStr("AudioStreamSynchronized") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamSynchronized{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	AudioStreamSynchronizedMaxStreams = 32
)

// Enums

func (me *AudioStreamSynchronized) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamSynchronized) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamSynchronized) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamSynchronized) SetStreamCount(stream_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamSynchronized.fnSetStreamCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamSynchronized) GetStreamCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamSynchronized.fnGetStreamCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamSynchronized) SetSyncStream(stream_index int64, audio_stream AudioStream) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream_index), audio_stream.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamSynchronized.fnSetSyncStream), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamSynchronized) GetSyncStream(stream_index int64) AudioStream {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAudioStream()
	pinner.Pin(&stream_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamSynchronized.fnGetSyncStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AudioStreamSynchronized) SetSyncStreamVolume(stream_index int64, volume_db float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream_index), gdc.ConstTypePtr(&volume_db)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamSynchronized.fnSetSyncStreamVolume), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamSynchronized) GetSyncStreamVolume(stream_index int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&stream_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamSynchronized.fnGetSyncStreamVolume), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
