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

type ptrsForAudioEffectLimiterList struct {
	fnSetCeilingDb     gdc.MethodBindPtr
	fnGetCeilingDb     gdc.MethodBindPtr
	fnSetThresholdDb   gdc.MethodBindPtr
	fnGetThresholdDb   gdc.MethodBindPtr
	fnSetSoftClipDb    gdc.MethodBindPtr
	fnGetSoftClipDb    gdc.MethodBindPtr
	fnSetSoftClipRatio gdc.MethodBindPtr
	fnGetSoftClipRatio gdc.MethodBindPtr
}

var ptrsForAudioEffectLimiter ptrsForAudioEffectLimiterList

func initAudioEffectLimiterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectLimiter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_ceiling_db")
		defer methodName.Destroy()
		ptrsForAudioEffectLimiter.fnSetCeilingDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_ceiling_db")
		defer methodName.Destroy()
		ptrsForAudioEffectLimiter.fnGetCeilingDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_threshold_db")
		defer methodName.Destroy()
		ptrsForAudioEffectLimiter.fnSetThresholdDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_threshold_db")
		defer methodName.Destroy()
		ptrsForAudioEffectLimiter.fnGetThresholdDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_soft_clip_db")
		defer methodName.Destroy()
		ptrsForAudioEffectLimiter.fnSetSoftClipDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_soft_clip_db")
		defer methodName.Destroy()
		ptrsForAudioEffectLimiter.fnGetSoftClipDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_soft_clip_ratio")
		defer methodName.Destroy()
		ptrsForAudioEffectLimiter.fnSetSoftClipRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_soft_clip_ratio")
		defer methodName.Destroy()
		ptrsForAudioEffectLimiter.fnGetSoftClipRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type AudioEffectLimiter struct {
	AudioEffect
}

func (me *AudioEffectLimiter) BaseClass() string {
	return "AudioEffectLimiter"
}

func NewAudioEffectLimiter() *AudioEffectLimiter {
	str := StringNameFromStr("AudioEffectLimiter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectLimiter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectLimiter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectLimiter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectLimiter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectLimiter) SetCeilingDb(ceiling float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ceiling)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectLimiter.fnSetCeilingDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectLimiter) GetCeilingDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectLimiter.fnGetCeilingDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectLimiter) SetThresholdDb(threshold float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectLimiter.fnSetThresholdDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectLimiter) GetThresholdDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectLimiter.fnGetThresholdDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectLimiter) SetSoftClipDb(soft_clip float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&soft_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectLimiter.fnSetSoftClipDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectLimiter) GetSoftClipDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectLimiter.fnGetSoftClipDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectLimiter) SetSoftClipRatio(soft_clip float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&soft_clip)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectLimiter.fnSetSoftClipRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectLimiter) GetSoftClipRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectLimiter.fnGetSoftClipRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
