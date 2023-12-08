# download-vscode

## 原因

由于使用Scoop安装/更新vscode的速度太慢所以创建此项目.

## 方案

为了提高下载速度, 使用CDN进行下载vscode的安装包.

## 特性

如果是下载的zip包,还可以转换成Scoop所需要的格式, 需要有7z命令(安装了Scoop的都有该命令). 在下载完成后将自动转换成Scoop需要的格式文件, 手动将其移到 `scoop\cache` 并重新执行 `scoop install` 即可
