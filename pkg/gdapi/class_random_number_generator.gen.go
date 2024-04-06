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

type ptrsForRandomNumberGeneratorList struct {
	fnSetSeed    gdc.MethodBindPtr
	fnGetSeed    gdc.MethodBindPtr
	fnSetState   gdc.MethodBindPtr
	fnGetState   gdc.MethodBindPtr
	fnRandi      gdc.MethodBindPtr
	fnRandf      gdc.MethodBindPtr
	fnRandfn     gdc.MethodBindPtr
	fnRandfRange gdc.MethodBindPtr
	fnRandiRange gdc.MethodBindPtr
	fnRandomize  gdc.MethodBindPtr
}

var ptrsForRandomNumberGenerator ptrsForRandomNumberGeneratorList

func initRandomNumberGeneratorPtrs(iface gdc.Interface) {

	className := StringNameFromStr("RandomNumberGenerator")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_seed")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnSetSeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_seed")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnGetSeed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_state")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnSetState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_state")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnGetState = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("randi")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnRandi = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("randf")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnRandf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
	}
	{
		methodName := StringNameFromStr("randfn")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnRandfn = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 837325100))
	}
	{
		methodName := StringNameFromStr("randf_range")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnRandfRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4269894367))
	}
	{
		methodName := StringNameFromStr("randi_range")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnRandiRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 50157827))
	}
	{
		methodName := StringNameFromStr("randomize")
		defer methodName.Destroy()
		ptrsForRandomNumberGenerator.fnRandomize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type RandomNumberGenerator struct {
	RefCounted
}

func (me *RandomNumberGenerator) BaseClass() string {
	return "RandomNumberGenerator"
}

func NewRandomNumberGenerator() *RandomNumberGenerator {
	str := StringNameFromStr("RandomNumberGenerator") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RandomNumberGenerator{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *RandomNumberGenerator) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RandomNumberGenerator) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RandomNumberGenerator) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RandomNumberGenerator) SetSeed(seed int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnSetSeed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RandomNumberGenerator) GetSeed() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnGetSeed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RandomNumberGenerator) SetState(state int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&state)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnSetState), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RandomNumberGenerator) GetState() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnGetState), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RandomNumberGenerator) Randi() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnRandi), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RandomNumberGenerator) Randf() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnRandf), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RandomNumberGenerator) Randfn(mean float64, deviation float64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mean), gdc.ConstTypePtr(&deviation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&mean)
	pinner.Pin(&deviation)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnRandfn), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RandomNumberGenerator) RandfRange(from float64, to float64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&from)
	pinner.Pin(&to)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnRandfRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RandomNumberGenerator) RandiRange(from int64, to int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&from)
	pinner.Pin(&to)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnRandiRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RandomNumberGenerator) Randomize() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRandomNumberGenerator.fnRandomize), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
