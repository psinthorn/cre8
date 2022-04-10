package render

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Render struct เพื่อเก็บค่าต่างๆ ที่สำคัญที่ต้องกำหนดและใช้ในการ Render
type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
}

// TemplateData ใช้เพื่อรับค่าต่างๆ ที่ต้องการใช้ในเทมเพลท
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

// GoTmpl สร้างเทมเพลทโดยใช้ go Template standard library
func (rd *Render) GoTmpl(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {

	// สร้างเทมเพลทจากไฟล์
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.html", rd.RootPath, view))
	if err != nil {
		return err
	}
	// สร้างตัวแปรรับ-ส่งข้อมูลที่ต้องการใช้ในเทมเพลทโดยส่งผ่าน temaplateData struct ที่ได้สร้างไว้
	td := &TemplateData{}
	// หากตัวแปร data มีข้อมูลมาด้วย ให้กำหนดให้ตัวแปร td
	if data != nil {
		td = data.(*TemplateData)
	}

	// ทำการ execute template หากมีข้อผิดพลาดให้คืนค่า err หากไม่มีคือนค่า nil
	if err := tmpl.Execute(w, &td); err != nil {
		return err
	}

	return nil
}

func (rd *Render) JetTmpl(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {

	return nil
}
