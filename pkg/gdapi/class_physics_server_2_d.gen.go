// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsServer2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer2D) BaseClass() string {
  return "PhysicsServer2D"
}



// Enums

type PhysicsServer2DSpaceParameter int
const (
  PhysicsServer2DSpaceParameterSpaceParamContactRecycleRadius PhysicsServer2DSpaceParameter = 0
  PhysicsServer2DSpaceParameterSpaceParamContactMaxSeparation PhysicsServer2DSpaceParameter = 1
  PhysicsServer2DSpaceParameterSpaceParamContactMaxAllowedPenetration PhysicsServer2DSpaceParameter = 2
  PhysicsServer2DSpaceParameterSpaceParamContactDefaultBias PhysicsServer2DSpaceParameter = 3
  PhysicsServer2DSpaceParameterSpaceParamBodyLinearVelocitySleepThreshold PhysicsServer2DSpaceParameter = 4
  PhysicsServer2DSpaceParameterSpaceParamBodyAngularVelocitySleepThreshold PhysicsServer2DSpaceParameter = 5
  PhysicsServer2DSpaceParameterSpaceParamBodyTimeToSleep PhysicsServer2DSpaceParameter = 6
  PhysicsServer2DSpaceParameterSpaceParamConstraintDefaultBias PhysicsServer2DSpaceParameter = 7
  PhysicsServer2DSpaceParameterSpaceParamSolverIterations PhysicsServer2DSpaceParameter = 8
)

type PhysicsServer2DShapeType int
const (
  PhysicsServer2DShapeTypeShapeWorldBoundary PhysicsServer2DShapeType = 0
  PhysicsServer2DShapeTypeShapeSeparationRay PhysicsServer2DShapeType = 1
  PhysicsServer2DShapeTypeShapeSegment PhysicsServer2DShapeType = 2
  PhysicsServer2DShapeTypeShapeCircle PhysicsServer2DShapeType = 3
  PhysicsServer2DShapeTypeShapeRectangle PhysicsServer2DShapeType = 4
  PhysicsServer2DShapeTypeShapeCapsule PhysicsServer2DShapeType = 5
  PhysicsServer2DShapeTypeShapeConvexPolygon PhysicsServer2DShapeType = 6
  PhysicsServer2DShapeTypeShapeConcavePolygon PhysicsServer2DShapeType = 7
  PhysicsServer2DShapeTypeShapeCustom PhysicsServer2DShapeType = 8
)

type PhysicsServer2DAreaParameter int
const (
  PhysicsServer2DAreaParameterAreaParamGravityOverrideMode PhysicsServer2DAreaParameter = 0
  PhysicsServer2DAreaParameterAreaParamGravity PhysicsServer2DAreaParameter = 1
  PhysicsServer2DAreaParameterAreaParamGravityVector PhysicsServer2DAreaParameter = 2
  PhysicsServer2DAreaParameterAreaParamGravityIsPoint PhysicsServer2DAreaParameter = 3
  PhysicsServer2DAreaParameterAreaParamGravityPointUnitDistance PhysicsServer2DAreaParameter = 4
  PhysicsServer2DAreaParameterAreaParamLinearDampOverrideMode PhysicsServer2DAreaParameter = 5
  PhysicsServer2DAreaParameterAreaParamLinearDamp PhysicsServer2DAreaParameter = 6
  PhysicsServer2DAreaParameterAreaParamAngularDampOverrideMode PhysicsServer2DAreaParameter = 7
  PhysicsServer2DAreaParameterAreaParamAngularDamp PhysicsServer2DAreaParameter = 8
  PhysicsServer2DAreaParameterAreaParamPriority PhysicsServer2DAreaParameter = 9
)

type PhysicsServer2DAreaSpaceOverrideMode int
const (
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideDisabled PhysicsServer2DAreaSpaceOverrideMode = 0
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideCombine PhysicsServer2DAreaSpaceOverrideMode = 1
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideCombineReplace PhysicsServer2DAreaSpaceOverrideMode = 2
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideReplace PhysicsServer2DAreaSpaceOverrideMode = 3
  PhysicsServer2DAreaSpaceOverrideModeAreaSpaceOverrideReplaceCombine PhysicsServer2DAreaSpaceOverrideMode = 4
)

type PhysicsServer2DBodyMode int
const (
  PhysicsServer2DBodyModeBodyModeStatic PhysicsServer2DBodyMode = 0
  PhysicsServer2DBodyModeBodyModeKinematic PhysicsServer2DBodyMode = 1
  PhysicsServer2DBodyModeBodyModeRigid PhysicsServer2DBodyMode = 2
  PhysicsServer2DBodyModeBodyModeRigidLinear PhysicsServer2DBodyMode = 3
)

type PhysicsServer2DBodyParameter int
const (
  PhysicsServer2DBodyParameterBodyParamBounce PhysicsServer2DBodyParameter = 0
  PhysicsServer2DBodyParameterBodyParamFriction PhysicsServer2DBodyParameter = 1
  PhysicsServer2DBodyParameterBodyParamMass PhysicsServer2DBodyParameter = 2
  PhysicsServer2DBodyParameterBodyParamInertia PhysicsServer2DBodyParameter = 3
  PhysicsServer2DBodyParameterBodyParamCenterOfMass PhysicsServer2DBodyParameter = 4
  PhysicsServer2DBodyParameterBodyParamGravityScale PhysicsServer2DBodyParameter = 5
  PhysicsServer2DBodyParameterBodyParamLinearDampMode PhysicsServer2DBodyParameter = 6
  PhysicsServer2DBodyParameterBodyParamAngularDampMode PhysicsServer2DBodyParameter = 7
  PhysicsServer2DBodyParameterBodyParamLinearDamp PhysicsServer2DBodyParameter = 8
  PhysicsServer2DBodyParameterBodyParamAngularDamp PhysicsServer2DBodyParameter = 9
  PhysicsServer2DBodyParameterBodyParamMax PhysicsServer2DBodyParameter = 10
)

type PhysicsServer2DBodyDampMode int
const (
  PhysicsServer2DBodyDampModeBodyDampModeCombine PhysicsServer2DBodyDampMode = 0
  PhysicsServer2DBodyDampModeBodyDampModeReplace PhysicsServer2DBodyDampMode = 1
)

type PhysicsServer2DBodyState int
const (
  PhysicsServer2DBodyStateBodyStateTransform PhysicsServer2DBodyState = 0
  PhysicsServer2DBodyStateBodyStateLinearVelocity PhysicsServer2DBodyState = 1
  PhysicsServer2DBodyStateBodyStateAngularVelocity PhysicsServer2DBodyState = 2
  PhysicsServer2DBodyStateBodyStateSleeping PhysicsServer2DBodyState = 3
  PhysicsServer2DBodyStateBodyStateCanSleep PhysicsServer2DBodyState = 4
)

type PhysicsServer2DJointType int
const (
  PhysicsServer2DJointTypeJointTypePin PhysicsServer2DJointType = 0
  PhysicsServer2DJointTypeJointTypeGroove PhysicsServer2DJointType = 1
  PhysicsServer2DJointTypeJointTypeDampedSpring PhysicsServer2DJointType = 2
  PhysicsServer2DJointTypeJointTypeMax PhysicsServer2DJointType = 3
)

type PhysicsServer2DJointParam int
const (
  PhysicsServer2DJointParamJointParamBias PhysicsServer2DJointParam = 0
  PhysicsServer2DJointParamJointParamMaxBias PhysicsServer2DJointParam = 1
  PhysicsServer2DJointParamJointParamMaxForce PhysicsServer2DJointParam = 2
)

type PhysicsServer2DPinJointParam int
const (
  PhysicsServer2DPinJointParamPinJointSoftness PhysicsServer2DPinJointParam = 0
)

type PhysicsServer2DDampedSpringParam int
const (
  PhysicsServer2DDampedSpringParamDampedSpringRestLength PhysicsServer2DDampedSpringParam = 0
  PhysicsServer2DDampedSpringParamDampedSpringStiffness PhysicsServer2DDampedSpringParam = 1
  PhysicsServer2DDampedSpringParamDampedSpringDamping PhysicsServer2DDampedSpringParam = 2
)

type PhysicsServer2DCCDMode int
const (
  PhysicsServer2DCCDModeCcdModeDisabled PhysicsServer2DCCDMode = 0
  PhysicsServer2DCCDModeCcdModeCastRay PhysicsServer2DCCDMode = 1
  PhysicsServer2DCCDModeCcdModeCastShape PhysicsServer2DCCDMode = 2
)

type PhysicsServer2DAreaBodyStatus int
const (
  PhysicsServer2DAreaBodyStatusAreaBodyAdded PhysicsServer2DAreaBodyStatus = 0
  PhysicsServer2DAreaBodyStatusAreaBodyRemoved PhysicsServer2DAreaBodyStatus = 1
)

type PhysicsServer2DProcessInfo int
const (
  PhysicsServer2DProcessInfoInfoActiveObjects PhysicsServer2DProcessInfo = 0
  PhysicsServer2DProcessInfoInfoCollisionPairs PhysicsServer2DProcessInfo = 1
  PhysicsServer2DProcessInfoInfoIslandCount PhysicsServer2DProcessInfo = 2
)

func (me *PhysicsServer2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsServer2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsServer2D) WorldBoundaryShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("world_boundary_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) SeparationRayShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("separation_ray_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) SegmentShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("segment_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) CircleShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("circle_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) RectangleShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rectangle_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) CapsuleShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("capsule_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) ConvexPolygonShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("convex_polygon_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) ConcavePolygonShapeCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("concave_polygon_shape_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) ShapeSetData(shape RID, data Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_set_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3175752987) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) ShapeGetType(shape RID, ) PhysicsServer2DShapeType {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_get_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1240598777) // FIXME: should cache?
  var ret PhysicsServer2DShapeType
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) ShapeGetData(shape RID, ) Variant {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shape_get_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4171304767) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) SpaceCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) SpaceSetActive(space RID, active bool, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) SpaceIsActive(space RID, ) bool {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) SpaceSetParam(space RID, param PhysicsServer2DSpaceParameter, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 949194586) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) SpaceGetParam(space RID, param PhysicsServer2DSpaceParameter, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 874111783) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) SpaceGetDirectState(space RID, ) PhysicsDirectSpaceState2D {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("space_get_direct_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3160173886) // FIXME: should cache?
  var ret PhysicsDirectSpaceState2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaSetSpace(area RID, space RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaGetSpace(area RID, ) RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaAddShape(area RID, shape RID, transform Transform2D, disabled bool, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_add_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 754862190) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(shape.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaSetShape(area RID, shape_idx int, shape RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2310537182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaSetShapeTransform(area RID, shape_idx int, transform Transform2D, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_shape_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 736082694) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaSetShapeDisabled(area RID, shape_idx int, disabled bool, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_shape_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2658558584) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaGetShapeCount(area RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaGetShape(area RID, shape_idx int, ) RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1066463050) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaGetShapeTransform(area RID, shape_idx int, ) Transform2D {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_shape_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1324854622) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaRemoveShape(area RID, shape_idx int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_remove_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaClearShapes(area RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_clear_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaSetCollisionLayer(area RID, layer int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaGetCollisionLayer(area RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaSetCollisionMask(area RID, mask int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaGetCollisionMask(area RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaSetParam(area RID, param PhysicsServer2DAreaParameter, value Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1257146028) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaSetTransform(area RID, transform Transform2D, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1246044741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaGetParam(area RID, param PhysicsServer2DAreaParameter, ) Variant {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3047435120) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaGetTransform(area RID, ) Transform2D {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 213527486) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaAttachObjectInstanceId(area RID, id int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_attach_object_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaGetObjectInstanceId(area RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_object_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaAttachCanvasInstanceId(area RID, id int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_attach_canvas_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaGetCanvasInstanceId(area RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_get_canvas_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) AreaSetMonitorCallback(area RID, callback Callable, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_monitor_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3379118538) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaSetAreaMonitorCallback(area RID, callback Callable, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_area_monitor_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3379118538) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) AreaSetMonitorable(area RID, monitorable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("area_set_monitorable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(area.AsCTypePtr()), gdc.ConstTypePtr(&monitorable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetSpace(body RID, space RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(space.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetSpace(body RID, ) RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetMode(body RID, mode PhysicsServer2DBodyMode, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1658067650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetMode(body RID, ) PhysicsServer2DBodyMode {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3261702585) // FIXME: should cache?
  var ret PhysicsServer2DBodyMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodyAddShape(body RID, shape RID, transform Transform2D, disabled bool, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 754862190) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(shape.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodySetShape(body RID, shape_idx int, shape RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2310537182) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodySetShapeTransform(body RID, shape_idx int, transform Transform2D, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_shape_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 736082694) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetShapeCount(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_shape_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodyGetShape(body RID, shape_idx int, ) RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1066463050) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodyGetShapeTransform(body RID, shape_idx int, ) Transform2D {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_shape_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1324854622) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodyRemoveShape(body RID, shape_idx int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_remove_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyClearShapes(body RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_clear_shapes")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodySetShapeDisabled(body RID, shape_idx int, disabled bool, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_shape_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2658558584) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodySetShapeAsOneWayCollision(body RID, shape_idx int, enable bool, margin float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_shape_as_one_way_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2556489974) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&shape_idx), gdc.ConstTypePtr(&enable), gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyAttachObjectInstanceId(body RID, id int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_attach_object_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetObjectInstanceId(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_object_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodyAttachCanvasInstanceId(body RID, id int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_attach_canvas_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetCanvasInstanceId(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_canvas_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetContinuousCollisionDetectionMode(body RID, mode PhysicsServer2DCCDMode, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_continuous_collision_detection_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1882257015) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetContinuousCollisionDetectionMode(body RID, ) PhysicsServer2DCCDMode {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_continuous_collision_detection_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2661282217) // FIXME: should cache?
  var ret PhysicsServer2DCCDMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetCollisionLayer(body RID, layer int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&layer), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetCollisionLayer(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_collision_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetCollisionMask(body RID, mask int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetCollisionMask(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetCollisionPriority(body RID, priority float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetCollisionPriority(body RID, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_collision_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetParam(body RID, param PhysicsServer2DBodyParameter, value Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2715630609) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetParam(body RID, param PhysicsServer2DBodyParameter, ) Variant {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3208033526) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodyResetMassProperties(body RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_reset_mass_properties")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodySetState(body RID, state PhysicsServer2DBodyState, value Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1706355209) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&state), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetState(body RID, state PhysicsServer2DBodyState, ) Variant {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4036367961) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&state), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodyApplyCentralImpulse(body RID, impulse Vector2, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_central_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(impulse.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyApplyTorqueImpulse(body RID, impulse float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_torque_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&impulse), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyApplyImpulse(body RID, impulse Vector2, position Vector2, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_impulse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 34330743) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(impulse.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyApplyCentralForce(body RID, force Vector2, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyApplyForce(body RID, force Vector2, position Vector2, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 34330743) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyApplyTorque(body RID, torque float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_apply_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyAddConstantCentralForce(body RID, force Vector2, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_constant_central_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyAddConstantForce(body RID, force Vector2, position Vector2, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 34330743) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyAddConstantTorque(body RID, torque float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodySetConstantForce(body RID, force Vector2, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(force.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetConstantForce(body RID, ) Vector2 {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_constant_force")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2440833711) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetConstantTorque(body RID, torque float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&torque), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetConstantTorque(body RID, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_constant_torque")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetAxisVelocity(body RID, axis_velocity Vector2, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_axis_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201125042) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(axis_velocity.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyAddCollisionException(body RID, excepted_body RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_add_collision_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(excepted_body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyRemoveCollisionException(body RID, excepted_body RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_remove_collision_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 395945892) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(excepted_body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodySetMaxContactsReported(body RID, amount int, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_max_contacts_reported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyGetMaxContactsReported(body RID, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_max_contacts_reported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetOmitForceIntegration(body RID, enable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_omit_force_integration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyIsOmittingForceIntegration(body RID, ) bool {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_is_omitting_force_integration")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodySetForceIntegrationCallback(body RID, callable Callable, userdata Variant, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_set_force_integration_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3059434249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(callable.AsCTypePtr()), gdc.ConstTypePtr(userdata.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) BodyTestMotion(body RID, parameters PhysicsTestMotionParameters2D, result PhysicsTestMotionResult2D, ) bool {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_test_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1699844009) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), gdc.ConstTypePtr(parameters.AsCTypePtr()), gdc.ConstTypePtr(result.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) BodyGetDirectState(body RID, ) PhysicsDirectBodyState2D {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("body_get_direct_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1191931871) // FIXME: should cache?
  var ret PhysicsDirectBodyState2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(body.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) JointCreate() RID {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_create")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) JointClear(joint RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) JointSetParam(joint RID, param PhysicsServer2DJointParam, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3972556514) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) JointGetParam(joint RID, param PhysicsServer2DJointParam, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4016448949) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) JointDisableCollisionsBetweenBodies(joint RID, disable bool, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_disable_collisions_between_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&disable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) JointIsDisabledCollisionsBetweenBodies(joint RID, ) bool {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_is_disabled_collisions_between_bodies")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) JointMakePin(joint RID, anchor Vector2, body_a RID, body_b RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_make_pin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2288600450) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(anchor.AsCTypePtr()), gdc.ConstTypePtr(body_a.AsCTypePtr()), gdc.ConstTypePtr(body_b.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) JointMakeGroove(joint RID, groove1_a Vector2, groove2_a Vector2, anchor_b Vector2, body_a RID, body_b RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_make_groove")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3573265764) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(groove1_a.AsCTypePtr()), gdc.ConstTypePtr(groove2_a.AsCTypePtr()), gdc.ConstTypePtr(anchor_b.AsCTypePtr()), gdc.ConstTypePtr(body_a.AsCTypePtr()), gdc.ConstTypePtr(body_b.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) JointMakeDampedSpring(joint RID, anchor_a Vector2, anchor_b Vector2, body_a RID, body_b RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_make_damped_spring")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 206603952) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(anchor_a.AsCTypePtr()), gdc.ConstTypePtr(anchor_b.AsCTypePtr()), gdc.ConstTypePtr(body_a.AsCTypePtr()), gdc.ConstTypePtr(body_b.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) PinJointSetParam(joint RID, param PhysicsServer2DPinJointParam, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pin_joint_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 550574241) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) PinJointGetParam(joint RID, param PhysicsServer2DPinJointParam, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("pin_joint_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 348281383) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) DampedSpringJointSetParam(joint RID, param PhysicsServer2DDampedSpringParam, value float32, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("damped_spring_joint_set_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 220564071) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) DampedSpringJointGetParam(joint RID, param PhysicsServer2DDampedSpringParam, ) float32 {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("damped_spring_joint_get_param")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2075871277) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), gdc.ConstTypePtr(&param), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) JointGetType(joint RID, ) PhysicsServer2DJointType {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("joint_get_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4262502231) // FIXME: should cache?
  var ret PhysicsServer2DJointType
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(joint.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsServer2D) FreeRid(rid RID, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("free_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) SetActive(active bool, )  {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicsServer2D) GetProcessInfo(process_info PhysicsServer2DProcessInfo, ) int {
  classNameV := StringNameFromStr("PhysicsServer2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 576496006) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&process_info), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
