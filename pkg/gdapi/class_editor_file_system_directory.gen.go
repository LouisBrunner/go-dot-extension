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

type ptrsForEditorFileSystemDirectoryList struct {
	fnGetSubdirCount            gdc.MethodBindPtr
	fnGetSubdir                 gdc.MethodBindPtr
	fnGetFileCount              gdc.MethodBindPtr
	fnGetFile                   gdc.MethodBindPtr
	fnGetFilePath               gdc.MethodBindPtr
	fnGetFileType               gdc.MethodBindPtr
	fnGetFileScriptClassName    gdc.MethodBindPtr
	fnGetFileScriptClassExtends gdc.MethodBindPtr
	fnGetFileImportIsValid      gdc.MethodBindPtr
	fnGetName                   gdc.MethodBindPtr
	fnGetPath                   gdc.MethodBindPtr
	fnGetParent                 gdc.MethodBindPtr
	fnFindFileIndex             gdc.MethodBindPtr
	fnFindDirIndex              gdc.MethodBindPtr
}

var ptrsForEditorFileSystemDirectory ptrsForEditorFileSystemDirectoryList

func initEditorFileSystemDirectoryPtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorFileSystemDirectory")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_subdir_count")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetSubdirCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_subdir")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetSubdir = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2330964164))
	}
	{
		methodName := StringNameFromStr("get_file_count")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetFileCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_file")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_file_path")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetFilePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_file_type")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetFileType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("get_file_script_class_name")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetFileScriptClassName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_file_script_class_extends")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetFileScriptClassExtends = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_file_import_is_valid")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetFileImportIsValid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_name")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("get_path")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_parent")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnGetParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 842323275))
	}
	{
		methodName := StringNameFromStr("find_file_index")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnFindFileIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("find_dir_index")
		defer methodName.Destroy()
		ptrsForEditorFileSystemDirectory.fnFindDirIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}

}

type EditorFileSystemDirectory struct {
	Object
}

func (me *EditorFileSystemDirectory) BaseClass() string {
	return "EditorFileSystemDirectory"
}

func NewEditorFileSystemDirectory() *EditorFileSystemDirectory {
	str := StringNameFromStr("EditorFileSystemDirectory") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorFileSystemDirectory{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *EditorFileSystemDirectory) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorFileSystemDirectory) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorFileSystemDirectory) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorFileSystemDirectory) GetSubdirCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetSubdirCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFileSystemDirectory) GetSubdir(idx int64) EditorFileSystemDirectory {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewEditorFileSystemDirectory()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetSubdir), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileSystemDirectory) GetFileCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetFileCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFileSystemDirectory) GetFile(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileSystemDirectory) GetFilePath(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetFilePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileSystemDirectory) GetFileType(idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetFileType), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileSystemDirectory) GetFileScriptClassName(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetFileScriptClassName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileSystemDirectory) GetFileScriptClassExtends(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetFileScriptClassExtends), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileSystemDirectory) GetFileImportIsValid(idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetFileImportIsValid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFileSystemDirectory) GetName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileSystemDirectory) GetPath() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileSystemDirectory) GetParent() EditorFileSystemDirectory {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewEditorFileSystemDirectory()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnGetParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorFileSystemDirectory) FindFileIndex(name String) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnFindFileIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *EditorFileSystemDirectory) FindDirIndex(name String) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorFileSystemDirectory.fnFindDirIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
