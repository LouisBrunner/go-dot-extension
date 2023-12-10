// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextMesh struct {
  obj gdc.ObjectPtr
}

func (me *TextMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextMesh) BaseClass() string {
  return "TextMesh"
}



// Enums

func (me *TextMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TextMesh) SetHorizontalAlignment(alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312603777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetHorizontalAlignment() HorizontalAlignment {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 341400642) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetVerticalAlignment(alignment VerticalAlignment, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1796458609) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetVerticalAlignment() VerticalAlignment {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertical_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3274884059) // FIXME: should cache?
  var ret VerticalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetText(text String, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetText() String {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetFont(font Font, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1262170328) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetFont() Font {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229501585) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetFontSize(font_size int, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&font_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetFontSize() int {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetLineSpacing(line_spacing float32, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line_spacing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetLineSpacing() float32 {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3289138044) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetAutowrapMode() TextServerAutowrapMode {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1549071663) // FIXME: should cache?
  var ret TextServerAutowrapMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetJustificationFlags(justification_flags TextServerJustificationFlag, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2877345813) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&justification_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetJustificationFlags() TextServerJustificationFlag {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1583363614) // FIXME: should cache?
  var ret TextServerJustificationFlag
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetDepth(depth float32, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetDepth() float32 {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_depth")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetWidth(width float32, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetWidth() float32 {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetPixelSize(pixel_size float32, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pixel_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixel_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetPixelSize() float32 {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pixel_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetOffset() Vector2 {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetCurveStep(curve_step float32, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve_step), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetCurveStep() float32 {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetTextDirection(direction TextServerDirection, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1418190634) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetTextDirection() TextServerDirection {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2516697328) // FIXME: should cache?
  var ret TextServerDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetLanguage() String {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 55961453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3385126229) // FIXME: should cache?
  var ret TextServerStructuredTextParser
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetStructuredTextBidiOverrideOptions(args Array, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(args.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) GetStructuredTextBidiOverrideOptions() Array {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextMesh) SetUppercase(enable bool, )  {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uppercase")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextMesh) IsUppercase() bool {
  classNameV := StringNameFromStr("TextMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_uppercase")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *TextMesh) GetPropText() String {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropText(value String) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropFont() Font {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropFont(value Font) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropFontSize() int {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropFontSize(value int) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropHorizontalAlignment() int {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropHorizontalAlignment(value int) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropVerticalAlignment() int {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropVerticalAlignment(value int) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropUppercase() bool {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropUppercase(value bool) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropLineSpacing() float32 {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropLineSpacing(value float32) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropAutowrapMode() int {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropAutowrapMode(value int) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropJustificationFlags() int {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropJustificationFlags(value int) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropPixelSize() float32 {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropPixelSize(value float32) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropCurveStep() float32 {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropCurveStep(value float32) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropDepth() float32 {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropDepth(value float32) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropWidth() float32 {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropWidth(value float32) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropOffset() Vector2 {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropTextDirection() int {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropTextDirection(value int) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropLanguage() String {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropLanguage(value String) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropStructuredTextBidiOverride() int {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropStructuredTextBidiOverride(value int) {
  panic("TODO: implement")
}

func (me *TextMesh) GetPropStructuredTextBidiOverrideOptions() Array {
  panic("TODO: implement")
}

func (me *TextMesh) SetPropStructuredTextBidiOverrideOptions(value Array) {
  panic("TODO: implement")
}