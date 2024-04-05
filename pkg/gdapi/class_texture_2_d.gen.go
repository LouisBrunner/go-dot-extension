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

type ptrsForTexture2DList struct {
  fnXGetWidth gdc.MethodBindPtr
  fnXGetHeight gdc.MethodBindPtr
  fnXIsPixelOpaque gdc.MethodBindPtr
  fnXHasAlpha gdc.MethodBindPtr
  fnXDraw gdc.MethodBindPtr
  fnXDrawRect gdc.MethodBindPtr
  fnXDrawRectRegion gdc.MethodBindPtr
  fnGetWidth gdc.MethodBindPtr
  fnGetHeight gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnHasAlpha gdc.MethodBindPtr
  fnDraw gdc.MethodBindPtr
  fnDrawRect gdc.MethodBindPtr
  fnDrawRectRegion gdc.MethodBindPtr
  fnGetImage gdc.MethodBindPtr
  fnCreatePlaceholder gdc.MethodBindPtr
}

var ptrsForTexture2D ptrsForTexture2DList

func initTexture2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Texture2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_width")
    defer methodName.Destroy()
    ptrsForTexture2D.fnGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_height")
    defer methodName.Destroy()
    ptrsForTexture2D.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForTexture2D.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
  }
  {
    methodName := StringNameFromStr("has_alpha")
    defer methodName.Destroy()
    ptrsForTexture2D.fnHasAlpha = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("draw")
    defer methodName.Destroy()
    ptrsForTexture2D.fnDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2729649137))
  }
  {
    methodName := StringNameFromStr("draw_rect")
    defer methodName.Destroy()
    ptrsForTexture2D.fnDrawRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3499451691))
  }
  {
    methodName := StringNameFromStr("draw_rect_region")
    defer methodName.Destroy()
    ptrsForTexture2D.fnDrawRectRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2963678660))
  }
  {
    methodName := StringNameFromStr("get_image")
    defer methodName.Destroy()
    ptrsForTexture2D.fnGetImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4190603485))
  }
  {
    methodName := StringNameFromStr("create_placeholder")
    defer methodName.Destroy()
    ptrsForTexture2D.fnCreatePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121922552))
  }
}

type Texture2D struct {
  Texture
}

func (me *Texture2D) BaseClass() string {
  return "Texture2D"
}

func NewTexture2D() *Texture2D {
  str := StringNameFromStr("Texture2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Texture2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Texture2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Texture2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Texture2D) GetWidth() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2D.fnGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture2D) GetHeight() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2D.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture2D) GetSize() Vector2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2D.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Texture2D) HasAlpha() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2D.fnHasAlpha), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture2D) Draw(canvas_item RID, position Vector2, modulate Color, transpose bool, )  {
  cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), position.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2D.fnDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Texture2D) DrawRect(canvas_item RID, rect Rect2, tile bool, modulate Color, transpose bool, )  {
  cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), rect.AsCTypePtr(), gdc.ConstTypePtr(&tile) , modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2D.fnDrawRect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Texture2D) DrawRectRegion(canvas_item RID, rect Rect2, src_rect Rect2, modulate Color, transpose bool, clip_uv bool, )  {
  cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), rect.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , gdc.ConstTypePtr(&clip_uv) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2D.fnDrawRectRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Texture2D) GetImage() Image {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2D.fnGetImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Texture2D) CreatePlaceholder() Resource {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTexture2D.fnCreatePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
