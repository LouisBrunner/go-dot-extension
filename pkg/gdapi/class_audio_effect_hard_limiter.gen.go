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

type ptrsForAudioEffectHardLimiterList struct {
	fnSetCeilingDb gdc.MethodBindPtr
	fnGetCeilingDb gdc.MethodBindPtr
	fnSetPreGainDb gdc.MethodBindPtr
	fnGetPreGainDb gdc.MethodBindPtr
	fnSetRelease   gdc.MethodBindPtr
	fnGetRelease   gdc.MethodBindPtr
}

var ptrsForAudioEffectHardLimiter ptrsForAudioEffectHardLimiterList

func initAudioEffectHardLimiterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectHardLimiter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_ceiling_db")
		defer methodName.Destroy()
		ptrsForAudioEffectHardLimiter.fnSetCeilingDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_ceiling_db")
		defer methodName.Destroy()
		ptrsForAudioEffectHardLimiter.fnGetCeilingDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_pre_gain_db")
		defer methodName.Destroy()
		ptrsForAudioEffectHardLimiter.fnSetPreGainDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pre_gain_db")
		defer methodName.Destroy()
		ptrsForAudioEffectHardLimiter.fnGetPreGainDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_release")
		defer methodName.Destroy()
		ptrsForAudioEffectHardLimiter.fnSetRelease = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_release")
		defer methodName.Destroy()
		ptrsForAudioEffectHardLimiter.fnGetRelease = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type AudioEffectHardLimiter struct {
	AudioEffect
}

func (me *AudioEffectHardLimiter) BaseClass() string {
	return "AudioEffectHardLimiter"
}

func NewAudioEffectHardLimiter() *AudioEffectHardLimiter {
	str := StringNameFromStr("AudioEffectHardLimiter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectHardLimiter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectHardLimiter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectHardLimiter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectHardLimiter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectHardLimiter) SetCeilingDb(ceiling float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ceiling)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectHardLimiter.fnSetCeilingDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectHardLimiter) GetCeilingDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectHardLimiter.fnGetCeilingDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectHardLimiter) SetPreGainDb(p_pre_gain float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_pre_gain)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectHardLimiter.fnSetPreGainDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectHardLimiter) GetPreGainDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectHardLimiter.fnGetPreGainDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectHardLimiter) SetRelease(p_release float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_release)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectHardLimiter.fnSetRelease), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectHardLimiter) GetRelease() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectHardLimiter.fnGetRelease), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
