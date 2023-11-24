# Загрузка ffmpeg
$ffmpegUrl = "https://www.gyan.dev/ffmpeg/builds/ffmpeg-release-essentials.zip"
$ffmpegZipPath = Join-Path -Path $Env:TEMP -ChildPath "ffmpeg.zip"
$ffmpegDestinationPath = Join-Path -Path $Env:USERPROFILE\AppData\Roaming -ChildPath "ffmpeg"

if (-not (Test-Path -Path $ffmpegDestinationPath)) {
    # Загрузка архива ffmpeg
    Invoke-WebRequest -Uri $ffmpegUrl -OutFile $ffmpegZipPath

    # Распаковка архива ffmpeg
    Expand-Archive -Path $ffmpegZipPath -DestinationPath $Env:USERPROFILE\AppData\Roaming

    # Переименование ffmpeg-файлов
    $ffmpegFiles = Get-ChildItem -Path $Env:USERPROFILE\AppData\Roaming -Filter "ffmpeg-*"
    foreach ($file in $ffmpegFiles) {
        $newName = Join-Path -Path $file.DirectoryName -ChildPath "ffmpeg"
        if (-not (Test-Path -Path $newName)) {
            Rename-Item -Path $file.FullName -NewName "ffmpeg"
        }
    }

    # Удаление архива ffmpeg
    Remove-Item -Path $ffmpegZipPath
}

# Добавление ffmpeg в переменные среды текущего пользователя
$oldPath = Get-ItemProperty -Path "HKCU:\Environment" -Name "Path"
$newPath = $oldPath.Path += ";%USERPROFILE%\AppData\Roaming\ffmpeg\bin"
Set-ItemProperty -Path "HKCU:\Environment" -Name "Path" -Value $newPath

# Запуск окна "Свойства системы" для обновления переменных среды
Start-Process SystemPropertiesAdvanced