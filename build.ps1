$sourcecode = "."
$target = "build\workshop-api"
# Windows, 64-bit
$env:GOOS = 'windows'; $env:GOARCH = 'amd64';             go build -o "$($target)-win64.exe" -ldflags "-s -w" $sourcecode
# Linux, 64-bit
$env:GOOS = 'linux';   $env:GOARCH = 'amd64';             go build -o "$($target)-linux64"   -ldflags "-s -w" $sourcecode
