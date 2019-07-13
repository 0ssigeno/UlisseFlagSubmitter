package httpSubmit
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"strings"
	"../../interfaceExploit"
	"crypto/tls"
)

func createEncode(flag string)[]byte{
	encode:=fmt.Sprintf("{\"flags\" : [\"%s\"]}",flag)
	return []byte(encode)
}

func checkResponseHttp(resp *http.Response) string{
	if (strings.Contains(resp.Status,"200")){ //TODO
		bodyy, _ := ioutil.ReadAll(resp.Body)
		body:=string(bodyy)
		fmt.Println(body)
		if strings.Contains(body,"ok"){
			return "OK"
		}else if strings.Contains(body,"old"){
			return "OLD"
		}else if strings.Contains(body,"invalid"){
			return "INVALID"
		}else if strings.Contains(body,"own"){
			return "OWN"
		}else {
			return "CANT PARSE RESULT"
		}
		return "OK"
	} else if(strings.Contains(resp.Status,"400")){
		return "MYPROBLEM"
	} else {
		return "SERVERPROBLEM"
	}
}


func postRequest(myJson []byte, url string ) *http.Response{
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(myJson))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(interfaceExploit.Username, interfaceExploit.Password)
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
