// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports

  "unsafe"





  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MovieWriter struct {
  obj gdc.ObjectPtr
}

func (me *MovieWriter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MovieWriter) BaseClass() string {
  return "MovieWriter"
}



// Enums

func (me *MovieWriter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MovieWriter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *MovieWriter) XGetAudioMixRate()  {
  panic("TODO: implement")
}

func  (me *MovieWriter) XGetAudioSpeakerMode()  {
  panic("TODO: implement")
}

func  (me *MovieWriter) XHandlesFile(path String, )  {
  panic("TODO: implement")
}

func  (me *MovieWriter) XWriteBegin(movie_size Vector2i, fps int, base_path String, )  {
  panic("TODO: implement")
}

func  (me *MovieWriter) XWriteFrame(frame_image Image, audio_frame_block unsafe.Pointer, )  {
  panic("TODO: implement")
}

func  (me *MovieWriter) XWriteEnd()  {
  panic("TODO: implement")
}

func  MovieWriterAddWriter(writer MovieWriter, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
