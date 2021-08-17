// AGPL License
// Copyright (c) 2021 ysicing <i@ysicing.me>

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ysicing/habor-sync/api"
	"k8s.io/klog/v2"
	"os"
)

var haborapi *api.Client

func init()  {
	c := api.NewClient("", "", "https://old.hub/api")
	haborapi =c
	klog.InitFlags(nil)
}

func main()  {
	ps, _, err := haborapi.Project.List()
	if err != nil {
		panic(err)
	}
	piddata := make(map[string]interface{})
	for _, p := range *ps {
		repos, _, err := haborapi.Repositories.List(p.ProjectID)
		if err != nil {
			continue
		}
		// Notice: 忽略某些项目
		if p.Name == "test" || len(*repos) == 0 {
			continue
		}
		for _, repo := range *repos {
			old := fmt.Sprintf("old.hub/%v", repo.Name)
			new := fmt.Sprintf("new.hub/%v", repo.Name)
			piddata[old] = new
		}
	}
	data, err := json.Marshal(piddata)
	if err != nil {
		panic(err)
	}
	if err := Writefile("image.json", string(data)); err != nil {
		panic(err)
	}
}

// Writefile 写文件
func Writefile(logpath, msg string) (err error) {
	file, err := os.OpenFile(logpath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(msg)
	write.Flush()
	return nil
}