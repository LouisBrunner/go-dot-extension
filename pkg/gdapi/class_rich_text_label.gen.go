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



// Enums

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

func (me *RichTextLabel) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RichTextLabel) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RichTextLabel) GetParsedText()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) AddText(text String, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetText(text String, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) AddImage(image Texture2D, width int, height int, color Color, inline_align InlineAlignment, region Rect2, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) Newline()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) RemoveParagraph(paragraph int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushFont(font Font, font_size int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushFontSize(font_size int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushNormal()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushBold()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushBoldItalics()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushItalics()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushMono()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushOutlineSize(outline_size int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushOutlineColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushParagraph(alignment HorizontalAlignment, base_direction ControlTextDirection, language String, st_parser TextServerStructuredTextParser, justification_flags TextServerJustificationFlag, tab_stops PackedFloat32Array, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushIndent(level int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushList(level int, type_ RichTextLabelListType, capitalize bool, bullet String, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushMeta(data Variant, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushHint(description String, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushUnderline()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushStrikethrough()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushTable(columns int, inline_align InlineAlignment, align_to_row int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushDropcap(string_ String, font Font, size int, dropcap_margins Rect2, color Color, outline_size int, outline_color Color, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetTableColumnExpand(column int, expand bool, ratio int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetCellRowBackgroundColor(odd_row_bg Color, even_row_bg Color, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetCellBorderColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetCellSizeOverride(min_size Vector2, max_size Vector2, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetCellPadding(padding Rect2, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushCell()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushFgcolor(fgcolor Color, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushBgcolor(bgcolor Color, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) PushCustomfx(effect RichTextEffect, env Dictionary, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) Pop()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) Clear()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetStructuredTextBidiOverride()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetStructuredTextBidiOverrideOptions(args Array, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetStructuredTextBidiOverrideOptions()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetTextDirection(direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetTextDirection()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetLanguage(language String, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetLanguage()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetAutowrapMode()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetMetaUnderline(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsMetaUnderlined()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetHintUnderline(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsHintUnderlined()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetScrollActive(active bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsScrollActive()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetScrollFollow(follow bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsScrollFollowing()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetVScrollBar()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) ScrollToLine(line int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) ScrollToParagraph(paragraph int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) ScrollToSelection()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetTabSize(spaces int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetTabSize()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetFitContent(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsFitContentEnabled()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetSelectionEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsSelectionEnabled()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetContextMenuEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsContextMenuEnabled()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetShortcutKeysEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsShortcutKeysEnabled()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetDeselectOnFocusLossEnabled(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsDeselectOnFocusLossEnabled()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetSelectionFrom()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetSelectionTo()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SelectAll()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetSelectedText()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) Deselect()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) ParseBbcode(bbcode String, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) AppendText(bbcode String, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetText()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsReady()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetThreaded(threaded bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsThreaded()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetProgressBarDelay(delay_ms int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetProgressBarDelay()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetVisibleCharacters(amount int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetVisibleCharacters()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetVisibleCharactersBehavior()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetVisibleCharactersBehavior(behavior TextServerVisibleCharactersBehavior, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetVisibleRatio(ratio float32, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetVisibleRatio()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetCharacterLine(character int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetCharacterParagraph(character int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetTotalCharacterCount()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetUseBbcode(enable bool, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsUsingBbcode()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetLineCount()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetVisibleLineCount()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetParagraphCount()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetVisibleParagraphCount()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetContentHeight()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetContentWidth()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetLineOffset(line int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetParagraphOffset(paragraph int, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) ParseExpressionsForValues(expressions PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) SetEffects(effects Array, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetEffects()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) InstallEffect(effect Variant, )  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) GetMenu()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) IsMenuVisible()  {
  panic("TODO: implement")
}

func  (me *RichTextLabel) MenuOption(option int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
