// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *RichTextLabel) GetParsedText() String {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parsed_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) AddText(text String, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) SetText(text String, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) AddImage(image Texture2D, width int, height int, color Color, inline_align InlineAlignment, region Rect2, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3346058748) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(image.AsCTypePtr()), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(region.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) Newline()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("newline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) RemoveParagraph(paragraph int, ) bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_paragraph")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3067735520) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paragraph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) PushFont(font Font, font_size int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3014009009) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushFontSize(font_size int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushNormal()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushBold()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_bold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushBoldItalics()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_bold_italics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushItalics()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_italics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushMono()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_mono")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushColor(color Color, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushOutlineSize(outline_size int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_outline_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&outline_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushOutlineColor(color Color, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_outline_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushParagraph(alignment HorizontalAlignment, base_direction ControlTextDirection, language String, st_parser TextServerStructuredTextParser, justification_flags TextServerJustificationFlag, tab_stops PackedFloat32Array, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_paragraph")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218895358) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), gdc.ConstTypePtr(&base_direction), gdc.ConstTypePtr(language.AsCTypePtr()), gdc.ConstTypePtr(&st_parser), gdc.ConstTypePtr(&justification_flags), gdc.ConstTypePtr(tab_stops.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushIndent(level int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_indent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&level), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushList(level int, type_ RichTextLabelListType, capitalize bool, bullet String, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4036303897) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&level), gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&capitalize), gdc.ConstTypePtr(bullet.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushMeta(data Variant, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1114965689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushHint(description String, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_hint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(description.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushUnderline()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_underline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushStrikethrough()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_strikethrough")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushTable(columns int, inline_align InlineAlignment, align_to_row int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_table")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1125058220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&columns), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&align_to_row), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushDropcap(string_ String, font Font, size int, dropcap_margins Rect2, color Color, outline_size int, outline_color Color, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_dropcap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 311501835) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(dropcap_margins.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), gdc.ConstTypePtr(&outline_size), gdc.ConstTypePtr(outline_color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) SetTableColumnExpand(column int, expand bool, ratio int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_table_column_expand")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4258957458) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&column), gdc.ConstTypePtr(&expand), gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) SetCellRowBackgroundColor(odd_row_bg Color, even_row_bg Color, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_row_background_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3465483165) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(odd_row_bg.AsCTypePtr()), gdc.ConstTypePtr(even_row_bg.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) SetCellBorderColor(color Color, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_border_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) SetCellSizeOverride(min_size Vector2, max_size Vector2, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_size_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3108078480) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(min_size.AsCTypePtr()), gdc.ConstTypePtr(max_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) SetCellPadding(padding Rect2, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cell_padding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(padding.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushCell()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_cell")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushFgcolor(fgcolor Color, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_fgcolor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(fgcolor.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushBgcolor(bgcolor Color, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_bgcolor")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bgcolor.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) PushCustomfx(effect RichTextEffect, env Dictionary, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("push_customfx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2337942958) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(effect.AsCTypePtr()), gdc.ConstTypePtr(env.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) Pop()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) Clear()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 55961453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3385126229) // FIXME: should cache?
  var ret TextServerStructuredTextParser
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetStructuredTextBidiOverrideOptions(args Array, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(args.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetStructuredTextBidiOverrideOptions() Array {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetTextDirection(direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 119160795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetTextDirection() ControlTextDirection {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797257663) // FIXME: should cache?
  var ret ControlTextDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetLanguage() String {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3289138044) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetAutowrapMode() TextServerAutowrapMode {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1549071663) // FIXME: should cache?
  var ret TextServerAutowrapMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetMetaUnderline(enable bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_meta_underline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsMetaUnderlined() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_meta_underlined")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetHintUnderline(enable bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hint_underline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsHintUnderlined() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hint_underlined")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetScrollActive(active bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsScrollActive() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_scroll_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetScrollFollow(follow bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_follow")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&follow), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsScrollFollowing() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_scroll_following")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetVScrollBar() VScrollBar {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_scroll_bar")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2630340773) // FIXME: should cache?
  var ret VScrollBar
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) ScrollToLine(line int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scroll_to_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) ScrollToParagraph(paragraph int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scroll_to_paragraph")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paragraph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) ScrollToSelection()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scroll_to_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) SetTabSize(spaces int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&spaces), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetTabSize() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetFitContent(enabled bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fit_content")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsFitContentEnabled() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_fit_content_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetSelectionEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsSelectionEnabled() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetContextMenuEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_context_menu_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsContextMenuEnabled() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_context_menu_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetShortcutKeysEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shortcut_keys_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsShortcutKeysEnabled() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_shortcut_keys_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetDeselectOnFocusLossEnabled(enable bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_deselect_on_focus_loss_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsDeselectOnFocusLossEnabled() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_deselect_on_focus_loss_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetSelectionFrom() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_from")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetSelectionTo() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selection_to")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SelectAll()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetSelectedText() String {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) Deselect()  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) ParseBbcode(bbcode String, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_bbcode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bbcode.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) AppendText(bbcode String, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("append_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bbcode.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetText() String {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) IsReady() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_ready")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetThreaded(threaded bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_threaded")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threaded), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsThreaded() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_threaded")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetProgressBarDelay(delay_ms int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_progress_bar_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delay_ms), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetProgressBarDelay() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_progress_bar_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetVisibleCharacters(amount int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible_characters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetVisibleCharacters() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_characters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetVisibleCharactersBehavior() TextServerVisibleCharactersBehavior {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_characters_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 258789322) // FIXME: should cache?
  var ret TextServerVisibleCharactersBehavior
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetVisibleCharactersBehavior(behavior TextServerVisibleCharactersBehavior, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible_characters_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3383839701) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&behavior), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) SetVisibleRatio(ratio float32, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetVisibleRatio() float32 {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetCharacterLine(character int, ) int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_character_line")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&character), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetCharacterParagraph(character int, ) int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_character_paragraph")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&character), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetTotalCharacterCount() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_character_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetUseBbcode(enable bool, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_bbcode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) IsUsingBbcode() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_bbcode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetLineCount() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetVisibleLineCount() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetParagraphCount() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_paragraph_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetVisibleParagraphCount() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_paragraph_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetContentHeight() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_content_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetContentWidth() int {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_content_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetLineOffset(line int, ) float32 {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4025615559) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) GetParagraphOffset(paragraph int, ) float32 {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_paragraph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4025615559) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&paragraph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) ParseExpressionsForValues(expressions PackedStringArray, ) Dictionary {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_expressions_for_values")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1522900837) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(expressions.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) SetEffects(effects Array, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_effects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(effects.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetEffects() Array {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_effects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) InstallEffect(effect Variant, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("install_effect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1114965689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(effect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RichTextLabel) GetMenu() PopupMenu {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_menu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 229722558) // FIXME: should cache?
  var ret PopupMenu
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) IsMenuVisible() bool {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_menu_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RichTextLabel) MenuOption(option int, )  {
  classNameV := StringNameFromStr("RichTextLabel")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("menu_option")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *RichTextLabel) GetPropBbcodeEnabled() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropBbcodeEnabled(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropText() String {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropText(value String) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropFitContent() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropFitContent(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropScrollActive() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropScrollActive(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropScrollFollowing() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropScrollFollowing(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropAutowrapMode() int {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropAutowrapMode(value int) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropTabSize() int {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropTabSize(value int) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropContextMenuEnabled() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropContextMenuEnabled(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropShortcutKeysEnabled() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropShortcutKeysEnabled(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropCustomEffects() RichTextEffect {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropCustomEffects(value RichTextEffect) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropMetaUnderlined() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropMetaUnderlined(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropHintUnderlined() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropHintUnderlined(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropThreaded() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropThreaded(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropProgressBarDelay() int {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropProgressBarDelay(value int) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropSelectionEnabled() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropSelectionEnabled(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropDeselectOnFocusLossEnabled() bool {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropDeselectOnFocusLossEnabled(value bool) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropVisibleCharacters() int {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropVisibleCharacters(value int) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropVisibleCharactersBehavior() int {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropVisibleCharactersBehavior(value int) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropVisibleRatio() float32 {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropVisibleRatio(value float32) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropTextDirection() int {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropTextDirection(value int) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropLanguage() String {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropLanguage(value String) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropStructuredTextBidiOverride() int {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropStructuredTextBidiOverride(value int) {
  panic("TODO: implement")
}

func (me *RichTextLabel) GetPropStructuredTextBidiOverrideOptions() Array {
  panic("TODO: implement")
}

func (me *RichTextLabel) SetPropStructuredTextBidiOverrideOptions(value Array) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API