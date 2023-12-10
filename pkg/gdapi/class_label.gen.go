// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Label struct {
  obj gdc.ObjectPtr
}

func (me *Label) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Label) BaseClass() string {
  return "Label"
}



// Enums

func (me *Label) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Label) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Label) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Label) SetHorizontalAlignment(alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312603777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetHorizontalAlignment() HorizontalAlignment {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 341400642) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetVerticalAlignment(alignment VerticalAlignment, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1796458609) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetVerticalAlignment() VerticalAlignment {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertical_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3274884059) // FIXME: should cache?
  var ret VerticalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetText(text String, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetText() String {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetLabelSettings(settings LabelSettings, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_label_settings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1030653839) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(settings.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetLabelSettings() LabelSettings {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_label_settings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 826676056) // FIXME: should cache?
  var ret LabelSettings
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetTextDirection(direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 119160795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetTextDirection() ControlTextDirection {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797257663) // FIXME: should cache?
  var ret ControlTextDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetLanguage() String {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3289138044) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetAutowrapMode() TextServerAutowrapMode {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1549071663) // FIXME: should cache?
  var ret TextServerAutowrapMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetJustificationFlags(justification_flags TextServerJustificationFlag, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2877345813) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&justification_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetJustificationFlags() TextServerJustificationFlag {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1583363614) // FIXME: should cache?
  var ret TextServerJustificationFlag
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetClipText(enable bool, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clip_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) IsClippingText() bool {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_clipping_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetTabStops(tab_stops PackedFloat32Array, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_stops")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tab_stops.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetTabStops() PackedFloat32Array {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_stops")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675695659) // FIXME: should cache?
  var ret PackedFloat32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1008890932) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetTextOverrunBehavior() TextServerOverrunBehavior {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3779142101) // FIXME: should cache?
  var ret TextServerOverrunBehavior
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetUppercase(enable bool, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uppercase")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) IsUppercase() bool {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_uppercase")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) GetLineHeight(line int, ) int {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 181039630) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) GetLineCount() int {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) GetVisibleLineCount() int {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_line_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) GetTotalCharacterCount() int {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_total_character_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetVisibleCharacters(amount int, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible_characters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetVisibleCharacters() int {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_characters")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) GetVisibleCharactersBehavior() TextServerVisibleCharactersBehavior {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_characters_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 258789322) // FIXME: should cache?
  var ret TextServerVisibleCharactersBehavior
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetVisibleCharactersBehavior(behavior TextServerVisibleCharactersBehavior, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible_characters_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3383839701) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&behavior), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) SetVisibleRatio(ratio float32, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_visible_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetVisibleRatio() float32 {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_visible_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetLinesSkipped(lines_skipped int, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_lines_skipped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lines_skipped), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetLinesSkipped() int {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_lines_skipped")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetMaxLinesVisible(lines_visible int, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_lines_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lines_visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetMaxLinesVisible() int {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_lines_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 55961453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3385126229) // FIXME: should cache?
  var ret TextServerStructuredTextParser
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label) SetStructuredTextBidiOverrideOptions(args Array, )  {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(args.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label) GetStructuredTextBidiOverrideOptions() Array {
  classNameV := StringNameFromStr("Label")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Label) GetPropText() String {
  panic("TODO: implement")
}

func (me *Label) SetPropText(value String) {
  panic("TODO: implement")
}

func (me *Label) GetPropLabelSettings() LabelSettings {
  panic("TODO: implement")
}

func (me *Label) SetPropLabelSettings(value LabelSettings) {
  panic("TODO: implement")
}

func (me *Label) GetPropHorizontalAlignment() int {
  panic("TODO: implement")
}

func (me *Label) SetPropHorizontalAlignment(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropVerticalAlignment() int {
  panic("TODO: implement")
}

func (me *Label) SetPropVerticalAlignment(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropAutowrapMode() int {
  panic("TODO: implement")
}

func (me *Label) SetPropAutowrapMode(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropJustificationFlags() int {
  panic("TODO: implement")
}

func (me *Label) SetPropJustificationFlags(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropClipText() bool {
  panic("TODO: implement")
}

func (me *Label) SetPropClipText(value bool) {
  panic("TODO: implement")
}

func (me *Label) GetPropTextOverrunBehavior() int {
  panic("TODO: implement")
}

func (me *Label) SetPropTextOverrunBehavior(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropUppercase() bool {
  panic("TODO: implement")
}

func (me *Label) SetPropUppercase(value bool) {
  panic("TODO: implement")
}

func (me *Label) GetPropTabStops() PackedFloat32Array {
  panic("TODO: implement")
}

func (me *Label) SetPropTabStops(value PackedFloat32Array) {
  panic("TODO: implement")
}

func (me *Label) GetPropLinesSkipped() int {
  panic("TODO: implement")
}

func (me *Label) SetPropLinesSkipped(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropMaxLinesVisible() int {
  panic("TODO: implement")
}

func (me *Label) SetPropMaxLinesVisible(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropVisibleCharacters() int {
  panic("TODO: implement")
}

func (me *Label) SetPropVisibleCharacters(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropVisibleCharactersBehavior() int {
  panic("TODO: implement")
}

func (me *Label) SetPropVisibleCharactersBehavior(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropVisibleRatio() float32 {
  panic("TODO: implement")
}

func (me *Label) SetPropVisibleRatio(value float32) {
  panic("TODO: implement")
}

func (me *Label) GetPropTextDirection() int {
  panic("TODO: implement")
}

func (me *Label) SetPropTextDirection(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropLanguage() String {
  panic("TODO: implement")
}

func (me *Label) SetPropLanguage(value String) {
  panic("TODO: implement")
}

func (me *Label) GetPropStructuredTextBidiOverride() int {
  panic("TODO: implement")
}

func (me *Label) SetPropStructuredTextBidiOverride(value int) {
  panic("TODO: implement")
}

func (me *Label) GetPropStructuredTextBidiOverrideOptions() Array {
  panic("TODO: implement")
}

func (me *Label) SetPropStructuredTextBidiOverrideOptions(value Array) {
  panic("TODO: implement")
}