package apiccore

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type element struct {
	elementType string
	elementName string
}

type htmlTag struct {
	layout string
	style  map[string]string
	elem   element
	width  string
}

//ToHTML return HTML representation of the given object
func ToHTML(obj interface{}, tag string) string {
	log.Println("Generating HTML:")
	var html strings.Builder
	html.WriteString("<html><head>\n")
	html.WriteString("<meta http-equiv=\"refresh\" content=\"180\">")
	html.WriteString("<link rel=\"stylesheet\" href=\"/static/apic.css\">")
	html.WriteString("</head><body>\n")
	processObject(obj, htmlTag{layout: tag}, &html)
	html.WriteString("\n</body>")
	html.WriteString("\n</html>")
	return html.String()
}

func processObject(obj interface{}, tag htmlTag, buff *strings.Builder) {
	valueO := reflect.ValueOf(obj)
	kind := valueO.Kind()
	switch kind {
	case reflect.Slice:
		log.Printf("Processing Slice, TAG:[%s]", tag)
		processSlice(valueO.Interface(), tag, buff)
	case reflect.Struct:
		log.Printf("Processing Struct, TAG:[%s]", tag)
		processStruct(valueO.Interface(), tag, buff)
	case reflect.Ptr:
		log.Printf("Processing PTR, TAG:[%s]", tag)
		processObject(valueO.Elem().Interface(), tag, buff)
	case reflect.String:
		processString(valueO.String(), tag, buff)
	case reflect.Int64:
		buff.WriteString(fmt.Sprintf("%d", valueO.Int()))
	case reflect.Float64:
		buff.WriteString(fmt.Sprintf("%.2f", valueO.Float()))
	default:
		log.Printf("Not Supported KIND [%v]", kind)
	}
}

func processString(s string, t htmlTag, buff *strings.Builder) {
	log.Printf("Processing a String TAG[%v]\n", t)
	if t.elem.elementType == "href" {
		buff.WriteString("<a href=\"")
		buff.WriteString(s)
		buff.WriteString("\">")
		buff.WriteString(t.elem.elementName)
		buff.WriteString("</a>")
	} else if t.style != nil {
		if val, ok := t.style[s]; ok {
			buff.WriteString("<div style=\"")
			buff.WriteString(val)
			buff.WriteString("\">")
			buff.WriteString(s)
			buff.WriteString("</div>")
		} else {
			buff.WriteString(s)
		}
	} else {
		buff.WriteString(s)
	}
}

func processStruct(obj interface{}, tag htmlTag, buff *strings.Builder) {
	valueO := reflect.ValueOf(obj)
	typeO := reflect.TypeOf(obj)
	switch tag.layout {
	case "DIV":
	case "ROW":
		log.Printf("Processing Struct Row, TAG:[%s] Fields:[%d]", tag, typeO.NumField())
		buff.WriteString("<tr>")
		for i := 0; i < typeO.NumField(); i++ {
			buff.WriteString("<td>")
			htmlTag := parseTag(typeO.Field(i).Tag.Get("html"))
			processObject(valueO.Field(i).Interface(), htmlTag, buff)
			buff.WriteString("</td>")
		}
		buff.WriteString("</tr>")
	case "TABLE":
		buff.WriteString("<div style=\"overflow-x:auto;\" class=\"div_struct_table\"><table class=\"struct\">")
		buff.WriteString("<tbody>")
		for i := 0; i < typeO.NumField(); i++ {
			ft := typeO.Field(i)
			buff.WriteString("<tr><td class=\"property_name\">")
			buff.WriteString(SplitByUpper(ft.Name) + ":")
			buff.WriteString("</td><td class=\"property_value\">")
			htmlTag := parseTag(typeO.Field(i).Tag.Get("html"))
			//log.Println(htmlTag)
			processObject(valueO.Field(i).Interface(), htmlTag, buff)
			buff.WriteString("</td></tr>")
		}
		buff.WriteString("</tbody>")
		buff.WriteString("</table></div>")
	}

}

func processSlice(obj interface{}, tag htmlTag, buff *strings.Builder) {
	value := reflect.ValueOf(obj)
	switch tag.layout {
	case "DIV":
		buff.WriteString("<div class=\"slice_main\" style=\"overflow-x:auto;\">")
		for i := 0; i < value.Len(); i++ {
			buff.WriteString("<div class=\"slice_item\" style=\"overflow-x:auto;\">")
			processObject(value.Index(i).Interface(), htmlTag{layout: "ROW"}, buff)
			buff.WriteString("</div>")
		}
		buff.WriteString("</div>")
	case "TABLE":
		if value.Len() > 0 {
			buff.WriteString("<div class=\"div_slice_table\" style=\"overflow-x:auto;\"><table class=\"listing\">")
			headers := getTableHeaders(value.Index(0).Interface())
			buff.WriteString(headers)
			buff.WriteString("<tbody>")
			for i := 0; i < value.Len(); i++ {
				processObject(value.Index(i).Interface(), htmlTag{layout: "ROW"}, buff)
			}
			buff.WriteString("</tbody>")
			buff.WriteString("</table></div>")
		} else {
			buff.WriteString("<div class=\"div_slice_nodata\">")
			buff.WriteString("No Data")
			buff.WriteString("</div>")
		}
	}
}

func parseTag(tag string) htmlTag {
	//log.Printf("Processing html tag [%s]\n", tag)
	var res = htmlTag{}
	for _, param := range strings.Split(tag, ";") {
		keyvalue := strings.Split(param, ":")
		switch keyvalue[0] {
		case "layout":
			res.layout = keyvalue[1]
		case "elem":
			e := strings.Split(keyvalue[1], "@")
			res.elem = element{e[0], e[1]}
		case "style":
			m := map[string]string{}
			for _, condition := range strings.Split(keyvalue[1], "|") {
				log.Println(condition)
				keyvalueCondition := strings.Split(condition, "@")
				log.Println(keyvalueCondition)
				m[keyvalueCondition[0]] = strings.ReplaceAll(strings.ReplaceAll(keyvalueCondition[1], "$$", ";"), "$", ":")
			}
			res.style = m
		case "width":
			res.width = keyvalue[1]
		}
	}
	return res
}

func getTableHeaders(obj interface{}) string {
	log.Printf("Received: %v\n", obj)
	valueO := reflect.ValueOf(obj)
	typeO := reflect.TypeOf(obj)
	log.Println(typeO)
	kind := valueO.Kind()
	log.Printf("Getting Headers from: %s", kind)
	switch kind {
	case reflect.Struct:
		var b strings.Builder
		b.WriteString("<thead><tr>")
		for i := 0; i < typeO.NumField(); i++ {
			ft := typeO.Field(i)
			htmlTag := parseTag(typeO.Field(i).Tag.Get("html"))
			style := ""
			if htmlTag.width != "" {
				style = fmt.Sprintf("style=\"width:%s\"", htmlTag.width)
			}
			b.WriteString("<th ")
			b.WriteString(style)
			b.WriteString(">")
			b.WriteString(SplitByUpper(ft.Name))
			b.WriteString("</th>")
		}
		b.WriteString("</tr></thead>")
		return b.String()
	default:
		panic("Not supported")
	}
}
