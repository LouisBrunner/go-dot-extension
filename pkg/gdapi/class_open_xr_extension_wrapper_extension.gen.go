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

type ptrsForOpenXRExtensionWrapperExtensionList struct {
	fnXGetRequestedExtensions                               gdc.MethodBindPtr
	fnXSetSystemPropertiesAndGetNextPointer                 gdc.MethodBindPtr
	fnXSetInstanceCreateInfoAndGetNextPointer               gdc.MethodBindPtr
	fnXSetSessionCreateAndGetNextPointer                    gdc.MethodBindPtr
	fnXSetSwapchainCreateInfoAndGetNextPointer              gdc.MethodBindPtr
	fnXSetHandJointLocationsAndGetNextPointer               gdc.MethodBindPtr
	fnXGetCompositionLayerCount                             gdc.MethodBindPtr
	fnXGetCompositionLayer                                  gdc.MethodBindPtr
	fnXGetCompositionLayerOrder                             gdc.MethodBindPtr
	fnXGetSuggestedTrackerNames                             gdc.MethodBindPtr
	fnXOnRegisterMetadata                                   gdc.MethodBindPtr
	fnXOnBeforeInstanceCreated                              gdc.MethodBindPtr
	fnXOnInstanceCreated                                    gdc.MethodBindPtr
	fnXOnInstanceDestroyed                                  gdc.MethodBindPtr
	fnXOnSessionCreated                                     gdc.MethodBindPtr
	fnXOnProcess                                            gdc.MethodBindPtr
	fnXOnPreRender                                          gdc.MethodBindPtr
	fnXOnSessionDestroyed                                   gdc.MethodBindPtr
	fnXOnStateIdle                                          gdc.MethodBindPtr
	fnXOnStateReady                                         gdc.MethodBindPtr
	fnXOnStateSynchronized                                  gdc.MethodBindPtr
	fnXOnStateVisible                                       gdc.MethodBindPtr
	fnXOnStateFocused                                       gdc.MethodBindPtr
	fnXOnStateStopping                                      gdc.MethodBindPtr
	fnXOnStateLossPending                                   gdc.MethodBindPtr
	fnXOnStateExiting                                       gdc.MethodBindPtr
	fnXOnEventPolled                                        gdc.MethodBindPtr
	fnXSetViewportCompositionLayerAndGetNextPointer         gdc.MethodBindPtr
	fnXGetViewportCompositionLayerExtensionProperties       gdc.MethodBindPtr
	fnXGetViewportCompositionLayerExtensionPropertyDefaults gdc.MethodBindPtr
	fnXOnViewportCompositionLayerDestroyed                  gdc.MethodBindPtr
	fnGetOpenxrApi                                          gdc.MethodBindPtr
	fnRegisterExtensionWrapper                              gdc.MethodBindPtr
}

var ptrsForOpenXRExtensionWrapperExtension ptrsForOpenXRExtensionWrapperExtensionList

func initOpenXRExtensionWrapperExtensionPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRExtensionWrapperExtension")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_openxr_api")
		defer methodName.Destroy()
		ptrsForOpenXRExtensionWrapperExtension.fnGetOpenxrApi = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1637791613))
	}
	{
		methodName := StringNameFromStr("register_extension_wrapper")
		defer methodName.Destroy()
		ptrsForOpenXRExtensionWrapperExtension.fnRegisterExtensionWrapper = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type OpenXRExtensionWrapperExtension struct {
	Object
}

func (me *OpenXRExtensionWrapperExtension) BaseClass() string {
	return "OpenXRExtensionWrapperExtension"
}

func NewOpenXRExtensionWrapperExtension() *OpenXRExtensionWrapperExtension {
	str := StringNameFromStr("OpenXRExtensionWrapperExtension") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRExtensionWrapperExtension{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OpenXRExtensionWrapperExtension) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRExtensionWrapperExtension) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRExtensionWrapperExtension) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRExtensionWrapperExtension) GetOpenxrApi() OpenXRAPIExtension {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewOpenXRAPIExtension()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRExtensionWrapperExtension.fnGetOpenxrApi), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *OpenXRExtensionWrapperExtension) RegisterExtensionWrapper() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRExtensionWrapperExtension.fnRegisterExtensionWrapper), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
