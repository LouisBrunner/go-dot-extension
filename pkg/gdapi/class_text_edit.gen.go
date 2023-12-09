// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextEdit struct {
  obj gdc.ObjectPtr
}

func (me *TextEdit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextEdit) BaseClass() string {
  return "TextEdit"
}



// Enums

type TextEditMenuItems int
const (
  TextEditMenuItemsMenuCut TextEditMenuItems = 0
  TextEditMenuItemsMenuCopy TextEditMenuItems = 1
  TextEditMenuItemsMenuPaste TextEditMenuItems = 2
  TextEditMenuItemsMenuClear TextEditMenuItems = 3
  TextEditMenuItemsMenuSelectAll TextEditMenuItems = 4
  TextEditMenuItemsMenuUndo TextEditMenuItems = 5
  TextEditMenuItemsMenuRedo TextEditMenuItems = 6
  TextEditMenuItemsMenuSubmenuTextDir TextEditMenuItems = 7
  TextEditMenuItemsMenuDirInherited TextEditMenuItems = 8
  TextEditMenuItemsMenuDirAuto TextEditMenuItems = 9
  TextEditMenuItemsMenuDirLtr TextEditMenuItems = 10
  TextEditMenuItemsMenuDirRtl TextEditMenuItems = 11
  TextEditMenuItemsMenuDisplayUcc TextEditMenuItems = 12
  TextEditMenuItemsMenuSubmenuInsertUcc TextEditMenuItems = 13
  TextEditMenuItemsMenuInsertLrm TextEditMenuItems = 14
  TextEditMenuItemsMenuInsertRlm TextEditMenuItems = 15
  TextEditMenuItemsMenuInsertLre TextEditMenuItems = 16
  TextEditMenuItemsMenuInsertRle TextEditMenuItems = 17
  TextEditMenuItemsMenuInsertLro TextEditMenuItems = 18
  TextEditMenuItemsMenuInsertRlo TextEditMenuItems = 19
  TextEditMenuItemsMenuInsertPdf TextEditMenuItems = 20
  TextEditMenuItemsMenuInsertAlm TextEditMenuItems = 21
  TextEditMenuItemsMenuInsertLri TextEditMenuItems = 22
  TextEditMenuItemsMenuInsertRli TextEditMenuItems = 23
  TextEditMenuItemsMenuInsertFsi TextEditMenuItems = 24
  TextEditMenuItemsMenuInsertPdi TextEditMenuItems = 25
  TextEditMenuItemsMenuInsertZwj TextEditMenuItems = 26
  TextEditMenuItemsMenuInsertZwnj TextEditMenuItems = 27
  TextEditMenuItemsMenuInsertWj TextEditMenuItems = 28
  TextEditMenuItemsMenuInsertShy TextEditMenuItems = 29
  TextEditMenuItemsMenuMax TextEditMenuItems = 30
)

type TextEditEditAction int
const (
  TextEditEditActionActionNone TextEditEditAction = 0
  TextEditEditActionActionTyping TextEditEditAction = 1
  TextEditEditActionActionBackspace TextEditEditAction = 2
  TextEditEditActionActionDelete TextEditEditAction = 3
)

type TextEditSearchFlags int
const (
  TextEditSearchFlagsSearchMatchCase TextEditSearchFlags = 1
  TextEditSearchFlagsSearchWholeWords TextEditSearchFlags = 2
  TextEditSearchFlagsSearchBackwards TextEditSearchFlags = 4
)

type TextEditCaretType int
const (
  TextEditCaretTypeCaretTypeLine TextEditCaretType = 0
  TextEditCaretTypeCaretTypeBlock TextEditCaretType = 1
)

type TextEditSelectionMode int
const (
  TextEditSelectionModeSelectionModeNone TextEditSelectionMode = 0
  TextEditSelectionModeSelectionModeShift TextEditSelectionMode = 1
  TextEditSelectionModeSelectionModePointer TextEditSelectionMode = 2
  TextEditSelectionModeSelectionModeWord TextEditSelectionMode = 3
  TextEditSelectionModeSelectionModeLine TextEditSelectionMode = 4
)

type TextEditLineWrappingMode int
const (
  TextEditLineWrappingModeLineWrappingNone TextEditLineWrappingMode = 0
  TextEditLineWrappingModeLineWrappingBoundary TextEditLineWrappingMode = 1
)

type TextEditGutterType int
const (
  TextEditGutterTypeGutterTypeString TextEditGutterType = 0
  TextEditGutterTypeGutterTypeIcon TextEditGutterType = 1
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

func  (me *TextEdit) XHandleUnicodeInput(unicode_char int, caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) XBackspace(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) XCut(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) XCopy(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) XPaste(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) XPastePrimaryClipboard(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) HasImeText()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetEditable(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsEditable()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetTextDirection(direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetTextDirection()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLanguage(language String, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLanguage()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetStructuredTextBidiOverride()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetStructuredTextBidiOverrideOptions(args Array, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetStructuredTextBidiOverrideOptions()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetTabSize(size int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetTabSize()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetOvertypeModeEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsOvertypeModeEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetContextMenuEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsContextMenuEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetShortcutKeysEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsShortcutKeysEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetVirtualKeyboardEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsVirtualKeyboardEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetMiddleMousePasteEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsMiddleMousePasteEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) Clear()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetText(text String, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetText()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineCount()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetPlaceholder(text String, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetPlaceholder()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLine(line int, new_text String, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLine(line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineWidth(line int, wrap_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineHeight()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetIndentLevel(line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetFirstNonWhitespaceColumn(line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SwapLines(from_line int, to_line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) InsertLineAt(line int, text String, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) InsertTextAtCaret(text String, caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) RemoveText(from_line int, from_column int, to_line int, to_column int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLastUnhiddenLine()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetNextVisibleLineOffsetFrom(line int, visible_amount int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetNextVisibleLineIndexOffsetFrom(line int, wrap_index int, visible_amount int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) Backspace(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) Cut(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) Copy(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) Paste(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) PastePrimaryClipboard(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) StartAction(action TextEditEditAction, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) EndAction()  {
  panic("TODO: implement")
}

func  (me *TextEdit) BeginComplexOperation()  {
  panic("TODO: implement")
}

func  (me *TextEdit) EndComplexOperation()  {
  panic("TODO: implement")
}

func  (me *TextEdit) HasUndo()  {
  panic("TODO: implement")
}

func  (me *TextEdit) HasRedo()  {
  panic("TODO: implement")
}

func  (me *TextEdit) Undo()  {
  panic("TODO: implement")
}

func  (me *TextEdit) Redo()  {
  panic("TODO: implement")
}

func  (me *TextEdit) ClearUndoHistory()  {
  panic("TODO: implement")
}

func  (me *TextEdit) TagSavedVersion()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetVersion()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSavedVersion()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetSearchText(search_text String, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetSearchFlags(flags int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) Search(text String, flags int, from_line int, from_colum int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetTooltipRequestFunc(callback Callable, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLocalMousePos()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetWordAtPos(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineColumnAtPos(position Vector2i, allow_out_of_bounds bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetPosAtLineColumn(line int, column int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetRectAtLineColumn(line int, column int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetMinimapLineAtPos(position Vector2i, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsDraggingCursor()  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsMouseOverSelection(edges bool, caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetCaretType(type_ TextEditCaretType, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetCaretType()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetCaretBlinkEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsCaretBlinkEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetCaretBlinkInterval(interval float32, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetCaretBlinkInterval()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetDrawCaretWhenEditableDisabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsDrawingCaretWhenEditableDisabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetMoveCaretOnRightClickEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsMoveCaretOnRightClickEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetCaretMidGraphemeEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsCaretMidGraphemeEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetMultipleCaretsEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsMultipleCaretsEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) AddCaret(line int, col int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) RemoveCaret(caret int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) RemoveSecondaryCarets()  {
  panic("TODO: implement")
}

func  (me *TextEdit) MergeOverlappingCarets()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetCaretCount()  {
  panic("TODO: implement")
}

func  (me *TextEdit) AddCaretAtCarets(below bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetCaretIndexEditOrder()  {
  panic("TODO: implement")
}

func  (me *TextEdit) AdjustCaretsAfterEdit(caret int, from_line int, from_col int, to_line int, to_col int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsCaretVisible(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetCaretDrawPos(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetCaretLine(line int, adjust_viewport bool, can_be_hidden bool, wrap_index int, caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetCaretLine(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetCaretColumn(column int, adjust_viewport bool, caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetCaretColumn(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetCaretWrapIndex(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetWordUnderCaret(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetSelectingEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsSelectingEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetDeselectOnFocusLossEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsDeselectOnFocusLossEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetDragAndDropSelectionEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsDragAndDropSelectionEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetSelectionMode(mode TextEditSelectionMode, line int, column int, caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSelectionMode()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SelectAll()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SelectWordUnderCaret(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) AddSelectionForNextOccurrence()  {
  panic("TODO: implement")
}

func  (me *TextEdit) Select(from_line int, from_column int, to_line int, to_column int, caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) HasSelection(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSelectedText(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSelectionLine(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSelectionColumn(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSelectionFromLine(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSelectionFromColumn(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSelectionToLine(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSelectionToColumn(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) Deselect(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) DeleteSelection(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineWrappingMode(mode TextEditLineWrappingMode, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineWrappingMode()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetAutowrapMode()  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsLineWrapped(line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineWrapCount(line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineWrapIndexAtColumn(line int, column int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineWrappedText(line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetSmoothScrollEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsSmoothScrollEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetVScrollBar()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetHScrollBar()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetVScroll(value float32, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetVScroll()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetHScroll(value int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetHScroll()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetScrollPastEndOfFileEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsScrollPastEndOfFileEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetVScrollSpeed(speed float32, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetVScrollSpeed()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetFitContentHeightEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsFitContentHeightEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetScrollPosForLine(line int, wrap_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineAsFirstVisible(line int, wrap_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetFirstVisibleLine()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineAsCenterVisible(line int, wrap_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineAsLastVisible(line int, wrap_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLastFullVisibleLine()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLastFullVisibleLineWrapIndex()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetVisibleLineCount()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetVisibleLineCountInRange(from_line int, to_line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetTotalVisibleLineCount()  {
  panic("TODO: implement")
}

func  (me *TextEdit) AdjustViewportToCaret(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) CenterViewportToCaret(caret_index int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetDrawMinimap(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsDrawingMinimap()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetMinimapWidth(width int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetMinimapWidth()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetMinimapVisibleLines()  {
  panic("TODO: implement")
}

func  (me *TextEdit) AddGutter(at int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) RemoveGutter(gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetGutterCount()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetGutterName(gutter int, name String, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetGutterName(gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetGutterType(gutter int, type_ TextEditGutterType, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetGutterType(gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetGutterWidth(gutter int, width int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetGutterWidth(gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetGutterDraw(gutter int, draw bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsGutterDrawn(gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetGutterClickable(gutter int, clickable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsGutterClickable(gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetGutterOverwritable(gutter int, overwritable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsGutterOverwritable(gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) MergeGutters(from_line int, to_line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetGutterCustomDraw(column int, draw_callback Callable, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetTotalGutterWidth()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineGutterMetadata(line int, gutter int, metadata Variant, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineGutterMetadata(line int, gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineGutterText(line int, gutter int, text String, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineGutterText(line int, gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineGutterIcon(line int, gutter int, icon Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineGutterIcon(line int, gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineGutterItemColor(line int, gutter int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineGutterItemColor(line int, gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineGutterClickable(line int, gutter int, clickable bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsLineGutterClickable(line int, gutter int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetLineBackgroundColor(line int, color Color, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetLineBackgroundColor(line int, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetSyntaxHighlighter(syntax_highlighter SyntaxHighlighter, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetSyntaxHighlighter()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetHighlightCurrentLine(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsHighlightCurrentLineEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetHighlightAllOccurrences(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsHighlightAllOccurrencesEnabled()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetDrawControlChars()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetDrawControlChars(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetDrawTabs(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsDrawingTabs()  {
  panic("TODO: implement")
}

func  (me *TextEdit) SetDrawSpaces(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsDrawingSpaces()  {
  panic("TODO: implement")
}

func  (me *TextEdit) GetMenu()  {
  panic("TODO: implement")
}

func  (me *TextEdit) IsMenuVisible()  {
  panic("TODO: implement")
}

func  (me *TextEdit) MenuOption(option int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
