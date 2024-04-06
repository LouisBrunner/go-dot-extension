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

type ptrsForViewportList struct {
	fnSetWorld2D                             gdc.MethodBindPtr
	fnGetWorld2D                             gdc.MethodBindPtr
	fnFindWorld2D                            gdc.MethodBindPtr
	fnSetCanvasTransform                     gdc.MethodBindPtr
	fnGetCanvasTransform                     gdc.MethodBindPtr
	fnSetGlobalCanvasTransform               gdc.MethodBindPtr
	fnGetGlobalCanvasTransform               gdc.MethodBindPtr
	fnGetFinalTransform                      gdc.MethodBindPtr
	fnGetScreenTransform                     gdc.MethodBindPtr
	fnGetVisibleRect                         gdc.MethodBindPtr
	fnSetTransparentBackground               gdc.MethodBindPtr
	fnHasTransparentBackground               gdc.MethodBindPtr
	fnSetUseHdr2D                            gdc.MethodBindPtr
	fnIsUsingHdr2D                           gdc.MethodBindPtr
	fnSetMsaa2D                              gdc.MethodBindPtr
	fnGetMsaa2D                              gdc.MethodBindPtr
	fnSetMsaa3D                              gdc.MethodBindPtr
	fnGetMsaa3D                              gdc.MethodBindPtr
	fnSetScreenSpaceAa                       gdc.MethodBindPtr
	fnGetScreenSpaceAa                       gdc.MethodBindPtr
	fnSetUseTaa                              gdc.MethodBindPtr
	fnIsUsingTaa                             gdc.MethodBindPtr
	fnSetUseDebanding                        gdc.MethodBindPtr
	fnIsUsingDebanding                       gdc.MethodBindPtr
	fnSetUseOcclusionCulling                 gdc.MethodBindPtr
	fnIsUsingOcclusionCulling                gdc.MethodBindPtr
	fnSetDebugDraw                           gdc.MethodBindPtr
	fnGetDebugDraw                           gdc.MethodBindPtr
	fnGetRenderInfo                          gdc.MethodBindPtr
	fnGetTexture                             gdc.MethodBindPtr
	fnSetPhysicsObjectPicking                gdc.MethodBindPtr
	fnGetPhysicsObjectPicking                gdc.MethodBindPtr
	fnSetPhysicsObjectPickingSort            gdc.MethodBindPtr
	fnGetPhysicsObjectPickingSort            gdc.MethodBindPtr
	fnGetViewportRid                         gdc.MethodBindPtr
	fnPushTextInput                          gdc.MethodBindPtr
	fnPushInput                              gdc.MethodBindPtr
	fnPushUnhandledInput                     gdc.MethodBindPtr
	fnGetCamera2D                            gdc.MethodBindPtr
	fnSetAsAudioListener2D                   gdc.MethodBindPtr
	fnIsAudioListener2D                      gdc.MethodBindPtr
	fnGetMousePosition                       gdc.MethodBindPtr
	fnWarpMouse                              gdc.MethodBindPtr
	fnUpdateMouseCursorState                 gdc.MethodBindPtr
	fnGuiGetDragData                         gdc.MethodBindPtr
	fnGuiIsDragging                          gdc.MethodBindPtr
	fnGuiIsDragSuccessful                    gdc.MethodBindPtr
	fnGuiReleaseFocus                        gdc.MethodBindPtr
	fnGuiGetFocusOwner                       gdc.MethodBindPtr
	fnSetDisableInput                        gdc.MethodBindPtr
	fnIsInputDisabled                        gdc.MethodBindPtr
	fnSetPositionalShadowAtlasSize           gdc.MethodBindPtr
	fnGetPositionalShadowAtlasSize           gdc.MethodBindPtr
	fnSetPositionalShadowAtlas16Bits         gdc.MethodBindPtr
	fnGetPositionalShadowAtlas16Bits         gdc.MethodBindPtr
	fnSetSnapControlsToPixels                gdc.MethodBindPtr
	fnIsSnapControlsToPixelsEnabled          gdc.MethodBindPtr
	fnSetSnap2DTransformsToPixel             gdc.MethodBindPtr
	fnIsSnap2DTransformsToPixelEnabled       gdc.MethodBindPtr
	fnSetSnap2DVerticesToPixel               gdc.MethodBindPtr
	fnIsSnap2DVerticesToPixelEnabled         gdc.MethodBindPtr
	fnSetPositionalShadowAtlasQuadrantSubdiv gdc.MethodBindPtr
	fnGetPositionalShadowAtlasQuadrantSubdiv gdc.MethodBindPtr
	fnSetInputAsHandled                      gdc.MethodBindPtr
	fnIsInputHandled                         gdc.MethodBindPtr
	fnSetHandleInputLocally                  gdc.MethodBindPtr
	fnIsHandlingInputLocally                 gdc.MethodBindPtr
	fnSetDefaultCanvasItemTextureFilter      gdc.MethodBindPtr
	fnGetDefaultCanvasItemTextureFilter      gdc.MethodBindPtr
	fnSetEmbeddingSubwindows                 gdc.MethodBindPtr
	fnIsEmbeddingSubwindows                  gdc.MethodBindPtr
	fnGetEmbeddedSubwindows                  gdc.MethodBindPtr
	fnSetCanvasCullMask                      gdc.MethodBindPtr
	fnGetCanvasCullMask                      gdc.MethodBindPtr
	fnSetCanvasCullMaskBit                   gdc.MethodBindPtr
	fnGetCanvasCullMaskBit                   gdc.MethodBindPtr
	fnSetDefaultCanvasItemTextureRepeat      gdc.MethodBindPtr
	fnGetDefaultCanvasItemTextureRepeat      gdc.MethodBindPtr
	fnSetSdfOversize                         gdc.MethodBindPtr
	fnGetSdfOversize                         gdc.MethodBindPtr
	fnSetSdfScale                            gdc.MethodBindPtr
	fnGetSdfScale                            gdc.MethodBindPtr
	fnSetMeshLodThreshold                    gdc.MethodBindPtr
	fnGetMeshLodThreshold                    gdc.MethodBindPtr
	fnSetWorld3D                             gdc.MethodBindPtr
	fnGetWorld3D                             gdc.MethodBindPtr
	fnFindWorld3D                            gdc.MethodBindPtr
	fnSetUseOwnWorld3D                       gdc.MethodBindPtr
	fnIsUsingOwnWorld3D                      gdc.MethodBindPtr
	fnGetCamera3D                            gdc.MethodBindPtr
	fnSetAsAudioListener3D                   gdc.MethodBindPtr
	fnIsAudioListener3D                      gdc.MethodBindPtr
	fnSetDisable3D                           gdc.MethodBindPtr
	fnIs3DDisabled                           gdc.MethodBindPtr
	fnSetUseXr                               gdc.MethodBindPtr
	fnIsUsingXr                              gdc.MethodBindPtr
	fnSetScaling3DMode                       gdc.MethodBindPtr
	fnGetScaling3DMode                       gdc.MethodBindPtr
	fnSetScaling3DScale                      gdc.MethodBindPtr
	fnGetScaling3DScale                      gdc.MethodBindPtr
	fnSetFsrSharpness                        gdc.MethodBindPtr
	fnGetFsrSharpness                        gdc.MethodBindPtr
	fnSetTextureMipmapBias                   gdc.MethodBindPtr
	fnGetTextureMipmapBias                   gdc.MethodBindPtr
	fnSetVrsMode                             gdc.MethodBindPtr
	fnGetVrsMode                             gdc.MethodBindPtr
	fnSetVrsTexture                          gdc.MethodBindPtr
	fnGetVrsTexture                          gdc.MethodBindPtr
}

var ptrsForViewport ptrsForViewportList

func initViewportPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Viewport")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_world_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnSetWorld2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2736080068))
	}
	{
		methodName := StringNameFromStr("get_world_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnGetWorld2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339128592))
	}
	{
		methodName := StringNameFromStr("find_world_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnFindWorld2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339128592))
	}
	{
		methodName := StringNameFromStr("set_canvas_transform")
		defer methodName.Destroy()
		ptrsForViewport.fnSetCanvasTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
	}
	{
		methodName := StringNameFromStr("get_canvas_transform")
		defer methodName.Destroy()
		ptrsForViewport.fnGetCanvasTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("set_global_canvas_transform")
		defer methodName.Destroy()
		ptrsForViewport.fnSetGlobalCanvasTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
	}
	{
		methodName := StringNameFromStr("get_global_canvas_transform")
		defer methodName.Destroy()
		ptrsForViewport.fnGetGlobalCanvasTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_final_transform")
		defer methodName.Destroy()
		ptrsForViewport.fnGetFinalTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_screen_transform")
		defer methodName.Destroy()
		ptrsForViewport.fnGetScreenTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("get_visible_rect")
		defer methodName.Destroy()
		ptrsForViewport.fnGetVisibleRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("set_transparent_background")
		defer methodName.Destroy()
		ptrsForViewport.fnSetTransparentBackground = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_transparent_background")
		defer methodName.Destroy()
		ptrsForViewport.fnHasTransparentBackground = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_hdr_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnSetUseHdr2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_hdr_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnIsUsingHdr2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_msaa_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnSetMsaa2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3330258708))
	}
	{
		methodName := StringNameFromStr("get_msaa_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnGetMsaa2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2542055527))
	}
	{
		methodName := StringNameFromStr("set_msaa_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnSetMsaa3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3330258708))
	}
	{
		methodName := StringNameFromStr("get_msaa_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnGetMsaa3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2542055527))
	}
	{
		methodName := StringNameFromStr("set_screen_space_aa")
		defer methodName.Destroy()
		ptrsForViewport.fnSetScreenSpaceAa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3544169389))
	}
	{
		methodName := StringNameFromStr("get_screen_space_aa")
		defer methodName.Destroy()
		ptrsForViewport.fnGetScreenSpaceAa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1390814124))
	}
	{
		methodName := StringNameFromStr("set_use_taa")
		defer methodName.Destroy()
		ptrsForViewport.fnSetUseTaa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_taa")
		defer methodName.Destroy()
		ptrsForViewport.fnIsUsingTaa = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_debanding")
		defer methodName.Destroy()
		ptrsForViewport.fnSetUseDebanding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_debanding")
		defer methodName.Destroy()
		ptrsForViewport.fnIsUsingDebanding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_occlusion_culling")
		defer methodName.Destroy()
		ptrsForViewport.fnSetUseOcclusionCulling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_occlusion_culling")
		defer methodName.Destroy()
		ptrsForViewport.fnIsUsingOcclusionCulling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_debug_draw")
		defer methodName.Destroy()
		ptrsForViewport.fnSetDebugDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1970246205))
	}
	{
		methodName := StringNameFromStr("get_debug_draw")
		defer methodName.Destroy()
		ptrsForViewport.fnGetDebugDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 579191299))
	}
	{
		methodName := StringNameFromStr("get_render_info")
		defer methodName.Destroy()
		ptrsForViewport.fnGetRenderInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 481977019))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForViewport.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1746695840))
	}
	{
		methodName := StringNameFromStr("set_physics_object_picking")
		defer methodName.Destroy()
		ptrsForViewport.fnSetPhysicsObjectPicking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_physics_object_picking")
		defer methodName.Destroy()
		ptrsForViewport.fnGetPhysicsObjectPicking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_physics_object_picking_sort")
		defer methodName.Destroy()
		ptrsForViewport.fnSetPhysicsObjectPickingSort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_physics_object_picking_sort")
		defer methodName.Destroy()
		ptrsForViewport.fnGetPhysicsObjectPickingSort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_viewport_rid")
		defer methodName.Destroy()
		ptrsForViewport.fnGetViewportRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("push_text_input")
		defer methodName.Destroy()
		ptrsForViewport.fnPushTextInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("push_input")
		defer methodName.Destroy()
		ptrsForViewport.fnPushInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3644664830))
	}
	{
		methodName := StringNameFromStr("push_unhandled_input")
		defer methodName.Destroy()
		ptrsForViewport.fnPushUnhandledInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3644664830))
	}
	{
		methodName := StringNameFromStr("get_camera_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnGetCamera2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3551466917))
	}
	{
		methodName := StringNameFromStr("set_as_audio_listener_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnSetAsAudioListener2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_audio_listener_2d")
		defer methodName.Destroy()
		ptrsForViewport.fnIsAudioListener2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_mouse_position")
		defer methodName.Destroy()
		ptrsForViewport.fnGetMousePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("warp_mouse")
		defer methodName.Destroy()
		ptrsForViewport.fnWarpMouse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("update_mouse_cursor_state")
		defer methodName.Destroy()
		ptrsForViewport.fnUpdateMouseCursorState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("gui_get_drag_data")
		defer methodName.Destroy()
		ptrsForViewport.fnGuiGetDragData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1214101251))
	}
	{
		methodName := StringNameFromStr("gui_is_dragging")
		defer methodName.Destroy()
		ptrsForViewport.fnGuiIsDragging = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("gui_is_drag_successful")
		defer methodName.Destroy()
		ptrsForViewport.fnGuiIsDragSuccessful = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("gui_release_focus")
		defer methodName.Destroy()
		ptrsForViewport.fnGuiReleaseFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("gui_get_focus_owner")
		defer methodName.Destroy()
		ptrsForViewport.fnGuiGetFocusOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2783021301))
	}
	{
		methodName := StringNameFromStr("set_disable_input")
		defer methodName.Destroy()
		ptrsForViewport.fnSetDisableInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_input_disabled")
		defer methodName.Destroy()
		ptrsForViewport.fnIsInputDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_positional_shadow_atlas_size")
		defer methodName.Destroy()
		ptrsForViewport.fnSetPositionalShadowAtlasSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_positional_shadow_atlas_size")
		defer methodName.Destroy()
		ptrsForViewport.fnGetPositionalShadowAtlasSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_positional_shadow_atlas_16_bits")
		defer methodName.Destroy()
		ptrsForViewport.fnSetPositionalShadowAtlas16Bits = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_positional_shadow_atlas_16_bits")
		defer methodName.Destroy()
		ptrsForViewport.fnGetPositionalShadowAtlas16Bits = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_snap_controls_to_pixels")
		defer methodName.Destroy()
		ptrsForViewport.fnSetSnapControlsToPixels = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_snap_controls_to_pixels_enabled")
		defer methodName.Destroy()
		ptrsForViewport.fnIsSnapControlsToPixelsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_snap_2d_transforms_to_pixel")
		defer methodName.Destroy()
		ptrsForViewport.fnSetSnap2DTransformsToPixel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_snap_2d_transforms_to_pixel_enabled")
		defer methodName.Destroy()
		ptrsForViewport.fnIsSnap2DTransformsToPixelEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_snap_2d_vertices_to_pixel")
		defer methodName.Destroy()
		ptrsForViewport.fnSetSnap2DVerticesToPixel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_snap_2d_vertices_to_pixel_enabled")
		defer methodName.Destroy()
		ptrsForViewport.fnIsSnap2DVerticesToPixelEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_positional_shadow_atlas_quadrant_subdiv")
		defer methodName.Destroy()
		ptrsForViewport.fnSetPositionalShadowAtlasQuadrantSubdiv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2596956071))
	}
	{
		methodName := StringNameFromStr("get_positional_shadow_atlas_quadrant_subdiv")
		defer methodName.Destroy()
		ptrsForViewport.fnGetPositionalShadowAtlasQuadrantSubdiv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2676778355))
	}
	{
		methodName := StringNameFromStr("set_input_as_handled")
		defer methodName.Destroy()
		ptrsForViewport.fnSetInputAsHandled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_input_handled")
		defer methodName.Destroy()
		ptrsForViewport.fnIsInputHandled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_handle_input_locally")
		defer methodName.Destroy()
		ptrsForViewport.fnSetHandleInputLocally = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_handling_input_locally")
		defer methodName.Destroy()
		ptrsForViewport.fnIsHandlingInputLocally = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_default_canvas_item_texture_filter")
		defer methodName.Destroy()
		ptrsForViewport.fnSetDefaultCanvasItemTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2815160100))
	}
	{
		methodName := StringNameFromStr("get_default_canvas_item_texture_filter")
		defer methodName.Destroy()
		ptrsForViewport.fnGetDefaultCanvasItemTextureFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 896601198))
	}
	{
		methodName := StringNameFromStr("set_embedding_subwindows")
		defer methodName.Destroy()
		ptrsForViewport.fnSetEmbeddingSubwindows = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_embedding_subwindows")
		defer methodName.Destroy()
		ptrsForViewport.fnIsEmbeddingSubwindows = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_embedded_subwindows")
		defer methodName.Destroy()
		ptrsForViewport.fnGetEmbeddedSubwindows = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_canvas_cull_mask")
		defer methodName.Destroy()
		ptrsForViewport.fnSetCanvasCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_canvas_cull_mask")
		defer methodName.Destroy()
		ptrsForViewport.fnGetCanvasCullMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_canvas_cull_mask_bit")
		defer methodName.Destroy()
		ptrsForViewport.fnSetCanvasCullMaskBit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_canvas_cull_mask_bit")
		defer methodName.Destroy()
		ptrsForViewport.fnGetCanvasCullMaskBit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_default_canvas_item_texture_repeat")
		defer methodName.Destroy()
		ptrsForViewport.fnSetDefaultCanvasItemTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1658513413))
	}
	{
		methodName := StringNameFromStr("get_default_canvas_item_texture_repeat")
		defer methodName.Destroy()
		ptrsForViewport.fnGetDefaultCanvasItemTextureRepeat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4049774160))
	}
	{
		methodName := StringNameFromStr("set_sdf_oversize")
		defer methodName.Destroy()
		ptrsForViewport.fnSetSdfOversize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2574159017))
	}
	{
		methodName := StringNameFromStr("get_sdf_oversize")
		defer methodName.Destroy()
		ptrsForViewport.fnGetSdfOversize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2631427510))
	}
	{
		methodName := StringNameFromStr("set_sdf_scale")
		defer methodName.Destroy()
		ptrsForViewport.fnSetSdfScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1402773951))
	}
	{
		methodName := StringNameFromStr("get_sdf_scale")
		defer methodName.Destroy()
		ptrsForViewport.fnGetSdfScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3162688184))
	}
	{
		methodName := StringNameFromStr("set_mesh_lod_threshold")
		defer methodName.Destroy()
		ptrsForViewport.fnSetMeshLodThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_mesh_lod_threshold")
		defer methodName.Destroy()
		ptrsForViewport.fnGetMeshLodThreshold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_world_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnSetWorld3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1400875337))
	}
	{
		methodName := StringNameFromStr("get_world_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnGetWorld3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 317588385))
	}
	{
		methodName := StringNameFromStr("find_world_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnFindWorld3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 317588385))
	}
	{
		methodName := StringNameFromStr("set_use_own_world_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnSetUseOwnWorld3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_own_world_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnIsUsingOwnWorld3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_camera_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnGetCamera3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2285090890))
	}
	{
		methodName := StringNameFromStr("set_as_audio_listener_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnSetAsAudioListener3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_audio_listener_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnIsAudioListener3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_disable_3d")
		defer methodName.Destroy()
		ptrsForViewport.fnSetDisable3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_3d_disabled")
		defer methodName.Destroy()
		ptrsForViewport.fnIs3DDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_xr")
		defer methodName.Destroy()
		ptrsForViewport.fnSetUseXr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_xr")
		defer methodName.Destroy()
		ptrsForViewport.fnIsUsingXr = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_scaling_3d_mode")
		defer methodName.Destroy()
		ptrsForViewport.fnSetScaling3DMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1531597597))
	}
	{
		methodName := StringNameFromStr("get_scaling_3d_mode")
		defer methodName.Destroy()
		ptrsForViewport.fnGetScaling3DMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2597660574))
	}
	{
		methodName := StringNameFromStr("set_scaling_3d_scale")
		defer methodName.Destroy()
		ptrsForViewport.fnSetScaling3DScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_scaling_3d_scale")
		defer methodName.Destroy()
		ptrsForViewport.fnGetScaling3DScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fsr_sharpness")
		defer methodName.Destroy()
		ptrsForViewport.fnSetFsrSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_fsr_sharpness")
		defer methodName.Destroy()
		ptrsForViewport.fnGetFsrSharpness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_texture_mipmap_bias")
		defer methodName.Destroy()
		ptrsForViewport.fnSetTextureMipmapBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_texture_mipmap_bias")
		defer methodName.Destroy()
		ptrsForViewport.fnGetTextureMipmapBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_vrs_mode")
		defer methodName.Destroy()
		ptrsForViewport.fnSetVrsMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2749867817))
	}
	{
		methodName := StringNameFromStr("get_vrs_mode")
		defer methodName.Destroy()
		ptrsForViewport.fnGetVrsMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 349660525))
	}
	{
		methodName := StringNameFromStr("set_vrs_texture")
		defer methodName.Destroy()
		ptrsForViewport.fnSetVrsTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_vrs_texture")
		defer methodName.Destroy()
		ptrsForViewport.fnGetVrsTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}

}

type Viewport struct {
	Node
}

func (me *Viewport) BaseClass() string {
	return "Viewport"
}

func NewViewport() *Viewport {
	str := StringNameFromStr("Viewport") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Viewport{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type ViewportPositionalShadowAtlasQuadrantSubdiv int

const (
	ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdivDisabled ViewportPositionalShadowAtlasQuadrantSubdiv = 0
	ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv1        ViewportPositionalShadowAtlasQuadrantSubdiv = 1
	ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv4        ViewportPositionalShadowAtlasQuadrantSubdiv = 2
	ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv16       ViewportPositionalShadowAtlasQuadrantSubdiv = 3
	ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv64       ViewportPositionalShadowAtlasQuadrantSubdiv = 4
	ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv256      ViewportPositionalShadowAtlasQuadrantSubdiv = 5
	ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdiv1024     ViewportPositionalShadowAtlasQuadrantSubdiv = 6
	ViewportPositionalShadowAtlasQuadrantSubdivShadowAtlasQuadrantSubdivMax      ViewportPositionalShadowAtlasQuadrantSubdiv = 7
)

type ViewportScaling3DMode int

const (
	ViewportScaling3DModeScaling3DModeBilinear ViewportScaling3DMode = 0
	ViewportScaling3DModeScaling3DModeFsr      ViewportScaling3DMode = 1
	ViewportScaling3DModeScaling3DModeFsr2     ViewportScaling3DMode = 2
	ViewportScaling3DModeScaling3DModeMax      ViewportScaling3DMode = 3
)

type ViewportMSAA int

const (
	ViewportMSAAMsaaDisabled ViewportMSAA = 0
	ViewportMSAAMsaa2X       ViewportMSAA = 1
	ViewportMSAAMsaa4X       ViewportMSAA = 2
	ViewportMSAAMsaa8X       ViewportMSAA = 3
	ViewportMSAAMsaaMax      ViewportMSAA = 4
)

type ViewportScreenSpaceAA int

const (
	ViewportScreenSpaceAAScreenSpaceAaDisabled ViewportScreenSpaceAA = 0
	ViewportScreenSpaceAAScreenSpaceAaFxaa     ViewportScreenSpaceAA = 1
	ViewportScreenSpaceAAScreenSpaceAaMax      ViewportScreenSpaceAA = 2
)

type ViewportRenderInfo int

const (
	ViewportRenderInfoRenderInfoObjectsInFrame    ViewportRenderInfo = 0
	ViewportRenderInfoRenderInfoPrimitivesInFrame ViewportRenderInfo = 1
	ViewportRenderInfoRenderInfoDrawCallsInFrame  ViewportRenderInfo = 2
	ViewportRenderInfoRenderInfoMax               ViewportRenderInfo = 3
)

type ViewportRenderInfoType int

const (
	ViewportRenderInfoTypeRenderInfoTypeVisible ViewportRenderInfoType = 0
	ViewportRenderInfoTypeRenderInfoTypeShadow  ViewportRenderInfoType = 1
	ViewportRenderInfoTypeRenderInfoTypeMax     ViewportRenderInfoType = 2
)

type ViewportDebugDraw int

const (
	ViewportDebugDrawDebugDrawDisabled                ViewportDebugDraw = 0
	ViewportDebugDrawDebugDrawUnshaded                ViewportDebugDraw = 1
	ViewportDebugDrawDebugDrawLighting                ViewportDebugDraw = 2
	ViewportDebugDrawDebugDrawOverdraw                ViewportDebugDraw = 3
	ViewportDebugDrawDebugDrawWireframe               ViewportDebugDraw = 4
	ViewportDebugDrawDebugDrawNormalBuffer            ViewportDebugDraw = 5
	ViewportDebugDrawDebugDrawVoxelGiAlbedo           ViewportDebugDraw = 6
	ViewportDebugDrawDebugDrawVoxelGiLighting         ViewportDebugDraw = 7
	ViewportDebugDrawDebugDrawVoxelGiEmission         ViewportDebugDraw = 8
	ViewportDebugDrawDebugDrawShadowAtlas             ViewportDebugDraw = 9
	ViewportDebugDrawDebugDrawDirectionalShadowAtlas  ViewportDebugDraw = 10
	ViewportDebugDrawDebugDrawSceneLuminance          ViewportDebugDraw = 11
	ViewportDebugDrawDebugDrawSsao                    ViewportDebugDraw = 12
	ViewportDebugDrawDebugDrawSsil                    ViewportDebugDraw = 13
	ViewportDebugDrawDebugDrawPssmSplits              ViewportDebugDraw = 14
	ViewportDebugDrawDebugDrawDecalAtlas              ViewportDebugDraw = 15
	ViewportDebugDrawDebugDrawSdfgi                   ViewportDebugDraw = 16
	ViewportDebugDrawDebugDrawSdfgiProbes             ViewportDebugDraw = 17
	ViewportDebugDrawDebugDrawGiBuffer                ViewportDebugDraw = 18
	ViewportDebugDrawDebugDrawDisableLod              ViewportDebugDraw = 19
	ViewportDebugDrawDebugDrawClusterOmniLights       ViewportDebugDraw = 20
	ViewportDebugDrawDebugDrawClusterSpotLights       ViewportDebugDraw = 21
	ViewportDebugDrawDebugDrawClusterDecals           ViewportDebugDraw = 22
	ViewportDebugDrawDebugDrawClusterReflectionProbes ViewportDebugDraw = 23
	ViewportDebugDrawDebugDrawOccluders               ViewportDebugDraw = 24
	ViewportDebugDrawDebugDrawMotionVectors           ViewportDebugDraw = 25
	ViewportDebugDrawDebugDrawInternalBuffer          ViewportDebugDraw = 26
)

type ViewportDefaultCanvasItemTextureFilter int

const (
	ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterNearest            ViewportDefaultCanvasItemTextureFilter = 0
	ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterLinear             ViewportDefaultCanvasItemTextureFilter = 1
	ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterLinearWithMipmaps  ViewportDefaultCanvasItemTextureFilter = 2
	ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterNearestWithMipmaps ViewportDefaultCanvasItemTextureFilter = 3
	ViewportDefaultCanvasItemTextureFilterDefaultCanvasItemTextureFilterMax                ViewportDefaultCanvasItemTextureFilter = 4
)

type ViewportDefaultCanvasItemTextureRepeat int

const (
	ViewportDefaultCanvasItemTextureRepeatDefaultCanvasItemTextureRepeatDisabled ViewportDefaultCanvasItemTextureRepeat = 0
	ViewportDefaultCanvasItemTextureRepeatDefaultCanvasItemTextureRepeatEnabled  ViewportDefaultCanvasItemTextureRepeat = 1
	ViewportDefaultCanvasItemTextureRepeatDefaultCanvasItemTextureRepeatMirror   ViewportDefaultCanvasItemTextureRepeat = 2
	ViewportDefaultCanvasItemTextureRepeatDefaultCanvasItemTextureRepeatMax      ViewportDefaultCanvasItemTextureRepeat = 3
)

type ViewportSDFOversize int

const (
	ViewportSDFOversizeSdfOversize100Percent ViewportSDFOversize = 0
	ViewportSDFOversizeSdfOversize120Percent ViewportSDFOversize = 1
	ViewportSDFOversizeSdfOversize150Percent ViewportSDFOversize = 2
	ViewportSDFOversizeSdfOversize200Percent ViewportSDFOversize = 3
	ViewportSDFOversizeSdfOversizeMax        ViewportSDFOversize = 4
)

type ViewportSDFScale int

const (
	ViewportSDFScaleSdfScale100Percent ViewportSDFScale = 0
	ViewportSDFScaleSdfScale50Percent  ViewportSDFScale = 1
	ViewportSDFScaleSdfScale25Percent  ViewportSDFScale = 2
	ViewportSDFScaleSdfScaleMax        ViewportSDFScale = 3
)

type ViewportVRSMode int

const (
	ViewportVRSModeVrsDisabled ViewportVRSMode = 0
	ViewportVRSModeVrsTexture  ViewportVRSMode = 1
	ViewportVRSModeVrsXr       ViewportVRSMode = 2
	ViewportVRSModeVrsMax      ViewportVRSMode = 3
)

func (me *Viewport) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Viewport) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Viewport) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Viewport) SetWorld2D(world_2d World2D) {
	cargs := []gdc.ConstTypePtr{world_2d.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetWorld2D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetWorld2D() World2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewWorld2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetWorld2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) FindWorld2D() World2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewWorld2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnFindWorld2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) SetCanvasTransform(xform Transform2D) {
	cargs := []gdc.ConstTypePtr{xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetCanvasTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetCanvasTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetCanvasTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) SetGlobalCanvasTransform(xform Transform2D) {
	cargs := []gdc.ConstTypePtr{xform.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetGlobalCanvasTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetGlobalCanvasTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetGlobalCanvasTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) GetFinalTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetFinalTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) GetScreenTransform() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetScreenTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) GetVisibleRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetVisibleRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) SetTransparentBackground(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetTransparentBackground), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) HasTransparentBackground() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnHasTransparentBackground), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetUseHdr2D(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetUseHdr2D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsUsingHdr2D() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsUsingHdr2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetMsaa2D(msaa ViewportMSAA) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msaa)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetMsaa2D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetMsaa2D() ViewportMSAA {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportMSAA

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetMsaa2D), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetMsaa3D(msaa ViewportMSAA) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&msaa)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetMsaa3D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetMsaa3D() ViewportMSAA {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportMSAA

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetMsaa3D), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetScreenSpaceAa(screen_space_aa ViewportScreenSpaceAA) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen_space_aa)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetScreenSpaceAa), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetScreenSpaceAa() ViewportScreenSpaceAA {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportScreenSpaceAA

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetScreenSpaceAa), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetUseTaa(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetUseTaa), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsUsingTaa() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsUsingTaa), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetUseDebanding(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetUseDebanding), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsUsingDebanding() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsUsingDebanding), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetUseOcclusionCulling(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetUseOcclusionCulling), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsUsingOcclusionCulling() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsUsingOcclusionCulling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetDebugDraw(debug_draw ViewportDebugDraw) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&debug_draw)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetDebugDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetDebugDraw() ViewportDebugDraw {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportDebugDraw

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetDebugDraw), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) GetRenderInfo(type_ ViewportRenderInfoType, info ViewportRenderInfo) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&info)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&type_)
	pinner.Pin(&info)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetRenderInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) GetTexture() ViewportTexture {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewViewportTexture()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) SetPhysicsObjectPicking(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetPhysicsObjectPicking), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetPhysicsObjectPicking() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetPhysicsObjectPicking), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetPhysicsObjectPickingSort(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetPhysicsObjectPickingSort), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetPhysicsObjectPickingSort() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetPhysicsObjectPickingSort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) GetViewportRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetViewportRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) PushTextInput(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnPushTextInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) PushInput(event InputEvent, in_local_coords bool) {
	cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), gdc.ConstTypePtr(&in_local_coords)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnPushInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) PushUnhandledInput(event InputEvent, in_local_coords bool) {
	cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), gdc.ConstTypePtr(&in_local_coords)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnPushUnhandledInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetCamera2D() Camera2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCamera2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetCamera2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) SetAsAudioListener2D(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetAsAudioListener2D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsAudioListener2D() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsAudioListener2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) GetMousePosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetMousePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) WarpMouse(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnWarpMouse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) UpdateMouseCursorState() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnUpdateMouseCursorState), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GuiGetDragData() Variant {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGuiGetDragData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) GuiIsDragging() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGuiIsDragging), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) GuiIsDragSuccessful() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGuiIsDragSuccessful), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) GuiReleaseFocus() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGuiReleaseFocus), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GuiGetFocusOwner() Control {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewControl()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGuiGetFocusOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) SetDisableInput(disable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetDisableInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsInputDisabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsInputDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetPositionalShadowAtlasSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetPositionalShadowAtlasSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetPositionalShadowAtlasSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetPositionalShadowAtlasSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetPositionalShadowAtlas16Bits(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetPositionalShadowAtlas16Bits), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetPositionalShadowAtlas16Bits() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetPositionalShadowAtlas16Bits), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetSnapControlsToPixels(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetSnapControlsToPixels), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsSnapControlsToPixelsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsSnapControlsToPixelsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetSnap2DTransformsToPixel(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetSnap2DTransformsToPixel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsSnap2DTransformsToPixelEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsSnap2DTransformsToPixelEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetSnap2DVerticesToPixel(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetSnap2DVerticesToPixel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsSnap2DVerticesToPixelEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsSnap2DVerticesToPixelEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetPositionalShadowAtlasQuadrantSubdiv(quadrant int64, subdiv ViewportPositionalShadowAtlasQuadrantSubdiv) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quadrant), gdc.ConstTypePtr(&subdiv)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetPositionalShadowAtlasQuadrantSubdiv), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetPositionalShadowAtlasQuadrantSubdiv(quadrant int64) ViewportPositionalShadowAtlasQuadrantSubdiv {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&quadrant)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportPositionalShadowAtlasQuadrantSubdiv
	pinner.Pin(&quadrant)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetPositionalShadowAtlasQuadrantSubdiv), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetInputAsHandled() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetInputAsHandled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsInputHandled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsInputHandled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetHandleInputLocally(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetHandleInputLocally), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsHandlingInputLocally() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsHandlingInputLocally), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetDefaultCanvasItemTextureFilter(mode ViewportDefaultCanvasItemTextureFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetDefaultCanvasItemTextureFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetDefaultCanvasItemTextureFilter() ViewportDefaultCanvasItemTextureFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportDefaultCanvasItemTextureFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetDefaultCanvasItemTextureFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetEmbeddingSubwindows(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetEmbeddingSubwindows), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsEmbeddingSubwindows() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsEmbeddingSubwindows), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) GetEmbeddedSubwindows() []Window {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetEmbeddedSubwindows), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Window](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Viewport) SetCanvasCullMask(mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetCanvasCullMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetCanvasCullMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetCanvasCullMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetCanvasCullMaskBit(layer int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetCanvasCullMaskBit), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetCanvasCullMaskBit(layer int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetCanvasCullMaskBit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetDefaultCanvasItemTextureRepeat(mode ViewportDefaultCanvasItemTextureRepeat) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetDefaultCanvasItemTextureRepeat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetDefaultCanvasItemTextureRepeat() ViewportDefaultCanvasItemTextureRepeat {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportDefaultCanvasItemTextureRepeat

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetDefaultCanvasItemTextureRepeat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetSdfOversize(oversize ViewportSDFOversize) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversize)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetSdfOversize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetSdfOversize() ViewportSDFOversize {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportSDFOversize

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetSdfOversize), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetSdfScale(scale ViewportSDFScale) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetSdfScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetSdfScale() ViewportSDFScale {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportSDFScale

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetSdfScale), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetMeshLodThreshold(pixels float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetMeshLodThreshold), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetMeshLodThreshold() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetMeshLodThreshold), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetWorld3D(world_3d World3D) {
	cargs := []gdc.ConstTypePtr{world_3d.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetWorld3D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetWorld3D() World3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewWorld3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetWorld3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) FindWorld3D() World3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewWorld3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnFindWorld3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) SetUseOwnWorld3D(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetUseOwnWorld3D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsUsingOwnWorld3D() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsUsingOwnWorld3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) GetCamera3D() Camera3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCamera3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetCamera3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Viewport) SetAsAudioListener3D(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetAsAudioListener3D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsAudioListener3D() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsAudioListener3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetDisable3D(disable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetDisable3D), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) Is3DDisabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIs3DDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetUseXr(use bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetUseXr), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) IsUsingXr() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnIsUsingXr), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetScaling3DMode(scaling_3d_mode ViewportScaling3DMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scaling_3d_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetScaling3DMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetScaling3DMode() ViewportScaling3DMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportScaling3DMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetScaling3DMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetScaling3DScale(scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetScaling3DScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetScaling3DScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetScaling3DScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetFsrSharpness(fsr_sharpness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fsr_sharpness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetFsrSharpness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetFsrSharpness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetFsrSharpness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetTextureMipmapBias(texture_mipmap_bias float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_mipmap_bias)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetTextureMipmapBias), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetTextureMipmapBias() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetTextureMipmapBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Viewport) SetVrsMode(mode ViewportVRSMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetVrsMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetVrsMode() ViewportVRSMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ViewportVRSMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetVrsMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Viewport) SetVrsTexture(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnSetVrsTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Viewport) GetVrsTexture() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForViewport.fnGetVrsTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ViewportSizeChangedSignalFn func()

func (me *Viewport) ConnectSizeChanged(subs SignalSubscribers, fn ViewportSizeChangedSignalFn) {
	sig := StringNameFromStr("size_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Viewport) DisconnectSizeChanged(subs SignalSubscribers, fn ViewportSizeChangedSignalFn) {
	sig := StringNameFromStr("size_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ViewportGuiFocusChangedSignalFn func(node Control)

func (me *Viewport) ConnectGuiFocusChanged(subs SignalSubscribers, fn ViewportGuiFocusChangedSignalFn) {
	sig := StringNameFromStr("gui_focus_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Viewport) DisconnectGuiFocusChanged(subs SignalSubscribers, fn ViewportGuiFocusChangedSignalFn) {
	sig := StringNameFromStr("gui_focus_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
