// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Label3D struct {
  obj gdc.ObjectPtr
}

func (me *Label3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Label3D) BaseClass() string {
  return "Label3D"
}



// Enums

type Label3DDrawFlags int
const (
  Label3DDrawFlagsFlagShaded Label3DDrawFlags = 0
  Label3DDrawFlagsFlagDoubleSided Label3DDrawFlags = 1
  Label3DDrawFlagsFlagDisableDepthTest Label3DDrawFlags = 2
  Label3DDrawFlagsFlagFixedSize Label3DDrawFlags = 3
  Label3DDrawFlagsFlagMax Label3DDrawFlags = 4
)

type Label3DAlphaCutMode int
const (
  Label3DAlphaCutModeAlphaCutDisabled Label3DAlphaCutMode = 0
  Label3DAlphaCutModeAlphaCutDiscard Label3DAlphaCutMode = 1
  Label3DAlphaCutModeAlphaCutOpaquePrepass Label3DAlphaCutMode = 2
  Label3DAlphaCutModeAlphaCutHash Label3DAlphaCutMode = 3
)

func (me *Label3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Label3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Label3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Label3D) SetHorizontalAlignment(alignment HorizontalAlignment, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2312603777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetHorizontalAlignment() HorizontalAlignment {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_horizontal_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 341400642) // FIXME: should cache?
  var ret HorizontalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetVerticalAlignment(alignment VerticalAlignment, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertical_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1796458609) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetVerticalAlignment() VerticalAlignment {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_vertical_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3274884059) // FIXME: should cache?
  var ret VerticalAlignment
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetModulate(modulate Color, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetModulate() Color {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetOutlineModulate(modulate Color, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outline_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(modulate.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetOutlineModulate() Color {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outline_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetText(text String, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetText() String {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetTextDirection(direction TextServerDirection, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1418190634) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetTextDirection() TextServerDirection {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2516697328) // FIXME: should cache?
  var ret TextServerDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetLanguage() String {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 55961453) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3385126229) // FIXME: should cache?
  var ret TextServerStructuredTextParser
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetStructuredTextBidiOverrideOptions(args Array, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(args.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetStructuredTextBidiOverrideOptions() Array {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_structured_text_bidi_override_options")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetUppercase(enable bool, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uppercase")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) IsUppercase() bool {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_uppercase")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetRenderPriority(priority int, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_render_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetRenderPriority() int {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_render_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetOutlineRenderPriority(priority int, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outline_render_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetOutlineRenderPriority() int {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outline_render_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetFont(font Font, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1262170328) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetFont() Font {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229501585) // FIXME: should cache?
  var ret Font
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetFontSize(size int, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetFontSize() int {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_font_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetOutlineSize(outline_size int, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outline_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&outline_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetOutlineSize() int {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outline_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetLineSpacing(line_spacing float32, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_line_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line_spacing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetLineSpacing() float32 {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_line_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3289138044) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetAutowrapMode() TextServerAutowrapMode {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_autowrap_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1549071663) // FIXME: should cache?
  var ret TextServerAutowrapMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetJustificationFlags(justification_flags TextServerJustificationFlag, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2877345813) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&justification_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetJustificationFlags() TextServerJustificationFlag {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_justification_flags")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1583363614) // FIXME: should cache?
  var ret TextServerJustificationFlag
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetWidth(width float32, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetWidth() float32 {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetPixelSize(pixel_size float32, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pixel_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixel_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetPixelSize() float32 {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pixel_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetOffset() Vector2 {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetDrawFlag(flag Label3DDrawFlags, enabled bool, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draw_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1285833066) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetDrawFlag(flag Label3DDrawFlags, ) bool {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_draw_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259226453) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_billboard_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4202036497) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetBillboardMode() BaseMaterial3DBillboardMode {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_billboard_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1283840139) // FIXME: should cache?
  var ret BaseMaterial3DBillboardMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetAlphaCutMode(mode Label3DAlphaCutMode, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_cut_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2549142916) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetAlphaCutMode() Label3DAlphaCutMode {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_cut_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 219468601) // FIXME: should cache?
  var ret Label3DAlphaCutMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetAlphaScissorThreshold(threshold float32, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_scissor_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetAlphaScissorThreshold() float32 {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_scissor_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetAlphaHashScale(threshold float32, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_hash_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetAlphaHashScale() float32 {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_hash_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3212649852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alpha_aa), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetAlphaAntialiasing() BaseMaterial3DAlphaAntiAliasing {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2889939400) // FIXME: should cache?
  var ret BaseMaterial3DAlphaAntiAliasing
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetAlphaAntialiasingEdge(edge float32, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_alpha_antialiasing_edge")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&edge), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetAlphaAntialiasingEdge() float32 {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_alpha_antialiasing_edge")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, )  {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 22904437) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Label3D) GetTextureFilter() BaseMaterial3DTextureFilter {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3289213076) // FIXME: should cache?
  var ret BaseMaterial3DTextureFilter
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Label3D) GenerateTriangleMesh() TriangleMesh {
  classNameV := StringNameFromStr("Label3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_triangle_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3476533166) // FIXME: should cache?
  var ret TriangleMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Label3D) GetPropPixelSize() float32 {
  panic("TODO: implement")
}

func (me *Label3D) SetPropPixelSize(value float32) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropOffset() Vector2 {
  panic("TODO: implement")
}

func (me *Label3D) SetPropOffset(value Vector2) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropBillboard() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropBillboard(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropShaded() bool {
  panic("TODO: implement")
}

func (me *Label3D) SetPropShaded(value bool) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropDoubleSided() bool {
  panic("TODO: implement")
}

func (me *Label3D) SetPropDoubleSided(value bool) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropNoDepthTest() bool {
  panic("TODO: implement")
}

func (me *Label3D) SetPropNoDepthTest(value bool) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropFixedSize() bool {
  panic("TODO: implement")
}

func (me *Label3D) SetPropFixedSize(value bool) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropAlphaCut() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropAlphaCut(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropAlphaScissorThreshold() float32 {
  panic("TODO: implement")
}

func (me *Label3D) SetPropAlphaScissorThreshold(value float32) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropAlphaHashScale() float32 {
  panic("TODO: implement")
}

func (me *Label3D) SetPropAlphaHashScale(value float32) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropAlphaAntialiasingMode() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropAlphaAntialiasingMode(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropAlphaAntialiasingEdge() float32 {
  panic("TODO: implement")
}

func (me *Label3D) SetPropAlphaAntialiasingEdge(value float32) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropTextureFilter() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropTextureFilter(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropRenderPriority() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropRenderPriority(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropOutlineRenderPriority() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropOutlineRenderPriority(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropModulate() Color {
  panic("TODO: implement")
}

func (me *Label3D) SetPropModulate(value Color) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropOutlineModulate() Color {
  panic("TODO: implement")
}

func (me *Label3D) SetPropOutlineModulate(value Color) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropText() String {
  panic("TODO: implement")
}

func (me *Label3D) SetPropText(value String) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropFont() Font {
  panic("TODO: implement")
}

func (me *Label3D) SetPropFont(value Font) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropFontSize() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropFontSize(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropOutlineSize() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropOutlineSize(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropHorizontalAlignment() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropHorizontalAlignment(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropVerticalAlignment() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropVerticalAlignment(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropUppercase() bool {
  panic("TODO: implement")
}

func (me *Label3D) SetPropUppercase(value bool) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropLineSpacing() float32 {
  panic("TODO: implement")
}

func (me *Label3D) SetPropLineSpacing(value float32) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropAutowrapMode() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropAutowrapMode(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropJustificationFlags() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropJustificationFlags(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropWidth() float32 {
  panic("TODO: implement")
}

func (me *Label3D) SetPropWidth(value float32) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropTextDirection() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropTextDirection(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropLanguage() String {
  panic("TODO: implement")
}

func (me *Label3D) SetPropLanguage(value String) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropStructuredTextBidiOverride() int {
  panic("TODO: implement")
}

func (me *Label3D) SetPropStructuredTextBidiOverride(value int) {
  panic("TODO: implement")
}

func (me *Label3D) GetPropStructuredTextBidiOverrideOptions() Array {
  panic("TODO: implement")
}

func (me *Label3D) SetPropStructuredTextBidiOverrideOptions(value Array) {
  panic("TODO: implement")
}