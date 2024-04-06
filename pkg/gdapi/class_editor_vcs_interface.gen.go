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

type ptrsForEditorVCSInterfaceList struct {
	fnXInitialize              gdc.MethodBindPtr
	fnXSetCredentials          gdc.MethodBindPtr
	fnXGetModifiedFilesData    gdc.MethodBindPtr
	fnXStageFile               gdc.MethodBindPtr
	fnXUnstageFile             gdc.MethodBindPtr
	fnXDiscardFile             gdc.MethodBindPtr
	fnXCommit                  gdc.MethodBindPtr
	fnXGetDiff                 gdc.MethodBindPtr
	fnXShutDown                gdc.MethodBindPtr
	fnXGetVcsName              gdc.MethodBindPtr
	fnXGetPreviousCommits      gdc.MethodBindPtr
	fnXGetBranchList           gdc.MethodBindPtr
	fnXGetRemotes              gdc.MethodBindPtr
	fnXCreateBranch            gdc.MethodBindPtr
	fnXRemoveBranch            gdc.MethodBindPtr
	fnXCreateRemote            gdc.MethodBindPtr
	fnXRemoveRemote            gdc.MethodBindPtr
	fnXGetCurrentBranchName    gdc.MethodBindPtr
	fnXCheckoutBranch          gdc.MethodBindPtr
	fnXPull                    gdc.MethodBindPtr
	fnXPush                    gdc.MethodBindPtr
	fnXFetch                   gdc.MethodBindPtr
	fnXGetLineDiff             gdc.MethodBindPtr
	fnCreateDiffLine           gdc.MethodBindPtr
	fnCreateDiffHunk           gdc.MethodBindPtr
	fnCreateDiffFile           gdc.MethodBindPtr
	fnCreateCommit             gdc.MethodBindPtr
	fnCreateStatusFile         gdc.MethodBindPtr
	fnAddDiffHunksIntoDiffFile gdc.MethodBindPtr
	fnAddLineDiffsIntoDiffHunk gdc.MethodBindPtr
	fnPopupError               gdc.MethodBindPtr
}

var ptrsForEditorVCSInterface ptrsForEditorVCSInterfaceList

func initEditorVCSInterfacePtrs(iface gdc.Interface) {

	className := StringNameFromStr("EditorVCSInterface")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create_diff_line")
		defer methodName.Destroy()
		ptrsForEditorVCSInterface.fnCreateDiffLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2901184053))
	}
	{
		methodName := StringNameFromStr("create_diff_hunk")
		defer methodName.Destroy()
		ptrsForEditorVCSInterface.fnCreateDiffHunk = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3784842090))
	}
	{
		methodName := StringNameFromStr("create_diff_file")
		defer methodName.Destroy()
		ptrsForEditorVCSInterface.fnCreateDiffFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2723227684))
	}
	{
		methodName := StringNameFromStr("create_commit")
		defer methodName.Destroy()
		ptrsForEditorVCSInterface.fnCreateCommit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1075983584))
	}
	{
		methodName := StringNameFromStr("create_status_file")
		defer methodName.Destroy()
		ptrsForEditorVCSInterface.fnCreateStatusFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1083471673))
	}
	{
		methodName := StringNameFromStr("add_diff_hunks_into_diff_file")
		defer methodName.Destroy()
		ptrsForEditorVCSInterface.fnAddDiffHunksIntoDiffFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015243225))
	}
	{
		methodName := StringNameFromStr("add_line_diffs_into_diff_hunk")
		defer methodName.Destroy()
		ptrsForEditorVCSInterface.fnAddLineDiffsIntoDiffHunk = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4015243225))
	}
	{
		methodName := StringNameFromStr("popup_error")
		defer methodName.Destroy()
		ptrsForEditorVCSInterface.fnPopupError = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}

}

type EditorVCSInterface struct {
	Object
}

func (me *EditorVCSInterface) BaseClass() string {
	return "EditorVCSInterface"
}

func NewEditorVCSInterface() *EditorVCSInterface {
	str := StringNameFromStr("EditorVCSInterface") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &EditorVCSInterface{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type EditorVCSInterfaceChangeType int

const (
	EditorVCSInterfaceChangeTypeChangeTypeNew        EditorVCSInterfaceChangeType = 0
	EditorVCSInterfaceChangeTypeChangeTypeModified   EditorVCSInterfaceChangeType = 1
	EditorVCSInterfaceChangeTypeChangeTypeRenamed    EditorVCSInterfaceChangeType = 2
	EditorVCSInterfaceChangeTypeChangeTypeDeleted    EditorVCSInterfaceChangeType = 3
	EditorVCSInterfaceChangeTypeChangeTypeTypechange EditorVCSInterfaceChangeType = 4
	EditorVCSInterfaceChangeTypeChangeTypeUnmerged   EditorVCSInterfaceChangeType = 5
)

type EditorVCSInterfaceTreeArea int

const (
	EditorVCSInterfaceTreeAreaTreeAreaCommit   EditorVCSInterfaceTreeArea = 0
	EditorVCSInterfaceTreeAreaTreeAreaStaged   EditorVCSInterfaceTreeArea = 1
	EditorVCSInterfaceTreeAreaTreeAreaUnstaged EditorVCSInterfaceTreeArea = 2
)

func (me *EditorVCSInterface) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *EditorVCSInterface) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *EditorVCSInterface) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *EditorVCSInterface) CreateDiffLine(new_line_no int64, old_line_no int64, content String, status String) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&new_line_no), gdc.ConstTypePtr(&old_line_no), content.AsCTypePtr(), status.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&new_line_no)
	pinner.Pin(&old_line_no)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorVCSInterface.fnCreateDiffLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorVCSInterface) CreateDiffHunk(old_start int64, new_start int64, old_lines int64, new_lines int64) Dictionary {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&old_start), gdc.ConstTypePtr(&new_start), gdc.ConstTypePtr(&old_lines), gdc.ConstTypePtr(&new_lines)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&old_start)
	pinner.Pin(&new_start)
	pinner.Pin(&old_lines)
	pinner.Pin(&new_lines)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorVCSInterface.fnCreateDiffHunk), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorVCSInterface) CreateDiffFile(new_file String, old_file String) Dictionary {
	cargs := []gdc.ConstTypePtr{new_file.AsCTypePtr(), old_file.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorVCSInterface.fnCreateDiffFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorVCSInterface) CreateCommit(msg String, author String, id String, unix_timestamp int64, offset_minutes int64) Dictionary {
	cargs := []gdc.ConstTypePtr{msg.AsCTypePtr(), author.AsCTypePtr(), id.AsCTypePtr(), gdc.ConstTypePtr(&unix_timestamp), gdc.ConstTypePtr(&offset_minutes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&unix_timestamp)
	pinner.Pin(&offset_minutes)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorVCSInterface.fnCreateCommit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorVCSInterface) CreateStatusFile(file_path String, change_type EditorVCSInterfaceChangeType, area EditorVCSInterfaceTreeArea) Dictionary {
	cargs := []gdc.ConstTypePtr{file_path.AsCTypePtr(), gdc.ConstTypePtr(&change_type), gdc.ConstTypePtr(&area)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&change_type)
	pinner.Pin(&area)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorVCSInterface.fnCreateStatusFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorVCSInterface) AddDiffHunksIntoDiffFile(diff_file Dictionary, diff_hunks []Dictionary) Dictionary {
	cargs := []gdc.ConstTypePtr{diff_file.AsCTypePtr(), gdc.ConstTypePtr(&diff_hunks)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&diff_hunks)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorVCSInterface.fnAddDiffHunksIntoDiffFile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorVCSInterface) AddLineDiffsIntoDiffHunk(diff_hunk Dictionary, line_diffs []Dictionary) Dictionary {
	cargs := []gdc.ConstTypePtr{diff_hunk.AsCTypePtr(), gdc.ConstTypePtr(&line_diffs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewDictionary()
	pinner.Pin(&line_diffs)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorVCSInterface.fnAddLineDiffsIntoDiffHunk), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *EditorVCSInterface) PopupError(msg String) {
	cargs := []gdc.ConstTypePtr{msg.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForEditorVCSInterface.fnPopupError), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
