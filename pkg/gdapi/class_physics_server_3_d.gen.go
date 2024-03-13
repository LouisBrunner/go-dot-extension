// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsServer3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer3D) BaseClass() string {
  return "PhysicsServer3D"
}



// Enums

type PhysicsServer3DJointType int
const (
  PhysicsServer3DJointTypeJointTypePin PhysicsServer3DJointType = 0
  PhysicsServer3DJointTypeJointTypeHinge PhysicsServer3DJointType = 1
  PhysicsServer3DJointTypeJointTypeSlider PhysicsServer3DJointType = 2
  PhysicsServer3DJointTypeJointTypeConeTwist PhysicsServer3DJointType = 3
  PhysicsServer3DJointTypeJointType6Dof PhysicsServer3DJointType = 4
  PhysicsServer3DJointTypeJointTypeMax PhysicsServer3DJointType = 5
)

type PhysicsServer3DPinJointParam int
const (
  PhysicsServer3DPinJointParamPinJointBias PhysicsServer3DPinJointParam = 0
  PhysicsServer3DPinJointParamPinJointDamping PhysicsServer3DPinJointParam = 1
  PhysicsServer3DPinJointParamPinJointImpulseClamp PhysicsServer3DPinJointParam = 2
)

type PhysicsServer3DHingeJointParam int
const (
  PhysicsServer3DHingeJointParamHingeJointBias PhysicsServer3DHingeJointParam = 0
  PhysicsServer3DHingeJointParamHingeJointLimitUpper PhysicsServer3DHingeJointParam = 1
  PhysicsServer3DHingeJointParamHingeJointLimitLower PhysicsServer3DHingeJointParam = 2
  PhysicsServer3DHingeJointParamHingeJointLimitBias PhysicsServer3DHingeJointParam = 3
  PhysicsServer3DHingeJointParamHingeJointLimitSoftness PhysicsServer3DHingeJointParam = 4
  PhysicsServer3DHingeJointParamHingeJointLimitRelaxation PhysicsServer3DHingeJointParam = 5
  PhysicsServer3DHingeJointParamHingeJointMotorTargetVelocity PhysicsServer3DHingeJointParam = 6
  PhysicsServer3DHingeJointParamHingeJointMotorMaxImpulse PhysicsServer3DHingeJointParam = 7
)

type PhysicsServer3DHingeJointFlag int
const (
  PhysicsServer3DHingeJointFlagHingeJointFlagUseLimit PhysicsServer3DHingeJointFlag = 0
  PhysicsServer3DHingeJointFlagHingeJointFlagEnableMotor PhysicsServer3DHingeJointFlag = 1
)

type PhysicsServer3DSliderJointParam int
const (
  PhysicsServer3DSliderJointParamSliderJointLinearLimitUpper PhysicsServer3DSliderJointParam = 0
  PhysicsServer3DSliderJointParamSliderJointLinearLimitLower PhysicsServer3DSliderJointParam = 1
  PhysicsServer3DSliderJointParamSliderJointLinearLimitSoftness PhysicsServer3DSliderJointParam = 2
  PhysicsServer3DSliderJointParamSliderJointLinearLimitRestitution PhysicsServer3DSliderJointParam = 3
  PhysicsServer3DSliderJointParamSliderJointLinearLimitDamping PhysicsServer3DSliderJointParam = 4
  PhysicsServer3DSliderJointParamSliderJointLinearMotionSoftness PhysicsServer3DSliderJointParam = 5
  PhysicsServer3DSliderJointParamSliderJointLinearMotionRestitution PhysicsServer3DSliderJointParam = 6
  PhysicsServer3DSliderJointParamSliderJointLinearMotionDamping PhysicsServer3DSliderJointParam = 7
  PhysicsServer3DSliderJointParamSliderJointLinearOrthogonalSoftness PhysicsServer3DSliderJointParam = 8
  PhysicsServer3DSliderJointParamSliderJointLinearOrthogonalRestitution PhysicsServer3DSliderJointParam = 9
  PhysicsServer3DSliderJointParamSliderJointLinearOrthogonalDamping PhysicsServer3DSliderJointParam = 10
  PhysicsServer3DSliderJointParamSliderJointAngularLimitUpper PhysicsServer3DSliderJointParam = 11
  PhysicsServer3DSliderJointParamSliderJointAngularLimitLower PhysicsServer3DSliderJointParam = 12
  PhysicsServer3DSliderJointParamSliderJointAngularLimitSoftness PhysicsServer3DSliderJointParam = 13
  PhysicsServer3DSliderJointParamSliderJointAngularLimitRestitution PhysicsServer3DSliderJointParam = 14
  PhysicsServer3DSliderJointParamSliderJointAngularLimitDamping PhysicsServer3DSliderJointParam = 15
  PhysicsServer3DSliderJointParamSliderJointAngularMotionSoftness PhysicsServer3DSliderJointParam = 16
  PhysicsServer3DSliderJointParamSliderJointAngularMotionRestitution PhysicsServer3DSliderJointParam = 17
  PhysicsServer3DSliderJointParamSliderJointAngularMotionDamping PhysicsServer3DSliderJointParam = 18
  PhysicsServer3DSliderJointParamSliderJointAngularOrthogonalSoftness PhysicsServer3DSliderJointParam = 19
  PhysicsServer3DSliderJointParamSliderJointAngularOrthogonalRestitution PhysicsServer3DSliderJointParam = 20
  PhysicsServer3DSliderJointParamSliderJointAngularOrthogonalDamping PhysicsServer3DSliderJointParam = 21
  PhysicsServer3DSliderJointParamSliderJointMax PhysicsServer3DSliderJointParam = 22
)

type PhysicsServer3DConeTwistJointParam int
const (
  PhysicsServer3DConeTwistJointParamConeTwistJointSwingSpan PhysicsServer3DConeTwistJointParam = 0
  PhysicsServer3DConeTwistJointParamConeTwistJointTwistSpan PhysicsServer3DConeTwistJointParam = 1
  PhysicsServer3DConeTwistJointParamConeTwistJointBias PhysicsServer3DConeTwistJointParam = 2
  PhysicsServer3DConeTwistJointParamConeTwistJointSoftness PhysicsServer3DConeTwistJointParam = 3
  PhysicsServer3DConeTwistJointParamConeTwistJointRelaxation PhysicsServer3DConeTwistJointParam = 4
)

type PhysicsServer3DG6DOFJointAxisParam int
const (
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearLowerLimit PhysicsServer3DG6DOFJointAxisParam = 0
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearUpperLimit PhysicsServer3DG6DOFJointAxisParam = 1
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearLimitSoftness PhysicsServer3DG6DOFJointAxisParam = 2
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearRestitution PhysicsServer3DG6DOFJointAxisParam = 3
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearDamping PhysicsServer3DG6DOFJointAxisParam = 4
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearMotorTargetVelocity PhysicsServer3DG6DOFJointAxisParam = 5
  PhysicsServer3DG6DOFJointAxisParamG6DofJointLinearMotorForceLimit PhysicsServer3DG6DOFJointAxisParam = 6
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularLowerLimit PhysicsServer3DG6DOFJointAxisParam = 10
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularUpperLimit PhysicsServer3DG6DOFJointAxisParam = 11
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularLimitSoftness PhysicsServer3DG6DOFJointAxisParam = 12
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularDamping PhysicsServer3DG6DOFJointAxisParam = 13
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularRestitution PhysicsServer3DG6DOFJointAxisParam = 14
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularForceLimit PhysicsServer3DG6DOFJointAxisParam = 15
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularErp PhysicsServer3DG6DOFJointAxisParam = 16
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularMotorTargetVelocity PhysicsServer3DG6DOFJointAxisParam = 17
  PhysicsServer3DG6DOFJointAxisParamG6DofJointAngularMotorForceLimit PhysicsServer3DG6DOFJointAxisParam = 18
)

type PhysicsServer3DG6DOFJointAxisFlag int
const (
  PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableLinearLimit PhysicsServer3DG6DOFJointAxisFlag = 0
  PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableAngularLimit PhysicsServer3DG6DOFJointAxisFlag = 1
  PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableMotor PhysicsServer3DG6DOFJointAxisFlag = 4
  PhysicsServer3DG6DOFJointAxisFlagG6DofJointFlagEnableLinearMotor PhysicsServer3DG6DOFJointAxisFlag = 5
)

type PhysicsServer3DShapeType int
const (
  PhysicsServer3DShapeTypeShapeWorldBoundary PhysicsServer3DShapeType = 0
  PhysicsServer3DShapeTypeShapeSeparationRay PhysicsServer3DShapeType = 1
  PhysicsServer3DShapeTypeShapeSphere PhysicsServer3DShapeType = 2
  PhysicsServer3DShapeTypeShapeBox PhysicsServer3DShapeType = 3
  PhysicsServer3DShapeTypeShapeCapsule PhysicsServer3DShapeType = 4
  PhysicsServer3DShapeTypeShapeCylinder PhysicsServer3DShapeType = 5
  PhysicsServer3DShapeTypeShapeConvexPolygon PhysicsServer3DShapeType = 6
  PhysicsServer3DShapeTypeShapeConcavePolygon PhysicsServer3DShapeType = 7
  PhysicsServer3DShapeTypeShapeHeightmap PhysicsServer3DShapeType = 8
  PhysicsServer3DShapeTypeShapeSoftBody PhysicsServer3DShapeType = 9
  PhysicsServer3DShapeTypeShapeCustom PhysicsServer3DShapeType = 10
)

type PhysicsServer3DAreaParameter int
const (
  PhysicsServer3DAreaParameterAreaParamGravityOverrideMode PhysicsServer3DAreaParameter = 0
  PhysicsServer3DAreaParameterAreaParamGravity PhysicsServer3DAreaParameter = 1
  PhysicsServer3DAreaParameterAreaParamGravityVector PhysicsServer3DAreaParameter = 2
  PhysicsServer3DAreaParameterAreaParamGravityIsPoint PhysicsServer3DAreaParameter = 3
  PhysicsServer3DAreaParameterAreaParamGravityPointUnitDistance PhysicsServer3DAreaParameter = 4
  PhysicsServer3DAreaParameterAreaParamLinearDampOverrideMode PhysicsServer3DAreaParameter = 5
  PhysicsServer3DAreaParameterAreaParamLinearDamp PhysicsServer3DAreaParameter = 6
  PhysicsServer3DAreaParameterAreaParamAngularDampOverrideMode PhysicsServer3DAreaParameter = 7
  PhysicsServer3DAreaParameterAreaParamAngularDamp PhysicsServer3DAreaParameter = 8
  PhysicsServer3DAreaParameterAreaParamPriority PhysicsServer3DAreaParameter = 9
  PhysicsServer3DAreaParameterAreaParamWindForceMagnitude PhysicsServer3DAreaParameter = 10
  PhysicsServer3DAreaParameterAreaParamWindSource PhysicsServer3DAreaParameter = 11
  PhysicsServer3DAreaParameterAreaParamWindDirection PhysicsServer3DAreaParameter = 12
  PhysicsServer3DAreaParameterAreaParamWindAttenuationFactor PhysicsServer3DAreaParameter = 13
)

type PhysicsServer3DAreaSpaceOverrideMode int
const (
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideDisabled PhysicsServer3DAreaSpaceOverrideMode = 0
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideCombine PhysicsServer3DAreaSpaceOverrideMode = 1
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideCombineReplace PhysicsServer3DAreaSpaceOverrideMode = 2
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideReplace PhysicsServer3DAreaSpaceOverrideMode = 3
  PhysicsServer3DAreaSpaceOverrideModeAreaSpaceOverrideReplaceCombine PhysicsServer3DAreaSpaceOverrideMode = 4
)

type PhysicsServer3DBodyMode int
const (
  PhysicsServer3DBodyModeBodyModeStatic PhysicsServer3DBodyMode = 0
  PhysicsServer3DBodyModeBodyModeKinematic PhysicsServer3DBodyMode = 1
  PhysicsServer3DBodyModeBodyModeRigid PhysicsServer3DBodyMode = 2
  PhysicsServer3DBodyModeBodyModeRigidLinear PhysicsServer3DBodyMode = 3
)

type PhysicsServer3DBodyParameter int
const (
  PhysicsServer3DBodyParameterBodyParamBounce PhysicsServer3DBodyParameter = 0
  PhysicsServer3DBodyParameterBodyParamFriction PhysicsServer3DBodyParameter = 1
  PhysicsServer3DBodyParameterBodyParamMass PhysicsServer3DBodyParameter = 2
  PhysicsServer3DBodyParameterBodyParamInertia PhysicsServer3DBodyParameter = 3
  PhysicsServer3DBodyParameterBodyParamCenterOfMass PhysicsServer3DBodyParameter = 4
  PhysicsServer3DBodyParameterBodyParamGravityScale PhysicsServer3DBodyParameter = 5
  PhysicsServer3DBodyParameterBodyParamLinearDampMode PhysicsServer3DBodyParameter = 6
  PhysicsServer3DBodyParameterBodyParamAngularDampMode PhysicsServer3DBodyParameter = 7
  PhysicsServer3DBodyParameterBodyParamLinearDamp PhysicsServer3DBodyParameter = 8
  PhysicsServer3DBodyParameterBodyParamAngularDamp PhysicsServer3DBodyParameter = 9
  PhysicsServer3DBodyParameterBodyParamMax PhysicsServer3DBodyParameter = 10
)

type PhysicsServer3DBodyDampMode int
const (
  PhysicsServer3DBodyDampModeBodyDampModeCombine PhysicsServer3DBodyDampMode = 0
  PhysicsServer3DBodyDampModeBodyDampModeReplace PhysicsServer3DBodyDampMode = 1
)

type PhysicsServer3DBodyState int
const (
  PhysicsServer3DBodyStateBodyStateTransform PhysicsServer3DBodyState = 0
  PhysicsServer3DBodyStateBodyStateLinearVelocity PhysicsServer3DBodyState = 1
  PhysicsServer3DBodyStateBodyStateAngularVelocity PhysicsServer3DBodyState = 2
  PhysicsServer3DBodyStateBodyStateSleeping PhysicsServer3DBodyState = 3
  PhysicsServer3DBodyStateBodyStateCanSleep PhysicsServer3DBodyState = 4
)

type PhysicsServer3DAreaBodyStatus int
const (
  PhysicsServer3DAreaBodyStatusAreaBodyAdded PhysicsServer3DAreaBodyStatus = 0
  PhysicsServer3DAreaBodyStatusAreaBodyRemoved PhysicsServer3DAreaBodyStatus = 1
)

type PhysicsServer3DProcessInfo int
const (
  PhysicsServer3DProcessInfoInfoActiveObjects PhysicsServer3DProcessInfo = 0
  PhysicsServer3DProcessInfoInfoCollisionPairs PhysicsServer3DProcessInfo = 1
  PhysicsServer3DProcessInfoInfoIslandCount PhysicsServer3DProcessInfo = 2
)

type PhysicsServer3DSpaceParameter int
const (
  PhysicsServer3DSpaceParameterSpaceParamContactRecycleRadius PhysicsServer3DSpaceParameter = 0
  PhysicsServer3DSpaceParameterSpaceParamContactMaxSeparation PhysicsServer3DSpaceParameter = 1
  PhysicsServer3DSpaceParameterSpaceParamContactMaxAllowedPenetration PhysicsServer3DSpaceParameter = 2
  PhysicsServer3DSpaceParameterSpaceParamContactDefaultBias PhysicsServer3DSpaceParameter = 3
  PhysicsServer3DSpaceParameterSpaceParamBodyLinearVelocitySleepThreshold PhysicsServer3DSpaceParameter = 4
  PhysicsServer3DSpaceParameterSpaceParamBodyAngularVelocitySleepThreshold PhysicsServer3DSpaceParameter = 5
  PhysicsServer3DSpaceParameterSpaceParamBodyTimeToSleep PhysicsServer3DSpaceParameter = 6
  PhysicsServer3DSpaceParameterSpaceParamSolverIterations PhysicsServer3DSpaceParameter = 7
)

type PhysicsServer3DBodyAxis int
const (
  PhysicsServer3DBodyAxisBodyAxisLinearX PhysicsServer3DBodyAxis = 1
  PhysicsServer3DBodyAxisBodyAxisLinearY PhysicsServer3DBodyAxis = 2
  PhysicsServer3DBodyAxisBodyAxisLinearZ PhysicsServer3DBodyAxis = 4
  PhysicsServer3DBodyAxisBodyAxisAngularX PhysicsServer3DBodyAxis = 8
  PhysicsServer3DBodyAxisBodyAxisAngularY PhysicsServer3DBodyAxis = 16
  PhysicsServer3DBodyAxisBodyAxisAngularZ PhysicsServer3DBodyAxis = 32
)

func (me *PhysicsServer3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsServer3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsServer3D) WorldBoundaryShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("world_boundary_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) SeparationRayShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("separation_ray_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) SphereShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sphere_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BoxShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("box_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) CapsuleShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("capsule_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) CylinderShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cylinder_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) ConvexPolygonShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convex_polygon_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) ConcavePolygonShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("concave_polygon_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) HeightmapShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("heightmap_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) CustomShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("custom_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) ShapeSetData(shape RID, data Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_set_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175752987) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) ShapeGetType(shape RID, ) PhysicsServer3DShapeType {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_get_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3418923367) // FIXME: should cache?
  var ret PhysicsServer3DShapeType
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) ShapeGetData(shape RID, ) Variant {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_get_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4171304767) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) SpaceCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) SpaceSetActive(space RID, active bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) SpaceIsActive(space RID, ) bool {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) SpaceSetParam(space RID, param PhysicsServer3DSpaceParameter, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2406017470) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) SpaceGetParam(space RID, param PhysicsServer3DSpaceParameter, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1523206731) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) SpaceGetDirectState(space RID, ) PhysicsDirectSpaceState3D {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_get_direct_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2048616813) // FIXME: should cache?
  var ret PhysicsDirectSpaceState3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaSetSpace(area RID, space RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaGetSpace(area RID, ) RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaAddShape(area RID, shape RID, transform Transform3D, disabled bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_add_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3711419014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(shape.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaSetShape(area RID, shape_idx int, shape RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2310537182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaSetShapeTransform(area RID, shape_idx int, transform Transform3D, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_shape_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675327471) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaSetShapeDisabled(area RID, shape_idx int, disabled bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_shape_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2658558584) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaGetShapeCount(area RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaGetShape(area RID, shape_idx int, ) RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1066463050) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaGetShapeTransform(area RID, shape_idx int, ) Transform3D {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_shape_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1050775521) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaRemoveShape(area RID, shape_idx int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_remove_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaClearShapes(area RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_clear_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaSetCollisionLayer(area RID, layer int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaGetCollisionLayer(area RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaSetCollisionMask(area RID, mask int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaGetCollisionMask(area RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaSetParam(area RID, param PhysicsServer3DAreaParameter, value Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2980114638) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaSetTransform(area RID, transform Transform3D, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3935195649) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaGetParam(area RID, param PhysicsServer3DAreaParameter, ) Variant {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 890056067) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaGetTransform(area RID, ) Transform3D {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1128465797) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaAttachObjectInstanceId(area RID, id int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_attach_object_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaGetObjectInstanceId(area RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_object_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) AreaSetMonitorCallback(area RID, callback Callable, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_monitor_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3379118538) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaSetAreaMonitorCallback(area RID, callback Callable, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_area_monitor_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3379118538) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaSetMonitorable(area RID, monitorable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_monitorable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&monitorable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) AreaSetRayPickable(area RID, enable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_ray_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetSpace(body RID, space RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetSpace(body RID, ) RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetMode(body RID, mode PhysicsServer3DBodyMode, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 606803466) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetMode(body RID, ) PhysicsServer3DBodyMode {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2488819728) // FIXME: should cache?
  var ret PhysicsServer3DBodyMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetCollisionLayer(body RID, layer int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetCollisionLayer(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetCollisionMask(body RID, mask int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetCollisionMask(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetCollisionPriority(body RID, priority float32, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetCollisionPriority(body RID, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodyAddShape(body RID, shape RID, transform Transform3D, disabled bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3711419014) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(shape.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodySetShape(body RID, shape_idx int, shape RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2310537182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodySetShapeTransform(body RID, shape_idx int, transform Transform3D, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_shape_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 675327471) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodySetShapeDisabled(body RID, shape_idx int, disabled bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_shape_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2658558584) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetShapeCount(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodyGetShape(body RID, shape_idx int, ) RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1066463050) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodyGetShapeTransform(body RID, shape_idx int, ) Transform3D {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_shape_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1050775521) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodyRemoveShape(body RID, shape_idx int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_remove_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyClearShapes(body RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_clear_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyAttachObjectInstanceId(body RID, id int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_attach_object_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetObjectInstanceId(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_object_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetEnableContinuousCollisionDetection(body RID, enable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_enable_continuous_collision_detection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyIsContinuousCollisionDetectionEnabled(body RID, ) bool {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_is_continuous_collision_detection_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetParam(body RID, param PhysicsServer3DBodyParameter, value Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 910941953) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetParam(body RID, param PhysicsServer3DBodyParameter, ) Variant {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3385027841) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodyResetMassProperties(body RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_reset_mass_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodySetState(body RID, state PhysicsServer3DBodyState, value Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 599977762) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&state), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetState(body RID, state PhysicsServer3DBodyState, ) Variant {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1850449534) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&state), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodyApplyCentralImpulse(body RID, impulse Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_central_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(impulse.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyApplyImpulse(body RID, impulse Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 390416203) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(impulse.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyApplyTorqueImpulse(body RID, impulse Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_torque_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(impulse.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyApplyCentralForce(body RID, force Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyApplyForce(body RID, force Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 390416203) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyApplyTorque(body RID, torque Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(torque.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyAddConstantCentralForce(body RID, force Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_constant_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyAddConstantForce(body RID, force Vector3, position Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 390416203) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyAddConstantTorque(body RID, torque Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(torque.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodySetConstantForce(body RID, force Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetConstantForce(body RID, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 531438156) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetConstantTorque(body RID, torque Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(torque.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetConstantTorque(body RID, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 531438156) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetAxisVelocity(body RID, axis_velocity Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_axis_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(axis_velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodySetAxisLock(body RID, axis PhysicsServer3DBodyAxis, lock bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_axis_lock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2020836892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&lock), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyIsAxisLocked(body RID, axis PhysicsServer3DBodyAxis, ) bool {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_is_axis_locked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 587853580) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&axis), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodyAddCollisionException(body RID, excepted_body RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_collision_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(excepted_body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyRemoveCollisionException(body RID, excepted_body RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_remove_collision_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(excepted_body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodySetMaxContactsReported(body RID, amount int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_max_contacts_reported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyGetMaxContactsReported(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_max_contacts_reported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetOmitForceIntegration(body RID, enable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_omit_force_integration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyIsOmittingForceIntegration(body RID, ) bool {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_is_omitting_force_integration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodySetForceIntegrationCallback(body RID, callable Callable, userdata Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_force_integration_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3059434249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(callable.AsCTypePtr()), gdc.ConstTypePtr(userdata.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodySetRayPickable(body RID, enable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_ray_pickable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) BodyTestMotion(body RID, parameters PhysicsTestMotionParameters3D, result PhysicsTestMotionResult3D, ) bool {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_test_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1944921792) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(parameters.AsCTypePtr()), gdc.ConstTypePtr(result.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) BodyGetDirectState(body RID, ) PhysicsDirectBodyState3D {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_direct_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3029727957) // FIXME: should cache?
  var ret PhysicsDirectBodyState3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) SoftBodyGetBounds(body RID, ) AABB {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("soft_body_get_bounds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 974181306) // FIXME: should cache?
  var ret AABB
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) JointCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) JointClear(joint RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) JointMakePin(joint RID, body_A RID, local_A Vector3, body_B RID, local_B Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_make_pin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4280171926) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(body_A.AsCTypePtr()), gdc.ConstTypePtr(local_A.AsCTypePtr()), gdc.ConstTypePtr(body_B.AsCTypePtr()), gdc.ConstTypePtr(local_B.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) PinJointSetParam(joint RID, param PhysicsServer3DPinJointParam, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pin_joint_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 810685294) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) PinJointGetParam(joint RID, param PhysicsServer3DPinJointParam, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pin_joint_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2817972347) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) PinJointSetLocalA(joint RID, local_A Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pin_joint_set_local_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(local_A.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) PinJointGetLocalA(joint RID, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pin_joint_get_local_a")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 531438156) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) PinJointSetLocalB(joint RID, local_B Vector3, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pin_joint_set_local_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3227306858) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(local_B.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) PinJointGetLocalB(joint RID, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pin_joint_get_local_b")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 531438156) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) JointMakeHinge(joint RID, body_A RID, hinge_A Transform3D, body_B RID, hinge_B Transform3D, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_make_hinge")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1684107643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(body_A.AsCTypePtr()), gdc.ConstTypePtr(hinge_A.AsCTypePtr()), gdc.ConstTypePtr(body_B.AsCTypePtr()), gdc.ConstTypePtr(hinge_B.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) HingeJointSetParam(joint RID, param PhysicsServer3DHingeJointParam, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hinge_joint_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3165502333) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) HingeJointGetParam(joint RID, param PhysicsServer3DHingeJointParam, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hinge_joint_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2129207581) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) HingeJointSetFlag(joint RID, flag PhysicsServer3DHingeJointFlag, enabled bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hinge_joint_set_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1601626188) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) HingeJointGetFlag(joint RID, flag PhysicsServer3DHingeJointFlag, ) bool {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("hinge_joint_get_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4165147865) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&flag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) JointMakeSlider(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_make_slider")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1684107643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(body_A.AsCTypePtr()), gdc.ConstTypePtr(local_ref_A.AsCTypePtr()), gdc.ConstTypePtr(body_B.AsCTypePtr()), gdc.ConstTypePtr(local_ref_B.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) SliderJointSetParam(joint RID, param PhysicsServer3DSliderJointParam, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("slider_joint_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2264833593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) SliderJointGetParam(joint RID, param PhysicsServer3DSliderJointParam, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("slider_joint_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3498644957) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) JointMakeConeTwist(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_make_cone_twist")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1684107643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(body_A.AsCTypePtr()), gdc.ConstTypePtr(local_ref_A.AsCTypePtr()), gdc.ConstTypePtr(body_B.AsCTypePtr()), gdc.ConstTypePtr(local_ref_B.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) ConeTwistJointSetParam(joint RID, param PhysicsServer3DConeTwistJointParam, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cone_twist_joint_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 808587618) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) ConeTwistJointGetParam(joint RID, param PhysicsServer3DConeTwistJointParam, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cone_twist_joint_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1134789658) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) JointGetType(joint RID, ) PhysicsServer3DJointType {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_get_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4290791900) // FIXME: should cache?
  var ret PhysicsServer3DJointType
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) JointSetSolverPriority(joint RID, priority int, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_set_solver_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) JointGetSolverPriority(joint RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_get_solver_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) JointDisableCollisionsBetweenBodies(joint RID, disable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_disable_collisions_between_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&disable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) JointIsDisabledCollisionsBetweenBodies(joint RID, ) bool {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_is_disabled_collisions_between_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) JointMakeGeneric6Dof(joint RID, body_A RID, local_ref_A Transform3D, body_B RID, local_ref_B Transform3D, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_make_generic_6dof")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1684107643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(body_A.AsCTypePtr()), gdc.ConstTypePtr(local_ref_A.AsCTypePtr()), gdc.ConstTypePtr(body_B.AsCTypePtr()), gdc.ConstTypePtr(local_ref_B.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) Generic6DofJointSetParam(joint RID, axis Vector3Axis, param PhysicsServer3DG6DOFJointAxisParam, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generic_6dof_joint_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2600081391) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) Generic6DofJointGetParam(joint RID, axis Vector3Axis, param PhysicsServer3DG6DOFJointAxisParam, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generic_6dof_joint_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 467122058) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) Generic6DofJointSetFlag(joint RID, axis Vector3Axis, flag PhysicsServer3DG6DOFJointAxisFlag, enable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generic_6dof_joint_set_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3570926903) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) Generic6DofJointGetFlag(joint RID, axis Vector3Axis, flag PhysicsServer3DG6DOFJointAxisFlag, ) bool {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generic_6dof_joint_get_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4158090196) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&axis), gdc.ConstTypePtr(&flag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer3D) FreeRid(rid RID, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("free_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) SetActive(active bool, )  {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer3D) GetProcessInfo(process_info PhysicsServer3DProcessInfo, ) int {
  classNameV := StringNameFromStr("PhysicsServer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1332958745) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&process_info), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
