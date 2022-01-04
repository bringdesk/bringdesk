
Set-Variable -Name "PRODUCT_NAME" -Value "Bring Desktop"
Set-Variable -Name "PRODUCT_VERSION" -Value "0.3"

Write-Host "=== Bring Desktop - MS-Windows ==="

# Section 1. Generate assets...
#
Write-Host "Step 1. Generate assets..."
$invokeExpressionOptions = @{
    Command = "go generate"
}
Invoke-Expression @invokeExpressionOptions

# Section 2. Compile executables...
#
Write-Host "Step 2. Compile executables..."

# Section 2.1. Compile X86_64 executable...
#
Write-Host "Step 2.1. Compile X86_64 executable..."
$Env:GOOS = "windows"
$Env:GOARCH = "amd64"
$Env:CGO_ENABLED = "1"
$Env:CC = "x86_64-w64-mingw32-gcc.exe"
$Env:CXX = "x86_64-w64-mingw32-g++.exe"
$ARCH = "amd64"
$invokeExpressionOptions = @{
    Command = "go build -o bring-desk-$ARCH.exe"
}
#[Environment]::GetEnvironmentVariables()
Invoke-Expression @invokeExpressionOptions

# Section 2.2. Compile X86 executable...
#
## FIX IT - Write-Host "Step 2.2. Compile X86 executable..."
## FIX IT - $Env:GOOS = "windows"
## FIX IT - $Env:GOARCH = "386"
## FIX IT - $Env:CGO_ENABLED = "1"
## FIX IT - $Env:CC = "gcc.exe"
## FIX IT - $Env:CXX = "g++.exe"
## FIX IT - $ARCH = "386"
## FIX IT - $invokeExpressionOptions = @{
## FIX IT -     Command = "go build -o bring-desk-$ARCH.exe"
## FIX IT - }
## FIX IT - #[Environment]::GetEnvironmentVariables()
## FIX IT - Invoke-Expression @invokeExpressionOptions

# Section 3. Make ZIP portable distribution package...
#
Write-Host "Step 3. Make ZIP portable distribution package..."
$TimeStamp = $(Get-Date -Format 'yyyyMMddHHmmtt')
$compressArchiveOptions = @{
    CompressionLevel = "Optimal"
    LiteralPath = "resources",
        "bring-desk-amd64.exe",
## FIX IT - "bring-desk-386.exe",
#       "libFLAC-8.dll", "libfreetype-6.dll", "libjpeg-9.dll", "libmodplug-1.dll", "libmpg123-0.dll", "libogg-0.dll",
#       "libopus-0.dll", "libopusfile-0.dll", "libpng16-16.dll", "libtiff-5.dll", "libvorbis-0.dll", "libvorbisfile-3.dll",
#       "libwebp-7.dll", "SDL2.dll", "SDL2_image.dll", "SDL2_mixer.dll", "SDL2_ttf.dll", "zlib1.dll",
        "ChangeLog", "LICENSE", "README.md"
    DestinationPath = "BringDesk-${PRODUCT_VERSION}-${TimeStamp}.zip"
}
Compress-Archive @compressArchiveOptions
