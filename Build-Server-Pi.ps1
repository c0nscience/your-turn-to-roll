cd .\frontend\
npm ci
npm run build
cd ..
Copy-Item -Recurse -Path .\frontend\dist -Destination .\server\pkg\public\
cd .\server\
GOOS=linux GOARCH=arm GOARM=5 go build -tags embed -o web.exe .\cmd\main.go
Remove-Item -Recurse .\pkg\public\dist
cd ..
Move-Item -Force .\server\web.exe -Destination .