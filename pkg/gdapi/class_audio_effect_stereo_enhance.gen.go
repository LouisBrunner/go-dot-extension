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

type ptrsForAudioEffectStereoEnhanceList struct {
	fnSetPanPullout  gdc.MethodBindPtr
	fnGetPanPullout  gdc.MethodBindPtr
	fnSetTimePullout gdc.MethodBindPtr
	fnGetTimePullout gdc.MethodBindPtr
	fnSetSurround    gdc.MethodBindPtr
	fnGetSurround    gdc.MethodBindPtr
}

var ptrsForAudioEffectStereoEnhance ptrsForAudioEffectStereoEnhanceList

func initAudioEffectStereoEnhancePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectStereoEnhance")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_pan_pullout")
		defer methodName.Destroy()
		ptrsForAudioEffectStereoEnhance.fnSetPanPullout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pan_pullout")
		defer methodName.Destroy()
		ptrsForAudioEffectStereoEnhance.fnGetPanPullout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_time_pullout")
		defer methodName.Destroy()
		ptrsForAudioEffectStereoEnhance.fnSetTimePullout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_time_pullout")
		defer methodName.Destroy()
		ptrsForAudioEffectStereoEnhance.fnGetTimePullout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_surround")
		defer methodName.Destroy()
		ptrsForAudioEffectStereoEnhance.fnSetSurround = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_surround")
		defer methodName.Destroy()
		ptrsForAudioEffectStereoEnhance.fnGetSurround = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type AudioEffectStereoEnhance struct {
	AudioEffect
}

func (me *AudioEffectStereoEnhance) BaseClass() string {
	return "AudioEffectStereoEnhance"
}

func NewAudioEffectStereoEnhance() *AudioEffectStereoEnhance {
	str := StringNameFromStr("AudioEffectStereoEnhance") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectStereoEnhance{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectStereoEnhance) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectStereoEnhance) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectStereoEnhance) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectStereoEnhance) SetPanPullout(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectStereoEnhance.fnSetPanPullout), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectStereoEnhance) GetPanPullout() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectStereoEnhance.fnGetPanPullout), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectStereoEnhance) SetTimePullout(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectStereoEnhance.fnSetTimePullout), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectStereoEnhance) GetTimePullout() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectStereoEnhance.fnGetTimePullout), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectStereoEnhance) SetSurround(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectStereoEnhance.fnSetSurround), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectStereoEnhance) GetSurround() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectStereoEnhance.fnGetSurround), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
