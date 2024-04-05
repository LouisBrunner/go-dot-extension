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

type ptrsForSprite3DList struct {
  fnSetTexture gdc.MethodBindPtr
  fnGetTexture gdc.MethodBindPtr
  fnSetRegionEnabled gdc.MethodBindPtr
  fnIsRegionEnabled gdc.MethodBindPtr
  fnSetRegionRect gdc.MethodBindPtr
  fnGetRegionRect gdc.MethodBindPtr
  fnSetFrame gdc.MethodBindPtr
  fnGetFrame gdc.MethodBindPtr
  fnSetFrameCoords gdc.MethodBindPtr
  fnGetFrameCoords gdc.MethodBindPtr
  fnSetVframes gdc.MethodBindPtr
  fnGetVframes gdc.MethodBindPtr
  fnSetHframes gdc.MethodBindPtr
  fnGetHframes gdc.MethodBindPtr
}

var ptrsForSprite3D ptrsForSprite3DList

func initSprite3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Sprite3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_texture")
    defer methodName.Destroy()
    ptrsForSprite3D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_texture")
    defer methodName.Destroy()
    ptrsForSprite3D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("set_region_enabled")
    defer methodName.Destroy()
    ptrsForSprite3D.fnSetRegionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_region_enabled")
    defer methodName.Destroy()
    ptrsForSprite3D.fnIsRegionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_region_rect")
    defer methodName.Destroy()
    ptrsForSprite3D.fnSetRegionRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
  }
  {
    methodName := StringNameFromStr("get_region_rect")
    defer methodName.Destroy()
    ptrsForSprite3D.fnGetRegionRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
  }
  {
    methodName := StringNameFromStr("set_frame")
    defer methodName.Destroy()
    ptrsForSprite3D.fnSetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_frame")
    defer methodName.Destroy()
    ptrsForSprite3D.fnGetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_frame_coords")
    defer methodName.Destroy()
    ptrsForSprite3D.fnSetFrameCoords = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_frame_coords")
    defer methodName.Destroy()
    ptrsForSprite3D.fnGetFrameCoords = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_vframes")
    defer methodName.Destroy()
    ptrsForSprite3D.fnSetVframes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_vframes")
    defer methodName.Destroy()
    ptrsForSprite3D.fnGetVframes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_hframes")
    defer methodName.Destroy()
    ptrsForSprite3D.fnSetHframes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_hframes")
    defer methodName.Destroy()
    ptrsForSprite3D.fnGetHframes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type Sprite3D struct {
  SpriteBase3D
}

func (me *Sprite3D) BaseClass() string {
  return "Sprite3D"
}

func NewSprite3D() *Sprite3D {
  str := StringNameFromStr("Sprite3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Sprite3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Sprite3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Sprite3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Sprite3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Sprite3D) SetTexture(texture Texture2D, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite3D) GetTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Sprite3D) SetRegionEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnSetRegionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite3D) IsRegionEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnIsRegionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite3D) SetRegionRect(rect Rect2, )  {
  cargs := []gdc.ConstTypePtr{rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnSetRegionRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite3D) GetRegionRect() Rect2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnGetRegionRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Sprite3D) SetFrame(frame int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnSetFrame), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite3D) GetFrame() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnGetFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite3D) SetFrameCoords(coords Vector2i, )  {
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnSetFrameCoords), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite3D) GetFrameCoords() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnGetFrameCoords), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Sprite3D) SetVframes(vframes int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vframes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnSetVframes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite3D) GetVframes() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnGetVframes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite3D) SetHframes(hframes int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hframes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnSetHframes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite3D) GetHframes() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite3D.fnGetHframes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type Sprite3DFrameChangedSignalFn func()

func (me *Sprite3D) ConnectFrameChanged(subs SignalSubscribers, fn Sprite3DFrameChangedSignalFn) {
  sig := StringNameFromStr("frame_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Sprite3D) DisconnectFrameChanged(subs SignalSubscribers, fn Sprite3DFrameChangedSignalFn) {
  sig := StringNameFromStr("frame_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Sprite3DTextureChangedSignalFn func()

func (me *Sprite3D) ConnectTextureChanged(subs SignalSubscribers, fn Sprite3DTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Sprite3D) DisconnectTextureChanged(subs SignalSubscribers, fn Sprite3DTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
