package main

import(
	"fmt";
	"android_burpsuit_injector/lib"
	"os"
	"github.com/fatih/color"
)

func main(){
	if len(os.Args) <= 1{
		fmt.Println(color.CyanString("Usage: ./abinject "), color.MagentaString("[apk file path]"))
		return
	}

	a := lib.Apk{FileLocation:os.Args[1]}

	if err := a.CheckFile(); err!=nil{
		color.Red(err.Error())
		return
	}

	if err:= a.Disassamble(); err!=nil{
		color.Red(err.Error())
		return
	}

	if err:=a.InsertNetworkConfig();err!=nil{
		color.Red(err.Error())
		return
	}

	if err:=a.ModifyAndroidManifest();err!=nil{
		color.Red(err.Error())
		return
	}

	if err:= a.Assamble(); err!=nil{
		color.Red(err.Error())
		return
	}

	if err:= a.StoreUnsignedApk(); err!=nil{
		color.Red(err.Error())
		return
	}

	if err:=a.Clean();err!=nil{
		color.Red(err.Error())
		return	
	}

	if err:=a.Zipalign();err!=nil{
		color.Red(err.Error())
		return	
	}

	if err:=a.SignApk();err!=nil{
		color.Red(err.Error())
		return	
	}

	color.Green("Done.\n")

}