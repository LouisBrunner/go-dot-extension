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

type ptrsForAudioEffectRecordList struct {
	fnSetRecordingActive gdc.MethodBindPtr
	fnIsRecordingActive  gdc.MethodBindPtr
	fnSetFormat          gdc.MethodBindPtr
	fnGetFormat          gdc.MethodBindPtr
	fnGetRecording       gdc.MethodBindPtr
}

var ptrsForAudioEffectRecord ptrsForAudioEffectRecordList

func initAudioEffectRecordPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectRecord")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_recording_active")
		defer methodName.Destroy()
		ptrsForAudioEffectRecord.fnSetRecordingActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_recording_active")
		defer methodName.Destroy()
		ptrsForAudioEffectRecord.fnIsRecordingActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_format")
		defer methodName.Destroy()
		ptrsForAudioEffectRecord.fnSetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 60648488))
	}
	{
		methodName := StringNameFromStr("get_format")
		defer methodName.Destroy()
		ptrsForAudioEffectRecord.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3151724922))
	}
	{
		methodName := StringNameFromStr("get_recording")
		defer methodName.Destroy()
		ptrsForAudioEffectRecord.fnGetRecording = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2964110865))
	}
}

type AudioEffectRecord struct {
	AudioEffect
}

func (me *AudioEffectRecord) BaseClass() string {
	return "AudioEffectRecord"
}

func NewAudioEffectRecord() *AudioEffectRecord {
	str := StringNameFromStr("AudioEffectRecord") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectRecord{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectRecord) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectRecord) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectRecord) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectRecord) SetRecordingActive(record bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&record)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectRecord.fnSetRecordingActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectRecord) IsRecordingActive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectRecord.fnIsRecordingActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectRecord) SetFormat(format AudioStreamWAVFormat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectRecord.fnSetFormat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectRecord) GetFormat() AudioStreamWAVFormat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioStreamWAVFormat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectRecord.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AudioEffectRecord) GetRecording() AudioStreamWAV {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAudioStreamWAV()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectRecord.fnGetRecording), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
