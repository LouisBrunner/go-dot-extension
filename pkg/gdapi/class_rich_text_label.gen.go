// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RichTextLabel struct {
  obj gdc.ObjectPtr
}

func (me *RichTextLabel) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RichTextLabel) BaseClass() string {
  return "RichTextLabel"
}

type RichTextLabelListType int
const (
  RichTextLabelListTypeListNumbers RichTextLabelListType = 0
  RichTextLabelListTypeListLetters RichTextLabelListType = 1
  RichTextLabelListTypeListRoman RichTextLabelListType = 2
  RichTextLabelListTypeListDots RichTextLabelListType = 3
)

type RichTextLabelMenuItems int
const (
  RichTextLabelMenuItemsMenuCopy RichTextLabelMenuItems = 0
  RichTextLabelMenuItemsMenuSelectAll RichTextLabelMenuItems = 1
  RichTextLabelMenuItemsMenuMax RichTextLabelMenuItems = 2
)

func  (me *RichTextLabel) GetParsedText() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) AddText(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetText(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) AddImage(image Texture2D, width int, height int, color Color, inline_align InlineAlignment, region Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) Newline() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) RemoveParagraph(paragraph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushFont(font Font, font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushFontSize(font_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushNormal() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushBold() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushBoldItalics() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushItalics() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushMono() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushOutlineSize(outline_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushOutlineColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushParagraph(alignment HorizontalAlignment, base_direction ControlTextDirection, language String, st_parser TextServerStructuredTextParser, justification_flags TextServerJustificationFlag, tab_stops PackedFloat32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushIndent(level int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushList(level int, type_ RichTextLabelListType, capitalize bool, bullet String, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushMeta(data Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushHint(description String, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushUnderline() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushStrikethrough() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushTable(columns int, inline_align InlineAlignment, align_to_row int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushDropcap(string String, font Font, size int, dropcap_margins Rect2, color Color, outline_size int, outline_color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetTableColumnExpand(column int, expand bool, ratio int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetCellRowBackgroundColor(odd_row_bg Color, even_row_bg Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetCellBorderColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetCellSizeOverride(min_size Vector2, max_size Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetCellPadding(padding Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushCell() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushFgcolor(fgcolor Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushBgcolor(bgcolor Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) PushCustomfx(effect RichTextEffect, env Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) Pop() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) Clear() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetStructuredTextBidiOverride() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetStructuredTextBidiOverrideOptions(args Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetStructuredTextBidiOverrideOptions() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetTextDirection(direction ControlTextDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetTextDirection() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetLanguage(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetLanguage() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetAutowrapMode() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetMetaUnderline(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsMetaUnderlined() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetHintUnderline(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsHintUnderlined() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetScrollActive(active bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsScrollActive() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetScrollFollow(follow bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsScrollFollowing() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetVScrollBar() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) ScrollToLine(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) ScrollToParagraph(paragraph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) ScrollToSelection() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetTabSize(spaces int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetTabSize() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetFitContent(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsFitContentEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetSelectionEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsSelectionEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetContextMenuEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsContextMenuEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetShortcutKeysEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsShortcutKeysEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetDeselectOnFocusLossEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsDeselectOnFocusLossEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetSelectionFrom() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetSelectionTo() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SelectAll() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetSelectedText() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) Deselect() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) ParseBbcode(bbcode String, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) AppendText(bbcode String, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetText() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsReady() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetThreaded(threaded bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsThreaded() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetProgressBarDelay(delay_ms int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetProgressBarDelay() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetVisibleCharacters(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetVisibleCharacters() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetVisibleCharactersBehavior() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetVisibleCharactersBehavior(behavior TextServerVisibleCharactersBehavior, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetVisibleRatio(ratio float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetVisibleRatio() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetCharacterLine(character int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetCharacterParagraph(character int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetTotalCharacterCount() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetUseBbcode(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsUsingBbcode() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetLineCount() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetVisibleLineCount() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetParagraphCount() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetVisibleParagraphCount() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetContentHeight() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetContentWidth() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetLineOffset(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetParagraphOffset(paragraph int, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) ParseExpressionsForValues(expressions PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) SetEffects(effects Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetEffects() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) InstallEffect(effect Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) GetMenu() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) IsMenuVisible() { // TODO: return value
  // TODO: implement
}

func  (me *RichTextLabel) MenuOption(option int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
