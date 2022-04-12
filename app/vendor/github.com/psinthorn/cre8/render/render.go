package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
)

// Render struct เพื่อเก็บค่าต่างๆ ที่สำคัญที่ต้องกำหนดและใช้ในการ Render
type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
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
func (rd *Render) PageTmpl(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(rd.Renderer) {
	case "go":
		return rd.GoTmpl(w, r, view, data)
	case "jet":
		return rd.JetTmpl(w, r, view, variables, data)
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

func (rd *Render) JetTmpl(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	var vars jet.VarMap
	// ประกาศตัวแปร vars โดยมี type เป็น jet.VarMap หากไม่เป็น type jet.VarMap จะไม่สามารถใช้ตัวแปรได้บน JetTemplate (เป็นข้อกำหนดของ JetTemplate)
	//
	if variables == nil {
		// หากไม่มีข้อมูลให้สร้างข้อมูลเปล่าขึ้นมาโดยกำหนดให้เป็น jet.VarMap type
		vars = make(jet.VarMap)
	} else {
		// หาก variables มีข้อมูลให้ทำการแปลงข้อมูลให้อยู่ในรูปแบบข้อมูล jet.VarMap type
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}

	t, err := rd.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))
	if err != nil {
		log.Println(err)
		return err
	}

	if err := t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
