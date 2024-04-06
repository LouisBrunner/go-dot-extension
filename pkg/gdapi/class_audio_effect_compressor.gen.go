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

type ptrsForAudioEffectCompressorList struct {
	fnSetThreshold gdc.MethodBindPtr
	fnGetThreshold gdc.MethodBindPtr
	fnSetRatio     gdc.MethodBindPtr
	fnGetRatio     gdc.MethodBindPtr
	fnSetGain      gdc.MethodBindPtr
	fnGetGain      gdc.MethodBindPtr
	fnSetAttackUs  gdc.MethodBindPtr
	fnGetAttackUs  gdc.MethodBindPtr
	fnSetReleaseMs gdc.MethodBindPtr
	fnGetReleaseMs gdc.MethodBindPtr
	fnSetMix       gdc.MethodBindPtr
	fnGetMix       gdc.MethodBindPtr
	fnSetSidechain gdc.MethodBindPtr
	fnGetSidechain gdc.MethodBindPtr
}

var ptrsForAudioEffectCompressor ptrsForAudioEffectCompressorList

func initAudioEffectCompressorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectCompressor")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_threshold")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnSetThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_threshold")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnGetThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_ratio")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnSetRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_ratio")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnGetRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_gain")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnSetGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_gain")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnGetGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_attack_us")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnSetAttackUs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_attack_us")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnGetAttackUs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_release_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnSetReleaseMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_release_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnGetReleaseMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_mix")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnSetMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_mix")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnGetMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_sidechain")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnSetSidechain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_sidechain")
		defer methodName.Destroy()
		ptrsForAudioEffectCompressor.fnGetSidechain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
}

type AudioEffectCompressor struct {
	AudioEffect
}

func (me *AudioEffectCompressor) BaseClass() string {
	return "AudioEffectCompressor"
}

func NewAudioEffectCompressor() *AudioEffectCompressor {
	str := StringNameFromStr("AudioEffectCompressor") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectCompressor{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectCompressor) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectCompressor) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectCompressor) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectCompressor) SetThreshold(threshold float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnSetThreshold), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectCompressor) GetThreshold() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnGetThreshold), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCompressor) SetRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnSetRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectCompressor) GetRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnGetRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCompressor) SetGain(gain float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gain)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnSetGain), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectCompressor) GetGain() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnGetGain), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCompressor) SetAttackUs(attack_us float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&attack_us)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnSetAttackUs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectCompressor) GetAttackUs() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnGetAttackUs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCompressor) SetReleaseMs(release_ms float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&release_ms)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnSetReleaseMs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectCompressor) GetReleaseMs() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnGetReleaseMs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCompressor) SetMix(mix float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnSetMix), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectCompressor) GetMix() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnGetMix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCompressor) SetSidechain(sidechain StringName) {
	cargs := []gdc.ConstTypePtr{sidechain.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnSetSidechain), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectCompressor) GetSidechain() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCompressor.fnGetSidechain), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
