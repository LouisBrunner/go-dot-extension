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

type ptrsForAspectRatioContainerList struct {
	fnSetRatio               gdc.MethodBindPtr
	fnGetRatio               gdc.MethodBindPtr
	fnSetStretchMode         gdc.MethodBindPtr
	fnGetStretchMode         gdc.MethodBindPtr
	fnSetAlignmentHorizontal gdc.MethodBindPtr
	fnGetAlignmentHorizontal gdc.MethodBindPtr
	fnSetAlignmentVertical   gdc.MethodBindPtr
	fnGetAlignmentVertical   gdc.MethodBindPtr
}

var ptrsForAspectRatioContainer ptrsForAspectRatioContainerList

func initAspectRatioContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AspectRatioContainer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_ratio")
		defer methodName.Destroy()
		ptrsForAspectRatioContainer.fnSetRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_ratio")
		defer methodName.Destroy()
		ptrsForAspectRatioContainer.fnGetRatio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_stretch_mode")
		defer methodName.Destroy()
		ptrsForAspectRatioContainer.fnSetStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1876743467))
	}
	{
		methodName := StringNameFromStr("get_stretch_mode")
		defer methodName.Destroy()
		ptrsForAspectRatioContainer.fnGetStretchMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3416449033))
	}
	{
		methodName := StringNameFromStr("set_alignment_horizontal")
		defer methodName.Destroy()
		ptrsForAspectRatioContainer.fnSetAlignmentHorizontal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2147829016))
	}
	{
		methodName := StringNameFromStr("get_alignment_horizontal")
		defer methodName.Destroy()
		ptrsForAspectRatioContainer.fnGetAlignmentHorizontal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3838875429))
	}
	{
		methodName := StringNameFromStr("set_alignment_vertical")
		defer methodName.Destroy()
		ptrsForAspectRatioContainer.fnSetAlignmentVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2147829016))
	}
	{
		methodName := StringNameFromStr("get_alignment_vertical")
		defer methodName.Destroy()
		ptrsForAspectRatioContainer.fnGetAlignmentVertical = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3838875429))
	}
}

type AspectRatioContainer struct {
	Container
}

func (me *AspectRatioContainer) BaseClass() string {
	return "AspectRatioContainer"
}

func NewAspectRatioContainer() *AspectRatioContainer {
	str := StringNameFromStr("AspectRatioContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AspectRatioContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AspectRatioContainerStretchMode int

const (
	AspectRatioContainerStretchModeStretchWidthControlsHeight AspectRatioContainerStretchMode = 0
	AspectRatioContainerStretchModeStretchHeightControlsWidth AspectRatioContainerStretchMode = 1
	AspectRatioContainerStretchModeStretchFit                 AspectRatioContainerStretchMode = 2
	AspectRatioContainerStretchModeStretchCover               AspectRatioContainerStretchMode = 3
)

type AspectRatioContainerAlignmentMode int

const (
	AspectRatioContainerAlignmentModeAlignmentBegin  AspectRatioContainerAlignmentMode = 0
	AspectRatioContainerAlignmentModeAlignmentCenter AspectRatioContainerAlignmentMode = 1
	AspectRatioContainerAlignmentModeAlignmentEnd    AspectRatioContainerAlignmentMode = 2
)

func (me *AspectRatioContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AspectRatioContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AspectRatioContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AspectRatioContainer) SetRatio(ratio float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAspectRatioContainer.fnSetRatio), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AspectRatioContainer) GetRatio() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAspectRatioContainer.fnGetRatio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AspectRatioContainer) SetStretchMode(stretch_mode AspectRatioContainerStretchMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stretch_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAspectRatioContainer.fnSetStretchMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AspectRatioContainer) GetStretchMode() AspectRatioContainerStretchMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AspectRatioContainerStretchMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAspectRatioContainer.fnGetStretchMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AspectRatioContainer) SetAlignmentHorizontal(alignment_horizontal AspectRatioContainerAlignmentMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment_horizontal)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAspectRatioContainer.fnSetAlignmentHorizontal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AspectRatioContainer) GetAlignmentHorizontal() AspectRatioContainerAlignmentMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AspectRatioContainerAlignmentMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAspectRatioContainer.fnGetAlignmentHorizontal), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AspectRatioContainer) SetAlignmentVertical(alignment_vertical AspectRatioContainerAlignmentMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment_vertical)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAspectRatioContainer.fnSetAlignmentVertical), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AspectRatioContainer) GetAlignmentVertical() AspectRatioContainerAlignmentMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AspectRatioContainerAlignmentMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAspectRatioContainer.fnGetAlignmentVertical), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
