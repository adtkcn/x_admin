// Package file_util 基于 os.OpenRoot 的安全文件操作封装。
//
// 所有操作都以程序运行目录作为安全根目录，传入的路径必须是相对路径：
//   - 绝对路径、"../" 逃逸、指向根目录之外的符号链接都会被拒绝并返回错误；
//   - 避免业务代码直接调用 os 包时误写、误删系统文件。
//
// 安全根目录句柄在包加载时打开一次，全程复用，不会每次操作都打开：
//   - Init() 只返回本次初始化的结果，可重复调用；
//   - Close() 在进程退出时调用一次，关闭后所有操作返回 os.ErrClosed，不会重新打开。
//
// 用法：
//
//	if err := file_util.Init(); err != nil { ... }   // 启动时
//	defer file_util.Close()                         // 退出时
//	err := file_util.WriteFile("uploads/a.txt", data)
//	b, err := file_util.ReadFile("uploads/a.txt")
package file_util

import (
	"os"
	"time"
)

const (
	// RootDir 安全根目录：程序运行目录
	RootDir = "."
	// FilePerm 新建文件的默认权限
	FilePerm os.FileMode = 0644
	// DirPerm 新建目录的默认权限
	DirPerm os.FileMode = 0755
)

// root 安全根目录句柄，包加载时打开一次，进程内全程复用；
// rootErr 保存初始化结果，之后不再改动，所有操作前只做一次错误判断。
var root, rootErr = os.OpenRoot(RootDir)

// Init 返回安全根目录的打开结果；句柄在包加载时已打开，可重复调用
func Init() error {
	return rootErr
}

// Close 关闭安全根目录，进程退出时调用一次。
// 关闭后所有文件操作返回 os.ErrClosed，不会重新打开。
func Close() error {
	if root == nil {
		return nil
	}
	return root.Close()
}

// OpenRoot 打开根目录内的子目录作为新的安全根目录。
// 返回的 *os.Root 由调用方负责关闭；子目录句柄独立于父目录，父目录关闭后仍可继续使用。
// 需要长时间持有子目录（例如使用 os.Root.FS()）时使用本函数。
func OpenRoot(path string) (*os.Root, error) {
	if rootErr != nil {
		return nil, rootErr
	}
	return root.OpenRoot(path)
}

// Open 只读打开文件，调用方负责关闭
func Open(path string) (*os.File, error) {
	if rootErr != nil {
		return nil, rootErr
	}
	return root.Open(path)
}

// Create 创建或截断文件（权限 FilePerm），调用方负责关闭
func Create(path string) (*os.File, error) {
	if rootErr != nil {
		return nil, rootErr
	}
	return root.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, FilePerm)
}

// OpenFile 按指定标志打开文件，调用方负责关闭。
// perm 只能是 0o777 内的权限位，否则返回错误。
func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error) {
	if rootErr != nil {
		return nil, rootErr
	}
	return root.OpenFile(path, flag, perm)
}

// ReadFile 读取文件全部内容
func ReadFile(path string) ([]byte, error) {
	if rootErr != nil {
		return nil, rootErr
	}
	return root.ReadFile(path)
}

// WriteFile 覆盖写入文件，文件不存在时创建（权限 FilePerm），父目录需已存在
func WriteFile(path string, data []byte) error {
	if rootErr != nil {
		return rootErr
	}
	return root.WriteFile(path, data, FilePerm)
}

// Mkdir 创建单级目录，父目录不存在时返回错误
func Mkdir(path string) error {
	if rootErr != nil {
		return rootErr
	}
	return root.Mkdir(path, DirPerm)
}

// MkdirAll 递归创建目录，已存在时返回 nil
func MkdirAll(path string) error {
	if rootErr != nil {
		return rootErr
	}
	return root.MkdirAll(path, DirPerm)
}

// Stat 读取文件信息，遇到符号链接时跟随链接
func Stat(path string) (os.FileInfo, error) {
	if rootErr != nil {
		return nil, rootErr
	}
	return root.Stat(path)
}

// Lstat 读取文件信息，遇到符号链接时返回链接本身的信息
func Lstat(path string) (os.FileInfo, error) {
	if rootErr != nil {
		return nil, rootErr
	}
	return root.Lstat(path)
}

// Chmod 修改文件权限
func Chmod(path string, mode os.FileMode) error {
	if rootErr != nil {
		return rootErr
	}
	return root.Chmod(path, mode)
}

// Chown 修改文件所属 uid/gid（Windows 不支持）
func Chown(path string, uid, gid int) error {
	if rootErr != nil {
		return rootErr
	}
	return root.Chown(path, uid, gid)
}

// Lchown 修改文件所属 uid/gid，符号链接则作用于链接本身（Windows 不支持）
func Lchown(path string, uid, gid int) error {
	if rootErr != nil {
		return rootErr
	}
	return root.Lchown(path, uid, gid)
}

// Chtimes 修改文件的访问时间和修改时间
func Chtimes(path string, atime, mtime time.Time) error {
	if rootErr != nil {
		return rootErr
	}
	return root.Chtimes(path, atime, mtime)
}

// Remove 删除文件或空目录
func Remove(path string) error {
	if rootErr != nil {
		return rootErr
	}
	return root.Remove(path)
}

// RemoveAll 递归删除目录及其下所有内容，路径不存在时返回 nil
func RemoveAll(path string) error {
	if rootErr != nil {
		return rootErr
	}
	return root.RemoveAll(path)
}

// Rename 重命名或移动文件/目录，源路径与目标路径必须在同一安全根目录内
func Rename(oldPath, newPath string) error {
	if rootErr != nil {
		return rootErr
	}
	return root.Rename(oldPath, newPath)
}

// Link 创建硬链接 newPath 指向 oldPath，路径不存在时返回错误
func Link(oldPath, newPath string) error {
	if rootErr != nil {
		return rootErr
	}
	return root.Link(oldPath, newPath)
}

// Symlink 创建符号链接 newPath 指向 oldPath。
// 注意：oldPath 不做校验，可以指向根目录之外，调用方需自行保证来源可信。
func Symlink(oldPath, newPath string) error {
	if rootErr != nil {
		return rootErr
	}
	return root.Symlink(oldPath, newPath)
}

// Readlink 读取符号链接指向的目标路径
func Readlink(path string) (string, error) {
	if rootErr != nil {
		return "", rootErr
	}
	return root.Readlink(path)
}
