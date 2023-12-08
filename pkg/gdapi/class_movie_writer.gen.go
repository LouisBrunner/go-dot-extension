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

func  (me *MovieWriter) XGetAudioMixRate() { // TODO: return value
  // TODO: implement
}

func  (me *MovieWriter) XGetAudioSpeakerMode() { // TODO: return value
  // TODO: implement
}

func  (me *MovieWriter) XHandlesFile(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *MovieWriter) XWriteBegin(movie_size Vector2i, fps int, base_path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *MovieWriter) XWriteFrame(frame_image Image, audio_frame_block unsafe.Pointer, ) { // TODO: return value
  // TODO: implement
}

func  (me *MovieWriter) XWriteEnd() { // TODO: return value
  // TODO: implement
}

func  MovieWriterAddWriter(writer MovieWriter, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
