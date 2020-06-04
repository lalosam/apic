package apiccore

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type htmlTag struct {
	layout string
	style  string
}

//ToHTML return HTML representation of the given object
func ToHTML(obj interface{}, tag string) string {
	log.Println("Generating HTML:")
	var html strings.Builder
	html.WriteString("<html><head>\n")
	html.WriteString("<link rel=\"stylesheet\" href=\"/static/apic.css\">")
	html.WriteString("</head><body>\n")
	processObject(obj, tag, &html)
	html.WriteString("\n</body>")
	html.WriteString("\n</html>")
	return html.String()
}

func processObject(obj interface{}, tag string, buff *strings.Builder) {
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
		buff.WriteString(valueO.String())
	case reflect.Int64:
		buff.WriteString(fmt.Sprintf("%d", valueO.Int()))
	case reflect.Float64:
		buff.WriteString(fmt.Sprintf("%.2f", valueO.Float()))
	default:
		log.Printf("Not Supported KIND [%v]", kind)
	}
}

func processStruct(obj interface{}, tag string, buff *strings.Builder) {
	valueO := reflect.ValueOf(obj)
	typeO := reflect.TypeOf(obj)
	switch tag {
	case "DIV":
	case "ROW":
		log.Printf("Processing Struct Row, TAG:[%s] Fields:[%d]", tag, typeO.NumField())
		buff.WriteString("<tr>")
		for i := 0; i < typeO.NumField(); i++ {
			buff.WriteString("<td>")
			htmlTag := parseTag(typeO.Field(i).Tag.Get("html"))
			log.Println(htmlTag)
			processObject(valueO.Field(i).Interface(), htmlTag.layout, buff)
			buff.WriteString("</td>")
		}
		buff.WriteString("</tr>")
	case "TABLE":
		buff.WriteString("<div style=\"overflow-x:auto;\"><table class=\"struct\">")
		buff.WriteString("<tbody>")
		for i := 0; i < typeO.NumField(); i++ {
			ft := typeO.Field(i)
			buff.WriteString("<tr><td class=\"property_name\">")
			buff.WriteString(ft.Name)
			buff.WriteString("</td><td class=\"property_value\">")
			htmlTag := parseTag(typeO.Field(i).Tag.Get("html"))
			log.Println(htmlTag)
			processObject(valueO.Field(i).Interface(), htmlTag.layout, buff)
			buff.WriteString("</td></tr>")
		}
		buff.WriteString("</tbody>")
		buff.WriteString("</table></div>")
	}

}

func processSlice(obj interface{}, tag string, buff *strings.Builder) {
	value := reflect.ValueOf(obj)
	switch tag {
	case "DIV":
		buff.WriteString("<div class=\"slice_main\" style=\"overflow-x:auto;\">")
		for i := 0; i < value.Len(); i++ {
			buff.WriteString("<div class=\"slice_item\" style=\"overflow-x:auto;\">")
			processObject(value.Index(i).Interface(), "ROW", buff)
			buff.WriteString("</div>")
		}
		buff.WriteString("</div>")
	case "TABLE":
		buff.WriteString("<div style=\"overflow-x:auto;\"><table class=\"listing\">")
		if value.Len() > 0 {
			headers := getTableHeaders(value.Index(0).Interface())
			buff.WriteString(headers)
		}
		buff.WriteString("<tbody>")
		for i := 0; i < value.Len(); i++ {
			processObject(value.Index(i).Interface(), "ROW", buff)
		}
		buff.WriteString("</tbody>")
		buff.WriteString("</table></div>")
	}
}

func convertToHTML(obj interface{}, openTag string, closeTag string, buffer *strings.Builder) {
	value := reflect.ValueOf(obj)
	kind := value.Kind()
	buffer.WriteString(openTag)
	switch kind {
	case reflect.Slice:
		buffer.WriteString("<table>")
		if value.Len() > 0 {
			headers := getTableHeaders(value.Index(0).Interface())
			buffer.WriteString(headers)
		}
		for i := 0; i < value.Len(); i++ {
			convertToHTML(value.Index(i), "<tr>", "</tr>", buffer)
		}
		buffer.WriteString("</table>")
	case reflect.Struct:
		buffer.WriteString("<td>STRUCT</td")
	}
	buffer.WriteString(closeTag)
}

func parseTag(tag string) htmlTag {
	log.Printf("Processing html tag [%s]\n", tag)
	var res = htmlTag{}
	for _, param := range strings.Split(tag, ";") {
		keyvalue := strings.Split(param, ":")
		switch keyvalue[0] {
		case "layout":
			res.layout = keyvalue[1]
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
			b.WriteString("<th>")
			b.WriteString(ft.Name)
			b.WriteString("</th>")
		}
		b.WriteString("</tr></thead>")
		return b.String()
	default:
		panic("Not supported")
	}
}
