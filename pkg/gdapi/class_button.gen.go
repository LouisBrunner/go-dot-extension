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

type ptrsForButtonList struct {
	fnSetText                  gdc.MethodBindPtr
	fnGetText                  gdc.MethodBindPtr
	fnSetTextOverrunBehavior   gdc.MethodBindPtr
	fnGetTextOverrunBehavior   gdc.MethodBindPtr
	fnSetTextDirection         gdc.MethodBindPtr
	fnGetTextDirection         gdc.MethodBindPtr
	fnSetLanguage              gdc.MethodBindPtr
	fnGetLanguage              gdc.MethodBindPtr
	fnSetButtonIcon            gdc.MethodBindPtr
	fnGetButtonIcon            gdc.MethodBindPtr
	fnSetFlat                  gdc.MethodBindPtr
	fnIsFlat                   gdc.MethodBindPtr
	fnSetClipText              gdc.MethodBindPtr
	fnGetClipText              gdc.MethodBindPtr
	fnSetTextAlignment         gdc.MethodBindPtr
	fnGetTextAlignment         gdc.MethodBindPtr
	fnSetIconAlignment         gdc.MethodBindPtr
	fnGetIconAlignment         gdc.MethodBindPtr
	fnSetVerticalIconAlignment gdc.MethodBindPtr
	fnGetVerticalIconAlignment gdc.MethodBindPtr
	fnSetExpandIcon            gdc.MethodBindPtr
	fnIsExpandIcon             gdc.MethodBindPtr
}

var ptrsForButton ptrsForButtonList

func initButtonPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Button")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_text")
		defer methodName.Destroy()
		ptrsForButton.fnSetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_text")
		defer methodName.Destroy()
		ptrsForButton.fnGetText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForButton.fnSetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1008890932))
	}
	{
		methodName := StringNameFromStr("get_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForButton.fnGetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3779142101))
	}
	{
		methodName := StringNameFromStr("set_text_direction")
		defer methodName.Destroy()
		ptrsForButton.fnSetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 119160795))
	}
	{
		methodName := StringNameFromStr("get_text_direction")
		defer methodName.Destroy()
		ptrsForButton.fnGetTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797257663))
	}
	{
		methodName := StringNameFromStr("set_language")
		defer methodName.Destroy()
		ptrsForButton.fnSetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_language")
		defer methodName.Destroy()
		ptrsForButton.fnGetLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_button_icon")
		defer methodName.Destroy()
		ptrsForButton.fnSetButtonIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4051416890))
	}
	{
		methodName := StringNameFromStr("get_button_icon")
		defer methodName.Destroy()
		ptrsForButton.fnGetButtonIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3635182373))
	}
	{
		methodName := StringNameFromStr("set_flat")
		defer methodName.Destroy()
		ptrsForButton.fnSetFlat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_flat")
		defer methodName.Destroy()
		ptrsForButton.fnIsFlat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_clip_text")
		defer methodName.Destroy()
		ptrsForButton.fnSetClipText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_clip_text")
		defer methodName.Destroy()
		ptrsForButton.fnGetClipText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_text_alignment")
		defer methodName.Destroy()
		ptrsForButton.fnSetTextAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312603777))
	}
	{
		methodName := StringNameFromStr("get_text_alignment")
		defer methodName.Destroy()
		ptrsForButton.fnGetTextAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 341400642))
	}
	{
		methodName := StringNameFromStr("set_icon_alignment")
		defer methodName.Destroy()
		ptrsForButton.fnSetIconAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2312603777))
	}
	{
		methodName := StringNameFromStr("get_icon_alignment")
		defer methodName.Destroy()
		ptrsForButton.fnGetIconAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 341400642))
	}
	{
		methodName := StringNameFromStr("set_vertical_icon_alignment")
		defer methodName.Destroy()
		ptrsForButton.fnSetVerticalIconAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1796458609))
	}
	{
		methodName := StringNameFromStr("get_vertical_icon_alignment")
		defer methodName.Destroy()
		ptrsForButton.fnGetVerticalIconAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3274884059))
	}
	{
		methodName := StringNameFromStr("set_expand_icon")
		defer methodName.Destroy()
		ptrsForButton.fnSetExpandIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_expand_icon")
		defer methodName.Destroy()
		ptrsForButton.fnIsExpandIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type Button struct {
	BaseButton
}

func (me *Button) BaseClass() string {
	return "Button"
}

func NewButton() *Button {
	str := StringNameFromStr("Button") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Button{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Button) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Button) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Button) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Button) SetText(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) GetText() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnGetText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Button) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) GetTextOverrunBehavior() TextServerOverrunBehavior {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerOverrunBehavior

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnGetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Button) SetTextDirection(direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) GetTextDirection() ControlTextDirection {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnGetTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Button) SetLanguage(language String) {
	cargs := []gdc.ConstTypePtr{language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) GetLanguage() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnGetLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Button) SetButtonIcon(texture Texture2D) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetButtonIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) GetButtonIcon() Texture2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnGetButtonIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Button) SetFlat(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetFlat), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) IsFlat() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnIsFlat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Button) SetClipText(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetClipText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) GetClipText() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnGetClipText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Button) SetTextAlignment(alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetTextAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) GetTextAlignment() HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnGetTextAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Button) SetIconAlignment(icon_alignment HorizontalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&icon_alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetIconAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) GetIconAlignment() HorizontalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret HorizontalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnGetIconAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Button) SetVerticalIconAlignment(vertical_icon_alignment VerticalAlignment) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertical_icon_alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetVerticalIconAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) GetVerticalIconAlignment() VerticalAlignment {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret VerticalAlignment

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnGetVerticalIconAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Button) SetExpandIcon(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnSetExpandIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Button) IsExpandIcon() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButton.fnIsExpandIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
