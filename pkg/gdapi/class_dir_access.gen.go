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

type ptrsForDirAccessList struct {
	fnOpen                     gdc.MethodBindPtr
	fnGetOpenError             gdc.MethodBindPtr
	fnListDirBegin             gdc.MethodBindPtr
	fnGetNext                  gdc.MethodBindPtr
	fnCurrentIsDir             gdc.MethodBindPtr
	fnListDirEnd               gdc.MethodBindPtr
	fnGetFiles                 gdc.MethodBindPtr
	fnGetFilesAt               gdc.MethodBindPtr
	fnGetDirectories           gdc.MethodBindPtr
	fnGetDirectoriesAt         gdc.MethodBindPtr
	fnGetDriveCount            gdc.MethodBindPtr
	fnGetDriveName             gdc.MethodBindPtr
	fnGetCurrentDrive          gdc.MethodBindPtr
	fnChangeDir                gdc.MethodBindPtr
	fnGetCurrentDir            gdc.MethodBindPtr
	fnMakeDir                  gdc.MethodBindPtr
	fnMakeDirAbsolute          gdc.MethodBindPtr
	fnMakeDirRecursive         gdc.MethodBindPtr
	fnMakeDirRecursiveAbsolute gdc.MethodBindPtr
	fnFileExists               gdc.MethodBindPtr
	fnDirExists                gdc.MethodBindPtr
	fnDirExistsAbsolute        gdc.MethodBindPtr
	fnGetSpaceLeft             gdc.MethodBindPtr
	fnCopy                     gdc.MethodBindPtr
	fnCopyAbsolute             gdc.MethodBindPtr
	fnRename                   gdc.MethodBindPtr
	fnRenameAbsolute           gdc.MethodBindPtr
	fnRemove                   gdc.MethodBindPtr
	fnRemoveAbsolute           gdc.MethodBindPtr
	fnSetIncludeNavigational   gdc.MethodBindPtr
	fnGetIncludeNavigational   gdc.MethodBindPtr
	fnSetIncludeHidden         gdc.MethodBindPtr
	fnGetIncludeHidden         gdc.MethodBindPtr
	fnIsCaseSensitive          gdc.MethodBindPtr
}

var ptrsForDirAccess ptrsForDirAccessList

func initDirAccessPtrs(iface gdc.Interface) {

	className := StringNameFromStr("DirAccess")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("open")
		defer methodName.Destroy()
		ptrsForDirAccess.fnOpen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1923528528))
	}
	{
		methodName := StringNameFromStr("get_open_error")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetOpenError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166280745))
	}
	{
		methodName := StringNameFromStr("list_dir_begin")
		defer methodName.Destroy()
		ptrsForDirAccess.fnListDirBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2610976713))
	}
	{
		methodName := StringNameFromStr("get_next")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetNext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("current_is_dir")
		defer methodName.Destroy()
		ptrsForDirAccess.fnCurrentIsDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("list_dir_end")
		defer methodName.Destroy()
		ptrsForDirAccess.fnListDirEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_files")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetFiles = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
	}
	{
		methodName := StringNameFromStr("get_files_at")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetFilesAt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3538744774))
	}
	{
		methodName := StringNameFromStr("get_directories")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetDirectories = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
	}
	{
		methodName := StringNameFromStr("get_directories_at")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetDirectoriesAt = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3538744774))
	}
	{
		methodName := StringNameFromStr("get_drive_count")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetDriveCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_drive_name")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetDriveName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 990163283))
	}
	{
		methodName := StringNameFromStr("get_current_drive")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetCurrentDrive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("change_dir")
		defer methodName.Destroy()
		ptrsForDirAccess.fnChangeDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("get_current_dir")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetCurrentDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1287308131))
	}
	{
		methodName := StringNameFromStr("make_dir")
		defer methodName.Destroy()
		ptrsForDirAccess.fnMakeDir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("make_dir_absolute")
		defer methodName.Destroy()
		ptrsForDirAccess.fnMakeDirAbsolute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("make_dir_recursive")
		defer methodName.Destroy()
		ptrsForDirAccess.fnMakeDirRecursive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("make_dir_recursive_absolute")
		defer methodName.Destroy()
		ptrsForDirAccess.fnMakeDirRecursiveAbsolute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("file_exists")
		defer methodName.Destroy()
		ptrsForDirAccess.fnFileExists = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
	}
	{
		methodName := StringNameFromStr("dir_exists")
		defer methodName.Destroy()
		ptrsForDirAccess.fnDirExists = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
	}
	{
		methodName := StringNameFromStr("dir_exists_absolute")
		defer methodName.Destroy()
		ptrsForDirAccess.fnDirExistsAbsolute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2323990056))
	}
	{
		methodName := StringNameFromStr("get_space_left")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetSpaceLeft = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("copy")
		defer methodName.Destroy()
		ptrsForDirAccess.fnCopy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1063198817))
	}
	{
		methodName := StringNameFromStr("copy_absolute")
		defer methodName.Destroy()
		ptrsForDirAccess.fnCopyAbsolute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1063198817))
	}
	{
		methodName := StringNameFromStr("rename")
		defer methodName.Destroy()
		ptrsForDirAccess.fnRename = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 852856452))
	}
	{
		methodName := StringNameFromStr("rename_absolute")
		defer methodName.Destroy()
		ptrsForDirAccess.fnRenameAbsolute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 852856452))
	}
	{
		methodName := StringNameFromStr("remove")
		defer methodName.Destroy()
		ptrsForDirAccess.fnRemove = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("remove_absolute")
		defer methodName.Destroy()
		ptrsForDirAccess.fnRemoveAbsolute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 166001499))
	}
	{
		methodName := StringNameFromStr("set_include_navigational")
		defer methodName.Destroy()
		ptrsForDirAccess.fnSetIncludeNavigational = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_include_navigational")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetIncludeNavigational = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_include_hidden")
		defer methodName.Destroy()
		ptrsForDirAccess.fnSetIncludeHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_include_hidden")
		defer methodName.Destroy()
		ptrsForDirAccess.fnGetIncludeHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_case_sensitive")
		defer methodName.Destroy()
		ptrsForDirAccess.fnIsCaseSensitive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
	}

}

type DirAccess struct {
	RefCounted
}

func (me *DirAccess) BaseClass() string {
	return "DirAccess"
}

func NewDirAccess() *DirAccess {
	str := StringNameFromStr("DirAccess") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &DirAccess{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *DirAccess) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *DirAccess) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *DirAccess) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func DirAccessOpen(path String) DirAccess {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDirAccess()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnOpen), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func DirAccessGetOpenError() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetOpenError), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DirAccess) ListDirBegin() Error {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnListDirBegin), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DirAccess) GetNext() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetNext), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DirAccess) CurrentIsDir() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnCurrentIsDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DirAccess) ListDirEnd() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnListDirEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DirAccess) GetFiles() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetFiles), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func DirAccessGetFilesAt(path String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetFilesAt), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DirAccess) GetDirectories() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetDirectories), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func DirAccessGetDirectoriesAt(path String) PackedStringArray {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetDirectoriesAt), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func DirAccessGetDriveCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetDriveCount), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func DirAccessGetDriveName(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetDriveName), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DirAccess) GetCurrentDrive() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetCurrentDrive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DirAccess) ChangeDir(to_dir String) Error {
	cargs := []gdc.ConstTypePtr{to_dir.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnChangeDir), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DirAccess) GetCurrentDir(include_drive bool) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&include_drive)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&include_drive)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetCurrentDir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *DirAccess) MakeDir(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnMakeDir), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func DirAccessMakeDirAbsolute(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnMakeDirAbsolute), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DirAccess) MakeDirRecursive(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnMakeDirRecursive), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func DirAccessMakeDirRecursiveAbsolute(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnMakeDirRecursiveAbsolute), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DirAccess) FileExists(path String) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnFileExists), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DirAccess) DirExists(path String) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnDirExists), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func DirAccessDirExistsAbsolute(path String) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnDirExistsAbsolute), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DirAccess) GetSpaceLeft() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetSpaceLeft), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DirAccess) Copy(from String, to String, chmod_flags int64) Error {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), gdc.ConstTypePtr(&chmod_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&chmod_flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnCopy), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func DirAccessCopyAbsolute(from String, to String, chmod_flags int64) Error {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), gdc.ConstTypePtr(&chmod_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&chmod_flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnCopyAbsolute), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DirAccess) Rename(from String, to String) Error {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnRename), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func DirAccessRenameAbsolute(from String, to String) Error {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnRenameAbsolute), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DirAccess) Remove(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnRemove), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func DirAccessRemoveAbsolute(path String) Error {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnRemoveAbsolute), nil, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *DirAccess) SetIncludeNavigational(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnSetIncludeNavigational), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DirAccess) GetIncludeNavigational() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetIncludeNavigational), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DirAccess) SetIncludeHidden(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnSetIncludeHidden), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *DirAccess) GetIncludeHidden() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnGetIncludeHidden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *DirAccess) IsCaseSensitive(path String) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForDirAccess.fnIsCaseSensitive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
