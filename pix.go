package openbank

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/stone-co/go-stone-openbank/types"
)

// PIXService handles communication with Stone Openbank API
type PIXService struct {
	client *Client
}

// GetOutboundPix is a service used to retrieve information details from a Pix.
func (s *PIXService) GetOutboundPix(id string) (*types.PIXOutBoundOutput, *Response, error) {

	path := fmt.Sprintf("/api/v1/pix/outbound_pix_payments/%s", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var pix types.PIXOutBoundOutput
	resp, err := s.client.Do(req, &pix)
	if err != nil {
		return nil, resp, err
	}

	return &pix, resp, err
}

// GetQRCodeData is a service used to retrieve information details from a Pix QRCode.
func (s *PIXService) GetQRCodeData(input types.GetQRCodeInput) (*types.QRCode, *Response, error) {
	const path = "/api/v1/pix/outbound_pix_payments/brcodes"

	req, err := s.client.NewAPIRequest(http.MethodGet, path, input)
	if err != nil {
		return nil, nil, err
	}

	var qrcode types.QRCode
	resp, err := s.client.Do(req, &qrcode)
	if err != nil {
		return nil, resp, err
	}

	return &qrcode, resp, err
}

// CreateDynamicQRCode is a service used to create a Dynamic Pix QRCode for receipt.
func (s *PIXService) CreateDynamicQRCode(input types.CreateDynamicQRCodeInput, idempotencyKey string) (*types.PIXInvoiceOutput, *Response, error) {
	const path = "/api/v1/pix_payment_invoices"

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, nil, err
	}

	if len(idempotencyKey) > idempotencyKeyMaxSize {
		return nil, nil, errors.New("invalid idempotency key")
	}
	req.Header.Add("x-stone-idempotency-key", idempotencyKey)

	var pixInvoiceOutput types.PIXInvoiceOutput
	resp, err := s.client.Do(req, &pixInvoiceOutput)
	if err != nil {
		return nil, resp, err
	}

	return &pixInvoiceOutput, resp, err
}

// GetEntries is a service used to retrieve all Pix entries.
func (s *PIXService) GetEntries(accountID string, idempotencyKey string) (*types.AllPixEntries, *Response, error) {
	path := fmt.Sprintf("/api/v1/pix/%s/entries", accountID)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	if len(idempotencyKey) > idempotencyKeyMaxSize {
		return nil, nil, errors.New("invalid idempotency key")
	}
	req.Header.Add("x-stone-idempotency-key", idempotencyKey)

	var allPixEntries types.AllPixEntries
	resp, err := s.client.Do(req, &allPixEntries)
	if err != nil {
		return nil, resp, err
	}

	return &allPixEntries, resp, err
}

// CreatePedingPayment is a service used to create a pending payment.
func (s *PIXService) CreatePedingPayment(input types.CreatePedingPaymentInput, idempotencyKey string) (*types.PendingPaymentOutput, *Response, error) {
	const path = "/api/v1/pix/outbound_pix_payments"

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, nil, err
	}

	if len(idempotencyKey) > idempotencyKeyMaxSize {
		return nil, nil, errors.New("invalid idempotency key")
	}
	req.Header.Add("x-stone-idempotency-key", idempotencyKey)

	var pendingPaymentOutput types.PendingPaymentOutput
	resp, err := s.client.Do(req, &pendingPaymentOutput)
	if err != nil {
		return nil, resp, err
	}

	return &pendingPaymentOutput, resp, err
}

// ConfirmPedingPayment is a service used to confirm a pending payment.
func (s *PIXService) ConfirmPedingPayment(input types.ConfirmPendingPaymentInput, idempotencyKey, pixID string) (*Response, error) {
	path := fmt.Sprintf("/api/v1/pix/outbound_pix_payments/%s/actions/confirm", pixID)

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, err
	}

	if len(idempotencyKey) > idempotencyKeyMaxSize {
		return nil, errors.New("invalid idempotency key")
	}
	req.Header.Add("x-stone-idempotency-key", idempotencyKey)

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
