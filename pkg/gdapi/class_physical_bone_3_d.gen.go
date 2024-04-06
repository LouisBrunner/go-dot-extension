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

type ptrsForPhysicalBone3DList struct {
	fnXIntegrateForces        gdc.MethodBindPtr
	fnApplyCentralImpulse     gdc.MethodBindPtr
	fnApplyImpulse            gdc.MethodBindPtr
	fnSetJointType            gdc.MethodBindPtr
	fnGetJointType            gdc.MethodBindPtr
	fnSetJointOffset          gdc.MethodBindPtr
	fnGetJointOffset          gdc.MethodBindPtr
	fnSetJointRotation        gdc.MethodBindPtr
	fnGetJointRotation        gdc.MethodBindPtr
	fnSetBodyOffset           gdc.MethodBindPtr
	fnGetBodyOffset           gdc.MethodBindPtr
	fnGetSimulatePhysics      gdc.MethodBindPtr
	fnIsSimulatingPhysics     gdc.MethodBindPtr
	fnGetBoneId               gdc.MethodBindPtr
	fnSetMass                 gdc.MethodBindPtr
	fnGetMass                 gdc.MethodBindPtr
	fnSetFriction             gdc.MethodBindPtr
	fnGetFriction             gdc.MethodBindPtr
	fnSetBounce               gdc.MethodBindPtr
	fnGetBounce               gdc.MethodBindPtr
	fnSetGravityScale         gdc.MethodBindPtr
	fnGetGravityScale         gdc.MethodBindPtr
	fnSetLinearDampMode       gdc.MethodBindPtr
	fnGetLinearDampMode       gdc.MethodBindPtr
	fnSetAngularDampMode      gdc.MethodBindPtr
	fnGetAngularDampMode      gdc.MethodBindPtr
	fnSetLinearDamp           gdc.MethodBindPtr
	fnGetLinearDamp           gdc.MethodBindPtr
	fnSetAngularDamp          gdc.MethodBindPtr
	fnGetAngularDamp          gdc.MethodBindPtr
	fnSetLinearVelocity       gdc.MethodBindPtr
	fnGetLinearVelocity       gdc.MethodBindPtr
	fnSetAngularVelocity      gdc.MethodBindPtr
	fnGetAngularVelocity      gdc.MethodBindPtr
	fnSetUseCustomIntegrator  gdc.MethodBindPtr
	fnIsUsingCustomIntegrator gdc.MethodBindPtr
	fnSetCanSleep             gdc.MethodBindPtr
	fnIsAbleToSleep           gdc.MethodBindPtr
}

var ptrsForPhysicalBone3D ptrsForPhysicalBone3DList

func initPhysicalBone3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicalBone3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("apply_central_impulse")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnApplyCentralImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("apply_impulse")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnApplyImpulse = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2754756483))
	}
	{
		methodName := StringNameFromStr("set_joint_type")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetJointType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2289552604))
	}
	{
		methodName := StringNameFromStr("get_joint_type")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetJointType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 931347320))
	}
	{
		methodName := StringNameFromStr("set_joint_offset")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetJointOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
	}
	{
		methodName := StringNameFromStr("get_joint_offset")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetJointOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("set_joint_rotation")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetJointRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_joint_rotation")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetJointRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_body_offset")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetBodyOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
	}
	{
		methodName := StringNameFromStr("get_body_offset")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetBodyOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("get_simulate_physics")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetSimulatePhysics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("is_simulating_physics")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnIsSimulatingPhysics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_bone_id")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetBoneId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_mass")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_mass")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_friction")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetFriction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_friction")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetFriction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_bounce")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetBounce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bounce")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetBounce = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_gravity_scale")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetGravityScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_gravity_scale")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetGravityScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_linear_damp_mode")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetLinearDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1244972221))
	}
	{
		methodName := StringNameFromStr("get_linear_damp_mode")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetLinearDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 205884699))
	}
	{
		methodName := StringNameFromStr("set_angular_damp_mode")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetAngularDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1244972221))
	}
	{
		methodName := StringNameFromStr("get_angular_damp_mode")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetAngularDampMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 205884699))
	}
	{
		methodName := StringNameFromStr("set_linear_damp")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_linear_damp")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetLinearDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_angular_damp")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_angular_damp")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetAngularDamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_linear_velocity")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_linear_velocity")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetLinearVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_angular_velocity")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_angular_velocity")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnGetAngularVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("set_use_custom_integrator")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetUseCustomIntegrator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_custom_integrator")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnIsUsingCustomIntegrator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_can_sleep")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnSetCanSleep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_able_to_sleep")
		defer methodName.Destroy()
		ptrsForPhysicalBone3D.fnIsAbleToSleep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type PhysicalBone3D struct {
	PhysicsBody3D
}

func (me *PhysicalBone3D) BaseClass() string {
	return "PhysicalBone3D"
}

func NewPhysicalBone3D() *PhysicalBone3D {
	str := StringNameFromStr("PhysicalBone3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicalBone3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type PhysicalBone3DDampMode int

const (
	PhysicalBone3DDampModeDampModeCombine PhysicalBone3DDampMode = 0
	PhysicalBone3DDampModeDampModeReplace PhysicalBone3DDampMode = 1
)

type PhysicalBone3DJointType int

const (
	PhysicalBone3DJointTypeJointTypeNone   PhysicalBone3DJointType = 0
	PhysicalBone3DJointTypeJointTypePin    PhysicalBone3DJointType = 1
	PhysicalBone3DJointTypeJointTypeCone   PhysicalBone3DJointType = 2
	PhysicalBone3DJointTypeJointTypeHinge  PhysicalBone3DJointType = 3
	PhysicalBone3DJointTypeJointTypeSlider PhysicalBone3DJointType = 4
	PhysicalBone3DJointTypeJointType6Dof   PhysicalBone3DJointType = 5
)

func (me *PhysicalBone3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicalBone3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicalBone3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicalBone3D) ApplyCentralImpulse(impulse Vector3) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnApplyCentralImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) ApplyImpulse(impulse Vector3, position Vector3) {
	cargs := []gdc.ConstTypePtr{impulse.AsCTypePtr(), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnApplyImpulse), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) SetJointType(joint_type PhysicalBone3DJointType) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetJointType), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetJointType() PhysicalBone3DJointType {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicalBone3DJointType

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetJointType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicalBone3D) SetJointOffset(offset Transform3D) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetJointOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetJointOffset() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetJointOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicalBone3D) SetJointRotation(euler Vector3) {
	cargs := []gdc.ConstTypePtr{euler.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetJointRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetJointRotation() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetJointRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicalBone3D) SetBodyOffset(offset Transform3D) {
	cargs := []gdc.ConstTypePtr{offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetBodyOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetBodyOffset() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetBodyOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicalBone3D) GetSimulatePhysics() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetSimulatePhysics), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) IsSimulatingPhysics() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnIsSimulatingPhysics), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) GetBoneId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetBoneId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) SetMass(mass float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetMass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetMass() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) SetFriction(friction float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&friction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetFriction), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetFriction() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetFriction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) SetBounce(bounce float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bounce)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetBounce), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetBounce() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetBounce), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) SetGravityScale(gravity_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gravity_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetGravityScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetGravityScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetGravityScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) SetLinearDampMode(linear_damp_mode PhysicalBone3DDampMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetLinearDampMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetLinearDampMode() PhysicalBone3DDampMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicalBone3DDampMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetLinearDampMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicalBone3D) SetAngularDampMode(angular_damp_mode PhysicalBone3DDampMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetAngularDampMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetAngularDampMode() PhysicalBone3DDampMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret PhysicalBone3DDampMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetAngularDampMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PhysicalBone3D) SetLinearDamp(linear_damp float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&linear_damp)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetLinearDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetLinearDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetLinearDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) SetAngularDamp(angular_damp float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angular_damp)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetAngularDamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetAngularDamp() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetAngularDamp), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) SetLinearVelocity(linear_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{linear_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetLinearVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetLinearVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetLinearVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicalBone3D) SetAngularVelocity(angular_velocity Vector3) {
	cargs := []gdc.ConstTypePtr{angular_velocity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetAngularVelocity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) GetAngularVelocity() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnGetAngularVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicalBone3D) SetUseCustomIntegrator(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetUseCustomIntegrator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) IsUsingCustomIntegrator() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnIsUsingCustomIntegrator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicalBone3D) SetCanSleep(able_to_sleep bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&able_to_sleep)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnSetCanSleep), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicalBone3D) IsAbleToSleep() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone3D.fnIsAbleToSleep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
