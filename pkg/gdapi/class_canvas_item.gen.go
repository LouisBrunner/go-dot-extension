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

type ptrsForCanvasItemList struct {
	fnXDraw                               gdc.MethodBindPtr
	fnGetCanvasItem                       gdc.MethodBindPtr
	fnSetVisible                          gdc.MethodBindPtr
	fnIsVisible                           gdc.MethodBindPtr
	fnIsVisibleInTree                     gdc.MethodBindPtr
	fnShow                                gdc.MethodBindPtr
	fnHide                                gdc.MethodBindPtr
	fnQueueRedraw                         gdc.MethodBindPtr
	fnMoveToFront                         gdc.MethodBindPtr
	fnSetAsTopLevel                       gdc.MethodBindPtr
	fnIsSetAsTopLevel                     gdc.MethodBindPtr
	fnSetLightMask                        gdc.MethodBindPtr
	fnGetLightMask                        gdc.MethodBindPtr
	fnSetModulate                         gdc.MethodBindPtr
	fnGetModulate                         gdc.MethodBindPtr
	fnSetSelfModulate                     gdc.MethodBindPtr
	fnGetSelfModulate                     gdc.MethodBindPtr
	fnSetZIndex                           gdc.MethodBindPtr
	fnGetZIndex                           gdc.MethodBindPtr
	fnSetZAsRelative                      gdc.MethodBindPtr
	fnIsZRelative                         gdc.MethodBindPtr
	fnSetYSortEnabled                     gdc.MethodBindPtr
	fnIsYSortEnabled                      gdc.MethodBindPtr
	fnSetDrawBehindParent                 gdc.MethodBindPtr
	fnIsDrawBehindParentEnabled           gdc.MethodBindPtr
	fnDrawLine                            gdc.MethodBindPtr
	fnDrawDashedLine                      gdc.MethodBindPtr
	fnDrawPolyline                        gdc.MethodBindPtr
	fnDrawPolylineColors                  gdc.MethodBindPtr
	fnDrawArc                             gdc.MethodBindPtr
	fnDrawMultiline                       gdc.MethodBindPtr
	fnDrawMultilineColors                 gdc.MethodBindPtr
	fnDrawRect                            gdc.MethodBindPtr
	fnDrawCircle                          gdc.MethodBindPtr
	fnDrawTexture                         gdc.MethodBindPtr
	fnDrawTextureRect                     gdc.MethodBindPtr
	fnDrawTextureRectRegion               gdc.MethodBindPtr
	fnDrawMsdfTextureRectRegion           gdc.MethodBindPtr
	fnDrawLcdTextureRectRegion            gdc.MethodBindPtr
	fnDrawStyleBox                        gdc.MethodBindPtr
	fnDrawPrimitive                       gdc.MethodBindPtr
	fnDrawPolygon                         gdc.MethodBindPtr
	fnDrawColoredPolygon                  gdc.MethodBindPtr
	fnDrawString                          gdc.MethodBindPtr
	fnDrawMultilineString                 gdc.MethodBindPtr
	fnDrawStringOutline                   gdc.MethodBindPtr
	fnDrawMultilineStringOutline          gdc.MethodBindPtr
	fnDrawChar                            gdc.MethodBindPtr
	fnDrawCharOutline                     gdc.MethodBindPtr
	fnDrawMesh                            gdc.MethodBindPtr
	fnDrawMultimesh                       gdc.MethodBindPtr
	fnDrawSetTransform                    gdc.MethodBindPtr
	fnDrawSetTransformMatrix              gdc.MethodBindPtr
	fnDrawAnimationSlice                  gdc.MethodBindPtr
	fnDrawEndAnimation                    gdc.MethodBindPtr
	fnGetTransform                        gdc.MethodBindPtr
	fnGetGlobalTransform                  gdc.MethodBindPtr
	fnGetGlobalTransformWithCanvas        gdc.MethodBindPtr
	fnGetViewportTransform                gdc.MethodBindPtr
	fnGetViewportRect                     gdc.MethodBindPtr
	fnGetCanvasTransform                  gdc.MethodBindPtr
	fnGetScreenTransform                  gdc.MethodBindPtr
	fnGetLocalMousePosition               gdc.MethodBindPtr
	fnGetGlobalMousePosition              gdc.MethodBindPtr
	fnGetCanvas                           gdc.MethodBindPtr
	fnGetWorld2D                          gdc.MethodBindPtr
	fnSetMaterial                         gdc.MethodBindPtr
	fnGetMaterial                         gdc.MethodBindPtr
	fnSetUseParentMaterial                gdc.MethodBindPtr
	fnGetUseParentMaterial                gdc.MethodBindPtr
	fnSetNotifyLocalTransform             gdc.MethodBindPtr
	fnIsLocalTransformNotificationEnabled gdc.MethodBindPtr
	fnSetNotifyTransform                  gdc.MethodBindPtr
	fnIsTransformNotificationEnabled      gdc.MethodBindPtr
	fnForceUpdateTransform                gdc.MethodBindPtr
	fnMakeCanvasPositionLocal             gdc.MethodBindPtr
	fnMakeInputLocal                      gdc.MethodBindPtr
	fnSetVisibilityLayer                  gdc.MethodBindPtr
	fnGetVisibilityLayer                  gdc.MethodBindPtr
	fnSetVisibilityLayerBit               gdc.MethodBindPtr
	fnGetVisibilityLayerBit               gdc.MethodBindPtr
	fnSetTextureFilter                    gdc.MethodBindPtr
	fnGetTextureFilter                    gdc.MethodBindPtr
	fnSetTextureRepeat                    gdc.MethodBindPtr
	fnGetTextureRepeat                    gdc.MethodBindPtr
	fnSetClipChildrenMode                 gdc.MethodBindPtr
	fnGetClipChildrenMode                 gdc.MethodBindPtr
}

var ptrsForCanvasItem ptrsForCanvasItemList

func initCanvasItemPtrs(iface gdc.Interface) {

	className := StringNameFromStr("CanvasItem")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_canvas_item")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetCanvasItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("set_visible")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_visible")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnIsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_visible_in_tree")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnIsVisibleInTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("show")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnShow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("hide")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnHide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("queue_redraw")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnQueueRedraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("move_to_front")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnMoveToFront = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_as_top_level")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetAsTopLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_set_as_top_level")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnIsSetAsTopLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_light_mask")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetLightMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_light_mask")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetLightMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_modulate")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_modulate")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_self_modulate")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetSelfModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_self_modulate")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetSelfModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_z_index")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetZIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_z_index")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetZIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_z_as_relative")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetZAsRelative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_z_relative")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnIsZRelative = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_y_sort_enabled")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetYSortEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_y_sort_enabled")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnIsYSortEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_draw_behind_parent")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetDrawBehindParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_draw_behind_parent_enabled")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnIsDrawBehindParentEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("draw_line")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1562330099))
	}
	{
		methodName := StringNameFromStr("draw_dashed_line")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawDashedLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 684651049))
	}
	{
		methodName := StringNameFromStr("draw_polyline")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawPolyline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3797364428))
	}
	{
		methodName := StringNameFromStr("draw_polyline_colors")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawPolylineColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2311979562))
	}
	{
		methodName := StringNameFromStr("draw_arc")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawArc = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4140652635))
	}
	{
		methodName := StringNameFromStr("draw_multiline")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawMultiline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2239075205))
	}
	{
		methodName := StringNameFromStr("draw_multiline_colors")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawMultilineColors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4072951537))
	}
	{
		methodName := StringNameFromStr("draw_rect")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2417231121))
	}
	{
		methodName := StringNameFromStr("draw_circle")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawCircle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3063020269))
	}
	{
		methodName := StringNameFromStr("draw_texture")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 520200117))
	}
	{
		methodName := StringNameFromStr("draw_texture_rect")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawTextureRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3832805018))
	}
	{
		methodName := StringNameFromStr("draw_texture_rect_region")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawTextureRectRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3883821411))
	}
	{
		methodName := StringNameFromStr("draw_msdf_texture_rect_region")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawMsdfTextureRectRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4219163252))
	}
	{
		methodName := StringNameFromStr("draw_lcd_texture_rect_region")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawLcdTextureRectRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3212350954))
	}
	{
		methodName := StringNameFromStr("draw_style_box")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawStyleBox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 388176283))
	}
	{
		methodName := StringNameFromStr("draw_primitive")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawPrimitive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3288481815))
	}
	{
		methodName := StringNameFromStr("draw_polygon")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 974537912))
	}
	{
		methodName := StringNameFromStr("draw_colored_polygon")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawColoredPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 15245644))
	}
	{
		methodName := StringNameFromStr("draw_string")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 728290553))
	}
	{
		methodName := StringNameFromStr("draw_multiline_string")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawMultilineString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1927038192))
	}
	{
		methodName := StringNameFromStr("draw_string_outline")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawStringOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 340562381))
	}
	{
		methodName := StringNameFromStr("draw_multiline_string_outline")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawMultilineStringOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1912318525))
	}
	{
		methodName := StringNameFromStr("draw_char")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawChar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3339793283))
	}
	{
		methodName := StringNameFromStr("draw_char_outline")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawCharOutline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3302344391))
	}
	{
		methodName := StringNameFromStr("draw_mesh")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 153818295))
	}
	{
		methodName := StringNameFromStr("draw_multimesh")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawMultimesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 937992368))
	}
	{
		methodName := StringNameFromStr("draw_set_transform")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawSetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 288975085))
	}
	{
		methodName := StringNameFromStr("draw_set_transform_matrix")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawSetTransformMatrix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
	}
	{
		methodName := StringNameFromStr("draw_animation_slice")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawAnimationSlice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3112831842))
	}
	{
		methodName := StringNameFromStr("draw_end_animation")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnDrawEndAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_transform")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_global_transform")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetGlobalTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_global_transform_with_canvas")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetGlobalTransformWithCanvas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_viewport_transform")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetViewportTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_viewport_rect")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetViewportRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("get_canvas_transform")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetCanvasTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_screen_transform")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetScreenTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_local_mouse_position")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetLocalMousePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_global_mouse_position")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetGlobalMousePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_canvas")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetCanvas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_world_2d")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetWorld2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339128592))
	}
	{
		methodName := StringNameFromStr("set_material")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
	}
	{
		methodName := StringNameFromStr("get_material")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
	}
	{
		methodName := StringNameFromStr("set_use_parent_material")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetUseParentMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_use_parent_material")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetUseParentMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_notify_local_transform")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetNotifyLocalTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_local_transform_notification_enabled")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnIsLocalTransformNotificationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_notify_transform")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetNotifyTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_transform_notification_enabled")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnIsTransformNotificationEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("force_update_transform")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnForceUpdateTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("make_canvas_position_local")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnMakeCanvasPositionLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2656412154))
	}
	{
		methodName := StringNameFromStr("make_input_local")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnMakeInputLocal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 811130057))
	}
	{
		methodName := StringNameFromStr("set_visibility_layer")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetVisibilityLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_visibility_layer")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetVisibilityLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_visibility_layer_bit")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetVisibilityLayerBit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_visibility_layer_bit")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetVisibilityLayerBit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_texture_filter")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1037999706))
	}
	{
		methodName := StringNameFromStr("get_texture_filter")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 121960042))
	}
	{
		methodName := StringNameFromStr("set_texture_repeat")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1716472974))
	}
	{
		methodName := StringNameFromStr("get_texture_repeat")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2667158319))
	}
	{
		methodName := StringNameFromStr("set_clip_children_mode")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnSetClipChildrenMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1319393776))
	}
	{
		methodName := StringNameFromStr("get_clip_children_mode")
		defer methodName.Destroy()
		ptrsForCanvasItem.fnGetClipChildrenMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3581808349))
	}

}

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
	CanvasItemNotificationTransformChanged      = 2000
	CanvasItemNotificationLocalTransformChanged = 35
	CanvasItemNotificationDraw                  = 30
	CanvasItemNotificationVisibilityChanged     = 31
	CanvasItemNotificationEnterCanvas           = 32
	CanvasItemNotificationExitCanvas            = 33
	CanvasItemNotificationWorld2DChanged        = 36
)

// Enums

type CanvasItemTextureFilter int

const (
	CanvasItemTextureFilterTextureFilterParentNode                    CanvasItemTextureFilter = 0
	CanvasItemTextureFilterTextureFilterNearest                       CanvasItemTextureFilter = 1
	CanvasItemTextureFilterTextureFilterLinear                        CanvasItemTextureFilter = 2
	CanvasItemTextureFilterTextureFilterNearestWithMipmaps            CanvasItemTextureFilter = 3
	CanvasItemTextureFilterTextureFilterLinearWithMipmaps             CanvasItemTextureFilter = 4
	CanvasItemTextureFilterTextureFilterNearestWithMipmapsAnisotropic CanvasItemTextureFilter = 5
	CanvasItemTextureFilterTextureFilterLinearWithMipmapsAnisotropic  CanvasItemTextureFilter = 6
	CanvasItemTextureFilterTextureFilterMax                           CanvasItemTextureFilter = 7
)

type CanvasItemTextureRepeat int

const (
	CanvasItemTextureRepeatTextureRepeatParentNode CanvasItemTextureRepeat = 0
	CanvasItemTextureRepeatTextureRepeatDisabled   CanvasItemTextureRepeat = 1
	CanvasItemTextureRepeatTextureRepeatEnabled    CanvasItemTextureRepeat = 2
	CanvasItemTextureRepeatTextureRepeatMirror     CanvasItemTextureRepeat = 3
	CanvasItemTextureRepeatTextureRepeatMax        CanvasItemTextureRepeat = 4
)

type CanvasItemClipChildrenMode int

const (
	CanvasItemClipChildrenModeClipChildrenDisabled CanvasItemClipChildrenMode = 0
	CanvasItemClipChildrenModeClipChildrenOnly     CanvasItemClipChildrenMode = 1
	CanvasItemClipChildrenModeClipChildrenAndDraw  CanvasItemClipChildrenMode = 2
	CanvasItemClipChildrenModeClipChildrenMax      CanvasItemClipChildrenMode = 3
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

func (me *CanvasItem) GetCanvasItem() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetCanvasItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) SetVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) IsVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnIsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) IsVisibleInTree() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnIsVisibleInTree), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) Show() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnShow), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) Hide() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnHide), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) QueueRedraw() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnQueueRedraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) MoveToFront() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnMoveToFront), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) SetAsTopLevel(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetAsTopLevel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) IsSetAsTopLevel() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnIsSetAsTopLevel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) SetLightMask(light_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&light_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetLightMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetLightMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetLightMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) SetModulate(modulate Color) {
	cargs := []gdc.ConstTypePtr{modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetModulate() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) SetSelfModulate(self_modulate Color) {
	cargs := []gdc.ConstTypePtr{self_modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetSelfModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetSelfModulate() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetSelfModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) SetZIndex(z_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&z_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetZIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetZIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetZIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) SetZAsRelative(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetZAsRelative), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) IsZRelative() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnIsZRelative), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) SetYSortEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetYSortEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) IsYSortEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnIsYSortEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) SetDrawBehindParent(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetDrawBehindParent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) IsDrawBehindParentEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnIsDrawBehindParentEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) DrawLine(from Vector2, to Vector2, color Color, width float64, antialiased bool) {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&antialiased)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawLine), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawDashedLine(from Vector2, to Vector2, color Color, width float64, dash float64, aligned bool) {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&dash), gdc.ConstTypePtr(&aligned)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawDashedLine), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawPolyline(points PackedVector2Array, color Color, width float64, antialiased bool) {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&antialiased)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawPolyline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawPolylineColors(points PackedVector2Array, colors PackedColorArray, width float64, antialiased bool) {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), colors.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&antialiased)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawPolylineColors), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawArc(center Vector2, radius float64, start_angle float64, end_angle float64, point_count int64, color Color, width float64, antialiased bool) {
	cargs := []gdc.ConstTypePtr{center.AsCTypePtr(), gdc.ConstTypePtr(&radius), gdc.ConstTypePtr(&start_angle), gdc.ConstTypePtr(&end_angle), gdc.ConstTypePtr(&point_count), color.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&antialiased)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawArc), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawMultiline(points PackedVector2Array, color Color, width float64) {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawMultiline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawMultilineColors(points PackedVector2Array, colors PackedColorArray, width float64) {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), colors.AsCTypePtr(), gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawMultilineColors), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawRect(rect Rect2, color Color, filled bool, width float64) {
	cargs := []gdc.ConstTypePtr{rect.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&filled), gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawCircle(position Vector2, radius float64, color Color) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&radius), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawCircle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawTexture(texture Texture2D, position Vector2, modulate Color) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), position.AsCTypePtr(), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawTextureRect(texture Texture2D, rect Rect2, tile bool, modulate Color, transpose bool) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), rect.AsCTypePtr(), gdc.ConstTypePtr(&tile), modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawTextureRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, transpose bool, clip_uv bool) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), rect.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&transpose), gdc.ConstTypePtr(&clip_uv)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawTextureRectRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawMsdfTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color, outline float64, pixel_range float64, scale float64) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), rect.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr(), gdc.ConstTypePtr(&outline), gdc.ConstTypePtr(&pixel_range), gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawMsdfTextureRectRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawLcdTextureRectRegion(texture Texture2D, rect Rect2, src_rect Rect2, modulate Color) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), rect.AsCTypePtr(), src_rect.AsCTypePtr(), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawLcdTextureRectRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawStyleBox(style_box StyleBox, rect Rect2) {
	cargs := []gdc.ConstTypePtr{style_box.AsCTypePtr(), rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawStyleBox), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawPrimitive(points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), colors.AsCTypePtr(), uvs.AsCTypePtr(), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawPrimitive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawPolygon(points PackedVector2Array, colors PackedColorArray, uvs PackedVector2Array, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), colors.AsCTypePtr(), uvs.AsCTypePtr(), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawColoredPolygon(points PackedVector2Array, color Color, uvs PackedVector2Array, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), color.AsCTypePtr(), uvs.AsCTypePtr(), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawColoredPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawString(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), modulate.AsCTypePtr(), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawString), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawMultilineString(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, max_lines int64, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), modulate.AsCTypePtr(), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawMultilineString), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawStringOutline(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, size int64, modulate Color, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&size), modulate.AsCTypePtr(), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawStringOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawMultilineStringOutline(font Font, pos Vector2, text String, alignment HorizontalAlignment, width float64, font_size int64, max_lines int64, size int64, modulate Color, brk_flags TextServerLineBreakFlag, justification_flags TextServerJustificationFlag, direction TextServerDirection, orientation TextServerOrientation) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), text.AsCTypePtr(), gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&max_lines), gdc.ConstTypePtr(&size), modulate.AsCTypePtr(), gdc.ConstTypePtr(&brk_flags), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawMultilineStringOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawChar(font Font, pos Vector2, char String, font_size int64, modulate Color) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), char.AsCTypePtr(), gdc.ConstTypePtr(&font_size), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawChar), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawCharOutline(font Font, pos Vector2, char String, font_size int64, size int64, modulate Color) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), pos.AsCTypePtr(), char.AsCTypePtr(), gdc.ConstTypePtr(&font_size), gdc.ConstTypePtr(&size), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawCharOutline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawMesh(mesh Mesh, texture Texture2D, transform Transform2D, modulate Color) {
	cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), texture.AsCTypePtr(), transform.AsCTypePtr(), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawMultimesh(multimesh MultiMesh, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{multimesh.AsCTypePtr(), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawMultimesh), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawSetTransform(position Vector2, rotation float64, scale Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&rotation), scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawSetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawSetTransformMatrix(xform Transform2D) {
	cargs := []gdc.ConstTypePtr{xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawSetTransformMatrix), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawAnimationSlice(animation_length float64, slice_begin float64, slice_end float64, offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&animation_length), gdc.ConstTypePtr(&slice_begin), gdc.ConstTypePtr(&slice_end), gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawAnimationSlice), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) DrawEndAnimation() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnDrawEndAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetGlobalTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetGlobalTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetGlobalTransformWithCanvas() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetGlobalTransformWithCanvas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetViewportTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetViewportTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetViewportRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetViewportRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetCanvasTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetCanvasTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetScreenTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetScreenTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetLocalMousePosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetLocalMousePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetGlobalMousePosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetGlobalMousePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetCanvas() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetCanvas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) GetWorld2D() World2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewWorld2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetWorld2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) SetMaterial(material Material) {
	cargs := []gdc.ConstTypePtr{material.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetMaterial() Material {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMaterial()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) SetUseParentMaterial(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetUseParentMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetUseParentMaterial() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetUseParentMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) SetNotifyLocalTransform(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetNotifyLocalTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) IsLocalTransformNotificationEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnIsLocalTransformNotificationEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) SetNotifyTransform(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetNotifyTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) IsTransformNotificationEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnIsTransformNotificationEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) ForceUpdateTransform() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnForceUpdateTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) MakeCanvasPositionLocal(screen_point Vector2) Vector2 {
	cargs := []gdc.ConstTypePtr{screen_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnMakeCanvasPositionLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) MakeInputLocal(event InputEvent) InputEvent {
	cargs := []gdc.ConstTypePtr{event.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInputEvent()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnMakeInputLocal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *CanvasItem) SetVisibilityLayer(layer int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetVisibilityLayer), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetVisibilityLayer() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetVisibilityLayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) SetVisibilityLayerBit(layer int64, enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetVisibilityLayerBit), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetVisibilityLayerBit(layer int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetVisibilityLayerBit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *CanvasItem) SetTextureFilter(mode CanvasItemTextureFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetTextureFilter() CanvasItemTextureFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CanvasItemTextureFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetTextureFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CanvasItem) SetTextureRepeat(mode CanvasItemTextureRepeat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetTextureRepeat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetTextureRepeat() CanvasItemTextureRepeat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CanvasItemTextureRepeat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetTextureRepeat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *CanvasItem) SetClipChildrenMode(mode CanvasItemClipChildrenMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnSetClipChildrenMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *CanvasItem) GetClipChildrenMode() CanvasItemClipChildrenMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret CanvasItemClipChildrenMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCanvasItem.fnGetClipChildrenMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
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
