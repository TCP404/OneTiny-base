package util

import (
	"fmt"
	"strings"

	"github.com/TCP404/OneTiny-base/core/model"
)

// TODO 对字符串拼接进行优化，推荐采用 strings.Builder 并提前设置好 Grow()
// https://juejin.cn/post/6844903713241301006

// GenerateIndexHTML 生成首页HTML内容，包括 head、fileList、
func GenerateIndexHTML(files []model.FileStruction, pathTitle string, IsAllowUpload bool) string {
	headHTML := indexHead(pathTitle)
	if IsAllowUpload {
		headHTML += upload()
	}
	fileListHTML := fileList(files)
	tailHTML := `<br /><hr /></body></html>`
	return headHTML + fileListHTML + tailHTML
}

func indexHead(pathTitle string) string {
	head := []string{
		`<!DOCTYPE html>`,
		`<html lang="zh-CN">`,
		`<head>`,
		`<meta charset="UTF-8">`,
		`<link rel="icon" href="data:image/ico;base64,aWNv">`,
		`<title>OneTiny Server</title>`,
		`</head>`,
		`<body>`,
		`<h1 style="display:inline;">Directory listing for `,
		pathTitle,
		`</h1><hr /><br />`,
	}
	return strings.Join(head, "")
}

func upload() string {
	form := []string{
		`<div style="position: absolute;right: 30px;top: 10px;font-size: 24px;width: 500px;border: 1px solid #000;padding: 7px;border-radius: 8px;">`,
		`<form action="/file/upload" method="post" enctype="multipart/form-data">`,
		`<input type="file" name="upload_file" style="width: 400px; float: left;">`,
		`<input type="submit" value="上传" style="float: right;">`,
		`</form>`,
		`</div>`,
	}
	return strings.Join(form, "")
}

func fileList(files []model.FileStruction) string {
	tableHead := `<table style="width:100%">`
	tHead := `<thead><tr><td style='width:30%'>文件名</td><td style="width:100px;text-align:center;">文件大小</td><td></td></tr></thead>`
	tBody := "<tbody>"
	tBody += "<tr><td><a href='../'> &nbsp;. . /</a></td></tr>"
	for i, f := range files {
		link := ""
		trHead := "<tr>"
		fileLink := fmt.Sprintf("<td>%d. <a href='%s'> %s </a></td>", i, f.URLRelPath, f.Name)
		fileSize := fmt.Sprintf("<td style='text-align:right'>%s</td>", f.Size)
		trTail := "</tr>"
		link = trHead + fileLink + fileSize + trTail
		tBody += link
	}
	tBody += `</tbody>`
	tableTail := `</table>`
	return tableHead + tHead + tBody + tableTail
}
