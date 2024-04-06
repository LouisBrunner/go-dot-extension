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

type ptrsForAudioEffectPitchShiftList struct {
	fnSetPitchScale   gdc.MethodBindPtr
	fnGetPitchScale   gdc.MethodBindPtr
	fnSetOversampling gdc.MethodBindPtr
	fnGetOversampling gdc.MethodBindPtr
	fnSetFftSize      gdc.MethodBindPtr
	fnGetFftSize      gdc.MethodBindPtr
}

var ptrsForAudioEffectPitchShift ptrsForAudioEffectPitchShiftList

func initAudioEffectPitchShiftPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectPitchShift")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_pitch_scale")
		defer methodName.Destroy()
		ptrsForAudioEffectPitchShift.fnSetPitchScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pitch_scale")
		defer methodName.Destroy()
		ptrsForAudioEffectPitchShift.fnGetPitchScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_oversampling")
		defer methodName.Destroy()
		ptrsForAudioEffectPitchShift.fnSetOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_oversampling")
		defer methodName.Destroy()
		ptrsForAudioEffectPitchShift.fnGetOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_fft_size")
		defer methodName.Destroy()
		ptrsForAudioEffectPitchShift.fnSetFftSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323518741))
	}
	{
		methodName := StringNameFromStr("get_fft_size")
		defer methodName.Destroy()
		ptrsForAudioEffectPitchShift.fnGetFftSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2361246789))
	}
}

type AudioEffectPitchShift struct {
	AudioEffect
}

func (me *AudioEffectPitchShift) BaseClass() string {
	return "AudioEffectPitchShift"
}

func NewAudioEffectPitchShift() *AudioEffectPitchShift {
	str := StringNameFromStr("AudioEffectPitchShift") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectPitchShift{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AudioEffectPitchShiftFFTSize int

const (
	AudioEffectPitchShiftFFTSizeFftSize256  AudioEffectPitchShiftFFTSize = 0
	AudioEffectPitchShiftFFTSizeFftSize512  AudioEffectPitchShiftFFTSize = 1
	AudioEffectPitchShiftFFTSizeFftSize1024 AudioEffectPitchShiftFFTSize = 2
	AudioEffectPitchShiftFFTSizeFftSize2048 AudioEffectPitchShiftFFTSize = 3
	AudioEffectPitchShiftFFTSizeFftSize4096 AudioEffectPitchShiftFFTSize = 4
	AudioEffectPitchShiftFFTSizeFftSizeMax  AudioEffectPitchShiftFFTSize = 5
)

func (me *AudioEffectPitchShift) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectPitchShift) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectPitchShift) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectPitchShift) SetPitchScale(rate float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rate)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPitchShift.fnSetPitchScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectPitchShift) GetPitchScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPitchShift.fnGetPitchScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectPitchShift) SetOversampling(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPitchShift.fnSetOversampling), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectPitchShift) GetOversampling() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPitchShift.fnGetOversampling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectPitchShift) SetFftSize(size AudioEffectPitchShiftFFTSize) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPitchShift.fnSetFftSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectPitchShift) GetFftSize() AudioEffectPitchShiftFFTSize {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioEffectPitchShiftFFTSize

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPitchShift.fnGetFftSize), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
