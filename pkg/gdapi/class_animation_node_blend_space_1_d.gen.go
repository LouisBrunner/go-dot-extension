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

type ptrsForAnimationNodeBlendSpace1DList struct {
	fnAddBlendPoint         gdc.MethodBindPtr
	fnSetBlendPointPosition gdc.MethodBindPtr
	fnGetBlendPointPosition gdc.MethodBindPtr
	fnSetBlendPointNode     gdc.MethodBindPtr
	fnGetBlendPointNode     gdc.MethodBindPtr
	fnRemoveBlendPoint      gdc.MethodBindPtr
	fnGetBlendPointCount    gdc.MethodBindPtr
	fnSetMinSpace           gdc.MethodBindPtr
	fnGetMinSpace           gdc.MethodBindPtr
	fnSetMaxSpace           gdc.MethodBindPtr
	fnGetMaxSpace           gdc.MethodBindPtr
	fnSetSnap               gdc.MethodBindPtr
	fnGetSnap               gdc.MethodBindPtr
	fnSetValueLabel         gdc.MethodBindPtr
	fnGetValueLabel         gdc.MethodBindPtr
	fnSetBlendMode          gdc.MethodBindPtr
	fnGetBlendMode          gdc.MethodBindPtr
	fnSetUseSync            gdc.MethodBindPtr
	fnIsUsingSync           gdc.MethodBindPtr
}

var ptrsForAnimationNodeBlendSpace1D ptrsForAnimationNodeBlendSpace1DList

func initAnimationNodeBlendSpace1DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNodeBlendSpace1D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_blend_point")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnAddBlendPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 285050433))
	}
	{
		methodName := StringNameFromStr("set_blend_point_position")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnSetBlendPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_blend_point_position")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnGetBlendPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_blend_point_node")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnSetBlendPointNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4240341528))
	}
	{
		methodName := StringNameFromStr("get_blend_point_node")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnGetBlendPointNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 665599029))
	}
	{
		methodName := StringNameFromStr("remove_blend_point")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnRemoveBlendPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_blend_point_count")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnGetBlendPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_min_space")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnSetMinSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_min_space")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnGetMinSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_max_space")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnSetMaxSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_max_space")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnGetMaxSpace = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_snap")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnSetSnap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_snap")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnGetSnap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_value_label")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnSetValueLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_value_label")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnGetValueLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_blend_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnSetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2600869457))
	}
	{
		methodName := StringNameFromStr("get_blend_mode")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnGetBlendMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1547667849))
	}
	{
		methodName := StringNameFromStr("set_use_sync")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnSetUseSync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_sync")
		defer methodName.Destroy()
		ptrsForAnimationNodeBlendSpace1D.fnIsUsingSync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type AnimationNodeBlendSpace1D struct {
	AnimationRootNode
}

func (me *AnimationNodeBlendSpace1D) BaseClass() string {
	return "AnimationNodeBlendSpace1D"
}

func NewAnimationNodeBlendSpace1D() *AnimationNodeBlendSpace1D {
	str := StringNameFromStr("AnimationNodeBlendSpace1D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationNodeBlendSpace1D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AnimationNodeBlendSpace1DBlendMode int

const (
	AnimationNodeBlendSpace1DBlendModeBlendModeInterpolated  AnimationNodeBlendSpace1DBlendMode = 0
	AnimationNodeBlendSpace1DBlendModeBlendModeDiscrete      AnimationNodeBlendSpace1DBlendMode = 1
	AnimationNodeBlendSpace1DBlendModeBlendModeDiscreteCarry AnimationNodeBlendSpace1DBlendMode = 2
)

func (me *AnimationNodeBlendSpace1D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationNodeBlendSpace1D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeBlendSpace1D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimationNodeBlendSpace1D) AddBlendPoint(node AnimationRootNode, pos float64, at_index int64) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), gdc.ConstTypePtr(&pos), gdc.ConstTypePtr(&at_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnAddBlendPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) SetBlendPointPosition(point int64, pos float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), gdc.ConstTypePtr(&pos)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnSetBlendPointPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) GetBlendPointPosition(point int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&point)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnGetBlendPointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeBlendSpace1D) SetBlendPointNode(point int64, node AnimationRootNode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point), node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnSetBlendPointNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) GetBlendPointNode(point int64) AnimationRootNode {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAnimationRootNode()
	pinner.Pin(&point)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnGetBlendPointNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeBlendSpace1D) RemoveBlendPoint(point int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&point)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnRemoveBlendPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) GetBlendPointCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnGetBlendPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeBlendSpace1D) SetMinSpace(min_space float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_space)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnSetMinSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) GetMinSpace() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnGetMinSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeBlendSpace1D) SetMaxSpace(max_space float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_space)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnSetMaxSpace), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) GetMaxSpace() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnGetMaxSpace), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeBlendSpace1D) SetSnap(snap float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&snap)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnSetSnap), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) GetSnap() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnGetSnap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeBlendSpace1D) SetValueLabel(text String) {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnSetValueLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) GetValueLabel() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnGetValueLabel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeBlendSpace1D) SetBlendMode(mode AnimationNodeBlendSpace1DBlendMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnSetBlendMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) GetBlendMode() AnimationNodeBlendSpace1DBlendMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationNodeBlendSpace1DBlendMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnGetBlendMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationNodeBlendSpace1D) SetUseSync(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnSetUseSync), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeBlendSpace1D) IsUsingSync() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeBlendSpace1D.fnIsUsingSync), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
