# install_libraw_windows.ps1

$ErrorActionPreference = "Stop"

$DepDir = "deps"
$InstallDir = Join-Path (Get-Location) "libraw_dist"
$LibRawUrl = "https://github.com/LibRaw/LibRaw.git"
$LibRawTag = "0.21.2"

Write-Host "开始安装 LibRaw..."

# 1. 创建依赖目录
if (-not (Test-Path $DepDir)) {
    New-Item -ItemType Directory -Path $DepDir | Out-Null
}

Set-Location $DepDir

# 2. 克隆 LibRaw
if (-not (Test-Path "LibRaw")) {
    Write-Host "正在克隆 LibRaw..."
    git clone $LibRawUrl
}

Set-Location "LibRaw"

# 3. 检出稳定版本
Write-Host "切换到版本 $LibRawTag..."
git checkout $LibRawTag

# 4. 创建 object 目录 (Makefile.mingw 需要)
if (-not (Test-Path "object")) {
    New-Item -ItemType Directory -Path "object" | Out-Null
}

# 修正 Makefile.mingw 中的 rm 命令 (Windows 没有 rm)
Write-Host "修正 Makefile.mingw..."
git checkout Makefile.mingw
(Get-Content Makefile.mingw) -replace "rm -f lib/libraw.a", "@echo Removing old lib..." | Set-Content Makefile.mingw

# 5. 编译
Write-Host "开始编译 (使用 Makefile.mingw)..."
mingw32-make -f Makefile.mingw library

# 6. 安装 (手动复制文件)
Write-Host "安装到 $InstallDir..."

# 创建安装目录结构
$IncludeDir = Join-Path $InstallDir "include"
$LibDir = Join-Path $InstallDir "lib"

if (-not (Test-Path $IncludeDir)) { New-Item -ItemType Directory -Path $IncludeDir -Force | Out-Null }
if (-not (Test-Path $LibDir)) { New-Item -ItemType Directory -Path $LibDir -Force | Out-Null }

# 复制头文件
# LibRaw 的头文件在 libraw/ 目录下，我们需要把整个 libraw 文件夹复制到 include/ 下
Copy-Item -Path "libraw" -Destination $IncludeDir -Recurse -Force

# 复制库文件
Copy-Item -Path "lib/libraw.a" -Destination $LibDir -Force

Set-Location ../..

# 获取绝对路径并替换反斜杠
$AbsInstallDir = $InstallDir -replace "\\", "/"

Write-Host "`n安装完成!"
Write-Host "头文件位置: $AbsInstallDir/include"
Write-Host "库文件位置: $AbsInstallDir/lib"
Write-Host "`n请修改 libraw.go 文件中的 cgo 指令："
Write-Host "// #cgo CFLAGS: -I${AbsInstallDir}/include" -ForegroundColor Green
Write-Host "// #cgo LDFLAGS: -L${AbsInstallDir}/lib -lraw -lws2_32" -ForegroundColor Green
# 注意：在 Windows 上通常需要链接 ws2_32
