// utils.go
package main

import (
	"fmt"
	"github.com/chownplusx/mustache"
	"github.com/chownplusx/web"
	md "github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

type Files []os.FileInfo

type ByModTime struct{ Files }

func (s Files) Len() int {
	return len(s)
}

func (s Files) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByModTime) Less(i, j int) bool {
	return s.Files[i].ModTime().Unix() > s.Files[j].ModTime().Unix()
}

func SortByModified(files []os.FileInfo) []os.FileInfo {
	sort.Sort(ByModTime{files})
	return files
}

func ReverseOrder(a []os.FileInfo) []os.FileInfo {
	//This reverses the order of the posts
	//But you can't tell just by looking ;)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func PostsInRange(start int, end int) string {
	filelist, _err := ioutil.ReadDir("posts")
	if _err != nil {
		fmt.Println(_err)
	}
	SortByModified(filelist)
	content := ""
	for i := start; i < end; i++ {
		name := filelist[i].Name()

		// This formats the date nice and pretty
		date := filelist[i].ModTime().Format("2006-01-02 15:04")
		post, _post := ioutil.ReadFile(path.Join("posts", filelist[i].Name()))
		if _post != nil {
			fmt.Println(_post)
		}
		permalink := "<a href=\"/post/" + name + "\">Permalink</a>"
		content += Loadmustache("perpost.mustache", &map[string]string{"title": name,
			"content": string(md.MarkdownCommon([]byte(post))),
			"details": "Posted at " + date + " " + permalink})
	}

	return content
}

func GetPostByName(name string) string {
	post, _post := ioutil.ReadFile(path.Join("posts", name))
	FileInfo, _file := os.Stat(path.Join("posts", name))
	if _post != nil {
		fmt.Println(_post)
	}
	if _file != nil {
		fmt.Println(_file)
	}
	date := FileInfo.ModTime().Format("2006-01-02 15:04")

	return Loadmustache("perpost.mustache", &map[string]string{"title": name,
		"content": string(md.MarkdownCommon([]byte(post))),
		"details": "Posted at " + date})
}

func Loadmustache(filename string, args *map[string]string) string {
	file, _err := ioutil.ReadFile("Mst/" + filename)
	if _err != nil {
		fmt.Println(_err)
		return "File not found"
	}
	data := mustache.Render(string(file), args)
	return data
}

func Sendstatic(ctx *web.Context, val string) {
	file, _err := ioutil.ReadFile("static/" + val)
	if _err != nil {
		return
	}
	filetype := strings.Split(val, ".")
	ctx.ContentType(filetype[len(filetype)-1])
	ctx.WriteString(string(file))
}
