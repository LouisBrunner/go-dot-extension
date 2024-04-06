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

type ptrsForStyleBoxList struct {
	fnXDraw               gdc.MethodBindPtr
	fnXGetDrawRect        gdc.MethodBindPtr
	fnXGetMinimumSize     gdc.MethodBindPtr
	fnXTestMask           gdc.MethodBindPtr
	fnGetMinimumSize      gdc.MethodBindPtr
	fnSetContentMargin    gdc.MethodBindPtr
	fnSetContentMarginAll gdc.MethodBindPtr
	fnGetContentMargin    gdc.MethodBindPtr
	fnGetMargin           gdc.MethodBindPtr
	fnGetOffset           gdc.MethodBindPtr
	fnDraw                gdc.MethodBindPtr
	fnGetCurrentItemDrawn gdc.MethodBindPtr
	fnTestMask            gdc.MethodBindPtr
}

var ptrsForStyleBox ptrsForStyleBoxList

func initStyleBoxPtrs(iface gdc.Interface) {

	className := StringNameFromStr("StyleBox")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_minimum_size")
		defer methodName.Destroy()
		ptrsForStyleBox.fnGetMinimumSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_content_margin")
		defer methodName.Destroy()
		ptrsForStyleBox.fnSetContentMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4290182280))
	}
	{
		methodName := StringNameFromStr("set_content_margin_all")
		defer methodName.Destroy()
		ptrsForStyleBox.fnSetContentMarginAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_content_margin")
		defer methodName.Destroy()
		ptrsForStyleBox.fnGetContentMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2869120046))
	}
	{
		methodName := StringNameFromStr("get_margin")
		defer methodName.Destroy()
		ptrsForStyleBox.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2869120046))
	}
	{
		methodName := StringNameFromStr("get_offset")
		defer methodName.Destroy()
		ptrsForStyleBox.fnGetOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("draw")
		defer methodName.Destroy()
		ptrsForStyleBox.fnDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2275962004))
	}
	{
		methodName := StringNameFromStr("get_current_item_drawn")
		defer methodName.Destroy()
		ptrsForStyleBox.fnGetCurrentItemDrawn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3213695180))
	}
	{
		methodName := StringNameFromStr("test_mask")
		defer methodName.Destroy()
		ptrsForStyleBox.fnTestMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3735564539))
	}

}

type StyleBox struct {
	Resource
}

func (me *StyleBox) BaseClass() string {
	return "StyleBox"
}

func NewStyleBox() *StyleBox {
	str := StringNameFromStr("StyleBox") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &StyleBox{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *StyleBox) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *StyleBox) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *StyleBox) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *StyleBox) GetMinimumSize() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBox.fnGetMinimumSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StyleBox) SetContentMargin(margin Side, offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBox.fnSetContentMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBox) SetContentMarginAll(offset float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBox.fnSetContentMarginAll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBox) GetContentMargin(margin Side) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&margin)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBox.fnGetContentMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *StyleBox) GetMargin(margin Side) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&margin)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBox.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *StyleBox) GetOffset() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBox.fnGetOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StyleBox) Draw(canvas_item RID, rect Rect2) {
	cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBox.fnDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *StyleBox) GetCurrentItemDrawn() CanvasItem {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewCanvasItem()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBox.fnGetCurrentItemDrawn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *StyleBox) TestMask(point Vector2, rect Rect2) bool {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForStyleBox.fnTestMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
