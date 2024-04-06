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

type ptrsForControlList struct {
	fnXHasPoint                   gdc.MethodBindPtr
	fnXStructuredTextParser       gdc.MethodBindPtr
	fnXGetMinimumSize             gdc.MethodBindPtr
	fnXGetTooltip                 gdc.MethodBindPtr
	fnXGetDragData                gdc.MethodBindPtr
	fnXCanDropData                gdc.MethodBindPtr
	fnXDropData                   gdc.MethodBindPtr
	fnXMakeCustomTooltip          gdc.MethodBindPtr
	fnXGuiInput                   gdc.MethodBindPtr
	fnAcceptEvent                 gdc.MethodBindPtr
	fnGetMinimumSize              gdc.MethodBindPtr
	fnGetCombinedMinimumSize      gdc.MethodBindPtr
	fnSetAnchorsPreset            gdc.MethodBindPtr
	fnSetOffsetsPreset            gdc.MethodBindPtr
	fnSetAnchorsAndOffsetsPreset  gdc.MethodBindPtr
	fnSetAnchor                   gdc.MethodBindPtr
	fnGetAnchor                   gdc.MethodBindPtr
	fnSetOffset                   gdc.MethodBindPtr
	fnGetOffset                   gdc.MethodBindPtr
	fnSetAnchorAndOffset          gdc.MethodBindPtr
	fnSetBegin                    gdc.MethodBindPtr
	fnSetEnd                      gdc.MethodBindPtr
	fnSetPosition                 gdc.MethodBindPtr
	fnSetSize                     gdc.MethodBindPtr
	fnResetSize                   gdc.MethodBindPtr
	fnSetCustomMinimumSize        gdc.MethodBindPtr
	fnSetGlobalPosition           gdc.MethodBindPtr
	fnSetRotation                 gdc.MethodBindPtr
	fnSetRotationDegrees          gdc.MethodBindPtr
	fnSetScale                    gdc.MethodBindPtr
	fnSetPivotOffset              gdc.MethodBindPtr
	fnGetBegin                    gdc.MethodBindPtr
	fnGetEnd                      gdc.MethodBindPtr
	fnGetPosition                 gdc.MethodBindPtr
	fnGetSize                     gdc.MethodBindPtr
	fnGetRotation                 gdc.MethodBindPtr
	fnGetRotationDegrees          gdc.MethodBindPtr
	fnGetScale                    gdc.MethodBindPtr
	fnGetPivotOffset              gdc.MethodBindPtr
	fnGetCustomMinimumSize        gdc.MethodBindPtr
	fnGetParentAreaSize           gdc.MethodBindPtr
	fnGetGlobalPosition           gdc.MethodBindPtr
	fnGetScreenPosition           gdc.MethodBindPtr
	fnGetRect                     gdc.MethodBindPtr
	fnGetGlobalRect               gdc.MethodBindPtr
	fnSetFocusMode                gdc.MethodBindPtr
	fnGetFocusMode                gdc.MethodBindPtr
	fnHasFocus                    gdc.MethodBindPtr
	fnGrabFocus                   gdc.MethodBindPtr
	fnReleaseFocus                gdc.MethodBindPtr
	fnFindPrevValidFocus          gdc.MethodBindPtr
	fnFindNextValidFocus          gdc.MethodBindPtr
	fnFindValidFocusNeighbor      gdc.MethodBindPtr
	fnSetHSizeFlags               gdc.MethodBindPtr
	fnGetHSizeFlags               gdc.MethodBindPtr
	fnSetStretchRatio             gdc.MethodBindPtr
	fnGetStretchRatio             gdc.MethodBindPtr
	fnSetVSizeFlags               gdc.MethodBindPtr
	fnGetVSizeFlags               gdc.MethodBindPtr
	fnSetTheme                    gdc.MethodBindPtr
	fnGetTheme                    gdc.MethodBindPtr
	fnSetThemeTypeVariation       gdc.MethodBindPtr
	fnGetThemeTypeVariation       gdc.MethodBindPtr
	fnBeginBulkThemeOverride      gdc.MethodBindPtr
	fnEndBulkThemeOverride        gdc.MethodBindPtr
	fnAddThemeIconOverride        gdc.MethodBindPtr
	fnAddThemeStyleboxOverride    gdc.MethodBindPtr
	fnAddThemeFontOverride        gdc.MethodBindPtr
	fnAddThemeFontSizeOverride    gdc.MethodBindPtr
	fnAddThemeColorOverride       gdc.MethodBindPtr
	fnAddThemeConstantOverride    gdc.MethodBindPtr
	fnRemoveThemeIconOverride     gdc.MethodBindPtr
	fnRemoveThemeStyleboxOverride gdc.MethodBindPtr
	fnRemoveThemeFontOverride     gdc.MethodBindPtr
	fnRemoveThemeFontSizeOverride gdc.MethodBindPtr
	fnRemoveThemeColorOverride    gdc.MethodBindPtr
	fnRemoveThemeConstantOverride gdc.MethodBindPtr
	fnGetThemeIcon                gdc.MethodBindPtr
	fnGetThemeStylebox            gdc.MethodBindPtr
	fnGetThemeFont                gdc.MethodBindPtr
	fnGetThemeFontSize            gdc.MethodBindPtr
	fnGetThemeColor               gdc.MethodBindPtr
	fnGetThemeConstant            gdc.MethodBindPtr
	fnHasThemeIconOverride        gdc.MethodBindPtr
	fnHasThemeStyleboxOverride    gdc.MethodBindPtr
	fnHasThemeFontOverride        gdc.MethodBindPtr
	fnHasThemeFontSizeOverride    gdc.MethodBindPtr
	fnHasThemeColorOverride       gdc.MethodBindPtr
	fnHasThemeConstantOverride    gdc.MethodBindPtr
	fnHasThemeIcon                gdc.MethodBindPtr
	fnHasThemeStylebox            gdc.MethodBindPtr
	fnHasThemeFont                gdc.MethodBindPtr
	fnHasThemeFontSize            gdc.MethodBindPtr
	fnHasThemeColor               gdc.MethodBindPtr
	fnHasThemeConstant            gdc.MethodBindPtr
	fnGetThemeDefaultBaseScale    gdc.MethodBindPtr
	fnGetThemeDefaultFont         gdc.MethodBindPtr
	fnGetThemeDefaultFontSize     gdc.MethodBindPtr
	fnGetParentControl            gdc.MethodBindPtr
	fnSetHGrowDirection           gdc.MethodBindPtr
	fnGetHGrowDirection           gdc.MethodBindPtr
	fnSetVGrowDirection           gdc.MethodBindPtr
	fnGetVGrowDirection           gdc.MethodBindPtr
	fnSetTooltipText              gdc.MethodBindPtr
	fnGetTooltipText              gdc.MethodBindPtr
	fnGetTooltip                  gdc.MethodBindPtr
	fnSetDefaultCursorShape       gdc.MethodBindPtr
	fnGetDefaultCursorShape       gdc.MethodBindPtr
	fnGetCursorShape              gdc.MethodBindPtr
	fnSetFocusNeighbor            gdc.MethodBindPtr
	fnGetFocusNeighbor            gdc.MethodBindPtr
	fnSetFocusNext                gdc.MethodBindPtr
	fnGetFocusNext                gdc.MethodBindPtr
	fnSetFocusPrevious            gdc.MethodBindPtr
	fnGetFocusPrevious            gdc.MethodBindPtr
	fnForceDrag                   gdc.MethodBindPtr
	fnSetMouseFilter              gdc.MethodBindPtr
	fnGetMouseFilter              gdc.MethodBindPtr
	fnSetForcePassScrollEvents    gdc.MethodBindPtr
	fnIsForcePassScrollEvents     gdc.MethodBindPtr
	fnSetClipContents             gdc.MethodBindPtr
	fnIsClippingContents          gdc.MethodBindPtr
	fnGrabClickFocus              gdc.MethodBindPtr
	fnSetDragForwarding           gdc.MethodBindPtr
	fnSetDragPreview              gdc.MethodBindPtr
	fnIsDragSuccessful            gdc.MethodBindPtr
	fnWarpMouse                   gdc.MethodBindPtr
	fnSetShortcutContext          gdc.MethodBindPtr
	fnGetShortcutContext          gdc.MethodBindPtr
	fnUpdateMinimumSize           gdc.MethodBindPtr
	fnSetLayoutDirection          gdc.MethodBindPtr
	fnGetLayoutDirection          gdc.MethodBindPtr
	fnIsLayoutRtl                 gdc.MethodBindPtr
	fnSetAutoTranslate            gdc.MethodBindPtr
	fnIsAutoTranslating           gdc.MethodBindPtr
	fnSetLocalizeNumeralSystem    gdc.MethodBindPtr
	fnIsLocalizingNumeralSystem   gdc.MethodBindPtr
}

var ptrsForControl ptrsForControlList

func initControlPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Control")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("accept_event")
		defer methodName.Destroy()
		ptrsForControl.fnAcceptEvent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_minimum_size")
		defer methodName.Destroy()
		ptrsForControl.fnGetMinimumSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_combined_minimum_size")
		defer methodName.Destroy()
		ptrsForControl.fnGetCombinedMinimumSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_anchors_preset")
		defer methodName.Destroy()
		ptrsForControl.fnSetAnchorsPreset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 509135270))
	}
	{
		methodName := StringNameFromStr("set_offsets_preset")
		defer methodName.Destroy()
		ptrsForControl.fnSetOffsetsPreset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3724524307))
	}
	{
		methodName := StringNameFromStr("set_anchors_and_offsets_preset")
		defer methodName.Destroy()
		ptrsForControl.fnSetAnchorsAndOffsetsPreset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3724524307))
	}
	{
		methodName := StringNameFromStr("set_anchor")
		defer methodName.Destroy()
		ptrsForControl.fnSetAnchor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2302782885))
	}
	{
		methodName := StringNameFromStr("get_anchor")
		defer methodName.Destroy()
		ptrsForControl.fnGetAnchor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2869120046))
	}
	{
		methodName := StringNameFromStr("set_offset")
		defer methodName.Destroy()
		ptrsForControl.fnSetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4290182280))
	}
	{
		methodName := StringNameFromStr("get_offset")
		defer methodName.Destroy()
		ptrsForControl.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2869120046))
	}
	{
		methodName := StringNameFromStr("set_anchor_and_offset")
		defer methodName.Destroy()
		ptrsForControl.fnSetAnchorAndOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4031722181))
	}
	{
		methodName := StringNameFromStr("set_begin")
		defer methodName.Destroy()
		ptrsForControl.fnSetBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_end")
		defer methodName.Destroy()
		ptrsForControl.fnSetEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_position")
		defer methodName.Destroy()
		ptrsForControl.fnSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2436320129))
	}
	{
		methodName := StringNameFromStr("set_size")
		defer methodName.Destroy()
		ptrsForControl.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2436320129))
	}
	{
		methodName := StringNameFromStr("reset_size")
		defer methodName.Destroy()
		ptrsForControl.fnResetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_custom_minimum_size")
		defer methodName.Destroy()
		ptrsForControl.fnSetCustomMinimumSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_global_position")
		defer methodName.Destroy()
		ptrsForControl.fnSetGlobalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2436320129))
	}
	{
		methodName := StringNameFromStr("set_rotation")
		defer methodName.Destroy()
		ptrsForControl.fnSetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_rotation_degrees")
		defer methodName.Destroy()
		ptrsForControl.fnSetRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("set_scale")
		defer methodName.Destroy()
		ptrsForControl.fnSetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_pivot_offset")
		defer methodName.Destroy()
		ptrsForControl.fnSetPivotOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_begin")
		defer methodName.Destroy()
		ptrsForControl.fnGetBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_end")
		defer methodName.Destroy()
		ptrsForControl.fnGetEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForControl.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_size")
		defer methodName.Destroy()
		ptrsForControl.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_rotation")
		defer methodName.Destroy()
		ptrsForControl.fnGetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_rotation_degrees")
		defer methodName.Destroy()
		ptrsForControl.fnGetRotationDegrees = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_scale")
		defer methodName.Destroy()
		ptrsForControl.fnGetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_pivot_offset")
		defer methodName.Destroy()
		ptrsForControl.fnGetPivotOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_custom_minimum_size")
		defer methodName.Destroy()
		ptrsForControl.fnGetCustomMinimumSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_parent_area_size")
		defer methodName.Destroy()
		ptrsForControl.fnGetParentAreaSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_global_position")
		defer methodName.Destroy()
		ptrsForControl.fnGetGlobalPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_screen_position")
		defer methodName.Destroy()
		ptrsForControl.fnGetScreenPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_rect")
		defer methodName.Destroy()
		ptrsForControl.fnGetRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("get_global_rect")
		defer methodName.Destroy()
		ptrsForControl.fnGetGlobalRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
	}
	{
		methodName := StringNameFromStr("set_focus_mode")
		defer methodName.Destroy()
		ptrsForControl.fnSetFocusMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3232914922))
	}
	{
		methodName := StringNameFromStr("get_focus_mode")
		defer methodName.Destroy()
		ptrsForControl.fnGetFocusMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2132829277))
	}
	{
		methodName := StringNameFromStr("has_focus")
		defer methodName.Destroy()
		ptrsForControl.fnHasFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("grab_focus")
		defer methodName.Destroy()
		ptrsForControl.fnGrabFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("release_focus")
		defer methodName.Destroy()
		ptrsForControl.fnReleaseFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("find_prev_valid_focus")
		defer methodName.Destroy()
		ptrsForControl.fnFindPrevValidFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2783021301))
	}
	{
		methodName := StringNameFromStr("find_next_valid_focus")
		defer methodName.Destroy()
		ptrsForControl.fnFindNextValidFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2783021301))
	}
	{
		methodName := StringNameFromStr("find_valid_focus_neighbor")
		defer methodName.Destroy()
		ptrsForControl.fnFindValidFocusNeighbor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1543910170))
	}
	{
		methodName := StringNameFromStr("set_h_size_flags")
		defer methodName.Destroy()
		ptrsForControl.fnSetHSizeFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 394851643))
	}
	{
		methodName := StringNameFromStr("get_h_size_flags")
		defer methodName.Destroy()
		ptrsForControl.fnGetHSizeFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3781367401))
	}
	{
		methodName := StringNameFromStr("set_stretch_ratio")
		defer methodName.Destroy()
		ptrsForControl.fnSetStretchRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_stretch_ratio")
		defer methodName.Destroy()
		ptrsForControl.fnGetStretchRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_v_size_flags")
		defer methodName.Destroy()
		ptrsForControl.fnSetVSizeFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 394851643))
	}
	{
		methodName := StringNameFromStr("get_v_size_flags")
		defer methodName.Destroy()
		ptrsForControl.fnGetVSizeFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3781367401))
	}
	{
		methodName := StringNameFromStr("set_theme")
		defer methodName.Destroy()
		ptrsForControl.fnSetTheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2326690814))
	}
	{
		methodName := StringNameFromStr("get_theme")
		defer methodName.Destroy()
		ptrsForControl.fnGetTheme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3846893731))
	}
	{
		methodName := StringNameFromStr("set_theme_type_variation")
		defer methodName.Destroy()
		ptrsForControl.fnSetThemeTypeVariation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_theme_type_variation")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeTypeVariation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("begin_bulk_theme_override")
		defer methodName.Destroy()
		ptrsForControl.fnBeginBulkThemeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("end_bulk_theme_override")
		defer methodName.Destroy()
		ptrsForControl.fnEndBulkThemeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("add_theme_icon_override")
		defer methodName.Destroy()
		ptrsForControl.fnAddThemeIconOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1373065600))
	}
	{
		methodName := StringNameFromStr("add_theme_stylebox_override")
		defer methodName.Destroy()
		ptrsForControl.fnAddThemeStyleboxOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4188838905))
	}
	{
		methodName := StringNameFromStr("add_theme_font_override")
		defer methodName.Destroy()
		ptrsForControl.fnAddThemeFontOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3518018674))
	}
	{
		methodName := StringNameFromStr("add_theme_font_size_override")
		defer methodName.Destroy()
		ptrsForControl.fnAddThemeFontSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2415702435))
	}
	{
		methodName := StringNameFromStr("add_theme_color_override")
		defer methodName.Destroy()
		ptrsForControl.fnAddThemeColorOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4260178595))
	}
	{
		methodName := StringNameFromStr("add_theme_constant_override")
		defer methodName.Destroy()
		ptrsForControl.fnAddThemeConstantOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2415702435))
	}
	{
		methodName := StringNameFromStr("remove_theme_icon_override")
		defer methodName.Destroy()
		ptrsForControl.fnRemoveThemeIconOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_stylebox_override")
		defer methodName.Destroy()
		ptrsForControl.fnRemoveThemeStyleboxOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_font_override")
		defer methodName.Destroy()
		ptrsForControl.fnRemoveThemeFontOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_font_size_override")
		defer methodName.Destroy()
		ptrsForControl.fnRemoveThemeFontSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_color_override")
		defer methodName.Destroy()
		ptrsForControl.fnRemoveThemeColorOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("remove_theme_constant_override")
		defer methodName.Destroy()
		ptrsForControl.fnRemoveThemeConstantOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_theme_icon")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3163973443))
	}
	{
		methodName := StringNameFromStr("get_theme_stylebox")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 604739069))
	}
	{
		methodName := StringNameFromStr("get_theme_font")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2826986490))
	}
	{
		methodName := StringNameFromStr("get_theme_font_size")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1327056374))
	}
	{
		methodName := StringNameFromStr("get_theme_color")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2798751242))
	}
	{
		methodName := StringNameFromStr("get_theme_constant")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1327056374))
	}
	{
		methodName := StringNameFromStr("has_theme_icon_override")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeIconOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_stylebox_override")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeStyleboxOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_font_override")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeFontOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_font_size_override")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeFontSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_color_override")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeColorOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_constant_override")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeConstantOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("has_theme_icon")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_stylebox")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeStylebox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_font")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_font_size")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_color")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("has_theme_constant")
		defer methodName.Destroy()
		ptrsForControl.fnHasThemeConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 866386512))
	}
	{
		methodName := StringNameFromStr("get_theme_default_base_scale")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeDefaultBaseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_theme_default_font")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeDefaultFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229501585))
	}
	{
		methodName := StringNameFromStr("get_theme_default_font_size")
		defer methodName.Destroy()
		ptrsForControl.fnGetThemeDefaultFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_parent_control")
		defer methodName.Destroy()
		ptrsForControl.fnGetParentControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2783021301))
	}
	{
		methodName := StringNameFromStr("set_h_grow_direction")
		defer methodName.Destroy()
		ptrsForControl.fnSetHGrowDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2022385301))
	}
	{
		methodName := StringNameFromStr("get_h_grow_direction")
		defer methodName.Destroy()
		ptrsForControl.fnGetHGrowDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635610155))
	}
	{
		methodName := StringNameFromStr("set_v_grow_direction")
		defer methodName.Destroy()
		ptrsForControl.fnSetVGrowDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2022385301))
	}
	{
		methodName := StringNameFromStr("get_v_grow_direction")
		defer methodName.Destroy()
		ptrsForControl.fnGetVGrowDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635610155))
	}
	{
		methodName := StringNameFromStr("set_tooltip_text")
		defer methodName.Destroy()
		ptrsForControl.fnSetTooltipText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_tooltip_text")
		defer methodName.Destroy()
		ptrsForControl.fnGetTooltipText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_tooltip")
		defer methodName.Destroy()
		ptrsForControl.fnGetTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2895288280))
	}
	{
		methodName := StringNameFromStr("set_default_cursor_shape")
		defer methodName.Destroy()
		ptrsForControl.fnSetDefaultCursorShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 217062046))
	}
	{
		methodName := StringNameFromStr("get_default_cursor_shape")
		defer methodName.Destroy()
		ptrsForControl.fnGetDefaultCursorShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2359535750))
	}
	{
		methodName := StringNameFromStr("get_cursor_shape")
		defer methodName.Destroy()
		ptrsForControl.fnGetCursorShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1395773853))
	}
	{
		methodName := StringNameFromStr("set_focus_neighbor")
		defer methodName.Destroy()
		ptrsForControl.fnSetFocusNeighbor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2024461774))
	}
	{
		methodName := StringNameFromStr("get_focus_neighbor")
		defer methodName.Destroy()
		ptrsForControl.fnGetFocusNeighbor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757935761))
	}
	{
		methodName := StringNameFromStr("set_focus_next")
		defer methodName.Destroy()
		ptrsForControl.fnSetFocusNext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_focus_next")
		defer methodName.Destroy()
		ptrsForControl.fnGetFocusNext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_focus_previous")
		defer methodName.Destroy()
		ptrsForControl.fnSetFocusPrevious = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_focus_previous")
		defer methodName.Destroy()
		ptrsForControl.fnGetFocusPrevious = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("force_drag")
		defer methodName.Destroy()
		ptrsForControl.fnForceDrag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3191844692))
	}
	{
		methodName := StringNameFromStr("set_mouse_filter")
		defer methodName.Destroy()
		ptrsForControl.fnSetMouseFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3891156122))
	}
	{
		methodName := StringNameFromStr("get_mouse_filter")
		defer methodName.Destroy()
		ptrsForControl.fnGetMouseFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1572545674))
	}
	{
		methodName := StringNameFromStr("set_force_pass_scroll_events")
		defer methodName.Destroy()
		ptrsForControl.fnSetForcePassScrollEvents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_force_pass_scroll_events")
		defer methodName.Destroy()
		ptrsForControl.fnIsForcePassScrollEvents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_clip_contents")
		defer methodName.Destroy()
		ptrsForControl.fnSetClipContents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_clipping_contents")
		defer methodName.Destroy()
		ptrsForControl.fnIsClippingContents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("grab_click_focus")
		defer methodName.Destroy()
		ptrsForControl.fnGrabClickFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_drag_forwarding")
		defer methodName.Destroy()
		ptrsForControl.fnSetDragForwarding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1076571380))
	}
	{
		methodName := StringNameFromStr("set_drag_preview")
		defer methodName.Destroy()
		ptrsForControl.fnSetDragPreview = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1496901182))
	}
	{
		methodName := StringNameFromStr("is_drag_successful")
		defer methodName.Destroy()
		ptrsForControl.fnIsDragSuccessful = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("warp_mouse")
		defer methodName.Destroy()
		ptrsForControl.fnWarpMouse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("set_shortcut_context")
		defer methodName.Destroy()
		ptrsForControl.fnSetShortcutContext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("get_shortcut_context")
		defer methodName.Destroy()
		ptrsForControl.fnGetShortcutContext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
	}
	{
		methodName := StringNameFromStr("update_minimum_size")
		defer methodName.Destroy()
		ptrsForControl.fnUpdateMinimumSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_layout_direction")
		defer methodName.Destroy()
		ptrsForControl.fnSetLayoutDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3310692370))
	}
	{
		methodName := StringNameFromStr("get_layout_direction")
		defer methodName.Destroy()
		ptrsForControl.fnGetLayoutDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1546772008))
	}
	{
		methodName := StringNameFromStr("is_layout_rtl")
		defer methodName.Destroy()
		ptrsForControl.fnIsLayoutRtl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_auto_translate")
		defer methodName.Destroy()
		ptrsForControl.fnSetAutoTranslate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_auto_translating")
		defer methodName.Destroy()
		ptrsForControl.fnIsAutoTranslating = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_localize_numeral_system")
		defer methodName.Destroy()
		ptrsForControl.fnSetLocalizeNumeralSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_localizing_numeral_system")
		defer methodName.Destroy()
		ptrsForControl.fnIsLocalizingNumeralSystem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type Control struct {
	CanvasItem
}

func (me *Control) BaseClass() string {
	return "Control"
}

func NewControl() *Control {
	str := StringNameFromStr("Control") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Control{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	ControlNotificationResized                = 40
	ControlNotificationMouseEnter             = 41
	ControlNotificationMouseExit              = 42
	ControlNotificationMouseEnterSelf         = 60
	ControlNotificationMouseExitSelf          = 61
	ControlNotificationFocusEnter             = 43
	ControlNotificationFocusExit              = 44
	ControlNotificationThemeChanged           = 45
	ControlNotificationScrollBegin            = 47
	ControlNotificationScrollEnd              = 48
	ControlNotificationLayoutDirectionChanged = 49
)

// Enums

type ControlFocusMode int

const (
	ControlFocusModeFocusNone  ControlFocusMode = 0
	ControlFocusModeFocusClick ControlFocusMode = 1
	ControlFocusModeFocusAll   ControlFocusMode = 2
)

type ControlCursorShape int

const (
	ControlCursorShapeCursorArrow        ControlCursorShape = 0
	ControlCursorShapeCursorIbeam        ControlCursorShape = 1
	ControlCursorShapeCursorPointingHand ControlCursorShape = 2
	ControlCursorShapeCursorCross        ControlCursorShape = 3
	ControlCursorShapeCursorWait         ControlCursorShape = 4
	ControlCursorShapeCursorBusy         ControlCursorShape = 5
	ControlCursorShapeCursorDrag         ControlCursorShape = 6
	ControlCursorShapeCursorCanDrop      ControlCursorShape = 7
	ControlCursorShapeCursorForbidden    ControlCursorShape = 8
	ControlCursorShapeCursorVsize        ControlCursorShape = 9
	ControlCursorShapeCursorHsize        ControlCursorShape = 10
	ControlCursorShapeCursorBdiagsize    ControlCursorShape = 11
	ControlCursorShapeCursorFdiagsize    ControlCursorShape = 12
	ControlCursorShapeCursorMove         ControlCursorShape = 13
	ControlCursorShapeCursorVsplit       ControlCursorShape = 14
	ControlCursorShapeCursorHsplit       ControlCursorShape = 15
	ControlCursorShapeCursorHelp         ControlCursorShape = 16
)

type ControlLayoutPreset int

const (
	ControlLayoutPresetPresetTopLeft      ControlLayoutPreset = 0
	ControlLayoutPresetPresetTopRight     ControlLayoutPreset = 1
	ControlLayoutPresetPresetBottomLeft   ControlLayoutPreset = 2
	ControlLayoutPresetPresetBottomRight  ControlLayoutPreset = 3
	ControlLayoutPresetPresetCenterLeft   ControlLayoutPreset = 4
	ControlLayoutPresetPresetCenterTop    ControlLayoutPreset = 5
	ControlLayoutPresetPresetCenterRight  ControlLayoutPreset = 6
	ControlLayoutPresetPresetCenterBottom ControlLayoutPreset = 7
	ControlLayoutPresetPresetCenter       ControlLayoutPreset = 8
	ControlLayoutPresetPresetLeftWide     ControlLayoutPreset = 9
	ControlLayoutPresetPresetTopWide      ControlLayoutPreset = 10
	ControlLayoutPresetPresetRightWide    ControlLayoutPreset = 11
	ControlLayoutPresetPresetBottomWide   ControlLayoutPreset = 12
	ControlLayoutPresetPresetVcenterWide  ControlLayoutPreset = 13
	ControlLayoutPresetPresetHcenterWide  ControlLayoutPreset = 14
	ControlLayoutPresetPresetFullRect     ControlLayoutPreset = 15
)

type ControlLayoutPresetMode int

const (
	ControlLayoutPresetModePresetModeMinsize    ControlLayoutPresetMode = 0
	ControlLayoutPresetModePresetModeKeepWidth  ControlLayoutPresetMode = 1
	ControlLayoutPresetModePresetModeKeepHeight ControlLayoutPresetMode = 2
	ControlLayoutPresetModePresetModeKeepSize   ControlLayoutPresetMode = 3
)

type ControlSizeFlags int

const (
	ControlSizeFlagsSizeShrinkBegin  ControlSizeFlags = 0
	ControlSizeFlagsSizeFill         ControlSizeFlags = 1
	ControlSizeFlagsSizeExpand       ControlSizeFlags = 2
	ControlSizeFlagsSizeExpandFill   ControlSizeFlags = 3
	ControlSizeFlagsSizeShrinkCenter ControlSizeFlags = 4
	ControlSizeFlagsSizeShrinkEnd    ControlSizeFlags = 8
)

type ControlMouseFilter int

const (
	ControlMouseFilterMouseFilterStop   ControlMouseFilter = 0
	ControlMouseFilterMouseFilterPass   ControlMouseFilter = 1
	ControlMouseFilterMouseFilterIgnore ControlMouseFilter = 2
)

type ControlGrowDirection int

const (
	ControlGrowDirectionGrowDirectionBegin ControlGrowDirection = 0
	ControlGrowDirectionGrowDirectionEnd   ControlGrowDirection = 1
	ControlGrowDirectionGrowDirectionBoth  ControlGrowDirection = 2
)

type ControlAnchor int

const (
	ControlAnchorAnchorBegin ControlAnchor = 0
	ControlAnchorAnchorEnd   ControlAnchor = 1
)

type ControlLayoutDirection int

const (
	ControlLayoutDirectionLayoutDirectionInherited ControlLayoutDirection = 0
	ControlLayoutDirectionLayoutDirectionLocale    ControlLayoutDirection = 1
	ControlLayoutDirectionLayoutDirectionLtr       ControlLayoutDirection = 2
	ControlLayoutDirectionLayoutDirectionRtl       ControlLayoutDirection = 3
)

type ControlTextDirection int

const (
	ControlTextDirectionTextDirectionInherited ControlTextDirection = 3
	ControlTextDirectionTextDirectionAuto      ControlTextDirection = 0
	ControlTextDirectionTextDirectionLtr       ControlTextDirection = 1
	ControlTextDirectionTextDirectionRtl       ControlTextDirection = 2
)

func (me *Control) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Control) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Control) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Control) AcceptEvent() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnAcceptEvent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetMinimumSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetMinimumSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetCombinedMinimumSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetCombinedMinimumSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) SetAnchorsPreset(preset ControlLayoutPreset, keep_offsets bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&preset), gdc.ConstTypePtr(&keep_offsets)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetAnchorsPreset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetOffsetsPreset(preset ControlLayoutPreset, resize_mode ControlLayoutPresetMode, margin int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&preset), gdc.ConstTypePtr(&resize_mode), gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetOffsetsPreset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetAnchorsAndOffsetsPreset(preset ControlLayoutPreset, resize_mode ControlLayoutPresetMode, margin int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&preset), gdc.ConstTypePtr(&resize_mode), gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetAnchorsAndOffsetsPreset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetAnchor(side Side, anchor float64, keep_offset bool, push_opposite_anchor bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), gdc.ConstTypePtr(&anchor), gdc.ConstTypePtr(&keep_offset), gdc.ConstTypePtr(&push_opposite_anchor)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetAnchor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetAnchor(side Side) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&side)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetAnchor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) SetOffset(side Side, offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetOffset(offset Side) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&offset)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) SetAnchorAndOffset(side Side, anchor float64, offset float64, push_opposite_anchor bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), gdc.ConstTypePtr(&anchor), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&push_opposite_anchor)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetAnchorAndOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetBegin(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetBegin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetEnd(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetPosition(position Vector2, keep_offsets bool) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&keep_offsets)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetSize(size Vector2, keep_offsets bool) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), gdc.ConstTypePtr(&keep_offsets)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) ResetSize() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnResetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetCustomMinimumSize(size Vector2) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetCustomMinimumSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetGlobalPosition(position Vector2, keep_offsets bool) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&keep_offsets)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetGlobalPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetRotation(radians float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetRotationDegrees(degrees float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetRotationDegrees), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetScale(scale Vector2) {
	cargs := []gdc.ConstTypePtr{scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetPivotOffset(pivot_offset Vector2) {
	cargs := []gdc.ConstTypePtr{pivot_offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetPivotOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetBegin() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetEnd() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetRotation() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) GetRotationDegrees() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetRotationDegrees), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) GetScale() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetPivotOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetPivotOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetCustomMinimumSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetCustomMinimumSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetParentAreaSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetParentAreaSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetGlobalPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetGlobalPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetScreenPosition() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetScreenPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetGlobalRect() Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetGlobalRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) SetFocusMode(mode ControlFocusMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetFocusMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetFocusMode() ControlFocusMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlFocusMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetFocusMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Control) HasFocus() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasFocus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) GrabFocus() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGrabFocus), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) ReleaseFocus() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnReleaseFocus), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) FindPrevValidFocus() Control {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewControl()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnFindPrevValidFocus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) FindNextValidFocus() Control {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewControl()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnFindNextValidFocus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) FindValidFocusNeighbor(side Side) Control {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewControl()
	pinner.Pin(&side)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnFindValidFocusNeighbor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) SetHSizeFlags(flags ControlSizeFlags) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetHSizeFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetHSizeFlags() ControlSizeFlags {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlSizeFlags

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetHSizeFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Control) SetStretchRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetStretchRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetStretchRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetStretchRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) SetVSizeFlags(flags ControlSizeFlags) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetVSizeFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetVSizeFlags() ControlSizeFlags {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlSizeFlags

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetVSizeFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Control) SetTheme(theme Theme) {
	cargs := []gdc.ConstTypePtr{theme.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetTheme), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetTheme() Theme {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTheme()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetTheme), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) SetThemeTypeVariation(theme_type StringName) {
	cargs := []gdc.ConstTypePtr{theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetThemeTypeVariation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetThemeTypeVariation() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeTypeVariation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) BeginBulkThemeOverride() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnBeginBulkThemeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) EndBulkThemeOverride() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnEndBulkThemeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) AddThemeIconOverride(name StringName, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnAddThemeIconOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) AddThemeStyleboxOverride(name StringName, stylebox StyleBox) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), stylebox.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnAddThemeStyleboxOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) AddThemeFontOverride(name StringName, font Font) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), font.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnAddThemeFontOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) AddThemeFontSizeOverride(name StringName, font_size int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnAddThemeFontSizeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) AddThemeColorOverride(name StringName, color Color) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnAddThemeColorOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) AddThemeConstantOverride(name StringName, constant int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&constant)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnAddThemeConstantOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) RemoveThemeIconOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnRemoveThemeIconOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) RemoveThemeStyleboxOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnRemoveThemeStyleboxOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) RemoveThemeFontOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnRemoveThemeFontOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) RemoveThemeFontSizeOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnRemoveThemeFontSizeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) RemoveThemeColorOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnRemoveThemeColorOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) RemoveThemeConstantOverride(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnRemoveThemeConstantOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetThemeIcon(name StringName, theme_type StringName) Texture2D {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetThemeStylebox(name StringName, theme_type StringName) StyleBox {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStyleBox()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeStylebox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetThemeFont(name StringName, theme_type StringName) Font {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetThemeFontSize(name StringName, theme_type StringName) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) GetThemeColor(name StringName, theme_type StringName) Color {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetThemeConstant(name StringName, theme_type StringName) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeIconOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeIconOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeStyleboxOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeStyleboxOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeFontOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeFontOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeFontSizeOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeFontSizeOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeColorOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeColorOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeConstantOverride(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeConstantOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeIcon(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeStylebox(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeStylebox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeFont(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeFontSize(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeColor(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) HasThemeConstant(name StringName, theme_type StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), theme_type.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnHasThemeConstant), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) GetThemeDefaultBaseScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeDefaultBaseScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) GetThemeDefaultFont() Font {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeDefaultFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetThemeDefaultFontSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetThemeDefaultFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) GetParentControl() Control {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewControl()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetParentControl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) SetHGrowDirection(direction ControlGrowDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetHGrowDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetHGrowDirection() ControlGrowDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlGrowDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetHGrowDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Control) SetVGrowDirection(direction ControlGrowDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetVGrowDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetVGrowDirection() ControlGrowDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlGrowDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetVGrowDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Control) SetTooltipText(hint String) {
	cargs := []gdc.ConstTypePtr{hint.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetTooltipText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetTooltipText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetTooltipText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) GetTooltip(at_position Vector2) String {
	cargs := []gdc.ConstTypePtr{at_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetTooltip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) SetDefaultCursorShape(shape ControlCursorShape) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetDefaultCursorShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetDefaultCursorShape() ControlCursorShape {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlCursorShape

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetDefaultCursorShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Control) GetCursorShape(position Vector2) ControlCursorShape {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlCursorShape

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetCursorShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Control) SetFocusNeighbor(side Side, neighbor NodePath) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side), neighbor.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetFocusNeighbor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetFocusNeighbor(side Side) NodePath {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&side)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()
	pinner.Pin(&side)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetFocusNeighbor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) SetFocusNext(next NodePath) {
	cargs := []gdc.ConstTypePtr{next.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetFocusNext), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetFocusNext() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetFocusNext), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) SetFocusPrevious(previous NodePath) {
	cargs := []gdc.ConstTypePtr{previous.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetFocusPrevious), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetFocusPrevious() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetFocusPrevious), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) ForceDrag(data Variant, preview Control) {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr(), preview.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnForceDrag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetMouseFilter(filter ControlMouseFilter) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetMouseFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetMouseFilter() ControlMouseFilter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlMouseFilter

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetMouseFilter), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Control) SetForcePassScrollEvents(force_pass_scroll_events bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force_pass_scroll_events)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetForcePassScrollEvents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) IsForcePassScrollEvents() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnIsForcePassScrollEvents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) SetClipContents(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetClipContents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) IsClippingContents() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnIsClippingContents), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) GrabClickFocus() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGrabClickFocus), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetDragForwarding(drag_func Callable, can_drop_func Callable, drop_func Callable) {
	cargs := []gdc.ConstTypePtr{drag_func.AsCTypePtr(), can_drop_func.AsCTypePtr(), drop_func.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetDragForwarding), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetDragPreview(control Control) {
	cargs := []gdc.ConstTypePtr{control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetDragPreview), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) IsDragSuccessful() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnIsDragSuccessful), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) WarpMouse(position Vector2) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnWarpMouse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetShortcutContext(node Node) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetShortcutContext), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetShortcutContext() Node {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetShortcutContext), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Control) UpdateMinimumSize() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnUpdateMinimumSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) SetLayoutDirection(direction ControlLayoutDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetLayoutDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) GetLayoutDirection() ControlLayoutDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlLayoutDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnGetLayoutDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Control) IsLayoutRtl() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnIsLayoutRtl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) SetAutoTranslate(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetAutoTranslate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) IsAutoTranslating() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnIsAutoTranslating), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Control) SetLocalizeNumeralSystem(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnSetLocalizeNumeralSystem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Control) IsLocalizingNumeralSystem() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForControl.fnIsLocalizingNumeralSystem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ControlResizedSignalFn func()

func (me *Control) ConnectResized(subs SignalSubscribers, fn ControlResizedSignalFn) {
	sig := StringNameFromStr("resized")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Control) DisconnectResized(subs SignalSubscribers, fn ControlResizedSignalFn) {
	sig := StringNameFromStr("resized")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ControlGuiInputSignalFn func(event InputEvent)

func (me *Control) ConnectGuiInput(subs SignalSubscribers, fn ControlGuiInputSignalFn) {
	sig := StringNameFromStr("gui_input")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Control) DisconnectGuiInput(subs SignalSubscribers, fn ControlGuiInputSignalFn) {
	sig := StringNameFromStr("gui_input")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ControlMouseEnteredSignalFn func()

func (me *Control) ConnectMouseEntered(subs SignalSubscribers, fn ControlMouseEnteredSignalFn) {
	sig := StringNameFromStr("mouse_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Control) DisconnectMouseEntered(subs SignalSubscribers, fn ControlMouseEnteredSignalFn) {
	sig := StringNameFromStr("mouse_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ControlMouseExitedSignalFn func()

func (me *Control) ConnectMouseExited(subs SignalSubscribers, fn ControlMouseExitedSignalFn) {
	sig := StringNameFromStr("mouse_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Control) DisconnectMouseExited(subs SignalSubscribers, fn ControlMouseExitedSignalFn) {
	sig := StringNameFromStr("mouse_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ControlFocusEnteredSignalFn func()

func (me *Control) ConnectFocusEntered(subs SignalSubscribers, fn ControlFocusEnteredSignalFn) {
	sig := StringNameFromStr("focus_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Control) DisconnectFocusEntered(subs SignalSubscribers, fn ControlFocusEnteredSignalFn) {
	sig := StringNameFromStr("focus_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ControlFocusExitedSignalFn func()

func (me *Control) ConnectFocusExited(subs SignalSubscribers, fn ControlFocusExitedSignalFn) {
	sig := StringNameFromStr("focus_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Control) DisconnectFocusExited(subs SignalSubscribers, fn ControlFocusExitedSignalFn) {
	sig := StringNameFromStr("focus_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ControlSizeFlagsChangedSignalFn func()

func (me *Control) ConnectSizeFlagsChanged(subs SignalSubscribers, fn ControlSizeFlagsChangedSignalFn) {
	sig := StringNameFromStr("size_flags_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Control) DisconnectSizeFlagsChanged(subs SignalSubscribers, fn ControlSizeFlagsChangedSignalFn) {
	sig := StringNameFromStr("size_flags_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ControlMinimumSizeChangedSignalFn func()

func (me *Control) ConnectMinimumSizeChanged(subs SignalSubscribers, fn ControlMinimumSizeChangedSignalFn) {
	sig := StringNameFromStr("minimum_size_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Control) DisconnectMinimumSizeChanged(subs SignalSubscribers, fn ControlMinimumSizeChangedSignalFn) {
	sig := StringNameFromStr("minimum_size_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type ControlThemeChangedSignalFn func()

func (me *Control) ConnectThemeChanged(subs SignalSubscribers, fn ControlThemeChangedSignalFn) {
	sig := StringNameFromStr("theme_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Control) DisconnectThemeChanged(subs SignalSubscribers, fn ControlThemeChangedSignalFn) {
	sig := StringNameFromStr("theme_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
