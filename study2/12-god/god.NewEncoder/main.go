package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	save()
}
func save() error {
	DATAFILE := "D:/cep/save"
	fmt.Println("Saving", DATAFILE)
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}
	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("Cannot create", DATAFILE)
		return err
	}
	defer saveTo.Close()
	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode("")
	if err != nil {
		fmt.Println("Cannot save to", DATAFILE)
		return err
	}
	return nil
}
