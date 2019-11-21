package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
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
	c,_:=json.Marshal(resmap)
	fmt.Println(string(c))
	//ss:=map[string]interface{}{}

	processlist(&resmap)

	cc,_:=json.Marshal(&resmap)
	fmt.Println(string(cc))
	d, _ := yaml.Marshal(&resmap)
	fmt.Println(string(d))
	ioutil.WriteFile("/Users/hq/go/src/gotest/src/company/test.yaml",d,os.ModePerm)
}
func updateMap( omap map[string]interface{},desmap map[string]interface{}) map[string]interface{} {
	for k,v:= range desmap {

		//如果是map说明没有递归完成，继续完成递归
		if t(v) {
			x,ok:=omap[k]
			if ok {
			omap[k] = updateMap(x.(map[string]interface{}),v.(map[string]interface{}))
			} else {
			omap[k] = 	updateMap(map[string]interface{}{},v.(map[string]interface{}))
			}
		}else {//终止条件
					//后者覆盖前者.
					omap[k]=desmap[k]

				}
			}
			return omap
		}

func processlist(resMap *map[string]interface{}) {
	var count int = len(*resMap)
	for k,v:= range *resMap {
		if strings.Contains(k,"[",) && strings.Contains(k,"]") && (strings.LastIndex(k,"]") + 1== len(k)) {
			x, ok := (*resMap)[string([]rune(k)[:strings.LastIndex(k,"[")])]
			if ok {
				if t(v) {
					xxx:=v.(map[string]interface{})
					processlist(&xxx)
				}
				x = append(x.([]interface{}), v)
				(*resMap)[string([]rune(k)[:strings.LastIndex(k,"[")])] = x
				delete(*resMap, k)
				count--;
				if count==0 {
					return
				}
				continue

			} else {
				if t(v) {
					xxx:=v.(map[string]interface{})
					processlist(&xxx)
				}
				x = append([]interface{}{}, v)
				(*resMap)[string([]rune(k)[:strings.LastIndex(k,"[")])] = x
				delete(*resMap, k)
				count--;
				if count==0 {
					return
				}
				continue
			}
		}

		if st(resMap) {
			return
		}
		if st(v) {
			count--;
			if count==0 {
				return
			}
		continue
		}
		count--;
		if count==0 {
			return
		}
		x:=v.(map[string]interface{})
			processlist(&x)
	}


}

/** error function
func processlist(resmap map[string]interface{})  map[string]interface{} {
		{

			for k,v:= range resmap {
			if strings.Contains(k,"[",) && strings.Contains(k,"]") && (strings.LastIndex(k,"]")== len(k)) {
			x,ok:=omap[string([]rune(k)[:len(k)])]
			if ok {
			x = append(x.([]interface{}), v)

		}else {
			x = append(x.([]interface{}), v)
			omap[string([]rune(k)[:len(k)])]=x
		}
		} else {
				if t(v) {
					x,ok:=omap[k]
					if ok {
						omap[k] = processlist(x.(map[string]interface{}),v.(map[string]interface{}))
					} else {
						omap[k] = 	processlist(map[string]interface{}{},v.(map[string]interface{}))
					}
				} else {//终止条件
					//后者覆盖前者.
					omap[k]=resmap[k]

				}

			}
		}
	}
	return omap
}
*/
//

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
