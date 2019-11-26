package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	//for i:=0;i<1000 ; i++ {
		var in string= "/Users/hq/go/src/gotest/src/company/test.properties"
		pmap := getpropertiesMap(in)
		resmap:=map[string]interface{}{}
		for k,v:= range pmap {
			skeys:= strings.Split(k,".")
			tempmap:=map[string]interface{}{}
			tempmap[skeys[len(skeys)-1]]=v
			var i int = len(skeys) -2
			for ;i>=0 ;i--  {
				tempmap= map[string]interface{}{skeys[i]:tempmap}
			}
			//每次新加入一行都需要进行调整已经生成的结果
			updateMap(resmap,tempmap)
		}
		processlist(&resmap)
		d, _ := yaml.Marshal(&resmap)
		fmt.Println(string(d))
		ioutil.WriteFile("/Users/hq/go/src/gotest/src/company/test.yaml",d,os.ModePerm)
//	}
}

func updateMap( omap map[string]interface{},desmap map[string]interface{}) map[string]interface{} {
	for k,v:= range desmap {
		//如果是map说明没有递归完成，继续完成递归
		if t(v) {
			x,ok:=omap[k]
			if ok {
			omap[k] = updateMap(x.(map[string]interface{}),v.(map[string]interface{}))
			} else {
			omap[k] = updateMap(map[string]interface{}{},v.(map[string]interface{}))
			}
		}else {//终止条件
		//后者覆盖前者.
		omap[k]=desmap[k]
		}
	}
	return omap
}
//数组 这里可以不用 指针的方式，纯粹是为了装b
func processlist(resMap *map[string]interface{}) {
	var count int = len(*resMap)
	for k,v:= range *resMap {
		//判断是否key包含了数组的符号
		if strings.Contains(k,"[",) && strings.Contains(k,"]") && (strings.LastIndex(k,"]") + 1== len(k)) {
			x, ok := (*resMap)[string([]rune(k)[:strings.LastIndex(k,"[")])]
			//已经存在
			if ok {
				//如果是v是map,处理它的孩子并把它加到数组里面
				if t(v) {
					xxx:=v.(map[string]interface{})
					processlist(&xxx)
				}
				x = append(x.([]interface{}), v)
				//创建一个数组子项，给后面生成yaml使用的数组
				(*resMap)[string([]rune(k)[:strings.LastIndex(k,"[")])] = x
				delete(*resMap, k)
				count--;
				//key已经全部遍历完成
				if count==0 {
					return
				}
				//该项处理完成 继续处理下一项
				continue
			} else {
				if t(v) {
					xxx:=v.(map[string]interface{})
					processlist(&xxx)
				}
				//首次
				x = append([]interface{}{}, v)
				(*resMap)[string([]rune(k)[:strings.LastIndex(k,"[")])] = x
				delete(*resMap, k)
				count--;
				//遍历完成
				if count==0 {
					return
				}
				continue
			}
		}
		//如果是字符串说明已经到达最后直接下次循环
		if st(v)  {
			count--;
			if count<=0 {
				return
			}
		continue
		}
		//因为加入了数组所以这里遍历到需要把它跳过
		if intert(v) {
			continue
		}
		//正常的情况不带数组标记的
		x:=v.(map[string]interface{})
			processlist(&x)
		if  0 == count--  {
			return
		}
	}
}

func intert(i interface{})bool {    //函数t 有一个参数i
	switch i.(type) {      //多选语句switch
	//主要用到了这个格式
	case []interface{}:
		return true
	}
	return false
}
func st(i interface{})bool {    //函数t 有一个参数i
	switch i.(type) {      //多选语句switch
	//主要用到了这个格式
	case string:
		return true
	}
	return false
}
func t(i interface{})bool {    //函数t 有一个参数i
	switch i.(type) {      //多选语句switch
	//主要用到了这个格式
	case map[string]interface{}:
		return true
	}
	return false
}
func it(i interface{})bool {    //函数t 有一个参数i
	switch i.(type) {      //多选语句switch
	//主要用到了这个格式
	case []int:
		return true
	}
	return false
}
//文件处理
func getpropertiesMap(in string) map[string]string  {
	content,err:= ioutil.ReadFile( in)
	if err != nil {
		log.Fatal(err)
	}
	var res = string(content)
	res = strings.Trim(res,"\n")
	 mp := make(map[string]string)
	for _,v:= range strings.Split(res,"\n") {
		if strings.HasPrefix(v,"#")|| v=="" {
			continue
		}
		key:= strings.Split(v,"=")[0]
		val:=strings.Split(v,"=")[1]
		val=strings.Trim(val," ")
		key=strings.Trim(key," ")
		mp[key]=val
	}
return mp
}
