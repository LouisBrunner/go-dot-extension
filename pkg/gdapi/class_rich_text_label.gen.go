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

type ptrsForRichTextLabelList struct {
	fnGetParsedText                        gdc.MethodBindPtr
	fnAddText                              gdc.MethodBindPtr
	fnSetText                              gdc.MethodBindPtr
	fnAddImage                             gdc.MethodBindPtr
	fnUpdateImage                          gdc.MethodBindPtr
	fnNewline                              gdc.MethodBindPtr
	fnRemoveParagraph                      gdc.MethodBindPtr
	fnPushFont                             gdc.MethodBindPtr
	fnPushFontSize                         gdc.MethodBindPtr
	fnPushNormal                           gdc.MethodBindPtr
	fnPushBold                             gdc.MethodBindPtr
	fnPushBoldItalics                      gdc.MethodBindPtr
	fnPushItalics                          gdc.MethodBindPtr
	fnPushMono                             gdc.MethodBindPtr
	fnPushColor                            gdc.MethodBindPtr
	fnPushOutlineSize                      gdc.MethodBindPtr
	fnPushOutlineColor                     gdc.MethodBindPtr
	fnPushParagraph                        gdc.MethodBindPtr
	fnPushIndent                           gdc.MethodBindPtr
	fnPushList                             gdc.MethodBindPtr
	fnPushMeta                             gdc.MethodBindPtr
	fnPushHint                             gdc.MethodBindPtr
	fnPushLanguage                         gdc.MethodBindPtr
	fnPushUnderline                        gdc.MethodBindPtr
	fnPushStrikethrough                    gdc.MethodBindPtr
	fnPushTable                            gdc.MethodBindPtr
	fnPushDropcap                          gdc.MethodBindPtr
	fnSetTableColumnExpand                 gdc.MethodBindPtr
	fnSetCellRowBackgroundColor            gdc.MethodBindPtr
	fnSetCellBorderColor                   gdc.MethodBindPtr
	fnSetCellSizeOverride                  gdc.MethodBindPtr
	fnSetCellPadding                       gdc.MethodBindPtr
	fnPushCell                             gdc.MethodBindPtr
	fnPushFgcolor                          gdc.MethodBindPtr
	fnPushBgcolor                          gdc.MethodBindPtr
	fnPushCustomfx                         gdc.MethodBindPtr
	fnPushContext                          gdc.MethodBindPtr
	fnPopContext                           gdc.MethodBindPtr
	fnPop                                  gdc.MethodBindPtr
	fnPopAll                               gdc.MethodBindPtr
	fnClear                                gdc.MethodBindPtr
	fnSetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnGetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnSetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnGetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnSetTextDirection                     gdc.MethodBindPtr
	fnGetTextDirection                     gdc.MethodBindPtr
	fnSetLanguage                          gdc.MethodBindPtr
	fnGetLanguage                          gdc.MethodBindPtr
	fnSetAutowrapMode                      gdc.MethodBindPtr
	fnGetAutowrapMode                      gdc.MethodBindPtr
	fnSetMetaUnderline                     gdc.MethodBindPtr
	fnIsMetaUnderlined                     gdc.MethodBindPtr
	fnSetHintUnderline                     gdc.MethodBindPtr
	fnIsHintUnderlined                     gdc.MethodBindPtr
	fnSetScrollActive                      gdc.MethodBindPtr
	fnIsScrollActive                       gdc.MethodBindPtr
	fnSetScrollFollow                      gdc.MethodBindPtr
	fnIsScrollFollowing                    gdc.MethodBindPtr
	fnGetVScrollBar                        gdc.MethodBindPtr
	fnScrollToLine                         gdc.MethodBindPtr
	fnScrollToParagraph                    gdc.MethodBindPtr
	fnScrollToSelection                    gdc.MethodBindPtr
	fnSetTabSize                           gdc.MethodBindPtr
	fnGetTabSize                           gdc.MethodBindPtr
	fnSetFitContent                        gdc.MethodBindPtr
	fnIsFitContentEnabled                  gdc.MethodBindPtr
	fnSetSelectionEnabled                  gdc.MethodBindPtr
	fnIsSelectionEnabled                   gdc.MethodBindPtr
	fnSetContextMenuEnabled                gdc.MethodBindPtr
	fnIsContextMenuEnabled                 gdc.MethodBindPtr
	fnSetShortcutKeysEnabled               gdc.MethodBindPtr
	fnIsShortcutKeysEnabled                gdc.MethodBindPtr
	fnSetDeselectOnFocusLossEnabled        gdc.MethodBindPtr
	fnIsDeselectOnFocusLossEnabled         gdc.MethodBindPtr
	fnSetDragAndDropSelectionEnabled       gdc.MethodBindPtr
	fnIsDragAndDropSelectionEnabled        gdc.MethodBindPtr
	fnGetSelectionFrom                     gdc.MethodBindPtr
	fnGetSelectionTo                       gdc.MethodBindPtr
	fnSelectAll                            gdc.MethodBindPtr
	fnGetSelectedText                      gdc.MethodBindPtr
	fnDeselect                             gdc.MethodBindPtr
	fnParseBbcode                          gdc.MethodBindPtr
	fnAppendText                           gdc.MethodBindPtr
	fnGetText                              gdc.MethodBindPtr
	fnIsReady                              gdc.MethodBindPtr
	fnSetThreaded                          gdc.MethodBindPtr
	fnIsThreaded                           gdc.MethodBindPtr
	fnSetProgressBarDelay                  gdc.MethodBindPtr
	fnGetProgressBarDelay                  gdc.MethodBindPtr
	fnSetVisibleCharacters                 gdc.MethodBindPtr
	fnGetVisibleCharacters                 gdc.MethodBindPtr
	fnGetVisibleCharactersBehavior         gdc.MethodBindPtr
	fnSetVisibleCharactersBehavior         gdc.MethodBindPtr
	fnSetVisibleRatio                      gdc.MethodBindPtr
	fnGetVisibleRatio                      gdc.MethodBindPtr
	fnGetCharacterLine                     gdc.MethodBindPtr
	fnGetCharacterParagraph                gdc.MethodBindPtr
	fnGetTotalCharacterCount               gdc.MethodBindPtr
	fnSetUseBbcode                         gdc.MethodBindPtr
	fnIsUsingBbcode                        gdc.MethodBindPtr
	fnGetLineCount                         gdc.MethodBindPtr
	fnGetVisibleLineCount                  gdc.MethodBindPtr
	fnGetParagraphCount                    gdc.MethodBindPtr
	fnGetVisibleParagraphCount             gdc.MethodBindPtr
	fnGetContentHeight                     gdc.MethodBindPtr
	fnGetContentWidth                      gdc.MethodBindPtr
	fnGetLineOffset                        gdc.MethodBindPtr
	fnGetParagraphOffset                   gdc.MethodBindPtr
	fnParseExpressionsForValues            gdc.MethodBindPtr
	fnSetEffects                           gdc.MethodBindPtr
	fnGetEffects                           gdc.MethodBindPtr
	fnInstallEffect                        gdc.MethodBindPtr
	fnGetMenu                              gdc.MethodBindPtr
	fnIsMenuVisible                        gdc.MethodBindPtr
	fnMenuOption                           gdc.MethodBindPtr
}

var ptrsForRichTextLabel ptrsForRichTextLabelList

func initRichTextLabelPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RichTextLabel")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_parsed_text")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetParsedText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("add_text")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnAddText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_text")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("add_image")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnAddImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3017663154))
	}
	{
		methodName := StringNameFromStr("update_image")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnUpdateImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 815048486))
	}
	{
		methodName := StringNameFromStr("newline")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnNewline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("remove_paragraph")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnRemoveParagraph = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3067735520))
	}
	{
		methodName := StringNameFromStr("push_font")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushFont = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2347424842))
	}
	{
		methodName := StringNameFromStr("push_font_size")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushFontSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("push_normal")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("push_bold")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushBold = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("push_bold_italics")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushBoldItalics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("push_italics")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushItalics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("push_mono")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushMono = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("push_color")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("push_outline_size")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushOutlineSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("push_outline_color")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushOutlineColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("push_paragraph")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushParagraph = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3089306873))
	}
	{
		methodName := StringNameFromStr("push_indent")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushIndent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("push_list")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3017143144))
	}
	{
		methodName := StringNameFromStr("push_meta")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushMeta = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1114965689))
	}
	{
		methodName := StringNameFromStr("push_hint")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("push_language")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("push_underline")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushUnderline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("push_strikethrough")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushStrikethrough = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("push_table")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushTable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2623499273))
	}
	{
		methodName := StringNameFromStr("push_dropcap")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushDropcap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4061635501))
	}
	{
		methodName := StringNameFromStr("set_table_column_expand")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetTableColumnExpand = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2185176273))
	}
	{
		methodName := StringNameFromStr("set_cell_row_background_color")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetCellRowBackgroundColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3465483165))
	}
	{
		methodName := StringNameFromStr("set_cell_border_color")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetCellBorderColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("set_cell_size_override")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetCellSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3108078480))
	}
	{
		methodName := StringNameFromStr("set_cell_padding")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetCellPadding = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2046264180))
	}
	{
		methodName := StringNameFromStr("push_cell")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushCell = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("push_fgcolor")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushFgcolor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("push_bgcolor")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushBgcolor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("push_customfx")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushCustomfx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2337942958))
	}
	{
		methodName := StringNameFromStr("push_context")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPushContext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("pop_context")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPopContext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("pop")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("pop_all")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnPopAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 55961453))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3385126229))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_text_direction")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 119160795))
	}
	{
		methodName := StringNameFromStr("get_text_direction")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797257663))
	}
	{
		methodName := StringNameFromStr("set_language")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_language")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_autowrap_mode")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3289138044))
	}
	{
		methodName := StringNameFromStr("get_autowrap_mode")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1549071663))
	}
	{
		methodName := StringNameFromStr("set_meta_underline")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetMetaUnderline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_meta_underlined")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsMetaUnderlined = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hint_underline")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetHintUnderline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hint_underlined")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsHintUnderlined = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_scroll_active")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetScrollActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_scroll_active")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsScrollActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_scroll_follow")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetScrollFollow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_scroll_following")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsScrollFollowing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_v_scroll_bar")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetVScrollBar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2630340773))
	}
	{
		methodName := StringNameFromStr("scroll_to_line")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnScrollToLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("scroll_to_paragraph")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnScrollToParagraph = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("scroll_to_selection")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnScrollToSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_tab_size")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetTabSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_tab_size")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetTabSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_fit_content")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetFitContent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_fit_content_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsFitContentEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_selection_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetSelectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_selection_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsSelectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_context_menu_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetContextMenuEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_context_menu_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsContextMenuEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_shortcut_keys_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetShortcutKeysEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_shortcut_keys_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsShortcutKeysEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_deselect_on_focus_loss_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetDeselectOnFocusLossEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_deselect_on_focus_loss_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsDeselectOnFocusLossEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_drag_and_drop_selection_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetDragAndDropSelectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_drag_and_drop_selection_enabled")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsDragAndDropSelectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_selection_from")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetSelectionFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_selection_to")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetSelectionTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("select_all")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSelectAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_selected_text")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetSelectedText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("deselect")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnDeselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("parse_bbcode")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnParseBbcode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("append_text")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnAppendText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_text")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("is_ready")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsReady = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_threaded")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetThreaded = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_threaded")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsThreaded = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_progress_bar_delay")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetProgressBarDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_progress_bar_delay")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetProgressBarDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_visible_characters")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetVisibleCharacters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_visible_characters")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetVisibleCharacters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_visible_characters_behavior")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetVisibleCharactersBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 258789322))
	}
	{
		methodName := StringNameFromStr("set_visible_characters_behavior")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetVisibleCharactersBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3383839701))
	}
	{
		methodName := StringNameFromStr("set_visible_ratio")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetVisibleRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visible_ratio")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetVisibleRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_character_line")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetCharacterLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
	}
	{
		methodName := StringNameFromStr("get_character_paragraph")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetCharacterParagraph = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
	}
	{
		methodName := StringNameFromStr("get_total_character_count")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetTotalCharacterCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_use_bbcode")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetUseBbcode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_bbcode")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsUsingBbcode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_line_count")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetLineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_visible_line_count")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetVisibleLineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_paragraph_count")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetParagraphCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_visible_paragraph_count")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetVisibleParagraphCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_content_height")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetContentHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_content_width")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetContentWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_line_offset")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetLineOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4025615559))
	}
	{
		methodName := StringNameFromStr("get_paragraph_offset")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetParagraphOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4025615559))
	}
	{
		methodName := StringNameFromStr("parse_expressions_for_values")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnParseExpressionsForValues = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1522900837))
	}
	{
		methodName := StringNameFromStr("set_effects")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnSetEffects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_effects")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetEffects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
	}
	{
		methodName := StringNameFromStr("install_effect")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnInstallEffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1114965689))
	}
	{
		methodName := StringNameFromStr("get_menu")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnGetMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 229722558))
	}
	{
		methodName := StringNameFromStr("is_menu_visible")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnIsMenuVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("menu_option")
		defer methodName.Destroy()
		ptrsForRichTextLabel.fnMenuOption = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
}

type RichTextLabel struct {
	Control
}

func (me *RichTextLabel) BaseClass() string {
	return "RichTextLabel"
}

func NewRichTextLabel() *RichTextLabel {
	str := StringNameFromStr("RichTextLabel") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RichTextLabel{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type RichTextLabelListType int

const (
	RichTextLabelListTypeListNumbers RichTextLabelListType = 0
	RichTextLabelListTypeListLetters RichTextLabelListType = 1
	RichTextLabelListTypeListRoman   RichTextLabelListType = 2
	RichTextLabelListTypeListDots    RichTextLabelListType = 3
)

type RichTextLabelMenuItems int

const (
	RichTextLabelMenuItemsMenuCopy      RichTextLabelMenuItems = 0
	RichTextLabelMenuItemsMenuSelectAll RichTextLabelMenuItems = 1
	RichTextLabelMenuItemsMenuMax       RichTextLabelMenuItems = 2
)

type RichTextLabelImageUpdateMask int

const (
	RichTextLabelImageUpdateMaskUpdateTexture        RichTextLabelImageUpdateMask = 1
	RichTextLabelImageUpdateMaskUpdateSize           RichTextLabelImageUpdateMask = 2
	RichTextLabelImageUpdateMaskUpdateColor          RichTextLabelImageUpdateMask = 4
	RichTextLabelImageUpdateMaskUpdateAlignment      RichTextLabelImageUpdateMask = 8
	RichTextLabelImageUpdateMaskUpdateRegion         RichTextLabelImageUpdateMask = 16
	RichTextLabelImageUpdateMaskUpdatePad            RichTextLabelImageUpdateMask = 32
	RichTextLabelImageUpdateMaskUpdateTooltip        RichTextLabelImageUpdateMask = 64
	RichTextLabelImageUpdateMaskUpdateWidthInPercent RichTextLabelImageUpdateMask = 128
)

func (me *RichTextLabel) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RichTextLabel) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RichTextLabel) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RichTextLabel) GetParsedText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetParsedText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RichTextLabel) AddText(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnAddText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) SetText(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) AddImage(image Texture2D, width int64, height int64, color Color, inline_align InlineAlignment, region Rect2, key Variant, pad bool, tooltip String, size_in_percent bool) {
	cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), color.AsCTypePtr(), gdc.ConstTypePtr(&inline_align), region.AsCTypePtr(), key.AsCTypePtr(), gdc.ConstTypePtr(&pad), tooltip.AsCTypePtr(), gdc.ConstTypePtr(&size_in_percent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnAddImage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) UpdateImage(key Variant, mask RichTextLabelImageUpdateMask, image Texture2D, width int64, height int64, color Color, inline_align InlineAlignment, region Rect2, pad bool, tooltip String, size_in_percent bool) {
	cargs := []gdc.ConstTypePtr{key.AsCTypePtr(), gdc.ConstTypePtr(&mask), image.AsCTypePtr(), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), color.AsCTypePtr(), gdc.ConstTypePtr(&inline_align), region.AsCTypePtr(), gdc.ConstTypePtr(&pad), tooltip.AsCTypePtr(), gdc.ConstTypePtr(&size_in_percent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnUpdateImage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) Newline() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnNewline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) RemoveParagraph(paragraph int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paragraph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&paragraph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnRemoveParagraph), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) PushFont(font Font, font_size int64) {
	cargs := []gdc.ConstTypePtr{font.AsCTypePtr(), gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushFont), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushFontSize(font_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushFontSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushNormal() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushNormal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushBold() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushBold), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushBoldItalics() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushBoldItalics), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushItalics() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushItalics), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushMono() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushMono), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushOutlineSize(outline_size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&outline_size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushOutlineSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushOutlineColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushOutlineColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushParagraph(alignment HorizontalAlignment, base_direction ControlTextDirection, language String, st_parser TextServerStructuredTextParser, justification_flags TextServerJustificationFlag, tab_stops PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&base_direction), language.AsCTypePtr(), gdc.ConstTypePtr(&st_parser), gdc.ConstTypePtr(&justification_flags), tab_stops.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushParagraph), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushIndent(level int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&level)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushIndent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushList(level int64, type_ RichTextLabelListType, capitalize bool, bullet String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&level), gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&capitalize), bullet.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushList), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushMeta(data Variant) {
	cargs := []gdc.ConstTypePtr{data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushMeta), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushHint(description String) {
	cargs := []gdc.ConstTypePtr{description.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushHint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushLanguage(language String) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushUnderline() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushUnderline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushStrikethrough() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushStrikethrough), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushTable(columns int64, inline_align InlineAlignment, align_to_row int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&columns), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&align_to_row)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushTable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushDropcap(string_ String, font Font, size int64, dropcap_margins Rect2, color Color, outline_size int64, outline_color Color) {
	cargs := []gdc.ConstTypePtr{string_.AsCTypePtr(), font.AsCTypePtr(), gdc.ConstTypePtr(&size), dropcap_margins.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&outline_size), outline_color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushDropcap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) SetTableColumnExpand(column int64, expand bool, ratio int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&expand), gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetTableColumnExpand), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) SetCellRowBackgroundColor(odd_row_bg Color, even_row_bg Color) {
	cargs := []gdc.ConstTypePtr{odd_row_bg.AsCTypePtr(), even_row_bg.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetCellRowBackgroundColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) SetCellBorderColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetCellBorderColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) SetCellSizeOverride(min_size Vector2, max_size Vector2) {
	cargs := []gdc.ConstTypePtr{min_size.AsCTypePtr(), max_size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetCellSizeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) SetCellPadding(padding Rect2) {
	cargs := []gdc.ConstTypePtr{padding.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetCellPadding), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushCell() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushCell), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushFgcolor(fgcolor Color) {
	cargs := []gdc.ConstTypePtr{fgcolor.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushFgcolor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushBgcolor(bgcolor Color) {
	cargs := []gdc.ConstTypePtr{bgcolor.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushBgcolor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushCustomfx(effect RichTextEffect, env Dictionary) {
	cargs := []gdc.ConstTypePtr{effect.AsCTypePtr(), env.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushCustomfx), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PushContext() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPushContext), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PopContext() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPopContext), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) Pop() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) PopAll() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnPopAll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerStructuredTextParser

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RichTextLabel) SetStructuredTextBidiOverrideOptions(args Array) {
	cargs := []gdc.ConstTypePtr{args.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetStructuredTextBidiOverrideOptions() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RichTextLabel) SetTextDirection(direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetTextDirection() ControlTextDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RichTextLabel) SetLanguage(language String) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetLanguage() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RichTextLabel) SetAutowrapMode(autowrap_mode TextServerAutowrapMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetAutowrapMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetAutowrapMode() TextServerAutowrapMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerAutowrapMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetAutowrapMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RichTextLabel) SetMetaUnderline(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetMetaUnderline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsMetaUnderlined() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsMetaUnderlined), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetHintUnderline(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetHintUnderline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsHintUnderlined() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsHintUnderlined), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetScrollActive(active bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetScrollActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsScrollActive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsScrollActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetScrollFollow(follow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&follow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetScrollFollow), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsScrollFollowing() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsScrollFollowing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetVScrollBar() VScrollBar {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVScrollBar()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetVScrollBar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RichTextLabel) ScrollToLine(line int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnScrollToLine), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) ScrollToParagraph(paragraph int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paragraph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnScrollToParagraph), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) ScrollToSelection() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnScrollToSelection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) SetTabSize(spaces int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spaces)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetTabSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetTabSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetTabSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetFitContent(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetFitContent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsFitContentEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsFitContentEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetSelectionEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetSelectionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsSelectionEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsSelectionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetContextMenuEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetContextMenuEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsContextMenuEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsContextMenuEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetShortcutKeysEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetShortcutKeysEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsShortcutKeysEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsShortcutKeysEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetDeselectOnFocusLossEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetDeselectOnFocusLossEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsDeselectOnFocusLossEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsDeselectOnFocusLossEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetDragAndDropSelectionEnabled(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetDragAndDropSelectionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsDragAndDropSelectionEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsDragAndDropSelectionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetSelectionFrom() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetSelectionFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetSelectionTo() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetSelectionTo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SelectAll() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSelectAll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetSelectedText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetSelectedText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RichTextLabel) Deselect() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnDeselect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) ParseBbcode(bbcode String) {
	cargs := []gdc.ConstTypePtr{bbcode.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnParseBbcode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) AppendText(bbcode String) {
	cargs := []gdc.ConstTypePtr{bbcode.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnAppendText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RichTextLabel) IsReady() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsReady), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetThreaded(threaded bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threaded)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetThreaded), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsThreaded() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsThreaded), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetProgressBarDelay(delay_ms int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay_ms)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetProgressBarDelay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetProgressBarDelay() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetProgressBarDelay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetVisibleCharacters(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetVisibleCharacters), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetVisibleCharacters() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetVisibleCharacters), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetVisibleCharactersBehavior() TextServerVisibleCharactersBehavior {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerVisibleCharactersBehavior

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetVisibleCharactersBehavior), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RichTextLabel) SetVisibleCharactersBehavior(behavior TextServerVisibleCharactersBehavior) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&behavior)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetVisibleCharactersBehavior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) SetVisibleRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetVisibleRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetVisibleRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetVisibleRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetCharacterLine(character int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&character)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&character)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetCharacterLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetCharacterParagraph(character int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&character)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&character)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetCharacterParagraph), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetTotalCharacterCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetTotalCharacterCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) SetUseBbcode(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetUseBbcode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) IsUsingBbcode() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsUsingBbcode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetLineCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetLineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetVisibleLineCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetVisibleLineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetParagraphCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetParagraphCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetVisibleParagraphCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetVisibleParagraphCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetContentHeight() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetContentHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetContentWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetContentWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetLineOffset(line int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetLineOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) GetParagraphOffset(paragraph int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paragraph)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&paragraph)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetParagraphOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) ParseExpressionsForValues(expressions PackedStringArray) Dictionary {
	cargs := []gdc.ConstTypePtr{expressions.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnParseExpressionsForValues), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RichTextLabel) SetEffects(effects Array) {
	cargs := []gdc.ConstTypePtr{effects.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnSetEffects), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetEffects() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetEffects), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RichTextLabel) InstallEffect(effect Variant) {
	cargs := []gdc.ConstTypePtr{effect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnInstallEffect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RichTextLabel) GetMenu() PopupMenu {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPopupMenu()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnGetMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RichTextLabel) IsMenuVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnIsMenuVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RichTextLabel) MenuOption(option int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRichTextLabel.fnMenuOption), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type RichTextLabelMetaClickedSignalFn func(meta Variant)

func (me *RichTextLabel) ConnectMetaClicked(subs SignalSubscribers, fn RichTextLabelMetaClickedSignalFn) {
	sig := StringNameFromStr("meta_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *RichTextLabel) DisconnectMetaClicked(subs SignalSubscribers, fn RichTextLabelMetaClickedSignalFn) {
	sig := StringNameFromStr("meta_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type RichTextLabelMetaHoverStartedSignalFn func(meta Variant)

func (me *RichTextLabel) ConnectMetaHoverStarted(subs SignalSubscribers, fn RichTextLabelMetaHoverStartedSignalFn) {
	sig := StringNameFromStr("meta_hover_started")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *RichTextLabel) DisconnectMetaHoverStarted(subs SignalSubscribers, fn RichTextLabelMetaHoverStartedSignalFn) {
	sig := StringNameFromStr("meta_hover_started")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type RichTextLabelMetaHoverEndedSignalFn func(meta Variant)

func (me *RichTextLabel) ConnectMetaHoverEnded(subs SignalSubscribers, fn RichTextLabelMetaHoverEndedSignalFn) {
	sig := StringNameFromStr("meta_hover_ended")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *RichTextLabel) DisconnectMetaHoverEnded(subs SignalSubscribers, fn RichTextLabelMetaHoverEndedSignalFn) {
	sig := StringNameFromStr("meta_hover_ended")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type RichTextLabelFinishedSignalFn func()

func (me *RichTextLabel) ConnectFinished(subs SignalSubscribers, fn RichTextLabelFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *RichTextLabel) DisconnectFinished(subs SignalSubscribers, fn RichTextLabelFinishedSignalFn) {
	sig := StringNameFromStr("finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
