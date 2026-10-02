#!/usr/bin/env sh
# Faaliyet Zaman Çizelgesi - Windows exe derleme betiği (Linux/macOS'ta çapraz derleme)
# Gereksinim: Go 1.22+ ve internet erişimi (ilk derlemede bağımlılıklar indirilir)
set -e
cd "$(dirname "$0")"
cp ../app/faaliyet-cizelgesi.html app.html
go mod tidy
mkdir -p ../dist
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-H windowsgui -s -w" -o ../dist/FaaliyetCizelgesi.exe .
echo "Tamamlandı: dist/FaaliyetCizelgesi.exe"
