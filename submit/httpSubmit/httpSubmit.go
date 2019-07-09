package httpSubmit
import (
	"fmt"
	"net/http"
//	"io/ioutil"
	"bytes"
	"strings"
)

func createEncode(flag string)[]byte{
	encode:=fmt.Sprintf("{'flags' : ['%s']}",flag)
	return []byte(encode)
}

func checkResponseHttp(resp *http.Response) string{
	if (strings.Contains(resp.Status,"200")){ //TODO
		//body, _ := ioutil.ReadAll(resp.Body)
	
		return "OK"
	}else{
		return "ERROR"
	}
}

func postRequest(myJson []byte, url string ) *http.Response{
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(myJson))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp

}

func SubmitFlagHttp(flag string, url string)string{
	myJson:=createEncode(flag)
	resp:=postRequest(myJson,url)
	return checkResponseHttp(resp)
}
