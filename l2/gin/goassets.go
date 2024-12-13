package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jessevdk/go-assets"
)

// go-assets（或类似库）在Go语言中嵌入静态资源（在这个例子中是HTML模板）的例子。
// 这种技术允许您在编译时将文件（如模板、图片、CSS等）嵌入到二进制文件中，从而在运行时无需依赖外部文件系统
// 这里定义了两个字符串变量，它们包含了HTML模板的内容
var (
	_Assetsbfa8d115ce0617d89507412d5393a462f8e9b003 = "<!doctype html>\n<body>\n  <p>Can you see this? → {{.Bar}}</p>\n</body>\n"
	_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!doctype html>\n<body>\n  <p>Hello, {{.Foo}}</p>\n</body>\n"
)

// // Assets变量是一个go-assets FileSystem的实例，它包含了嵌入的文件信息
// // 第一个map定义了文件系统的目录结构 / 第二个map定义了每个文件的具体信息，包括路径、权限、修改时间和内容
var Assets = assets.NewFileSystem(map[string][]string{"/": {"html"}, "/html": {"bar.tmpl", "index.tmpl"}}, map[string]*assets.File{
	"/": {
		Path:     "/",
		FileMode: 0x800001ed,                                 // 文件权限，通常是Unix风格的权限表示
		Mtime:    time.Unix(1524365738, 1524365738517125470), // 修改时间
		Data:     nil,                                        // 目录没有数据内容
	}, "/html": {
		Path:     "/html",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1524365491, 1524365491289799093),
		Data:     nil,
	}, "/html/bar.tmpl": {
		Path:     "/html/bar.tmpl",
		FileMode: 0x1a4, // 通常表示文件是可读的
		Mtime:    time.Unix(1524365491, 1524365491289611557),
		Data:     []byte(_Assetsbfa8d115ce0617d89507412d5393a462f8e9b003), // 文件内容
	}, "/html/index.tmpl": {
		Path:     "/html/index.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1524365491, 1524365491289995821),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	},
}, "")

func main() {
	r := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", gin.H{
			"Foo": "World",
		})
	})
	r.GET("/bar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/bar.tmpl", gin.H{
			"Bar": "World",
		})
	})
	r.Run(":8080")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
