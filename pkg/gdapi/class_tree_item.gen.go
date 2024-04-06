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

type ptrsForTreeItemList struct {
	fnSetCellMode                          gdc.MethodBindPtr
	fnGetCellMode                          gdc.MethodBindPtr
	fnSetEditMultiline                     gdc.MethodBindPtr
	fnIsEditMultiline                      gdc.MethodBindPtr
	fnSetChecked                           gdc.MethodBindPtr
	fnSetIndeterminate                     gdc.MethodBindPtr
	fnIsChecked                            gdc.MethodBindPtr
	fnIsIndeterminate                      gdc.MethodBindPtr
	fnPropagateCheck                       gdc.MethodBindPtr
	fnSetText                              gdc.MethodBindPtr
	fnGetText                              gdc.MethodBindPtr
	fnSetTextDirection                     gdc.MethodBindPtr
	fnGetTextDirection                     gdc.MethodBindPtr
	fnSetAutowrapMode                      gdc.MethodBindPtr
	fnGetAutowrapMode                      gdc.MethodBindPtr
	fnSetTextOverrunBehavior               gdc.MethodBindPtr
	fnGetTextOverrunBehavior               gdc.MethodBindPtr
	fnSetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnGetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnSetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnGetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnSetLanguage                          gdc.MethodBindPtr
	fnGetLanguage                          gdc.MethodBindPtr
	fnSetSuffix                            gdc.MethodBindPtr
	fnGetSuffix                            gdc.MethodBindPtr
	fnSetIcon                              gdc.MethodBindPtr
	fnGetIcon                              gdc.MethodBindPtr
	fnSetIconRegion                        gdc.MethodBindPtr
	fnGetIconRegion                        gdc.MethodBindPtr
	fnSetIconMaxWidth                      gdc.MethodBindPtr
	fnGetIconMaxWidth                      gdc.MethodBindPtr
	fnSetIconModulate                      gdc.MethodBindPtr
	fnGetIconModulate                      gdc.MethodBindPtr
	fnSetRange                             gdc.MethodBindPtr
	fnGetRange                             gdc.MethodBindPtr
	fnSetRangeConfig                       gdc.MethodBindPtr
	fnGetRangeConfig                       gdc.MethodBindPtr
	fnSetMetadata                          gdc.MethodBindPtr
	fnGetMetadata                          gdc.MethodBindPtr
	fnSetCustomDraw                        gdc.MethodBindPtr
	fnSetCollapsed                         gdc.MethodBindPtr
	fnIsCollapsed                          gdc.MethodBindPtr
	fnSetCollapsedRecursive                gdc.MethodBindPtr
	fnIsAnyCollapsed                       gdc.MethodBindPtr
	fnSetVisible                           gdc.MethodBindPtr
	fnIsVisible                            gdc.MethodBindPtr
	fnUncollapseTree                       gdc.MethodBindPtr
	fnSetCustomMinimumHeight               gdc.MethodBindPtr
	fnGetCustomMinimumHeight               gdc.MethodBindPtr
	fnSetSelectable                        gdc.MethodBindPtr
	fnIsSelectable                         gdc.MethodBindPtr
	fnIsSelected                           gdc.MethodBindPtr
	fnSelect                               gdc.MethodBindPtr
	fnDeselect                             gdc.MethodBindPtr
	fnSetEditable                          gdc.MethodBindPtr
	fnIsEditable                           gdc.MethodBindPtr
	fnSetCustomColor                       gdc.MethodBindPtr
	fnGetCustomColor                       gdc.MethodBindPtr
	fnClearCustomColor                     gdc.MethodBindPtr
	fnSetCustomFont                        gdc.MethodBindPtr
	fnGetCustomFont                        gdc.MethodBindPtr
	fnSetCustomFontSize                    gdc.MethodBindPtr
	fnGetCustomFontSize                    gdc.MethodBindPtr
	fnSetCustomBgColor                     gdc.MethodBindPtr
	fnClearCustomBgColor                   gdc.MethodBindPtr
	fnGetCustomBgColor                     gdc.MethodBindPtr
	fnSetCustomAsButton                    gdc.MethodBindPtr
	fnIsCustomSetAsButton                  gdc.MethodBindPtr
	fnAddButton                            gdc.MethodBindPtr
	fnGetButtonCount                       gdc.MethodBindPtr
	fnGetButtonTooltipText                 gdc.MethodBindPtr
	fnGetButtonId                          gdc.MethodBindPtr
	fnGetButtonById                        gdc.MethodBindPtr
	fnGetButton                            gdc.MethodBindPtr
	fnSetButtonTooltipText                 gdc.MethodBindPtr
	fnSetButton                            gdc.MethodBindPtr
	fnEraseButton                          gdc.MethodBindPtr
	fnSetButtonDisabled                    gdc.MethodBindPtr
	fnSetButtonColor                       gdc.MethodBindPtr
	fnIsButtonDisabled                     gdc.MethodBindPtr
	fnSetTooltipText                       gdc.MethodBindPtr
	fnGetTooltipText                       gdc.MethodBindPtr
	fnSetTextAlignment                     gdc.MethodBindPtr
	fnGetTextAlignment                     gdc.MethodBindPtr
	fnSetExpandRight                       gdc.MethodBindPtr
	fnGetExpandRight                       gdc.MethodBindPtr
	fnSetDisableFolding                    gdc.MethodBindPtr
	fnIsFoldingDisabled                    gdc.MethodBindPtr
	fnCreateChild                          gdc.MethodBindPtr
	fnAddChild                             gdc.MethodBindPtr
	fnRemoveChild                          gdc.MethodBindPtr
	fnGetTree                              gdc.MethodBindPtr
	fnGetNext                              gdc.MethodBindPtr
	fnGetPrev                              gdc.MethodBindPtr
	fnGetParent                            gdc.MethodBindPtr
	fnGetFirstChild                        gdc.MethodBindPtr
	fnGetNextInTree                        gdc.MethodBindPtr
	fnGetPrevInTree                        gdc.MethodBindPtr
	fnGetNextVisible                       gdc.MethodBindPtr
	fnGetPrevVisible                       gdc.MethodBindPtr
	fnGetChild                             gdc.MethodBindPtr
	fnGetChildCount                        gdc.MethodBindPtr
	fnGetChildren                          gdc.MethodBindPtr
	fnGetIndex                             gdc.MethodBindPtr
	fnMoveBefore                           gdc.MethodBindPtr
	fnMoveAfter                            gdc.MethodBindPtr
	fnCallRecursive                        gdc.MethodBindPtr
}

var ptrsForTreeItem ptrsForTreeItemList

func initTreeItemPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TreeItem")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_cell_mode")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCellMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 289920701))
	}
	{
		methodName := StringNameFromStr("get_cell_mode")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetCellMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3406114978))
	}
	{
		methodName := StringNameFromStr("set_edit_multiline")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetEditMultiline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_edit_multiline")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsEditMultiline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_checked")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("set_indeterminate")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetIndeterminate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_checked")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("is_indeterminate")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsIndeterminate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("propagate_check")
		defer methodName.Destroy()
		ptrsForTreeItem.fnPropagateCheck = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 972357352))
	}
	{
		methodName := StringNameFromStr("set_text")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_text")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_text_direction")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1707680378))
	}
	{
		methodName := StringNameFromStr("get_text_direction")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4235602388))
	}
	{
		methodName := StringNameFromStr("set_autowrap_mode")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3633006561))
	}
	{
		methodName := StringNameFromStr("get_autowrap_mode")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2902757236))
	}
	{
		methodName := StringNameFromStr("set_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1940772195))
	}
	{
		methodName := StringNameFromStr("get_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3782727860))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 868756907))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3377823772))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 537221740))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 663333327))
	}
	{
		methodName := StringNameFromStr("set_language")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_language")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_suffix")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetSuffix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_suffix")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetSuffix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_icon")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
	}
	{
		methodName := StringNameFromStr("get_icon")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
	}
	{
		methodName := StringNameFromStr("set_icon_region")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetIconRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1356297692))
	}
	{
		methodName := StringNameFromStr("get_icon_region")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetIconRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3327874267))
	}
	{
		methodName := StringNameFromStr("set_icon_max_width")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetIconMaxWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_icon_max_width")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetIconMaxWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_icon_modulate")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetIconModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_icon_modulate")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetIconModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("set_range")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_range")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_range_config")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetRangeConfig = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1547181014))
	}
	{
		methodName := StringNameFromStr("get_range_config")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetRangeConfig = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3554694381))
	}
	{
		methodName := StringNameFromStr("set_metadata")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
	}
	{
		methodName := StringNameFromStr("get_metadata")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
	}
	{
		methodName := StringNameFromStr("set_custom_draw")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCustomDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 272420368))
	}
	{
		methodName := StringNameFromStr("set_collapsed")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCollapsed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collapsed")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsCollapsed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_collapsed_recursive")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCollapsedRecursive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_any_collapsed")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsAnyCollapsed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2595650253))
	}
	{
		methodName := StringNameFromStr("set_visible")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_visible")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("uncollapse_tree")
		defer methodName.Destroy()
		ptrsForTreeItem.fnUncollapseTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_custom_minimum_height")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCustomMinimumHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_custom_minimum_height")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetCustomMinimumHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_selectable")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetSelectable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_selectable")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsSelectable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("is_selected")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3067735520))
	}
	{
		methodName := StringNameFromStr("select")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSelect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("deselect")
		defer methodName.Destroy()
		ptrsForTreeItem.fnDeselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_editable")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_editable")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3067735520))
	}
	{
		methodName := StringNameFromStr("set_custom_color")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_custom_color")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("clear_custom_color")
		defer methodName.Destroy()
		ptrsForTreeItem.fnClearCustomColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_custom_font")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCustomFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2637609184))
	}
	{
		methodName := StringNameFromStr("get_custom_font")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetCustomFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4244553094))
	}
	{
		methodName := StringNameFromStr("set_custom_font_size")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCustomFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_custom_font_size")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetCustomFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_custom_bg_color")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCustomBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 894174518))
	}
	{
		methodName := StringNameFromStr("clear_custom_bg_color")
		defer methodName.Destroy()
		ptrsForTreeItem.fnClearCustomBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_custom_bg_color")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetCustomBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("set_custom_as_button")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetCustomAsButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_custom_set_as_button")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsCustomSetAsButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("add_button")
		defer methodName.Destroy()
		ptrsForTreeItem.fnAddButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1688223362))
	}
	{
		methodName := StringNameFromStr("get_button_count")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetButtonCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_button_tooltip_text")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetButtonTooltipText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1391810591))
	}
	{
		methodName := StringNameFromStr("get_button_id")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetButtonId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("get_button_by_id")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetButtonById = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("get_button")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2584904275))
	}
	{
		methodName := StringNameFromStr("set_button_tooltip_text")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetButtonTooltipText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2285447957))
	}
	{
		methodName := StringNameFromStr("set_button")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 176101966))
	}
	{
		methodName := StringNameFromStr("erase_button")
		defer methodName.Destroy()
		ptrsForTreeItem.fnEraseButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_button_disabled")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetButtonDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1383440665))
	}
	{
		methodName := StringNameFromStr("set_button_color")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetButtonColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3733378741))
	}
	{
		methodName := StringNameFromStr("is_button_disabled")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsButtonDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}
	{
		methodName := StringNameFromStr("set_tooltip_text")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetTooltipText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_tooltip_text")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetTooltipText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_text_alignment")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetTextAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3276431499))
	}
	{
		methodName := StringNameFromStr("get_text_alignment")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetTextAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4171562184))
	}
	{
		methodName := StringNameFromStr("set_expand_right")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetExpandRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_expand_right")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetExpandRight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_disable_folding")
		defer methodName.Destroy()
		ptrsForTreeItem.fnSetDisableFolding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_folding_disabled")
		defer methodName.Destroy()
		ptrsForTreeItem.fnIsFoldingDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("create_child")
		defer methodName.Destroy()
		ptrsForTreeItem.fnCreateChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 954243986))
	}
	{
		methodName := StringNameFromStr("add_child")
		defer methodName.Destroy()
		ptrsForTreeItem.fnAddChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1819951137))
	}
	{
		methodName := StringNameFromStr("remove_child")
		defer methodName.Destroy()
		ptrsForTreeItem.fnRemoveChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1819951137))
	}
	{
		methodName := StringNameFromStr("get_tree")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2243340556))
	}
	{
		methodName := StringNameFromStr("get_next")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetNext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1514277247))
	}
	{
		methodName := StringNameFromStr("get_prev")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetPrev = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2768121250))
	}
	{
		methodName := StringNameFromStr("get_parent")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1514277247))
	}
	{
		methodName := StringNameFromStr("get_first_child")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetFirstChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1514277247))
	}
	{
		methodName := StringNameFromStr("get_next_in_tree")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetNextInTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1666920593))
	}
	{
		methodName := StringNameFromStr("get_prev_in_tree")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetPrevInTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1666920593))
	}
	{
		methodName := StringNameFromStr("get_next_visible")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetNextVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1666920593))
	}
	{
		methodName := StringNameFromStr("get_prev_visible")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetPrevVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1666920593))
	}
	{
		methodName := StringNameFromStr("get_child")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 306700752))
	}
	{
		methodName := StringNameFromStr("get_child_count")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetChildCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_children")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetChildren = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("get_index")
		defer methodName.Destroy()
		ptrsForTreeItem.fnGetIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("move_before")
		defer methodName.Destroy()
		ptrsForTreeItem.fnMoveBefore = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1819951137))
	}
	{
		methodName := StringNameFromStr("move_after")
		defer methodName.Destroy()
		ptrsForTreeItem.fnMoveAfter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1819951137))
	}
	{
		methodName := StringNameFromStr("call_recursive")
		defer methodName.Destroy()
		ptrsForTreeItem.fnCallRecursive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2866548813))
	}
}

type TreeItem struct {
	Object
}

func (me *TreeItem) BaseClass() string {
	return "TreeItem"
}

func NewTreeItem() *TreeItem {
	str := StringNameFromStr("TreeItem") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TreeItem{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TreeItemTreeCellMode int

const (
	TreeItemTreeCellModeCellModeString TreeItemTreeCellMode = 0
	TreeItemTreeCellModeCellModeCheck  TreeItemTreeCellMode = 1
	TreeItemTreeCellModeCellModeRange  TreeItemTreeCellMode = 2
	TreeItemTreeCellModeCellModeIcon   TreeItemTreeCellMode = 3
	TreeItemTreeCellModeCellModeCustom TreeItemTreeCellMode = 4
)

func (me *TreeItem) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TreeItem) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TreeItem) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TreeItem) SetCellMode(column int64, mode TreeItemTreeCellMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCellMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetCellMode(column int64) TreeItemTreeCellMode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TreeItemTreeCellMode
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetCellMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TreeItem) SetEditMultiline(column int64, multiline bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&multiline)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetEditMultiline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsEditMultiline(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsEditMultiline), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetChecked(column int64, checked bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&checked)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetChecked), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) SetIndeterminate(column int64, indeterminate bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&indeterminate)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetIndeterminate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsChecked(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsChecked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) IsIndeterminate(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsIndeterminate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) PropagateCheck(column int64, emit_signal bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&emit_signal)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnPropagateCheck), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) SetText(column int64, text String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetText(column int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetTextDirection(column int64, direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetTextDirection(column int64) ControlTextDirection {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TreeItem) SetAutowrapMode(column int64, autowrap_mode TextServerAutowrapMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&autowrap_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetAutowrapMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetAutowrapMode(column int64) TextServerAutowrapMode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerAutowrapMode
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetAutowrapMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TreeItem) SetTextOverrunBehavior(column int64, overrun_behavior TextServerOverrunBehavior) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&overrun_behavior)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetTextOverrunBehavior(column int64) TextServerOverrunBehavior {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerOverrunBehavior
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TreeItem) SetStructuredTextBidiOverride(column int64, parser TextServerStructuredTextParser) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&parser)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetStructuredTextBidiOverride(column int64) TextServerStructuredTextParser {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerStructuredTextParser
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TreeItem) SetStructuredTextBidiOverrideOptions(column int64, args Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), args.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetStructuredTextBidiOverrideOptions(column int64) Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetLanguage(column int64, language String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetLanguage(column int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetSuffix(column int64, text String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetSuffix), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetSuffix(column int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetSuffix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetIcon(column int64, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetIcon(column int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetIconRegion(column int64, region Rect2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetIconRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetIconRegion(column int64) Rect2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetIconRegion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetIconMaxWidth(column int64, width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetIconMaxWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetIconMaxWidth(column int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetIconMaxWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetIconModulate(column int64, modulate Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetIconModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetIconModulate(column int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetIconModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetRange(column int64, value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetRange(column int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetRangeConfig(column int64, min float64, max float64, step float64, expr bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&min), gdc.ConstTypePtr(&max), gdc.ConstTypePtr(&step), gdc.ConstTypePtr(&expr)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetRangeConfig), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetRangeConfig(column int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetRangeConfig), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetMetadata(column int64, meta Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), meta.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetMetadata), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetMetadata(column int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetMetadata), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetCustomDraw(column int64, object Object, callback StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), object.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCustomDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) SetCollapsed(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCollapsed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsCollapsed() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsCollapsed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetCollapsedRecursive(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCollapsedRecursive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsAnyCollapsed(only_visible bool) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&only_visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&only_visible)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsAnyCollapsed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetVisible(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) UncollapseTree() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnUncollapseTree), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) SetCustomMinimumHeight(height int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCustomMinimumHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetCustomMinimumHeight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetCustomMinimumHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetSelectable(column int64, selectable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&selectable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetSelectable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsSelectable(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsSelectable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) IsSelected(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) Select(column int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSelect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) Deselect(column int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnDeselect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) SetEditable(column int64, enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetEditable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsEditable(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsEditable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetCustomColor(column int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCustomColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetCustomColor(column int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetCustomColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) ClearCustomColor(column int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnClearCustomColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) SetCustomFont(column int64, font Font) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), font.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCustomFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetCustomFont(column int64) Font {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFont()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetCustomFont), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetCustomFontSize(column int64, font_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCustomFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetCustomFontSize(column int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetCustomFontSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetCustomBgColor(column int64, color Color, just_outline bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), color.AsCTypePtr(), gdc.ConstTypePtr(&just_outline)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCustomBgColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) ClearCustomBgColor(column int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnClearCustomBgColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetCustomBgColor(column int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetCustomBgColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetCustomAsButton(column int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetCustomAsButton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsCustomSetAsButton(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsCustomSetAsButton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) AddButton(column int64, button Texture2D, id int64, disabled bool, tooltip_text String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), button.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&disabled), tooltip_text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnAddButton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetButtonCount(column int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetButtonCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) GetButtonTooltipText(column int64, button_index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&column)
	pinner.Pin(&button_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetButtonTooltipText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetButtonId(column int64, button_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&column)
	pinner.Pin(&button_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetButtonId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) GetButtonById(column int64, id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&column)
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetButtonById), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) GetButton(column int64, button_index int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&column)
	pinner.Pin(&button_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetButton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetButtonTooltipText(column int64, button_index int64, tooltip String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), tooltip.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetButtonTooltipText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) SetButton(column int64, button_index int64, button Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), button.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetButton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) EraseButton(column int64, button_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnEraseButton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) SetButtonDisabled(column int64, button_index int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetButtonDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) SetButtonColor(column int64, button_index int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetButtonColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsButtonDisabled(column int64, button_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&button_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)
	pinner.Pin(&button_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsButtonDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetTooltipText(column int64, tooltip String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), tooltip.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetTooltipText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetTooltipText(column int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetTooltipText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) SetTextAlignment(column int64, text_alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&text_alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetTextAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetTextAlignment(column int64) HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetTextAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TreeItem) SetExpandRight(column int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetExpandRight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetExpandRight(column int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetExpandRight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) SetDisableFolding(disable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnSetDisableFolding), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) IsFoldingDisabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnIsFoldingDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) CreateChild(index int64) TreeItem {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnCreateChild), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) AddChild(child TreeItem) {
	cargs := []gdc.ConstTypePtr{child.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnAddChild), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) RemoveChild(child TreeItem) {
	cargs := []gdc.ConstTypePtr{child.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnRemoveChild), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) GetTree() Tree {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTree()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetTree), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetNext() TreeItem {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetNext), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetPrev() TreeItem {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetPrev), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetParent() TreeItem {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetFirstChild() TreeItem {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetFirstChild), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetNextInTree(wrap bool) TreeItem {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()
	pinner.Pin(&wrap)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetNextInTree), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetPrevInTree(wrap bool) TreeItem {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()
	pinner.Pin(&wrap)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetPrevInTree), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetNextVisible(wrap bool) TreeItem {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()
	pinner.Pin(&wrap)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetNextVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetPrevVisible(wrap bool) TreeItem {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&wrap)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()
	pinner.Pin(&wrap)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetPrevVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetChild(index int64) TreeItem {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTreeItem()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetChild), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TreeItem) GetChildCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetChildCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) GetChildren() []TreeItem {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetChildren), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[TreeItem](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *TreeItem) GetIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnGetIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TreeItem) MoveBefore(item TreeItem) {
	cargs := []gdc.ConstTypePtr{item.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnMoveBefore), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) MoveAfter(item TreeItem) {
	cargs := []gdc.ConstTypePtr{item.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTreeItem.fnMoveAfter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TreeItem) CallRecursive(method StringName, varargs ...Variant) {
	cargs := make([]gdc.ConstVariantPtr, 0, 1+len(varargs))
	var0 := method.AsVariant()
	defer var0.Destroy()
	cargs = append(cargs, var0.AsCPtr())
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}

	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForTreeItem.fnCallRecursive), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), nil, cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return
	}

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
