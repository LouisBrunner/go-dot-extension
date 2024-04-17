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

type ptrsForTextEditList struct {
	fnXHandleUnicodeInput                  gdc.MethodBindPtr
	fnXBackspace                           gdc.MethodBindPtr
	fnXCut                                 gdc.MethodBindPtr
	fnXCopy                                gdc.MethodBindPtr
	fnXPaste                               gdc.MethodBindPtr
	fnXPastePrimaryClipboard               gdc.MethodBindPtr
	fnHasImeText                           gdc.MethodBindPtr
	fnCancelIme                            gdc.MethodBindPtr
	fnApplyIme                             gdc.MethodBindPtr
	fnSetEditable                          gdc.MethodBindPtr
	fnIsEditable                           gdc.MethodBindPtr
	fnSetTextDirection                     gdc.MethodBindPtr
	fnGetTextDirection                     gdc.MethodBindPtr
	fnSetLanguage                          gdc.MethodBindPtr
	fnGetLanguage                          gdc.MethodBindPtr
	fnSetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnGetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnSetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnGetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnSetTabSize                           gdc.MethodBindPtr
	fnGetTabSize                           gdc.MethodBindPtr
	fnSetIndentWrappedLines                gdc.MethodBindPtr
	fnIsIndentWrappedLines                 gdc.MethodBindPtr
	fnSetOvertypeModeEnabled               gdc.MethodBindPtr
	fnIsOvertypeModeEnabled                gdc.MethodBindPtr
	fnSetContextMenuEnabled                gdc.MethodBindPtr
	fnIsContextMenuEnabled                 gdc.MethodBindPtr
	fnSetShortcutKeysEnabled               gdc.MethodBindPtr
	fnIsShortcutKeysEnabled                gdc.MethodBindPtr
	fnSetVirtualKeyboardEnabled            gdc.MethodBindPtr
	fnIsVirtualKeyboardEnabled             gdc.MethodBindPtr
	fnSetMiddleMousePasteEnabled           gdc.MethodBindPtr
	fnIsMiddleMousePasteEnabled            gdc.MethodBindPtr
	fnClear                                gdc.MethodBindPtr
	fnSetText                              gdc.MethodBindPtr
	fnGetText                              gdc.MethodBindPtr
	fnGetLineCount                         gdc.MethodBindPtr
	fnSetPlaceholder                       gdc.MethodBindPtr
	fnGetPlaceholder                       gdc.MethodBindPtr
	fnSetLine                              gdc.MethodBindPtr
	fnGetLine                              gdc.MethodBindPtr
	fnGetLineWidth                         gdc.MethodBindPtr
	fnGetLineHeight                        gdc.MethodBindPtr
	fnGetIndentLevel                       gdc.MethodBindPtr
	fnGetFirstNonWhitespaceColumn          gdc.MethodBindPtr
	fnSwapLines                            gdc.MethodBindPtr
	fnInsertLineAt                         gdc.MethodBindPtr
	fnInsertTextAtCaret                    gdc.MethodBindPtr
	fnRemoveText                           gdc.MethodBindPtr
	fnGetLastUnhiddenLine                  gdc.MethodBindPtr
	fnGetNextVisibleLineOffsetFrom         gdc.MethodBindPtr
	fnGetNextVisibleLineIndexOffsetFrom    gdc.MethodBindPtr
	fnBackspace                            gdc.MethodBindPtr
	fnCut                                  gdc.MethodBindPtr
	fnCopy                                 gdc.MethodBindPtr
	fnPaste                                gdc.MethodBindPtr
	fnPastePrimaryClipboard                gdc.MethodBindPtr
	fnStartAction                          gdc.MethodBindPtr
	fnEndAction                            gdc.MethodBindPtr
	fnBeginComplexOperation                gdc.MethodBindPtr
	fnEndComplexOperation                  gdc.MethodBindPtr
	fnHasUndo                              gdc.MethodBindPtr
	fnHasRedo                              gdc.MethodBindPtr
	fnUndo                                 gdc.MethodBindPtr
	fnRedo                                 gdc.MethodBindPtr
	fnClearUndoHistory                     gdc.MethodBindPtr
	fnTagSavedVersion                      gdc.MethodBindPtr
	fnGetVersion                           gdc.MethodBindPtr
	fnGetSavedVersion                      gdc.MethodBindPtr
	fnSetSearchText                        gdc.MethodBindPtr
	fnSetSearchFlags                       gdc.MethodBindPtr
	fnSearch                               gdc.MethodBindPtr
	fnSetTooltipRequestFunc                gdc.MethodBindPtr
	fnGetLocalMousePos                     gdc.MethodBindPtr
	fnGetWordAtPos                         gdc.MethodBindPtr
	fnGetLineColumnAtPos                   gdc.MethodBindPtr
	fnGetPosAtLineColumn                   gdc.MethodBindPtr
	fnGetRectAtLineColumn                  gdc.MethodBindPtr
	fnGetMinimapLineAtPos                  gdc.MethodBindPtr
	fnIsDraggingCursor                     gdc.MethodBindPtr
	fnIsMouseOverSelection                 gdc.MethodBindPtr
	fnSetCaretType                         gdc.MethodBindPtr
	fnGetCaretType                         gdc.MethodBindPtr
	fnSetCaretBlinkEnabled                 gdc.MethodBindPtr
	fnIsCaretBlinkEnabled                  gdc.MethodBindPtr
	fnSetCaretBlinkInterval                gdc.MethodBindPtr
	fnGetCaretBlinkInterval                gdc.MethodBindPtr
	fnSetDrawCaretWhenEditableDisabled     gdc.MethodBindPtr
	fnIsDrawingCaretWhenEditableDisabled   gdc.MethodBindPtr
	fnSetMoveCaretOnRightClickEnabled      gdc.MethodBindPtr
	fnIsMoveCaretOnRightClickEnabled       gdc.MethodBindPtr
	fnSetCaretMidGraphemeEnabled           gdc.MethodBindPtr
	fnIsCaretMidGraphemeEnabled            gdc.MethodBindPtr
	fnSetMultipleCaretsEnabled             gdc.MethodBindPtr
	fnIsMultipleCaretsEnabled              gdc.MethodBindPtr
	fnAddCaret                             gdc.MethodBindPtr
	fnRemoveCaret                          gdc.MethodBindPtr
	fnRemoveSecondaryCarets                gdc.MethodBindPtr
	fnMergeOverlappingCarets               gdc.MethodBindPtr
	fnGetCaretCount                        gdc.MethodBindPtr
	fnAddCaretAtCarets                     gdc.MethodBindPtr
	fnGetCaretIndexEditOrder               gdc.MethodBindPtr
	fnAdjustCaretsAfterEdit                gdc.MethodBindPtr
	fnIsCaretVisible                       gdc.MethodBindPtr
	fnGetCaretDrawPos                      gdc.MethodBindPtr
	fnSetCaretLine                         gdc.MethodBindPtr
	fnGetCaretLine                         gdc.MethodBindPtr
	fnSetCaretColumn                       gdc.MethodBindPtr
	fnGetCaretColumn                       gdc.MethodBindPtr
	fnGetCaretWrapIndex                    gdc.MethodBindPtr
	fnGetWordUnderCaret                    gdc.MethodBindPtr
	fnSetSelectingEnabled                  gdc.MethodBindPtr
	fnIsSelectingEnabled                   gdc.MethodBindPtr
	fnSetDeselectOnFocusLossEnabled        gdc.MethodBindPtr
	fnIsDeselectOnFocusLossEnabled         gdc.MethodBindPtr
	fnSetDragAndDropSelectionEnabled       gdc.MethodBindPtr
	fnIsDragAndDropSelectionEnabled        gdc.MethodBindPtr
	fnSetSelectionMode                     gdc.MethodBindPtr
	fnGetSelectionMode                     gdc.MethodBindPtr
	fnSelectAll                            gdc.MethodBindPtr
	fnSelectWordUnderCaret                 gdc.MethodBindPtr
	fnAddSelectionForNextOccurrence        gdc.MethodBindPtr
	fnSkipSelectionForNextOccurrence       gdc.MethodBindPtr
	fnSelect                               gdc.MethodBindPtr
	fnHasSelection                         gdc.MethodBindPtr
	fnGetSelectedText                      gdc.MethodBindPtr
	fnGetSelectionLine                     gdc.MethodBindPtr
	fnGetSelectionColumn                   gdc.MethodBindPtr
	fnGetSelectionFromLine                 gdc.MethodBindPtr
	fnGetSelectionFromColumn               gdc.MethodBindPtr
	fnGetSelectionToLine                   gdc.MethodBindPtr
	fnGetSelectionToColumn                 gdc.MethodBindPtr
	fnDeselect                             gdc.MethodBindPtr
	fnDeleteSelection                      gdc.MethodBindPtr
	fnSetLineWrappingMode                  gdc.MethodBindPtr
	fnGetLineWrappingMode                  gdc.MethodBindPtr
	fnSetAutowrapMode                      gdc.MethodBindPtr
	fnGetAutowrapMode                      gdc.MethodBindPtr
	fnIsLineWrapped                        gdc.MethodBindPtr
	fnGetLineWrapCount                     gdc.MethodBindPtr
	fnGetLineWrapIndexAtColumn             gdc.MethodBindPtr
	fnGetLineWrappedText                   gdc.MethodBindPtr
	fnSetSmoothScrollEnabled               gdc.MethodBindPtr
	fnIsSmoothScrollEnabled                gdc.MethodBindPtr
	fnGetVScrollBar                        gdc.MethodBindPtr
	fnGetHScrollBar                        gdc.MethodBindPtr
	fnSetVScroll                           gdc.MethodBindPtr
	fnGetVScroll                           gdc.MethodBindPtr
	fnSetHScroll                           gdc.MethodBindPtr
	fnGetHScroll                           gdc.MethodBindPtr
	fnSetScrollPastEndOfFileEnabled        gdc.MethodBindPtr
	fnIsScrollPastEndOfFileEnabled         gdc.MethodBindPtr
	fnSetVScrollSpeed                      gdc.MethodBindPtr
	fnGetVScrollSpeed                      gdc.MethodBindPtr
	fnSetFitContentHeightEnabled           gdc.MethodBindPtr
	fnIsFitContentHeightEnabled            gdc.MethodBindPtr
	fnGetScrollPosForLine                  gdc.MethodBindPtr
	fnSetLineAsFirstVisible                gdc.MethodBindPtr
	fnGetFirstVisibleLine                  gdc.MethodBindPtr
	fnSetLineAsCenterVisible               gdc.MethodBindPtr
	fnSetLineAsLastVisible                 gdc.MethodBindPtr
	fnGetLastFullVisibleLine               gdc.MethodBindPtr
	fnGetLastFullVisibleLineWrapIndex      gdc.MethodBindPtr
	fnGetVisibleLineCount                  gdc.MethodBindPtr
	fnGetVisibleLineCountInRange           gdc.MethodBindPtr
	fnGetTotalVisibleLineCount             gdc.MethodBindPtr
	fnAdjustViewportToCaret                gdc.MethodBindPtr
	fnCenterViewportToCaret                gdc.MethodBindPtr
	fnSetDrawMinimap                       gdc.MethodBindPtr
	fnIsDrawingMinimap                     gdc.MethodBindPtr
	fnSetMinimapWidth                      gdc.MethodBindPtr
	fnGetMinimapWidth                      gdc.MethodBindPtr
	fnGetMinimapVisibleLines               gdc.MethodBindPtr
	fnAddGutter                            gdc.MethodBindPtr
	fnRemoveGutter                         gdc.MethodBindPtr
	fnGetGutterCount                       gdc.MethodBindPtr
	fnSetGutterName                        gdc.MethodBindPtr
	fnGetGutterName                        gdc.MethodBindPtr
	fnSetGutterType                        gdc.MethodBindPtr
	fnGetGutterType                        gdc.MethodBindPtr
	fnSetGutterWidth                       gdc.MethodBindPtr
	fnGetGutterWidth                       gdc.MethodBindPtr
	fnSetGutterDraw                        gdc.MethodBindPtr
	fnIsGutterDrawn                        gdc.MethodBindPtr
	fnSetGutterClickable                   gdc.MethodBindPtr
	fnIsGutterClickable                    gdc.MethodBindPtr
	fnSetGutterOverwritable                gdc.MethodBindPtr
	fnIsGutterOverwritable                 gdc.MethodBindPtr
	fnMergeGutters                         gdc.MethodBindPtr
	fnSetGutterCustomDraw                  gdc.MethodBindPtr
	fnGetTotalGutterWidth                  gdc.MethodBindPtr
	fnSetLineGutterMetadata                gdc.MethodBindPtr
	fnGetLineGutterMetadata                gdc.MethodBindPtr
	fnSetLineGutterText                    gdc.MethodBindPtr
	fnGetLineGutterText                    gdc.MethodBindPtr
	fnSetLineGutterIcon                    gdc.MethodBindPtr
	fnGetLineGutterIcon                    gdc.MethodBindPtr
	fnSetLineGutterItemColor               gdc.MethodBindPtr
	fnGetLineGutterItemColor               gdc.MethodBindPtr
	fnSetLineGutterClickable               gdc.MethodBindPtr
	fnIsLineGutterClickable                gdc.MethodBindPtr
	fnSetLineBackgroundColor               gdc.MethodBindPtr
	fnGetLineBackgroundColor               gdc.MethodBindPtr
	fnSetSyntaxHighlighter                 gdc.MethodBindPtr
	fnGetSyntaxHighlighter                 gdc.MethodBindPtr
	fnSetHighlightCurrentLine              gdc.MethodBindPtr
	fnIsHighlightCurrentLineEnabled        gdc.MethodBindPtr
	fnSetHighlightAllOccurrences           gdc.MethodBindPtr
	fnIsHighlightAllOccurrencesEnabled     gdc.MethodBindPtr
	fnGetDrawControlChars                  gdc.MethodBindPtr
	fnSetDrawControlChars                  gdc.MethodBindPtr
	fnSetDrawTabs                          gdc.MethodBindPtr
	fnIsDrawingTabs                        gdc.MethodBindPtr
	fnSetDrawSpaces                        gdc.MethodBindPtr
	fnIsDrawingSpaces                      gdc.MethodBindPtr
	fnGetMenu                              gdc.MethodBindPtr
	fnIsMenuVisible                        gdc.MethodBindPtr
	fnMenuOption                           gdc.MethodBindPtr
}

var ptrsForTextEdit ptrsForTextEditList

func initTextEditPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TextEdit")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("has_ime_text")
		defer methodName.Destroy()
		ptrsForTextEdit.fnHasImeText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("cancel_ime")
		defer methodName.Destroy()
		ptrsForTextEdit.fnCancelIme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("apply_ime")
		defer methodName.Destroy()
		ptrsForTextEdit.fnApplyIme = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_editable")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_editable")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_text_direction")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 119160795))
	}
	{
		methodName := StringNameFromStr("get_text_direction")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797257663))
	}
	{
		methodName := StringNameFromStr("set_language")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_language")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 55961453))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3385126229))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_tab_size")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetTabSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_tab_size")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetTabSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_indent_wrapped_lines")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetIndentWrappedLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_indent_wrapped_lines")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsIndentWrappedLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_overtype_mode_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetOvertypeModeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_overtype_mode_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsOvertypeModeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_context_menu_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetContextMenuEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_context_menu_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsContextMenuEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_shortcut_keys_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetShortcutKeysEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_shortcut_keys_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsShortcutKeysEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_virtual_keyboard_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetVirtualKeyboardEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_virtual_keyboard_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsVirtualKeyboardEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_middle_mouse_paste_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetMiddleMousePasteEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_middle_mouse_paste_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsMiddleMousePasteEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForTextEdit.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_text")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_text")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_line_count")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_placeholder")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetPlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_placeholder")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetPlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_line_width")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 688195400))
	}
	{
		methodName := StringNameFromStr("get_line_height")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_indent_level")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetIndentLevel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_first_non_whitespace_column")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetFirstNonWhitespaceColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("swap_lines")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSwapLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("insert_line_at")
		defer methodName.Destroy()
		ptrsForTextEdit.fnInsertLineAt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("insert_text_at_caret")
		defer methodName.Destroy()
		ptrsForTextEdit.fnInsertTextAtCaret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2697778442))
	}
	{
		methodName := StringNameFromStr("remove_text")
		defer methodName.Destroy()
		ptrsForTextEdit.fnRemoveText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4275841770))
	}
	{
		methodName := StringNameFromStr("get_last_unhidden_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLastUnhiddenLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_next_visible_line_offset_from")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetNextVisibleLineOffsetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("get_next_visible_line_index_offset_from")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetNextVisibleLineIndexOffsetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3386475622))
	}
	{
		methodName := StringNameFromStr("backspace")
		defer methodName.Destroy()
		ptrsForTextEdit.fnBackspace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("cut")
		defer methodName.Destroy()
		ptrsForTextEdit.fnCut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("copy")
		defer methodName.Destroy()
		ptrsForTextEdit.fnCopy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("paste")
		defer methodName.Destroy()
		ptrsForTextEdit.fnPaste = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("paste_primary_clipboard")
		defer methodName.Destroy()
		ptrsForTextEdit.fnPastePrimaryClipboard = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("start_action")
		defer methodName.Destroy()
		ptrsForTextEdit.fnStartAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2834827583))
	}
	{
		methodName := StringNameFromStr("end_action")
		defer methodName.Destroy()
		ptrsForTextEdit.fnEndAction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("begin_complex_operation")
		defer methodName.Destroy()
		ptrsForTextEdit.fnBeginComplexOperation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("end_complex_operation")
		defer methodName.Destroy()
		ptrsForTextEdit.fnEndComplexOperation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("has_undo")
		defer methodName.Destroy()
		ptrsForTextEdit.fnHasUndo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("has_redo")
		defer methodName.Destroy()
		ptrsForTextEdit.fnHasRedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("undo")
		defer methodName.Destroy()
		ptrsForTextEdit.fnUndo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("redo")
		defer methodName.Destroy()
		ptrsForTextEdit.fnRedo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear_undo_history")
		defer methodName.Destroy()
		ptrsForTextEdit.fnClearUndoHistory = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("tag_saved_version")
		defer methodName.Destroy()
		ptrsForTextEdit.fnTagSavedVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_version")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_saved_version")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSavedVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_search_text")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetSearchText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_search_flags")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetSearchFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("search")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSearch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1203739136))
	}
	{
		methodName := StringNameFromStr("set_tooltip_request_func")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetTooltipRequestFunc = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1611583062))
	}
	{
		methodName := StringNameFromStr("get_local_mouse_pos")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLocalMousePos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_word_at_pos")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetWordAtPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3674420000))
	}
	{
		methodName := StringNameFromStr("get_line_column_at_pos")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineColumnAtPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 239517838))
	}
	{
		methodName := StringNameFromStr("get_pos_at_line_column")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetPosAtLineColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 410388347))
	}
	{
		methodName := StringNameFromStr("get_rect_at_line_column")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetRectAtLineColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3256618057))
	}
	{
		methodName := StringNameFromStr("get_minimap_line_at_pos")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetMinimapLineAtPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2485466453))
	}
	{
		methodName := StringNameFromStr("is_dragging_cursor")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsDraggingCursor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_mouse_over_selection")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsMouseOverSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1840282309))
	}
	{
		methodName := StringNameFromStr("set_caret_type")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetCaretType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1211596914))
	}
	{
		methodName := StringNameFromStr("get_caret_type")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetCaretType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2830252959))
	}
	{
		methodName := StringNameFromStr("set_caret_blink_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetCaretBlinkEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_caret_blink_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsCaretBlinkEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_caret_blink_interval")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetCaretBlinkInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_caret_blink_interval")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetCaretBlinkInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_draw_caret_when_editable_disabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetDrawCaretWhenEditableDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_drawing_caret_when_editable_disabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsDrawingCaretWhenEditableDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_move_caret_on_right_click_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetMoveCaretOnRightClickEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_move_caret_on_right_click_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsMoveCaretOnRightClickEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_caret_mid_grapheme_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetCaretMidGraphemeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_caret_mid_grapheme_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsCaretMidGraphemeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_multiple_carets_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetMultipleCaretsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_multiple_carets_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsMultipleCaretsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("add_caret")
		defer methodName.Destroy()
		ptrsForTextEdit.fnAddCaret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 50157827))
	}
	{
		methodName := StringNameFromStr("remove_caret")
		defer methodName.Destroy()
		ptrsForTextEdit.fnRemoveCaret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("remove_secondary_carets")
		defer methodName.Destroy()
		ptrsForTextEdit.fnRemoveSecondaryCarets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("merge_overlapping_carets")
		defer methodName.Destroy()
		ptrsForTextEdit.fnMergeOverlappingCarets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_caret_count")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetCaretCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_caret_at_carets")
		defer methodName.Destroy()
		ptrsForTextEdit.fnAddCaretAtCarets = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_caret_index_edit_order")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetCaretIndexEditOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("adjust_carets_after_edit")
		defer methodName.Destroy()
		ptrsForTextEdit.fnAdjustCaretsAfterEdit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1770277138))
	}
	{
		methodName := StringNameFromStr("is_caret_visible")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsCaretVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1051549951))
	}
	{
		methodName := StringNameFromStr("get_caret_draw_pos")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetCaretDrawPos = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 478253731))
	}
	{
		methodName := StringNameFromStr("set_caret_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetCaretLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1302582944))
	}
	{
		methodName := StringNameFromStr("get_caret_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetCaretLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("set_caret_column")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetCaretColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3796796178))
	}
	{
		methodName := StringNameFromStr("get_caret_column")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetCaretColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("get_caret_wrap_index")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetCaretWrapIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("get_word_under_caret")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetWordUnderCaret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3929349208))
	}
	{
		methodName := StringNameFromStr("set_selecting_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetSelectingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_selecting_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsSelectingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_deselect_on_focus_loss_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetDeselectOnFocusLossEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_deselect_on_focus_loss_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsDeselectOnFocusLossEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_drag_and_drop_selection_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetDragAndDropSelectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_drag_and_drop_selection_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsDragAndDropSelectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_selection_mode")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetSelectionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1443345937))
	}
	{
		methodName := StringNameFromStr("get_selection_mode")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSelectionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3750106938))
	}
	{
		methodName := StringNameFromStr("select_all")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSelectAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("select_word_under_caret")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSelectWordUnderCaret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("add_selection_for_next_occurrence")
		defer methodName.Destroy()
		ptrsForTextEdit.fnAddSelectionForNextOccurrence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("skip_selection_for_next_occurrence")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSkipSelectionForNextOccurrence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("select")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSelect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2560984452))
	}
	{
		methodName := StringNameFromStr("has_selection")
		defer methodName.Destroy()
		ptrsForTextEdit.fnHasSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2824505868))
	}
	{
		methodName := StringNameFromStr("get_selected_text")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSelectedText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2309358862))
	}
	{
		methodName := StringNameFromStr("get_selection_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSelectionLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("get_selection_column")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSelectionColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("get_selection_from_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSelectionFromLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("get_selection_from_column")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSelectionFromColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("get_selection_to_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSelectionToLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("get_selection_to_column")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSelectionToColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("deselect")
		defer methodName.Destroy()
		ptrsForTextEdit.fnDeselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("delete_selection")
		defer methodName.Destroy()
		ptrsForTextEdit.fnDeleteSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("set_line_wrapping_mode")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineWrappingMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2525115309))
	}
	{
		methodName := StringNameFromStr("get_line_wrapping_mode")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineWrappingMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3562716114))
	}
	{
		methodName := StringNameFromStr("set_autowrap_mode")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3289138044))
	}
	{
		methodName := StringNameFromStr("get_autowrap_mode")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1549071663))
	}
	{
		methodName := StringNameFromStr("is_line_wrapped")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsLineWrapped = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_line_wrap_count")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineWrapCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_line_wrap_index_at_column")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineWrapIndexAtColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("get_line_wrapped_text")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineWrappedText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 647634434))
	}
	{
		methodName := StringNameFromStr("set_smooth_scroll_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetSmoothScrollEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_smooth_scroll_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsSmoothScrollEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_v_scroll_bar")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetVScrollBar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3226026593))
	}
	{
		methodName := StringNameFromStr("get_h_scroll_bar")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetHScrollBar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3774687988))
	}
	{
		methodName := StringNameFromStr("set_v_scroll")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetVScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_v_scroll")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetVScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_h_scroll")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetHScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_h_scroll")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetHScroll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_scroll_past_end_of_file_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetScrollPastEndOfFileEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_scroll_past_end_of_file_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsScrollPastEndOfFileEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_v_scroll_speed")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetVScrollSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_v_scroll_speed")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetVScrollSpeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_fit_content_height_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetFitContentHeightEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_fit_content_height_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsFitContentHeightEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_scroll_pos_for_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetScrollPosForLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3929084198))
	}
	{
		methodName := StringNameFromStr("set_line_as_first_visible")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineAsFirstVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2230941749))
	}
	{
		methodName := StringNameFromStr("get_first_visible_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetFirstVisibleLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_line_as_center_visible")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineAsCenterVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2230941749))
	}
	{
		methodName := StringNameFromStr("set_line_as_last_visible")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineAsLastVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2230941749))
	}
	{
		methodName := StringNameFromStr("get_last_full_visible_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLastFullVisibleLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_last_full_visible_line_wrap_index")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLastFullVisibleLineWrapIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_visible_line_count")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetVisibleLineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_visible_line_count_in_range")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetVisibleLineCountInRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3175239445))
	}
	{
		methodName := StringNameFromStr("get_total_visible_line_count")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetTotalVisibleLineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("adjust_viewport_to_caret")
		defer methodName.Destroy()
		ptrsForTextEdit.fnAdjustViewportToCaret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
	}
	{
		methodName := StringNameFromStr("center_viewport_to_caret")
		defer methodName.Destroy()
		ptrsForTextEdit.fnCenterViewportToCaret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1995695955))
	}
	{
		methodName := StringNameFromStr("set_draw_minimap")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetDrawMinimap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_drawing_minimap")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsDrawingMinimap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_minimap_width")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetMinimapWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_minimap_width")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetMinimapWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_minimap_visible_lines")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetMinimapVisibleLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_gutter")
		defer methodName.Destroy()
		ptrsForTextEdit.fnAddGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
	}
	{
		methodName := StringNameFromStr("remove_gutter")
		defer methodName.Destroy()
		ptrsForTextEdit.fnRemoveGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_gutter_count")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetGutterCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_gutter_name")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetGutterName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_gutter_name")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetGutterName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_gutter_type")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetGutterType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1088959071))
	}
	{
		methodName := StringNameFromStr("get_gutter_type")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetGutterType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1159699127))
	}
	{
		methodName := StringNameFromStr("set_gutter_width")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetGutterWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_gutter_width")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetGutterWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_gutter_draw")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetGutterDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_gutter_drawn")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsGutterDrawn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_gutter_clickable")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetGutterClickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_gutter_clickable")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsGutterClickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_gutter_overwritable")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetGutterOverwritable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_gutter_overwritable")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsGutterOverwritable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("merge_gutters")
		defer methodName.Destroy()
		ptrsForTextEdit.fnMergeGutters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_gutter_custom_draw")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetGutterCustomDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 957362965))
	}
	{
		methodName := StringNameFromStr("get_total_gutter_width")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetTotalGutterWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_line_gutter_metadata")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineGutterMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2060538656))
	}
	{
		methodName := StringNameFromStr("get_line_gutter_metadata")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineGutterMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 678354945))
	}
	{
		methodName := StringNameFromStr("set_line_gutter_text")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineGutterText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2285447957))
	}
	{
		methodName := StringNameFromStr("get_line_gutter_text")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineGutterText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1391810591))
	}
	{
		methodName := StringNameFromStr("set_line_gutter_icon")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineGutterIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 176101966))
	}
	{
		methodName := StringNameFromStr("get_line_gutter_icon")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineGutterIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2584904275))
	}
	{
		methodName := StringNameFromStr("set_line_gutter_item_color")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineGutterItemColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3733378741))
	}
	{
		methodName := StringNameFromStr("get_line_gutter_item_color")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineGutterItemColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2165839948))
	}
	{
		methodName := StringNameFromStr("set_line_gutter_clickable")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineGutterClickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1383440665))
	}
	{
		methodName := StringNameFromStr("is_line_gutter_clickable")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsLineGutterClickable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
	}
	{
		methodName := StringNameFromStr("set_line_background_color")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetLineBackgroundColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_line_background_color")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetLineBackgroundColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("set_syntax_highlighter")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetSyntaxHighlighter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2765644541))
	}
	{
		methodName := StringNameFromStr("get_syntax_highlighter")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetSyntaxHighlighter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2721131626))
	}
	{
		methodName := StringNameFromStr("set_highlight_current_line")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetHighlightCurrentLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_highlight_current_line_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsHighlightCurrentLineEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_highlight_all_occurrences")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetHighlightAllOccurrences = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_highlight_all_occurrences_enabled")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsHighlightAllOccurrencesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_draw_control_chars")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetDrawControlChars = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_draw_control_chars")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetDrawControlChars = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_draw_tabs")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetDrawTabs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_drawing_tabs")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsDrawingTabs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_draw_spaces")
		defer methodName.Destroy()
		ptrsForTextEdit.fnSetDrawSpaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_drawing_spaces")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsDrawingSpaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_menu")
		defer methodName.Destroy()
		ptrsForTextEdit.fnGetMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 229722558))
	}
	{
		methodName := StringNameFromStr("is_menu_visible")
		defer methodName.Destroy()
		ptrsForTextEdit.fnIsMenuVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("menu_option")
		defer methodName.Destroy()
		ptrsForTextEdit.fnMenuOption = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}

}

type TextEdit struct {
	Control
}

func (me *TextEdit) BaseClass() string {
	return "TextEdit"
}

func NewTextEdit() *TextEdit {
	str := StringNameFromStr("TextEdit") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TextEdit{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TextEditMenuItems int

const (
	TextEditMenuItemsMenuCut              TextEditMenuItems = 0
	TextEditMenuItemsMenuCopy             TextEditMenuItems = 1
	TextEditMenuItemsMenuPaste            TextEditMenuItems = 2
	TextEditMenuItemsMenuClear            TextEditMenuItems = 3
	TextEditMenuItemsMenuSelectAll        TextEditMenuItems = 4
	TextEditMenuItemsMenuUndo             TextEditMenuItems = 5
	TextEditMenuItemsMenuRedo             TextEditMenuItems = 6
	TextEditMenuItemsMenuSubmenuTextDir   TextEditMenuItems = 7
	TextEditMenuItemsMenuDirInherited     TextEditMenuItems = 8
	TextEditMenuItemsMenuDirAuto          TextEditMenuItems = 9
	TextEditMenuItemsMenuDirLtr           TextEditMenuItems = 10
	TextEditMenuItemsMenuDirRtl           TextEditMenuItems = 11
	TextEditMenuItemsMenuDisplayUcc       TextEditMenuItems = 12
	TextEditMenuItemsMenuSubmenuInsertUcc TextEditMenuItems = 13
	TextEditMenuItemsMenuInsertLrm        TextEditMenuItems = 14
	TextEditMenuItemsMenuInsertRlm        TextEditMenuItems = 15
	TextEditMenuItemsMenuInsertLre        TextEditMenuItems = 16
	TextEditMenuItemsMenuInsertRle        TextEditMenuItems = 17
	TextEditMenuItemsMenuInsertLro        TextEditMenuItems = 18
	TextEditMenuItemsMenuInsertRlo        TextEditMenuItems = 19
	TextEditMenuItemsMenuInsertPdf        TextEditMenuItems = 20
	TextEditMenuItemsMenuInsertAlm        TextEditMenuItems = 21
	TextEditMenuItemsMenuInsertLri        TextEditMenuItems = 22
	TextEditMenuItemsMenuInsertRli        TextEditMenuItems = 23
	TextEditMenuItemsMenuInsertFsi        TextEditMenuItems = 24
	TextEditMenuItemsMenuInsertPdi        TextEditMenuItems = 25
	TextEditMenuItemsMenuInsertZwj        TextEditMenuItems = 26
	TextEditMenuItemsMenuInsertZwnj       TextEditMenuItems = 27
	TextEditMenuItemsMenuInsertWj         TextEditMenuItems = 28
	TextEditMenuItemsMenuInsertShy        TextEditMenuItems = 29
	TextEditMenuItemsMenuMax              TextEditMenuItems = 30
)

type TextEditEditAction int

const (
	TextEditEditActionActionNone      TextEditEditAction = 0
	TextEditEditActionActionTyping    TextEditEditAction = 1
	TextEditEditActionActionBackspace TextEditEditAction = 2
	TextEditEditActionActionDelete    TextEditEditAction = 3
)

type TextEditSearchFlags int

const (
	TextEditSearchFlagsSearchMatchCase  TextEditSearchFlags = 1
	TextEditSearchFlagsSearchWholeWords TextEditSearchFlags = 2
	TextEditSearchFlagsSearchBackwards  TextEditSearchFlags = 4
)

type TextEditCaretType int

const (
	TextEditCaretTypeCaretTypeLine  TextEditCaretType = 0
	TextEditCaretTypeCaretTypeBlock TextEditCaretType = 1
)

type TextEditSelectionMode int

const (
	TextEditSelectionModeSelectionModeNone    TextEditSelectionMode = 0
	TextEditSelectionModeSelectionModeShift   TextEditSelectionMode = 1
	TextEditSelectionModeSelectionModePointer TextEditSelectionMode = 2
	TextEditSelectionModeSelectionModeWord    TextEditSelectionMode = 3
	TextEditSelectionModeSelectionModeLine    TextEditSelectionMode = 4
)

type TextEditLineWrappingMode int

const (
	TextEditLineWrappingModeLineWrappingNone     TextEditLineWrappingMode = 0
	TextEditLineWrappingModeLineWrappingBoundary TextEditLineWrappingMode = 1
)

type TextEditGutterType int

const (
	TextEditGutterTypeGutterTypeString TextEditGutterType = 0
	TextEditGutterTypeGutterTypeIcon   TextEditGutterType = 1
	TextEditGutterTypeGutterTypeCustom TextEditGutterType = 2
)

func (me *TextEdit) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TextEdit) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TextEdit) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TextEdit) HasImeText() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnHasImeText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) CancelIme() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnCancelIme), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) ApplyIme() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnApplyIme), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SetEditable(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetEditable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsEditable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsEditable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetTextDirection(direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetTextDirection() ControlTextDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextEdit) SetLanguage(language String) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLanguage() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerStructuredTextParser

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextEdit) SetStructuredTextBidiOverrideOptions(args Array) {
	cargs := []gdc.ConstTypePtr{args.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetStructuredTextBidiOverrideOptions() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetTabSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetTabSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetTabSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetTabSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetIndentWrappedLines(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetIndentWrappedLines), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsIndentWrappedLines() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsIndentWrappedLines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetOvertypeModeEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetOvertypeModeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsOvertypeModeEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsOvertypeModeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetContextMenuEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetContextMenuEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsContextMenuEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsContextMenuEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetShortcutKeysEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetShortcutKeysEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsShortcutKeysEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsShortcutKeysEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetVirtualKeyboardEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetVirtualKeyboardEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsVirtualKeyboardEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsVirtualKeyboardEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetMiddleMousePasteEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetMiddleMousePasteEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsMiddleMousePasteEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsMiddleMousePasteEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SetText(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) GetLineCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetPlaceholder(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetPlaceholder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetPlaceholder() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetPlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetLine(line int64, new_text String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), new_text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLine), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLine(line int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) GetLineWidth(line int64, wrap_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&wrap_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&line)
	pinner.Pin(&wrap_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetLineHeight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetIndentLevel(line int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetIndentLevel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetFirstNonWhitespaceColumn(line int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetFirstNonWhitespaceColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SwapLines(from_line int64, to_line int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line), gdc.ConstTypePtr(&to_line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSwapLines), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) InsertLineAt(line int64, text String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnInsertLineAt), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) InsertTextAtCaret(text String, caret_index int64) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnInsertTextAtCaret), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) RemoveText(from_line int64, from_column int64, to_line int64, to_column int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line), gdc.ConstTypePtr(&from_column), gdc.ConstTypePtr(&to_line), gdc.ConstTypePtr(&to_column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnRemoveText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLastUnhiddenLine() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLastUnhiddenLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetNextVisibleLineOffsetFrom(line int64, visible_amount int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&visible_amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&line)
	pinner.Pin(&visible_amount)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetNextVisibleLineOffsetFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetNextVisibleLineIndexOffsetFrom(line int64, wrap_index int64, visible_amount int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&wrap_index), gdc.ConstTypePtr(&visible_amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&line)
	pinner.Pin(&wrap_index)
	pinner.Pin(&visible_amount)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetNextVisibleLineIndexOffsetFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) Backspace(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnBackspace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) Cut(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnCut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) Copy(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnCopy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) Paste(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnPaste), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) PastePrimaryClipboard(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnPastePrimaryClipboard), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) StartAction(action TextEditEditAction) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&action)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnStartAction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) EndAction() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnEndAction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) BeginComplexOperation() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnBeginComplexOperation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) EndComplexOperation() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnEndComplexOperation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) HasUndo() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnHasUndo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) HasRedo() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnHasRedo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) Undo() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnUndo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) Redo() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnRedo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) ClearUndoHistory() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnClearUndoHistory), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) TagSavedVersion() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnTagSavedVersion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetVersion() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetVersion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetSavedVersion() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSavedVersion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetSearchText(search_text String) {
	cargs := []gdc.ConstTypePtr{search_text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetSearchText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SetSearchFlags(flags int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetSearchFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) Search(text String, flags int64, from_line int64, from_colum int64) Vector2i {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), gdc.ConstTypePtr(&flags), gdc.ConstTypePtr(&from_line), gdc.ConstTypePtr(&from_colum)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&flags)
	pinner.Pin(&from_line)
	pinner.Pin(&from_colum)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSearch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetTooltipRequestFunc(callback Callable) {
	cargs := []gdc.ConstTypePtr{callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetTooltipRequestFunc), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLocalMousePos() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLocalMousePos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) GetWordAtPos(position Vector2) String {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetWordAtPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) GetLineColumnAtPos(position Vector2i, allow_out_of_bounds bool) Vector2i {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&allow_out_of_bounds)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&allow_out_of_bounds)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineColumnAtPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) GetPosAtLineColumn(line int64, column int64) Vector2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()
	pinner.Pin(&line)
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetPosAtLineColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) GetRectAtLineColumn(line int64, column int64) Rect2i {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2i()
	pinner.Pin(&line)
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetRectAtLineColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) GetMinimapLineAtPos(position Vector2i) int64 {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetMinimapLineAtPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) IsDraggingCursor() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsDraggingCursor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) IsMouseOverSelection(edges bool, caret_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edges), gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&edges)
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsMouseOverSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetCaretType(type_ TextEditCaretType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetCaretType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetCaretType() TextEditCaretType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextEditCaretType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetCaretType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextEdit) SetCaretBlinkEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetCaretBlinkEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsCaretBlinkEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsCaretBlinkEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetCaretBlinkInterval(interval float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetCaretBlinkInterval), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetCaretBlinkInterval() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetCaretBlinkInterval), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetDrawCaretWhenEditableDisabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetDrawCaretWhenEditableDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsDrawingCaretWhenEditableDisabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsDrawingCaretWhenEditableDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetMoveCaretOnRightClickEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetMoveCaretOnRightClickEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsMoveCaretOnRightClickEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsMoveCaretOnRightClickEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetCaretMidGraphemeEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetCaretMidGraphemeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsCaretMidGraphemeEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsCaretMidGraphemeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetMultipleCaretsEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetMultipleCaretsEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsMultipleCaretsEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsMultipleCaretsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) AddCaret(line int64, col int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&col)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&line)
	pinner.Pin(&col)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnAddCaret), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) RemoveCaret(caret int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnRemoveCaret), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) RemoveSecondaryCarets() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnRemoveSecondaryCarets), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) MergeOverlappingCarets() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnMergeOverlappingCarets), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetCaretCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetCaretCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) AddCaretAtCarets(below bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&below)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnAddCaretAtCarets), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetCaretIndexEditOrder() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetCaretIndexEditOrder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) AdjustCaretsAfterEdit(caret int64, from_line int64, from_col int64, to_line int64, to_col int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret), gdc.ConstTypePtr(&from_line), gdc.ConstTypePtr(&from_col), gdc.ConstTypePtr(&to_line), gdc.ConstTypePtr(&to_col)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnAdjustCaretsAfterEdit), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsCaretVisible(caret_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsCaretVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetCaretDrawPos(caret_index int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetCaretDrawPos), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetCaretLine(line int64, adjust_viewport bool, can_be_hidden bool, wrap_index int64, caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&adjust_viewport), gdc.ConstTypePtr(&can_be_hidden), gdc.ConstTypePtr(&wrap_index), gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetCaretLine), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetCaretLine(caret_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetCaretLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetCaretColumn(column int64, adjust_viewport bool, caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&adjust_viewport), gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetCaretColumn), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetCaretColumn(caret_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetCaretColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetCaretWrapIndex(caret_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetCaretWrapIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetWordUnderCaret(caret_index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetWordUnderCaret), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetSelectingEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetSelectingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsSelectingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsSelectingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetDeselectOnFocusLossEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetDeselectOnFocusLossEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsDeselectOnFocusLossEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsDeselectOnFocusLossEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetDragAndDropSelectionEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetDragAndDropSelectionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsDragAndDropSelectionEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsDragAndDropSelectionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetSelectionMode(mode TextEditSelectionMode, line int64, column int64, caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetSelectionMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetSelectionMode() TextEditSelectionMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextEditSelectionMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSelectionMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextEdit) SelectAll() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSelectAll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SelectWordUnderCaret(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSelectWordUnderCaret), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) AddSelectionForNextOccurrence() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnAddSelectionForNextOccurrence), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SkipSelectionForNextOccurrence() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSkipSelectionForNextOccurrence), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) Select(from_line int64, from_column int64, to_line int64, to_column int64, caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line), gdc.ConstTypePtr(&from_column), gdc.ConstTypePtr(&to_line), gdc.ConstTypePtr(&to_column), gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSelect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) HasSelection(caret_index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnHasSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetSelectedText(caret_index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSelectedText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) GetSelectionLine(caret_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSelectionLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetSelectionColumn(caret_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSelectionColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetSelectionFromLine(caret_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSelectionFromLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetSelectionFromColumn(caret_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSelectionFromColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetSelectionToLine(caret_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSelectionToLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetSelectionToColumn(caret_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&caret_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSelectionToColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) Deselect(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnDeselect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) DeleteSelection(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnDeleteSelection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SetLineWrappingMode(mode TextEditLineWrappingMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineWrappingMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLineWrappingMode() TextEditLineWrappingMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextEditLineWrappingMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineWrappingMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextEdit) SetAutowrapMode(autowrap_mode TextServerAutowrapMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetAutowrapMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetAutowrapMode() TextServerAutowrapMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerAutowrapMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetAutowrapMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextEdit) IsLineWrapped(line int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsLineWrapped), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetLineWrapCount(line int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineWrapCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetLineWrapIndexAtColumn(line int64, column int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&column)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&line)
	pinner.Pin(&column)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineWrapIndexAtColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetLineWrappedText(line int64) PackedStringArray {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineWrappedText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetSmoothScrollEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetSmoothScrollEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsSmoothScrollEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsSmoothScrollEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetVScrollBar() VScrollBar {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVScrollBar()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetVScrollBar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) GetHScrollBar() HScrollBar {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewHScrollBar()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetHScrollBar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetVScroll(value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetVScroll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetVScroll() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetVScroll), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetHScroll(value int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetHScroll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetHScroll() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetHScroll), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetScrollPastEndOfFileEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetScrollPastEndOfFileEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsScrollPastEndOfFileEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsScrollPastEndOfFileEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetVScrollSpeed(speed float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&speed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetVScrollSpeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetVScrollSpeed() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetVScrollSpeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetFitContentHeightEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetFitContentHeightEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsFitContentHeightEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsFitContentHeightEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetScrollPosForLine(line int64, wrap_index int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&wrap_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&line)
	pinner.Pin(&wrap_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetScrollPosForLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetLineAsFirstVisible(line int64, wrap_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&wrap_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineAsFirstVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetFirstVisibleLine() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetFirstVisibleLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetLineAsCenterVisible(line int64, wrap_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&wrap_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineAsCenterVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SetLineAsLastVisible(line int64, wrap_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&wrap_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineAsLastVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLastFullVisibleLine() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLastFullVisibleLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetLastFullVisibleLineWrapIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLastFullVisibleLineWrapIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetVisibleLineCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetVisibleLineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetVisibleLineCountInRange(from_line int64, to_line int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line), gdc.ConstTypePtr(&to_line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&from_line)
	pinner.Pin(&to_line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetVisibleLineCountInRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetTotalVisibleLineCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetTotalVisibleLineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) AdjustViewportToCaret(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnAdjustViewportToCaret), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) CenterViewportToCaret(caret_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&caret_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnCenterViewportToCaret), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SetDrawMinimap(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetDrawMinimap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsDrawingMinimap() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsDrawingMinimap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetMinimapWidth(width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetMinimapWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetMinimapWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetMinimapWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetMinimapVisibleLines() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetMinimapVisibleLines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) AddGutter(at int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&at)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnAddGutter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) RemoveGutter(gutter int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnRemoveGutter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetGutterCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetGutterCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetGutterName(gutter int64, name String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetGutterName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetGutterName(gutter int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetGutterName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetGutterType(gutter int64, type_ TextEditGutterType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter), gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetGutterType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetGutterType(gutter int64) TextEditGutterType {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextEditGutterType
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetGutterType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TextEdit) SetGutterWidth(gutter int64, width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter), gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetGutterWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetGutterWidth(gutter int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetGutterWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetGutterDraw(gutter int64, draw bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter), gdc.ConstTypePtr(&draw)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetGutterDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsGutterDrawn(gutter int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsGutterDrawn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetGutterClickable(gutter int64, clickable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter), gdc.ConstTypePtr(&clickable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetGutterClickable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsGutterClickable(gutter int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsGutterClickable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetGutterOverwritable(gutter int64, overwritable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter), gdc.ConstTypePtr(&overwritable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetGutterOverwritable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsGutterOverwritable(gutter int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsGutterOverwritable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) MergeGutters(from_line int64, to_line int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line), gdc.ConstTypePtr(&to_line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnMergeGutters), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SetGutterCustomDraw(column int64, draw_callback Callable) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), draw_callback.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetGutterCustomDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetTotalGutterWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetTotalGutterWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetLineGutterMetadata(line int64, gutter int64, metadata Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter), metadata.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineGutterMetadata), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLineGutterMetadata(line int64, gutter int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&line)
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineGutterMetadata), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetLineGutterText(line int64, gutter int64, text String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineGutterText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLineGutterText(line int64, gutter int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&line)
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineGutterText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetLineGutterIcon(line int64, gutter int64, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineGutterIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLineGutterIcon(line int64, gutter int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&line)
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineGutterIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetLineGutterItemColor(line int64, gutter int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineGutterItemColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLineGutterItemColor(line int64, gutter int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&line)
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineGutterItemColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetLineGutterClickable(line int64, gutter int64, clickable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter), gdc.ConstTypePtr(&clickable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineGutterClickable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsLineGutterClickable(line int64, gutter int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), gdc.ConstTypePtr(&gutter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&line)
	pinner.Pin(&gutter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsLineGutterClickable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetLineBackgroundColor(line int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetLineBackgroundColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetLineBackgroundColor(line int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetLineBackgroundColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetSyntaxHighlighter(syntax_highlighter SyntaxHighlighter) {
	cargs := []gdc.ConstTypePtr{syntax_highlighter.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetSyntaxHighlighter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) GetSyntaxHighlighter() SyntaxHighlighter {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSyntaxHighlighter()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetSyntaxHighlighter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) SetHighlightCurrentLine(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetHighlightCurrentLine), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsHighlightCurrentLineEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsHighlightCurrentLineEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetHighlightAllOccurrences(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetHighlightAllOccurrences), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsHighlightAllOccurrencesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsHighlightAllOccurrencesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetDrawControlChars() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetDrawControlChars), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetDrawControlChars(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetDrawControlChars), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) SetDrawTabs(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetDrawTabs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsDrawingTabs() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsDrawingTabs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) SetDrawSpaces(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnSetDrawSpaces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TextEdit) IsDrawingSpaces() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsDrawingSpaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) GetMenu() PopupMenu {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPopupMenu()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnGetMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TextEdit) IsMenuVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnIsMenuVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TextEdit) MenuOption(option int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTextEdit.fnMenuOption), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TextEditTextSetSignalFn func()

func (me *TextEdit) ConnectTextSet(subs SignalSubscribers, fn TextEditTextSetSignalFn) {
	sig := StringNameFromStr("text_set")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectTextSet(subs SignalSubscribers, fn TextEditTextSetSignalFn) {
	sig := StringNameFromStr("text_set")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditTextChangedSignalFn func()

func (me *TextEdit) ConnectTextChanged(subs SignalSubscribers, fn TextEditTextChangedSignalFn) {
	sig := StringNameFromStr("text_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectTextChanged(subs SignalSubscribers, fn TextEditTextChangedSignalFn) {
	sig := StringNameFromStr("text_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditLinesEditedFromSignalFn func(from_line int, to_line int)

func (me *TextEdit) ConnectLinesEditedFrom(subs SignalSubscribers, fn TextEditLinesEditedFromSignalFn) {
	sig := StringNameFromStr("lines_edited_from")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectLinesEditedFrom(subs SignalSubscribers, fn TextEditLinesEditedFromSignalFn) {
	sig := StringNameFromStr("lines_edited_from")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditCaretChangedSignalFn func()

func (me *TextEdit) ConnectCaretChanged(subs SignalSubscribers, fn TextEditCaretChangedSignalFn) {
	sig := StringNameFromStr("caret_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectCaretChanged(subs SignalSubscribers, fn TextEditCaretChangedSignalFn) {
	sig := StringNameFromStr("caret_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditGutterClickedSignalFn func(line int, gutter int)

func (me *TextEdit) ConnectGutterClicked(subs SignalSubscribers, fn TextEditGutterClickedSignalFn) {
	sig := StringNameFromStr("gutter_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectGutterClicked(subs SignalSubscribers, fn TextEditGutterClickedSignalFn) {
	sig := StringNameFromStr("gutter_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditGutterAddedSignalFn func()

func (me *TextEdit) ConnectGutterAdded(subs SignalSubscribers, fn TextEditGutterAddedSignalFn) {
	sig := StringNameFromStr("gutter_added")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectGutterAdded(subs SignalSubscribers, fn TextEditGutterAddedSignalFn) {
	sig := StringNameFromStr("gutter_added")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TextEditGutterRemovedSignalFn func()

func (me *TextEdit) ConnectGutterRemoved(subs SignalSubscribers, fn TextEditGutterRemovedSignalFn) {
	sig := StringNameFromStr("gutter_removed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TextEdit) DisconnectGutterRemoved(subs SignalSubscribers, fn TextEditGutterRemovedSignalFn) {
	sig := StringNameFromStr("gutter_removed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
