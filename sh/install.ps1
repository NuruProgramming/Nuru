# Hii ni skripti ya kusanikisha programu ya Nuru
# Hii skripti ni ya "Operating System" ya "Windows" pekee
# Skripti hii imejaribiwa kwenye toleo la Kumi na kuendelea ya "Windows"

param (
    [Parameter(HelpMessage="The version to be downloaded from GitHub (default: latest)")]
    [ValidateNotNullOrEmpty()]
    [string]$Version = "latest",

    [Parameter(HelpMessage="The base path to be used when installing (default: C:\Program Files\nuru)")]
    [ValidateNotNullOrEmpty()]
    [string]$Prefix = "C:\Program Files\nuru",

    [Parameter(HelpMessage="Show this help message")]
    [switch]$Help
)

$RELEASE_URL = "https://github.com/NuruProgramming/Nuru/releases"
$LATEST_URL = 'https://api.github.com/repos/NuruProgramming/Nuru/releases/latest'

function Get-Help {
    $HelpMessage = @"
Usage install.ps1 [Options]

Options:
    -Version The version to be downloaded from GitHub (default: latest)
    -Prefix  The base path to be used when installing (default: C:\Program Files\nuru)
    -Help    Show this help message
"@

    Write-Host $HelpMessage
}

function Get-Arch {
    $ARCH = [System.Runtime.InteropServices.RuntimeInformation, mscorlib]::OSArchitecture.ToString().ToLower()
    Write-Verbose "Detected system architecture: $ARCH"

    switch ($ARCH) {
        "x86" { return "i386" }
        "x64" { return "amd64" }
        "arm64" { return "arm64" }
        Default {
            Write-Error "Unsupported architecture: $ARCH"
            exit 1
        }
    }
}

function Download-File {
    param (
        [Parameter(Mandatory=$true, HelpMessage="The URL from where the file is to be downloaded.")]
        [ValidateNotNullOrEmpty()]
        [string]$URL,

        [Parameter(HelpMessage="Specify the local path to save the file")]
        [string]$Save,

        [Parameter(HelpMessage="Specify if the content is RESTful or not")]
        [switch]$IsReq
    )

    Write-Verbose "Downloading file from '$URL'"

    # Ensure TLS 1.2 for secure communication
    [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

    try {
        if ($IsReq) {
            if ($Save) {
                Invoke-WebRequest -Uri $URL -OutFile $Save
                return $Save
            } else {
                return Invoke-WebRequest -Uri $URL
            }
        } else {
            return Invoke-RestMethod -Uri $URL
        }
    } catch {
        Write-Host "Error while downloading: $_"
        exit 2
    }
}

function Main {
    if ($Help) {
        Get-Help
        exit 0
    }

    $ARCH = Get-Arch

    $Version = $Version.Trim()

    if ($Version -eq "latest") {
        Write-Verbose "Fetching latest version tag from GitHub..."
        $latestRelease = Download-File -URL $LATEST_URL
        $Version = $latestRelease.tag_name
        Write-Host "Latest version: $Version"
    }

    if (-not $Version) {
        Write-Host "Error: Version cannot be null or empty."
        exit 2
    }

    # Ensure the installation path exists
    if (!(Test-Path -Path $Prefix)) {
        Write-Verbose "Creating directory at: $Prefix"
        try {
            New-Item -ItemType Directory -Path $Prefix -Force
        } catch {
            Write-Host "Error creating directory: $_"
            exit 2
        }
    }

    $RFileName = "nuru_Windows_$ARCH"
    $TempDir = "$env:TEMP\$RFileName"
    $ZIP_FileName = "$TempDir.zip"
    $ZIP_URL = "$RELEASE_URL/download/$Version/$RFileName.zip"

    # Download the zip file containing the release
    $ZIP_FILE = Download-File -IsReq -URL $ZIP_URL -Save $ZIP_FileName

    Write-Verbose "Extracting $ZIP_FILE to $TempDir"
    try {
        Expand-Archive -Path $ZIP_FILE -DestinationPath $TempDir
    } catch {
        Write-Host "Error while extracting the archive: $_"
        exit 2
    }

    # Clean up the downloaded zip file
    Remove-Item -Path $ZIP_FILE -Force

    # Copy the executable to the desired installation path
    Write-Verbose "Copying nuru.exe to $Prefix"
    try {
        Copy-Item -Path "$TempDir/nuru.exe" -Destination "$Prefix/nuru.exe" -Force
    } catch {
        Write-Host "Error while copying the executable: $_"
        exit 2
    } finally {
        # Cleanup temporary files after installation
        Remove-Item -Path $TempDir -Recurse -Force
    }

    Write-Host "Installation completed successfully."
}

Main
