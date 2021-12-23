// package main
/*
首先定义了类型HandlerFunc，这是提供给框架用户的，用来定义路由映射的处理方法。我们在Engine中，添加了一张路由映射表router，key 由请求方法和静态路由地址构成，例如GET-/、GET-/hello、POST-/hello，这样针对相同的路由，如果请求方法不同,可以映射不同的处理方法(Handler)，value 是用户映射的处理方法。

当用户调用(*Engine).GET()方法时，会将路由和处理方法注册到映射表 router 中，(*Engine).Run()方法，是 ListenAndServe 的包装。

Engine实现的 ServeHTTP 方法的作用就是，解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND 。
 */

package main

import (
	"fmt"
	"log"
	"net/http"

	"gee_cache"
)

// gee-web测试
//type student struct {
//	Name string
//	Age  int8
//}
//
//func FormatAsDate(t time.Time) string {
//	year, month, day := t.Date()
//	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
//}
//
//func main() {
//	r := gee_web.Default()
//	r.SetFuncMap(template.FuncMap{
//		"FormatAsDate": FormatAsDate,
//	})
//	r.LoadHTMLGlob("templates/*")
//	r.Static("/assets", "./static")
//
//	stu1 := &student{Name: "Geektutu", Age: 20}
//	stu2 := &student{Name: "Jack", Age: 22}
//	r.GET("/", func(c *gee_web.Context) {
//		c.HTML(http.StatusOK, "css.tmpl", nil)
//	})
//	r.GET("/students", func(c *gee_web.Context) {
//		c.HTML(http.StatusOK, "arr.tmpl", gee_web.H{
//			"title":  "gee",
//			"stuArr": [2]*student{stu1, stu2},
//		})
//	})
//
//	r.GET("/date", func(c *gee_web.Context) {
//		c.HTML(http.StatusOK, "custom_func.tmpl", gee_web.H{
//			"title": "gee",
//			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
//		})
//	})
//
//	// index out of range for testing Recovery()
//	r.GET("/panic", func(c *gee_web.Context) {
//		names := []string{"geektutu"}
//		c.String(http.StatusOK, names[100])
//	})
//
//	r.Run(":9999")
//}


// gee-cache 测试
var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	gee_cache.NewGroup("scores", 2<<10, gee_cache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:9999"
	peers := gee_cache.NewHTTPPool(addr)
	log.Println("geecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}