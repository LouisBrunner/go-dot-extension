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

type ptrsForOpenXRInteractionProfileMetadataList struct {
	fnRegisterProfileRename      gdc.MethodBindPtr
	fnRegisterTopLevelPath       gdc.MethodBindPtr
	fnRegisterInteractionProfile gdc.MethodBindPtr
	fnRegisterIoPath             gdc.MethodBindPtr
}

var ptrsForOpenXRInteractionProfileMetadata ptrsForOpenXRInteractionProfileMetadataList

func initOpenXRInteractionProfileMetadataPtrs(iface gdc.Interface) {

	className := StringNameFromStr("OpenXRInteractionProfileMetadata")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("register_profile_rename")
		defer methodName.Destroy()
		ptrsForOpenXRInteractionProfileMetadata.fnRegisterProfileRename = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3186203200))
	}
	{
		methodName := StringNameFromStr("register_top_level_path")
		defer methodName.Destroy()
		ptrsForOpenXRInteractionProfileMetadata.fnRegisterTopLevelPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 254767734))
	}
	{
		methodName := StringNameFromStr("register_interaction_profile")
		defer methodName.Destroy()
		ptrsForOpenXRInteractionProfileMetadata.fnRegisterInteractionProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 254767734))
	}
	{
		methodName := StringNameFromStr("register_io_path")
		defer methodName.Destroy()
		ptrsForOpenXRInteractionProfileMetadata.fnRegisterIoPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3443511926))
	}

}

type OpenXRInteractionProfileMetadata struct {
	Object
}

func (me *OpenXRInteractionProfileMetadata) BaseClass() string {
	return "OpenXRInteractionProfileMetadata"
}

func NewOpenXRInteractionProfileMetadata() *OpenXRInteractionProfileMetadata {
	str := StringNameFromStr("OpenXRInteractionProfileMetadata") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &OpenXRInteractionProfileMetadata{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *OpenXRInteractionProfileMetadata) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *OpenXRInteractionProfileMetadata) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *OpenXRInteractionProfileMetadata) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *OpenXRInteractionProfileMetadata) RegisterProfileRename(old_name String, new_name String) {
	cargs := []gdc.ConstTypePtr{old_name.AsCTypePtr(), new_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfileMetadata.fnRegisterProfileRename), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRInteractionProfileMetadata) RegisterTopLevelPath(display_name String, openxr_path String, openxr_extension_name String) {
	cargs := []gdc.ConstTypePtr{display_name.AsCTypePtr(), openxr_path.AsCTypePtr(), openxr_extension_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfileMetadata.fnRegisterTopLevelPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRInteractionProfileMetadata) RegisterInteractionProfile(display_name String, openxr_path String, openxr_extension_name String) {
	cargs := []gdc.ConstTypePtr{display_name.AsCTypePtr(), openxr_path.AsCTypePtr(), openxr_extension_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfileMetadata.fnRegisterInteractionProfile), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *OpenXRInteractionProfileMetadata) RegisterIoPath(interaction_profile String, display_name String, toplevel_path String, openxr_path String, openxr_extension_name String, action_type OpenXRActionActionType) {
	cargs := []gdc.ConstTypePtr{interaction_profile.AsCTypePtr(), display_name.AsCTypePtr(), toplevel_path.AsCTypePtr(), openxr_path.AsCTypePtr(), openxr_extension_name.AsCTypePtr(), gdc.ConstTypePtr(&action_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOpenXRInteractionProfileMetadata.fnRegisterIoPath), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
