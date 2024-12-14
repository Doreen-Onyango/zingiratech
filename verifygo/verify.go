package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	verifyV2 "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	client    *twilio.RestClient
	verifySid string
)

func init() {
	// Load environment variables early to ensure credentials are available
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read Twilio credentials from environment variables
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	verifySid = os.Getenv("TWILIO_VERIFY_SID")

	if accountSid == "" || authToken == "" || verifySid == "" {
		log.Fatal("Error: Missing Twilio credentials in environment variables")
	}

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

func main() {
	// Define command-line flag for phone number
	phoneNumber := flag.String("phone", "", "The phone number to verify (e.g., +1234567890)")
	flag.Parse()

	// Ensure a phone number was provided
	if *phoneNumber == "" {
		fmt.Println("Error: Please provide a phone number using the -phone flag")
		os.Exit(1)
	}

	// Step 1: Send verification SMS
	if err := sendVerification(*phoneNumber, "sms"); err != nil {
		fmt.Printf("Error sending verification: %v\n", err)
		return
	}

	// Step 2: Get OTP from user input
	otpCode := getOTPFromUser()

	// Step 3: Verify OTP
	if err := verifyOTP(*phoneNumber, otpCode); err != nil {
		fmt.Printf("Error verifying OTP: %v\n", err)
		return
	}
}

func sendVerification(to, channel string) error {
	params := &verifyV2.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel(channel)

	// Send the verification request to Twilio
	verification, err := client.VerifyV2.CreateVerification(verifySid, params)
	if err != nil {
		return fmt.Errorf("failed to send verification: %v", err)
	}

	fmt.Printf("Verification sent: %v\n", *verification.Status)
	return nil
}

func verifyOTP(to, code string) error {
	params := &verifyV2.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	// Verify the OTP
	verificationCheck, err := client.VerifyV2.CreateVerificationCheck(verifySid, params)
	if err != nil {
		return fmt.Errorf("failed to verify OTP: %v", err)
	}

	fmt.Printf("Verification status: %v\n", *verificationCheck.Status)
	return nil
}

func getOTPFromUser() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter the OTP: ")
	otpCode, _ := reader.ReadString('\n')
	// Trim any newline or extra spaces
	return otpCode[:len(otpCode)-1] // Remove trailing newline character
}
