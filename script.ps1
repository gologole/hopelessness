# Получение пути к корневой папке скрипта
$scriptRoot = Split-Path -Parent -Path $MyInvocation.MyCommand.Definition

# Поиск папки "ffmpeg"
$ffmpegPath = Join-Path -Path $scriptRoot -ChildPath "ffmpeg"

if (Test-Path -Path $ffmpegPath) {
    # Поиск папки "bin" внутри папки "ffmpeg"
    $ffmpegBinPath = Join-Path -Path $ffmpegPath -ChildPath "bin"

    if (Test-Path -Path $ffmpegBinPath) {
        # Получение текущих переменных среды
        $envPath = [Environment]::GetEnvironmentVariable("Path", "User")

        # Добавление пути к папке "bin" в переменную среды
        $newPath = "$envPath;$ffmpegBinPath"
        [Environment]::SetEnvironmentVariable("Path", $newPath, "User")

        Write-Host "Путь к папке 'bin' внутри папки 'ffmpeg' добавлен в переменную среды."
    } else {
        Write-Host "Папка 'bin' не найдена внутри папки 'ffmpeg'."
    }
} else {
    Write-Host "Папка 'ffmpeg' не найдена в корневой папке скрипта."
}