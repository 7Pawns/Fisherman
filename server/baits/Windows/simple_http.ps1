$hookUrl = "%s"

while ($true) {
    Start-Sleep 1
    Invoke-WebRequest -Uri $hookUrl
}