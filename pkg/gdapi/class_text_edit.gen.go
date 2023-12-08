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

func  (me *TextEdit) XHandleUnicodeInput(unicode_char int, caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) XBackspace(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) XCut(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) XCopy(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) XPaste(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) XPastePrimaryClipboard(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) HasImeText() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetEditable(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsEditable() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetTextDirection(direction ControlTextDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetTextDirection() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLanguage(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLanguage() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetStructuredTextBidiOverride() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetStructuredTextBidiOverrideOptions(args Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetStructuredTextBidiOverrideOptions() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetTabSize(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetTabSize() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetOvertypeModeEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsOvertypeModeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetContextMenuEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsContextMenuEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetShortcutKeysEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsShortcutKeysEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetVirtualKeyboardEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsVirtualKeyboardEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetMiddleMousePasteEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsMiddleMousePasteEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetText(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetText() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineCount() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetPlaceholder(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetPlaceholder() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLine(line int, new_text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLine(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineWidth(line int, wrap_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineHeight() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetIndentLevel(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetFirstNonWhitespaceColumn(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SwapLines(from_line int, to_line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) InsertLineAt(line int, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) InsertTextAtCaret(text String, caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) RemoveText(from_line int, from_column int, to_line int, to_column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLastUnhiddenLine() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetNextVisibleLineOffsetFrom(line int, visible_amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetNextVisibleLineIndexOffsetFrom(line int, wrap_index int, visible_amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Backspace(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Cut(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Copy(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Paste(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) PastePrimaryClipboard(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) StartAction(action TextEditEditAction, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) EndAction() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) BeginComplexOperation() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) EndComplexOperation() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) HasUndo() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) HasRedo() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Undo() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Redo() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) ClearUndoHistory() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) TagSavedVersion() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetVersion() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSavedVersion() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetSearchText(search_text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetSearchFlags(flags int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Search(text String, flags int, from_line int, from_colum int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetTooltipRequestFunc(callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLocalMousePos() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetWordAtPos(position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineColumnAtPos(position Vector2i, allow_out_of_bounds bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetPosAtLineColumn(line int, column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetRectAtLineColumn(line int, column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetMinimapLineAtPos(position Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsDraggingCursor() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsMouseOverSelection(edges bool, caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetCaretType(type_ TextEditCaretType, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetCaretType() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetCaretBlinkEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsCaretBlinkEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetCaretBlinkInterval(interval float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetCaretBlinkInterval() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetDrawCaretWhenEditableDisabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsDrawingCaretWhenEditableDisabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetMoveCaretOnRightClickEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsMoveCaretOnRightClickEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetCaretMidGraphemeEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsCaretMidGraphemeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetMultipleCaretsEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsMultipleCaretsEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) AddCaret(line int, col int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) RemoveCaret(caret int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) RemoveSecondaryCarets() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) MergeOverlappingCarets() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetCaretCount() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) AddCaretAtCarets(below bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetCaretIndexEditOrder() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) AdjustCaretsAfterEdit(caret int, from_line int, from_col int, to_line int, to_col int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsCaretVisible(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetCaretDrawPos(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetCaretLine(line int, adjust_viewport bool, can_be_hidden bool, wrap_index int, caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetCaretLine(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetCaretColumn(column int, adjust_viewport bool, caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetCaretColumn(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetCaretWrapIndex(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetWordUnderCaret(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetSelectingEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsSelectingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetDeselectOnFocusLossEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsDeselectOnFocusLossEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetDragAndDropSelectionEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsDragAndDropSelectionEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetSelectionMode(mode TextEditSelectionMode, line int, column int, caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSelectionMode() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SelectAll() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SelectWordUnderCaret(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) AddSelectionForNextOccurrence() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Select(from_line int, from_column int, to_line int, to_column int, caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) HasSelection(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSelectedText(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSelectionLine(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSelectionColumn(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSelectionFromLine(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSelectionFromColumn(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSelectionToLine(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSelectionToColumn(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) Deselect(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) DeleteSelection(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineWrappingMode(mode TextEditLineWrappingMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineWrappingMode() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetAutowrapMode() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsLineWrapped(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineWrapCount(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineWrapIndexAtColumn(line int, column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineWrappedText(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetSmoothScrollEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsSmoothScrollEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetVScrollBar() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetHScrollBar() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetVScroll(value float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetVScroll() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetHScroll(value int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetHScroll() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetScrollPastEndOfFileEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsScrollPastEndOfFileEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetVScrollSpeed(speed float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetVScrollSpeed() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetFitContentHeightEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsFitContentHeightEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetScrollPosForLine(line int, wrap_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineAsFirstVisible(line int, wrap_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetFirstVisibleLine() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineAsCenterVisible(line int, wrap_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineAsLastVisible(line int, wrap_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLastFullVisibleLine() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLastFullVisibleLineWrapIndex() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetVisibleLineCount() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetVisibleLineCountInRange(from_line int, to_line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetTotalVisibleLineCount() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) AdjustViewportToCaret(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) CenterViewportToCaret(caret_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetDrawMinimap(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsDrawingMinimap() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetMinimapWidth(width int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetMinimapWidth() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetMinimapVisibleLines() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) AddGutter(at int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) RemoveGutter(gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetGutterCount() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetGutterName(gutter int, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetGutterName(gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetGutterType(gutter int, type_ TextEditGutterType, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetGutterType(gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetGutterWidth(gutter int, width int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetGutterWidth(gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetGutterDraw(gutter int, draw bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsGutterDrawn(gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetGutterClickable(gutter int, clickable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsGutterClickable(gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetGutterOverwritable(gutter int, overwritable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsGutterOverwritable(gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) MergeGutters(from_line int, to_line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetGutterCustomDraw(column int, draw_callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetTotalGutterWidth() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineGutterMetadata(line int, gutter int, metadata Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineGutterMetadata(line int, gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineGutterText(line int, gutter int, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineGutterText(line int, gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineGutterIcon(line int, gutter int, icon Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineGutterIcon(line int, gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineGutterItemColor(line int, gutter int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineGutterItemColor(line int, gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineGutterClickable(line int, gutter int, clickable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsLineGutterClickable(line int, gutter int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetLineBackgroundColor(line int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetLineBackgroundColor(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetSyntaxHighlighter(syntax_highlighter SyntaxHighlighter, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetSyntaxHighlighter() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetHighlightCurrentLine(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsHighlightCurrentLineEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetHighlightAllOccurrences(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsHighlightAllOccurrencesEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetDrawControlChars() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetDrawControlChars(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetDrawTabs(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsDrawingTabs() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) SetDrawSpaces(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsDrawingSpaces() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) GetMenu() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) IsMenuVisible() { // TODO: return value
  // TODO: implement
}

func  (me *TextEdit) MenuOption(option int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
