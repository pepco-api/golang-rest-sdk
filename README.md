# Golang SDK for Pasargad IPG

Golang SDK for Pasargad Internet Payment Gateway (RESTful API)

## Installation

```bash
go get -u github.com/pepco-api/golang-rest-sdk
```

## Usage
 - Read API Documentation, [Click Here! (دانلود مستندات کامل درگاه پرداخت)](https://www.pep.co.ir/wp-content/uploads/2019/06/1-__PEP_IPG_REST-13971020.Ver3_.00.pdf)
 - Save your private key into an `.xml` file inside your project directory.
 
 
## Redirect User to Payment Gateway
```go
package main
import (
	"fmt"
	"github.com/pepco-api/golang-rest-sdk"
)

func main() {	
    // Create an object from PasargadAPI struct
    // e.q: pasargadApi := pasargad.PasargadAPI(4412312,123456,"https://pep.co.ir/ipgtest","cert.xml");
    pasargadApi := pasargad.PasargadAPI(YOUR_MERCHANT_CODE, YOUR_TERMINAL_ID, REDIRECT_URL, CERT_FILE_HERE)
    request := pasargad.CreatePaymentRequest{
        Amount:        15000,
        InvoiceNumber: "4029",
        InvoiceDate:   "2021/08/17 16:12:00",
        //Mobile:        "09121231234",   // Optional
        //Email:         "xxxx@xxxx.xxx", // Optional
    }
    response, err := pasargadApi.Redirect(request)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(response)
    }
}
```

This method, will return a `string` like this:

```
// output:
 https://pep.shaparak.ir/payment.aspx?n=LySl+5tYkxL5qNMBRthW7DWzV8e3ALnTJUqiCS0V/io=
// Redirect User to the generated URL to make payment
```

## Checking and Verifying Transaction
After Payment Process, User is going to be returned to your REDIRECT_URL.

payment gateway is going to answer the payment result with sending below parameters to your redirectURL (as `QueryString` parameters):
 - InvoiceNumber (iN field) 
 - InvoiceDate (iD field) 
 - TransactionReferenceID (tref field) 
 
Store this information in a proper data storage and let's check transaction result by sending a check api request to the Bank:

```go
package main
import (
	"fmt"
	"github.com/pepco-api/golang-rest-sdk"
)

func main() {	
    pasargadApi := pasargad.PasargadAPI(4412312,123456,"https://pep.co.ir/ipgtest","cert.xml");
    request := pasargad.CreateCheckTransactionRequest{
        InvoiceNumber:          "4029",
        InvoiceDate:            "2021/08/17 16:12:00",
        TransactionReferenceID: "637648135297288707", // optional
    }
    response, err := pasargadApi.CheckTransaction(request)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(response)
    }
}
```

Successful result:
```json
{
    "TraceNumber": 13,
    "ReferenceNumber": 100200300400500,
    "TransactionDate": "2021/08/08 11:58:23",
    "Action": "1003",
    "TransactionReferenceID": "636843820118990203",
    "InvoiceNumber": "4029",
    "InvoiceDate": "2021/08/08 11:54:03",
    "MerchantCode": 100123,
    "TerminalCode": 200123,
    "Amount": 15000,
    "IsSuccess": true,
    "Message": " "
}
```
Our code will return `CheckTransactionResponse` struct as response or `err` in case of any errors.

If you got `IsSuccess` with `true` value, so everything is O.K!

Now, for your successful transaction, you should call `VerifyPayment()` method to keep the money and Bank makes sure the checking process done properly:


```go
package main
import (
	"fmt"
	"github.com/pepco-api/golang-rest-sdk"
)

func main() {	
	pasargadApi := pasargad.PasargadAPI(4412312,123456,"https://pep.co.ir/ipgtest","cert.xml");
	request := pasargad.CreateVerifyPaymentRequest{
		Amount:        15000,
		InvoiceDate:   "2021/08/17 16:12:00",,
		InvoiceNumber: "4029",
	}
	response, err := pasargadApi.VerifyPayment(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
```


...and the successful response looks like this response:
```json
{
 "IsSuccess": true,
 "Message": " ",
 "MaskedCardNumber": "5022-29**-****-2328",
 "HashedCardNumber": "2DDB1E270C598677AE328AA37C2970E3075E1DB....",
 "ShaparakRefNumber": "100200300400500"
}
```
Our code will return `VerifyPaymentResponse` struct as response or `err` in case of any errors.

## Payment Refund
If for any reason, you decided to cancel an order in early hours after taking the order (maximum 2 hours later), you can refund the client payment to his/her bank card.

for this, use `Refund()` method:


```go
package main
import (
	"fmt"
	"github.com/pepco-api/golang-rest-sdk"
)

func main() {
	pasargadApi := pasargad.PasargadAPI(4412312,123456,"https://pep.co.ir/ipgtest","cert.xml");
	request := pasargad.CreateRefundRequest{
		InvoiceDate:   "2021/08/17 16:12:00",
		InvoiceNumber: "4029",
	}
	response, err := pasargadApi.Refund(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}
```
Our code will return `RefundResponse` struct as response or `err` in case of any errors.

# Support
Please use your credentials to login into [Support Panel](https://my.pep.co.ir)
Contact Author/Maintainer: [Reza Seyf](https://twitter.com/seyfcode)