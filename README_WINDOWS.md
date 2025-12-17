# Windows 安装指南

本指南介绍了如何在 Windows 环境下(使用 MinGW-w64)安装 LibRaw 依赖并配置 Go 项目。

## 1. 环境要求

- **MinGW-w64**: 提供 gcc, g++ 和 mingw32-make。
  - 验证: `gcc --version`, `mingw32-make --version`
- **CMake**: 用于构建系统配置。
  - 验证: `cmake --version`
- **Git**: 用于下载源码。
- **Go**: 编程语言环境。

## 2. 自动安装脚本

项目根目录下提供了一个 PowerShell 脚本 `install_libraw_windows.ps1`，它可以自动完成下载、编译和安装过程。

在 PowerShell 中运行：

```powershell
powershell -ExecutionPolicy Bypass -File install_libraw_windows.ps1
```

该脚本执行以下操作：
1. 在 `deps/` 目录下克隆 LibRaw 源码。
2. 使用 `mingw32-make` 和 `Makefile.mingw` 编译静态库 (`libraw.a`)。
3. 将编译好的库文件复制到 `C:\mingw64\lib\`。
4. 将头文件复制到 `C:\mingw64\include\`。

> **注意**: 脚本需要写入 `C:\mingw64` 目录的权限。如果没有权限，请以管理员身份运行 PowerShell，或者手动复制文件。

## 3. Go CGO 配置

在 Windows 上使用 LibRaw，需要在 `libraw.go` 中配置正确的链接参数。

已修改的配置如下：

```go
// #cgo LDFLAGS: -lraw -lws2_32 -lstdc++
```

- `-lraw`: 链接 LibRaw 静态库。
- `-lws2_32`: LibRaw 在 Windows 上依赖 Winsock 库。
- `-lstdc++`: 因为 LibRaw 是 C++ 编写的，Go 链接器需要显式链接 C++ 标准库。

## 4. 验证安装

您可以尝试编译示例程序来验证安装是否成功：

```powershell
go build ./cmd/thumb_example/main.go
```

如果生成了 `main.exe` 且无报错，则说明安装成功。
