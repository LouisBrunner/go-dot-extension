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

type ptrsForGeneric6DOFJoint3DList struct {
	fnSetParamX gdc.MethodBindPtr
	fnGetParamX gdc.MethodBindPtr
	fnSetParamY gdc.MethodBindPtr
	fnGetParamY gdc.MethodBindPtr
	fnSetParamZ gdc.MethodBindPtr
	fnGetParamZ gdc.MethodBindPtr
	fnSetFlagX  gdc.MethodBindPtr
	fnGetFlagX  gdc.MethodBindPtr
	fnSetFlagY  gdc.MethodBindPtr
	fnGetFlagY  gdc.MethodBindPtr
	fnSetFlagZ  gdc.MethodBindPtr
	fnGetFlagZ  gdc.MethodBindPtr
}

var ptrsForGeneric6DOFJoint3D ptrsForGeneric6DOFJoint3DList

func initGeneric6DOFJoint3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Generic6DOFJoint3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_param_x")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnSetParamX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2018184242))
	}
	{
		methodName := StringNameFromStr("get_param_x")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnGetParamX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2599835054))
	}
	{
		methodName := StringNameFromStr("set_param_y")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnSetParamY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2018184242))
	}
	{
		methodName := StringNameFromStr("get_param_y")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnGetParamY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2599835054))
	}
	{
		methodName := StringNameFromStr("set_param_z")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnSetParamZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2018184242))
	}
	{
		methodName := StringNameFromStr("get_param_z")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnGetParamZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2599835054))
	}
	{
		methodName := StringNameFromStr("set_flag_x")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnSetFlagX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2451594564))
	}
	{
		methodName := StringNameFromStr("get_flag_x")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnGetFlagX = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2122427807))
	}
	{
		methodName := StringNameFromStr("set_flag_y")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnSetFlagY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2451594564))
	}
	{
		methodName := StringNameFromStr("get_flag_y")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnGetFlagY = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2122427807))
	}
	{
		methodName := StringNameFromStr("set_flag_z")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnSetFlagZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2451594564))
	}
	{
		methodName := StringNameFromStr("get_flag_z")
		defer methodName.Destroy()
		ptrsForGeneric6DOFJoint3D.fnGetFlagZ = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2122427807))
	}
}

type Generic6DOFJoint3D struct {
	Joint3D
}

func (me *Generic6DOFJoint3D) BaseClass() string {
	return "Generic6DOFJoint3D"
}

func NewGeneric6DOFJoint3D() *Generic6DOFJoint3D {
	str := StringNameFromStr("Generic6DOFJoint3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Generic6DOFJoint3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type Generic6DOFJoint3DParam int

const (
	Generic6DOFJoint3DParamParamLinearLowerLimit              Generic6DOFJoint3DParam = 0
	Generic6DOFJoint3DParamParamLinearUpperLimit              Generic6DOFJoint3DParam = 1
	Generic6DOFJoint3DParamParamLinearLimitSoftness           Generic6DOFJoint3DParam = 2
	Generic6DOFJoint3DParamParamLinearRestitution             Generic6DOFJoint3DParam = 3
	Generic6DOFJoint3DParamParamLinearDamping                 Generic6DOFJoint3DParam = 4
	Generic6DOFJoint3DParamParamLinearMotorTargetVelocity     Generic6DOFJoint3DParam = 5
	Generic6DOFJoint3DParamParamLinearMotorForceLimit         Generic6DOFJoint3DParam = 6
	Generic6DOFJoint3DParamParamLinearSpringStiffness         Generic6DOFJoint3DParam = 7
	Generic6DOFJoint3DParamParamLinearSpringDamping           Generic6DOFJoint3DParam = 8
	Generic6DOFJoint3DParamParamLinearSpringEquilibriumPoint  Generic6DOFJoint3DParam = 9
	Generic6DOFJoint3DParamParamAngularLowerLimit             Generic6DOFJoint3DParam = 10
	Generic6DOFJoint3DParamParamAngularUpperLimit             Generic6DOFJoint3DParam = 11
	Generic6DOFJoint3DParamParamAngularLimitSoftness          Generic6DOFJoint3DParam = 12
	Generic6DOFJoint3DParamParamAngularDamping                Generic6DOFJoint3DParam = 13
	Generic6DOFJoint3DParamParamAngularRestitution            Generic6DOFJoint3DParam = 14
	Generic6DOFJoint3DParamParamAngularForceLimit             Generic6DOFJoint3DParam = 15
	Generic6DOFJoint3DParamParamAngularErp                    Generic6DOFJoint3DParam = 16
	Generic6DOFJoint3DParamParamAngularMotorTargetVelocity    Generic6DOFJoint3DParam = 17
	Generic6DOFJoint3DParamParamAngularMotorForceLimit        Generic6DOFJoint3DParam = 18
	Generic6DOFJoint3DParamParamAngularSpringStiffness        Generic6DOFJoint3DParam = 19
	Generic6DOFJoint3DParamParamAngularSpringDamping          Generic6DOFJoint3DParam = 20
	Generic6DOFJoint3DParamParamAngularSpringEquilibriumPoint Generic6DOFJoint3DParam = 21
	Generic6DOFJoint3DParamParamMax                           Generic6DOFJoint3DParam = 22
)

type Generic6DOFJoint3DFlag int

const (
	Generic6DOFJoint3DFlagFlagEnableLinearLimit   Generic6DOFJoint3DFlag = 0
	Generic6DOFJoint3DFlagFlagEnableAngularLimit  Generic6DOFJoint3DFlag = 1
	Generic6DOFJoint3DFlagFlagEnableLinearSpring  Generic6DOFJoint3DFlag = 3
	Generic6DOFJoint3DFlagFlagEnableAngularSpring Generic6DOFJoint3DFlag = 2
	Generic6DOFJoint3DFlagFlagEnableMotor         Generic6DOFJoint3DFlag = 4
	Generic6DOFJoint3DFlagFlagEnableLinearMotor   Generic6DOFJoint3DFlag = 5
	Generic6DOFJoint3DFlagFlagMax                 Generic6DOFJoint3DFlag = 6
)

func (me *Generic6DOFJoint3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Generic6DOFJoint3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Generic6DOFJoint3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *Generic6DOFJoint3D) SetParamX(param Generic6DOFJoint3DParam, value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnSetParamX), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Generic6DOFJoint3D) GetParamX(param Generic6DOFJoint3DParam) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnGetParamX), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Generic6DOFJoint3D) SetParamY(param Generic6DOFJoint3DParam, value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnSetParamY), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Generic6DOFJoint3D) GetParamY(param Generic6DOFJoint3DParam) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnGetParamY), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Generic6DOFJoint3D) SetParamZ(param Generic6DOFJoint3DParam, value float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnSetParamZ), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Generic6DOFJoint3D) GetParamZ(param Generic6DOFJoint3DParam) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&param)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&param)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnGetParamZ), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Generic6DOFJoint3D) SetFlagX(flag Generic6DOFJoint3DFlag, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnSetFlagX), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Generic6DOFJoint3D) GetFlagX(flag Generic6DOFJoint3DFlag) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnGetFlagX), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Generic6DOFJoint3D) SetFlagY(flag Generic6DOFJoint3DFlag, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnSetFlagY), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Generic6DOFJoint3D) GetFlagY(flag Generic6DOFJoint3DFlag) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnGetFlagY), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Generic6DOFJoint3D) SetFlagZ(flag Generic6DOFJoint3DFlag, value bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&value)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnSetFlagZ), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Generic6DOFJoint3D) GetFlagZ(flag Generic6DOFJoint3DFlag) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&flag)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGeneric6DOFJoint3D.fnGetFlagZ), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
