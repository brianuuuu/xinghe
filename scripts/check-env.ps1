#Requires -Version 5.1
$ErrorActionPreference = "Continue"

Write-Host "=== Xinghe P0 Local Environment Check ===" -ForegroundColor Cyan

$tools = @(
    @{ Name = "go";     Args = @("version") },
    @{ Name = "git";    Args = @("--version") },
    @{ Name = "node";   Args = @("-v") },
    @{ Name = "npm";    Args = @("-v") },
    @{ Name = "codex";  Args = @("--version") },
    @{ Name = "claude"; Args = @("--version") }
)

foreach ($t in $tools) {
    $cmd = Get-Command $t.Name -ErrorAction SilentlyContinue
    if (-not $cmd) {
        Write-Host ("[MISSING] {0}" -f $t.Name) -ForegroundColor Yellow
        continue
    }
    $out = & $t.Name @($t.Args) 2>&1 | Select-Object -First 1
    Write-Host ("[OK]      {0,-8} -> {1} ({2})" -f $t.Name, $out, $cmd.Source) -ForegroundColor Green
}
