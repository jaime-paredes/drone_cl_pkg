package internal

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jaime-paredes/drone_cl_pkg/helpers"
	"github.com/jaime-paredes/drone_cl_pkg/models"
	"github.com/jaime-paredes/drone_cl_pkg/pkg"
	"software.sslmate.com/src/go-pkcs12"
)

//Data of a certificate with node revoked
type Data struct {
	Revoked bool `json:"revoked"`
}

//CertCheck receive an array with all companies in production, get its certificate and will check if was revoked
func CertCheck(companies []models.Companies) {
	var headers map[string]string
	body := []byte{}
	var wg sync.WaitGroup
	wg.Add(len(companies))
	for _, cpn := range companies {
		go func(cpn models.Companies) {
			defer wg.Done()
			trackingAPI := os.Getenv("trackingAPI")
			certURL := fmt.Sprintf("%s/v1/companies/%d.json?expand=[certificate]", trackingAPI, cpn.CpnID)
			request := helpers.Request{
				Method:  "GET",
				BaseURL: certURL,
				Headers: headers,
				Body:    body,
			}
			respCert, err := helpers.APICall(request)
			if err != nil {
				fmt.Println("Error calling API Tracking", err.Error())
				return
			}
			bodyCert := (respCert.Body)

			var comp models.Company
			if err = json.Unmarshal(bodyCert, &comp); err != nil {
				fmt.Println(err)
			}
			fmt.Println(comp.CpnID)
			p12, err := base64.StdEncoding.DecodeString(comp.Certificate.Base64)
			if err != nil {
				fmt.Println("Error decoding certificate", err.Error())
				return
			}
			_, cert, err := pkcs12.Decode(p12, comp.Certificate.Pass)

			revoked, result, err := pkg.RevCheck(cert)
			if err != nil {
				fmt.Println("Error checking revocation certificate")
				return
			}
			// mongoResp := apicall.APICall(request apicall.Request)
			if !revoked && result {
				fmt.Println("Certificate ok")
			} else if revoked && !result {
				fmt.Println("error checking revocation")
			} else if revoked && result {
				mongoURL := fmt.Sprintf("%s/v1/companies/%d/certificate.json", trackingAPI, comp.CpnID)
				body, _ = json.Marshal(Data{Revoked: true})
				request := helpers.Request{
					Method:  "PUT",
					BaseURL: mongoURL,
					Body:    body,
				}
				respMongo, err := helpers.APICall(request)
				if err != nil {
					fmt.Println("Error calling PUT API Tracking")
				}
				fmt.Println(respMongo)
				droneCL := os.Getenv("droneCLAPI")
				rdsURL := fmt.Sprintf("%s/v1/companies/%d/certificate.json", droneCL, comp.CpnID)
				reqRds := helpers.Request{
					Method:  "GET",
					BaseURL: rdsURL,
					Body:    body,
				}
				respRds, err := helpers.APICall(reqRds)
				if err != nil {
					fmt.Println("Error calling RDS tracking")
				}
				fmt.Println(respRds)
				notifyURL := os.Getenv("notifyAPI")
				currentTime := time.Now()
				msg := fmt.Sprintf("Empresa: %s (cpn_id: %d, Fecha: %s, Certificado Revocado)", comp.CpnName, comp.CpnID, currentTime.Format("2006-01-02 15:04:05"))
				message := models.TrackData{Msg: msg}
				bodyNotify, _ := json.Marshal(models.Notification{CpnID: comp.CpnID, TrackType: "Certificado Revocado", TrackData: message})
				reqNotify := helpers.Request{
					Method:  "POST",
					BaseURL: notifyURL,
					Body:    bodyNotify,
				}
				respNotify, err := helpers.APICall(reqNotify)
				if err != nil {
					fmt.Println("Error calling NotifyTrack")
				}
				fmt.Println(respNotify.StatusCode)
				fmt.Println("Certificated is revoked")
			} else {
				fmt.Println("Certificate Ok")
			}
		}(cpn)
	}
	wg.Wait()
}
