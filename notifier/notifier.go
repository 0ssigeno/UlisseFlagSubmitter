package notifier

import (
	"../execExploits"
	"../interfaceExploit"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"path/filepath"
	"strings"
)

//TODO controllo solo i file .teams
func Loop() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					completeName := strings.Split(event.Name, ".")
					if len(completeName) == 1 {
						break
					}
					if completeName[1] == interfaceExploit.TeamFileExtension {
						exploit := execExploits.CreateExploit(filepath.Base(completeName[0]))
						execExploits.AddExploitToList(exploit)
						fmt.Printf("Added exploit %s ", filepath.Base(completeName[0]))
					}
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					completeName := strings.Split(event.Name, ".")
					if len(completeName) == 1 {
						break
					}
					if completeName[1] == interfaceExploit.TeamFileExtension {
						execExploits.RemoveExploitFromList(filepath.Base(completeName[0]))
						fmt.Printf("Removed exploit %s ", filepath.Base(completeName[0]))
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(interfaceExploit.DirExploits)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
