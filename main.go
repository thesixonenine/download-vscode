package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

// main download vscode use cdn at cn
// https://code.visualstudio.com/sha/download?build=stable&os=win32-x64-archive
// https://az764295.vo.msecnd.net/stable/1a5daa3a0231a0fbba4f14db7ec463cf99d7768e/VSCode-win32-x64-1.84.2.zip
// https://vscode.cdn.azure.cn/stable/1a5daa3a0231a0fbba4f14db7ec463cf99d7768e/VSCode-win32-x64-1.84.2.zip
// vscode#1.84.2#https_update.code.visualstudio.com_1.84.2_win32-x64-archive_stable_dl.7z
func main() {
	fmt.Println(os.Args)
	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	resp, _ := client.Get("https://code.visualstudio.com/sha/download?build=stable&os=win32-x64-archive")
	location, _ := resp.Location()
	host := location.Host
	queryPath := location.Path
	originUrl := location.Scheme + "://" + host + queryPath
	cdnUrl := location.Scheme + "://vscode.cdn.azure.cn" + queryPath
	split := strings.Split(queryPath, "/")
	fileName := split[len(split)-1]
	ext := path.Ext(fileName)
	fileN := strings.TrimSuffix(fileName, ext)
	fmt.Printf("最新文件名:\n%s|%s\n", fileN, fileName)
	fmt.Printf("原始下载链接:\n%s\n", originUrl)
	fmt.Printf("CDN下载链接:\n%s\n", cdnUrl)
	// 开始下载文件
	resp, err := http.Get(cdnUrl)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if resp.StatusCode == 404 {
		fmt.Printf("CDN还未缓存文件[%s],是否继续使用原始链接下载?[y/n] ", fileName)
		flag := ""
		_, err := fmt.Scanln(&flag)
		if err != nil {
			log.Fatalf("输入错误: %s", err.Error())
		}
		flag = strings.TrimSpace(flag)
		if strings.EqualFold(flag, "y") {
			resp, err = http.Get(originUrl)
		} else {
			return
		}
	}
	if resp.StatusCode != 200 {
		log.Fatalf("文件[%s]下载失败,code[%d],msg[%s]", fileName, resp.StatusCode, resp.Status)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatalf(err.Error())
		}
	}()
	out, err := os.Create(fileName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Fatalf(err.Error())
		}
	}()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func ZIPTo7z(fn string) {
	format := "mkdir tmp && Copy-Item FILE_NAME.zip tmp && cd tmp && 7z x FILE_NAME.zip && Remove-Item FILE_NAME.zip && 7z a -t7z FILE_NAME.7z * && Copy-Item FILE_NAME.7z ../ && cd .. && Remove-Item -Force -Recurse tmp"
	command := strings.ReplaceAll(format, "FILE_NAME", fn)
	exec.Command(command)
}
