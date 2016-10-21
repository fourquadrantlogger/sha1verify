package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
	"crypto/sha1"
	"strconv"
	"strings"

)

var (
	folder=flag.String("folder","","文件夹路径")
	rootpath=""
	ignorefilepath=flag.String("gitignore",".gitignore","忽略指定的文件或文件夹的配置文件")

	writeFile=flag.String("outfile","sha1.txt","输出文件名")
	files =make(map[string]struct{
		Sha1 string
		Size int64})
	ignoreconfigs=make(map[string]interface{})
)

func walk(folderpath string)  {
	err := filepath.Walk(folderpath, func(fpath string, f os.FileInfo, err error) error {
		if(Ignored(fpath)){
			return nil
		}

		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}else{
			fullpath := fpath
			bytes,_:=ioutil.ReadFile(fullpath)
			sha1bytes:=sha1.Sum(bytes)
			files[fullpath]=struct{Sha1 string
					       Size int64}{Sha1:fmt.Sprintf("%x",sha1bytes), Size:f.Size(),}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func wildmatch(parten string, str string) bool {
	//fmt.Println(parten,str)
	if(len(str)<len(parten)){
		return false
	}
	if(str[:len(parten)]==parten){
		return true
	}
	return false
}

func Ignored(parentpath string)bool{
	fullpath := parentpath
	shortpath:=parentpath[len(rootpath):]

	for key,_:=range ignoreconfigs{
		if(wildmatch(key,shortpath)){
			return true
		}
		if(wildmatch(key,fullpath)){
			return true
		}
	}
	return false
}
func main() {
	flag.Parse()
	f,err:=os.Open(*ignorefilepath)
	if(err!=nil){
		ignoreconfigs=nil
	}else {
		configbytes,_:=ioutil.ReadAll(f)
		lines:=strings.Split(string(configbytes),"\n")
		for index,line:=range lines{
			ignoreconfigs[line]=index
		}
	}
	//fmt.Println(ignoreconfigs)
	if(*folder==""){
		arg0,_:=os.Getwd()
		f,_:=os.Stat(arg0)
		if(f.IsDir()){

			rootpath=arg0
			*folder=rootpath
		}else {
			rootpath=filepath.Dir(arg0)
			*folder=rootpath
		}
	}else {
		f,_:=os.Stat(*folder)
		if(f.IsDir()){
			rootpath=*folder
			*folder=rootpath
		}else {
			rootpath=filepath.Dir(*folder)
			*folder=rootpath
		}
	}
	//fmt.Println("rootpath",rootpath)
	walk(*folder)
	resultfile:="sha1 results"
	for key,value:=range files{
		fmt.Println(key,value.Size,"\t",value.Sha1)
		resultfile+="\n"+key+","+strconv.FormatInt(value.Size,10)+","+value.Sha1;
	}
	ioutil.WriteFile(*writeFile,[]byte(resultfile),os.ModePerm)
}
