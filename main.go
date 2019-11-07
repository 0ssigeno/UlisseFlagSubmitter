package main

import (
	"./execExploits"
	"./interfaceExploit"
	"./notifier"
	"./submit"
	"fmt"
	"os"
	"strconv"
	"time"
)

//se il flusso principale quitta, quittano pure le gooroutine

func LoopPrintExploits() {
	tick := 0
	for {
		fmt.Printf("Print number %d \n", tick)
		fmt.Println(interfaceExploit.PPExploits())
		time.Sleep(time.Duration(interfaceExploit.TimePrint) * time.Second)
		tick += 1
	}
}

func manageInput() {
	var command string
	var name string
	var team string
	_, err := fmt.Scanln(&command, &name, &team)
	if err != nil {
		//Stampare un bell'errore :D
		fmt.Print("Required: command name team")
		os.Exit(2)
	}
	switch command {
	case interfaceExploit.COMMANDS:
		fmt.Println("> start nomeExploit")
		fmt.Println("> stop nomeExploit")
		fmt.Println("> add nomeExploit numTeam")
		fmt.Println("> remove nomeExploit numTeam")
		fmt.Println("> time  exploit/submit/print/timeout valore")
	case interfaceExploit.STOP:
		index, name := interfaceExploit.GetExploit(name)
		execExploits.StopExploit(index)
		fmt.Printf("> Exploit %s Stopped\n", name)
	case interfaceExploit.START:
		index, name := interfaceExploit.GetExploit(name)
		execExploits.StartExploit(index)
		fmt.Printf("> Exploit %s Started\n", name)
	case interfaceExploit.ADDG:
		indexExp, name := interfaceExploit.GetExploit(name)
		execExploits.AddTeam(indexExp, team)
		fmt.Printf("> Exploit %s Added team %s\n", name, team)
	case interfaceExploit.REMOVEG:
		indexExp, name := interfaceExploit.GetExploit(name)
		execExploits.RemoveTeam(indexExp, team)
		fmt.Printf("> Exploit %s Removed team %s\n", name, team)
	case interfaceExploit.CHANGETIME:
		switch name {
		case "exploit":
			interfaceExploit.TimeExploit, _ = strconv.Atoi(team)
			fmt.Printf("> Changed %s time to %s\n", name, team)
		case "submit":
			interfaceExploit.TimeSubmit, _ = strconv.Atoi(team)
			fmt.Printf("> Changed %s time to %s\n", name, team)
		case "print":
			interfaceExploit.TimePrint, _ = strconv.Atoi(team)
			fmt.Printf("> Changed %s time to %s\n", name, team)
		case "timeout":
			interfaceExploit.TimeTimeout, _ = strconv.Atoi(team)
			fmt.Printf("> Changed %s time to %s\n", name, team)
		default:
			fmt.Println("> Command not valid")

		}
	default:
		fmt.Println("> Command not valid")
	}
}

func main() {
	fmt.Printf("Exploit executed every %d seconds \n", interfaceExploit.TimeExploit)
	fmt.Printf("Submit executed every %d seconds \n", interfaceExploit.TimeSubmit)
	fmt.Printf("Print executed every %d seconds \n", interfaceExploit.TimePrint)
	fmt.Printf("Timeout for exploits after %d seconds \n", interfaceExploit.TimeTimeout)
	go execExploits.Loop()
	go submit.Loop(interfaceExploit.HTTP) //TODO modificare HTTP con il protocollo corretto
	go LoopPrintExploits()
	go notifier.Loop()
	for {
		manageInput()
	}
}
