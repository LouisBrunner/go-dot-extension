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

type SpriteFrames struct {
  Resource
}

func (me *SpriteFrames) BaseClass() string {
  return "SpriteFrames"
}

func NewSpriteFrames() *SpriteFrames {
  str := StringNameFromStr("SpriteFrames") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SpriteFrames{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SpriteFrames) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SpriteFrames) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpriteFrames) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SpriteFrames) AddAnimation(anim StringName, )  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteFrames) HasAnimation(anim StringName, ) bool {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteFrames) RemoveAnimation(anim StringName, )  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteFrames) RenameAnimation(anim StringName, newname StringName, )  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rename_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), newname.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteFrames) GetAnimationNames() PackedStringArray {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_names")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1139954409) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SpriteFrames) SetAnimationSpeed(anim StringName, fps float64, )  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animation_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4135858297) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&fps) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteFrames) GetAnimationSpeed(anim StringName, ) float64 {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_speed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2349060816) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteFrames) SetAnimationLoop(anim StringName, loop bool, )  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animation_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2524380260) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&loop) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteFrames) GetAnimationLoop(anim StringName, ) bool {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteFrames) AddFrame(anim StringName, texture Texture2D, duration float64, at_position int64, )  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1351332740) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), texture.AsCTypePtr(), gdc.ConstTypePtr(&duration) , gdc.ConstTypePtr(&at_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteFrames) SetFrame(anim StringName, idx int64, texture Texture2D, duration float64, )  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 56804795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&idx) , texture.AsCTypePtr(), gdc.ConstTypePtr(&duration) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteFrames) RemoveFrame(anim StringName, idx int64, )  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_frame")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2415702435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteFrames) GetFrameCount(anim StringName, ) int64 {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2458036349) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteFrames) GetFrameTexture(anim StringName, idx int64, ) Texture2D {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2900517879) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SpriteFrames) GetFrameDuration(anim StringName, idx int64, ) float64 {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frame_duration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1129309260) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpriteFrames) Clear(anim StringName, )  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{anim.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpriteFrames) ClearAll()  {
  classNameV := StringNameFromStr("SpriteFrames")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
