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

type ptrsForAudioEffectChorusList struct {
	fnSetVoiceCount    gdc.MethodBindPtr
	fnGetVoiceCount    gdc.MethodBindPtr
	fnSetVoiceDelayMs  gdc.MethodBindPtr
	fnGetVoiceDelayMs  gdc.MethodBindPtr
	fnSetVoiceRateHz   gdc.MethodBindPtr
	fnGetVoiceRateHz   gdc.MethodBindPtr
	fnSetVoiceDepthMs  gdc.MethodBindPtr
	fnGetVoiceDepthMs  gdc.MethodBindPtr
	fnSetVoiceLevelDb  gdc.MethodBindPtr
	fnGetVoiceLevelDb  gdc.MethodBindPtr
	fnSetVoiceCutoffHz gdc.MethodBindPtr
	fnGetVoiceCutoffHz gdc.MethodBindPtr
	fnSetVoicePan      gdc.MethodBindPtr
	fnGetVoicePan      gdc.MethodBindPtr
	fnSetWet           gdc.MethodBindPtr
	fnGetWet           gdc.MethodBindPtr
	fnSetDry           gdc.MethodBindPtr
	fnGetDry           gdc.MethodBindPtr
}

var ptrsForAudioEffectChorus ptrsForAudioEffectChorusList

func initAudioEffectChorusPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectChorus")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_voice_count")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnSetVoiceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_voice_count")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnGetVoiceCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_voice_delay_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnSetVoiceDelayMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_voice_delay_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnGetVoiceDelayMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_voice_rate_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnSetVoiceRateHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_voice_rate_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnGetVoiceRateHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_voice_depth_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnSetVoiceDepthMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_voice_depth_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnGetVoiceDepthMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_voice_level_db")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnSetVoiceLevelDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_voice_level_db")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnGetVoiceLevelDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_voice_cutoff_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnSetVoiceCutoffHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_voice_cutoff_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnGetVoiceCutoffHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_voice_pan")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnSetVoicePan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_voice_pan")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnGetVoicePan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_wet")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnSetWet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_wet")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnGetWet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_dry")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnSetDry = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_dry")
		defer methodName.Destroy()
		ptrsForAudioEffectChorus.fnGetDry = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type AudioEffectChorus struct {
	AudioEffect
}

func (me *AudioEffectChorus) BaseClass() string {
	return "AudioEffectChorus"
}

func NewAudioEffectChorus() *AudioEffectChorus {
	str := StringNameFromStr("AudioEffectChorus") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectChorus{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectChorus) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectChorus) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectChorus) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectChorus) SetVoiceCount(voices int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voices)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnSetVoiceCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectChorus) GetVoiceCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnGetVoiceCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectChorus) SetVoiceDelayMs(voice_idx int64, delay_ms float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&delay_ms)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnSetVoiceDelayMs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectChorus) GetVoiceDelayMs(voice_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&voice_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnGetVoiceDelayMs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectChorus) SetVoiceRateHz(voice_idx int64, rate_hz float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&rate_hz)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnSetVoiceRateHz), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectChorus) GetVoiceRateHz(voice_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&voice_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnGetVoiceRateHz), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectChorus) SetVoiceDepthMs(voice_idx int64, depth_ms float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&depth_ms)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnSetVoiceDepthMs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectChorus) GetVoiceDepthMs(voice_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&voice_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnGetVoiceDepthMs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectChorus) SetVoiceLevelDb(voice_idx int64, level_db float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&level_db)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnSetVoiceLevelDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectChorus) GetVoiceLevelDb(voice_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&voice_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnGetVoiceLevelDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectChorus) SetVoiceCutoffHz(voice_idx int64, cutoff_hz float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&cutoff_hz)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnSetVoiceCutoffHz), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectChorus) GetVoiceCutoffHz(voice_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&voice_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnGetVoiceCutoffHz), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectChorus) SetVoicePan(voice_idx int64, pan float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&pan)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnSetVoicePan), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectChorus) GetVoicePan(voice_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&voice_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnGetVoicePan), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectChorus) SetWet(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnSetWet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectChorus) GetWet() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnGetWet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectChorus) SetDry(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnSetDry), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectChorus) GetDry() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectChorus.fnGetDry), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
