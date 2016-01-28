package twilio

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var accountSid, authToken, apiURL, fromNum string

func init() {
	accountSid = os.Getenv("TWILIO_SID")
	authToken = os.Getenv("TWILIO_TOKEN")
	apiURL = "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	fromNum = os.Getenv("TWILIO_NUMBER")
}

func Send(toNum, msg string) error {
	v := url.Values{}
	v.Set("To", toNum)
	v.Set("From", fromNum)
	v.Set("Body", msg)

	rb := *strings.NewReader(v.Encode())
	req, err := http.NewRequest("POST", apiURL, &rb)
	if err != nil {
		return err
	}

	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil {
			return err
		}
	} else {
		return errors.New(resp.Status)
	}
	return nil
}
