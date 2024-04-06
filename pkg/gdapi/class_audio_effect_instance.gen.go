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

type ptrsForAudioEffectInstanceList struct {
	fnXProcess        gdc.MethodBindPtr
	fnXProcessSilence gdc.MethodBindPtr
}

var ptrsForAudioEffectInstance ptrsForAudioEffectInstanceList

func initAudioEffectInstancePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectInstance")
	defer className.Destroy()

}

type AudioEffectInstance struct {
	RefCounted
}

func (me *AudioEffectInstance) BaseClass() string {
	return "AudioEffectInstance"
}

func NewAudioEffectInstance() *AudioEffectInstance {
	str := StringNameFromStr("AudioEffectInstance") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectInstance{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectInstance) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectInstance) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectInstance) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
