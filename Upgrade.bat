@echo off
git pull && powershell -command "Stop-service -Force -name "AnonimasuRobot" -ErrorAction SilentlyContinue; go mod tidy; go build; Start-service -name "AnonimasuRobot""
:: Hail Hydra