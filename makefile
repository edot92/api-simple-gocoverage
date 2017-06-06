#simple build to all platform and copy conf :)
#for linux you must set permision to excecute
all: build

build: app

app:
	env  GOOS=darwin GOARCH=amd64 go build  -o ./build/darwin_amd64/apiedy
	cp -f -R  conf build/darwin_amd64 
	env  GOOS=linux GOARCH=386 go build  -o ./build/linux_386/apiedy
	cp -f -R  conf build/darwin_amd64 
	env  GOOS=linux GOARCH=amd64 go build  -o  ./build/linux_amd64/apiedy
	cp -f -R conf build/linux_amd64
	env  GOOS=linux GOARCH=arm go build  -o ./build/linux_arm/apiedy
	cp -f -R conf build/linux_arm
	env  GOOS=windows GOARCH=386 go build  -o ./build/windows_386/apiedy.exe
	cp -f -R conf build/windows_386
	env  GOOS=windows GOARCH=amd64 go build  -o ./build/windows_amd64/apiedy.exe
	cp -f -R conf build/windows_amd64
	env  GOOS=linux GOARCH=mips64 go build  -o ./build/linux_mips64/apiedy
	cp -f -R conf build/linux_mips64
	env  GOOS=linux GOARCH=mips64le go build  -o ./build/linux_mips64le/apiedy
	cp -f -R conf build/linux_mips64le
	env  GOOS=linux GOARCH=mips go build  -o ./build/linux_mips/apiedy
	cp -f -R conf build/linux_mips
	env  GOOS=linux GOARCH=mipsle go build  -o ./build/linux_mipsle/apiedy
	cp -f -R conf build/linux_mipsle
