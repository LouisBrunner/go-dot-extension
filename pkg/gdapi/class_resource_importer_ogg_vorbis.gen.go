// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterOggVorbis struct {
  ResourceImporter
}

func (me *ResourceImporterOggVorbis) BaseClass() string {
  return "ResourceImporterOggVorbis"
}

func NewResourceImporterOggVorbis() *ResourceImporterOggVorbis {
  str := StringNameFromStr("ResourceImporterOggVorbis") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterOggVorbis{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterOggVorbis) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterOggVorbis) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterOggVorbis) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  ResourceImporterOggVorbisLoadFromBuffer(buffer PackedByteArray, ) AudioStreamOggVorbis {
  classNameV := StringNameFromStr("ResourceImporterOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 354904730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  ret := NewAudioStreamOggVorbis()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  ResourceImporterOggVorbisLoadFromFile(path String, ) AudioStreamOggVorbis {
  classNameV := StringNameFromStr("ResourceImporterOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797568536) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewAudioStreamOggVorbis()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
