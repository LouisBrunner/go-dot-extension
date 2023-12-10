// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CanvasItem struct {
  obj gdc.ObjectPtr
}

func (me *CanvasItem) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasItem) BaseClass() string {
  return "CanvasItem"
}



// Constants

var (
  CanvasItemNotificationTransformChanged = "2000" // TODO: construct correctly
  CanvasItemNotificationLocalTransformChanged = "35" // TODO: construct correctly
  CanvasItemNotificationDraw = "30" // TODO: construct correctly
  CanvasItemNotificationVisibilityChanged = "31" // TODO: construct correctly
  CanvasItemNotificationEnterCanvas = "32" // TODO: construct correctly
  CanvasItemNotificationExitCanvas = "33" // TODO: construct correctly
  CanvasItemNotificationWorld2DChanged = "36" // TODO: construct correctly
)

// Enums

type CanvasItemTextureFilter int
const (
  CanvasItemTextureFilterTextureFilterParentNode CanvasItemTextureFilter = 0
  CanvasItemTextureFilterTextureFilterNearest CanvasItemTextureFilter = 1
  CanvasItemTextureFilterTextureFilterLinear CanvasItemTextureFilter = 2
  CanvasItemTextureFilterTextureFilterNearestWithMipmaps CanvasItemTextureFilter = 3
  CanvasItemTextureFilterTextureFilterLinearWithMipmaps CanvasItemTextureFilter = 4
  CanvasItemTextureFilterTextureFilterNearestWithMipmapsAnisotropic CanvasItemTextureFilter = 5
  CanvasItemTextureFilterTextureFilterLinearWithMipmapsAnisotropic CanvasItemTextureFilter = 6
  CanvasItemTextureFilterTextureFilterMax CanvasItemTextureFilter = 7
)

type CanvasItemTextureRepeat int
const (
  CanvasItemTextureRepeatTextureRepeatParentNode CanvasItemTextureRepeat = 0
  CanvasItemTextureRepeatTextureRepeatDisabled CanvasItemTextureRepeat = 1
  CanvasItemTextureRepeatTextureRepeatEnabled CanvasItemTextureRepeat = 2
  CanvasItemTextureRepeatTextureRepeatMirror CanvasItemTextureRepeat = 3
  CanvasItemTextureRepeatTextureRepeatMax CanvasItemTextureRepeat = 4
)

type CanvasItemClipChildrenMode int
const (
  CanvasItemClipChildrenModeClipChildrenDisabled CanvasItemClipChildrenMode = 0
  CanvasItemClipChildrenModeClipChildrenOnly CanvasItemClipChildrenMode = 1
  CanvasItemClipChildrenModeClipChildrenAndDraw CanvasItemClipChildrenMode = 2
  CanvasItemClipChildrenModeClipChildrenMax CanvasItemClipChildrenMode = 3
)

func (me *CanvasItem) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CanvasItem) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CanvasItem) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CanvasItem) GetCanvasItem() RID {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetVisible(visible bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) IsVisible() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) IsVisibleInTree() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible_in_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) Show()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("show")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) Hide()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) QueueRedraw()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_redraw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) MoveToFront()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_to_front")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) SetAsTopLevel(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_top_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) IsSetAsTopLevel() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_set_as_top_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetLightMask(light_mask int, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetLightMask() int {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetModulate(modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetModulate() Color {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetSelfModulate(self_modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_self_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(self_modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetSelfModulate() Color {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_self_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetZIndex(z_index int, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&z_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetZIndex() int {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetZAsRelative(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_z_as_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) IsZRelative() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_z_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetYSortEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_y_sort_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) IsYSortEnabled() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_y_sort_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetDrawBehindParent(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_behind_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) IsDrawBehindParentEnabled() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draw_behind_parent_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) DrawLine(from Vector2, to Vector2, color Color, width float32, antialiased bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2516941890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&antialiased), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawDashedLine(from Vector2, to Vector2, color Color, width float32, dash float32, aligned bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_dashed_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2175215884) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(from.AsCTypePtr()), gdc.ConstTypePtr(to.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&dash), gdc.ConstTypePtr(&aligned), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawPolyline(points PackedVector2Array, color Color, width float32, antialiased bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_polyline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4175878946) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&antialiased), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawPolylineColors(points PackedVector2Array, colors PackedColorArray, width float32, antialiased bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_polyline_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2239164197) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), gdc.ConstTypePtr(colors.AsCTypePtr()), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&antialiased), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawArc(center Vector2, radius float32, start_angle float32, end_angle float32, point_count int, color Color, width float32, antialiased bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_arc")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3486841771) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(center.AsCTypePtr()), gdc.ConstTypePtr(&radius), gdc.ConstTypePtr(&start_angle), gdc.ConstTypePtr(&end_angle), gdc.ConstTypePtr(&point_count), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&antialiased), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawMultiline(points PackedVector2Array, color Color, width float32, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4230657331) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawMultilineColors(points PackedVector2Array, colors PackedColorArray, width float32, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 235933050) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), gdc.ConstTypePtr(colors.AsCTypePtr()), gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawRect(rect Rect2, color Color, filled bool, width float32, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 84391229) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&filled), gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawCircle(position Vector2, radius float32, color Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_circle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3063020269) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&radius), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawTexture(texture Texture2D, position Vector2, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1695860435) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawTextureRect(texture Texture2D, rect Rect2, tile bool, modulate Color, transpose bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_texture_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3204081724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), gdc.ConstTypePtr(&tile), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&transpose), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, transpose bool, clip_uv bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_texture_rect_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3196597532) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), gdc.ConstTypePtr(src_rect.AsCTypePtr()), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&transpose), gdc.ConstTypePtr(&clip_uv), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawMsdfTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, outline float32, pixel_range float32, scale float32, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_msdf_texture_rect_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2672026175) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), gdc.ConstTypePtr(src_rect.AsCTypePtr()), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&outline), gdc.ConstTypePtr(&pixel_range), gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawLcdTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_lcd_texture_rect_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 169610548) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), gdc.ConstTypePtr(src_rect.AsCTypePtr()), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawStyleBox(style_box StyleBox, rect Rect2, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_style_box")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 388176283) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(style_box.AsCTypePtr()), gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawPrimitive(points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_primitive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2248678295) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), gdc.ConstTypePtr(colors.AsCTypePtr()), gdc.ConstTypePtr(uvs.AsCTypePtr()), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawPolygon(points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2683625537) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), gdc.ConstTypePtr(colors.AsCTypePtr()), gdc.ConstTypePtr(uvs.AsCTypePtr()), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawColoredPolygon(points PackedVector2Array, color Color, uvs PackedVector2Array, texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_colored_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1659099617) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(uvs.AsCTypePtr()), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawString(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2552080639) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawMultilineString(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4002645436) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawStringOutline(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_string_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 850005221) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawMultilineStringOutline(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, size int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline_string_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3717870722) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(modulate.AsCTypePtr()), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawChar(font Font, pos Vector2, char String, font_size int, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_char")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2329089032) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(char.AsCTypePtr()), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawCharOutline(font Font, pos Vector2, char String, font_size int, size int, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_char_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 419453826) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(char.AsCTypePtr()), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawMesh(mesh Mesh, texture Texture2D, transform Transform2D, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1634855856) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawMultimesh(multimesh MultiMesh, texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multimesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 937992368) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(multimesh.AsCTypePtr()), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawSetTransform(position Vector2, rotation float32, scale Vector2, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3283884939) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&rotation), gdc.ConstTypePtr(scale.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawSetTransformMatrix(xform Transform2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_set_transform_matrix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(xform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawAnimationSlice(animation_length float32, slice_begin float32, slice_end float32, offset float32, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_animation_slice")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2295343543) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&animation_length), gdc.ConstTypePtr(&slice_begin), gdc.ConstTypePtr(&slice_end), gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) DrawEndAnimation()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_end_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetGlobalTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetGlobalTransformWithCanvas() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_transform_with_canvas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetViewportTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_viewport_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetViewportRect() Rect2 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_viewport_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetCanvasTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetScreenTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetLocalMousePosition() Vector2 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_mouse_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetGlobalMousePosition() Vector2 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_mouse_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetCanvas() RID {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) GetWorld2D() World2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339128592) // FIXME: should cache?
  var ret World2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetMaterial() Material {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  var ret Material
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetUseParentMaterial(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_parent_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetUseParentMaterial() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_parent_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetNotifyLocalTransform(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_notify_local_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) IsLocalTransformNotificationEnabled() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_local_transform_notification_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetNotifyTransform(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_notify_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) IsTransformNotificationEnabled() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_transform_notification_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) ForceUpdateTransform()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) MakeCanvasPositionLocal(screen_point Vector2, ) Vector2 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_canvas_position_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2656412154) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(screen_point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) MakeInputLocal(event InputEvent, ) InputEvent {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_input_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 811130057) // FIXME: should cache?
  var ret InputEvent
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(event.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetVisibilityLayer(layer int, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetVisibilityLayer() int {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetVisibilityLayerBit(layer int, enabled bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_layer_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetVisibilityLayerBit(layer int, ) bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_layer_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetTextureFilter(mode CanvasItemTextureFilter, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1037999706) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetTextureFilter() CanvasItemTextureFilter {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121960042) // FIXME: should cache?
  var ret CanvasItemTextureFilter
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetTextureRepeat(mode CanvasItemTextureRepeat, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1716472974) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetTextureRepeat() CanvasItemTextureRepeat {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2667158319) // FIXME: should cache?
  var ret CanvasItemTextureRepeat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CanvasItem) SetClipChildrenMode(mode CanvasItemClipChildrenMode, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clip_children_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1319393776) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CanvasItem) GetClipChildrenMode() CanvasItemClipChildrenMode {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clip_children_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3581808349) // FIXME: should cache?
  var ret CanvasItemClipChildrenMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CanvasItem) GetPropVisible() bool {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropVisible(value bool) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropModulate() Color {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropModulate(value Color) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropSelfModulate() Color {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropSelfModulate(value Color) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropShowBehindParent() bool {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropShowBehindParent(value bool) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropTopLevel() bool {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropTopLevel(value bool) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropClipChildren() int {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropClipChildren(value int) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropLightMask() int {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropLightMask(value int) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropVisibilityLayer() int {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropVisibilityLayer(value int) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropZIndex() int {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropZIndex(value int) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropZAsRelative() bool {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropZAsRelative(value bool) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropYSortEnabled() bool {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropYSortEnabled(value bool) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropTextureFilter() int {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropTextureFilter(value int) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropTextureRepeat() int {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropTextureRepeat(value int) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropMaterial() any {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropMaterial(value any) {
  panic("TODO: implement")
}

func (me *CanvasItem) GetPropUseParentMaterial() bool {
  panic("TODO: implement")
}

func (me *CanvasItem) SetPropUseParentMaterial(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API