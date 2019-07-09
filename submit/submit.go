package submit
import (
	"../interfaceExploit"
	"./httpSubmit"
	"time"
	"fmt"
)
func submitSingleFlag(typeS string,indexFlag int,indexTeam int,indexExploit int){
	if(typeS == interfaceExploit.HTTP){
		url:=interfaceExploit.UrlSubmit
		flag:=interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flag[indexFlag].Flag
		status:=httpSubmit.SubmitFlagHttp(flag,url)
		interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flag[indexFlag].Status=status
		if status=="OK"{
			interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flags+=1
			interfaceExploit.Exploits[indexExploit].Flags+=1
		}
	}else{
		//TODO
	}
}



func submitSingleTeam( typeS string,indexTeam int, indexExploit int){
	for indexFlag,flag:= range interfaceExploit.Exploits[indexExploit].Teams[indexTeam].Flag{
		if flag.Status=="NEW" || flag.Status=="ERROR"{
			go submitSingleFlag(typeS,indexFlag,indexTeam,indexExploit)
		}else{
			break //non le guardo veramente tutte
		}
	}

}
func submit(typeS string,indexExploit int){
	for indexTeam, team := range interfaceExploit.Exploits[indexExploit].Teams{
		if(team.Flag[0].Status=="NEW"|| team.Flag[0].Status=="ERROR"){
			go submitSingleTeam(typeS,indexTeam,indexExploit)

		}else{
			break;
		}
	}
}



func Loop(typeSubmit string){
	for{
		fmt.Println("Executing Submits")
		for indexExploit, exploit:=range interfaceExploit.Exploits{
			if(exploit.Active ){
				go submit(typeSubmit,indexExploit)
			}
		}
		time.Sleep(time.Duration(interfaceExploit.TimeSubmit) * time.Second)
	}
}
