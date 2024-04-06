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

type ptrsForAudioEffectDelayList struct {
	fnSetDry             gdc.MethodBindPtr
	fnGetDry             gdc.MethodBindPtr
	fnSetTap1Active      gdc.MethodBindPtr
	fnIsTap1Active       gdc.MethodBindPtr
	fnSetTap1DelayMs     gdc.MethodBindPtr
	fnGetTap1DelayMs     gdc.MethodBindPtr
	fnSetTap1LevelDb     gdc.MethodBindPtr
	fnGetTap1LevelDb     gdc.MethodBindPtr
	fnSetTap1Pan         gdc.MethodBindPtr
	fnGetTap1Pan         gdc.MethodBindPtr
	fnSetTap2Active      gdc.MethodBindPtr
	fnIsTap2Active       gdc.MethodBindPtr
	fnSetTap2DelayMs     gdc.MethodBindPtr
	fnGetTap2DelayMs     gdc.MethodBindPtr
	fnSetTap2LevelDb     gdc.MethodBindPtr
	fnGetTap2LevelDb     gdc.MethodBindPtr
	fnSetTap2Pan         gdc.MethodBindPtr
	fnGetTap2Pan         gdc.MethodBindPtr
	fnSetFeedbackActive  gdc.MethodBindPtr
	fnIsFeedbackActive   gdc.MethodBindPtr
	fnSetFeedbackDelayMs gdc.MethodBindPtr
	fnGetFeedbackDelayMs gdc.MethodBindPtr
	fnSetFeedbackLevelDb gdc.MethodBindPtr
	fnGetFeedbackLevelDb gdc.MethodBindPtr
	fnSetFeedbackLowpass gdc.MethodBindPtr
	fnGetFeedbackLowpass gdc.MethodBindPtr
}

var ptrsForAudioEffectDelay ptrsForAudioEffectDelayList

func initAudioEffectDelayPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectDelay")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_dry")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetDry = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_dry")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetDry = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("set_tap1_active")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetTap1Active = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_tap1_active")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnIsTap1Active = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tap1_delay_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetTap1DelayMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_tap1_delay_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetTap1DelayMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_tap1_level_db")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetTap1LevelDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_tap1_level_db")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetTap1LevelDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_tap1_pan")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetTap1Pan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_tap1_pan")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetTap1Pan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_tap2_active")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetTap2Active = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_tap2_active")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnIsTap2Active = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tap2_delay_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetTap2DelayMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_tap2_delay_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetTap2DelayMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_tap2_level_db")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetTap2LevelDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_tap2_level_db")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetTap2LevelDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_tap2_pan")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetTap2Pan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_tap2_pan")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetTap2Pan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_feedback_active")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetFeedbackActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_feedback_active")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnIsFeedbackActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_feedback_delay_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetFeedbackDelayMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_feedback_delay_ms")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetFeedbackDelayMs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_feedback_level_db")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetFeedbackLevelDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_feedback_level_db")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetFeedbackLevelDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_feedback_lowpass")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnSetFeedbackLowpass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_feedback_lowpass")
		defer methodName.Destroy()
		ptrsForAudioEffectDelay.fnGetFeedbackLowpass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type AudioEffectDelay struct {
	AudioEffect
}

func (me *AudioEffectDelay) BaseClass() string {
	return "AudioEffectDelay"
}

func NewAudioEffectDelay() *AudioEffectDelay {
	str := StringNameFromStr("AudioEffectDelay") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectDelay{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectDelay) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectDelay) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectDelay) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectDelay) SetDry(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetDry), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetDry() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetDry), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetTap1Active(amount bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetTap1Active), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) IsTap1Active() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnIsTap1Active), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetTap1DelayMs(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetTap1DelayMs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetTap1DelayMs() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetTap1DelayMs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetTap1LevelDb(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetTap1LevelDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetTap1LevelDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetTap1LevelDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetTap1Pan(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetTap1Pan), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetTap1Pan() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetTap1Pan), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetTap2Active(amount bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetTap2Active), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) IsTap2Active() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnIsTap2Active), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetTap2DelayMs(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetTap2DelayMs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetTap2DelayMs() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetTap2DelayMs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetTap2LevelDb(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetTap2LevelDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetTap2LevelDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetTap2LevelDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetTap2Pan(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetTap2Pan), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetTap2Pan() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetTap2Pan), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetFeedbackActive(amount bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetFeedbackActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) IsFeedbackActive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnIsFeedbackActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetFeedbackDelayMs(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetFeedbackDelayMs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetFeedbackDelayMs() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetFeedbackDelayMs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetFeedbackLevelDb(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetFeedbackLevelDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetFeedbackLevelDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetFeedbackLevelDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectDelay) SetFeedbackLowpass(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnSetFeedbackLowpass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectDelay) GetFeedbackLowpass() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectDelay.fnGetFeedbackLowpass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
