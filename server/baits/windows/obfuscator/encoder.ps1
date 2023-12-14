# Return a base64 string of the supplied command

param (
    [string]$command = ""
)

$bytes = [System.Text.Encoding]::Unicode.GetBytes($command)
$encodedCommand = [Convert]::ToBase64String($bytes)

powershell.exe -EncodedCommand $encodedCommand