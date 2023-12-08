// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LineEdit struct {
  obj gdc.ObjectPtr
}

func (me *LineEdit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LineEdit) BaseClass() string {
  return "LineEdit"
}

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

func  (me *LineEdit) SetHorizontalAlignment(alignment HorizontalAlignment, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetHorizontalAlignment() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) Select(from int, to int, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SelectAll() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) Deselect() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) HasSelection() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetSelectedText() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetSelectionFromColumn() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetSelectionToColumn() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetText(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetText() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetDrawControlChars() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetDrawControlChars(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetTextDirection(direction ControlTextDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetTextDirection() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetLanguage(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetLanguage() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetStructuredTextBidiOverride() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetStructuredTextBidiOverrideOptions(args Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetStructuredTextBidiOverrideOptions() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetPlaceholder(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetPlaceholder() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetCaretColumn(position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetCaretColumn() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetScrollOffset() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetExpandToTextLengthEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsExpandToTextLengthEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetCaretBlinkEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsCaretBlinkEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetCaretMidGraphemeEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsCaretMidGraphemeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetCaretForceDisplayed(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsCaretForceDisplayed() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetCaretBlinkInterval(interval float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetCaretBlinkInterval() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetMaxLength(chars int, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetMaxLength() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) InsertTextAtCaret(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) DeleteCharAtCaret() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) DeleteText(from_column int, to_column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetEditable(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsEditable() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetSecret(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsSecret() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetSecretCharacter(character String, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetSecretCharacter() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) MenuOption(option int, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetMenu() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsMenuVisible() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetContextMenuEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsContextMenuEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetVirtualKeyboardEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsVirtualKeyboardEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetVirtualKeyboardType(type_ LineEditVirtualKeyboardType, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetVirtualKeyboardType() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetClearButtonEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsClearButtonEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetShortcutKeysEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsShortcutKeysEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetMiddleMousePasteEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsMiddleMousePasteEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetSelectingEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsSelectingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetDeselectOnFocusLossEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsDeselectOnFocusLossEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetRightIcon(icon Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) GetRightIcon() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetFlat(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsFlat() { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) SetSelectAllOnFocus(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *LineEdit) IsSelectAllOnFocus() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
