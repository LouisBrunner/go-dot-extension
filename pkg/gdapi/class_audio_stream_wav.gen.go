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

type ptrsForAudioStreamWAVList struct {
	fnSetData      gdc.MethodBindPtr
	fnGetData      gdc.MethodBindPtr
	fnSetFormat    gdc.MethodBindPtr
	fnGetFormat    gdc.MethodBindPtr
	fnSetLoopMode  gdc.MethodBindPtr
	fnGetLoopMode  gdc.MethodBindPtr
	fnSetLoopBegin gdc.MethodBindPtr
	fnGetLoopBegin gdc.MethodBindPtr
	fnSetLoopEnd   gdc.MethodBindPtr
	fnGetLoopEnd   gdc.MethodBindPtr
	fnSetMixRate   gdc.MethodBindPtr
	fnGetMixRate   gdc.MethodBindPtr
	fnSetStereo    gdc.MethodBindPtr
	fnIsStereo     gdc.MethodBindPtr
	fnSaveToWav    gdc.MethodBindPtr
}

var ptrsForAudioStreamWAV ptrsForAudioStreamWAVList

func initAudioStreamWAVPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamWAV")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_data")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnSetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2971499966))
	}
	{
		methodName := StringNameFromStr("get_data")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2362200018))
	}
	{
		methodName := StringNameFromStr("set_format")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnSetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 60648488))
	}
	{
		methodName := StringNameFromStr("get_format")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3151724922))
	}
	{
		methodName := StringNameFromStr("set_loop_mode")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnSetLoopMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2444882972))
	}
	{
		methodName := StringNameFromStr("get_loop_mode")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnGetLoopMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 393560655))
	}
	{
		methodName := StringNameFromStr("set_loop_begin")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnSetLoopBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_loop_begin")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnGetLoopBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_loop_end")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnSetLoopEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_loop_end")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnGetLoopEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_mix_rate")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnSetMixRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_mix_rate")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnGetMixRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_stereo")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnSetStereo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_stereo")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnIsStereo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("save_to_wav")
		defer methodName.Destroy()
		ptrsForAudioStreamWAV.fnSaveToWav = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
}

type AudioStreamWAV struct {
	AudioStream
}

func (me *AudioStreamWAV) BaseClass() string {
	return "AudioStreamWAV"
}

func NewAudioStreamWAV() *AudioStreamWAV {
	str := StringNameFromStr("AudioStreamWAV") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamWAV{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AudioStreamWAVFormat int

const (
	AudioStreamWAVFormatFormat8Bits    AudioStreamWAVFormat = 0
	AudioStreamWAVFormatFormat16Bits   AudioStreamWAVFormat = 1
	AudioStreamWAVFormatFormatImaAdpcm AudioStreamWAVFormat = 2
)

type AudioStreamWAVLoopMode int

const (
	AudioStreamWAVLoopModeLoopDisabled AudioStreamWAVLoopMode = 0
	AudioStreamWAVLoopModeLoopForward  AudioStreamWAVLoopMode = 1
	AudioStreamWAVLoopModeLoopPingpong AudioStreamWAVLoopMode = 2
	AudioStreamWAVLoopModeLoopBackward AudioStreamWAVLoopMode = 3
)

func (me *AudioStreamWAV) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamWAV) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamWAV) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamWAV) SetData(data PackedByteArray) {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnSetData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamWAV) GetData() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AudioStreamWAV) SetFormat(format AudioStreamWAVFormat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnSetFormat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamWAV) GetFormat() AudioStreamWAVFormat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioStreamWAVFormat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AudioStreamWAV) SetLoopMode(loop_mode AudioStreamWAVLoopMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnSetLoopMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamWAV) GetLoopMode() AudioStreamWAVLoopMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioStreamWAVLoopMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnGetLoopMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AudioStreamWAV) SetLoopBegin(loop_begin int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_begin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnSetLoopBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamWAV) GetLoopBegin() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnGetLoopBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamWAV) SetLoopEnd(loop_end int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_end)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnSetLoopEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamWAV) GetLoopEnd() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnGetLoopEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamWAV) SetMixRate(mix_rate int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix_rate)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnSetMixRate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamWAV) GetMixRate() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnGetMixRate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamWAV) SetStereo(stereo bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stereo)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnSetStereo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamWAV) IsStereo() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnIsStereo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamWAV) SaveToWav(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamWAV.fnSaveToWav), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
