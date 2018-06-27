// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/resources.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common3 "github.com/hyperledger/mchain/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// APIResource represents an API resource in the peer whose ACL is determined by the policy_ref field
type APIResource struct {
	PolicyRef string `protobuf:"bytes,1,opt,name=policy_ref,json=policyRef" json:"policy_ref,omitempty"`
}

func (m *APIResource) Reset()                    { *m = APIResource{} }
func (m *APIResource) String() string            { return proto.CompactTextString(m) }
func (*APIResource) ProtoMessage()               {}
func (*APIResource) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *APIResource) GetPolicyRef() string {
	if m != nil {
		return m.PolicyRef
	}
	return ""
}

// ChaincodeIdentifier identifies a piece of chaincode.  For a peer to accept invocations of
// this chaincode, the hash of the installed code must match, as must the version string
// included with the install command.
type ChaincodeIdentifier struct {
	Hash    []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *ChaincodeIdentifier) Reset()                    { *m = ChaincodeIdentifier{} }
func (m *ChaincodeIdentifier) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeIdentifier) ProtoMessage()               {}
func (*ChaincodeIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *ChaincodeIdentifier) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ChaincodeIdentifier) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// ChaincodeValidation instructs the peer how transactions for this chaincode should be
// validated.  The only validation mechanism which ships with mchain today is the standard
// 'vscc' validation mechanism.  This built in validation method utilizes an endorsement policy
// which checks that a sufficient number of signatures have been included.  The 'arguement'
// field encodes any parameters required by the validation implementation.
type ChaincodeValidation struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Argument []byte `protobuf:"bytes,2,opt,name=argument,proto3" json:"argument,omitempty"`
}

func (m *ChaincodeValidation) Reset()                    { *m = ChaincodeValidation{} }
func (m *ChaincodeValidation) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeValidation) ProtoMessage()               {}
func (*ChaincodeValidation) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *ChaincodeValidation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeValidation) GetArgument() []byte {
	if m != nil {
		return m.Argument
	}
	return nil
}

// VSCCArgs is passed (marshaled) as a parameter to the VSCC imlementation via the
// argument field of the ChaincodeValidation message.
type VSCCArgs struct {
	EndorsementPolicyRef string `protobuf:"bytes,1,opt,name=endorsement_policy_ref,json=endorsementPolicyRef" json:"endorsement_policy_ref,omitempty"`
}

func (m *VSCCArgs) Reset()                    { *m = VSCCArgs{} }
func (m *VSCCArgs) String() string            { return proto.CompactTextString(m) }
func (*VSCCArgs) ProtoMessage()               {}
func (*VSCCArgs) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *VSCCArgs) GetEndorsementPolicyRef() string {
	if m != nil {
		return m.EndorsementPolicyRef
	}
	return ""
}

// ChaincodeEndorsement instructs the peer how transactions should be endorsed.  The only
// endorsement mechanism which ships with the mchain today is the standard 'escc' mechanism.
// This code simply simulates the proposal to generate a RW set, then signs the result
// using the peer's local signing identity.
type ChaincodeEndorsement struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ChaincodeEndorsement) Reset()                    { *m = ChaincodeEndorsement{} }
func (m *ChaincodeEndorsement) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeEndorsement) ProtoMessage()               {}
func (*ChaincodeEndorsement) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *ChaincodeEndorsement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// ConfigTree encapsulates channel and resources configuration of a channel.
// Both configurations are represented as common.Config
type ConfigTree struct {
	ChannelConfig   *common3.Config `protobuf:"bytes,1,opt,name=channel_config,json=channelConfig" json:"channel_config,omitempty"`
	ResourcesConfig *common3.Config `protobuf:"bytes,2,opt,name=resources_config,json=resourcesConfig" json:"resources_config,omitempty"`
}

func (m *ConfigTree) Reset()                    { *m = ConfigTree{} }
func (m *ConfigTree) String() string            { return proto.CompactTextString(m) }
func (*ConfigTree) ProtoMessage()               {}
func (*ConfigTree) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *ConfigTree) GetChannelConfig() *common3.Config {
	if m != nil {
		return m.ChannelConfig
	}
	return nil
}

func (m *ConfigTree) GetResourcesConfig() *common3.Config {
	if m != nil {
		return m.ResourcesConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*APIResource)(nil), "protos.APIResource")
	proto.RegisterType((*ChaincodeIdentifier)(nil), "protos.ChaincodeIdentifier")
	proto.RegisterType((*ChaincodeValidation)(nil), "protos.ChaincodeValidation")
	proto.RegisterType((*VSCCArgs)(nil), "protos.VSCCArgs")
	proto.RegisterType((*ChaincodeEndorsement)(nil), "protos.ChaincodeEndorsement")
	proto.RegisterType((*ConfigTree)(nil), "protos.ConfigTree")
}

func init() { proto.RegisterFile("peer/resources.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4b, 0x4b, 0xeb, 0x40,
	0x14, 0xc7, 0x69, 0xb9, 0xdc, 0xdb, 0x9e, 0xf6, 0x56, 0x19, 0xab, 0x94, 0x82, 0x20, 0x59, 0xa9,
	0x48, 0x02, 0x3e, 0x16, 0xee, 0xac, 0xa1, 0x8b, 0xae, 0x2c, 0x51, 0xba, 0x70, 0x53, 0xa6, 0x93,
	0x93, 0x64, 0x20, 0x99, 0x09, 0x67, 0x52, 0xb1, 0x1b, 0x3f, 0xbb, 0x64, 0xa6, 0x8d, 0x05, 0xbb,
	0xca, 0x79, 0xfc, 0xce, 0x3f, 0xe7, 0x31, 0x30, 0x2c, 0x11, 0x29, 0x20, 0x34, 0x7a, 0x4d, 0x02,
	0x8d, 0x5f, 0x92, 0xae, 0x34, 0xfb, 0x6b, 0x3f, 0x66, 0x7c, 0x2a, 0x74, 0x51, 0x68, 0x15, 0x08,
	0xad, 0x12, 0x99, 0x56, 0x9f, 0x2e, 0xed, 0xdd, 0x40, 0x6f, 0x32, 0x9f, 0x45, 0xdb, 0x22, 0x76,
	0x0e, 0x50, 0xea, 0x5c, 0x8a, 0xcd, 0x92, 0x30, 0x19, 0xb5, 0x2e, 0x5a, 0x97, 0xdd, 0xa8, 0xeb,
	0x22, 0x11, 0x26, 0x5e, 0x08, 0x27, 0x61, 0xc6, 0xa5, 0x12, 0x3a, 0xc6, 0x59, 0x8c, 0xaa, 0x92,
	0x89, 0x44, 0x62, 0x0c, 0xfe, 0x64, 0xdc, 0x64, 0x96, 0xef, 0x47, 0xd6, 0x66, 0x23, 0xf8, 0xf7,
	0x81, 0x64, 0xa4, 0x56, 0xa3, 0xb6, 0x95, 0xd9, 0xb9, 0xde, 0x74, 0x4f, 0x64, 0xc1, 0x73, 0x19,
	0xf3, 0x4a, 0x6a, 0x55, 0x8b, 0x28, 0x5e, 0xe0, 0xf6, 0xa7, 0xd6, 0x66, 0x63, 0xe8, 0x70, 0x4a,
	0xd7, 0x05, 0xaa, 0xca, 0xaa, 0xf4, 0xa3, 0xc6, 0xf7, 0x9e, 0xa0, 0xb3, 0x78, 0x0d, 0xc3, 0x09,
	0xa5, 0x86, 0xdd, 0xc3, 0x19, 0xaa, 0x58, 0x93, 0xc1, 0x3a, 0xb5, 0xfc, 0x35, 0xc2, 0x70, 0x2f,
	0x3b, 0x6f, 0xa6, 0xb9, 0x86, 0x61, 0xd3, 0xc8, 0xf4, 0x07, 0x38, 0xd4, 0x89, 0xf7, 0x05, 0x10,
	0xda, 0xcd, 0xbd, 0x11, 0x22, 0x7b, 0x80, 0x81, 0xc8, 0xb8, 0x52, 0x98, 0x2f, 0xdd, 0x3e, 0x2d,
	0xdb, 0xbb, 0x1d, 0xf8, 0x6e, 0xcb, 0xbe, 0x63, 0xa3, 0xff, 0x5b, 0xca, 0xb9, 0xec, 0x11, 0x8e,
	0x9b, 0xf3, 0xec, 0x0a, 0xdb, 0x07, 0x0b, 0x8f, 0x1a, 0xce, 0x05, 0x9e, 0x5f, 0xc0, 0xd3, 0x94,
	0xfa, 0xd9, 0xa6, 0x44, 0xca, 0x31, 0x4e, 0x91, 0xfc, 0x84, 0xaf, 0x48, 0x0a, 0x77, 0x47, 0xe3,
	0xd7, 0xc7, 0x7f, 0xbf, 0x4a, 0x65, 0x95, 0xad, 0x57, 0xb5, 0x58, 0xb0, 0x87, 0x06, 0x0e, 0x0d,
	0x1c, 0x1a, 0xd4, 0xe8, 0xca, 0xbd, 0x8b, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0xfe,
	0xe3, 0xfb, 0x36, 0x02, 0x00, 0x00,
}
