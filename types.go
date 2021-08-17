package pasargad

// General Types ================================================
// ErrorResponse is the API error response.
type ErrorResponse struct {
	IsSuccess bool   `json:"IsSuccess"` // status of request (True or False)
	Message   string `json:"Message"`   // Message - The response of server.
}

// Error returns a formatted error string.
func (m ErrorResponse) Error() string {
	return "Error Message: " + m.Message
}

// Redirect Data Types ==========================================
// CreatePaymentRequest is the format of our data to send to bank
type CreatePaymentRequest struct {
	Amount          int64  `json:"amount"`          // invoice amount
	InvoiceNumber   string `json:"invoiceNumber"`   // invoice number
	InvoiceDate     string `json:"invoiceDate"`     // invoice date
	Action          string `json:"action"`          // request action identifier
	Mobile          string `json:"mobile"`          // mobile number of the user
	Email           string `json:"email"`           // email address of the user
	MerchantCode    int64  `json:"merchantCode"`    // merchant code
	TerminalCode    int64  `json:"terminalCode"`    // terminal code
	RedirectAddress string `json:"redirectAddress"` // redirect url
	TimeStamp       string `json:"timeStamp"`       // Current timestamp (Y/m/d H:i:s)
}

// Get Redirect Request and it's parameters
func (m *CreatePaymentRequest) GetRedirectRequest() CreatePaymentRequest {
	return CreatePaymentRequest{
		Amount:          m.Amount,
		InvoiceNumber:   m.InvoiceNumber,
		InvoiceDate:     m.InvoiceDate,
		Mobile:          m.Mobile,
		Email:           m.Email,
		Action:          m.Action,
		MerchantCode:    m.MerchantCode,
		TerminalCode:    m.TerminalCode,
		RedirectAddress: m.RedirectAddress,
		TimeStamp:       m.TimeStamp,
	}
}

// RedirectResponse data type
type RedirectResponse struct {
	IsSuccess bool   `json:"IsSuccess"` // status of request (True or False)
	Message   string `json:"Message"`   // Message - The response of server
	Token     string `json:"Token"`     // Token (for successful requests)
}

// Check Transaction Data Types ===================================================
// CreateCheckTransactionRequest is the struct to create a checkTransaction request
type CreateCheckTransactionRequest struct {
	TransactionReferenceID string `json:"transactionReferenceID"` // transaction reference id
	InvoiceNumber          string `json:"invoiceNumber"`          // invoice number
	InvoiceDate            string `json:"invoiceDate"`            // invoice date
	TerminalCode           int64  `json:"terminalCode"`           // terminal code
	MerchantCode           int64  `json:"merchantCode"`           // merchant code
}

// GetCheckTransactionRequest and parameters
func (m *CreateCheckTransactionRequest) GetCheckTransactionRequest() CreateCheckTransactionRequest {
	return CreateCheckTransactionRequest{
		InvoiceNumber:          m.InvoiceNumber,
		InvoiceDate:            m.InvoiceDate,
		TransactionReferenceID: m.TransactionReferenceID,
		MerchantCode:           m.MerchantCode,
		TerminalCode:           m.TerminalCode,
	}
}

// RedirectResponse data type
type CheckTransactionResponse struct {
	IsSuccess              bool   `json:"IsSuccess"`              // status of request (True or False)
	Message                string `json:"Message"`                // Message - The response of server
	ReferenceNumber        int64  `json:"ReferenceNumber"`        // Reference number
	TraceNumber            int64  `json:"TraceNumber"`            // trace number
	TransactionDate        string `json:"TransactionDate"`        // transaction date
	Action                 string `json:"Action"`                 // action identifier
	TransactionReferenceID string `json:"TransactionReferenceID"` // transaaction internal reference id
	InvoiceNumber          string `json:"InvoiceNumber"`          // invoice number
	InvoiceDate            string `json:"InvoiceDate"`            // invoice date
	MerchantCode           int64  `json:"MerchantCode"`           // merchant code
	TerminalCode           int64  `json:"TerminalCode"`           // terminal code
	Amount                 int64  `json:"amount"`                 // transaction amount
}

// Verify Payment Data Types ================================================
// CreateVerifyPaymentRequest is the struct to create a VerifyPayment request
type CreateVerifyPaymentRequest struct {
	Amount        int64  `json:"amount"`        // invoice amount
	InvoiceNumber string `json:"invoiceNumber"` // invoice number
	InvoiceDate   string `json:"invoiceDate"`   // invoice date
	TerminalCode  int64  `json:"terminalCode"`  // terminal code
	MerchantCode  int64  `json:"merchantCode"`  // merchant code
	TimeStamp     string `json:"timeStamp"`     // Current timestamp (Y/m/d H:i:s)
}

// GetVerifyPaymentRequest and parameters
func (m *CreateVerifyPaymentRequest) GetVerifyPaymentRequest() CreateVerifyPaymentRequest {
	return CreateVerifyPaymentRequest{
		Amount:        m.Amount,
		InvoiceNumber: m.InvoiceNumber,
		InvoiceDate:   m.InvoiceDate,
		MerchantCode:  m.MerchantCode,
		TerminalCode:  m.TerminalCode,
		TimeStamp:     m.TimeStamp,
	}
}

// VerifyPaymentResponse
// VerifyPaymentResponse data type
type VerifyPaymentResponse struct {
	IsSuccess         bool   `json:"IsSuccess"`         // status of request (True or False)
	Message           string `json:"Message"`           // Message - The response of server
	MaskedCardNumber  string `json:"MaskedCardNumber"`  // Masked Card Number (like: 5022-29**-****-2328)
	HashedCardNumber  string `json:"HashedCardNumber"`  // Hashed Card Number (like: 2DDB1E270C598.....)
	ShaparakRefNumber string `json:"ShaparakRefNumber"` // Shaparak Reference ID (like: 100200300400500)
}

// Refund Data Types ==================================================
// CreateRefundRequest is the struct to create a Refund Payment to user
type CreateRefundRequest struct {
	InvoiceNumber string `json:"invoiceNumber"` // invoice number
	InvoiceDate   string `json:"invoiceDate"`   // invoice date
	TerminalCode  int64  `json:"terminalCode"`  // terminal code
	MerchantCode  int64  `json:"merchantCode"`  // merchant code
	TimeStamp     string `json:"timeStamp"`     // Current timestamp (Y/m/d H:i:s)
}

// GetRefundRequest and parameters
func (m *CreateRefundRequest) GetRefundRequest() CreateRefundRequest {
	return CreateRefundRequest{
		InvoiceNumber: m.InvoiceNumber,
		InvoiceDate:   m.InvoiceDate,
		MerchantCode:  m.MerchantCode,
		TerminalCode:  m.TerminalCode,
		TimeStamp:     m.TimeStamp,
	}
}

// RefundResponse data type
type RefundResponse struct {
	IsSuccess bool   `json:"IsSuccess"` // status of request (True or False)
	Message   string `json:"Message"`   // Message - The response of server
}
