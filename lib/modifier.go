package lib

import(
	"os";
	"fmt";
	"github.com/fatih/color";
	"strings";
	"io/ioutil"
)

func (a *Apk) InsertNetworkConfig() error{
	color.Green("Inserting Network Config.\n")
	xmlDir := fmt.Sprintf("%s/res/xml",a.Directoryname)
	_, err := os.Lstat(xmlDir)
	if err !=nil{
		err = os.Mkdir(xmlDir, os.ModePerm)
		if err != nil{
			return err
		}
	}
	f, err := os.Create(fmt.Sprintf("%s/network_security_config.xml",xmlDir))
    if err != nil {
        return err
    }
    defer f.Close()

    _, err2 := f.WriteString("<network-security-config>\n    <base-config>\n        <trust-anchors>\n            <!-- Trust preinstalled CAs -->\n            <certificates src=\"system\" />\n            <!-- Additionally trust user added CAs -->\n            <certificates src=\"user\" />\n        </trust-anchors>\n    </base-config>\n</network-security-config>")
	return err2
}


func (a *Apk) ModifyAndroidManifest() error {
	color.Green("Modifying AndroidManifest.xml.\n")

	mainfiestFile := fmt.Sprintf("%s/AndroidManifest.xml",a.Directoryname)

    content, err := ioutil.ReadFile(mainfiestFile)
    if err != nil {
        return err
    }

	
	txt := strings.ReplaceAll(string(content),"android:allowBackup=\"false\"","android:allowBackup=\"true\"") 
	txt = strings.ReplaceAll(txt,"android:extractNativeLibs=\"false\"","android:extractNativeLibs=\"true\"") 
	txt = strings.ReplaceAll(txt,"<application","<application android:networkSecurityConfig=\"@xml/network_security_config\"") 
	err = ioutil.WriteFile(mainfiestFile, []byte(txt), 0644)

	return err

}

