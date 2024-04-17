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

type ptrsForEditorFileDialogList struct {
	fnClearFilters               gdc.MethodBindPtr
	fnAddFilter                  gdc.MethodBindPtr
	fnSetFilters                 gdc.MethodBindPtr
	fnGetFilters                 gdc.MethodBindPtr
	fnGetOptionName              gdc.MethodBindPtr
	fnGetOptionValues            gdc.MethodBindPtr
	fnGetOptionDefault           gdc.MethodBindPtr
	fnSetOptionName              gdc.MethodBindPtr
	fnSetOptionValues            gdc.MethodBindPtr
	fnSetOptionDefault           gdc.MethodBindPtr
	fnSetOptionCount             gdc.MethodBindPtr
	fnGetOptionCount             gdc.MethodBindPtr
	fnAddOption                  gdc.MethodBindPtr
	fnGetSelectedOptions         gdc.MethodBindPtr
	fnGetCurrentDir              gdc.MethodBindPtr
	fnGetCurrentFile             gdc.MethodBindPtr
	fnGetCurrentPath             gdc.MethodBindPtr
	fnSetCurrentDir              gdc.MethodBindPtr
	fnSetCurrentFile             gdc.MethodBindPtr
	fnSetCurrentPath             gdc.MethodBindPtr
	fnSetFileMode                gdc.MethodBindPtr
	fnGetFileMode                gdc.MethodBindPtr
	fnGetVbox                    gdc.MethodBindPtr
	fnGetLineEdit                gdc.MethodBindPtr
	fnSetAccess                  gdc.MethodBindPtr
	fnGetAccess                  gdc.MethodBindPtr
	fnSetShowHiddenFiles         gdc.MethodBindPtr
	fnIsShowingHiddenFiles       gdc.MethodBindPtr
	fnSetDisplayMode             gdc.MethodBindPtr
	fnGetDisplayMode             gdc.MethodBindPtr
	fnSetDisableOverwriteWarning gdc.MethodBindPtr
	fnIsOverwriteWarningDisabled gdc.MethodBindPtr
	fnAddSideMenu                gdc.MethodBindPtr
	fnInvalidate                 gdc.MethodBindPtr
}

var ptrsForEditorFileDialog ptrsForEditorFileDialogList

func initEditorFileDialogPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorFileDialog")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("clear_filters")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnClearFilters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("add_filter")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnAddFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3388804757))
	}
	{
		methodName := StringNameFromStr("set_filters")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetFilters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015028928))
	}
	{
		methodName := StringNameFromStr("get_filters")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetFilters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("get_option_name")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetOptionName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_option_values")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetOptionValues = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 647634434))
	}
	{
		methodName := StringNameFromStr("get_option_default")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetOptionDefault = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_option_name")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetOptionName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("set_option_values")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetOptionValues = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3353661094))
	}
	{
		methodName := StringNameFromStr("set_option_default")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetOptionDefault = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_option_count")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetOptionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_option_count")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetOptionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("add_option")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnAddOption = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 149592325))
	}
	{
		methodName := StringNameFromStr("get_selected_options")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetSelectedOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
	}
	{
		methodName := StringNameFromStr("get_current_dir")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetCurrentDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_current_file")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetCurrentFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_current_path")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetCurrentPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_current_dir")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetCurrentDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_current_file")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetCurrentFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_current_path")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetCurrentPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("set_file_mode")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetFileMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 274150415))
	}
	{
		methodName := StringNameFromStr("get_file_mode")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetFileMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2681044145))
	}
	{
		methodName := StringNameFromStr("get_vbox")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetVbox = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 915758477))
	}
	{
		methodName := StringNameFromStr("get_line_edit")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetLineEdit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4071694264))
	}
	{
		methodName := StringNameFromStr("set_access")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetAccess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3882893764))
	}
	{
		methodName := StringNameFromStr("get_access")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetAccess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 778734016))
	}
	{
		methodName := StringNameFromStr("set_show_hidden_files")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetShowHiddenFiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_showing_hidden_files")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnIsShowingHiddenFiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_display_mode")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetDisplayMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3049004050))
	}
	{
		methodName := StringNameFromStr("get_display_mode")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnGetDisplayMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3517174669))
	}
	{
		methodName := StringNameFromStr("set_disable_overwrite_warning")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnSetDisableOverwriteWarning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_overwrite_warning_disabled")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnIsOverwriteWarningDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("add_side_menu")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnAddSideMenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 402368861))
	}
	{
		methodName := StringNameFromStr("invalidate")
		defer methodName.Destroy()
		ptrsForEditorFileDialog.fnInvalidate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type EditorFileDialog struct {
	ConfirmationDialog
}

func (me *EditorFileDialog) BaseClass() string {
	return "EditorFileDialog"
}

func NewEditorFileDialog() *EditorFileDialog {
	str := StringNameFromStr("EditorFileDialog") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorFileDialog{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type EditorFileDialogFileMode int

const (
	EditorFileDialogFileModeFileModeOpenFile  EditorFileDialogFileMode = 0
	EditorFileDialogFileModeFileModeOpenFiles EditorFileDialogFileMode = 1
	EditorFileDialogFileModeFileModeOpenDir   EditorFileDialogFileMode = 2
	EditorFileDialogFileModeFileModeOpenAny   EditorFileDialogFileMode = 3
	EditorFileDialogFileModeFileModeSaveFile  EditorFileDialogFileMode = 4
)

type EditorFileDialogAccess int

const (
	EditorFileDialogAccessAccessResources  EditorFileDialogAccess = 0
	EditorFileDialogAccessAccessUserdata   EditorFileDialogAccess = 1
	EditorFileDialogAccessAccessFilesystem EditorFileDialogAccess = 2
)

type EditorFileDialogDisplayMode int

const (
	EditorFileDialogDisplayModeDisplayThumbnails EditorFileDialogDisplayMode = 0
	EditorFileDialogDisplayModeDisplayList       EditorFileDialogDisplayMode = 1
)

func (me *EditorFileDialog) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorFileDialog) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorFileDialog) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorFileDialog) ClearFilters() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnClearFilters), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) AddFilter(filter String, description String) {
	cargs := []gdc.ConstTypePtr{filter.AsCTypePtr(), description.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnAddFilter), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) SetFilters(filters PackedStringArray) {
	cargs := []gdc.ConstTypePtr{filters.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetFilters), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) GetFilters() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetFilters), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileDialog) GetOptionName(option int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&option)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetOptionName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileDialog) GetOptionValues(option int64) PackedStringArray {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()
	pinner.Pin(&option)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetOptionValues), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileDialog) GetOptionDefault(option int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&option)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetOptionDefault), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFileDialog) SetOptionName(option int64, name String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetOptionName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) SetOptionValues(option int64, values PackedStringArray) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option), values.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetOptionValues), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) SetOptionDefault(option int64, default_value_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option), gdc.ConstTypePtr(&default_value_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetOptionDefault), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) SetOptionCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetOptionCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) GetOptionCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetOptionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFileDialog) AddOption(name String, values PackedStringArray, default_value_index int64) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), values.AsCTypePtr(), gdc.ConstTypePtr(&default_value_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnAddOption), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) GetSelectedOptions() Dictionary {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetSelectedOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileDialog) GetCurrentDir() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetCurrentDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileDialog) GetCurrentFile() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetCurrentFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileDialog) GetCurrentPath() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetCurrentPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileDialog) SetCurrentDir(dir String) {
	cargs := []gdc.ConstTypePtr{dir.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetCurrentDir), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) SetCurrentFile(file String) {
	cargs := []gdc.ConstTypePtr{file.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetCurrentFile), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) SetCurrentPath(path String) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetCurrentPath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) SetFileMode(mode EditorFileDialogFileMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetFileMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) GetFileMode() EditorFileDialogFileMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret EditorFileDialogFileMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetFileMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *EditorFileDialog) GetVbox() VBoxContainer {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVBoxContainer()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetVbox), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileDialog) GetLineEdit() LineEdit {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewLineEdit()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetLineEdit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileDialog) SetAccess(access EditorFileDialogAccess) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&access)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetAccess), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) GetAccess() EditorFileDialogAccess {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret EditorFileDialogAccess

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetAccess), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *EditorFileDialog) SetShowHiddenFiles(show bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetShowHiddenFiles), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) IsShowingHiddenFiles() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnIsShowingHiddenFiles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFileDialog) SetDisplayMode(mode EditorFileDialogDisplayMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetDisplayMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) GetDisplayMode() EditorFileDialogDisplayMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret EditorFileDialogDisplayMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnGetDisplayMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *EditorFileDialog) SetDisableOverwriteWarning(disable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnSetDisableOverwriteWarning), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) IsOverwriteWarningDisabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnIsOverwriteWarningDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFileDialog) AddSideMenu(menu Control, title String) {
	cargs := []gdc.ConstTypePtr{menu.AsCTypePtr(), title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnAddSideMenu), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *EditorFileDialog) Invalidate() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileDialog.fnInvalidate), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type EditorFileDialogFileSelectedSignalFn func(path String)

func (me *EditorFileDialog) ConnectFileSelected(subs SignalSubscribers, fn EditorFileDialogFileSelectedSignalFn) {
	sig := StringNameFromStr("file_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileDialog) DisconnectFileSelected(subs SignalSubscribers, fn EditorFileDialogFileSelectedSignalFn) {
	sig := StringNameFromStr("file_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type EditorFileDialogFilesSelectedSignalFn func(paths PackedStringArray)

func (me *EditorFileDialog) ConnectFilesSelected(subs SignalSubscribers, fn EditorFileDialogFilesSelectedSignalFn) {
	sig := StringNameFromStr("files_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileDialog) DisconnectFilesSelected(subs SignalSubscribers, fn EditorFileDialogFilesSelectedSignalFn) {
	sig := StringNameFromStr("files_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type EditorFileDialogDirSelectedSignalFn func(dir String)

func (me *EditorFileDialog) ConnectDirSelected(subs SignalSubscribers, fn EditorFileDialogDirSelectedSignalFn) {
	sig := StringNameFromStr("dir_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *EditorFileDialog) DisconnectDirSelected(subs SignalSubscribers, fn EditorFileDialogDirSelectedSignalFn) {
	sig := StringNameFromStr("dir_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
