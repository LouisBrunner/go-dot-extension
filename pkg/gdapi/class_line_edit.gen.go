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

func (me *LineEdit) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LineEdit) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *LineEdit) SetHorizontalAlignment(alignment HorizontalAlignment, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetHorizontalAlignment()  {
  panic("TODO: implement")
}

func  (me *LineEdit) Clear()  {
  panic("TODO: implement")
}

func  (me *LineEdit) Select(from int, to int, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) SelectAll()  {
  panic("TODO: implement")
}

func  (me *LineEdit) Deselect()  {
  panic("TODO: implement")
}

func  (me *LineEdit) HasSelection()  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetSelectedText()  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetSelectionFromColumn()  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetSelectionToColumn()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetText(text String, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetText()  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetDrawControlChars()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetDrawControlChars(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetTextDirection(direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetTextDirection()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetLanguage(language String, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetLanguage()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetStructuredTextBidiOverride()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetStructuredTextBidiOverrideOptions(args Array, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetStructuredTextBidiOverrideOptions()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetPlaceholder(text String, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetPlaceholder()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetCaretColumn(position int, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetCaretColumn()  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetScrollOffset()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetExpandToTextLengthEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsExpandToTextLengthEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetCaretBlinkEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsCaretBlinkEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetCaretMidGraphemeEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsCaretMidGraphemeEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetCaretForceDisplayed(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsCaretForceDisplayed()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetCaretBlinkInterval(interval float32, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetCaretBlinkInterval()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetMaxLength(chars int, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetMaxLength()  {
  panic("TODO: implement")
}

func  (me *LineEdit) InsertTextAtCaret(text String, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) DeleteCharAtCaret()  {
  panic("TODO: implement")
}

func  (me *LineEdit) DeleteText(from_column int, to_column int, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetEditable(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsEditable()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetSecret(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsSecret()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetSecretCharacter(character String, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetSecretCharacter()  {
  panic("TODO: implement")
}

func  (me *LineEdit) MenuOption(option int, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetMenu()  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsMenuVisible()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetContextMenuEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsContextMenuEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetVirtualKeyboardEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsVirtualKeyboardEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetVirtualKeyboardType(type_ LineEditVirtualKeyboardType, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetVirtualKeyboardType()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetClearButtonEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsClearButtonEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetShortcutKeysEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsShortcutKeysEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetMiddleMousePasteEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsMiddleMousePasteEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetSelectingEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsSelectingEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetDeselectOnFocusLossEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsDeselectOnFocusLossEnabled()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetRightIcon(icon Texture2D, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) GetRightIcon()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetFlat(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsFlat()  {
  panic("TODO: implement")
}

func  (me *LineEdit) SetSelectAllOnFocus(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *LineEdit) IsSelectAllOnFocus()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
