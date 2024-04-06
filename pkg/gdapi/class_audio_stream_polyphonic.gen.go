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

type ptrsForAudioStreamPolyphonicList struct {
	fnSetPolyphony gdc.MethodBindPtr
	fnGetPolyphony gdc.MethodBindPtr
}

var ptrsForAudioStreamPolyphonic ptrsForAudioStreamPolyphonicList

func initAudioStreamPolyphonicPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamPolyphonic")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_polyphony")
		defer methodName.Destroy()
		ptrsForAudioStreamPolyphonic.fnSetPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_polyphony")
		defer methodName.Destroy()
		ptrsForAudioStreamPolyphonic.fnGetPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
}

type AudioStreamPolyphonic struct {
	AudioStream
}

func (me *AudioStreamPolyphonic) BaseClass() string {
	return "AudioStreamPolyphonic"
}

func NewAudioStreamPolyphonic() *AudioStreamPolyphonic {
	str := StringNameFromStr("AudioStreamPolyphonic") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamPolyphonic{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioStreamPolyphonic) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamPolyphonic) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPolyphonic) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamPolyphonic) SetPolyphony(voices int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voices)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPolyphonic.fnSetPolyphony), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamPolyphonic) GetPolyphony() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPolyphonic.fnGetPolyphony), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
