package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FirebaseImage struct {
	Token string `json:"downloadTokens"`
}

func FetchFirebaseImage(urlFirebase string) (string, error) {
	url := urlFirebase
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	fmt.Println(res)
	fmt.Println(res.Body)
	firebaseImage := FirebaseImage{}
	err = json.Unmarshal(body, &firebaseImage)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s?alt=media&token=%s", urlFirebase, firebaseImage.Token), nil
}
