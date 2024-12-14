package main

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	TWILIO_ACCOUNT_SID string             = os.Getenv("TWILIO_ACCOUNT_SID")
	TWILIO_AUTH_TOKEN  string             = os.Getenv("TWILIO_AUTH_TOKEN")
	VERIFY_SERVICE_SID string             = os.Getenv("VERIFY_SERVICE_SID")
	client             *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})
)

func sendOtp(to string) {
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sent verification '%s'\n", *resp.Sid)
	}
}
