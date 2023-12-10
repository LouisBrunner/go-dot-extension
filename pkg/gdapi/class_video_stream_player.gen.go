// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VideoStreamPlayer struct {
  obj gdc.ObjectPtr
}

func (me *VideoStreamPlayer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VideoStreamPlayer) BaseClass() string {
  return "VideoStreamPlayer"
}



// Enums

func (me *VideoStreamPlayer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VideoStreamPlayer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VideoStreamPlayer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VideoStreamPlayer) SetStream(stream VideoStream, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2317102564) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(stream.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) GetStream() VideoStream {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 438621487) // FIXME: should cache?
  var ret VideoStream
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) Play()  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) Stop()  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) IsPlaying() bool {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) SetPaused(paused bool, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paused), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) IsPaused() bool {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) SetVolume(volume float32, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volume")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&volume), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) GetVolume() float32 {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volume")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) SetVolumeDb(db float32, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&db), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) GetVolumeDb() float32 {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) SetAudioTrack(track int, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_audio_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) GetAudioTrack() int {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_audio_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) GetStreamName() String {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) SetStreamPosition(position float32, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) GetStreamPosition() float32 {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) SetAutoplay(enabled bool, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) HasAutoplay() bool {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) SetExpand(enable bool, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_expand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) HasExpand() bool {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_expand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) SetBufferingMsec(msec int, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffering_msec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) GetBufferingMsec() int {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffering_msec")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) SetBus(bus StringName, )  {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bus.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VideoStreamPlayer) GetBus() StringName {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VideoStreamPlayer) GetVideoTexture() Texture2D {
  classNameV := StringNameFromStr("VideoStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_video_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VideoStreamPlayer) GetPropAudioTrack() int {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropAudioTrack(value int) {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) GetPropStream() VideoStream {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropStream(value VideoStream) {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) GetPropVolumeDb() float32 {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropVolumeDb(value float32) {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) GetPropVolume() float32 {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropVolume(value float32) {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) GetPropAutoplay() bool {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropAutoplay(value bool) {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) GetPropPaused() bool {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropPaused(value bool) {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) GetPropExpand() bool {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropExpand(value bool) {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) GetPropBufferingMsec() int {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropBufferingMsec(value int) {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) GetPropStreamPosition() float32 {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropStreamPosition(value float32) {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) GetPropBus() StringName {
  panic("TODO: implement")
}

func (me *VideoStreamPlayer) SetPropBus(value StringName) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API