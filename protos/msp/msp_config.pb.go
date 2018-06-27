// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msp/msp_config.proto

package msp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MSPConfig collects all the configuration information for
// an MSP. The Config field should be unmarshalled in a way
// that depends on the Type
type MSPConfig struct {
	// Type holds the type of the MSP; the default one would
	// be of type FABRIC implementing an X.509 based provider
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Config is MSP dependent configuration info
	Config []byte `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (m *MSPConfig) Reset()                    { *m = MSPConfig{} }
func (m *MSPConfig) String() string            { return proto.CompactTextString(m) }
func (*MSPConfig) ProtoMessage()               {}
func (*MSPConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MSPConfig) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MSPConfig) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

// FabricMSPConfig collects all the configuration information for
// a Fabric MSP.
// Here we assume a default certificate validation policy, where
// any certificate signed by any of the listed rootCA certs would
// be considered as valid under this MSP.
// This MSP may or may not come with a signing identity. If it does,
// it can also issue signing identities. If it does not, it can only
// be used to validate and verify certificates.
type FabricMSPConfig struct {
	// Name holds the identifier of the MSP; MSP identifier
	// is chosen by the application that governs this MSP.
	// For example, and assuming the default implementation of MSP,
	// that is X.509-based and considers a single Issuer,
	// this can refer to the Subject OU field or the Issuer OU field.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of root certificates trusted by this MSP
	// they are used upon certificate validation (see
	// comment for IntermediateCerts below)
	RootCerts [][]byte `protobuf:"bytes,2,rep,name=root_certs,json=rootCerts,proto3" json:"root_certs,omitempty"`
	// List of intermediate certificates trusted by this MSP;
	// they are used upon certificate validation as follows:
	// validation attempts to build a path from the certificate
	// to be validated (which is at one end of the path) and
	// one of the certs in the RootCerts field (which is at
	// the other end of the path). If the path is longer than
	// 2, certificates in the middle are searched within the
	// IntermediateCerts pool
	IntermediateCerts [][]byte `protobuf:"bytes,3,rep,name=intermediate_certs,json=intermediateCerts,proto3" json:"intermediate_certs,omitempty"`
	// Identity denoting the administrator of this MSP
	Admins [][]byte `protobuf:"bytes,4,rep,name=admins,proto3" json:"admins,omitempty"`
	// Identity revocation list
	RevocationList [][]byte `protobuf:"bytes,5,rep,name=revocation_list,json=revocationList,proto3" json:"revocation_list,omitempty"`
	// SigningIdentity holds information on the signing identity
	// this peer is to use, and which is to be imported by the
	// MSP defined before
	SigningIdentity *SigningIdentityInfo `protobuf:"bytes,6,opt,name=signing_identity,json=signingIdentity" json:"signing_identity,omitempty"`
	// OrganizationalUnitIdentifiers holds one or more
	// mchain organizational unit identifiers that belong to
	// this MSP configuration
	OrganizationalUnitIdentifiers []*FabricOUIdentifier `protobuf:"bytes,7,rep,name=organizational_unit_identifiers,json=organizationalUnitIdentifiers" json:"organizational_unit_identifiers,omitempty"`
	// FabricCryptoConfig contains the configuration parameters
	// for the cryptographic algorithms used by this MSP
	CryptoConfig *FabricCryptoConfig `protobuf:"bytes,8,opt,name=crypto_config,json=cryptoConfig" json:"crypto_config,omitempty"`
	// List of TLS root certificates trusted by this MSP.
	// They are returned by GetTLSRootCerts.
	TlsRootCerts [][]byte `protobuf:"bytes,9,rep,name=tls_root_certs,json=tlsRootCerts,proto3" json:"tls_root_certs,omitempty"`
	// List of TLS intermediate certificates trusted by this MSP;
	// They are returned by GetTLSIntermediateCerts.
	TlsIntermediateCerts [][]byte `protobuf:"bytes,10,rep,name=tls_intermediate_certs,json=tlsIntermediateCerts,proto3" json:"tls_intermediate_certs,omitempty"`
	// FabricNodeOUs contains the configuration to distinguish clients from peers from orderers
	// based on the OUs.
	FabricNodeOUs *FabricNodeOUs `protobuf:"bytes,11,opt,name=FabricNodeOUs" json:"FabricNodeOUs,omitempty"`
}

func (m *FabricMSPConfig) Reset()                    { *m = FabricMSPConfig{} }
func (m *FabricMSPConfig) String() string            { return proto.CompactTextString(m) }
func (*FabricMSPConfig) ProtoMessage()               {}
func (*FabricMSPConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *FabricMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FabricMSPConfig) GetRootCerts() [][]byte {
	if m != nil {
		return m.RootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetIntermediateCerts() [][]byte {
	if m != nil {
		return m.IntermediateCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetAdmins() [][]byte {
	if m != nil {
		return m.Admins
	}
	return nil
}

func (m *FabricMSPConfig) GetRevocationList() [][]byte {
	if m != nil {
		return m.RevocationList
	}
	return nil
}

func (m *FabricMSPConfig) GetSigningIdentity() *SigningIdentityInfo {
	if m != nil {
		return m.SigningIdentity
	}
	return nil
}

func (m *FabricMSPConfig) GetOrganizationalUnitIdentifiers() []*FabricOUIdentifier {
	if m != nil {
		return m.OrganizationalUnitIdentifiers
	}
	return nil
}

func (m *FabricMSPConfig) GetCryptoConfig() *FabricCryptoConfig {
	if m != nil {
		return m.CryptoConfig
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsRootCerts() [][]byte {
	if m != nil {
		return m.TlsRootCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetTlsIntermediateCerts() [][]byte {
	if m != nil {
		return m.TlsIntermediateCerts
	}
	return nil
}

func (m *FabricMSPConfig) GetFabricNodeOUs() *FabricNodeOUs {
	if m != nil {
		return m.FabricNodeOUs
	}
	return nil
}

// FabricCryptoConfig contains configuration parameters
// for the cryptographic algorithms used by the MSP
// this configuration refers to
type FabricCryptoConfig struct {
	// SignatureHashFamily is a string representing the hash family to be used
	// during sign and verify operations.
	// Allowed values are "SHA2" and "SHA3".
	SignatureHashFamily string `protobuf:"bytes,1,opt,name=signature_hash_family,json=signatureHashFamily" json:"signature_hash_family,omitempty"`
	// IdentityIdentifierHashFunction is a string representing the hash function
	// to be used during the computation of the identity identifier of an MSP identity.
	// Allowed values are "SHA256", "SHA384" and "SHA3_256", "SHA3_384".
	IdentityIdentifierHashFunction string `protobuf:"bytes,2,opt,name=identity_identifier_hash_function,json=identityIdentifierHashFunction" json:"identity_identifier_hash_function,omitempty"`
}

func (m *FabricCryptoConfig) Reset()                    { *m = FabricCryptoConfig{} }
func (m *FabricCryptoConfig) String() string            { return proto.CompactTextString(m) }
func (*FabricCryptoConfig) ProtoMessage()               {}
func (*FabricCryptoConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *FabricCryptoConfig) GetSignatureHashFamily() string {
	if m != nil {
		return m.SignatureHashFamily
	}
	return ""
}

func (m *FabricCryptoConfig) GetIdentityIdentifierHashFunction() string {
	if m != nil {
		return m.IdentityIdentifierHashFunction
	}
	return ""
}

// IdemixMSPConfig collects all the configuration information for
// an Idemix MSP.
type IdemixMSPConfig struct {
	// Name holds the identifier of the MSP
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// IPk represents the (serialized) issuer public key
	IPk []byte `protobuf:"bytes,2,opt,name=IPk,proto3" json:"IPk,omitempty"`
	// signer may contain crypto material to configure a default signer
	Signer *IdemixMSPSignerConfig `protobuf:"bytes,3,opt,name=signer" json:"signer,omitempty"`
}

func (m *IdemixMSPConfig) Reset()                    { *m = IdemixMSPConfig{} }
func (m *IdemixMSPConfig) String() string            { return proto.CompactTextString(m) }
func (*IdemixMSPConfig) ProtoMessage()               {}
func (*IdemixMSPConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *IdemixMSPConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *IdemixMSPConfig) GetIPk() []byte {
	if m != nil {
		return m.IPk
	}
	return nil
}

func (m *IdemixMSPConfig) GetSigner() *IdemixMSPSignerConfig {
	if m != nil {
		return m.Signer
	}
	return nil
}

// IdemixMSPSIgnerConfig contains the crypto material to set up an idemix signing identity
type IdemixMSPSignerConfig struct {
	// Cred represents the serialized idemix credential of the default signer
	Cred []byte `protobuf:"bytes,1,opt,name=Cred,proto3" json:"Cred,omitempty"`
	// Sk is the secret key of the default signer, corresponding to credential Cred
	Sk []byte `protobuf:"bytes,2,opt,name=Sk,proto3" json:"Sk,omitempty"`
	// organizational_unit_identifier defines the organizational unit the default signer is in
	OrganizationalUnitIdentifier string `protobuf:"bytes,3,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier" json:"organizational_unit_identifier,omitempty"`
	// is_admin defines whether the default signer is admin or not
	IsAdmin bool `protobuf:"varint,4,opt,name=is_admin,json=isAdmin" json:"is_admin,omitempty"`
}

func (m *IdemixMSPSignerConfig) Reset()                    { *m = IdemixMSPSignerConfig{} }
func (m *IdemixMSPSignerConfig) String() string            { return proto.CompactTextString(m) }
func (*IdemixMSPSignerConfig) ProtoMessage()               {}
func (*IdemixMSPSignerConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *IdemixMSPSignerConfig) GetCred() []byte {
	if m != nil {
		return m.Cred
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetSk() []byte {
	if m != nil {
		return m.Sk
	}
	return nil
}

func (m *IdemixMSPSignerConfig) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

func (m *IdemixMSPSignerConfig) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

// SigningIdentityInfo represents the configuration information
// related to the signing identity the peer is to use for generating
// endorsements
type SigningIdentityInfo struct {
	// PublicSigner carries the public information of the signing
	// identity. For an X.509 provider this would be represented by
	// an X.509 certificate
	PublicSigner []byte `protobuf:"bytes,1,opt,name=public_signer,json=publicSigner,proto3" json:"public_signer,omitempty"`
	// PrivateSigner denotes a reference to the private key of the
	// peer's signing identity
	PrivateSigner *KeyInfo `protobuf:"bytes,2,opt,name=private_signer,json=privateSigner" json:"private_signer,omitempty"`
}

func (m *SigningIdentityInfo) Reset()                    { *m = SigningIdentityInfo{} }
func (m *SigningIdentityInfo) String() string            { return proto.CompactTextString(m) }
func (*SigningIdentityInfo) ProtoMessage()               {}
func (*SigningIdentityInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SigningIdentityInfo) GetPublicSigner() []byte {
	if m != nil {
		return m.PublicSigner
	}
	return nil
}

func (m *SigningIdentityInfo) GetPrivateSigner() *KeyInfo {
	if m != nil {
		return m.PrivateSigner
	}
	return nil
}

// KeyInfo represents a (secret) key that is either already stored
// in the bccsp/keystore or key material to be imported to the
// bccsp key-store. In later versions it may contain also a
// keystore identifier
type KeyInfo struct {
	// Identifier of the key inside the default keystore; this for
	// the case of Software BCCSP as well as the HSM BCCSP would be
	// the SKI of the key
	KeyIdentifier string `protobuf:"bytes,1,opt,name=key_identifier,json=keyIdentifier" json:"key_identifier,omitempty"`
	// KeyMaterial (optional) for the key to be imported; this is
	// properly encoded key bytes, prefixed by the type of the key
	KeyMaterial []byte `protobuf:"bytes,2,opt,name=key_material,json=keyMaterial,proto3" json:"key_material,omitempty"`
}

func (m *KeyInfo) Reset()                    { *m = KeyInfo{} }
func (m *KeyInfo) String() string            { return proto.CompactTextString(m) }
func (*KeyInfo) ProtoMessage()               {}
func (*KeyInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *KeyInfo) GetKeyIdentifier() string {
	if m != nil {
		return m.KeyIdentifier
	}
	return ""
}

func (m *KeyInfo) GetKeyMaterial() []byte {
	if m != nil {
		return m.KeyMaterial
	}
	return nil
}

// FabricOUIdentifier represents an organizational unit and
// its related chain of trust identifier.
type FabricOUIdentifier struct {
	// Certificate represents the second certificate in a certification chain.
	// (Notice that the first certificate in a certification chain is supposed
	// to be the certificate of an identity).
	// It must correspond to the certificate of root or intermediate CA
	// recognized by the MSP this message belongs to.
	// Starting from this certificate, a certification chain is computed
	// and bound to the OrganizationUnitIdentifier specified
	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// OrganizationUnitIdentifier defines the organizational unit under the
	// MSP identified with MSPIdentifier
	OrganizationalUnitIdentifier string `protobuf:"bytes,2,opt,name=organizational_unit_identifier,json=organizationalUnitIdentifier" json:"organizational_unit_identifier,omitempty"`
}

func (m *FabricOUIdentifier) Reset()                    { *m = FabricOUIdentifier{} }
func (m *FabricOUIdentifier) String() string            { return proto.CompactTextString(m) }
func (*FabricOUIdentifier) ProtoMessage()               {}
func (*FabricOUIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *FabricOUIdentifier) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *FabricOUIdentifier) GetOrganizationalUnitIdentifier() string {
	if m != nil {
		return m.OrganizationalUnitIdentifier
	}
	return ""
}

// FabricNodeOUs contains configuration to tell apart clients from peers from orderers
// based on OUs. If NodeOUs recognition is enabled then an msp identity
// that does not contain any of the specified OU will be considered invalid.
type FabricNodeOUs struct {
	// If true then an msp identity that does not contain any of the specified OU will be considered invalid.
	Enable bool `protobuf:"varint,1,opt,name=Enable" json:"Enable,omitempty"`
	// OU Identifier of the clients
	ClientOUIdentifier *FabricOUIdentifier `protobuf:"bytes,2,opt,name=clientOUIdentifier" json:"clientOUIdentifier,omitempty"`
	// OU Identifier of the peers
	PeerOUIdentifier *FabricOUIdentifier `protobuf:"bytes,3,opt,name=peerOUIdentifier" json:"peerOUIdentifier,omitempty"`
}

func (m *FabricNodeOUs) Reset()                    { *m = FabricNodeOUs{} }
func (m *FabricNodeOUs) String() string            { return proto.CompactTextString(m) }
func (*FabricNodeOUs) ProtoMessage()               {}
func (*FabricNodeOUs) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *FabricNodeOUs) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *FabricNodeOUs) GetClientOUIdentifier() *FabricOUIdentifier {
	if m != nil {
		return m.ClientOUIdentifier
	}
	return nil
}

func (m *FabricNodeOUs) GetPeerOUIdentifier() *FabricOUIdentifier {
	if m != nil {
		return m.PeerOUIdentifier
	}
	return nil
}

func init() {
	proto.RegisterType((*MSPConfig)(nil), "msp.MSPConfig")
	proto.RegisterType((*FabricMSPConfig)(nil), "msp.FabricMSPConfig")
	proto.RegisterType((*FabricCryptoConfig)(nil), "msp.FabricCryptoConfig")
	proto.RegisterType((*IdemixMSPConfig)(nil), "msp.IdemixMSPConfig")
	proto.RegisterType((*IdemixMSPSignerConfig)(nil), "msp.IdemixMSPSignerConfig")
	proto.RegisterType((*SigningIdentityInfo)(nil), "msp.SigningIdentityInfo")
	proto.RegisterType((*KeyInfo)(nil), "msp.KeyInfo")
	proto.RegisterType((*FabricOUIdentifier)(nil), "msp.FabricOUIdentifier")
	proto.RegisterType((*FabricNodeOUs)(nil), "msp.FabricNodeOUs")
}

func init() { proto.RegisterFile("msp/msp_config.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x6f, 0x1a, 0x49,
	0x10, 0xd5, 0x00, 0xc6, 0x50, 0x0c, 0xe0, 0x6d, 0x7f, 0xec, 0xec, 0x6a, 0xed, 0xc5, 0xb3, 0x1b,
	0x85, 0x4b, 0x40, 0xc2, 0x91, 0x92, 0x43, 0x2e, 0x31, 0x89, 0x13, 0x94, 0x38, 0xb6, 0x06, 0xf9,
	0x92, 0xcb, 0xa8, 0x19, 0x9a, 0xa1, 0xc5, 0x7c, 0xa9, 0xbb, 0xb1, 0x42, 0x94, 0x7f, 0x91, 0x6b,
	0x7e, 0x43, 0x6e, 0xf9, 0x7f, 0x51, 0x7f, 0x80, 0x87, 0x98, 0x90, 0xdc, 0xba, 0xeb, 0xbd, 0x57,
	0x5d, 0xf3, 0xaa, 0xba, 0x07, 0x0e, 0x62, 0x9e, 0x75, 0x63, 0x9e, 0xf9, 0x41, 0x9a, 0x4c, 0x68,
	0xd8, 0xc9, 0x58, 0x2a, 0x52, 0x54, 0x8c, 0x79, 0xe6, 0x3e, 0x81, 0xea, 0xe5, 0xf0, 0xba, 0xaf,
	0xe2, 0x08, 0x41, 0x49, 0x2c, 0x32, 0xe2, 0x58, 0x2d, 0xab, 0xbd, 0xe3, 0xa9, 0x35, 0x3a, 0x82,
	0xb2, 0x56, 0x39, 0x85, 0x96, 0xd5, 0xb6, 0x3d, 0xb3, 0x73, 0xbf, 0x96, 0xa0, 0x79, 0x81, 0x47,
	0x8c, 0x06, 0x6b, 0xfa, 0x04, 0xc7, 0x5a, 0x5f, 0xf5, 0xd4, 0x1a, 0x1d, 0x03, 0xb0, 0x34, 0x15,
	0x7e, 0x40, 0x98, 0xe0, 0x4e, 0xa1, 0x55, 0x6c, 0xdb, 0x5e, 0x55, 0x46, 0xfa, 0x32, 0x80, 0x1e,
	0x01, 0xa2, 0x89, 0x20, 0x2c, 0x26, 0x63, 0x8a, 0x05, 0x31, 0xb4, 0xa2, 0xa2, 0xfd, 0x91, 0x47,
	0x34, 0xfd, 0x08, 0xca, 0x78, 0x1c, 0xd3, 0x84, 0x3b, 0x25, 0x45, 0x31, 0x3b, 0xf4, 0x10, 0x9a,
	0x8c, 0xdc, 0xa6, 0x01, 0x16, 0x34, 0x4d, 0xfc, 0x88, 0x72, 0xe1, 0xec, 0x28, 0x42, 0xe3, 0x2e,
	0xfc, 0x96, 0x72, 0x81, 0xfa, 0xb0, 0xc7, 0x69, 0x98, 0xd0, 0x24, 0xf4, 0xe9, 0x98, 0x24, 0x82,
	0x8a, 0x85, 0x53, 0x6e, 0x59, 0xed, 0x5a, 0xcf, 0xe9, 0xc4, 0x3c, 0xeb, 0x0c, 0x35, 0x38, 0x30,
	0xd8, 0x20, 0x99, 0xa4, 0x5e, 0x93, 0xaf, 0x07, 0x91, 0x0f, 0xff, 0xa6, 0x2c, 0xc4, 0x09, 0xfd,
	0xa8, 0x12, 0xe3, 0xc8, 0x9f, 0x27, 0x54, 0x98, 0x84, 0x13, 0x4a, 0x18, 0x77, 0x76, 0x5b, 0xc5,
	0x76, 0xad, 0xf7, 0xa7, 0xca, 0xa9, 0x6d, 0xba, 0xba, 0x19, 0xac, 0x70, 0xef, 0x78, 0x5d, 0x7f,
	0x93, 0x50, 0x71, 0x87, 0x72, 0xf4, 0x0c, 0xea, 0x01, 0x5b, 0x64, 0x22, 0x35, 0x1d, 0x73, 0x2a,
	0xaa, 0xc4, 0x7c, 0xba, 0xbe, 0xc2, 0xb5, 0xf1, 0x9e, 0x1d, 0xe4, 0x76, 0xe8, 0x7f, 0x68, 0x88,
	0x88, 0xfb, 0x39, 0xdb, 0xab, 0xca, 0x0b, 0x5b, 0x44, 0xdc, 0x5b, 0x39, 0xff, 0x18, 0x8e, 0x24,
	0x6b, 0x83, 0xfb, 0xa0, 0xd8, 0x07, 0x22, 0xe2, 0x83, 0x7b, 0x0d, 0x78, 0x0a, 0x75, 0x7d, 0xfe,
	0xbb, 0x74, 0x4c, 0xae, 0x6e, 0xb8, 0x53, 0x53, 0x95, 0xa1, 0x5c, 0x65, 0x06, 0xf1, 0xd6, 0x89,
	0xee, 0x67, 0x0b, 0xd0, 0xfd, 0xd2, 0x51, 0x0f, 0x0e, 0xa5, 0xbd, 0x58, 0xcc, 0x19, 0xf1, 0xa7,
	0x98, 0x4f, 0xfd, 0x09, 0x8e, 0x69, 0xb4, 0x30, 0x43, 0xb4, 0xbf, 0x02, 0x5f, 0x63, 0x3e, 0xbd,
	0x50, 0x10, 0x1a, 0xc0, 0xe9, 0xb2, 0x79, 0x39, 0xd3, 0x8d, 0x7a, 0x9e, 0x04, 0xd2, 0x54, 0x35,
	0xae, 0x55, 0xef, 0x64, 0x49, 0xbc, 0xb3, 0x57, 0x25, 0x32, 0x2c, 0x77, 0x06, 0xcd, 0xc1, 0x98,
	0xc4, 0xf4, 0xc3, 0xf6, 0x29, 0xde, 0x83, 0xe2, 0xe0, 0x7a, 0x66, 0xae, 0x80, 0x5c, 0xa2, 0x1e,
	0x94, 0x65, 0x69, 0x84, 0x39, 0x45, 0xe5, 0xc0, 0xdf, 0xca, 0x81, 0x55, 0xae, 0xa1, 0xc2, 0x4c,
	0x7b, 0x0c, 0xd3, 0xfd, 0x62, 0xc1, 0xe1, 0x46, 0x86, 0x3c, 0xb3, 0xcf, 0xc8, 0x58, 0x9d, 0x69,
	0x7b, 0x6a, 0x8d, 0x1a, 0x50, 0x18, 0x2e, 0x8f, 0x2c, 0x0c, 0x67, 0xe8, 0x05, 0x9c, 0x6c, 0x9f,
	0x3a, 0x55, 0x49, 0xd5, 0xfb, 0x67, 0xdb, 0x6c, 0xa1, 0xbf, 0xa0, 0x42, 0xb9, 0xaf, 0xae, 0x8d,
	0x53, 0x6a, 0x59, 0xed, 0x8a, 0xb7, 0x4b, 0xf9, 0x73, 0xb9, 0x75, 0x53, 0xd8, 0xdf, 0x30, 0xfe,
	0xe8, 0x3f, 0xa8, 0x67, 0xf3, 0x51, 0x44, 0x03, 0xdf, 0x7c, 0xb0, 0x2e, 0xd2, 0xd6, 0x41, 0xfd,
	0x19, 0xe8, 0x0c, 0x1a, 0x19, 0xa3, 0xb7, 0x72, 0x88, 0x0c, 0xab, 0xa0, 0x6c, 0xb1, 0x95, 0x2d,
	0x6f, 0x88, 0xbe, 0x49, 0x75, 0xc3, 0xd1, 0x22, 0x77, 0x08, 0xbb, 0x06, 0x41, 0x0f, 0xa0, 0x31,
	0x23, 0xf9, 0x6e, 0x1a, 0xfb, 0xeb, 0x33, 0x92, 0x6b, 0x1d, 0x3a, 0x05, 0x5b, 0xd2, 0x62, 0x2c,
	0x08, 0xa3, 0x38, 0x32, 0xee, 0xd4, 0x66, 0x64, 0x71, 0x69, 0x42, 0xee, 0xa7, 0xe5, 0x98, 0xe5,
	0x2f, 0x1c, 0x6a, 0x41, 0x4d, 0x0e, 0x37, 0x9d, 0xd0, 0x00, 0x0b, 0x62, 0x3e, 0x21, 0x1f, 0xfa,
	0x0d, 0x7b, 0x0b, 0xbf, 0xb6, 0xd7, 0xfd, 0x66, 0xfd, 0x70, 0x41, 0xe4, 0x93, 0xf5, 0x32, 0xc1,
	0xa3, 0x48, 0x1f, 0x5a, 0xf1, 0xcc, 0x0e, 0xbd, 0x02, 0x14, 0x44, 0x94, 0x24, 0x22, 0x5f, 0xa7,
	0x71, 0xed, 0xa7, 0xef, 0xc6, 0x06, 0x89, 0x7c, 0xd2, 0x32, 0x42, 0xd8, 0x5a, 0x9a, 0xe2, 0xf6,
	0x34, 0xf7, 0x04, 0xe7, 0x3e, 0x9c, 0xa6, 0x2c, 0xec, 0x4c, 0x17, 0x19, 0x61, 0x11, 0x19, 0x87,
	0x84, 0x75, 0x26, 0x4a, 0xa7, 0x7f, 0x16, 0x5c, 0x66, 0x3a, 0xdf, 0xbb, 0xe4, 0x99, 0x1e, 0xd8,
	0x6b, 0x1c, 0xcc, 0x70, 0x48, 0xde, 0xb7, 0x43, 0x2a, 0xa6, 0xf3, 0x51, 0x27, 0x48, 0xe3, 0x6e,
	0x4e, 0xdb, 0xd5, 0xda, 0xae, 0xd6, 0xca, 0x5f, 0xcf, 0xa8, 0xac, 0xd6, 0x67, 0xdf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x77, 0x69, 0x37, 0x77, 0x8c, 0x06, 0x00, 0x00,
}