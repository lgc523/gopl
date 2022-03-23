package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	f1 := &Friend{FName: "蔡流进", Job: "记者"}
	f2 := &Friend{FName: "陈夏京", Job: "气象厅总厅预报组二组科长"}
	p0 := &Person{"李时雨", []*Friend{f1, f2}}
	p1 := &Person{"韩气峻", []*Friend{f2, f1}}

	text := `<气象厅的人们>{{- /* comment */ -}}{{range .}}
男主{{.UserName -}} 的女朋友{{ range .Friends }}
- {{.FName}}: {{.Job -}}
{{- end -}}
{{$chp:=11}}
2022-03-22,我看到了第{{$chp}}集。
{{- end -}}`

	_ = `<气象厅的人们>{{range .}}
男主{{.UserName}} 的女朋友{{ with .Friends }} {{- range . }}
- {{ .FName }}: {{.Job}}
{{- end }}
{{- end}}
{{end}}
netflix-2022 搞笑浪漫爱情韩剧...
`

	tmpl, err := template.New("range_template").Parse(text)
	if err != nil {
		log.Fatalf("template parse err:%v", err)
	}
	err = tmpl.Execute(os.Stdout, []*Person{p0, p1})
	if err != nil {
		log.Fatalln(err)
	}

}

type Friend struct {
	FName string
	Job   string
}

type Person struct {
	UserName string
	Friends  []*Friend
}
