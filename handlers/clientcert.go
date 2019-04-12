package handlers

import (
	"encoding/pem"
	"net/http"
	"strings"

	"code.cloudfoundry.org/gorouter/config"

	"code.cloudfoundry.org/gorouter/logger"
	"github.com/uber-go/zap"
	"github.com/urfave/negroni"
)

const xfcc = "X-Forwarded-Client-Cert"

type clientCert struct {
	skipSanitization  func(req *http.Request) bool
	forceDeleteHeader func(req *http.Request) (bool, error)
	forwardingMode    string
	logger            logger.Logger
}

func NewClientCert(skipSanitization func(req *http.Request) bool, forceDeleteHeader func(req *http.Request) (bool, error), forwardingMode string, logger logger.Logger) negroni.Handler {
	return &clientCert{
		skipSanitization:  skipSanitization,
		forceDeleteHeader: forceDeleteHeader,
		forwardingMode:    forwardingMode,
		logger:            logger,
	}
}

func (c *clientCert) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	skip := c.skipSanitization(r)
	if !skip {
		switch c.forwardingMode {
		case config.FORWARD:
			// TODO: Consider putting these in readable booleans like "isTLS" and "hasPeerCertificates"
			if r.TLS == nil || len(r.TLS.PeerCertificates) == 0 {
				r.Header.Del(xfcc)
			}
		case config.SANITIZE_SET:
			r.Header.Del(xfcc)
			if r.TLS != nil {
				sanitizeHeader(r)
			}
		}
	}

	delete, err := c.forceDeleteHeader(r)
	if err != nil {
		c.logger.Error("signature-validation-failed", zap.Error(err))
		writeStatus(
			rw,
			http.StatusBadRequest,
			"Failed to validate Route Service Signature for x-forwarded-client-cert",
			c.logger,
		)
		return
	}
	if delete {
		r.Header.Del(xfcc)
	}
	next(rw, r)
}

//TODO: rename this function
// maybe setXFCCHeaderWithTlsCertificate ?
func sanitizeHeader(r *http.Request) {
	// we only care about the first cert at this moment
	if len(r.TLS.PeerCertificates) > 0 {
		cert := r.TLS.PeerCertificates[0]
		b := pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}
		certPEM := pem.EncodeToMemory(&b)
		r.Header.Add(xfcc, sanitize(certPEM))
	}
}

func sanitize(cert []byte) string {
	s := string(cert)
	r := strings.NewReplacer("-----BEGIN CERTIFICATE-----", "",
		"-----END CERTIFICATE-----", "",
		"\n", "")
	return r.Replace(s)
}
