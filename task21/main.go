package main

import "fmt"

// AuthService - interface which is being used in our system
type AuthService interface {
	ParseJSON()
}

// ThirdPartyAuthService - some implementation (ParseXML) which was coded by other devs, cannot be changed
type ThirdPartyAuthService struct {
}

func (s *ThirdPartyAuthService) ParseXml() {
	fmt.Println("Parsing XML auth data")
}

// AdapterAuthService - adaptor which adapt ThirdPartyAuthService to our interface AuthService
type AdapterAuthService struct {
	authService ThirdPartyAuthService
}

func (s AdapterAuthService) ParseJSON() {
	fmt.Println("Converting JSON data to XML...")
	s.authService.ParseXml()
}

func main() {
	a := AdapterAuthService{}
	a.ParseJSON()
}
