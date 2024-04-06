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

type ptrsForDisplayServerList struct {
	fnHasFeature                        gdc.MethodBindPtr
	fnGetName                           gdc.MethodBindPtr
	fnGlobalMenuSetPopupCallbacks       gdc.MethodBindPtr
	fnGlobalMenuAddSubmenuItem          gdc.MethodBindPtr
	fnGlobalMenuAddItem                 gdc.MethodBindPtr
	fnGlobalMenuAddCheckItem            gdc.MethodBindPtr
	fnGlobalMenuAddIconItem             gdc.MethodBindPtr
	fnGlobalMenuAddIconCheckItem        gdc.MethodBindPtr
	fnGlobalMenuAddRadioCheckItem       gdc.MethodBindPtr
	fnGlobalMenuAddIconRadioCheckItem   gdc.MethodBindPtr
	fnGlobalMenuAddMultistateItem       gdc.MethodBindPtr
	fnGlobalMenuAddSeparator            gdc.MethodBindPtr
	fnGlobalMenuGetItemIndexFromText    gdc.MethodBindPtr
	fnGlobalMenuGetItemIndexFromTag     gdc.MethodBindPtr
	fnGlobalMenuIsItemChecked           gdc.MethodBindPtr
	fnGlobalMenuIsItemCheckable         gdc.MethodBindPtr
	fnGlobalMenuIsItemRadioCheckable    gdc.MethodBindPtr
	fnGlobalMenuGetItemCallback         gdc.MethodBindPtr
	fnGlobalMenuGetItemKeyCallback      gdc.MethodBindPtr
	fnGlobalMenuGetItemTag              gdc.MethodBindPtr
	fnGlobalMenuGetItemText             gdc.MethodBindPtr
	fnGlobalMenuGetItemSubmenu          gdc.MethodBindPtr
	fnGlobalMenuGetItemAccelerator      gdc.MethodBindPtr
	fnGlobalMenuIsItemDisabled          gdc.MethodBindPtr
	fnGlobalMenuIsItemHidden            gdc.MethodBindPtr
	fnGlobalMenuGetItemTooltip          gdc.MethodBindPtr
	fnGlobalMenuGetItemState            gdc.MethodBindPtr
	fnGlobalMenuGetItemMaxStates        gdc.MethodBindPtr
	fnGlobalMenuGetItemIcon             gdc.MethodBindPtr
	fnGlobalMenuGetItemIndentationLevel gdc.MethodBindPtr
	fnGlobalMenuSetItemChecked          gdc.MethodBindPtr
	fnGlobalMenuSetItemCheckable        gdc.MethodBindPtr
	fnGlobalMenuSetItemRadioCheckable   gdc.MethodBindPtr
	fnGlobalMenuSetItemCallback         gdc.MethodBindPtr
	fnGlobalMenuSetItemHoverCallbacks   gdc.MethodBindPtr
	fnGlobalMenuSetItemKeyCallback      gdc.MethodBindPtr
	fnGlobalMenuSetItemTag              gdc.MethodBindPtr
	fnGlobalMenuSetItemText             gdc.MethodBindPtr
	fnGlobalMenuSetItemSubmenu          gdc.MethodBindPtr
	fnGlobalMenuSetItemAccelerator      gdc.MethodBindPtr
	fnGlobalMenuSetItemDisabled         gdc.MethodBindPtr
	fnGlobalMenuSetItemHidden           gdc.MethodBindPtr
	fnGlobalMenuSetItemTooltip          gdc.MethodBindPtr
	fnGlobalMenuSetItemState            gdc.MethodBindPtr
	fnGlobalMenuSetItemMaxStates        gdc.MethodBindPtr
	fnGlobalMenuSetItemIcon             gdc.MethodBindPtr
	fnGlobalMenuSetItemIndentationLevel gdc.MethodBindPtr
	fnGlobalMenuGetItemCount            gdc.MethodBindPtr
	fnGlobalMenuRemoveItem              gdc.MethodBindPtr
	fnGlobalMenuClear                   gdc.MethodBindPtr
	fnTtsIsSpeaking                     gdc.MethodBindPtr
	fnTtsIsPaused                       gdc.MethodBindPtr
	fnTtsGetVoices                      gdc.MethodBindPtr
	fnTtsGetVoicesForLanguage           gdc.MethodBindPtr
	fnTtsSpeak                          gdc.MethodBindPtr
	fnTtsPause                          gdc.MethodBindPtr
	fnTtsResume                         gdc.MethodBindPtr
	fnTtsStop                           gdc.MethodBindPtr
	fnTtsSetUtteranceCallback           gdc.MethodBindPtr
	fnIsDarkModeSupported               gdc.MethodBindPtr
	fnIsDarkMode                        gdc.MethodBindPtr
	fnGetAccentColor                    gdc.MethodBindPtr
	fnMouseSetMode                      gdc.MethodBindPtr
	fnMouseGetMode                      gdc.MethodBindPtr
	fnWarpMouse                         gdc.MethodBindPtr
	fnMouseGetPosition                  gdc.MethodBindPtr
	fnMouseGetButtonState               gdc.MethodBindPtr
	fnClipboardSet                      gdc.MethodBindPtr
	fnClipboardGet                      gdc.MethodBindPtr
	fnClipboardGetImage                 gdc.MethodBindPtr
	fnClipboardHas                      gdc.MethodBindPtr
	fnClipboardHasImage                 gdc.MethodBindPtr
	fnClipboardSetPrimary               gdc.MethodBindPtr
	fnClipboardGetPrimary               gdc.MethodBindPtr
	fnGetDisplayCutouts                 gdc.MethodBindPtr
	fnGetDisplaySafeArea                gdc.MethodBindPtr
	fnGetScreenCount                    gdc.MethodBindPtr
	fnGetPrimaryScreen                  gdc.MethodBindPtr
	fnGetKeyboardFocusScreen            gdc.MethodBindPtr
	fnGetScreenFromRect                 gdc.MethodBindPtr
	fnScreenGetPosition                 gdc.MethodBindPtr
	fnScreenGetSize                     gdc.MethodBindPtr
	fnScreenGetUsableRect               gdc.MethodBindPtr
	fnScreenGetDpi                      gdc.MethodBindPtr
	fnScreenGetScale                    gdc.MethodBindPtr
	fnIsTouchscreenAvailable            gdc.MethodBindPtr
	fnScreenGetMaxScale                 gdc.MethodBindPtr
	fnScreenGetRefreshRate              gdc.MethodBindPtr
	fnScreenGetPixel                    gdc.MethodBindPtr
	fnScreenGetImage                    gdc.MethodBindPtr
	fnScreenSetOrientation              gdc.MethodBindPtr
	fnScreenGetOrientation              gdc.MethodBindPtr
	fnScreenSetKeepOn                   gdc.MethodBindPtr
	fnScreenIsKeptOn                    gdc.MethodBindPtr
	fnGetWindowList                     gdc.MethodBindPtr
	fnGetWindowAtScreenPosition         gdc.MethodBindPtr
	fnWindowGetNativeHandle             gdc.MethodBindPtr
	fnWindowGetActivePopup              gdc.MethodBindPtr
	fnWindowSetPopupSafeRect            gdc.MethodBindPtr
	fnWindowGetPopupSafeRect            gdc.MethodBindPtr
	fnWindowSetTitle                    gdc.MethodBindPtr
	fnWindowGetTitleSize                gdc.MethodBindPtr
	fnWindowSetMousePassthrough         gdc.MethodBindPtr
	fnWindowGetCurrentScreen            gdc.MethodBindPtr
	fnWindowSetCurrentScreen            gdc.MethodBindPtr
	fnWindowGetPosition                 gdc.MethodBindPtr
	fnWindowGetPositionWithDecorations  gdc.MethodBindPtr
	fnWindowSetPosition                 gdc.MethodBindPtr
	fnWindowGetSize                     gdc.MethodBindPtr
	fnWindowSetSize                     gdc.MethodBindPtr
	fnWindowSetRectChangedCallback      gdc.MethodBindPtr
	fnWindowSetWindowEventCallback      gdc.MethodBindPtr
	fnWindowSetInputEventCallback       gdc.MethodBindPtr
	fnWindowSetInputTextCallback        gdc.MethodBindPtr
	fnWindowSetDropFilesCallback        gdc.MethodBindPtr
	fnWindowGetAttachedInstanceId       gdc.MethodBindPtr
	fnWindowGetMaxSize                  gdc.MethodBindPtr
	fnWindowSetMaxSize                  gdc.MethodBindPtr
	fnWindowGetMinSize                  gdc.MethodBindPtr
	fnWindowSetMinSize                  gdc.MethodBindPtr
	fnWindowGetSizeWithDecorations      gdc.MethodBindPtr
	fnWindowGetMode                     gdc.MethodBindPtr
	fnWindowSetMode                     gdc.MethodBindPtr
	fnWindowSetFlag                     gdc.MethodBindPtr
	fnWindowGetFlag                     gdc.MethodBindPtr
	fnWindowSetWindowButtonsOffset      gdc.MethodBindPtr
	fnWindowGetSafeTitleMargins         gdc.MethodBindPtr
	fnWindowRequestAttention            gdc.MethodBindPtr
	fnWindowMoveToForeground            gdc.MethodBindPtr
	fnWindowIsFocused                   gdc.MethodBindPtr
	fnWindowCanDraw                     gdc.MethodBindPtr
	fnWindowSetTransient                gdc.MethodBindPtr
	fnWindowSetExclusive                gdc.MethodBindPtr
	fnWindowSetImeActive                gdc.MethodBindPtr
	fnWindowSetImePosition              gdc.MethodBindPtr
	fnWindowSetVsyncMode                gdc.MethodBindPtr
	fnWindowGetVsyncMode                gdc.MethodBindPtr
	fnWindowIsMaximizeAllowed           gdc.MethodBindPtr
	fnWindowMaximizeOnTitleDblClick     gdc.MethodBindPtr
	fnWindowMinimizeOnTitleDblClick     gdc.MethodBindPtr
	fnImeGetSelection                   gdc.MethodBindPtr
	fnImeGetText                        gdc.MethodBindPtr
	fnVirtualKeyboardShow               gdc.MethodBindPtr
	fnVirtualKeyboardHide               gdc.MethodBindPtr
	fnVirtualKeyboardGetHeight          gdc.MethodBindPtr
	fnCursorSetShape                    gdc.MethodBindPtr
	fnCursorGetShape                    gdc.MethodBindPtr
	fnCursorSetCustomImage              gdc.MethodBindPtr
	fnGetSwapCancelOk                   gdc.MethodBindPtr
	fnEnableForStealingFocus            gdc.MethodBindPtr
	fnDialogShow                        gdc.MethodBindPtr
	fnDialogInputText                   gdc.MethodBindPtr
	fnFileDialogShow                    gdc.MethodBindPtr
	fnKeyboardGetLayoutCount            gdc.MethodBindPtr
	fnKeyboardGetCurrentLayout          gdc.MethodBindPtr
	fnKeyboardSetCurrentLayout          gdc.MethodBindPtr
	fnKeyboardGetLayoutLanguage         gdc.MethodBindPtr
	fnKeyboardGetLayoutName             gdc.MethodBindPtr
	fnKeyboardGetKeycodeFromPhysical    gdc.MethodBindPtr
	fnKeyboardGetLabelFromPhysical      gdc.MethodBindPtr
	fnProcessEvents                     gdc.MethodBindPtr
	fnForceProcessAndDropEvents         gdc.MethodBindPtr
	fnSetNativeIcon                     gdc.MethodBindPtr
	fnSetIcon                           gdc.MethodBindPtr
	fnTabletGetDriverCount              gdc.MethodBindPtr
	fnTabletGetDriverName               gdc.MethodBindPtr
	fnTabletGetCurrentDriver            gdc.MethodBindPtr
	fnTabletSetCurrentDriver            gdc.MethodBindPtr
}

var ptrsForDisplayServer ptrsForDisplayServerList

func initDisplayServerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("DisplayServer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("has_feature")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnHasFeature = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334065950))
	}
	{
		methodName := StringNameFromStr("get_name")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("global_menu_set_popup_callbacks")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetPopupCallbacks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3893727526))
	}
	{
		methodName := StringNameFromStr("global_menu_add_submenu_item")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuAddSubmenuItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2828985934))
	}
	{
		methodName := StringNameFromStr("global_menu_add_item")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuAddItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3401266716))
	}
	{
		methodName := StringNameFromStr("global_menu_add_check_item")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuAddCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3401266716))
	}
	{
		methodName := StringNameFromStr("global_menu_add_icon_item")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuAddIconItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4245856523))
	}
	{
		methodName := StringNameFromStr("global_menu_add_icon_check_item")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuAddIconCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4245856523))
	}
	{
		methodName := StringNameFromStr("global_menu_add_radio_check_item")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuAddRadioCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3401266716))
	}
	{
		methodName := StringNameFromStr("global_menu_add_icon_radio_check_item")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuAddIconRadioCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4245856523))
	}
	{
		methodName := StringNameFromStr("global_menu_add_multistate_item")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuAddMultistateItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3431222859))
	}
	{
		methodName := StringNameFromStr("global_menu_add_separator")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuAddSeparator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3214812433))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_index_from_text")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemIndexFromText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878152881))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_index_from_tag")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemIndexFromTag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2941063483))
	}
	{
		methodName := StringNameFromStr("global_menu_is_item_checked")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuIsItemChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3511468594))
	}
	{
		methodName := StringNameFromStr("global_menu_is_item_checkable")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuIsItemCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3511468594))
	}
	{
		methodName := StringNameFromStr("global_menu_is_item_radio_checkable")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuIsItemRadioCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3511468594))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 748666903))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_key_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemKeyCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 748666903))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_tag")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemTag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 330672633))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_text")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 591067909))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_submenu")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemSubmenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 591067909))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_accelerator")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemAccelerator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 936065394))
	}
	{
		methodName := StringNameFromStr("global_menu_is_item_disabled")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuIsItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3511468594))
	}
	{
		methodName := StringNameFromStr("global_menu_is_item_hidden")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuIsItemHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3511468594))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_tooltip")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 591067909))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_state")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3422818498))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_max_states")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemMaxStates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3422818498))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_icon")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3591713183))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_indentation_level")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemIndentationLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3422818498))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_checked")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4108344793))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_checkable")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4108344793))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_radio_checkable")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemRadioCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4108344793))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3809915389))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_hover_callbacks")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemHoverCallbacks = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3809915389))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_key_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemKeyCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3809915389))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_tag")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemTag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 453659863))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_text")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 965966136))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_submenu")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemSubmenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 965966136))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_accelerator")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemAccelerator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 566943293))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_disabled")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4108344793))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_hidden")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4108344793))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_tooltip")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 965966136))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_state")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3474840532))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_max_states")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemMaxStates = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3474840532))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_icon")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3201338066))
	}
	{
		methodName := StringNameFromStr("global_menu_set_item_indentation_level")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuSetItemIndentationLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3474840532))
	}
	{
		methodName := StringNameFromStr("global_menu_get_item_count")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuGetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("global_menu_remove_item")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuRemoveItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2956805083))
	}
	{
		methodName := StringNameFromStr("global_menu_clear")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGlobalMenuClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("tts_is_speaking")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTtsIsSpeaking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("tts_is_paused")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTtsIsPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("tts_get_voices")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTtsGetVoices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("tts_get_voices_for_language")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTtsGetVoicesForLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291131558))
	}
	{
		methodName := StringNameFromStr("tts_speak")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTtsSpeak = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 903992738))
	}
	{
		methodName := StringNameFromStr("tts_pause")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTtsPause = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("tts_resume")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTtsResume = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("tts_stop")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTtsStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("tts_set_utterance_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTtsSetUtteranceCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 109679083))
	}
	{
		methodName := StringNameFromStr("is_dark_mode_supported")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnIsDarkModeSupported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_dark_mode")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnIsDarkMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_accent_color")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetAccentColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("mouse_set_mode")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnMouseSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 348288463))
	}
	{
		methodName := StringNameFromStr("mouse_get_mode")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnMouseGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1353961651))
	}
	{
		methodName := StringNameFromStr("warp_mouse")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWarpMouse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("mouse_get_position")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnMouseGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("mouse_get_button_state")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnMouseGetButtonState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2512161324))
	}
	{
		methodName := StringNameFromStr("clipboard_set")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnClipboardSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("clipboard_get")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnClipboardGet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("clipboard_get_image")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnClipboardGetImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4190603485))
	}
	{
		methodName := StringNameFromStr("clipboard_has")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnClipboardHas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("clipboard_has_image")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnClipboardHasImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("clipboard_set_primary")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnClipboardSetPrimary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("clipboard_get_primary")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnClipboardGetPrimary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_display_cutouts")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetDisplayCutouts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_display_safe_area")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetDisplaySafeArea = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 410525958))
	}
	{
		methodName := StringNameFromStr("get_screen_count")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetScreenCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_primary_screen")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetPrimaryScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_keyboard_focus_screen")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetKeyboardFocusScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_screen_from_rect")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetScreenFromRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 741354659))
	}
	{
		methodName := StringNameFromStr("screen_get_position")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1725937825))
	}
	{
		methodName := StringNameFromStr("screen_get_size")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1725937825))
	}
	{
		methodName := StringNameFromStr("screen_get_usable_rect")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetUsableRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2439012528))
	}
	{
		methodName := StringNameFromStr("screen_get_dpi")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetDpi = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 181039630))
	}
	{
		methodName := StringNameFromStr("screen_get_scale")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 909105437))
	}
	{
		methodName := StringNameFromStr("is_touchscreen_available")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnIsTouchscreenAvailable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3323674545))
	}
	{
		methodName := StringNameFromStr("screen_get_max_scale")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetMaxScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("screen_get_refresh_rate")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetRefreshRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 909105437))
	}
	{
		methodName := StringNameFromStr("screen_get_pixel")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetPixel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1532707496))
	}
	{
		methodName := StringNameFromStr("screen_get_image")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3813388802))
	}
	{
		methodName := StringNameFromStr("screen_set_orientation")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenSetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2211511631))
	}
	{
		methodName := StringNameFromStr("screen_get_orientation")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenGetOrientation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 133818562))
	}
	{
		methodName := StringNameFromStr("screen_set_keep_on")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenSetKeepOn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("screen_is_kept_on")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnScreenIsKeptOn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_window_list")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetWindowList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
	}
	{
		methodName := StringNameFromStr("get_window_at_screen_position")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetWindowAtScreenPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
	}
	{
		methodName := StringNameFromStr("window_get_native_handle")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetNativeHandle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1096425680))
	}
	{
		methodName := StringNameFromStr("window_get_active_popup")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetActivePopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("window_set_popup_safe_rect")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetPopupSafeRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3317281434))
	}
	{
		methodName := StringNameFromStr("window_get_popup_safe_rect")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetPopupSafeRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2161169500))
	}
	{
		methodName := StringNameFromStr("window_set_title")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 441246282))
	}
	{
		methodName := StringNameFromStr("window_get_title_size")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetTitleSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2925301799))
	}
	{
		methodName := StringNameFromStr("window_set_mouse_passthrough")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetMousePassthrough = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1993637420))
	}
	{
		methodName := StringNameFromStr("window_get_current_screen")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetCurrentScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("window_set_current_screen")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetCurrentScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2230941749))
	}
	{
		methodName := StringNameFromStr("window_get_position")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 763922886))
	}
	{
		methodName := StringNameFromStr("window_get_position_with_decorations")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetPositionWithDecorations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 763922886))
	}
	{
		methodName := StringNameFromStr("window_set_position")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2019273902))
	}
	{
		methodName := StringNameFromStr("window_get_size")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 763922886))
	}
	{
		methodName := StringNameFromStr("window_set_size")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2019273902))
	}
	{
		methodName := StringNameFromStr("window_set_rect_changed_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetRectChangedCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1091192925))
	}
	{
		methodName := StringNameFromStr("window_set_window_event_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetWindowEventCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1091192925))
	}
	{
		methodName := StringNameFromStr("window_set_input_event_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetInputEventCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1091192925))
	}
	{
		methodName := StringNameFromStr("window_set_input_text_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetInputTextCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1091192925))
	}
	{
		methodName := StringNameFromStr("window_set_drop_files_callback")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetDropFilesCallback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1091192925))
	}
	{
		methodName := StringNameFromStr("window_get_attached_instance_id")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetAttachedInstanceId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("window_get_max_size")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 763922886))
	}
	{
		methodName := StringNameFromStr("window_set_max_size")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetMaxSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2019273902))
	}
	{
		methodName := StringNameFromStr("window_get_min_size")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetMinSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 763922886))
	}
	{
		methodName := StringNameFromStr("window_set_min_size")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetMinSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2019273902))
	}
	{
		methodName := StringNameFromStr("window_get_size_with_decorations")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetSizeWithDecorations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 763922886))
	}
	{
		methodName := StringNameFromStr("window_get_mode")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2185728461))
	}
	{
		methodName := StringNameFromStr("window_set_mode")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1319965401))
	}
	{
		methodName := StringNameFromStr("window_set_flag")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 254894155))
	}
	{
		methodName := StringNameFromStr("window_get_flag")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetFlag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 802816991))
	}
	{
		methodName := StringNameFromStr("window_set_window_buttons_offset")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetWindowButtonsOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2019273902))
	}
	{
		methodName := StringNameFromStr("window_get_safe_title_margins")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetSafeTitleMargins = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2295066620))
	}
	{
		methodName := StringNameFromStr("window_request_attention")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowRequestAttention = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
	}
	{
		methodName := StringNameFromStr("window_move_to_foreground")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowMoveToForeground = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
	}
	{
		methodName := StringNameFromStr("window_is_focused")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowIsFocused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1051549951))
	}
	{
		methodName := StringNameFromStr("window_can_draw")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowCanDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1051549951))
	}
	{
		methodName := StringNameFromStr("window_set_transient")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetTransient = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("window_set_exclusive")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetExclusive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("window_set_ime_active")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetImeActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1661950165))
	}
	{
		methodName := StringNameFromStr("window_set_ime_position")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetImePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2019273902))
	}
	{
		methodName := StringNameFromStr("window_set_vsync_mode")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowSetVsyncMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2179333492))
	}
	{
		methodName := StringNameFromStr("window_get_vsync_mode")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowGetVsyncMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 578873795))
	}
	{
		methodName := StringNameFromStr("window_is_maximize_allowed")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowIsMaximizeAllowed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1051549951))
	}
	{
		methodName := StringNameFromStr("window_maximize_on_title_dbl_click")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowMaximizeOnTitleDblClick = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("window_minimize_on_title_dbl_click")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnWindowMinimizeOnTitleDblClick = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("ime_get_selection")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnImeGetSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("ime_get_text")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnImeGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("virtual_keyboard_show")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnVirtualKeyboardShow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3042891259))
	}
	{
		methodName := StringNameFromStr("virtual_keyboard_hide")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnVirtualKeyboardHide = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("virtual_keyboard_get_height")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnVirtualKeyboardGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("cursor_set_shape")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnCursorSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2026291549))
	}
	{
		methodName := StringNameFromStr("cursor_get_shape")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnCursorGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1087724927))
	}
	{
		methodName := StringNameFromStr("cursor_set_custom_image")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnCursorSetCustomImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1816663697))
	}
	{
		methodName := StringNameFromStr("get_swap_cancel_ok")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnGetSwapCancelOk = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("enable_for_stealing_focus")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnEnableForStealingFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("dialog_show")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnDialogShow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4115553226))
	}
	{
		methodName := StringNameFromStr("dialog_input_text")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnDialogInputText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3088703427))
	}
	{
		methodName := StringNameFromStr("file_dialog_show")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnFileDialogShow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1531299078))
	}
	{
		methodName := StringNameFromStr("keyboard_get_layout_count")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnKeyboardGetLayoutCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("keyboard_get_current_layout")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnKeyboardGetCurrentLayout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("keyboard_set_current_layout")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnKeyboardSetCurrentLayout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("keyboard_get_layout_language")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnKeyboardGetLayoutLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("keyboard_get_layout_name")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnKeyboardGetLayoutName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("keyboard_get_keycode_from_physical")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnKeyboardGetKeycodeFromPhysical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3447613187))
	}
	{
		methodName := StringNameFromStr("keyboard_get_label_from_physical")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnKeyboardGetLabelFromPhysical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3447613187))
	}
	{
		methodName := StringNameFromStr("process_events")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnProcessEvents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("force_process_and_drop_events")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnForceProcessAndDropEvents = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_native_icon")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnSetNativeIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_icon")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnSetIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 532598488))
	}
	{
		methodName := StringNameFromStr("tablet_get_driver_count")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTabletGetDriverCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("tablet_get_driver_name")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTabletGetDriverName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("tablet_get_current_driver")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTabletGetCurrentDriver = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("tablet_set_current_driver")
		defer methodName.Destroy()
		ptrsForDisplayServer.fnTabletSetCurrentDriver = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
}

type DisplayServer struct {
	Object
}

func (me *DisplayServer) BaseClass() string {
	return "DisplayServer"
}

func NewDisplayServer() *DisplayServer {
	str := StringNameFromStr("DisplayServer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &DisplayServer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	DisplayServerScreenWithMouseFocus    = "-4" // TODO: construct correctly
	DisplayServerScreenWithKeyboardFocus = "-3" // TODO: construct correctly
	DisplayServerScreenPrimary           = "-2" // TODO: construct correctly
	DisplayServerScreenOfMainWindow      = "-1" // TODO: construct correctly
	DisplayServerMainWindowId            = "0"  // TODO: construct correctly
	DisplayServerInvalidWindowId         = "-1" // TODO: construct correctly
)

// Enums

type DisplayServerFeature int

const (
	DisplayServerFeatureFeatureGlobalMenu         DisplayServerFeature = 0
	DisplayServerFeatureFeatureSubwindows         DisplayServerFeature = 1
	DisplayServerFeatureFeatureTouchscreen        DisplayServerFeature = 2
	DisplayServerFeatureFeatureMouse              DisplayServerFeature = 3
	DisplayServerFeatureFeatureMouseWarp          DisplayServerFeature = 4
	DisplayServerFeatureFeatureClipboard          DisplayServerFeature = 5
	DisplayServerFeatureFeatureVirtualKeyboard    DisplayServerFeature = 6
	DisplayServerFeatureFeatureCursorShape        DisplayServerFeature = 7
	DisplayServerFeatureFeatureCustomCursorShape  DisplayServerFeature = 8
	DisplayServerFeatureFeatureNativeDialog       DisplayServerFeature = 9
	DisplayServerFeatureFeatureIme                DisplayServerFeature = 10
	DisplayServerFeatureFeatureWindowTransparency DisplayServerFeature = 11
	DisplayServerFeatureFeatureHidpi              DisplayServerFeature = 12
	DisplayServerFeatureFeatureIcon               DisplayServerFeature = 13
	DisplayServerFeatureFeatureNativeIcon         DisplayServerFeature = 14
	DisplayServerFeatureFeatureOrientation        DisplayServerFeature = 15
	DisplayServerFeatureFeatureSwapBuffers        DisplayServerFeature = 16
	DisplayServerFeatureFeatureClipboardPrimary   DisplayServerFeature = 18
	DisplayServerFeatureFeatureTextToSpeech       DisplayServerFeature = 19
	DisplayServerFeatureFeatureExtendToTitle      DisplayServerFeature = 20
	DisplayServerFeatureFeatureScreenCapture      DisplayServerFeature = 21
)

type DisplayServerMouseMode int

const (
	DisplayServerMouseModeMouseModeVisible        DisplayServerMouseMode = 0
	DisplayServerMouseModeMouseModeHidden         DisplayServerMouseMode = 1
	DisplayServerMouseModeMouseModeCaptured       DisplayServerMouseMode = 2
	DisplayServerMouseModeMouseModeConfined       DisplayServerMouseMode = 3
	DisplayServerMouseModeMouseModeConfinedHidden DisplayServerMouseMode = 4
)

type DisplayServerScreenOrientation int

const (
	DisplayServerScreenOrientationScreenLandscape        DisplayServerScreenOrientation = 0
	DisplayServerScreenOrientationScreenPortrait         DisplayServerScreenOrientation = 1
	DisplayServerScreenOrientationScreenReverseLandscape DisplayServerScreenOrientation = 2
	DisplayServerScreenOrientationScreenReversePortrait  DisplayServerScreenOrientation = 3
	DisplayServerScreenOrientationScreenSensorLandscape  DisplayServerScreenOrientation = 4
	DisplayServerScreenOrientationScreenSensorPortrait   DisplayServerScreenOrientation = 5
	DisplayServerScreenOrientationScreenSensor           DisplayServerScreenOrientation = 6
)

type DisplayServerVirtualKeyboardType int

const (
	DisplayServerVirtualKeyboardTypeKeyboardTypeDefault       DisplayServerVirtualKeyboardType = 0
	DisplayServerVirtualKeyboardTypeKeyboardTypeMultiline     DisplayServerVirtualKeyboardType = 1
	DisplayServerVirtualKeyboardTypeKeyboardTypeNumber        DisplayServerVirtualKeyboardType = 2
	DisplayServerVirtualKeyboardTypeKeyboardTypeNumberDecimal DisplayServerVirtualKeyboardType = 3
	DisplayServerVirtualKeyboardTypeKeyboardTypePhone         DisplayServerVirtualKeyboardType = 4
	DisplayServerVirtualKeyboardTypeKeyboardTypeEmailAddress  DisplayServerVirtualKeyboardType = 5
	DisplayServerVirtualKeyboardTypeKeyboardTypePassword      DisplayServerVirtualKeyboardType = 6
	DisplayServerVirtualKeyboardTypeKeyboardTypeUrl           DisplayServerVirtualKeyboardType = 7
)

type DisplayServerCursorShape int

const (
	DisplayServerCursorShapeCursorArrow        DisplayServerCursorShape = 0
	DisplayServerCursorShapeCursorIbeam        DisplayServerCursorShape = 1
	DisplayServerCursorShapeCursorPointingHand DisplayServerCursorShape = 2
	DisplayServerCursorShapeCursorCross        DisplayServerCursorShape = 3
	DisplayServerCursorShapeCursorWait         DisplayServerCursorShape = 4
	DisplayServerCursorShapeCursorBusy         DisplayServerCursorShape = 5
	DisplayServerCursorShapeCursorDrag         DisplayServerCursorShape = 6
	DisplayServerCursorShapeCursorCanDrop      DisplayServerCursorShape = 7
	DisplayServerCursorShapeCursorForbidden    DisplayServerCursorShape = 8
	DisplayServerCursorShapeCursorVsize        DisplayServerCursorShape = 9
	DisplayServerCursorShapeCursorHsize        DisplayServerCursorShape = 10
	DisplayServerCursorShapeCursorBdiagsize    DisplayServerCursorShape = 11
	DisplayServerCursorShapeCursorFdiagsize    DisplayServerCursorShape = 12
	DisplayServerCursorShapeCursorMove         DisplayServerCursorShape = 13
	DisplayServerCursorShapeCursorVsplit       DisplayServerCursorShape = 14
	DisplayServerCursorShapeCursorHsplit       DisplayServerCursorShape = 15
	DisplayServerCursorShapeCursorHelp         DisplayServerCursorShape = 16
	DisplayServerCursorShapeCursorMax          DisplayServerCursorShape = 17
)

type DisplayServerFileDialogMode int

const (
	DisplayServerFileDialogModeFileDialogModeOpenFile  DisplayServerFileDialogMode = 0
	DisplayServerFileDialogModeFileDialogModeOpenFiles DisplayServerFileDialogMode = 1
	DisplayServerFileDialogModeFileDialogModeOpenDir   DisplayServerFileDialogMode = 2
	DisplayServerFileDialogModeFileDialogModeOpenAny   DisplayServerFileDialogMode = 3
	DisplayServerFileDialogModeFileDialogModeSaveFile  DisplayServerFileDialogMode = 4
)

type DisplayServerWindowMode int

const (
	DisplayServerWindowModeWindowModeWindowed            DisplayServerWindowMode = 0
	DisplayServerWindowModeWindowModeMinimized           DisplayServerWindowMode = 1
	DisplayServerWindowModeWindowModeMaximized           DisplayServerWindowMode = 2
	DisplayServerWindowModeWindowModeFullscreen          DisplayServerWindowMode = 3
	DisplayServerWindowModeWindowModeExclusiveFullscreen DisplayServerWindowMode = 4
)

type DisplayServerWindowFlags int

const (
	DisplayServerWindowFlagsWindowFlagResizeDisabled   DisplayServerWindowFlags = 0
	DisplayServerWindowFlagsWindowFlagBorderless       DisplayServerWindowFlags = 1
	DisplayServerWindowFlagsWindowFlagAlwaysOnTop      DisplayServerWindowFlags = 2
	DisplayServerWindowFlagsWindowFlagTransparent      DisplayServerWindowFlags = 3
	DisplayServerWindowFlagsWindowFlagNoFocus          DisplayServerWindowFlags = 4
	DisplayServerWindowFlagsWindowFlagPopup            DisplayServerWindowFlags = 5
	DisplayServerWindowFlagsWindowFlagExtendToTitle    DisplayServerWindowFlags = 6
	DisplayServerWindowFlagsWindowFlagMousePassthrough DisplayServerWindowFlags = 7
	DisplayServerWindowFlagsWindowFlagMax              DisplayServerWindowFlags = 8
)

type DisplayServerWindowEvent int

const (
	DisplayServerWindowEventWindowEventMouseEnter     DisplayServerWindowEvent = 0
	DisplayServerWindowEventWindowEventMouseExit      DisplayServerWindowEvent = 1
	DisplayServerWindowEventWindowEventFocusIn        DisplayServerWindowEvent = 2
	DisplayServerWindowEventWindowEventFocusOut       DisplayServerWindowEvent = 3
	DisplayServerWindowEventWindowEventCloseRequest   DisplayServerWindowEvent = 4
	DisplayServerWindowEventWindowEventGoBackRequest  DisplayServerWindowEvent = 5
	DisplayServerWindowEventWindowEventDpiChange      DisplayServerWindowEvent = 6
	DisplayServerWindowEventWindowEventTitlebarChange DisplayServerWindowEvent = 7
)

type DisplayServerVSyncMode int

const (
	DisplayServerVSyncModeVsyncDisabled DisplayServerVSyncMode = 0
	DisplayServerVSyncModeVsyncEnabled  DisplayServerVSyncMode = 1
	DisplayServerVSyncModeVsyncAdaptive DisplayServerVSyncMode = 2
	DisplayServerVSyncModeVsyncMailbox  DisplayServerVSyncMode = 3
)

type DisplayServerHandleType int

const (
	DisplayServerHandleTypeDisplayHandle DisplayServerHandleType = 0
	DisplayServerHandleTypeWindowHandle  DisplayServerHandleType = 1
	DisplayServerHandleTypeWindowView    DisplayServerHandleType = 2
	DisplayServerHandleTypeOpenglContext DisplayServerHandleType = 3
)

type DisplayServerTTSUtteranceEvent int

const (
	DisplayServerTTSUtteranceEventTtsUtteranceStarted  DisplayServerTTSUtteranceEvent = 0
	DisplayServerTTSUtteranceEventTtsUtteranceEnded    DisplayServerTTSUtteranceEvent = 1
	DisplayServerTTSUtteranceEventTtsUtteranceCanceled DisplayServerTTSUtteranceEvent = 2
	DisplayServerTTSUtteranceEventTtsUtteranceBoundary DisplayServerTTSUtteranceEvent = 3
)

func (me *DisplayServer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *DisplayServer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *DisplayServer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *DisplayServer) HasFeature(feature DisplayServerFeature) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&feature)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnHasFeature), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GetName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GlobalMenuSetPopupCallbacks(menu_root String, open_callback Callable, close_callback Callable) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), open_callback.AsCTypePtr(), close_callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetPopupCallbacks), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuAddSubmenuItem(menu_root String, label String, submenu String, index int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), label.AsCTypePtr(), submenu.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuAddSubmenuItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuAddItem(menu_root String, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuAddItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuAddCheckItem(menu_root String, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuAddCheckItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuAddIconItem(menu_root String, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), icon.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuAddIconItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuAddIconCheckItem(menu_root String, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), icon.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuAddIconCheckItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuAddRadioCheckItem(menu_root String, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuAddRadioCheckItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuAddIconRadioCheckItem(menu_root String, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), icon.AsCTypePtr(), label.AsCTypePtr(), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuAddIconRadioCheckItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuAddMultistateItem(menu_root String, label String, max_states int64, default_state int64, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), label.AsCTypePtr(), gdc.ConstTypePtr(&max_states), gdc.ConstTypePtr(&default_state), callback.AsCTypePtr(), key_callback.AsCTypePtr(), tag.AsCTypePtr(), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&max_states)
	pinner.Pin(&default_state)
	pinner.Pin(&accelerator)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuAddMultistateItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuAddSeparator(menu_root String, index int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuAddSeparator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuGetItemIndexFromText(menu_root String, text String) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemIndexFromText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuGetItemIndexFromTag(menu_root String, tag Variant) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), tag.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemIndexFromTag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuIsItemChecked(menu_root String, idx int64) bool {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuIsItemChecked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuIsItemCheckable(menu_root String, idx int64) bool {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuIsItemCheckable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuIsItemRadioCheckable(menu_root String, idx int64) bool {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuIsItemRadioCheckable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuGetItemCallback(menu_root String, idx int64) Callable {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCallable()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemCallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GlobalMenuGetItemKeyCallback(menu_root String, idx int64) Callable {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCallable()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemKeyCallback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GlobalMenuGetItemTag(menu_root String, idx int64) Variant {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemTag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GlobalMenuGetItemText(menu_root String, idx int64) String {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GlobalMenuGetItemSubmenu(menu_root String, idx int64) String {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemSubmenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GlobalMenuGetItemAccelerator(menu_root String, idx int64) Key {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemAccelerator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) GlobalMenuIsItemDisabled(menu_root String, idx int64) bool {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuIsItemDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuIsItemHidden(menu_root String, idx int64) bool {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuIsItemHidden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuGetItemTooltip(menu_root String, idx int64) String {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemTooltip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GlobalMenuGetItemState(menu_root String, idx int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuGetItemMaxStates(menu_root String, idx int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemMaxStates), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuGetItemIcon(menu_root String, idx int64) Texture2D {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GlobalMenuGetItemIndentationLevel(menu_root String, idx int64) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemIndentationLevel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuSetItemChecked(menu_root String, idx int64, checked bool) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&checked)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemChecked), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemCheckable(menu_root String, idx int64, checkable bool) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&checkable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemCheckable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemRadioCheckable(menu_root String, idx int64, checkable bool) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&checkable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemRadioCheckable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemCallback(menu_root String, idx int64, callback Callable) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemHoverCallbacks(menu_root String, idx int64, callback Callable) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemHoverCallbacks), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemKeyCallback(menu_root String, idx int64, key_callback Callable) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), key_callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemKeyCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemTag(menu_root String, idx int64, tag Variant) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), tag.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemTag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemText(menu_root String, idx int64, text String) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemSubmenu(menu_root String, idx int64, submenu String) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), submenu.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemSubmenu), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemAccelerator(menu_root String, idx int64, keycode Key) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&keycode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemAccelerator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemDisabled(menu_root String, idx int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemHidden(menu_root String, idx int64, hidden bool) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemHidden), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemTooltip(menu_root String, idx int64, tooltip String) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), tooltip.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemTooltip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemState(menu_root String, idx int64, state int64) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&state)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemState), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemMaxStates(menu_root String, idx int64, max_states int64) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&max_states)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemMaxStates), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemIcon(menu_root String, idx int64, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuSetItemIndentationLevel(menu_root String, idx int64, level int64) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&level)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuSetItemIndentationLevel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuGetItemCount(menu_root String) int64 {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuGetItemCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GlobalMenuRemoveItem(menu_root String, idx int64) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr(), gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuRemoveItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GlobalMenuClear(menu_root String) {
	cargs := []gdc.ConstTypePtr{menu_root.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGlobalMenuClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) TtsIsSpeaking() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTtsIsSpeaking), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) TtsIsPaused() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTtsIsPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) TtsGetVoices() []Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTtsGetVoices), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *DisplayServer) TtsGetVoicesForLanguage(language String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTtsGetVoicesForLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) TtsSpeak(text String, voice String, volume int64, pitch float64, rate float64, utterance_id int64, interrupt bool) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), voice.AsCTypePtr(), gdc.ConstTypePtr(&volume), gdc.ConstTypePtr(&pitch), gdc.ConstTypePtr(&rate), gdc.ConstTypePtr(&utterance_id), gdc.ConstTypePtr(&interrupt)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTtsSpeak), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) TtsPause() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTtsPause), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) TtsResume() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTtsResume), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) TtsStop() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTtsStop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) TtsSetUtteranceCallback(event DisplayServerTTSUtteranceEvent, callable Callable) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&event), callable.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTtsSetUtteranceCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) IsDarkModeSupported() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnIsDarkModeSupported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) IsDarkMode() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnIsDarkMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GetAccentColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetAccentColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) MouseSetMode(mouse_mode DisplayServerMouseMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mouse_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnMouseSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) MouseGetMode() DisplayServerMouseMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret DisplayServerMouseMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnMouseGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) WarpMouse(position Vector2i) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWarpMouse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) MouseGetPosition() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnMouseGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) MouseGetButtonState() MouseButtonMask {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret MouseButtonMask

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnMouseGetButtonState), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) ClipboardSet(clipboard String) {
	cargs := []gdc.ConstTypePtr{clipboard.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnClipboardSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) ClipboardGet() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnClipboardGet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) ClipboardGetImage() Image {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImage()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnClipboardGetImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) ClipboardHas() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnClipboardHas), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) ClipboardHasImage() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnClipboardHasImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) ClipboardSetPrimary(clipboard_primary String) {
	cargs := []gdc.ConstTypePtr{clipboard_primary.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnClipboardSetPrimary), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) ClipboardGetPrimary() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnClipboardGetPrimary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GetDisplayCutouts() []Rect2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetDisplayCutouts), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Rect2](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *DisplayServer) GetDisplaySafeArea() Rect2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetDisplaySafeArea), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GetScreenCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetScreenCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GetPrimaryScreen() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetPrimaryScreen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GetKeyboardFocusScreen() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetKeyboardFocusScreen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GetScreenFromRect(rect Rect2) int64 {
	cargs := []gdc.ConstTypePtr{rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetScreenFromRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) ScreenGetPosition(screen int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) ScreenGetSize(screen int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) ScreenGetUsableRect(screen int64) Rect2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2i()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetUsableRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) ScreenGetDpi(screen int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetDpi), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) ScreenGetScale(screen int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) IsTouchscreenAvailable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnIsTouchscreenAvailable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) ScreenGetMaxScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetMaxScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) ScreenGetRefreshRate(screen int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetRefreshRate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) ScreenGetPixel(position Vector2i) Color {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetPixel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) ScreenGetImage(screen int64) Image {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImage()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) ScreenSetOrientation(orientation DisplayServerScreenOrientation, screen int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&orientation), gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenSetOrientation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) ScreenGetOrientation(screen int64) DisplayServerScreenOrientation {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret DisplayServerScreenOrientation
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenGetOrientation), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) ScreenSetKeepOn(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenSetKeepOn), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) ScreenIsKeptOn() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnScreenIsKeptOn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) GetWindowList() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetWindowList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) GetWindowAtScreenPosition(position Vector2i) int64 {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetWindowAtScreenPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowGetNativeHandle(handle_type DisplayServerHandleType, window_id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&handle_type), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&handle_type)
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetNativeHandle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowGetActivePopup() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetActivePopup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowSetPopupSafeRect(window int64, rect Rect2i) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window), rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetPopupSafeRect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetPopupSafeRect(window int64) Rect2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2i()
	pinner.Pin(&window)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetPopupSafeRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) WindowSetTitle(title String, window_id int64) {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetTitleSize(title String, window_id int64) Vector2i {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetTitleSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) WindowSetMousePassthrough(region PackedVector2Array, window_id int64) {
	cargs := []gdc.ConstTypePtr{region.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetMousePassthrough), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetCurrentScreen(window_id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetCurrentScreen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowSetCurrentScreen(screen int64, window_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetCurrentScreen), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetPosition(window_id int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) WindowGetPositionWithDecorations(window_id int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetPositionWithDecorations), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) WindowSetPosition(position Vector2i, window_id int64) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetSize(window_id int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) WindowSetSize(size Vector2i, window_id int64) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetRectChangedCallback(callback Callable, window_id int64) {
	cargs := []gdc.ConstTypePtr{callback.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetRectChangedCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetWindowEventCallback(callback Callable, window_id int64) {
	cargs := []gdc.ConstTypePtr{callback.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetWindowEventCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetInputEventCallback(callback Callable, window_id int64) {
	cargs := []gdc.ConstTypePtr{callback.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetInputEventCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetInputTextCallback(callback Callable, window_id int64) {
	cargs := []gdc.ConstTypePtr{callback.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetInputTextCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetDropFilesCallback(callback Callable, window_id int64) {
	cargs := []gdc.ConstTypePtr{callback.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetDropFilesCallback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetAttachedInstanceId(window_id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetAttachedInstanceId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowGetMaxSize(window_id int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetMaxSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) WindowSetMaxSize(max_size Vector2i, window_id int64) {
	cargs := []gdc.ConstTypePtr{max_size.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetMaxSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetMinSize(window_id int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetMinSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) WindowSetMinSize(min_size Vector2i, window_id int64) {
	cargs := []gdc.ConstTypePtr{min_size.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetMinSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetSizeWithDecorations(window_id int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetSizeWithDecorations), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) WindowGetMode(window_id int64) DisplayServerWindowMode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret DisplayServerWindowMode
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) WindowSetMode(mode DisplayServerWindowMode, window_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetFlag(flag DisplayServerWindowFlags, enabled bool, window_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enabled), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetFlag), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetFlag(flag DisplayServerWindowFlags, window_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&flag)
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetFlag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowSetWindowButtonsOffset(offset Vector2i, window_id int64) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetWindowButtonsOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetSafeTitleMargins(window_id int64) Vector3i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3i()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetSafeTitleMargins), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) WindowRequestAttention(window_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowRequestAttention), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowMoveToForeground(window_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowMoveToForeground), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowIsFocused(window_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowIsFocused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowCanDraw(window_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowCanDraw), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowSetTransient(window_id int64, parent_window_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), gdc.ConstTypePtr(&parent_window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetTransient), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetExclusive(window_id int64, exclusive bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), gdc.ConstTypePtr(&exclusive)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetExclusive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetImeActive(active bool, window_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetImeActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetImePosition(position Vector2i, window_id int64) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetImePosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowSetVsyncMode(vsync_mode DisplayServerVSyncMode, window_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vsync_mode), gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowSetVsyncMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) WindowGetVsyncMode(window_id int64) DisplayServerVSyncMode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret DisplayServerVSyncMode
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowGetVsyncMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) WindowIsMaximizeAllowed(window_id int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&window_id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowIsMaximizeAllowed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowMaximizeOnTitleDblClick() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowMaximizeOnTitleDblClick), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) WindowMinimizeOnTitleDblClick() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnWindowMinimizeOnTitleDblClick), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) ImeGetSelection() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnImeGetSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) ImeGetText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnImeGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) VirtualKeyboardShow(existing_text String, position Rect2, type_ DisplayServerVirtualKeyboardType, max_length int64, cursor_start int64, cursor_end int64) {
	cargs := []gdc.ConstTypePtr{existing_text.AsCTypePtr(), position.AsCTypePtr(), gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&max_length), gdc.ConstTypePtr(&cursor_start), gdc.ConstTypePtr(&cursor_end)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnVirtualKeyboardShow), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) VirtualKeyboardHide() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnVirtualKeyboardHide), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) VirtualKeyboardGetHeight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnVirtualKeyboardGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) CursorSetShape(shape DisplayServerCursorShape) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnCursorSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) CursorGetShape() DisplayServerCursorShape {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret DisplayServerCursorShape

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnCursorGetShape), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) CursorSetCustomImage(cursor Resource, shape DisplayServerCursorShape, hotspot Vector2) {
	cargs := []gdc.ConstTypePtr{cursor.AsCTypePtr(), gdc.ConstTypePtr(&shape), hotspot.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnCursorSetCustomImage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) GetSwapCancelOk() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnGetSwapCancelOk), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) EnableForStealingFocus(process_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&process_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnEnableForStealingFocus), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) DialogShow(title String, description String, buttons PackedStringArray, callback Callable) Error {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr(), description.AsCTypePtr(), buttons.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnDialogShow), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) DialogInputText(title String, description String, existing_text String, callback Callable) Error {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr(), description.AsCTypePtr(), existing_text.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnDialogInputText), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) FileDialogShow(title String, current_directory String, filename String, show_hidden bool, mode DisplayServerFileDialogMode, filters PackedStringArray, callback Callable) Error {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr(), current_directory.AsCTypePtr(), filename.AsCTypePtr(), gdc.ConstTypePtr(&show_hidden), gdc.ConstTypePtr(&mode), filters.AsCTypePtr(), callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&show_hidden)
	pinner.Pin(&mode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnFileDialogShow), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) KeyboardGetLayoutCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnKeyboardGetLayoutCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) KeyboardGetCurrentLayout() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnKeyboardGetCurrentLayout), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) KeyboardSetCurrentLayout(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnKeyboardSetCurrentLayout), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) KeyboardGetLayoutLanguage(index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnKeyboardGetLayoutLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) KeyboardGetLayoutName(index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnKeyboardGetLayoutName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) KeyboardGetKeycodeFromPhysical(keycode Key) Key {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key
	pinner.Pin(&keycode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnKeyboardGetKeycodeFromPhysical), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) KeyboardGetLabelFromPhysical(keycode Key) Key {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key
	pinner.Pin(&keycode)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnKeyboardGetLabelFromPhysical), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DisplayServer) ProcessEvents() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnProcessEvents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) ForceProcessAndDropEvents() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnForceProcessAndDropEvents), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) SetNativeIcon(filename String) {
	cargs := []gdc.ConstTypePtr{filename.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnSetNativeIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) SetIcon(image Image) {
	cargs := []gdc.ConstTypePtr{image.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnSetIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DisplayServer) TabletGetDriverCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTabletGetDriverCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DisplayServer) TabletGetDriverName(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTabletGetDriverName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) TabletGetCurrentDriver() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTabletGetCurrentDriver), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DisplayServer) TabletSetCurrentDriver(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDisplayServer.fnTabletSetCurrentDriver), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
