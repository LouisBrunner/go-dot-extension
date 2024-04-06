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

type ptrsForAudioEffectPannerList struct {
	fnSetPan gdc.MethodBindPtr
	fnGetPan gdc.MethodBindPtr
}

var ptrsForAudioEffectPanner ptrsForAudioEffectPannerList

func initAudioEffectPannerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectPanner")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_pan")
		defer methodName.Destroy()
		ptrsForAudioEffectPanner.fnSetPan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_pan")
		defer methodName.Destroy()
		ptrsForAudioEffectPanner.fnGetPan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type AudioEffectPanner struct {
	AudioEffect
}

func (me *AudioEffectPanner) BaseClass() string {
	return "AudioEffectPanner"
}

func NewAudioEffectPanner() *AudioEffectPanner {
	str := StringNameFromStr("AudioEffectPanner") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectPanner{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectPanner) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectPanner) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectPanner) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectPanner) SetPan(cpanume float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cpanume)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPanner.fnSetPan), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectPanner) GetPan() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPanner.fnGetPan), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
