package lib

import(
	"os/exec";
	"os";
	"github.com/fatih/color";
	"fmt"
)


func (a *Apk) Disassamble() (error){

	color.Green("Decompilng apk.\n")

	exePath,err := GetExePath()
	if err != nil{
		return err
	}

	cmd := exec.Command("java", "-jar",fmt.Sprintf("%s/apktool.jar",exePath),"d", a.FileLocation)
    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
    err = cmd.Run()

	return err
}


func (a *Apk) Assamble() (error){
	
	color.Green("Compiling apk.\n")

	exePath,err := GetExePath()
	if err != nil{
		return err
	}


	cmd := exec.Command("java", "-jar", fmt.Sprintf("%s/apktool.jar",exePath),"b", a.Directoryname)
    cmd.Stdout = os.Stdout
	
    err = cmd.Run()

    if err != nil {
		
		color.Magenta("Re-Compiling using --use-aapt2 flag.")
		
		cmd := exec.Command("java", "-jar", fmt.Sprintf("%s/apktool.jar",exePath),"b", a.Directoryname, "--use-aapt2")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
        
		return err
    }

	return nil
}