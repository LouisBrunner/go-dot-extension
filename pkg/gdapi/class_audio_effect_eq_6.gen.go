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

type ptrsForAudioEffectEQ6List struct {
}

var ptrsForAudioEffectEQ6 ptrsForAudioEffectEQ6List

func initAudioEffectEQ6Ptrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectEQ6")
	defer className.Destroy()
}

type AudioEffectEQ6 struct {
	AudioEffectEQ
}

func (me *AudioEffectEQ6) BaseClass() string {
	return "AudioEffectEQ6"
}

func NewAudioEffectEQ6() *AudioEffectEQ6 {
	str := StringNameFromStr("AudioEffectEQ6") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectEQ6{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectEQ6) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectEQ6) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ6) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
