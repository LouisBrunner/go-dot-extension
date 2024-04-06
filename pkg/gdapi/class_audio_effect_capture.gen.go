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

type ptrsForAudioEffectCaptureList struct {
	fnCanGetBuffer          gdc.MethodBindPtr
	fnGetBuffer             gdc.MethodBindPtr
	fnClearBuffer           gdc.MethodBindPtr
	fnSetBufferLength       gdc.MethodBindPtr
	fnGetBufferLength       gdc.MethodBindPtr
	fnGetFramesAvailable    gdc.MethodBindPtr
	fnGetDiscardedFrames    gdc.MethodBindPtr
	fnGetBufferLengthFrames gdc.MethodBindPtr
	fnGetPushedFrames       gdc.MethodBindPtr
}

var ptrsForAudioEffectCapture ptrsForAudioEffectCaptureList

func initAudioEffectCapturePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectCapture")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("can_get_buffer")
		defer methodName.Destroy()
		ptrsForAudioEffectCapture.fnCanGetBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_buffer")
		defer methodName.Destroy()
		ptrsForAudioEffectCapture.fnGetBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2649534757))
	}
	{
		methodName := StringNameFromStr("clear_buffer")
		defer methodName.Destroy()
		ptrsForAudioEffectCapture.fnClearBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_buffer_length")
		defer methodName.Destroy()
		ptrsForAudioEffectCapture.fnSetBufferLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_buffer_length")
		defer methodName.Destroy()
		ptrsForAudioEffectCapture.fnGetBufferLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("get_frames_available")
		defer methodName.Destroy()
		ptrsForAudioEffectCapture.fnGetFramesAvailable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_discarded_frames")
		defer methodName.Destroy()
		ptrsForAudioEffectCapture.fnGetDiscardedFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_buffer_length_frames")
		defer methodName.Destroy()
		ptrsForAudioEffectCapture.fnGetBufferLengthFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_pushed_frames")
		defer methodName.Destroy()
		ptrsForAudioEffectCapture.fnGetPushedFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type AudioEffectCapture struct {
	AudioEffect
}

func (me *AudioEffectCapture) BaseClass() string {
	return "AudioEffectCapture"
}

func NewAudioEffectCapture() *AudioEffectCapture {
	str := StringNameFromStr("AudioEffectCapture") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectCapture{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectCapture) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectCapture) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectCapture) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectCapture) CanGetBuffer(frames int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&frames)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCapture.fnCanGetBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCapture) GetBuffer(frames int64) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()
	pinner.Pin(&frames)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCapture.fnGetBuffer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AudioEffectCapture) ClearBuffer() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCapture.fnClearBuffer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectCapture) SetBufferLength(buffer_length_seconds float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_length_seconds)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCapture.fnSetBufferLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectCapture) GetBufferLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCapture.fnGetBufferLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCapture) GetFramesAvailable() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCapture.fnGetFramesAvailable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCapture) GetDiscardedFrames() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCapture.fnGetDiscardedFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCapture) GetBufferLengthFrames() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCapture.fnGetBufferLengthFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectCapture) GetPushedFrames() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectCapture.fnGetPushedFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
