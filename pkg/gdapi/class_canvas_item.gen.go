// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasItem struct {
  obj gdc.ObjectPtr
}

func (me *CanvasItem) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CanvasItem) BaseClass() string {
  return "CanvasItem"
}

// TODO: needed?
// const (
// // )

var (
  CanvasItemNotificationTransformChanged = "2000" // TODO: construct correctly
  CanvasItemNotificationLocalTransformChanged = "35" // TODO: construct correctly
  CanvasItemNotificationDraw = "30" // TODO: construct correctly
  CanvasItemNotificationVisibilityChanged = "31" // TODO: construct correctly
  CanvasItemNotificationEnterCanvas = "32" // TODO: construct correctly
  CanvasItemNotificationExitCanvas = "33" // TODO: construct correctly
  CanvasItemNotificationWorld2DChanged = "36" // TODO: construct correctly
)

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

func  (me *CanvasItem) XDraw() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetCanvasItem() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetVisible(visible bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) IsVisible() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) IsVisibleInTree() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) Show() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) Hide() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) QueueRedraw() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) MoveToFront() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetAsTopLevel(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) IsSetAsTopLevel() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetLightMask(light_mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetLightMask() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetModulate(modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetModulate() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetSelfModulate(self_modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetSelfModulate() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetZIndex(z_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetZIndex() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetZAsRelative(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) IsZRelative() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetYSortEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) IsYSortEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetDrawBehindParent(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) IsDrawBehindParentEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawLine(from Vector2, to Vector2, color Color, width float32, antialiased bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawDashedLine(from Vector2, to Vector2, color Color, width float32, dash float32, aligned bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawPolyline(points PackedVector2Array, color Color, width float32, antialiased bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawPolylineColors(points PackedVector2Array, colors PackedColorArray, width float32, antialiased bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawArc(center Vector2, radius float32, start_angle float32, end_angle float32, point_count int, color Color, width float32, antialiased bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawMultiline(points PackedVector2Array, color Color, width float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawMultilineColors(points PackedVector2Array, colors PackedColorArray, width float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawRect(rect Rect2, color Color, filled bool, width float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawCircle(position Vector2, radius float32, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawTexture(texture Texture2D, position Vector2, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawTextureRect(texture Texture2D, rect Rect2, tile bool, modulate Color, transpose bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, transpose bool, clip_uv bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawMsdfTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, outline float32, pixel_range float32, scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawLcdTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawStyleBox(style_box StyleBox, rect Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawPrimitive(points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawPolygon(points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawColoredPolygon(points PackedVector2Array, color Color, uvs PackedVector2Array, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawString(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawMultilineString(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawStringOutline(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, size int, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawMultilineStringOutline(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float32, font_size int, max_lines int, size int, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawChar(font Font, pos Vector2, char String, font_size int, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawCharOutline(font Font, pos Vector2, char String, font_size int, size int, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawMesh(mesh Mesh, texture Texture2D, transform Transform2D, modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawMultimesh(multimesh MultiMesh, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawSetTransform(position Vector2, rotation float32, scale Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawSetTransformMatrix(xform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawAnimationSlice(animation_length float32, slice_begin float32, slice_end float32, offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) DrawEndAnimation() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetTransform() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetGlobalTransform() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetGlobalTransformWithCanvas() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetViewportTransform() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetViewportRect() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetCanvasTransform() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetScreenTransform() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetLocalMousePosition() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetGlobalMousePosition() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetCanvas() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetWorld2D() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetMaterial(material Material, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetMaterial() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetUseParentMaterial(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetUseParentMaterial() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetNotifyLocalTransform(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) IsLocalTransformNotificationEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetNotifyTransform(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) IsTransformNotificationEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) ForceUpdateTransform() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) MakeCanvasPositionLocal(screen_point Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) MakeInputLocal(event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetVisibilityLayer(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetVisibilityLayer() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetVisibilityLayerBit(layer int, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetVisibilityLayerBit(layer int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetTextureFilter(mode CanvasItemTextureFilter, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetTextureFilter() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetTextureRepeat(mode CanvasItemTextureRepeat, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetTextureRepeat() { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) SetClipChildrenMode(mode CanvasItemClipChildrenMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *CanvasItem) GetClipChildrenMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
