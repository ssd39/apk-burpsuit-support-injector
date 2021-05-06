# apk-burpsuit-support-injector

Build Instruction:

go get
go build

Prerequisite

- to build tool from source make sure you installed go in your system
- java jre required for apktool
- set android path in your enviorment variable so zipalign and jarsigner can be executable from any path in your system.


Usage:
./abinject [path to your apk file]

It will return [given app name]-injected.apk in current directory.
