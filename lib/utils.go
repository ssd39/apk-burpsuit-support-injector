package lib

import (
	"os";
	"errors";
	"fmt";
	"strings";
	"io/ioutil";
	"path/filepath"   	
)

const (
	MAX_BYTES_TO_READ   = 2 * 1024
)


func (a *Apk) CheckFile() error {
	fi, err := os.Lstat(a.FileLocation)
	if err != nil{
		return err
	}
	if fi.Mode()&os.ModeSymlink == 0  && fi.Mode()&os.ModeDir == 0 && fi.Mode()&os.ModeSocket == 0 &&  fi.Mode()&os.ModeCharDevice == 0 && fi.Mode()&os.ModeDevice == 0 && fi.Mode()&os.ModeNamedPipe == 0{
			
		file, err := os.OpenFile(a.FileLocation, os.O_RDONLY, 0666)
		if err != nil{
			return err
		}
		defer file.Close()
	
		contentByte := make([]byte, MAX_BYTES_TO_READ)
		numByte, err := file.Read(contentByte)
		if err != nil{
			return err
		}
		contentByte = contentByte[:numByte]
		
		if numByte > 49  && HasPrefix(contentByte, "PK\x03\x04") && HasPrefix(contentByte[30:],"AndroidManifest.xml"){
			a.Filename = fi.Name()
			splitName:= strings.Split(fi.Name(),".")
			a.Directoryname = strings.Join(splitName[0:len(splitName)-1],".")
			return nil
		}

	}	
	return errors.New(fmt.Sprintf("%s: not valid apk file.", a.FileLocation))
}

func HasPrefix(s []byte, prefix string) bool {
	return len(s) >= len(prefix) && Equal(s[:len(prefix)], prefix)
}

func Equal(a []byte, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range []byte(b) {
		if v != a[i] {
			return false
		}
	}
	return true
}

func (a *Apk) StoreUnsignedApk() error{

	input, err := ioutil.ReadFile(fmt.Sprintf("%s/dist/%s",a.Directoryname,a.Filename))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s-unsigned.apk",a.Directoryname), input, 0644)
	return err
}

func (a *Apk) Clean() error{
	err := os.RemoveAll(a.Directoryname)
	return err
}


func GetExePath() (string,error){
	ex, err := os.Executable()                                                                                                             
    if err != nil {                                                                                                                        
        return "",err                                                                                                                      
    }                                                                                                                                      
    exePath := filepath.Dir(ex) 
	return exePath,nil
}