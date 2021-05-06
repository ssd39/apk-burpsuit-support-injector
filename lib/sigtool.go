package lib

import(
	"os/exec";
	"os";
	"github.com/fatih/color";
	"fmt";
	
)


func (a *Apk) Zipalign() error {
	color.Green("Zip aligning.\n")
	unsigned_apk := fmt.Sprintf("%s-unsigned.apk",a.Directoryname)
	
	cmd := exec.Command("zipalign", "-p", "-f","4", unsigned_apk,fmt.Sprintf("%s-injected.apk",a.Directoryname))
    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
    err := cmd.Run()
	if err != nil {
        return err
    }
	e := os.Remove(unsigned_apk)
	return e
}



func (a *Apk) SignApk() error {
	color.Green("Signing Apk.\n")
	injected_apk:=fmt.Sprintf("%s-injected.apk",a.Directoryname)
	
	exePath,err := GetExePath()
	if err != nil{
		return err
	}

	cmd := exec.Command("jarsigner","-sigalg", "SHA1withRSA", "-digestalg", "SHA1", "-keystore", fmt.Sprintf("%s/sig.keystore",exePath), "-storepass", "homealone", injected_apk, "ddl")
    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
    err = cmd.Run()
	return err
}