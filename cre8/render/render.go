package render

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFTokern      string
	Secure          bool
	Port            string
	ServerName      string
}

// PageTmpl check what is template engine you choose
// เพื่อตรวจสอบว่าใช้ เทมเพลทเอนจิ้นอะไร เช่น go หรือ jet
func (rd *Render) PageTmpl(w http.ResponseWriter, r *http.Request, view string, varialbes, data interface{}) error {
	switch strings.ToLower(rd.Renderer) {
	case "go":
		return rd.GoTmpl(w, r, view, data)
	case "jet":
		return rd.JetTmpl(w, r, view, data)
	}
	return nil
}

func (rd *Render) GoTmpl(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.html", rd.Renderer, view))
	if err != nil {
		return err
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	if err := tmpl.Execute(w, &td); err != nil {
		return err
	}

	return nil
}

func (rd *Render) JetTmpl(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {

	return nil
}
