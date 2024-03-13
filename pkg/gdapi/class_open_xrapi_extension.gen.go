// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OpenXRAPIExtension struct {
  obj gdc.ObjectPtr
}

func (me *OpenXRAPIExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OpenXRAPIExtension) BaseClass() string {
  return "OpenXRAPIExtension"
}



// Enums

func (me *OpenXRAPIExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OpenXRAPIExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OpenXRAPIExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OpenXRAPIExtension) GetInstance() int {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) GetSystemId() int {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_system_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) GetSession() int {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_session")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) TransformFromPose(pose unsafe.Pointer, ) Transform3D {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("transform_from_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3255299855) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pose), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) XrResult(result int, format String, args Array, ) bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("xr_result")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3886436197) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&result), gdc.ConstTypePtr(format.AsCTypePtr()), gdc.ConstTypePtr(args.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  OpenXRAPIExtensionOpenxrIsEnabled(check_run_in_editor bool, ) bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("openxr_is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2703660260) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&check_run_in_editor), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) GetInstanceProcAddr(name String, ) int {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_proc_addr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1597066294) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) GetErrorString(result int, ) String {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_error_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 990163283) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&result), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) GetSwapchainFormatName(swapchain_format int, ) String {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swapchain_format_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 990163283) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&swapchain_format), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) IsInitialized() bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_initialized")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) IsRunning() bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_running")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) GetPlaySpace() int {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_play_space")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) GetNextFrameTime() int {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_next_frame_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OpenXRAPIExtension) CanRender() bool {
  classNameV := StringNameFromStr("OpenXRAPIExtension")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("can_render")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
