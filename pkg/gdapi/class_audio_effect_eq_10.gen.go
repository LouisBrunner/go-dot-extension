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

type ptrsForAudioEffectEQ10List struct {
}

var ptrsForAudioEffectEQ10 ptrsForAudioEffectEQ10List

func initAudioEffectEQ10Ptrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectEQ10")
	defer className.Destroy()
}

type AudioEffectEQ10 struct {
	AudioEffectEQ
}

func (me *AudioEffectEQ10) BaseClass() string {
	return "AudioEffectEQ10"
}

func NewAudioEffectEQ10() *AudioEffectEQ10 {
	str := StringNameFromStr("AudioEffectEQ10") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectEQ10{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectEQ10) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectEQ10) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ10) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
