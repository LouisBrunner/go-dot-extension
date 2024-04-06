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

type ptrsForCurve2DList struct {
	fnGetPointCount           gdc.MethodBindPtr
	fnSetPointCount           gdc.MethodBindPtr
	fnAddPoint                gdc.MethodBindPtr
	fnSetPointPosition        gdc.MethodBindPtr
	fnGetPointPosition        gdc.MethodBindPtr
	fnSetPointIn              gdc.MethodBindPtr
	fnGetPointIn              gdc.MethodBindPtr
	fnSetPointOut             gdc.MethodBindPtr
	fnGetPointOut             gdc.MethodBindPtr
	fnRemovePoint             gdc.MethodBindPtr
	fnClearPoints             gdc.MethodBindPtr
	fnSample                  gdc.MethodBindPtr
	fnSamplef                 gdc.MethodBindPtr
	fnSetBakeInterval         gdc.MethodBindPtr
	fnGetBakeInterval         gdc.MethodBindPtr
	fnGetBakedLength          gdc.MethodBindPtr
	fnSampleBaked             gdc.MethodBindPtr
	fnSampleBakedWithRotation gdc.MethodBindPtr
	fnGetBakedPoints          gdc.MethodBindPtr
	fnGetClosestPoint         gdc.MethodBindPtr
	fnGetClosestOffset        gdc.MethodBindPtr
	fnTessellate              gdc.MethodBindPtr
	fnTessellateEvenLength    gdc.MethodBindPtr
}

var ptrsForCurve2D ptrsForCurve2DList

func initCurve2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Curve2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_point_count")
		defer methodName.Destroy()
		ptrsForCurve2D.fnGetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_point_count")
		defer methodName.Destroy()
		ptrsForCurve2D.fnSetPointCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("add_point")
		defer methodName.Destroy()
		ptrsForCurve2D.fnAddPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4175465202))
	}
	{
		methodName := StringNameFromStr("set_point_position")
		defer methodName.Destroy()
		ptrsForCurve2D.fnSetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
	}
	{
		methodName := StringNameFromStr("get_point_position")
		defer methodName.Destroy()
		ptrsForCurve2D.fnGetPointPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("set_point_in")
		defer methodName.Destroy()
		ptrsForCurve2D.fnSetPointIn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
	}
	{
		methodName := StringNameFromStr("get_point_in")
		defer methodName.Destroy()
		ptrsForCurve2D.fnGetPointIn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("set_point_out")
		defer methodName.Destroy()
		ptrsForCurve2D.fnSetPointOut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
	}
	{
		methodName := StringNameFromStr("get_point_out")
		defer methodName.Destroy()
		ptrsForCurve2D.fnGetPointOut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("remove_point")
		defer methodName.Destroy()
		ptrsForCurve2D.fnRemovePoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("clear_points")
		defer methodName.Destroy()
		ptrsForCurve2D.fnClearPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("sample")
		defer methodName.Destroy()
		ptrsForCurve2D.fnSample = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 26514310))
	}
	{
		methodName := StringNameFromStr("samplef")
		defer methodName.Destroy()
		ptrsForCurve2D.fnSamplef = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3588506812))
	}
	{
		methodName := StringNameFromStr("set_bake_interval")
		defer methodName.Destroy()
		ptrsForCurve2D.fnSetBakeInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bake_interval")
		defer methodName.Destroy()
		ptrsForCurve2D.fnGetBakeInterval = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_baked_length")
		defer methodName.Destroy()
		ptrsForCurve2D.fnGetBakedLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("sample_baked")
		defer methodName.Destroy()
		ptrsForCurve2D.fnSampleBaked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3464257706))
	}
	{
		methodName := StringNameFromStr("sample_baked_with_rotation")
		defer methodName.Destroy()
		ptrsForCurve2D.fnSampleBakedWithRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3296056341))
	}
	{
		methodName := StringNameFromStr("get_baked_points")
		defer methodName.Destroy()
		ptrsForCurve2D.fnGetBakedPoints = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
	}
	{
		methodName := StringNameFromStr("get_closest_point")
		defer methodName.Destroy()
		ptrsForCurve2D.fnGetClosestPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2656412154))
	}
	{
		methodName := StringNameFromStr("get_closest_offset")
		defer methodName.Destroy()
		ptrsForCurve2D.fnGetClosestOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2276447920))
	}
	{
		methodName := StringNameFromStr("tessellate")
		defer methodName.Destroy()
		ptrsForCurve2D.fnTessellate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 958145977))
	}
	{
		methodName := StringNameFromStr("tessellate_even_length")
		defer methodName.Destroy()
		ptrsForCurve2D.fnTessellateEvenLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2319761637))
	}

}

type Curve2D struct {
	Resource
}

func (me *Curve2D) BaseClass() string {
	return "Curve2D"
}

func NewCurve2D() *Curve2D {
	str := StringNameFromStr("Curve2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Curve2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Curve2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Curve2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Curve2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Curve2D) GetPointCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnGetPointCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Curve2D) SetPointCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnSetPointCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Curve2D) AddPoint(position Vector2, in Vector2, out Vector2, index int64) {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), in.AsCTypePtr(), out.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnAddPoint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Curve2D) SetPointPosition(idx int64, position Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnSetPointPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Curve2D) GetPointPosition(idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnGetPointPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) SetPointIn(idx int64, position Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnSetPointIn), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Curve2D) GetPointIn(idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnGetPointIn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) SetPointOut(idx int64, position Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnSetPointOut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Curve2D) GetPointOut(idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnGetPointOut), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) RemovePoint(idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnRemovePoint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Curve2D) ClearPoints() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnClearPoints), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Curve2D) Sample(idx int64, t float64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&t)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&idx)
	pinner.Pin(&t)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnSample), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) Samplef(fofs float64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fofs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&fofs)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnSamplef), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) SetBakeInterval(distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnSetBakeInterval), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Curve2D) GetBakeInterval() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnGetBakeInterval), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Curve2D) GetBakedLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnGetBakedLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Curve2D) SampleBaked(offset float64, cubic bool) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&cubic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&offset)
	pinner.Pin(&cubic)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnSampleBaked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) SampleBakedWithRotation(offset float64, cubic bool) Transform2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&cubic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()
	pinner.Pin(&offset)
	pinner.Pin(&cubic)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnSampleBakedWithRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) GetBakedPoints() PackedVector2Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnGetBakedPoints), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) GetClosestPoint(to_point Vector2) Vector2 {
	cargs := []gdc.ConstTypePtr{to_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnGetClosestPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) GetClosestOffset(to_point Vector2) float64 {
	cargs := []gdc.ConstTypePtr{to_point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnGetClosestOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Curve2D) Tessellate(max_stages int64, tolerance_degrees float64) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_stages), gdc.ConstTypePtr(&tolerance_degrees)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()
	pinner.Pin(&max_stages)
	pinner.Pin(&tolerance_degrees)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnTessellate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Curve2D) TessellateEvenLength(max_stages int64, tolerance_length float64) PackedVector2Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_stages), gdc.ConstTypePtr(&tolerance_length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedVector2Array()
	pinner.Pin(&max_stages)
	pinner.Pin(&tolerance_length)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCurve2D.fnTessellateEvenLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
