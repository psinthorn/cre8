package cre8

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/psinthorn/cre8/render"
)

const version = "1.0.0"

// New เพื่อตรวจสอบ folder path ที่ทำงานอยู่ให้ปัจจุบัน โดยจะคืนค่ากลับมาเป็น string
func (c *Cre8) New(rootPath string) error {
	folderConfig := initFolders{
		rootPath:    rootPath,
		folderNames: []string{"models", "views", "controllers", "data", "public", "static", "middleware", "logs", "tmp"},
	}

	if err := c.Init(folderConfig); err != nil {
		return err
	}

	// check .env file is exists if not create new file
	if err := c.checkDotEnv(rootPath); err != nil {
		return err
	}

	// read .env file by using joho/godotenv lib
	if err := godotenv.Load(rootPath + "/.env"); err != nil {
		return nil
	}

	// create log
	infoLog, errorLog := c.startLoggers()
	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.RootPath = rootPath
	c.Routes = c.routes().(*chi.Mux)

	c.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	views := jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)
	c.JetViews = views
	c.createRender()

	return nil
}

func (c *Cre8) Init(fd initFolders) error {
	rootPath := fd.rootPath
	for _, folderName := range fd.folderNames {
		err := c.createFolderIfNotExist(rootPath + "/" + folderName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cre8) ListenAndServe() {
	serv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     c.ErrorLog,
		Handler:      c.Routes,
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	c.InfoLog.Printf("Listen and Serve on port: %s", c.config.port)
	err := serv.ListenAndServe()
	c.ErrorLog.Fatal(err)
}

// checkDotEnv check is .ENV file is available
func (c *Cre8) checkDotEnv(rootPath string) error {
	err := c.createFileIfNotExists(fmt.Sprintf("%s/.env", rootPath))
	if err != nil {
		return err
	}
	return nil
}

func (c *Cre8) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}

func (c *Cre8) createRender() {
	pageRender := render.Render{
		Port:     c.config.port,
		Renderer: c.config.renderer,
		RootPath: c.RootPath,
		JetViews: c.JetViews,
	}

	c.Render = &pageRender
}
