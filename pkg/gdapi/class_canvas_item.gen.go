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

type CanvasItem struct {
  Node
}

func (me *CanvasItem) BaseClass() string {
  return "CanvasItem"
}

func NewCanvasItem() *CanvasItem {
  str := StringNameFromStr("CanvasItem") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CanvasItem{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) SetVisible(visible bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) IsVisible() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) IsVisibleInTree() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_visible_in_tree")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) Show()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("show")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) Hide()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) QueueRedraw()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("queue_redraw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) MoveToFront()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_to_front")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) SetAsTopLevel(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_as_top_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) IsSetAsTopLevel() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_set_as_top_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) SetLightMask(light_mask int64, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetLightMask() int64 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_light_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) SetModulate(modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetModulate() Color {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) SetSelfModulate(self_modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_self_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{self_modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetSelfModulate() Color {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_self_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) SetZIndex(z_index int64, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&z_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetZIndex() int64 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_z_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) SetZAsRelative(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_z_as_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) IsZRelative() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_z_relative")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) SetYSortEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_y_sort_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) IsYSortEnabled() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_y_sort_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) SetDrawBehindParent(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_behind_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) IsDrawBehindParentEnabled() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draw_behind_parent_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) DrawLine(from Vector2, to Vector2, color Color, width float64, antialiased bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1562330099) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&antialiased) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawDashedLine(from Vector2, to Vector2, color Color, width float64, dash float64, aligned bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_dashed_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 684651049) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&dash) , gdc.ConstTypePtr(&aligned) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawPolyline(points PackedVector2Array, color Color, width float64, antialiased bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_polyline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3797364428) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&antialiased) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawPolylineColors(points PackedVector2Array, colors PackedColorArray, width float64, antialiased bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_polyline_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2311979562) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), colors.AsCTypePtr(), gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&antialiased) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawArc(center Vector2, radius float64, start_angle float64, end_angle float64, point_count int64, color Color, width float64, antialiased bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_arc")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4140652635) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{center.AsCTypePtr(), gdc.ConstTypePtr(&radius) , gdc.ConstTypePtr(&start_angle) , gdc.ConstTypePtr(&end_angle) , gdc.ConstTypePtr(&point_count) , color.AsCTypePtr(), gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&antialiased) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawMultiline(points PackedVector2Array, color Color, width float64, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2239075205) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawMultilineColors(points PackedVector2Array, colors PackedColorArray, width float64, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline_colors")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4072951537) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), colors.AsCTypePtr(), gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawRect(rect Rect2, color Color, filled bool, width float64, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2417231121) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{rect.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&filled) , gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawCircle(position Vector2, radius float64, color Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_circle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3063020269) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&radius) , color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawTexture(texture Texture2D, position Vector2, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 520200117) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), position.AsCTypePtr(), modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawTextureRect(texture Texture2D, rect Rect2, tile bool, modulate Color, transpose bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_texture_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3832805018) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), rect.AsCTypePtr(), gdc.ConstTypePtr(&tile) , modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, transpose bool, clip_uv bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_texture_rect_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3883821411) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), rect.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose) , gdc.ConstTypePtr(&clip_uv) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawMsdfTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, outline float64, pixel_range float64, scale float64, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_msdf_texture_rect_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4219163252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), rect.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&outline) , gdc.ConstTypePtr(&pixel_range) , gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawLcdTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_lcd_texture_rect_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3212350954) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), rect.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawStyleBox(style_box StyleBox, rect Rect2, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_style_box")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 388176283) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{style_box.AsCTypePtr(), rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawPrimitive(points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_primitive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3288481815) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), colors.AsCTypePtr(), uvs.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawPolygon(points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 974537912) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), colors.AsCTypePtr(), uvs.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawColoredPolygon(points PackedVector2Array, color Color, uvs PackedVector2Array, texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_colored_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 15245644) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), color.AsCTypePtr(), uvs.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawString(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 728290553) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment) , gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&font_size) , modulate.AsCTypePtr(), gdc.ConstTypePtr(&justification_flags) , gdc.ConstTypePtr(&direction) , gdc.ConstTypePtr(&orientation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawMultilineString(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, max_lines int64, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1927038192) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment) , gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&font_size) , gdc.ConstTypePtr(&max_lines) , modulate.AsCTypePtr(), gdc.ConstTypePtr(&brk_flags) , gdc.ConstTypePtr(&justification_flags) , gdc.ConstTypePtr(&direction) , gdc.ConstTypePtr(&orientation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawStringOutline(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, size int64, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_string_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 340562381) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment) , gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&font_size) , gdc.ConstTypePtr(&size) , modulate.AsCTypePtr(), gdc.ConstTypePtr(&justification_flags) , gdc.ConstTypePtr(&direction) , gdc.ConstTypePtr(&orientation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawMultilineStringOutline(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, max_lines int64, size int64, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multiline_string_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1912318525) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment) , gdc.ConstTypePtr(&width) , gdc.ConstTypePtr(&font_size) , gdc.ConstTypePtr(&max_lines) , gdc.ConstTypePtr(&size) , modulate.AsCTypePtr(), gdc.ConstTypePtr(&brk_flags) , gdc.ConstTypePtr(&justification_flags) , gdc.ConstTypePtr(&direction) , gdc.ConstTypePtr(&orientation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawChar(font Font, pos Vector2, char String, font_size int64, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_char")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3339793283) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), char.AsCTypePtr(), gdc.ConstTypePtr(&font_size) , modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawCharOutline(font Font, pos Vector2, char String, font_size int64, size int64, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_char_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3302344391) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), char.AsCTypePtr(), gdc.ConstTypePtr(&font_size) , gdc.ConstTypePtr(&size) , modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawMesh(mesh Mesh, texture Texture2D, transform Transform2D, modulate Color, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 153818295) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), texture.AsCTypePtr(), transform.AsCTypePtr(), modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawMultimesh(multimesh MultiMesh, texture Texture2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_multimesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 937992368) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawSetTransform(position Vector2, rotation float64, scale Vector2, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 288975085) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&rotation) , scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawSetTransformMatrix(xform Transform2D, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_set_transform_matrix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawAnimationSlice(animation_length float64, slice_begin float64, slice_end float64, offset float64, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_animation_slice")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3112831842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&animation_length) , gdc.ConstTypePtr(&slice_begin) , gdc.ConstTypePtr(&slice_end) , gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) DrawEndAnimation()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_end_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetGlobalTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetGlobalTransformWithCanvas() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_transform_with_canvas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetViewportTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_viewport_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetViewportRect() Rect2 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_viewport_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetCanvasTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetScreenTransform() Transform2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814499831) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetLocalMousePosition() Vector2 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_local_mouse_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetGlobalMousePosition() Vector2 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_mouse_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetCanvas() RID {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_canvas")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) GetWorld2D() World2D {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_2d")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339128592) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewWorld2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetMaterial() Material {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) SetUseParentMaterial(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_parent_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetUseParentMaterial() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_parent_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) SetNotifyLocalTransform(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_notify_local_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) IsLocalTransformNotificationEnabled() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_local_transform_notification_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) SetNotifyTransform(enable bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_notify_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) IsTransformNotificationEnabled() bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_transform_notification_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) ForceUpdateTransform()  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) MakeCanvasPositionLocal(screen_point Vector2, ) Vector2 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_canvas_position_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2656412154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{screen_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) MakeInputLocal(event InputEvent, ) InputEvent {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_input_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 811130057) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInputEvent()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CanvasItem) SetVisibilityLayer(layer int64, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetVisibilityLayer() int64 {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) SetVisibilityLayerBit(layer int64, enabled bool, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visibility_layer_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetVisibilityLayerBit(layer int64, ) bool {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visibility_layer_bit")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&layer)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CanvasItem) SetTextureFilter(mode CanvasItemTextureFilter, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1037999706) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetTextureFilter() CanvasItemTextureFilter {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 121960042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CanvasItemTextureFilter

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CanvasItem) SetTextureRepeat(mode CanvasItemTextureRepeat, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1716472974) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetTextureRepeat() CanvasItemTextureRepeat {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2667158319) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CanvasItemTextureRepeat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CanvasItem) SetClipChildrenMode(mode CanvasItemClipChildrenMode, )  {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clip_children_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1319393776) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CanvasItem) GetClipChildrenMode() CanvasItemClipChildrenMode {
  classNameV := StringNameFromStr("CanvasItem")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clip_children_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3581808349) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret CanvasItemClipChildrenMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CanvasItemDrawSignalFn func()

func (me *CanvasItem) ConnectDraw(subs SignalSubscribers, fn CanvasItemDrawSignalFn) {
  sig := StringNameFromStr("draw")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CanvasItem) DisconnectDraw(subs SignalSubscribers, fn CanvasItemDrawSignalFn) {
  sig := StringNameFromStr("draw")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CanvasItemVisibilityChangedSignalFn func()

func (me *CanvasItem) ConnectVisibilityChanged(subs SignalSubscribers, fn CanvasItemVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CanvasItem) DisconnectVisibilityChanged(subs SignalSubscribers, fn CanvasItemVisibilityChangedSignalFn) {
  sig := StringNameFromStr("visibility_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CanvasItemHiddenSignalFn func()

func (me *CanvasItem) ConnectHidden(subs SignalSubscribers, fn CanvasItemHiddenSignalFn) {
  sig := StringNameFromStr("hidden")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CanvasItem) DisconnectHidden(subs SignalSubscribers, fn CanvasItemHiddenSignalFn) {
  sig := StringNameFromStr("hidden")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CanvasItemItemRectChangedSignalFn func()

func (me *CanvasItem) ConnectItemRectChanged(subs SignalSubscribers, fn CanvasItemItemRectChangedSignalFn) {
  sig := StringNameFromStr("item_rect_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CanvasItem) DisconnectItemRectChanged(subs SignalSubscribers, fn CanvasItemItemRectChangedSignalFn) {
  sig := StringNameFromStr("item_rect_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
