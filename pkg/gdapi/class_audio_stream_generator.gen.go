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

type ptrsForAudioStreamGeneratorList struct {
	fnSetMixRate      gdc.MethodBindPtr
	fnGetMixRate      gdc.MethodBindPtr
	fnSetBufferLength gdc.MethodBindPtr
	fnGetBufferLength gdc.MethodBindPtr
}

var ptrsForAudioStreamGenerator ptrsForAudioStreamGeneratorList

func initAudioStreamGeneratorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamGenerator")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_mix_rate")
		defer methodName.Destroy()
		ptrsForAudioStreamGenerator.fnSetMixRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_mix_rate")
		defer methodName.Destroy()
		ptrsForAudioStreamGenerator.fnGetMixRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_buffer_length")
		defer methodName.Destroy()
		ptrsForAudioStreamGenerator.fnSetBufferLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_buffer_length")
		defer methodName.Destroy()
		ptrsForAudioStreamGenerator.fnGetBufferLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type AudioStreamGenerator struct {
	AudioStream
}

func (me *AudioStreamGenerator) BaseClass() string {
	return "AudioStreamGenerator"
}

func NewAudioStreamGenerator() *AudioStreamGenerator {
	str := StringNameFromStr("AudioStreamGenerator") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamGenerator{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioStreamGenerator) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamGenerator) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamGenerator) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamGenerator) SetMixRate(hz float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGenerator.fnSetMixRate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamGenerator) GetMixRate() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGenerator.fnGetMixRate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamGenerator) SetBufferLength(seconds float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGenerator.fnSetBufferLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamGenerator) GetBufferLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamGenerator.fnGetBufferLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
