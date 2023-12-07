package main

import (
	"fmt"
	"net/http"
	"os"
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
	a := location.Scheme + "://" + location.Host + location.Path
	fmt.Println(a)
	a = location.Scheme + "://vscode.cdn.azure.cn" + location.Path
	fmt.Println(a)
}
