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

type ptrsForRootMotionViewList struct {
	fnSetAnimationPath gdc.MethodBindPtr
	fnGetAnimationPath gdc.MethodBindPtr
	fnSetColor         gdc.MethodBindPtr
	fnGetColor         gdc.MethodBindPtr
	fnSetCellSize      gdc.MethodBindPtr
	fnGetCellSize      gdc.MethodBindPtr
	fnSetRadius        gdc.MethodBindPtr
	fnGetRadius        gdc.MethodBindPtr
	fnSetZeroY         gdc.MethodBindPtr
	fnGetZeroY         gdc.MethodBindPtr
}

var ptrsForRootMotionView ptrsForRootMotionViewList

func initRootMotionViewPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RootMotionView")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_animation_path")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnSetAnimationPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_animation_path")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnGetAnimationPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_color")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnSetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
	}
	{
		methodName := StringNameFromStr("get_color")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnGetColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
	}
	{
		methodName := StringNameFromStr("set_cell_size")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnSetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_cell_size")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnGetCellSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_zero_y")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnSetZeroY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_zero_y")
		defer methodName.Destroy()
		ptrsForRootMotionView.fnGetZeroY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type RootMotionView struct {
	VisualInstance3D
}

func (me *RootMotionView) BaseClass() string {
	return "RootMotionView"
}

func NewRootMotionView() *RootMotionView {
	str := StringNameFromStr("RootMotionView") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RootMotionView{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RootMotionView) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RootMotionView) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RootMotionView) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RootMotionView) SetAnimationPath(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnSetAnimationPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RootMotionView) GetAnimationPath() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnGetAnimationPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RootMotionView) SetColor(color Color) {
	cargs := []gdc.ConstTypePtr{color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnSetColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RootMotionView) GetColor() Color {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnGetColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RootMotionView) SetCellSize(size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnSetCellSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RootMotionView) GetCellSize() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnGetCellSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RootMotionView) SetRadius(size float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RootMotionView) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RootMotionView) SetZeroY(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnSetZeroY), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RootMotionView) GetZeroY() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRootMotionView.fnGetZeroY), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
