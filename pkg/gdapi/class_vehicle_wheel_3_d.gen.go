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

type ptrsForVehicleWheel3DList struct {
	fnSetRadius               gdc.MethodBindPtr
	fnGetRadius               gdc.MethodBindPtr
	fnSetSuspensionRestLength gdc.MethodBindPtr
	fnGetSuspensionRestLength gdc.MethodBindPtr
	fnSetSuspensionTravel     gdc.MethodBindPtr
	fnGetSuspensionTravel     gdc.MethodBindPtr
	fnSetSuspensionStiffness  gdc.MethodBindPtr
	fnGetSuspensionStiffness  gdc.MethodBindPtr
	fnSetSuspensionMaxForce   gdc.MethodBindPtr
	fnGetSuspensionMaxForce   gdc.MethodBindPtr
	fnSetDampingCompression   gdc.MethodBindPtr
	fnGetDampingCompression   gdc.MethodBindPtr
	fnSetDampingRelaxation    gdc.MethodBindPtr
	fnGetDampingRelaxation    gdc.MethodBindPtr
	fnSetUseAsTraction        gdc.MethodBindPtr
	fnIsUsedAsTraction        gdc.MethodBindPtr
	fnSetUseAsSteering        gdc.MethodBindPtr
	fnIsUsedAsSteering        gdc.MethodBindPtr
	fnSetFrictionSlip         gdc.MethodBindPtr
	fnGetFrictionSlip         gdc.MethodBindPtr
	fnIsInContact             gdc.MethodBindPtr
	fnGetContactBody          gdc.MethodBindPtr
	fnSetRollInfluence        gdc.MethodBindPtr
	fnGetRollInfluence        gdc.MethodBindPtr
	fnGetSkidinfo             gdc.MethodBindPtr
	fnGetRpm                  gdc.MethodBindPtr
	fnSetEngineForce          gdc.MethodBindPtr
	fnGetEngineForce          gdc.MethodBindPtr
	fnSetBrake                gdc.MethodBindPtr
	fnGetBrake                gdc.MethodBindPtr
	fnSetSteering             gdc.MethodBindPtr
	fnGetSteering             gdc.MethodBindPtr
}

var ptrsForVehicleWheel3D ptrsForVehicleWheel3DList

func initVehicleWheel3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VehicleWheel3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_radius")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_radius")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_suspension_rest_length")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetSuspensionRestLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_suspension_rest_length")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetSuspensionRestLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_suspension_travel")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetSuspensionTravel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_suspension_travel")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetSuspensionTravel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_suspension_stiffness")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetSuspensionStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_suspension_stiffness")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetSuspensionStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_suspension_max_force")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetSuspensionMaxForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_suspension_max_force")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetSuspensionMaxForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_damping_compression")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetDampingCompression = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_damping_compression")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetDampingCompression = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_damping_relaxation")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetDampingRelaxation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_damping_relaxation")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetDampingRelaxation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_use_as_traction")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetUseAsTraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_used_as_traction")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnIsUsedAsTraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_as_steering")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetUseAsSteering = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_used_as_steering")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnIsUsedAsSteering = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_friction_slip")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetFrictionSlip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_friction_slip")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetFrictionSlip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("is_in_contact")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnIsInContact = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_contact_body")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetContactBody = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 151077316))
	}
	{
		methodName := StringNameFromStr("set_roll_influence")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetRollInfluence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_roll_influence")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetRollInfluence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_skidinfo")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetSkidinfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_rpm")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetRpm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_engine_force")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetEngineForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_engine_force")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetEngineForce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_brake")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetBrake = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_brake")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetBrake = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_steering")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnSetSteering = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_steering")
		defer methodName.Destroy()
		ptrsForVehicleWheel3D.fnGetSteering = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type VehicleWheel3D struct {
	Node3D
}

func (me *VehicleWheel3D) BaseClass() string {
	return "VehicleWheel3D"
}

func NewVehicleWheel3D() *VehicleWheel3D {
	str := StringNameFromStr("VehicleWheel3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &VehicleWheel3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *VehicleWheel3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *VehicleWheel3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *VehicleWheel3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *VehicleWheel3D) SetRadius(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetRadius() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetSuspensionRestLength(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetSuspensionRestLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetSuspensionRestLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetSuspensionRestLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetSuspensionTravel(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetSuspensionTravel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetSuspensionTravel() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetSuspensionTravel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetSuspensionStiffness(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetSuspensionStiffness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetSuspensionStiffness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetSuspensionStiffness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetSuspensionMaxForce(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetSuspensionMaxForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetSuspensionMaxForce() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetSuspensionMaxForce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetDampingCompression(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetDampingCompression), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetDampingCompression() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetDampingCompression), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetDampingRelaxation(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetDampingRelaxation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetDampingRelaxation() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetDampingRelaxation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetUseAsTraction(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetUseAsTraction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) IsUsedAsTraction() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnIsUsedAsTraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetUseAsSteering(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetUseAsSteering), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) IsUsedAsSteering() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnIsUsedAsSteering), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetFrictionSlip(length float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetFrictionSlip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetFrictionSlip() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetFrictionSlip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) IsInContact() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnIsInContact), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) GetContactBody() Node3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetContactBody), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VehicleWheel3D) SetRollInfluence(roll_influence float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&roll_influence)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetRollInfluence), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetRollInfluence() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetRollInfluence), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) GetSkidinfo() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetSkidinfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) GetRpm() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetRpm), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetEngineForce(engine_force float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&engine_force)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetEngineForce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetEngineForce() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetEngineForce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetBrake(brake float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&brake)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetBrake), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetBrake() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetBrake), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VehicleWheel3D) SetSteering(steering float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&steering)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnSetSteering), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VehicleWheel3D) GetSteering() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVehicleWheel3D.fnGetSteering), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
