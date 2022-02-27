package main  
 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
)

type QuestionAndAnswer struct {
	Context string `json:"context"`
	Question string `json:"question"`
}

func saveQuestionAndAnswerQuestion(filename string, inputContext string, inputQuestion string) {
	data := &QuestionAndAnswer{
		Context: inputContext,
		Question: inputQuestion,
	}

	buf, err1 := json.Marshal(data)
	if err1 !=nil {
		panic(err1)
	}
	inputFilePath := "/home/ubuntu/server_inputs/" + filename
	err2 := ioutil.WriteFile(inputFilePath, buf, 0644)
	if err2 !=nil {
		panic(err2)
	}
}

func runAndGetQuestionAndAnswerResult(filename string) (string) {
	cmd := exec.Command("python3", "/home/ubuntu/gogogo/python_backend/run_bert_answers_questions.py", "-f", filename, "-o", filename)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err1 := cmd.Run()
	if err1 != nil {
		fmt.Println(fmt.Sprint(err1) + ": " + stderr.String())
		return ""
	}

	outputFilePath := "/home/ubuntu/server_outputs/" + filename
	jsonFile, err2 := os.Open(outputFilePath)
    // if we os.Open returns an error then handle it
    if err2 != nil {
        fmt.Println(err2)
		return ""
    }
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var answers map[string]string
    json.Unmarshal([]byte(byteValue), &answers)
	return answers["1"]
}

func main() {
	saveQuestionAndAnswerQuestion(
		"0.json",
		"Ontario is reporting 1,003 people hospitalized with COVID-19 on Friday, marking the lowest number the province has seen since last December — the beginning of the Omicron wave. The number of people in intensive care units also dropped, with the province reporting 297 people in ICUs. That marks the first time that figure has been under 300 since Jan. 5. The province is set to lift its proof of vaccination requirements next week as part of Ontario's reopening plan due to positive trends in public health indicators but the province's top doctor said mask requirements will remain in place for the time being. Chief Medical Officer of Health Dr. Kieran Moore said at a Thursday news briefing that masks remain 'an important tool in our tool box' when it comes to reducing transmission of the virus. Moore said given positive trends in public health indicators and the province's high vaccination rate, health officials are actively reviewing all directives to health-care providers, and hope to provide an update in the coming weeks. The number of hospitalizations reported Friday is down from 1,066 the day before and from 1,281 at the same time last week.",
		"Why are vaccination requirements lifted?",
	)
	answer := runAndGetQuestionAndAnswerResult("0.json")
	fmt.Println(answer)
}