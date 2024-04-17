// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForLabelList struct {
	fnSetHorizontalAlignment               gdc.MethodBindPtr
	fnGetHorizontalAlignment               gdc.MethodBindPtr
	fnSetVerticalAlignment                 gdc.MethodBindPtr
	fnGetVerticalAlignment                 gdc.MethodBindPtr
	fnSetText                              gdc.MethodBindPtr
	fnGetText                              gdc.MethodBindPtr
	fnSetLabelSettings                     gdc.MethodBindPtr
	fnGetLabelSettings                     gdc.MethodBindPtr
	fnSetTextDirection                     gdc.MethodBindPtr
	fnGetTextDirection                     gdc.MethodBindPtr
	fnSetLanguage                          gdc.MethodBindPtr
	fnGetLanguage                          gdc.MethodBindPtr
	fnSetAutowrapMode                      gdc.MethodBindPtr
	fnGetAutowrapMode                      gdc.MethodBindPtr
	fnSetJustificationFlags                gdc.MethodBindPtr
	fnGetJustificationFlags                gdc.MethodBindPtr
	fnSetClipText                          gdc.MethodBindPtr
	fnIsClippingText                       gdc.MethodBindPtr
	fnSetTabStops                          gdc.MethodBindPtr
	fnGetTabStops                          gdc.MethodBindPtr
	fnSetTextOverrunBehavior               gdc.MethodBindPtr
	fnGetTextOverrunBehavior               gdc.MethodBindPtr
	fnSetEllipsisChar                      gdc.MethodBindPtr
	fnGetEllipsisChar                      gdc.MethodBindPtr
	fnSetUppercase                         gdc.MethodBindPtr
	fnIsUppercase                          gdc.MethodBindPtr
	fnGetLineHeight                        gdc.MethodBindPtr
	fnGetLineCount                         gdc.MethodBindPtr
	fnGetVisibleLineCount                  gdc.MethodBindPtr
	fnGetTotalCharacterCount               gdc.MethodBindPtr
	fnSetVisibleCharacters                 gdc.MethodBindPtr
	fnGetVisibleCharacters                 gdc.MethodBindPtr
	fnGetVisibleCharactersBehavior         gdc.MethodBindPtr
	fnSetVisibleCharactersBehavior         gdc.MethodBindPtr
	fnSetVisibleRatio                      gdc.MethodBindPtr
	fnGetVisibleRatio                      gdc.MethodBindPtr
	fnSetLinesSkipped                      gdc.MethodBindPtr
	fnGetLinesSkipped                      gdc.MethodBindPtr
	fnSetMaxLinesVisible                   gdc.MethodBindPtr
	fnGetMaxLinesVisible                   gdc.MethodBindPtr
	fnSetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnGetStructuredTextBidiOverride        gdc.MethodBindPtr
	fnSetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnGetStructuredTextBidiOverrideOptions gdc.MethodBindPtr
	fnGetCharacterBounds                   gdc.MethodBindPtr
}

var ptrsForLabel ptrsForLabelList

func initLabelPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Label")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForLabel.fnSetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312603777))
	}
	{
		methodName := StringNameFromStr("get_horizontal_alignment")
		defer methodName.Destroy()
		ptrsForLabel.fnGetHorizontalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 341400642))
	}
	{
		methodName := StringNameFromStr("set_vertical_alignment")
		defer methodName.Destroy()
		ptrsForLabel.fnSetVerticalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1796458609))
	}
	{
		methodName := StringNameFromStr("get_vertical_alignment")
		defer methodName.Destroy()
		ptrsForLabel.fnGetVerticalAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3274884059))
	}
	{
		methodName := StringNameFromStr("set_text")
		defer methodName.Destroy()
		ptrsForLabel.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_text")
		defer methodName.Destroy()
		ptrsForLabel.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_label_settings")
		defer methodName.Destroy()
		ptrsForLabel.fnSetLabelSettings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1030653839))
	}
	{
		methodName := StringNameFromStr("get_label_settings")
		defer methodName.Destroy()
		ptrsForLabel.fnGetLabelSettings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 826676056))
	}
	{
		methodName := StringNameFromStr("set_text_direction")
		defer methodName.Destroy()
		ptrsForLabel.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 119160795))
	}
	{
		methodName := StringNameFromStr("get_text_direction")
		defer methodName.Destroy()
		ptrsForLabel.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797257663))
	}
	{
		methodName := StringNameFromStr("set_language")
		defer methodName.Destroy()
		ptrsForLabel.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_language")
		defer methodName.Destroy()
		ptrsForLabel.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_autowrap_mode")
		defer methodName.Destroy()
		ptrsForLabel.fnSetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3289138044))
	}
	{
		methodName := StringNameFromStr("get_autowrap_mode")
		defer methodName.Destroy()
		ptrsForLabel.fnGetAutowrapMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1549071663))
	}
	{
		methodName := StringNameFromStr("set_justification_flags")
		defer methodName.Destroy()
		ptrsForLabel.fnSetJustificationFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2877345813))
	}
	{
		methodName := StringNameFromStr("get_justification_flags")
		defer methodName.Destroy()
		ptrsForLabel.fnGetJustificationFlags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1583363614))
	}
	{
		methodName := StringNameFromStr("set_clip_text")
		defer methodName.Destroy()
		ptrsForLabel.fnSetClipText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_clipping_text")
		defer methodName.Destroy()
		ptrsForLabel.fnIsClippingText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tab_stops")
		defer methodName.Destroy()
		ptrsForLabel.fnSetTabStops = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
	}
	{
		methodName := StringNameFromStr("get_tab_stops")
		defer methodName.Destroy()
		ptrsForLabel.fnGetTabStops = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675695659))
	}
	{
		methodName := StringNameFromStr("set_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForLabel.fnSetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1008890932))
	}
	{
		methodName := StringNameFromStr("get_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForLabel.fnGetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3779142101))
	}
	{
		methodName := StringNameFromStr("set_ellipsis_char")
		defer methodName.Destroy()
		ptrsForLabel.fnSetEllipsisChar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_ellipsis_char")
		defer methodName.Destroy()
		ptrsForLabel.fnGetEllipsisChar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_uppercase")
		defer methodName.Destroy()
		ptrsForLabel.fnSetUppercase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_uppercase")
		defer methodName.Destroy()
		ptrsForLabel.fnIsUppercase = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_line_height")
		defer methodName.Destroy()
		ptrsForLabel.fnGetLineHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 181039630))
	}
	{
		methodName := StringNameFromStr("get_line_count")
		defer methodName.Destroy()
		ptrsForLabel.fnGetLineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_visible_line_count")
		defer methodName.Destroy()
		ptrsForLabel.fnGetVisibleLineCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_total_character_count")
		defer methodName.Destroy()
		ptrsForLabel.fnGetTotalCharacterCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_visible_characters")
		defer methodName.Destroy()
		ptrsForLabel.fnSetVisibleCharacters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_visible_characters")
		defer methodName.Destroy()
		ptrsForLabel.fnGetVisibleCharacters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_visible_characters_behavior")
		defer methodName.Destroy()
		ptrsForLabel.fnGetVisibleCharactersBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 258789322))
	}
	{
		methodName := StringNameFromStr("set_visible_characters_behavior")
		defer methodName.Destroy()
		ptrsForLabel.fnSetVisibleCharactersBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3383839701))
	}
	{
		methodName := StringNameFromStr("set_visible_ratio")
		defer methodName.Destroy()
		ptrsForLabel.fnSetVisibleRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_visible_ratio")
		defer methodName.Destroy()
		ptrsForLabel.fnGetVisibleRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_lines_skipped")
		defer methodName.Destroy()
		ptrsForLabel.fnSetLinesSkipped = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_lines_skipped")
		defer methodName.Destroy()
		ptrsForLabel.fnGetLinesSkipped = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_max_lines_visible")
		defer methodName.Destroy()
		ptrsForLabel.fnSetMaxLinesVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_lines_visible")
		defer methodName.Destroy()
		ptrsForLabel.fnGetMaxLinesVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForLabel.fnSetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 55961453))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override")
		defer methodName.Destroy()
		ptrsForLabel.fnGetStructuredTextBidiOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3385126229))
	}
	{
		methodName := StringNameFromStr("set_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForLabel.fnSetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_structured_text_bidi_override_options")
		defer methodName.Destroy()
		ptrsForLabel.fnGetStructuredTextBidiOverrideOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("get_character_bounds")
		defer methodName.Destroy()
		ptrsForLabel.fnGetCharacterBounds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3327874267))
	}

}

type Label struct {
	Control
}

func (me *Label) BaseClass() string {
	return "Label"
}

func NewLabel() *Label {
	str := StringNameFromStr("Label") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Label{}
	obj.SetBaseObject(objPtr)
	return obj
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

func (me *Label) SetHorizontalAlignment(alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetHorizontalAlignment() HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetHorizontalAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label) SetVerticalAlignment(alignment VerticalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetVerticalAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetVerticalAlignment() VerticalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VerticalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetVerticalAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label) SetText(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label) SetLabelSettings(settings LabelSettings) {
	cargs := []gdc.ConstTypePtr{settings.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetLabelSettings), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetLabelSettings() LabelSettings {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewLabelSettings()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetLabelSettings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label) SetTextDirection(direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetTextDirection() ControlTextDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label) SetLanguage(language String) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetLanguage() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label) SetAutowrapMode(autowrap_mode TextServerAutowrapMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&autowrap_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetAutowrapMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetAutowrapMode() TextServerAutowrapMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerAutowrapMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetAutowrapMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label) SetJustificationFlags(justification_flags TextServerJustificationFlag) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&justification_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetJustificationFlags), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetJustificationFlags() TextServerJustificationFlag {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerJustificationFlag

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetJustificationFlags), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label) SetClipText(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetClipText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) IsClippingText() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnIsClippingText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) SetTabStops(tab_stops PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{tab_stops.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetTabStops), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetTabStops() PackedFloat32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedFloat32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetTabStops), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetTextOverrunBehavior() TextServerOverrunBehavior {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerOverrunBehavior

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label) SetEllipsisChar(char String) {
	cargs := []gdc.ConstTypePtr{char.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetEllipsisChar), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetEllipsisChar() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetEllipsisChar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label) SetUppercase(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetUppercase), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) IsUppercase() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnIsUppercase), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) GetLineHeight(line int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&line)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetLineHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) GetLineCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetLineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) GetVisibleLineCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetVisibleLineCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) GetTotalCharacterCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetTotalCharacterCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) SetVisibleCharacters(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetVisibleCharacters), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetVisibleCharacters() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetVisibleCharacters), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) GetVisibleCharactersBehavior() TextServerVisibleCharactersBehavior {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerVisibleCharactersBehavior

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetVisibleCharactersBehavior), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label) SetVisibleCharactersBehavior(behavior TextServerVisibleCharactersBehavior) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&behavior)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetVisibleCharactersBehavior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) SetVisibleRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetVisibleRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetVisibleRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetVisibleRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) SetLinesSkipped(lines_skipped int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lines_skipped)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetLinesSkipped), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetLinesSkipped() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetLinesSkipped), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) SetMaxLinesVisible(lines_visible int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lines_visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetMaxLinesVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetMaxLinesVisible() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetMaxLinesVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Label) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetStructuredTextBidiOverride() TextServerStructuredTextParser {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerStructuredTextParser

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetStructuredTextBidiOverride), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Label) SetStructuredTextBidiOverrideOptions(args Array) {
	cargs := []gdc.ConstTypePtr{args.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnSetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Label) GetStructuredTextBidiOverrideOptions() Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetStructuredTextBidiOverrideOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Label) GetCharacterBounds(pos int64) Rect2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pos)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()
	pinner.Pin(&pos)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForLabel.fnGetCharacterBounds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
