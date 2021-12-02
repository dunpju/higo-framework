package errcode

import (
	"fmt"
	"github.com/dengpju/higo-utils/utils"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"text/template"
)

var relativeWd = "/app/errcode"

type CodeBuilder struct {
	Name         string
	Len          int
	FuncName     string
	KeyValueDocs []utils.KeyValueDoc
}

func NewCodeBuilder(name string) *CodeBuilder {
	return &CodeBuilder{Name: name}
}

func (this *CodeBuilder) parse(doc string) *CodeBuilder {
	messagePattern := `(@Message\(").*?("\))`
	messageRegexp := regexp.MustCompile(messagePattern)
	if messageRegexp == nil {
		log.Fatalln("regexp err")
	}
	messages := messageRegexp.FindAllStringSubmatch(doc, -1)
	message := make([]string, 0)
	for _, msg := range messages {
		msg0 := strings.Replace(msg[0], "\n", "", -1)
		msg0 = strings.Replace(msg0, `@Message("`, "", -1)
		msg0 = strings.Replace(msg0, `")`, "", -1)
		message = append(message, msg0)
	}

	codePattern := `(const\s+).*?(;)`
	codeRegexp := regexp.MustCompile(codePattern)
	if codeRegexp == nil {
		log.Fatalln("regexp err")
	}
	codes := codeRegexp.FindAllStringSubmatch(doc, -1)
	codd := make([]utils.KeyValueDoc, 0)
	for i, cod := range codes {
		cod0 := strings.Replace(cod[0], "\n", "", -1)
		cod0 = strings.Replace(cod0, `const `, "", -1)
		cod0 = strings.Replace(cod0, `;`, "", -1)
		codSplit := strings.Split(cod0, "=")
		key := utils.CaseToCamel(strings.ToLower(strings.Trim(codSplit[0], "")))
		value := strings.ToLower(strings.Trim(codSplit[1], ""))
		codd = append(codd, *utils.NewKeyValueDoc(key, value, message[i]))
	}
	this.KeyValueDocs = codd
	return this
}

func (this *CodeBuilder) generate() {
	keyValueDocs := this.KeyValueDocs
	this.Len = len(this.KeyValueDocs) - 1
	if len(keyValueDocs) > 0 {
		wd, err := os.Getwd()
		if nil != err {
			panic(err)
		}
		OutDir := wd + relativeWd
		utils.Dir(OutDir).Create()
		File := OutDir + "/" + strings.Replace(keyValueDocs[0].Value.(string), " ", "", -1) + ".go"
		this.FuncName = "code" + strings.Replace(keyValueDocs[0].Value.(string), " ", "", -1)
		if utils.FileExist(File) {
			log.Println(File + " already existed")
			return
		}
		outFile := utils.NewFile(File, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
		defer outFile.Close()
		tpl := this.Template("code.tpl")
		tmpl, err := template.New("code.tpl").Parse(tpl)
		if err != nil {
			panic(err)
		}
		//生成
		err = tmpl.Execute(outFile.File(), this)
		if err != nil {
			panic(err)
		}
		fmt.Println("code file " + File + " generate success!")
	}
}

func (this *CodeBuilder) Template(tplfile string) string {
	return tpl(tplfile)
}

func tpl(tplfile string) string {
	_, file, _, _ := runtime.Caller(0)
	file = path.Dir(file) + utils.PathSeparator() + tplfile
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	context, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(context)
}

var funcNames []string

func Builder() {
	if len(funcNames) == 0 {
		wd, err := os.Getwd()
		if nil != err {
			panic(err)
		}
		files := utils.Dir(wd + relativeWd + "/md").Suffix("md").Scan().Get()
		for _, filePath := range files {
			fileContext := string(utils.ReadFile(filePath).ReadAll())
			build := NewCodeBuilder("CodeErrorCode").parse(fileContext)
			build.generate()
			funcNames = append(funcNames, build.FuncName)
		}
		if len(funcNames) > 0 {
			NewAutoload(funcNames).generate()
		}
	}
}

type Autoload struct {
	FuncNames []string
}

func NewAutoload(funcNames []string) *Autoload {
	return &Autoload{FuncNames: funcNames}
}

func (this *Autoload) generate() {
	wd, err := os.Getwd()
	if nil != err {
		panic(err)
	}
	OutDir := wd + relativeWd
	utils.Dir(OutDir).Create()
	File := OutDir + "/" + "Autoload.go"
	outFile := utils.NewFile(File, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	defer outFile.Close()
	tpl := this.Template("autoload.tpl")
	tmpl, err := template.New("autoload.tpl").Parse(tpl)
	if err != nil {
		panic(err)
	}
	//生成
	err = tmpl.Execute(outFile.File(), this)
	if err != nil {
		panic(err)
	}
	fmt.Println("Autoload file " + File + " generate success!")
}

func (this *Autoload) Template(tplfile string) string {
	return tpl(tplfile)
}
