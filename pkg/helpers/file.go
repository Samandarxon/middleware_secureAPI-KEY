package helpers

import (
	"encoding/json"
	"essy_travel/models"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

func Read(path string) ([]interface{}, error) {
	// time.Sleep(time.Second * 10)
	body, err := os.ReadFile(path)
	if err != nil {

		return nil, err
	}

	var response []interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	fmt.Println("Read")
	return response, nil
}

func ReadAirport(path string) ([]models.UpAirport, error) {
	// time.Sleep(time.Second * 10)
	body, err := os.ReadFile(path)
	if err != nil {

		return nil, err
	}

	var response []models.UpAirport
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	fmt.Println("Read")
	return response, nil
}
func ReadCity(path string) ([]models.UpCity, error) {
	// time.Sleep(time.Second * 10)
	body, err := os.ReadFile(path)
	if err != nil {

		return nil, err
	}

	var response []models.UpCity
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	fmt.Println("Read")
	return response, nil
}
func ReadCountry(path string) ([]models.UpCountry, error) {
	// time.Sleep(time.Second * 10)
	body, err := os.ReadFile(path)
	if err != nil {

		return nil, err
	}

	var response []models.UpCountry
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	fmt.Println("Read")
	return response, nil
}
func Write(path string, data []interface{}) error {

	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, body, os.ModePerm)
	if err != nil {
		return err
	}

	return err
}

func Upload(file *multipart.FileHeader, base string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	n := fmt.Sprintf("%d - %s", time.Now().UTC().Unix(), file.Filename)
	dst := fmt.Sprintf("%s/%s", base, n)
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	fmt.Println("Uploded")
	return n, err
}
