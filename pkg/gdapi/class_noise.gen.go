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

type ptrsForNoiseList struct {
	fnGetNoise1D         gdc.MethodBindPtr
	fnGetNoise2D         gdc.MethodBindPtr
	fnGetNoise2Dv        gdc.MethodBindPtr
	fnGetNoise3D         gdc.MethodBindPtr
	fnGetNoise3Dv        gdc.MethodBindPtr
	fnGetImage           gdc.MethodBindPtr
	fnGetSeamlessImage   gdc.MethodBindPtr
	fnGetImage3D         gdc.MethodBindPtr
	fnGetSeamlessImage3D gdc.MethodBindPtr
}

var ptrsForNoise ptrsForNoiseList

func initNoisePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Noise")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_noise_1d")
		defer methodName.Destroy()
		ptrsForNoise.fnGetNoise1D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3919130443))
	}
	{
		methodName := StringNameFromStr("get_noise_2d")
		defer methodName.Destroy()
		ptrsForNoise.fnGetNoise2D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2753205203))
	}
	{
		methodName := StringNameFromStr("get_noise_2dv")
		defer methodName.Destroy()
		ptrsForNoise.fnGetNoise2Dv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2276447920))
	}
	{
		methodName := StringNameFromStr("get_noise_3d")
		defer methodName.Destroy()
		ptrsForNoise.fnGetNoise3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 973811851))
	}
	{
		methodName := StringNameFromStr("get_noise_3dv")
		defer methodName.Destroy()
		ptrsForNoise.fnGetNoise3Dv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1109078154))
	}
	{
		methodName := StringNameFromStr("get_image")
		defer methodName.Destroy()
		ptrsForNoise.fnGetImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3180683109))
	}
	{
		methodName := StringNameFromStr("get_seamless_image")
		defer methodName.Destroy()
		ptrsForNoise.fnGetSeamlessImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2770743602))
	}
	{
		methodName := StringNameFromStr("get_image_3d")
		defer methodName.Destroy()
		ptrsForNoise.fnGetImage3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3977814329))
	}
	{
		methodName := StringNameFromStr("get_seamless_image_3d")
		defer methodName.Destroy()
		ptrsForNoise.fnGetSeamlessImage3D = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 451006340))
	}
}

type Noise struct {
	Resource
}

func (me *Noise) BaseClass() string {
	return "Noise"
}

func NewNoise() *Noise {
	str := StringNameFromStr("Noise") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Noise{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *Noise) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Noise) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Noise) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Noise) GetNoise1D(x float64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&x)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoise.fnGetNoise1D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Noise) GetNoise2D(x float64, y float64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&x)
	pinner.Pin(&y)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoise.fnGetNoise2D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Noise) GetNoise2Dv(v Vector2) float64 {
	cargs := []gdc.ConstTypePtr{v.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoise.fnGetNoise2Dv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Noise) GetNoise3D(x float64, y float64, z float64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&x)
	pinner.Pin(&y)
	pinner.Pin(&z)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoise.fnGetNoise3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Noise) GetNoise3Dv(v Vector3) float64 {
	cargs := []gdc.ConstTypePtr{v.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoise.fnGetNoise3Dv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Noise) GetImage(width int64, height int64, invert bool, in_3d_space bool, normalize bool) Image {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&in_3d_space), gdc.ConstTypePtr(&normalize)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImage()
	pinner.Pin(&width)
	pinner.Pin(&height)
	pinner.Pin(&invert)
	pinner.Pin(&in_3d_space)
	pinner.Pin(&normalize)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoise.fnGetImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Noise) GetSeamlessImage(width int64, height int64, invert bool, in_3d_space bool, skirt float64, normalize bool) Image {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&in_3d_space), gdc.ConstTypePtr(&skirt), gdc.ConstTypePtr(&normalize)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewImage()
	pinner.Pin(&width)
	pinner.Pin(&height)
	pinner.Pin(&invert)
	pinner.Pin(&in_3d_space)
	pinner.Pin(&skirt)
	pinner.Pin(&normalize)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoise.fnGetSeamlessImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Noise) GetImage3D(width int64, height int64, depth int64, invert bool, normalize bool) []Image {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&depth), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&normalize)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&width)
	pinner.Pin(&height)
	pinner.Pin(&depth)
	pinner.Pin(&invert)
	pinner.Pin(&normalize)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoise.fnGetImage3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Image](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Noise) GetSeamlessImage3D(width int64, height int64, depth int64, invert bool, skirt float64, normalize bool) []Image {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&depth), gdc.ConstTypePtr(&invert), gdc.ConstTypePtr(&skirt), gdc.ConstTypePtr(&normalize)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&width)
	pinner.Pin(&height)
	pinner.Pin(&depth)
	pinner.Pin(&invert)
	pinner.Pin(&skirt)
	pinner.Pin(&normalize)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNoise.fnGetSeamlessImage3D), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Image](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Signals
