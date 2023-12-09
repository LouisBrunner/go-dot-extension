// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type EditorVCSInterface struct {
  obj gdc.ObjectPtr
}

func (me *EditorVCSInterface) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *EditorVCSInterface) BaseClass() string {
  return "EditorVCSInterface"
}



// Enums

type EditorVCSInterfaceChangeType int
const (
  EditorVCSInterfaceChangeTypeChangeTypeNew EditorVCSInterfaceChangeType = 0
  EditorVCSInterfaceChangeTypeChangeTypeModified EditorVCSInterfaceChangeType = 1
  EditorVCSInterfaceChangeTypeChangeTypeRenamed EditorVCSInterfaceChangeType = 2
  EditorVCSInterfaceChangeTypeChangeTypeDeleted EditorVCSInterfaceChangeType = 3
  EditorVCSInterfaceChangeTypeChangeTypeTypechange EditorVCSInterfaceChangeType = 4
  EditorVCSInterfaceChangeTypeChangeTypeUnmerged EditorVCSInterfaceChangeType = 5
)

type EditorVCSInterfaceTreeArea int
const (
  EditorVCSInterfaceTreeAreaTreeAreaCommit EditorVCSInterfaceTreeArea = 0
  EditorVCSInterfaceTreeAreaTreeAreaStaged EditorVCSInterfaceTreeArea = 1
  EditorVCSInterfaceTreeAreaTreeAreaUnstaged EditorVCSInterfaceTreeArea = 2
)

func (me *EditorVCSInterface) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *EditorVCSInterface) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *EditorVCSInterface) XInitialize(project_path String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XSetCredentials(username String, password String, ssh_public_key_path String, ssh_private_key_path String, ssh_passphrase String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XGetModifiedFilesData()  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XStageFile(file_path String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XUnstageFile(file_path String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XDiscardFile(file_path String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XCommit(msg String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XGetDiff(identifier String, area int, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XShutDown()  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XGetVcsName()  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XGetPreviousCommits(max_commits int, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XGetBranchList()  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XGetRemotes()  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XCreateBranch(branch_name String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XRemoveBranch(branch_name String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XCreateRemote(remote_name String, remote_url String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XRemoveRemote(remote_name String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XGetCurrentBranchName()  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XCheckoutBranch(branch_name String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XPull(remote String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XPush(remote String, force bool, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XFetch(remote String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) XGetLineDiff(file_path String, text String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) CreateDiffLine(new_line_no int, old_line_no int, content String, status String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) CreateDiffHunk(old_start int, new_start int, old_lines int, new_lines int, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) CreateDiffFile(new_file String, old_file String, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) CreateCommit(msg String, author String, id String, unix_timestamp int, offset_minutes int, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) CreateStatusFile(file_path String, change_type EditorVCSInterfaceChangeType, area EditorVCSInterfaceTreeArea, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) AddDiffHunksIntoDiffFile(diff_file Dictionary, diff_hunks Dictionary, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) AddLineDiffsIntoDiffHunk(diff_hunk Dictionary, line_diffs Dictionary, )  {
  panic("TODO: implement")
}

func  (me *EditorVCSInterface) PopupError(msg String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
