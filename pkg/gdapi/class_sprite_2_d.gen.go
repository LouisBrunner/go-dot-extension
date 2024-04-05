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

type ptrsForSprite2DList struct {
  fnSetTexture gdc.MethodBindPtr
  fnGetTexture gdc.MethodBindPtr
  fnSetCentered gdc.MethodBindPtr
  fnIsCentered gdc.MethodBindPtr
  fnSetOffset gdc.MethodBindPtr
  fnGetOffset gdc.MethodBindPtr
  fnSetFlipH gdc.MethodBindPtr
  fnIsFlippedH gdc.MethodBindPtr
  fnSetFlipV gdc.MethodBindPtr
  fnIsFlippedV gdc.MethodBindPtr
  fnSetRegionEnabled gdc.MethodBindPtr
  fnIsRegionEnabled gdc.MethodBindPtr
  fnIsPixelOpaque gdc.MethodBindPtr
  fnSetRegionRect gdc.MethodBindPtr
  fnGetRegionRect gdc.MethodBindPtr
  fnSetRegionFilterClipEnabled gdc.MethodBindPtr
  fnIsRegionFilterClipEnabled gdc.MethodBindPtr
  fnSetFrame gdc.MethodBindPtr
  fnGetFrame gdc.MethodBindPtr
  fnSetFrameCoords gdc.MethodBindPtr
  fnGetFrameCoords gdc.MethodBindPtr
  fnSetVframes gdc.MethodBindPtr
  fnGetVframes gdc.MethodBindPtr
  fnSetHframes gdc.MethodBindPtr
  fnGetHframes gdc.MethodBindPtr
  fnGetRect gdc.MethodBindPtr
}

var ptrsForSprite2D ptrsForSprite2DList

func initSprite2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Sprite2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_texture")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_texture")
    defer methodName.Destroy()
    ptrsForSprite2D.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
  }
  {
    methodName := StringNameFromStr("set_centered")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_centered")
    defer methodName.Destroy()
    ptrsForSprite2D.fnIsCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_offset")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
  }
  {
    methodName := StringNameFromStr("get_offset")
    defer methodName.Destroy()
    ptrsForSprite2D.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("set_flip_h")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetFlipH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_flipped_h")
    defer methodName.Destroy()
    ptrsForSprite2D.fnIsFlippedH = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_flip_v")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetFlipV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_flipped_v")
    defer methodName.Destroy()
    ptrsForSprite2D.fnIsFlippedV = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_region_enabled")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetRegionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_region_enabled")
    defer methodName.Destroy()
    ptrsForSprite2D.fnIsRegionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("is_pixel_opaque")
    defer methodName.Destroy()
    ptrsForSprite2D.fnIsPixelOpaque = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 556197845))
  }
  {
    methodName := StringNameFromStr("set_region_rect")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetRegionRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
  }
  {
    methodName := StringNameFromStr("get_region_rect")
    defer methodName.Destroy()
    ptrsForSprite2D.fnGetRegionRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
  }
  {
    methodName := StringNameFromStr("set_region_filter_clip_enabled")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetRegionFilterClipEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_region_filter_clip_enabled")
    defer methodName.Destroy()
    ptrsForSprite2D.fnIsRegionFilterClipEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_frame")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_frame")
    defer methodName.Destroy()
    ptrsForSprite2D.fnGetFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_frame_coords")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetFrameCoords = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
  {
    methodName := StringNameFromStr("get_frame_coords")
    defer methodName.Destroy()
    ptrsForSprite2D.fnGetFrameCoords = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
  }
  {
    methodName := StringNameFromStr("set_vframes")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetVframes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_vframes")
    defer methodName.Destroy()
    ptrsForSprite2D.fnGetVframes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_hframes")
    defer methodName.Destroy()
    ptrsForSprite2D.fnSetHframes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_hframes")
    defer methodName.Destroy()
    ptrsForSprite2D.fnGetHframes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_rect")
    defer methodName.Destroy()
    ptrsForSprite2D.fnGetRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
  }
}

type Sprite2D struct {
  Node2D
}

func (me *Sprite2D) BaseClass() string {
  return "Sprite2D"
}

func NewSprite2D() *Sprite2D {
  str := StringNameFromStr("Sprite2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Sprite2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Sprite2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Sprite2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Sprite2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Sprite2D) SetTexture(texture Texture2D, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) GetTexture() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Sprite2D) SetCentered(centered bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&centered) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetCentered), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) IsCentered() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnIsCentered), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite2D) SetOffset(offset Vector2, )  {
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) GetOffset() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Sprite2D) SetFlipH(flip_h bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_h) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetFlipH), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) IsFlippedH() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnIsFlippedH), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite2D) SetFlipV(flip_v bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_v) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetFlipV), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) IsFlippedV() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnIsFlippedV), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite2D) SetRegionEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetRegionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) IsRegionEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnIsRegionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite2D) IsPixelOpaque(pos Vector2, ) bool {
  cargs := []gdc.ConstTypePtr{pos.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnIsPixelOpaque), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite2D) SetRegionRect(rect Rect2, )  {
  cargs := []gdc.ConstTypePtr{rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetRegionRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) GetRegionRect() Rect2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnGetRegionRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Sprite2D) SetRegionFilterClipEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetRegionFilterClipEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) IsRegionFilterClipEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnIsRegionFilterClipEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite2D) SetFrame(frame int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&frame) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetFrame), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) GetFrame() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnGetFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite2D) SetFrameCoords(coords Vector2i, )  {
  cargs := []gdc.ConstTypePtr{coords.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetFrameCoords), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) GetFrameCoords() Vector2i {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnGetFrameCoords), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Sprite2D) SetVframes(vframes int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vframes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetVframes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) GetVframes() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnGetVframes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite2D) SetHframes(hframes int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hframes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnSetHframes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Sprite2D) GetHframes() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnGetHframes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Sprite2D) GetRect() Rect2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSprite2D.fnGetRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type Sprite2DFrameChangedSignalFn func()

func (me *Sprite2D) ConnectFrameChanged(subs SignalSubscribers, fn Sprite2DFrameChangedSignalFn) {
  sig := StringNameFromStr("frame_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Sprite2D) DisconnectFrameChanged(subs SignalSubscribers, fn Sprite2DFrameChangedSignalFn) {
  sig := StringNameFromStr("frame_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Sprite2DTextureChangedSignalFn func()

func (me *Sprite2D) ConnectTextureChanged(subs SignalSubscribers, fn Sprite2DTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Sprite2D) DisconnectTextureChanged(subs SignalSubscribers, fn Sprite2DTextureChangedSignalFn) {
  sig := StringNameFromStr("texture_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
