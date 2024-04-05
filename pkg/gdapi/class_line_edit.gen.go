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

type ptrsForLineEditList struct {
  fnSetHorizontalAlignment gdc.MethodBindPtr
  fnGetHorizontalAlignment gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
  fnSelect gdc.MethodBindPtr
  fnSelectAll gdc.MethodBindPtr
  fnDeselect gdc.MethodBindPtr
  fnHasSelection gdc.MethodBindPtr
  fnGetSelectedText gdc.MethodBindPtr
  fnGetSelectionFromColumn gdc.MethodBindPtr
  fnGetSelectionToColumn gdc.MethodBindPtr
  fnSetText gdc.MethodBindPtr
  fnGetText gdc.MethodBindPtr
  fnGetDrawControlChars gdc.MethodBindPtr
  fnSetDrawControlChars gdc.MethodBindPtr
  fnSetTextDirection gdc.MethodBindPtr
  fnGetTextDirection gdc.MethodBindPtr
  fnSetLanguage gdc.MethodBindPtr
  fnGetLanguage gdc.MethodBindPtr
  fnSetStructuredTextBidiOverride gdc.MethodBindPtr
  fnGetStructuredTextBidiOverride gdc.MethodBindPtr
  fnSetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
  fnGetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
  fnSetPlaceholder gdc.MethodBindPtr
  fnGetPlaceholder gdc.MethodBindPtr
  fnSetCaretColumn gdc.MethodBindPtr
  fnGetCaretColumn gdc.MethodBindPtr
  fnGetScrollOffset gdc.MethodBindPtr
  fnSetExpandToTextLengthEnabled gdc.MethodBindPtr
  fnIsExpandToTextLengthEnabled gdc.MethodBindPtr
  fnSetCaretBlinkEnabled gdc.MethodBindPtr
  fnIsCaretBlinkEnabled gdc.MethodBindPtr
  fnSetCaretMidGraphemeEnabled gdc.MethodBindPtr
  fnIsCaretMidGraphemeEnabled gdc.MethodBindPtr
  fnSetCaretForceDisplayed gdc.MethodBindPtr
  fnIsCaretForceDisplayed gdc.MethodBindPtr
  fnSetCaretBlinkInterval gdc.MethodBindPtr
  fnGetCaretBlinkInterval gdc.MethodBindPtr
  fnSetMaxLength gdc.MethodBindPtr
  fnGetMaxLength gdc.MethodBindPtr
  fnInsertTextAtCaret gdc.MethodBindPtr
  fnDeleteCharAtCaret gdc.MethodBindPtr
  fnDeleteText gdc.MethodBindPtr
  fnSetEditable gdc.MethodBindPtr
  fnIsEditable gdc.MethodBindPtr
  fnSetSecret gdc.MethodBindPtr
  fnIsSecret gdc.MethodBindPtr
  fnSetSecretCharacter gdc.MethodBindPtr
  fnGetSecretCharacter gdc.MethodBindPtr
  fnMenuOption gdc.MethodBindPtr
  fnGetMenu gdc.MethodBindPtr
  fnIsMenuVisible gdc.MethodBindPtr
  fnSetContextMenuEnabled gdc.MethodBindPtr
  fnIsContextMenuEnabled gdc.MethodBindPtr
  fnSetVirtualKeyboardEnabled gdc.MethodBindPtr
  fnIsVirtualKeyboardEnabled gdc.MethodBindPtr
  fnSetVirtualKeyboardType gdc.MethodBindPtr
  fnGetVirtualKeyboardType gdc.MethodBindPtr
  fnSetClearButtonEnabled gdc.MethodBindPtr
  fnIsClearButtonEnabled gdc.MethodBindPtr
  fnSetShortcutKeysEnabled gdc.MethodBindPtr
  fnIsShortcutKeysEnabled gdc.MethodBindPtr
  fnSetMiddleMousePasteEnabled gdc.MethodBindPtr
  fnIsMiddleMousePasteEnabled gdc.MethodBindPtr
  fnSetSelectingEnabled gdc.MethodBindPtr
  fnIsSelectingEnabled gdc.MethodBindPtr
  fnSetDeselectOnFocusLossEnabled gdc.MethodBindPtr
  fnIsDeselectOnFocusLossEnabled gdc.MethodBindPtr
  fnSetDragAndDropSelectionEnabled gdc.MethodBindPtr
  fnIsDragAndDropSelectionEnabled gdc.MethodBindPtr
  fnSetRightIcon gdc.MethodBindPtr
  fnGetRightIcon gdc.MethodBindPtr
  fnSetFlat gdc.MethodBindPtr
  fnIsFlat gdc.MethodBindPtr
  fnSetSelectAllOnFocus gdc.MethodBindPtr
  fnIsSelectAllOnFocus gdc.MethodBindPtr
}

var ptrsForLineEdit ptrsForLineEditList

func initLineEditPtrs(iface gdc.Interface) {

  className := StringNameFromStr("LineEdit")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_horizontal_alignment")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312603777))
  }
  {
    methodName := StringNameFromStr("get_horizontal_alignment")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 341400642))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForLineEdit.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("select")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSelect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1328111411))
  }
  {
    methodName := StringNameFromStr("select_all")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSelectAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("deselect")
    defer methodName.Destroy()
    ptrsForLineEdit.fnDeselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("has_selection")
    defer methodName.Destroy()
    ptrsForLineEdit.fnHasSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_selected_text")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetSelectedText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
  }
  {
    methodName := StringNameFromStr("get_selection_from_column")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetSelectionFromColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_selection_to_column")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetSelectionToColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_text")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_text")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_draw_control_chars")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetDrawControlChars = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_draw_control_chars")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetDrawControlChars = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_text_direction")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 119160795))
  }
  {
    methodName := StringNameFromStr("get_text_direction")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797257663))
  }
  {
    methodName := StringNameFromStr("set_language")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_language")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_structured_text_bidi_override")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 55961453))
  }
  {
    methodName := StringNameFromStr("get_structured_text_bidi_override")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3385126229))
  }
  {
    methodName := StringNameFromStr("set_structured_text_bidi_override_options")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_structured_text_bidi_override_options")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("set_placeholder")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetPlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_placeholder")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetPlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_caret_column")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetCaretColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_caret_column")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetCaretColumn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_scroll_offset")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetScrollOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_expand_to_text_length_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetExpandToTextLengthEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_expand_to_text_length_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsExpandToTextLengthEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_caret_blink_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetCaretBlinkEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_caret_blink_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsCaretBlinkEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_caret_mid_grapheme_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetCaretMidGraphemeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_caret_mid_grapheme_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsCaretMidGraphemeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_caret_force_displayed")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetCaretForceDisplayed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_caret_force_displayed")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsCaretForceDisplayed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_caret_blink_interval")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetCaretBlinkInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_caret_blink_interval")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetCaretBlinkInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max_length")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetMaxLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_length")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetMaxLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("insert_text_at_caret")
    defer methodName.Destroy()
    ptrsForLineEdit.fnInsertTextAtCaret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("delete_char_at_caret")
    defer methodName.Destroy()
    ptrsForLineEdit.fnDeleteCharAtCaret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("delete_text")
    defer methodName.Destroy()
    ptrsForLineEdit.fnDeleteText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("set_editable")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_editable")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsEditable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_secret")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetSecret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_secret")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsSecret = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_secret_character")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetSecretCharacter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_secret_character")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetSecretCharacter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("menu_option")
    defer methodName.Destroy()
    ptrsForLineEdit.fnMenuOption = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_menu")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 229722558))
  }
  {
    methodName := StringNameFromStr("is_menu_visible")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsMenuVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_context_menu_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetContextMenuEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_context_menu_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsContextMenuEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_virtual_keyboard_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetVirtualKeyboardEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_virtual_keyboard_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsVirtualKeyboardEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_virtual_keyboard_type")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetVirtualKeyboardType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2696893573))
  }
  {
    methodName := StringNameFromStr("get_virtual_keyboard_type")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetVirtualKeyboardType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1928699316))
  }
  {
    methodName := StringNameFromStr("set_clear_button_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetClearButtonEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_clear_button_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsClearButtonEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_shortcut_keys_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetShortcutKeysEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_shortcut_keys_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsShortcutKeysEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_middle_mouse_paste_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetMiddleMousePasteEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_middle_mouse_paste_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsMiddleMousePasteEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_selecting_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetSelectingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_selecting_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsSelectingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_deselect_on_focus_loss_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetDeselectOnFocusLossEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_deselect_on_focus_loss_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsDeselectOnFocusLossEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_drag_and_drop_selection_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetDragAndDropSelectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_drag_and_drop_selection_enabled")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsDragAndDropSelectionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_right_icon")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetRightIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
  }
  {
    methodName := StringNameFromStr("get_right_icon")
    defer methodName.Destroy()
    ptrsForLineEdit.fnGetRightIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 255860311))
  }
  {
    methodName := StringNameFromStr("set_flat")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetFlat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_flat")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsFlat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_select_all_on_focus")
    defer methodName.Destroy()
    ptrsForLineEdit.fnSetSelectAllOnFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_select_all_on_focus")
    defer methodName.Destroy()
    ptrsForLineEdit.fnIsSelectAllOnFocus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type LineEdit struct {
  Control
}

func (me *LineEdit) BaseClass() string {
  return "LineEdit"
}

func NewLineEdit() *LineEdit {
  str := StringNameFromStr("LineEdit") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &LineEdit{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type LineEditMenuItems int
const (
  LineEditMenuItemsMenuCut LineEditMenuItems = 0
  LineEditMenuItemsMenuCopy LineEditMenuItems = 1
  LineEditMenuItemsMenuPaste LineEditMenuItems = 2
  LineEditMenuItemsMenuClear LineEditMenuItems = 3
  LineEditMenuItemsMenuSelectAll LineEditMenuItems = 4
  LineEditMenuItemsMenuUndo LineEditMenuItems = 5
  LineEditMenuItemsMenuRedo LineEditMenuItems = 6
  LineEditMenuItemsMenuSubmenuTextDir LineEditMenuItems = 7
  LineEditMenuItemsMenuDirInherited LineEditMenuItems = 8
  LineEditMenuItemsMenuDirAuto LineEditMenuItems = 9
  LineEditMenuItemsMenuDirLtr LineEditMenuItems = 10
  LineEditMenuItemsMenuDirRtl LineEditMenuItems = 11
  LineEditMenuItemsMenuDisplayUcc LineEditMenuItems = 12
  LineEditMenuItemsMenuSubmenuInsertUcc LineEditMenuItems = 13
  LineEditMenuItemsMenuInsertLrm LineEditMenuItems = 14
  LineEditMenuItemsMenuInsertRlm LineEditMenuItems = 15
  LineEditMenuItemsMenuInsertLre LineEditMenuItems = 16
  LineEditMenuItemsMenuInsertRle LineEditMenuItems = 17
  LineEditMenuItemsMenuInsertLro LineEditMenuItems = 18
  LineEditMenuItemsMenuInsertRlo LineEditMenuItems = 19
  LineEditMenuItemsMenuInsertPdf LineEditMenuItems = 20
  LineEditMenuItemsMenuInsertAlm LineEditMenuItems = 21
  LineEditMenuItemsMenuInsertLri LineEditMenuItems = 22
  LineEditMenuItemsMenuInsertRli LineEditMenuItems = 23
  LineEditMenuItemsMenuInsertFsi LineEditMenuItems = 24
  LineEditMenuItemsMenuInsertPdi LineEditMenuItems = 25
  LineEditMenuItemsMenuInsertZwj LineEditMenuItems = 26
  LineEditMenuItemsMenuInsertZwnj LineEditMenuItems = 27
  LineEditMenuItemsMenuInsertWj LineEditMenuItems = 28
  LineEditMenuItemsMenuInsertShy LineEditMenuItems = 29
  LineEditMenuItemsMenuMax LineEditMenuItems = 30
)

type LineEditVirtualKeyboardType int
const (
  LineEditVirtualKeyboardTypeKeyboardTypeDefault LineEditVirtualKeyboardType = 0
  LineEditVirtualKeyboardTypeKeyboardTypeMultiline LineEditVirtualKeyboardType = 1
  LineEditVirtualKeyboardTypeKeyboardTypeNumber LineEditVirtualKeyboardType = 2
  LineEditVirtualKeyboardTypeKeyboardTypeNumberDecimal LineEditVirtualKeyboardType = 3
  LineEditVirtualKeyboardTypeKeyboardTypePhone LineEditVirtualKeyboardType = 4
  LineEditVirtualKeyboardTypeKeyboardTypeEmailAddress LineEditVirtualKeyboardType = 5
  LineEditVirtualKeyboardTypeKeyboardTypePassword LineEditVirtualKeyboardType = 6
  LineEditVirtualKeyboardTypeKeyboardTypeUrl LineEditVirtualKeyboardType = 7
)

func (me *LineEdit) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *LineEdit) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LineEdit) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *LineEdit) SetHorizontalAlignment(alignment HorizontalAlignment, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetHorizontalAlignment() HorizontalAlignment {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret HorizontalAlignment

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LineEdit) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) Select(from int64, to int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from) , gdc.ConstTypePtr(&to) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSelect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) SelectAll()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSelectAll), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) Deselect()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnDeselect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) HasSelection() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnHasSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) GetSelectedText() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetSelectedText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LineEdit) GetSelectionFromColumn() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetSelectionFromColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) GetSelectionToColumn() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetSelectionToColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetText(text String, )  {
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetText() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LineEdit) GetDrawControlChars() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetDrawControlChars), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetDrawControlChars(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetDrawControlChars), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) SetTextDirection(direction ControlTextDirection, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetTextDirection() ControlTextDirection {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ControlTextDirection

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LineEdit) SetLanguage(language String, )  {
  cargs := []gdc.ConstTypePtr{language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetLanguage() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LineEdit) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerStructuredTextParser

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LineEdit) SetStructuredTextBidiOverrideOptions(args Array, )  {
  cargs := []gdc.ConstTypePtr{args.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetStructuredTextBidiOverrideOptions() Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LineEdit) SetPlaceholder(text String, )  {
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetPlaceholder), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetPlaceholder() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetPlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LineEdit) SetCaretColumn(position int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetCaretColumn), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetCaretColumn() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetCaretColumn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) GetScrollOffset() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetScrollOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetExpandToTextLengthEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetExpandToTextLengthEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsExpandToTextLengthEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsExpandToTextLengthEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetCaretBlinkEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetCaretBlinkEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsCaretBlinkEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsCaretBlinkEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetCaretMidGraphemeEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetCaretMidGraphemeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsCaretMidGraphemeEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsCaretMidGraphemeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetCaretForceDisplayed(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetCaretForceDisplayed), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsCaretForceDisplayed() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsCaretForceDisplayed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetCaretBlinkInterval(interval float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interval) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetCaretBlinkInterval), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetCaretBlinkInterval() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetCaretBlinkInterval), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetMaxLength(chars int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&chars) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetMaxLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetMaxLength() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetMaxLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) InsertTextAtCaret(text String, )  {
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnInsertTextAtCaret), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) DeleteCharAtCaret()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnDeleteCharAtCaret), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) DeleteText(from_column int64, to_column int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_column) , gdc.ConstTypePtr(&to_column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnDeleteText), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) SetEditable(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetEditable), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsEditable() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsEditable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetSecret(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetSecret), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsSecret() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsSecret), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetSecretCharacter(character String, )  {
  cargs := []gdc.ConstTypePtr{character.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetSecretCharacter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetSecretCharacter() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetSecretCharacter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LineEdit) MenuOption(option int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnMenuOption), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetMenu() PopupMenu {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPopupMenu()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetMenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LineEdit) IsMenuVisible() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsMenuVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetContextMenuEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetContextMenuEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsContextMenuEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsContextMenuEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetVirtualKeyboardEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetVirtualKeyboardEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsVirtualKeyboardEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsVirtualKeyboardEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetVirtualKeyboardType(type_ LineEditVirtualKeyboardType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetVirtualKeyboardType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetVirtualKeyboardType() LineEditVirtualKeyboardType {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret LineEditVirtualKeyboardType

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetVirtualKeyboardType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *LineEdit) SetClearButtonEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetClearButtonEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsClearButtonEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsClearButtonEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetShortcutKeysEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetShortcutKeysEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsShortcutKeysEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsShortcutKeysEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetMiddleMousePasteEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetMiddleMousePasteEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsMiddleMousePasteEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsMiddleMousePasteEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetSelectingEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetSelectingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsSelectingEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsSelectingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetDeselectOnFocusLossEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetDeselectOnFocusLossEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsDeselectOnFocusLossEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsDeselectOnFocusLossEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetDragAndDropSelectionEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetDragAndDropSelectionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsDragAndDropSelectionEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsDragAndDropSelectionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetRightIcon(icon Texture2D, )  {
  cargs := []gdc.ConstTypePtr{icon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetRightIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) GetRightIcon() Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnGetRightIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *LineEdit) SetFlat(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetFlat), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsFlat() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsFlat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *LineEdit) SetSelectAllOnFocus(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnSetSelectAllOnFocus), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *LineEdit) IsSelectAllOnFocus() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLineEdit.fnIsSelectAllOnFocus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type LineEditTextChangedSignalFn func(new_text String, )

func (me *LineEdit) ConnectTextChanged(subs SignalSubscribers, fn LineEditTextChangedSignalFn) {
  sig := StringNameFromStr("text_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *LineEdit) DisconnectTextChanged(subs SignalSubscribers, fn LineEditTextChangedSignalFn) {
  sig := StringNameFromStr("text_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type LineEditTextChangeRejectedSignalFn func(rejected_substring String, )

func (me *LineEdit) ConnectTextChangeRejected(subs SignalSubscribers, fn LineEditTextChangeRejectedSignalFn) {
  sig := StringNameFromStr("text_change_rejected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *LineEdit) DisconnectTextChangeRejected(subs SignalSubscribers, fn LineEditTextChangeRejectedSignalFn) {
  sig := StringNameFromStr("text_change_rejected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type LineEditTextSubmittedSignalFn func(new_text String, )

func (me *LineEdit) ConnectTextSubmitted(subs SignalSubscribers, fn LineEditTextSubmittedSignalFn) {
  sig := StringNameFromStr("text_submitted")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *LineEdit) DisconnectTextSubmitted(subs SignalSubscribers, fn LineEditTextSubmittedSignalFn) {
  sig := StringNameFromStr("text_submitted")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
