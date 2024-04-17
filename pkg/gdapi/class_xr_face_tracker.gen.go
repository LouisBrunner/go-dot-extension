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

type ptrsForXRFaceTrackerList struct {
	fnGetBlendShape  gdc.MethodBindPtr
	fnSetBlendShape  gdc.MethodBindPtr
	fnGetBlendShapes gdc.MethodBindPtr
	fnSetBlendShapes gdc.MethodBindPtr
}

var ptrsForXRFaceTracker ptrsForXRFaceTrackerList

func initXRFaceTrackerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("XRFaceTracker")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_blend_shape")
		defer methodName.Destroy()
		ptrsForXRFaceTracker.fnGetBlendShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 330010046))
	}
	{
		methodName := StringNameFromStr("set_blend_shape")
		defer methodName.Destroy()
		ptrsForXRFaceTracker.fnSetBlendShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2352588791))
	}
	{
		methodName := StringNameFromStr("get_blend_shapes")
		defer methodName.Destroy()
		ptrsForXRFaceTracker.fnGetBlendShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 675695659))
	}
	{
		methodName := StringNameFromStr("set_blend_shapes")
		defer methodName.Destroy()
		ptrsForXRFaceTracker.fnSetBlendShapes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2899603908))
	}

}

type XRFaceTracker struct {
	RefCounted
}

func (me *XRFaceTracker) BaseClass() string {
	return "XRFaceTracker"
}

func NewXRFaceTracker() *XRFaceTracker {
	str := StringNameFromStr("XRFaceTracker") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &XRFaceTracker{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type XRFaceTrackerBlendShapeEntry int

const (
	XRFaceTrackerBlendShapeEntryFtEyeLookOutRight       XRFaceTrackerBlendShapeEntry = 0
	XRFaceTrackerBlendShapeEntryFtEyeLookInRight        XRFaceTrackerBlendShapeEntry = 1
	XRFaceTrackerBlendShapeEntryFtEyeLookUpRight        XRFaceTrackerBlendShapeEntry = 2
	XRFaceTrackerBlendShapeEntryFtEyeLookDownRight      XRFaceTrackerBlendShapeEntry = 3
	XRFaceTrackerBlendShapeEntryFtEyeLookOutLeft        XRFaceTrackerBlendShapeEntry = 4
	XRFaceTrackerBlendShapeEntryFtEyeLookInLeft         XRFaceTrackerBlendShapeEntry = 5
	XRFaceTrackerBlendShapeEntryFtEyeLookUpLeft         XRFaceTrackerBlendShapeEntry = 6
	XRFaceTrackerBlendShapeEntryFtEyeLookDownLeft       XRFaceTrackerBlendShapeEntry = 7
	XRFaceTrackerBlendShapeEntryFtEyeClosedRight        XRFaceTrackerBlendShapeEntry = 8
	XRFaceTrackerBlendShapeEntryFtEyeClosedLeft         XRFaceTrackerBlendShapeEntry = 9
	XRFaceTrackerBlendShapeEntryFtEyeSquintRight        XRFaceTrackerBlendShapeEntry = 10
	XRFaceTrackerBlendShapeEntryFtEyeSquintLeft         XRFaceTrackerBlendShapeEntry = 11
	XRFaceTrackerBlendShapeEntryFtEyeWideRight          XRFaceTrackerBlendShapeEntry = 12
	XRFaceTrackerBlendShapeEntryFtEyeWideLeft           XRFaceTrackerBlendShapeEntry = 13
	XRFaceTrackerBlendShapeEntryFtEyeDilationRight      XRFaceTrackerBlendShapeEntry = 14
	XRFaceTrackerBlendShapeEntryFtEyeDilationLeft       XRFaceTrackerBlendShapeEntry = 15
	XRFaceTrackerBlendShapeEntryFtEyeConstrictRight     XRFaceTrackerBlendShapeEntry = 16
	XRFaceTrackerBlendShapeEntryFtEyeConstrictLeft      XRFaceTrackerBlendShapeEntry = 17
	XRFaceTrackerBlendShapeEntryFtBrowPinchRight        XRFaceTrackerBlendShapeEntry = 18
	XRFaceTrackerBlendShapeEntryFtBrowPinchLeft         XRFaceTrackerBlendShapeEntry = 19
	XRFaceTrackerBlendShapeEntryFtBrowLowererRight      XRFaceTrackerBlendShapeEntry = 20
	XRFaceTrackerBlendShapeEntryFtBrowLowererLeft       XRFaceTrackerBlendShapeEntry = 21
	XRFaceTrackerBlendShapeEntryFtBrowInnerUpRight      XRFaceTrackerBlendShapeEntry = 22
	XRFaceTrackerBlendShapeEntryFtBrowInnerUpLeft       XRFaceTrackerBlendShapeEntry = 23
	XRFaceTrackerBlendShapeEntryFtBrowOuterUpRight      XRFaceTrackerBlendShapeEntry = 24
	XRFaceTrackerBlendShapeEntryFtBrowOuterUpLeft       XRFaceTrackerBlendShapeEntry = 25
	XRFaceTrackerBlendShapeEntryFtNoseSneerRight        XRFaceTrackerBlendShapeEntry = 26
	XRFaceTrackerBlendShapeEntryFtNoseSneerLeft         XRFaceTrackerBlendShapeEntry = 27
	XRFaceTrackerBlendShapeEntryFtNasalDilationRight    XRFaceTrackerBlendShapeEntry = 28
	XRFaceTrackerBlendShapeEntryFtNasalDilationLeft     XRFaceTrackerBlendShapeEntry = 29
	XRFaceTrackerBlendShapeEntryFtNasalConstrictRight   XRFaceTrackerBlendShapeEntry = 30
	XRFaceTrackerBlendShapeEntryFtNasalConstrictLeft    XRFaceTrackerBlendShapeEntry = 31
	XRFaceTrackerBlendShapeEntryFtCheekSquintRight      XRFaceTrackerBlendShapeEntry = 32
	XRFaceTrackerBlendShapeEntryFtCheekSquintLeft       XRFaceTrackerBlendShapeEntry = 33
	XRFaceTrackerBlendShapeEntryFtCheekPuffRight        XRFaceTrackerBlendShapeEntry = 34
	XRFaceTrackerBlendShapeEntryFtCheekPuffLeft         XRFaceTrackerBlendShapeEntry = 35
	XRFaceTrackerBlendShapeEntryFtCheekSuckRight        XRFaceTrackerBlendShapeEntry = 36
	XRFaceTrackerBlendShapeEntryFtCheekSuckLeft         XRFaceTrackerBlendShapeEntry = 37
	XRFaceTrackerBlendShapeEntryFtJawOpen               XRFaceTrackerBlendShapeEntry = 38
	XRFaceTrackerBlendShapeEntryFtMouthClosed           XRFaceTrackerBlendShapeEntry = 39
	XRFaceTrackerBlendShapeEntryFtJawRight              XRFaceTrackerBlendShapeEntry = 40
	XRFaceTrackerBlendShapeEntryFtJawLeft               XRFaceTrackerBlendShapeEntry = 41
	XRFaceTrackerBlendShapeEntryFtJawForward            XRFaceTrackerBlendShapeEntry = 42
	XRFaceTrackerBlendShapeEntryFtJawBackward           XRFaceTrackerBlendShapeEntry = 43
	XRFaceTrackerBlendShapeEntryFtJawClench             XRFaceTrackerBlendShapeEntry = 44
	XRFaceTrackerBlendShapeEntryFtJawMandibleRaise      XRFaceTrackerBlendShapeEntry = 45
	XRFaceTrackerBlendShapeEntryFtLipSuckUpperRight     XRFaceTrackerBlendShapeEntry = 46
	XRFaceTrackerBlendShapeEntryFtLipSuckUpperLeft      XRFaceTrackerBlendShapeEntry = 47
	XRFaceTrackerBlendShapeEntryFtLipSuckLowerRight     XRFaceTrackerBlendShapeEntry = 48
	XRFaceTrackerBlendShapeEntryFtLipSuckLowerLeft      XRFaceTrackerBlendShapeEntry = 49
	XRFaceTrackerBlendShapeEntryFtLipSuckCornerRight    XRFaceTrackerBlendShapeEntry = 50
	XRFaceTrackerBlendShapeEntryFtLipSuckCornerLeft     XRFaceTrackerBlendShapeEntry = 51
	XRFaceTrackerBlendShapeEntryFtLipFunnelUpperRight   XRFaceTrackerBlendShapeEntry = 52
	XRFaceTrackerBlendShapeEntryFtLipFunnelUpperLeft    XRFaceTrackerBlendShapeEntry = 53
	XRFaceTrackerBlendShapeEntryFtLipFunnelLowerRight   XRFaceTrackerBlendShapeEntry = 54
	XRFaceTrackerBlendShapeEntryFtLipFunnelLowerLeft    XRFaceTrackerBlendShapeEntry = 55
	XRFaceTrackerBlendShapeEntryFtLipPuckerUpperRight   XRFaceTrackerBlendShapeEntry = 56
	XRFaceTrackerBlendShapeEntryFtLipPuckerUpperLeft    XRFaceTrackerBlendShapeEntry = 57
	XRFaceTrackerBlendShapeEntryFtLipPuckerLowerRight   XRFaceTrackerBlendShapeEntry = 58
	XRFaceTrackerBlendShapeEntryFtLipPuckerLowerLeft    XRFaceTrackerBlendShapeEntry = 59
	XRFaceTrackerBlendShapeEntryFtMouthUpperUpRight     XRFaceTrackerBlendShapeEntry = 60
	XRFaceTrackerBlendShapeEntryFtMouthUpperUpLeft      XRFaceTrackerBlendShapeEntry = 61
	XRFaceTrackerBlendShapeEntryFtMouthLowerDownRight   XRFaceTrackerBlendShapeEntry = 62
	XRFaceTrackerBlendShapeEntryFtMouthLowerDownLeft    XRFaceTrackerBlendShapeEntry = 63
	XRFaceTrackerBlendShapeEntryFtMouthUpperDeepenRight XRFaceTrackerBlendShapeEntry = 64
	XRFaceTrackerBlendShapeEntryFtMouthUpperDeepenLeft  XRFaceTrackerBlendShapeEntry = 65
	XRFaceTrackerBlendShapeEntryFtMouthUpperRight       XRFaceTrackerBlendShapeEntry = 66
	XRFaceTrackerBlendShapeEntryFtMouthUpperLeft        XRFaceTrackerBlendShapeEntry = 67
	XRFaceTrackerBlendShapeEntryFtMouthLowerRight       XRFaceTrackerBlendShapeEntry = 68
	XRFaceTrackerBlendShapeEntryFtMouthLowerLeft        XRFaceTrackerBlendShapeEntry = 69
	XRFaceTrackerBlendShapeEntryFtMouthCornerPullRight  XRFaceTrackerBlendShapeEntry = 70
	XRFaceTrackerBlendShapeEntryFtMouthCornerPullLeft   XRFaceTrackerBlendShapeEntry = 71
	XRFaceTrackerBlendShapeEntryFtMouthCornerSlantRight XRFaceTrackerBlendShapeEntry = 72
	XRFaceTrackerBlendShapeEntryFtMouthCornerSlantLeft  XRFaceTrackerBlendShapeEntry = 73
	XRFaceTrackerBlendShapeEntryFtMouthFrownRight       XRFaceTrackerBlendShapeEntry = 74
	XRFaceTrackerBlendShapeEntryFtMouthFrownLeft        XRFaceTrackerBlendShapeEntry = 75
	XRFaceTrackerBlendShapeEntryFtMouthStretchRight     XRFaceTrackerBlendShapeEntry = 76
	XRFaceTrackerBlendShapeEntryFtMouthStretchLeft      XRFaceTrackerBlendShapeEntry = 77
	XRFaceTrackerBlendShapeEntryFtMouthDimpleRight      XRFaceTrackerBlendShapeEntry = 78
	XRFaceTrackerBlendShapeEntryFtMouthDimpleLeft       XRFaceTrackerBlendShapeEntry = 79
	XRFaceTrackerBlendShapeEntryFtMouthRaiserUpper      XRFaceTrackerBlendShapeEntry = 80
	XRFaceTrackerBlendShapeEntryFtMouthRaiserLower      XRFaceTrackerBlendShapeEntry = 81
	XRFaceTrackerBlendShapeEntryFtMouthPressRight       XRFaceTrackerBlendShapeEntry = 82
	XRFaceTrackerBlendShapeEntryFtMouthPressLeft        XRFaceTrackerBlendShapeEntry = 83
	XRFaceTrackerBlendShapeEntryFtMouthTightenerRight   XRFaceTrackerBlendShapeEntry = 84
	XRFaceTrackerBlendShapeEntryFtMouthTightenerLeft    XRFaceTrackerBlendShapeEntry = 85
	XRFaceTrackerBlendShapeEntryFtTongueOut             XRFaceTrackerBlendShapeEntry = 86
	XRFaceTrackerBlendShapeEntryFtTongueUp              XRFaceTrackerBlendShapeEntry = 87
	XRFaceTrackerBlendShapeEntryFtTongueDown            XRFaceTrackerBlendShapeEntry = 88
	XRFaceTrackerBlendShapeEntryFtTongueRight           XRFaceTrackerBlendShapeEntry = 89
	XRFaceTrackerBlendShapeEntryFtTongueLeft            XRFaceTrackerBlendShapeEntry = 90
	XRFaceTrackerBlendShapeEntryFtTongueRoll            XRFaceTrackerBlendShapeEntry = 91
	XRFaceTrackerBlendShapeEntryFtTongueBlendDown       XRFaceTrackerBlendShapeEntry = 92
	XRFaceTrackerBlendShapeEntryFtTongueCurlUp          XRFaceTrackerBlendShapeEntry = 93
	XRFaceTrackerBlendShapeEntryFtTongueSquish          XRFaceTrackerBlendShapeEntry = 94
	XRFaceTrackerBlendShapeEntryFtTongueFlat            XRFaceTrackerBlendShapeEntry = 95
	XRFaceTrackerBlendShapeEntryFtTongueTwistRight      XRFaceTrackerBlendShapeEntry = 96
	XRFaceTrackerBlendShapeEntryFtTongueTwistLeft       XRFaceTrackerBlendShapeEntry = 97
	XRFaceTrackerBlendShapeEntryFtSoftPalateClose       XRFaceTrackerBlendShapeEntry = 98
	XRFaceTrackerBlendShapeEntryFtThroatSwallow         XRFaceTrackerBlendShapeEntry = 99
	XRFaceTrackerBlendShapeEntryFtNeckFlexRight         XRFaceTrackerBlendShapeEntry = 100
	XRFaceTrackerBlendShapeEntryFtNeckFlexLeft          XRFaceTrackerBlendShapeEntry = 101
	XRFaceTrackerBlendShapeEntryFtEyeClosed             XRFaceTrackerBlendShapeEntry = 102
	XRFaceTrackerBlendShapeEntryFtEyeWide               XRFaceTrackerBlendShapeEntry = 103
	XRFaceTrackerBlendShapeEntryFtEyeSquint             XRFaceTrackerBlendShapeEntry = 104
	XRFaceTrackerBlendShapeEntryFtEyeDilation           XRFaceTrackerBlendShapeEntry = 105
	XRFaceTrackerBlendShapeEntryFtEyeConstrict          XRFaceTrackerBlendShapeEntry = 106
	XRFaceTrackerBlendShapeEntryFtBrowDownRight         XRFaceTrackerBlendShapeEntry = 107
	XRFaceTrackerBlendShapeEntryFtBrowDownLeft          XRFaceTrackerBlendShapeEntry = 108
	XRFaceTrackerBlendShapeEntryFtBrowDown              XRFaceTrackerBlendShapeEntry = 109
	XRFaceTrackerBlendShapeEntryFtBrowUpRight           XRFaceTrackerBlendShapeEntry = 110
	XRFaceTrackerBlendShapeEntryFtBrowUpLeft            XRFaceTrackerBlendShapeEntry = 111
	XRFaceTrackerBlendShapeEntryFtBrowUp                XRFaceTrackerBlendShapeEntry = 112
	XRFaceTrackerBlendShapeEntryFtNoseSneer             XRFaceTrackerBlendShapeEntry = 113
	XRFaceTrackerBlendShapeEntryFtNasalDilation         XRFaceTrackerBlendShapeEntry = 114
	XRFaceTrackerBlendShapeEntryFtNasalConstrict        XRFaceTrackerBlendShapeEntry = 115
	XRFaceTrackerBlendShapeEntryFtCheekPuff             XRFaceTrackerBlendShapeEntry = 116
	XRFaceTrackerBlendShapeEntryFtCheekSuck             XRFaceTrackerBlendShapeEntry = 117
	XRFaceTrackerBlendShapeEntryFtCheekSquint           XRFaceTrackerBlendShapeEntry = 118
	XRFaceTrackerBlendShapeEntryFtLipSuckUpper          XRFaceTrackerBlendShapeEntry = 119
	XRFaceTrackerBlendShapeEntryFtLipSuckLower          XRFaceTrackerBlendShapeEntry = 120
	XRFaceTrackerBlendShapeEntryFtLipSuck               XRFaceTrackerBlendShapeEntry = 121
	XRFaceTrackerBlendShapeEntryFtLipFunnelUpper        XRFaceTrackerBlendShapeEntry = 122
	XRFaceTrackerBlendShapeEntryFtLipFunnelLower        XRFaceTrackerBlendShapeEntry = 123
	XRFaceTrackerBlendShapeEntryFtLipFunnel             XRFaceTrackerBlendShapeEntry = 124
	XRFaceTrackerBlendShapeEntryFtLipPuckerUpper        XRFaceTrackerBlendShapeEntry = 125
	XRFaceTrackerBlendShapeEntryFtLipPuckerLower        XRFaceTrackerBlendShapeEntry = 126
	XRFaceTrackerBlendShapeEntryFtLipPucker             XRFaceTrackerBlendShapeEntry = 127
	XRFaceTrackerBlendShapeEntryFtMouthUpperUp          XRFaceTrackerBlendShapeEntry = 128
	XRFaceTrackerBlendShapeEntryFtMouthLowerDown        XRFaceTrackerBlendShapeEntry = 129
	XRFaceTrackerBlendShapeEntryFtMouthOpen             XRFaceTrackerBlendShapeEntry = 130
	XRFaceTrackerBlendShapeEntryFtMouthRight            XRFaceTrackerBlendShapeEntry = 131
	XRFaceTrackerBlendShapeEntryFtMouthLeft             XRFaceTrackerBlendShapeEntry = 132
	XRFaceTrackerBlendShapeEntryFtMouthSmileRight       XRFaceTrackerBlendShapeEntry = 133
	XRFaceTrackerBlendShapeEntryFtMouthSmileLeft        XRFaceTrackerBlendShapeEntry = 134
	XRFaceTrackerBlendShapeEntryFtMouthSmile            XRFaceTrackerBlendShapeEntry = 135
	XRFaceTrackerBlendShapeEntryFtMouthSadRight         XRFaceTrackerBlendShapeEntry = 136
	XRFaceTrackerBlendShapeEntryFtMouthSadLeft          XRFaceTrackerBlendShapeEntry = 137
	XRFaceTrackerBlendShapeEntryFtMouthSad              XRFaceTrackerBlendShapeEntry = 138
	XRFaceTrackerBlendShapeEntryFtMouthStretch          XRFaceTrackerBlendShapeEntry = 139
	XRFaceTrackerBlendShapeEntryFtMouthDimple           XRFaceTrackerBlendShapeEntry = 140
	XRFaceTrackerBlendShapeEntryFtMouthTightener        XRFaceTrackerBlendShapeEntry = 141
	XRFaceTrackerBlendShapeEntryFtMouthPress            XRFaceTrackerBlendShapeEntry = 142
	XRFaceTrackerBlendShapeEntryFtMax                   XRFaceTrackerBlendShapeEntry = 143
)

func (me *XRFaceTracker) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *XRFaceTracker) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *XRFaceTracker) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *XRFaceTracker) GetBlendShape(blend_shape XRFaceTrackerBlendShapeEntry) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_shape)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&blend_shape)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRFaceTracker.fnGetBlendShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *XRFaceTracker) SetBlendShape(blend_shape XRFaceTrackerBlendShapeEntry, weight float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&blend_shape), gdc.ConstTypePtr(&weight)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRFaceTracker.fnSetBlendShape), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *XRFaceTracker) GetBlendShapes() PackedFloat32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedFloat32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRFaceTracker.fnGetBlendShapes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *XRFaceTracker) SetBlendShapes(weights PackedFloat32Array) {
	cargs := []gdc.ConstTypePtr{weights.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForXRFaceTracker.fnSetBlendShapes), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
