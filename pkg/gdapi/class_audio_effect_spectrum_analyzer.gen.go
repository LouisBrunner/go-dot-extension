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

type ptrsForAudioEffectSpectrumAnalyzerList struct {
	fnSetBufferLength gdc.MethodBindPtr
	fnGetBufferLength gdc.MethodBindPtr
	fnSetTapBackPos   gdc.MethodBindPtr
	fnGetTapBackPos   gdc.MethodBindPtr
	fnSetFftSize      gdc.MethodBindPtr
	fnGetFftSize      gdc.MethodBindPtr
}

var ptrsForAudioEffectSpectrumAnalyzer ptrsForAudioEffectSpectrumAnalyzerList

func initAudioEffectSpectrumAnalyzerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectSpectrumAnalyzer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_buffer_length")
		defer methodName.Destroy()
		ptrsForAudioEffectSpectrumAnalyzer.fnSetBufferLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_buffer_length")
		defer methodName.Destroy()
		ptrsForAudioEffectSpectrumAnalyzer.fnGetBufferLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_tap_back_pos")
		defer methodName.Destroy()
		ptrsForAudioEffectSpectrumAnalyzer.fnSetTapBackPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_tap_back_pos")
		defer methodName.Destroy()
		ptrsForAudioEffectSpectrumAnalyzer.fnGetTapBackPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fft_size")
		defer methodName.Destroy()
		ptrsForAudioEffectSpectrumAnalyzer.fnSetFftSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1202879215))
	}
	{
		methodName := StringNameFromStr("get_fft_size")
		defer methodName.Destroy()
		ptrsForAudioEffectSpectrumAnalyzer.fnGetFftSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3925405343))
	}

}

type AudioEffectSpectrumAnalyzer struct {
	AudioEffect
}

func (me *AudioEffectSpectrumAnalyzer) BaseClass() string {
	return "AudioEffectSpectrumAnalyzer"
}

func NewAudioEffectSpectrumAnalyzer() *AudioEffectSpectrumAnalyzer {
	str := StringNameFromStr("AudioEffectSpectrumAnalyzer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectSpectrumAnalyzer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AudioEffectSpectrumAnalyzerFFTSize int

const (
	AudioEffectSpectrumAnalyzerFFTSizeFftSize256  AudioEffectSpectrumAnalyzerFFTSize = 0
	AudioEffectSpectrumAnalyzerFFTSizeFftSize512  AudioEffectSpectrumAnalyzerFFTSize = 1
	AudioEffectSpectrumAnalyzerFFTSizeFftSize1024 AudioEffectSpectrumAnalyzerFFTSize = 2
	AudioEffectSpectrumAnalyzerFFTSizeFftSize2048 AudioEffectSpectrumAnalyzerFFTSize = 3
	AudioEffectSpectrumAnalyzerFFTSizeFftSize4096 AudioEffectSpectrumAnalyzerFFTSize = 4
	AudioEffectSpectrumAnalyzerFFTSizeFftSizeMax  AudioEffectSpectrumAnalyzerFFTSize = 5
)

func (me *AudioEffectSpectrumAnalyzer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectSpectrumAnalyzer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectSpectrumAnalyzer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectSpectrumAnalyzer) SetBufferLength(seconds float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectSpectrumAnalyzer.fnSetBufferLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectSpectrumAnalyzer) GetBufferLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectSpectrumAnalyzer.fnGetBufferLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectSpectrumAnalyzer) SetTapBackPos(seconds float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectSpectrumAnalyzer.fnSetTapBackPos), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectSpectrumAnalyzer) GetTapBackPos() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectSpectrumAnalyzer.fnGetTapBackPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectSpectrumAnalyzer) SetFftSize(size AudioEffectSpectrumAnalyzerFFTSize) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectSpectrumAnalyzer.fnSetFftSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectSpectrumAnalyzer) GetFftSize() AudioEffectSpectrumAnalyzerFFTSize {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioEffectSpectrumAnalyzerFFTSize

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectSpectrumAnalyzer.fnGetFftSize), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
