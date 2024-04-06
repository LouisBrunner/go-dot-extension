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

type ptrsForWindowList struct {
	fnXGetContentsMinimumSize       gdc.MethodBindPtr
	fnSetTitle                      gdc.MethodBindPtr
	fnGetTitle                      gdc.MethodBindPtr
	fnGetWindowId                   gdc.MethodBindPtr
	fnSetInitialPosition            gdc.MethodBindPtr
	fnGetInitialPosition            gdc.MethodBindPtr
	fnSetCurrentScreen              gdc.MethodBindPtr
	fnGetCurrentScreen              gdc.MethodBindPtr
	fnSetPosition                   gdc.MethodBindPtr
	fnGetPosition                   gdc.MethodBindPtr
	fnMoveToCenter                  gdc.MethodBindPtr
	fnSetSize                       gdc.MethodBindPtr
	fnGetSize                       gdc.MethodBindPtr
	fnResetSize                     gdc.MethodBindPtr
	fnGetPositionWithDecorations    gdc.MethodBindPtr
	fnGetSizeWithDecorations        gdc.MethodBindPtr
	fnSetMaxSize                    gdc.MethodBindPtr
	fnGetMaxSize                    gdc.MethodBindPtr
	fnSetMinSize                    gdc.MethodBindPtr
	fnGetMinSize                    gdc.MethodBindPtr
	fnSetMode                       gdc.MethodBindPtr
	fnGetMode                       gdc.MethodBindPtr
	fnSetFlag                       gdc.MethodBindPtr
	fnGetFlag                       gdc.MethodBindPtr
	fnIsMaximizeAllowed             gdc.MethodBindPtr
	fnRequestAttention              gdc.MethodBindPtr
	fnMoveToForeground              gdc.MethodBindPtr
	fnSetVisible                    gdc.MethodBindPtr
	fnIsVisible                     gdc.MethodBindPtr
	fnHide                          gdc.MethodBindPtr
	fnShow                          gdc.MethodBindPtr
	fnSetTransient                  gdc.MethodBindPtr
	fnIsTransient                   gdc.MethodBindPtr
	fnSetExclusive                  gdc.MethodBindPtr
	fnIsExclusive                   gdc.MethodBindPtr
	fnSetUnparentWhenInvisible      gdc.MethodBindPtr
	fnCanDraw                       gdc.MethodBindPtr
	fnHasFocus                      gdc.MethodBindPtr
	fnGrabFocus                     gdc.MethodBindPtr
	fnSetImeActive                  gdc.MethodBindPtr
	fnSetImePosition                gdc.MethodBindPtr
	fnIsEmbedded                    gdc.MethodBindPtr
	fnGetContentsMinimumSize        gdc.MethodBindPtr
	fnSetContentScaleSize           gdc.MethodBindPtr
	fnGetContentScaleSize           gdc.MethodBindPtr
	fnSetContentScaleMode           gdc.MethodBindPtr
	fnGetContentScaleMode           gdc.MethodBindPtr
	fnSetContentScaleAspect         gdc.MethodBindPtr
	fnGetContentScaleAspect         gdc.MethodBindPtr
	fnSetContentScaleStretch        gdc.MethodBindPtr
	fnGetContentScaleStretch        gdc.MethodBindPtr
	fnSetKeepTitleVisible           gdc.MethodBindPtr
	fnGetKeepTitleVisible           gdc.MethodBindPtr
	fnSetContentScaleFactor         gdc.MethodBindPtr
	fnGetContentScaleFactor         gdc.MethodBindPtr
	fnSetUseFontOversampling        gdc.MethodBindPtr
	fnIsUsingFontOversampling       gdc.MethodBindPtr
	fnSetMousePassthroughPolygon    gdc.MethodBindPtr
	fnGetMousePassthroughPolygon    gdc.MethodBindPtr
	fnSetWrapControls               gdc.MethodBindPtr
	fnIsWrappingControls            gdc.MethodBindPtr
	fnChildControlsChanged          gdc.MethodBindPtr
	fnSetTheme                      gdc.MethodBindPtr
	fnGetTheme                      gdc.MethodBindPtr
	fnSetThemeTypeVariation         gdc.MethodBindPtr
	fnGetThemeTypeVariation         gdc.MethodBindPtr
	fnBeginBulkThemeOverride        gdc.MethodBindPtr
	fnEndBulkThemeOverride          gdc.MethodBindPtr
	fnAddThemeIconOverride          gdc.MethodBindPtr
	fnAddThemeStyleboxOverride      gdc.MethodBindPtr
	fnAddThemeFontOverride          gdc.MethodBindPtr
	fnAddThemeFontSizeOverride      gdc.MethodBindPtr
	fnAddThemeColorOverride         gdc.MethodBindPtr
	fnAddThemeConstantOverride      gdc.MethodBindPtr
	fnRemoveThemeIconOverride       gdc.MethodBindPtr
	fnRemoveThemeStyleboxOverride   gdc.MethodBindPtr
	fnRemoveThemeFontOverride       gdc.MethodBindPtr
	fnRemoveThemeFontSizeOverride   gdc.MethodBindPtr
	fnRemoveThemeColorOverride      gdc.MethodBindPtr
	fnRemoveThemeConstantOverride   gdc.MethodBindPtr
	fnGetThemeIcon                  gdc.MethodBindPtr
	fnGetThemeStylebox              gdc.MethodBindPtr
	fnGetThemeFont                  gdc.MethodBindPtr
	fnGetThemeFontSize              gdc.MethodBindPtr
	fnGetThemeColor                 gdc.MethodBindPtr
	fnGetThemeConstant              gdc.MethodBindPtr
	fnHasThemeIconOverride          gdc.MethodBindPtr
	fnHasThemeStyleboxOverride      gdc.MethodBindPtr
	fnHasThemeFontOverride          gdc.MethodBindPtr
	fnHasThemeFontSizeOverride      gdc.MethodBindPtr
	fnHasThemeColorOverride         gdc.MethodBindPtr
	fnHasThemeConstantOverride      gdc.MethodBindPtr
	fnHasThemeIcon                  gdc.MethodBindPtr
	fnHasThemeStylebox              gdc.MethodBindPtr
	fnHasThemeFont                  gdc.MethodBindPtr
	fnHasThemeFontSize              gdc.MethodBindPtr
	fnHasThemeColor                 gdc.MethodBindPtr
	fnHasThemeConstant              gdc.MethodBindPtr
	fnGetThemeDefaultBaseScale      gdc.MethodBindPtr
	fnGetThemeDefaultFont           gdc.MethodBindPtr
	fnGetThemeDefaultFontSize       gdc.MethodBindPtr
	fnSetLayoutDirection            gdc.MethodBindPtr
	fnGetLayoutDirection            gdc.MethodBindPtr
	fnIsLayoutRtl                   gdc.MethodBindPtr
	fnSetAutoTranslate              gdc.MethodBindPtr
	fnIsAutoTranslating             gdc.MethodBindPtr
	fnPopup                         gdc.MethodBindPtr
	fnPopupOnParent                 gdc.MethodBindPtr
	fnPopupCentered                 gdc.MethodBindPtr
	fnPopupCenteredRatio            gdc.MethodBindPtr
	fnPopupCenteredClamped          gdc.MethodBindPtr
	fnPopupExclusive                gdc.MethodBindPtr
	fnPopupExclusiveOnParent        gdc.MethodBindPtr
	fnPopupExclusiveCentered        gdc.MethodBindPtr
	fnPopupExclusiveCenteredRatio   gdc.MethodBindPtr
	fnPopupExclusiveCenteredClamped gdc.MethodBindPtr
}

var ptrsForWindow ptrsForWindowList

func initWindowPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Window")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_title")
		defer methodName.Destroy()
		ptrsForWindow.fnSetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_title")
		defer methodName.Destroy()
		ptrsForWindow.fnGetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_window_id")
		defer methodName.Destroy()
		ptrsForWindow.fnGetWindowId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_initial_position")
		defer methodName.Destroy()
		ptrsForWindow.fnSetInitialPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4084468099))
	}
	{
		methodName := StringNameFromStr("get_initial_position")
		defer methodName.Destroy()
		ptrsForWindow.fnGetInitialPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4294066647))
	}
	{
		methodName := StringNameFromStr("set_current_screen")
		defer methodName.Destroy()
		ptrsForWindow.fnSetCurrentScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_current_screen")
		defer methodName.Destroy()
		ptrsForWindow.fnGetCurrentScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForWindow.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForWindow.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("move_to_center")
		defer methodName.Destroy()
		ptrsForWindow.fnMoveToCenter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForWindow.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForWindow.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("reset_size")
		defer methodName.Destroy()
		ptrsForWindow.fnResetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_position_with_decorations")
		defer methodName.Destroy()
		ptrsForWindow.fnGetPositionWithDecorations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("get_size_with_decorations")
		defer methodName.Destroy()
		ptrsForWindow.fnGetSizeWithDecorations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("set_max_size")
		defer methodName.Destroy()
		ptrsForWindow.fnSetMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_max_size")
		defer methodName.Destroy()
		ptrsForWindow.fnGetMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("set_min_size")
		defer methodName.Destroy()
		ptrsForWindow.fnSetMinSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_min_size")
		defer methodName.Destroy()
		ptrsForWindow.fnGetMinSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("set_mode")
		defer methodName.Destroy()
		ptrsForWindow.fnSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3095236531))
	}
	{
		methodName := StringNameFromStr("get_mode")
		defer methodName.Destroy()
		ptrsForWindow.fnGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2566346114))
	}
	{
		methodName := StringNameFromStr("set_flag")
		defer methodName.Destroy()
		ptrsForWindow.fnSetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3426449779))
	}
	{
		methodName := StringNameFromStr("get_flag")
		defer methodName.Destroy()
		ptrsForWindow.fnGetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3062752289))
	}
	{
		methodName := StringNameFromStr("is_maximize_allowed")
		defer methodName.Destroy()
		ptrsForWindow.fnIsMaximizeAllowed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("request_attention")
		defer methodName.Destroy()
		ptrsForWindow.fnRequestAttention = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("move_to_foreground")
		defer methodName.Destroy()
		ptrsForWindow.fnMoveToForeground = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_visible")
		defer methodName.Destroy()
		ptrsForWindow.fnSetVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_visible")
		defer methodName.Destroy()
		ptrsForWindow.fnIsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("hide")
		defer methodName.Destroy()
		ptrsForWindow.fnHide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("show")
		defer methodName.Destroy()
		ptrsForWindow.fnShow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_transient")
		defer methodName.Destroy()
		ptrsForWindow.fnSetTransient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_transient")
		defer methodName.Destroy()
		ptrsForWindow.fnIsTransient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_exclusive")
		defer methodName.Destroy()
		ptrsForWindow.fnSetExclusive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_exclusive")
		defer methodName.Destroy()
		ptrsForWindow.fnIsExclusive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_unparent_when_invisible")
		defer methodName.Destroy()
		ptrsForWindow.fnSetUnparentWhenInvisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("can_draw")
		defer methodName.Destroy()
		ptrsForWindow.fnCanDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("has_focus")
		defer methodName.Destroy()
		ptrsForWindow.fnHasFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("grab_focus")
		defer methodName.Destroy()
		ptrsForWindow.fnGrabFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_ime_active")
		defer methodName.Destroy()
		ptrsForWindow.fnSetImeActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_ime_position")
		defer methodName.Destroy()
		ptrsForWindow.fnSetImePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("is_embedded")
		defer methodName.Destroy()
		ptrsForWindow.fnIsEmbedded = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_contents_minimum_size")
		defer methodName.Destroy()
		ptrsForWindow.fnGetContentsMinimumSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_content_scale_size")
		defer methodName.Destroy()
		ptrsForWindow.fnSetContentScaleSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_content_scale_size")
		defer methodName.Destroy()
		ptrsForWindow.fnGetContentScaleSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("set_content_scale_mode")
		defer methodName.Destroy()
		ptrsForWindow.fnSetContentScaleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2937716473))
	}
	{
		methodName := StringNameFromStr("get_content_scale_mode")
		defer methodName.Destroy()
		ptrsForWindow.fnGetContentScaleMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 161585230))
	}
	{
		methodName := StringNameFromStr("set_content_scale_aspect")
		defer methodName.Destroy()
		ptrsForWindow.fnSetContentScaleAspect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2370399418))
	}
	{
		methodName := StringNameFromStr("get_content_scale_aspect")
		defer methodName.Destroy()
		ptrsForWindow.fnGetContentScaleAspect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4158790715))
	}
	{
		methodName := StringNameFromStr("set_content_scale_stretch")
		defer methodName.Destroy()
		ptrsForWindow.fnSetContentScaleStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 349355940))
	}
	{
		methodName := StringNameFromStr("get_content_scale_stretch")
		defer methodName.Destroy()
		ptrsForWindow.fnGetContentScaleStretch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 536857316))
	}
	{
		methodName := StringNameFromStr("set_keep_title_visible")
		defer methodName.Destroy()
		ptrsForWindow.fnSetKeepTitleVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_keep_title_visible")
		defer methodName.Destroy()
		ptrsForWindow.fnGetKeepTitleVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_content_scale_factor")
		defer methodName.Destroy()
		ptrsForWindow.fnSetContentScaleFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_content_scale_factor")
		defer methodName.Destroy()
		ptrsForWindow.fnGetContentScaleFactor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_use_font_oversampling")
		defer methodName.Destroy()
		ptrsForWindow.fnSetUseFontOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_font_oversampling")
		defer methodName.Destroy()
		ptrsForWindow.fnIsUsingFontOversampling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_mouse_passthrough_polygon")
		defer methodName.Destroy()
		ptrsForWindow.fnSetMousePassthroughPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
	}
	{
		methodName := StringNameFromStr("get_mouse_passthrough_polygon")
		defer methodName.Destroy()
		ptrsForWindow.fnGetMousePassthroughPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("set_wrap_controls")
		defer methodName.Destroy()
		ptrsForWindow.fnSetWrapControls = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_wrapping_controls")
		defer methodName.Destroy()
		ptrsForWindow.fnIsWrappingControls = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("child_controls_changed")
		defer methodName.Destroy()
		ptrsForWindow.fnChildControlsChanged = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_theme")
		defer methodName.Destroy()
		ptrsForWindow.fnSetTheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2326690814))
	}
	{
		methodName := StringNameFromStr("get_theme")
		defer methodName.Destroy()
		ptrsForWindow.fnGetTheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3846893731))
	}
	{
		methodName := StringNameFromStr("set_theme_type_variation")
		defer methodName.Destroy()
		ptrsForWindow.fnSetThemeTypeVariation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_theme_type_variation")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeTypeVariation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("begin_bulk_theme_override")
		defer methodName.Destroy()
		ptrsForWindow.fnBeginBulkThemeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("end_bulk_theme_override")
		defer methodName.Destroy()
		ptrsForWindow.fnEndBulkThemeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("add_theme_icon_override")
		defer methodName.Destroy()
		ptrsForWindow.fnAddThemeIconOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1373065600))
	}
	{
		methodName := StringNameFromStr("add_theme_stylebox_override")
		defer methodName.Destroy()
		ptrsForWindow.fnAddThemeStyleboxOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4188838905))
	}
	{
		methodName := StringNameFromStr("add_theme_font_override")
		defer methodName.Destroy()
		ptrsForWindow.fnAddThemeFontOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3518018674))
	}
	{
		methodName := StringNameFromStr("add_theme_font_size_override")
		defer methodName.Destroy()
		ptrsForWindow.fnAddThemeFontSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2415702435))
	}
	{
		methodName := StringNameFromStr("add_theme_color_override")
		defer methodName.Destroy()
		ptrsForWindow.fnAddThemeColorOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4260178595))
	}
	{
		methodName := StringNameFromStr("add_theme_constant_override")
		defer methodName.Destroy()
		ptrsForWindow.fnAddThemeConstantOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2415702435))
	}
	{
		methodName := StringNameFromStr("remove_theme_icon_override")
		defer methodName.Destroy()
		ptrsForWindow.fnRemoveThemeIconOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_stylebox_override")
		defer methodName.Destroy()
		ptrsForWindow.fnRemoveThemeStyleboxOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_font_override")
		defer methodName.Destroy()
		ptrsForWindow.fnRemoveThemeFontOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_font_size_override")
		defer methodName.Destroy()
		ptrsForWindow.fnRemoveThemeFontSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_color_override")
		defer methodName.Destroy()
		ptrsForWindow.fnRemoveThemeColorOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_constant_override")
		defer methodName.Destroy()
		ptrsForWindow.fnRemoveThemeConstantOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_theme_icon")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3163973443))
	}
	{
		methodName := StringNameFromStr("get_theme_stylebox")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 604739069))
	}
	{
		methodName := StringNameFromStr("get_theme_font")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2826986490))
	}
	{
		methodName := StringNameFromStr("get_theme_font_size")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1327056374))
	}
	{
		methodName := StringNameFromStr("get_theme_color")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2798751242))
	}
	{
		methodName := StringNameFromStr("get_theme_constant")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1327056374))
	}
	{
		methodName := StringNameFromStr("has_theme_icon_override")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeIconOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_stylebox_override")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeStyleboxOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_font_override")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeFontOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_font_size_override")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeFontSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_color_override")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeColorOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_constant_override")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeConstantOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_icon")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_stylebox")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_font")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_font_size")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_color")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_constant")
		defer methodName.Destroy()
		ptrsForWindow.fnHasThemeConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("get_theme_default_base_scale")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeDefaultBaseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_theme_default_font")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeDefaultFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229501585))
	}
	{
		methodName := StringNameFromStr("get_theme_default_font_size")
		defer methodName.Destroy()
		ptrsForWindow.fnGetThemeDefaultFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_layout_direction")
		defer methodName.Destroy()
		ptrsForWindow.fnSetLayoutDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3094704184))
	}
	{
		methodName := StringNameFromStr("get_layout_direction")
		defer methodName.Destroy()
		ptrsForWindow.fnGetLayoutDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3909617982))
	}
	{
		methodName := StringNameFromStr("is_layout_rtl")
		defer methodName.Destroy()
		ptrsForWindow.fnIsLayoutRtl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_auto_translate")
		defer methodName.Destroy()
		ptrsForWindow.fnSetAutoTranslate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_auto_translating")
		defer methodName.Destroy()
		ptrsForWindow.fnIsAutoTranslating = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("popup")
		defer methodName.Destroy()
		ptrsForWindow.fnPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1680304321))
	}
	{
		methodName := StringNameFromStr("popup_on_parent")
		defer methodName.Destroy()
		ptrsForWindow.fnPopupOnParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1763793166))
	}
	{
		methodName := StringNameFromStr("popup_centered")
		defer methodName.Destroy()
		ptrsForWindow.fnPopupCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3447975422))
	}
	{
		methodName := StringNameFromStr("popup_centered_ratio")
		defer methodName.Destroy()
		ptrsForWindow.fnPopupCenteredRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1014814997))
	}
	{
		methodName := StringNameFromStr("popup_centered_clamped")
		defer methodName.Destroy()
		ptrsForWindow.fnPopupCenteredClamped = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2613752477))
	}
	{
		methodName := StringNameFromStr("popup_exclusive")
		defer methodName.Destroy()
		ptrsForWindow.fnPopupExclusive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2134721627))
	}
	{
		methodName := StringNameFromStr("popup_exclusive_on_parent")
		defer methodName.Destroy()
		ptrsForWindow.fnPopupExclusiveOnParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2344671043))
	}
	{
		methodName := StringNameFromStr("popup_exclusive_centered")
		defer methodName.Destroy()
		ptrsForWindow.fnPopupExclusiveCentered = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3357594017))
	}
	{
		methodName := StringNameFromStr("popup_exclusive_centered_ratio")
		defer methodName.Destroy()
		ptrsForWindow.fnPopupExclusiveCenteredRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2284776287))
	}
	{
		methodName := StringNameFromStr("popup_exclusive_centered_clamped")
		defer methodName.Destroy()
		ptrsForWindow.fnPopupExclusiveCenteredClamped = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2612708785))
	}

}

type Window struct {
	Viewport
}

func (me *Window) BaseClass() string {
	return "Window"
}

func NewWindow() *Window {
	str := StringNameFromStr("Window") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Window{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	WindowNotificationVisibilityChanged = 30
	WindowNotificationThemeChanged      = 32
)

// Enums

type WindowMode int

const (
	WindowModeModeWindowed            WindowMode = 0
	WindowModeModeMinimized           WindowMode = 1
	WindowModeModeMaximized           WindowMode = 2
	WindowModeModeFullscreen          WindowMode = 3
	WindowModeModeExclusiveFullscreen WindowMode = 4
)

type WindowFlags int

const (
	WindowFlagsFlagResizeDisabled   WindowFlags = 0
	WindowFlagsFlagBorderless       WindowFlags = 1
	WindowFlagsFlagAlwaysOnTop      WindowFlags = 2
	WindowFlagsFlagTransparent      WindowFlags = 3
	WindowFlagsFlagNoFocus          WindowFlags = 4
	WindowFlagsFlagPopup            WindowFlags = 5
	WindowFlagsFlagExtendToTitle    WindowFlags = 6
	WindowFlagsFlagMousePassthrough WindowFlags = 7
	WindowFlagsFlagMax              WindowFlags = 8
)

type WindowContentScaleMode int

const (
	WindowContentScaleModeContentScaleModeDisabled    WindowContentScaleMode = 0
	WindowContentScaleModeContentScaleModeCanvasItems WindowContentScaleMode = 1
	WindowContentScaleModeContentScaleModeViewport    WindowContentScaleMode = 2
)

type WindowContentScaleAspect int

const (
	WindowContentScaleAspectContentScaleAspectIgnore     WindowContentScaleAspect = 0
	WindowContentScaleAspectContentScaleAspectKeep       WindowContentScaleAspect = 1
	WindowContentScaleAspectContentScaleAspectKeepWidth  WindowContentScaleAspect = 2
	WindowContentScaleAspectContentScaleAspectKeepHeight WindowContentScaleAspect = 3
	WindowContentScaleAspectContentScaleAspectExpand     WindowContentScaleAspect = 4
)

type WindowContentScaleStretch int

const (
	WindowContentScaleStretchContentScaleStretchFractional WindowContentScaleStretch = 0
	WindowContentScaleStretchContentScaleStretchInteger    WindowContentScaleStretch = 1
)

type WindowLayoutDirection int

const (
	WindowLayoutDirectionLayoutDirectionInherited WindowLayoutDirection = 0
	WindowLayoutDirectionLayoutDirectionLocale    WindowLayoutDirection = 1
	WindowLayoutDirectionLayoutDirectionLtr       WindowLayoutDirection = 2
	WindowLayoutDirectionLayoutDirectionRtl       WindowLayoutDirection = 3
)

type WindowWindowInitialPosition int

const (
	WindowWindowInitialPositionWindowInitialPositionAbsolute                      WindowWindowInitialPosition = 0
	WindowWindowInitialPositionWindowInitialPositionCenterPrimaryScreen           WindowWindowInitialPosition = 1
	WindowWindowInitialPositionWindowInitialPositionCenterMainWindowScreen        WindowWindowInitialPosition = 2
	WindowWindowInitialPositionWindowInitialPositionCenterOtherScreen             WindowWindowInitialPosition = 3
	WindowWindowInitialPositionWindowInitialPositionCenterScreenWithMouseFocus    WindowWindowInitialPosition = 4
	WindowWindowInitialPositionWindowInitialPositionCenterScreenWithKeyboardFocus WindowWindowInitialPosition = 5
)

func (me *Window) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Window) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Window) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Window) SetTitle(title String) {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetTitle() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) GetWindowId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetWindowId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) SetInitialPosition(initial_position WindowWindowInitialPosition) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&initial_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetInitialPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetInitialPosition() WindowWindowInitialPosition {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WindowWindowInitialPosition

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetInitialPosition), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Window) SetCurrentScreen(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetCurrentScreen), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetCurrentScreen() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetCurrentScreen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) SetPosition(position Vector2i) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetPosition() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) MoveToCenter() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnMoveToCenter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) SetSize(size Vector2i) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetSize() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) ResetSize() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnResetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetPositionWithDecorations() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetPositionWithDecorations), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) GetSizeWithDecorations() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetSizeWithDecorations), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) SetMaxSize(max_size Vector2i) {
	cargs := []gdc.ConstTypePtr{max_size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetMaxSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetMaxSize() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetMaxSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) SetMinSize(min_size Vector2i) {
	cargs := []gdc.ConstTypePtr{min_size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetMinSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetMinSize() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetMinSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) SetMode(mode WindowMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetMode() WindowMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WindowMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Window) SetFlag(flag WindowFlags, enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetFlag(flag WindowFlags) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) IsMaximizeAllowed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnIsMaximizeAllowed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) RequestAttention() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnRequestAttention), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) MoveToForeground() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnMoveToForeground), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) SetVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) IsVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnIsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) Hide() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHide), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) Show() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnShow), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) SetTransient(transient bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&transient)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetTransient), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) IsTransient() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnIsTransient), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) SetExclusive(exclusive bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclusive)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetExclusive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) IsExclusive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnIsExclusive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) SetUnparentWhenInvisible(unparent bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unparent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetUnparentWhenInvisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) CanDraw() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnCanDraw), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasFocus() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasFocus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) GrabFocus() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGrabFocus), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) SetImeActive(active bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetImeActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) SetImePosition(position Vector2i) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetImePosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) IsEmbedded() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnIsEmbedded), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) GetContentsMinimumSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetContentsMinimumSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) SetContentScaleSize(size Vector2i) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetContentScaleSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetContentScaleSize() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetContentScaleSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) SetContentScaleMode(mode WindowContentScaleMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetContentScaleMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetContentScaleMode() WindowContentScaleMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WindowContentScaleMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetContentScaleMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Window) SetContentScaleAspect(aspect WindowContentScaleAspect) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&aspect)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetContentScaleAspect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetContentScaleAspect() WindowContentScaleAspect {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WindowContentScaleAspect

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetContentScaleAspect), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Window) SetContentScaleStretch(stretch WindowContentScaleStretch) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetContentScaleStretch), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetContentScaleStretch() WindowContentScaleStretch {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WindowContentScaleStretch

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetContentScaleStretch), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Window) SetKeepTitleVisible(title_visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&title_visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetKeepTitleVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetKeepTitleVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetKeepTitleVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) SetContentScaleFactor(factor float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&factor)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetContentScaleFactor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetContentScaleFactor() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetContentScaleFactor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) SetUseFontOversampling(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetUseFontOversampling), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) IsUsingFontOversampling() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnIsUsingFontOversampling), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) SetMousePassthroughPolygon(polygon PackedVector2Array) {
	cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetMousePassthroughPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetMousePassthroughPolygon() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetMousePassthroughPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) SetWrapControls(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetWrapControls), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) IsWrappingControls() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnIsWrappingControls), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) ChildControlsChanged() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnChildControlsChanged), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) SetTheme(theme Theme) {
	cargs := []gdc.ConstTypePtr{theme.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetTheme), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetTheme() Theme {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTheme()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetTheme), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) SetThemeTypeVariation(theme_type StringName) {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetThemeTypeVariation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetThemeTypeVariation() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeTypeVariation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) BeginBulkThemeOverride() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnBeginBulkThemeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) EndBulkThemeOverride() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnEndBulkThemeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) AddThemeIconOverride(name StringName, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnAddThemeIconOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) AddThemeStyleboxOverride(name StringName, stylebox StyleBox) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), stylebox.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnAddThemeStyleboxOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) AddThemeFontOverride(name StringName, font Font) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), font.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnAddThemeFontOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) AddThemeFontSizeOverride(name StringName, font_size int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnAddThemeFontSizeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) AddThemeColorOverride(name StringName, color Color) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnAddThemeColorOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) AddThemeConstantOverride(name StringName, constant int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&constant)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnAddThemeConstantOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) RemoveThemeIconOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnRemoveThemeIconOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) RemoveThemeStyleboxOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnRemoveThemeStyleboxOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) RemoveThemeFontOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnRemoveThemeFontOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) RemoveThemeFontSizeOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnRemoveThemeFontSizeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) RemoveThemeColorOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnRemoveThemeColorOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) RemoveThemeConstantOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnRemoveThemeConstantOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetThemeIcon(name StringName, theme_type StringName) Texture2D {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) GetThemeStylebox(name StringName, theme_type StringName) StyleBox {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStyleBox()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeStylebox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) GetThemeFont(name StringName, theme_type StringName) Font {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) GetThemeFontSize(name StringName, theme_type StringName) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) GetThemeColor(name StringName, theme_type StringName) Color {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) GetThemeConstant(name StringName, theme_type StringName) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeIconOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeIconOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeStyleboxOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeStyleboxOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeFontOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeFontOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeFontSizeOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeFontSizeOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeColorOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeColorOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeConstantOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeConstantOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeIcon(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeStylebox(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeStylebox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeFont(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeFontSize(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeColor(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) HasThemeConstant(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnHasThemeConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) GetThemeDefaultBaseScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeDefaultBaseScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) GetThemeDefaultFont() Font {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeDefaultFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Window) GetThemeDefaultFontSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetThemeDefaultFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) SetLayoutDirection(direction WindowLayoutDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetLayoutDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) GetLayoutDirection() WindowLayoutDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret WindowLayoutDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnGetLayoutDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Window) IsLayoutRtl() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnIsLayoutRtl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) SetAutoTranslate(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnSetAutoTranslate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) IsAutoTranslating() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnIsAutoTranslating), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Window) Popup(rect Rect2i) {
	cargs := []gdc.ConstTypePtr{rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) PopupOnParent(parent_rect Rect2i) {
	cargs := []gdc.ConstTypePtr{parent_rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopupOnParent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) PopupCentered(minsize Vector2i) {
	cargs := []gdc.ConstTypePtr{minsize.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopupCentered), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) PopupCenteredRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopupCenteredRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) PopupCenteredClamped(minsize Vector2i, fallback_ratio float64) {
	cargs := []gdc.ConstTypePtr{minsize.AsCTypePtr(), gdc.ConstTypePtr(&fallback_ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopupCenteredClamped), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) PopupExclusive(from_node Node, rect Rect2i) {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopupExclusive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) PopupExclusiveOnParent(from_node Node, parent_rect Rect2i) {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), parent_rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopupExclusiveOnParent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) PopupExclusiveCentered(from_node Node, minsize Vector2i) {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), minsize.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopupExclusiveCentered), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) PopupExclusiveCenteredRatio(from_node Node, ratio float64) {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopupExclusiveCenteredRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Window) PopupExclusiveCenteredClamped(from_node Node, minsize Vector2i, fallback_ratio float64) {
	cargs := []gdc.ConstTypePtr{from_node.AsCTypePtr(), minsize.AsCTypePtr(), gdc.ConstTypePtr(&fallback_ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForWindow.fnPopupExclusiveCenteredClamped), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type WindowWindowInputSignalFn func(event InputEvent)

func (me *Window) ConnectWindowInput(subs SignalSubscribers, fn WindowWindowInputSignalFn) {
	sig := StringNameFromStr("window_input")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectWindowInput(subs SignalSubscribers, fn WindowWindowInputSignalFn) {
	sig := StringNameFromStr("window_input")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowFilesDroppedSignalFn func(files PackedStringArray)

func (me *Window) ConnectFilesDropped(subs SignalSubscribers, fn WindowFilesDroppedSignalFn) {
	sig := StringNameFromStr("files_dropped")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectFilesDropped(subs SignalSubscribers, fn WindowFilesDroppedSignalFn) {
	sig := StringNameFromStr("files_dropped")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowMouseEnteredSignalFn func()

func (me *Window) ConnectMouseEntered(subs SignalSubscribers, fn WindowMouseEnteredSignalFn) {
	sig := StringNameFromStr("mouse_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectMouseEntered(subs SignalSubscribers, fn WindowMouseEnteredSignalFn) {
	sig := StringNameFromStr("mouse_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowMouseExitedSignalFn func()

func (me *Window) ConnectMouseExited(subs SignalSubscribers, fn WindowMouseExitedSignalFn) {
	sig := StringNameFromStr("mouse_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectMouseExited(subs SignalSubscribers, fn WindowMouseExitedSignalFn) {
	sig := StringNameFromStr("mouse_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowFocusEnteredSignalFn func()

func (me *Window) ConnectFocusEntered(subs SignalSubscribers, fn WindowFocusEnteredSignalFn) {
	sig := StringNameFromStr("focus_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectFocusEntered(subs SignalSubscribers, fn WindowFocusEnteredSignalFn) {
	sig := StringNameFromStr("focus_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowFocusExitedSignalFn func()

func (me *Window) ConnectFocusExited(subs SignalSubscribers, fn WindowFocusExitedSignalFn) {
	sig := StringNameFromStr("focus_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectFocusExited(subs SignalSubscribers, fn WindowFocusExitedSignalFn) {
	sig := StringNameFromStr("focus_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowCloseRequestedSignalFn func()

func (me *Window) ConnectCloseRequested(subs SignalSubscribers, fn WindowCloseRequestedSignalFn) {
	sig := StringNameFromStr("close_requested")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectCloseRequested(subs SignalSubscribers, fn WindowCloseRequestedSignalFn) {
	sig := StringNameFromStr("close_requested")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowGoBackRequestedSignalFn func()

func (me *Window) ConnectGoBackRequested(subs SignalSubscribers, fn WindowGoBackRequestedSignalFn) {
	sig := StringNameFromStr("go_back_requested")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectGoBackRequested(subs SignalSubscribers, fn WindowGoBackRequestedSignalFn) {
	sig := StringNameFromStr("go_back_requested")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowVisibilityChangedSignalFn func()

func (me *Window) ConnectVisibilityChanged(subs SignalSubscribers, fn WindowVisibilityChangedSignalFn) {
	sig := StringNameFromStr("visibility_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectVisibilityChanged(subs SignalSubscribers, fn WindowVisibilityChangedSignalFn) {
	sig := StringNameFromStr("visibility_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowAboutToPopupSignalFn func()

func (me *Window) ConnectAboutToPopup(subs SignalSubscribers, fn WindowAboutToPopupSignalFn) {
	sig := StringNameFromStr("about_to_popup")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectAboutToPopup(subs SignalSubscribers, fn WindowAboutToPopupSignalFn) {
	sig := StringNameFromStr("about_to_popup")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowThemeChangedSignalFn func()

func (me *Window) ConnectThemeChanged(subs SignalSubscribers, fn WindowThemeChangedSignalFn) {
	sig := StringNameFromStr("theme_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectThemeChanged(subs SignalSubscribers, fn WindowThemeChangedSignalFn) {
	sig := StringNameFromStr("theme_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowDpiChangedSignalFn func()

func (me *Window) ConnectDpiChanged(subs SignalSubscribers, fn WindowDpiChangedSignalFn) {
	sig := StringNameFromStr("dpi_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectDpiChanged(subs SignalSubscribers, fn WindowDpiChangedSignalFn) {
	sig := StringNameFromStr("dpi_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type WindowTitlebarChangedSignalFn func()

func (me *Window) ConnectTitlebarChanged(subs SignalSubscribers, fn WindowTitlebarChangedSignalFn) {
	sig := StringNameFromStr("titlebar_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Window) DisconnectTitlebarChanged(subs SignalSubscribers, fn WindowTitlebarChangedSignalFn) {
	sig := StringNameFromStr("titlebar_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
