package authentication

import (
	"github.com/go-ldap/ldap/v3"
)

// ProductionLDAPClientFactory the production implementation of an ldap connection factory.
type ProductionLDAPClientFactory struct {
	addr string
	opts []ldap.DialOpt
}

// NewProductionLDAPClientFactory create a concrete ldap connection factory.
func NewProductionLDAPClientFactory(addr string, opts ...ldap.DialOpt) *ProductionLDAPClientFactory {
	return &ProductionLDAPClientFactory{addr, opts}
}

// Dial creates a client from the default LDAP URL when successful.
func (f *ProductionLDAPClientFactory) Dial() (client LDAPClient, err error) {
	return ldap.DialURL(f.addr, f.opts...)
}

// DialURL creates a client from an LDAP URL when successful.
func (f *ProductionLDAPClientFactory) DialURL(addr string) (client LDAPClient, err error) {
	return ldap.DialURL(addr, f.opts...)
}
