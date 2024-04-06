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

type ptrsForAudioStreamMP3List struct {
	fnSetData       gdc.MethodBindPtr
	fnGetData       gdc.MethodBindPtr
	fnSetLoop       gdc.MethodBindPtr
	fnHasLoop       gdc.MethodBindPtr
	fnSetLoopOffset gdc.MethodBindPtr
	fnGetLoopOffset gdc.MethodBindPtr
	fnSetBpm        gdc.MethodBindPtr
	fnGetBpm        gdc.MethodBindPtr
	fnSetBeatCount  gdc.MethodBindPtr
	fnGetBeatCount  gdc.MethodBindPtr
	fnSetBarBeats   gdc.MethodBindPtr
	fnGetBarBeats   gdc.MethodBindPtr
}

var ptrsForAudioStreamMP3 ptrsForAudioStreamMP3List

func initAudioStreamMP3Ptrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamMP3")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_data")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnSetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2971499966))
	}
	{
		methodName := StringNameFromStr("get_data")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2362200018))
	}
	{
		methodName := StringNameFromStr("set_loop")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnSetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_loop")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnHasLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_loop_offset")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnSetLoopOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_loop_offset")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnGetLoopOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_bpm")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnSetBpm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bpm")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnGetBpm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_beat_count")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnSetBeatCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_beat_count")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnGetBeatCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_bar_beats")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnSetBarBeats = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_bar_beats")
		defer methodName.Destroy()
		ptrsForAudioStreamMP3.fnGetBarBeats = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type AudioStreamMP3 struct {
	AudioStream
}

func (me *AudioStreamMP3) BaseClass() string {
	return "AudioStreamMP3"
}

func NewAudioStreamMP3() *AudioStreamMP3 {
	str := StringNameFromStr("AudioStreamMP3") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamMP3{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioStreamMP3) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamMP3) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamMP3) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamMP3) SetData(data PackedByteArray) {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnSetData), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamMP3) GetData() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AudioStreamMP3) SetLoop(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnSetLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamMP3) HasLoop() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnHasLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamMP3) SetLoopOffset(seconds float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnSetLoopOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamMP3) GetLoopOffset() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnGetLoopOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamMP3) SetBpm(bpm float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bpm)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnSetBpm), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamMP3) GetBpm() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnGetBpm), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamMP3) SetBeatCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnSetBeatCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamMP3) GetBeatCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnGetBeatCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStreamMP3) SetBarBeats(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnSetBarBeats), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamMP3) GetBarBeats() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamMP3.fnGetBarBeats), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
