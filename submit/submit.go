package submit

import (
	"../interfaceExploit"
	"./httpSubmit"
	"fmt"
	"sync"
)

//Mappa di handler
var handler map[string]func(int, int, int)

//Invocazione della funzione associata all'handler
func submitSingleFlag(typeS string, indexFlag int, indexTeam int, indexExploit int) {
	handler[typeS](indexFlag, indexTeam, indexExploit)
}

//registrazione degli handler
func register() {
	handler[interfaceExploit.HTTP] = _handleHTTP
}

func submitSingleTeam(typeS string, indexTeam int, indexExploit int) {
	for indexFlag, flag := range interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flag {
		if flag.Status == "NEW" || flag.Status == "ERROR" {
			submitSingleFlag(typeS, indexFlag, indexTeam, indexExploit)
		} else {
			//break //non le guardo veramente tutte
		}
	}

}
func submit(typeS string, indexExploit int) {

	//Producer
	teams := make(chan int, 20)
	go func() {
		for indexTeam, _ := range interfaceExploit.Exploits[indexExploit].Teams {
			teams <- indexTeam
		}
	}()

	//Creo il gruppo
	submitters := 4
	wg := sync.WaitGroup{}
	wg.Add(submitters)

	//Partono i submitters
	for i := 0; i < submitters; i++ {
		go func() {
			defer wg.Done()
			for indexTeam := range teams {
				//TODO capire questo controllo
				if interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flag[0].Status == "NEW" ||
					interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flag[0].Status == "ERROR" {

					submitSingleTeam(typeS, indexTeam, indexExploit)
				}
			}
		}()
	}
	wg.Wait()
}

func Loop(typeSubmit string) {
	handler = make(map[string]func(int, int, int))
	register()
	for {
		fmt.Println("Executing Submits")
		for indexExploit, exploit := range interfaceExploit.Exploits {
			if exploit.Active {
				submit(typeSubmit, indexExploit)
			}
		}
	}
}

func _handleHTTP(indexFlag int, indexTeam int, indexExploit int) {
	url := interfaceExploit.UrlSubmit
	flag := interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flag[indexFlag].Flag
	status := httpSubmit.SubmitFlagHttp(flag, url)
	interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flag[indexFlag].Status = status
	if status == "OK" {
		interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flags += 1
		interfaceExploit.Exploits[indexExploit].Flags += 1
	}
}
