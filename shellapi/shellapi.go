// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package shellapi provides API definitions for accessing
//shell32.dll.
package shellapi

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/shell32"
)

var (
	DragQueryFile func(T.HDROP, T.UINT, VString, T.UINT) (T.UINT, error)

	DragQueryPoint func(T.HDROP, *T.POINT) (T.BOOL, error)

	DragFinish func(T.HDROP)

	DragAcceptFiles func(T.HWND, T.BOOL)

	ShellExecute func(wnd T.HWND,
		operation, file, parameters, directory VString,
		showCmd T.INT) (T.HINSTANCE, error)

	FindExecutable func(
		file, directory VString, result OVString) (T.HINSTANCE, error)

	CommandLineToArgvW func(
		cmdLine T.WString, numArgs *int) (*T.WString, error)
	//TODO(t):*WString was *LPWSTR

	ShellAbout func(
		wnd T.HWND,
		app, otherStuff VString,
		icon T.HICON) (T.INT, error)

	DuplicateIcon func(Inst T.HINSTANCE, icon T.HICON) (T.HICON, error)

	ExtractAssociatedIcon func(
		inst T.HINSTANCE, iconPath VString, icon *T.WORD) (T.HICON, error)

	ExtractAssociatedIconEx func(
		inst T.HINSTANCE,
		iconPath VString,
		iconIndex *T.WORD,
		iconId *T.WORD) (T.HICON, error)

	ExtractIcon func(
		inst T.HINSTANCE,
		exeFileName VString,
		iconIndex T.UINT) (T.HICON, error)

	ExtractIconEx func(
		file VString,
		iconIndex int,
		large, small *T.HICON,
		icons T.UINT) (T.UINT, error)

	SHFreeNameMappings func(hNameMappings T.HANDLE)

	WinExecError func(
		wnd T.HWND,
		err int,
		fileName,
		title VString)

	SHCreateProcessAsUserW func(
		scpi *T.SHCREATEPROCESSINFOW) (T.BOOL, error)

	SHGetFileInfoA func(
		path T.AString,
		fileAttributes T.DWORD,
		sfi *T.SHFILEINFOA,
		fileInfo, flags T.UINT) (T.DWORD_PTR, error)

	SHGetFileInfoW func(
		path T.WString,
		fileAttributes T.DWORD,
		sfi *T.SHFILEINFOW,
		fileInfo, flags T.UINT) (T.DWORD_PTR, error)

	SHGetDiskFreeSpaceEx func(
		directoryName VString,
		freeBytesAvailableToCaller,
		totalNumberOfBytes,
		totalNumberOfFreeBytes *T.ULARGE_INTEGER) (T.BOOL, error)

	SHGetNewLinkInfo func(
		linkTo, dir VString,
		name OVString,
		mustCopy *T.BOOL,
		flags T.UINT) (T.BOOL, error)

	SHInvokePrinterCommand func(
		wnd T.HWND,
		action T.UINT,
		buf1, buf2 VString,
		modal T.BOOL) (T.BOOL, error)

	IsLFNDrive func(Path VString) (T.BOOL, error)

	SHEnumerateUnreadMailAccounts func(
		keyUser T.HKEY,
		index T.DWORD,
		mailAddress OVString,
		cMailAddress int)

	SHGetUnreadMailCount func(
		keyUser T.HKEY,
		mailAddress VString,
		count *T.DWORD,
		fileTime *T.FILETIME,
		shellExecuteCommand OVString,
		cShellExecuteCommand int)

	SHSetUnreadMailCount func(
		mailAddress VString,
		count T.DWORD,
		shellExecuteCommand T.AString)

	SHTestTokenMembership func(
		token T.HANDLE, rid T.ULONG) (T.BOOL, error)

	SHGetImageList func(
		imageList int, riid T.REFIID, Obj **T.VOID)
)

var ShellApiANSIApis = outside.Apis{
	{"ExtractAssociatedIconA", &ExtractAssociatedIcon},
	{"ExtractAssociatedIconExA", &ExtractAssociatedIconEx},
	{"ExtractIconA", &ExtractIcon},
	{"ExtractIconExA", &ExtractIconEx},
	{"FindExecutableA", &FindExecutable},
	{"ShellAboutA", &ShellAbout},
	{"ShellExecuteA", &ShellExecute},
	{"SHGetDiskFreeSpaceExA", &SHGetDiskFreeSpaceEx},
	{"SHGetFileInfoA", &SHGetFileInfoA},
	{"SHInvokePrinterCommandA", &SHInvokePrinterCommand},
}

var ShellApiUnicodeApis = outside.Apis{
	{"ExtractAssociatedIconW", &ExtractAssociatedIcon},
	{"ExtractAssociatedIconExW", &ExtractAssociatedIconEx},
	{"ExtractIconW", &ExtractIcon},
	{"ExtractIconExW", &ExtractIconEx},
	{"FindExecutableW", &FindExecutable},
	{"ShellAboutW", &ShellAbout},
	{"ShellExecuteW", &ShellExecute},
	{"SHGetDiskFreeSpaceExW", &SHGetDiskFreeSpaceEx},
	{"SHGetFileInfoW", &SHGetFileInfoW},
	{"SHInvokePrinterCommandW", &SHInvokePrinterCommand},
}

//NOTR(t):Not implemented
//{"WinExecErrorA", &WinExecError},
//{"WinExecErrorW", &WinExecError},

//TODO(t): Not on XP
//{"SHEnumerateUnreadMailAccounts", &SHEnumerateUnreadMailAccounts},
//{"SHGetUnreadMailCount", &SHGetUnreadMailCount},
//{"SHSetUnreadMailCount", &SHSetUnreadMailCount},

var ShellApiApis = outside.Apis{
	{"CommandLineToArgvW", &CommandLineToArgvW},
	{"DragAcceptFiles", &DragAcceptFiles},
	{"DragFinish", &DragFinish},
	{"DragQueryFile", &DragQueryFile},
	{"DragQueryPoint", &DragQueryPoint},
	{"DuplicateIcon", &DuplicateIcon},
	{"IsLFNDrive", &IsLFNDrive},
	{"SHCreateProcessAsUserW", &SHCreateProcessAsUserW},
	{"SHFreeNameMappings", &SHFreeNameMappings},
	{"SHGetImageList", &SHGetImageList},
	{"SHGetNewLinkInfo", &SHGetNewLinkInfo},
	{"SHTestTokenMembership", &SHTestTokenMembership},
}
