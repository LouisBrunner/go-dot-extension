// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Texture2D struct {
  obj gdc.ObjectPtr
}

func (me *Texture2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture2D) BaseClass() string {
  return "Texture2D"
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

func  (me *Texture2D) GetWidth() int {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Texture2D) GetHeight() int {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Texture2D) GetSize() Vector2 {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Texture2D) HasAlpha() bool {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_alpha")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Texture2D) Draw(canvas_item RID, position Vector2, modulate Color, transpose bool, )  {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1115460088) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&transpose), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Texture2D) DrawRect(canvas_item RID, rect Rect2, tile bool, modulate Color, transpose bool, )  {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 575156982) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), gdc.ConstTypePtr(&tile), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&transpose), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Texture2D) DrawRectRegion(canvas_item RID, rect Rect2, src_rect Rect2, modulate Color, transpose bool, clip_uv bool, )  {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_rect_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1066564656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), gdc.ConstTypePtr(src_rect.AsCTypePtr()), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&transpose), gdc.ConstTypePtr(&clip_uv), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Texture2D) GetImage() Image {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4190603485) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Texture2D) CreatePlaceholder() Resource {
  classNameV := StringNameFromStr("Texture2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_placeholder")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121922552) // FIXME: should cache?
  var ret Resource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties