// index.go
package main

import (
	"fmt"
	"github.com/chownplusx/web"
	"strconv"
)

func IndexLoadGet(ctx *web.Context, val string) string {
	content := ""
	if ctx.Params["start"] == "" {
		content = PostsInRange(0, 5)
	} else {
		start, _interr := strconv.Atoi(ctx.Params["start"])
		if _interr != nil {
			fmt.Println(_interr)
		}
		content = PostsInRange(start, start+5)
	}
	mapping := map[string]string{"title": "Proggin: Index", "content": content}
	return Loadmustache("frame.mustache", &mapping)
}

func GetSinglePost(val string) string {
	content := GetPostByName(val)
	mapping := map[string]string{"title": "Proggin: Index", "content": content}
	return Loadmustache("frame.mustache", &mapping)
}
