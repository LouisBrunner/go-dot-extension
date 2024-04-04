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
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture2D) GetHeight() int64 {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture2D) GetSize() Vector2 {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Texture2D) HasAlpha() bool {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Texture2D) Draw(canvas_item RID, position Vector2, modulate Color, transpose bool, )  {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2729649137) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), position.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Texture2D) DrawRect(canvas_item RID, rect Rect2, tile bool, modulate Color, transpose bool, )  {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3499451691) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), rect.AsCTypePtr(), gdc.ConstTypePtr(&tile) , modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Texture2D) DrawRectRegion(canvas_item RID, rect Rect2, src_rect Rect2, modulate Color, transpose bool, clip_uv bool, )  {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_rect_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2963678660) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), rect.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , gdc.ConstTypePtr(&clip_uv) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Texture2D) GetImage() Image {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4190603485) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImage()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Texture2D) CreatePlaceholder() Resource {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121922552) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
