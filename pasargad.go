package pasargad

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// The API URLs
const (
	URL_GET_TOKEN = "https://pep.shaparak.ir/Api/v1/Payment/GetToken"

	// Redirect User with token to this URL.
	// e.q: https://pep.shaparak.ir/payment.aspx?n=Token
	URL_PAYMENT_GATEWAY   = "https://pep.shaparak.ir/payment.aspx"
	URL_CHECK_TRANSACTION = "https://pep.shaparak.ir/Api/v1/Payment/CheckTransactionResult"
	URL_VERIFY_PAYMENT    = "https://pep.shaparak.ir/Api/v1/Payment/VerifyPayment"
	URL_REFUND            = "https://pep.shaparak.ir/Api/v1/Payment/RefundPayment"
)

const ACTION_PAYMENT = "1003"

// The API dateTime format for Pasargad
const dateTimeFormat = "2006/01/02 15:04:05"

// HTTPClient is HTTP client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// PasargadPaymentAPI for rest
type PasargadPaymentAPI struct {
	httpClient        HTTPClient
	merchantCode      int64
	terminalId        int64
	redirectUrl       string
	certificationFile string
	sign              string
}

// PasargadAPI creates a new PasargadPaymentAPI instance.
func PasargadAPI(merchantCode int64, terminalId int64, redirectUrl string, certificationFile string) *PasargadPaymentAPI {
	return PasargadAPIClient(merchantCode, terminalId, redirectUrl, certificationFile, &http.Client{})
}

// PasargadAPIClient creates a new PasargadPaymentAPI instance
// and allows you to pass a http.Client.
func PasargadAPIClient(merchantCode int64, terminalId int64, redirectUrl string, certificationFile string, httpClient HTTPClient) *PasargadPaymentAPI {
	return &PasargadPaymentAPI{
		httpClient:        httpClient,
		merchantCode:      merchantCode,
		terminalId:        terminalId,
		redirectUrl:       redirectUrl,
		certificationFile: certificationFile,
	}
}

// makeRequest is our RequestBuilder object (used in other packages of pepco-api)
func (m *PasargadPaymentAPI) makeRequest(url, method string, body interface{}, resp interface{}) error {
	var data []byte
	if body != nil {
		var err error
		data, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Sign", m.sign)

	r, err := m.httpClient.Do(req)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	// From Here --------------------
	respData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if r.StatusCode != 200 {
		var errRes ErrorResponse
		err = json.Unmarshal(respData, &errRes)
		if err != nil {
			return err
		}
		return errRes
	} else if resp != nil {
		err = json.Unmarshal(respData, &resp)
		if err != nil {
			return err
		}
	}
	// To here ---------------------

	return nil
}

// SetSign sets new  key.
func (m *PasargadPaymentAPI) SetSign(sign string) {
	m.sign = sign
}

// Sign Data with RSA key
func (m *PasargadPaymentAPI) signData(body interface{}) error {
	var data []byte
	if body != nil {
		var err error
		// Getting Data ready for signing with PKCS1
		data, err = json.Marshal(body)
		if err != nil {
			return err
		}
		// Convert XML to Public/Private Keys
		res, err := m.convertXmlToKey()
		if err != nil {
			return err
		}
		// Creating Signer with our private key...
		signer, err := NewSigner(res)
		if err != nil {
			return err
		}

		// ...and finally, siging data.
		signedMessage, err := signer.SignBase64(data)
		if err != nil {
			return err
		}
		m.SetSign(signedMessage)
	}
	return nil
}

// Get Current timestamp in Y/m/d H:i:s format
func (m *PasargadPaymentAPI) getTimestamp() string {
	now := time.Now()
	return now.Format(dateTimeFormat)
}

// Generate Payment URL
func (m *PasargadPaymentAPI) Redirect(request CreatePaymentRequest) (string, error) {
	requestBody := request.GetRedirectRequest()
	requestBody.Action = ACTION_PAYMENT
	requestBody.MerchantCode = m.merchantCode
	requestBody.TerminalCode = m.terminalId
	requestBody.RedirectAddress = m.redirectUrl
	requestBody.TimeStamp = m.getTimestamp()

	m.signData(requestBody)
	var resp RedirectResponse
	err := m.makeRequest(URL_GET_TOKEN, "POST", requestBody, &resp)

	if err != nil {
		return "", err
	}

	if resp.IsSuccess == false {
		requestError := errors.New(resp.Message)
		return "", requestError
	}
	// In this stage, we got a successful response from Pasargad IPG
	var redirectAddress string = URL_PAYMENT_GATEWAY + "?n=" + resp.Token
	return redirectAddress, nil
}

// CheckTransaction method
func (m *PasargadPaymentAPI) CheckTransaction(request CreateCheckTransactionRequest) (*CheckTransactionResponse, error) {
	requestBody := request.GetCheckTransactionRequest()
	requestBody.MerchantCode = m.merchantCode
	requestBody.TerminalCode = m.terminalId

	m.signData(requestBody)
	var resp CheckTransactionResponse
	err := m.makeRequest(URL_CHECK_TRANSACTION, "POST", requestBody, &resp)

	if err != nil {
		return nil, err
	}

	if resp.IsSuccess == false {
		requestError := errors.New(resp.Message)
		return nil, requestError
	}
	return &resp, nil
}

// VerifyPayment method
func (m *PasargadPaymentAPI) VerifyPayment(request CreateVerifyPaymentRequest) (*VerifyPaymentResponse, error) {
	requestBody := request.GetVerifyPaymentRequest()
	requestBody.MerchantCode = m.merchantCode
	requestBody.TerminalCode = m.terminalId
	requestBody.TimeStamp = m.getTimestamp()

	m.signData(requestBody)
	var resp VerifyPaymentResponse
	err := m.makeRequest(URL_VERIFY_PAYMENT, "POST", requestBody, &resp)

	if err != nil {
		return nil, err
	}

	if resp.IsSuccess == false {
		requestError := errors.New(resp.Message)
		return nil, requestError
	}
	return &resp, nil
}

// Refund method
func (m *PasargadPaymentAPI) Refund(request CreateRefundRequest) (*RefundResponse, error) {
	requestBody := request.GetRefundRequest()
	requestBody.MerchantCode = m.merchantCode
	requestBody.TerminalCode = m.terminalId
	requestBody.TimeStamp = m.getTimestamp()

	m.signData(requestBody)
	var resp RefundResponse
	err := m.makeRequest(URL_REFUND, "POST", requestBody, &resp)

	if err != nil {
		return nil, err
	}

	if resp.IsSuccess == false {
		requestError := errors.New(resp.Message)
		return nil, requestError
	}
	return &resp, nil
}
