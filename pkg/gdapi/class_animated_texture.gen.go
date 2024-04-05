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

type ptrsForAnimatedTextureList struct {
  fnSetFrames gdc.MethodBindPtr
  fnGetFrames gdc.MethodBindPtr
  fnSetCurrentFrame gdc.MethodBindPtr
  fnGetCurrentFrame gdc.MethodBindPtr
  fnSetPause gdc.MethodBindPtr
  fnGetPause gdc.MethodBindPtr
  fnSetOneShot gdc.MethodBindPtr
  fnGetOneShot gdc.MethodBindPtr
  fnSetSpeedScale gdc.MethodBindPtr
  fnGetSpeedScale gdc.MethodBindPtr
  fnSetFrameTexture gdc.MethodBindPtr
  fnGetFrameTexture gdc.MethodBindPtr
  fnSetFrameDuration gdc.MethodBindPtr
  fnGetFrameDuration gdc.MethodBindPtr
}

var ptrsForAnimatedTexture ptrsForAnimatedTextureList

func initAnimatedTexturePtrs(iface gdc.Interface) {

  className := StringNameFromStr("AnimatedTexture")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_frames")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnSetFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_frames")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnGetFrames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_current_frame")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnSetCurrentFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_current_frame")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnGetCurrentFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_pause")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnSetPause = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_pause")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnGetPause = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_one_shot")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnSetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_one_shot")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnGetOneShot = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_speed_scale")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnSetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_speed_scale")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnGetSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_frame_texture")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnSetFrameTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
  }
  {
    methodName := StringNameFromStr("get_frame_texture")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnGetFrameTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
  }
  {
    methodName := StringNameFromStr("set_frame_duration")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnSetFrameDuration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("get_frame_duration")
    defer methodName.Destroy()
    ptrsForAnimatedTexture.fnGetFrameDuration = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
}

type AnimatedTexture struct {
  Texture2D
}

func (me *AnimatedTexture) BaseClass() string {
  return "AnimatedTexture"
}

func NewAnimatedTexture() *AnimatedTexture {
  str := StringNameFromStr("AnimatedTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimatedTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  AnimatedTextureMaxFrames = "256" // TODO: construct correctly
)

// Enums

func (me *AnimatedTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimatedTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimatedTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimatedTexture) SetFrames(frames int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frames) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnSetFrames), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedTexture) GetFrames() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnGetFrames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedTexture) SetCurrentFrame(frame int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnSetCurrentFrame), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedTexture) GetCurrentFrame() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnGetCurrentFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedTexture) SetPause(pause bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pause) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnSetPause), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedTexture) GetPause() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnGetPause), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedTexture) SetOneShot(one_shot bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&one_shot) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnSetOneShot), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedTexture) GetOneShot() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnGetOneShot), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedTexture) SetSpeedScale(scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnSetSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedTexture) GetSpeedScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnGetSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimatedTexture) SetFrameTexture(frame int64, texture Texture2D, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame) , texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnSetFrameTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedTexture) GetFrameTexture(frame int64, ) Texture2D {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&frame)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnGetFrameTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimatedTexture) SetFrameDuration(frame int64, duration float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame) , gdc.ConstTypePtr(&duration) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnSetFrameDuration), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimatedTexture) GetFrameDuration(frame int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&frame)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimatedTexture.fnGetFrameDuration), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
