package hconfig

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

// 应用配置
type App struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`

	Host string `json:"host"`
	Port int    `json:"port"`

	LogLevel string `json:"logLevel"`
}

// 启动服务
func (app *App) Start(preFunc func(app *App)) {
	log.Println("******************************************************************************************")
	defer printStack()
	app.Args()
	if preFunc != nil {
		preFunc(app)
	}
	app.ListenAndServe()
}

// 加载命令行参数
func (app *App) Args() {
	var (
		// 服务名与端口号
		name string
		port int
	)
	flag.StringVar(&name, "name", app.Name, fmt.Sprintf("Set Application name. Default '%s'", app.Name))
	flag.IntVar(&port, "port", app.Port, fmt.Sprintf("Set Port. Default is %d", app.Port))

	flag.Parse()

	app.Name = name
	app.Port = port
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	log.Printf("==> %s\n", string(buf[:n]))

	if err := recover(); err != nil {
		log.Fatalf("** Main Fatalf-> %s", err)
	}
}

func (app *App) ListenAndServe() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		w.Write([]byte("MicroService:" + app.Name))
	})
	log.Printf("MicroService: %s  ListenAndServe %s:%d   Start server http://127.0.0.1:%d", app.Name, app.Host, app.Port, app.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Host, app.Port), nil))
}
