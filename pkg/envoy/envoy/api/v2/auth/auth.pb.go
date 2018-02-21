// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	envoy/api/v2/auth/auth.proto
	envoy/api/v2/auth/cert.proto

It has these top-level messages:
	AuthAction
	TlsParameters
	TlsCertificate
	TlsSessionTicketKeys
	CertificateValidationContext
	CommonTlsContext
	UpstreamTlsContext
	DownstreamTlsContext
	SdsSecretConfig
	Secret
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Should we do white-list or black-list style access control.
type AuthAction_ActionType int32

const (
	// Request matches all rules are allowed, otherwise denied.
	AuthAction_ALLOW AuthAction_ActionType = 0
	// Request matches all rules or missing required auth fields are denied,
	// otherwise allowed.
	AuthAction_DENY AuthAction_ActionType = 1
)

var AuthAction_ActionType_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}
var AuthAction_ActionType_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x AuthAction_ActionType) String() string {
	return proto.EnumName(AuthAction_ActionType_name, int32(x))
}
func (AuthAction_ActionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type AuthAction struct {
	ActionType AuthAction_ActionType `protobuf:"varint,1,opt,name=action_type,json=actionType,enum=envoy.api.v2.auth.AuthAction_ActionType" json:"action_type,omitempty"`
	// List of rules
	Rules []*AuthAction_Rule `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
}

func (m *AuthAction) Reset()                    { *m = AuthAction{} }
func (m *AuthAction) String() string            { return proto.CompactTextString(m) }
func (*AuthAction) ProtoMessage()               {}
func (*AuthAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthAction) GetActionType() AuthAction_ActionType {
	if m != nil {
		return m.ActionType
	}
	return AuthAction_ALLOW
}

func (m *AuthAction) GetRules() []*AuthAction_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Logic AND that requires all rules match.
type AuthAction_AndRule struct {
	Rules []*AuthAction_Rule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *AuthAction_AndRule) Reset()                    { *m = AuthAction_AndRule{} }
func (m *AuthAction_AndRule) String() string            { return proto.CompactTextString(m) }
func (*AuthAction_AndRule) ProtoMessage()               {}
func (*AuthAction_AndRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *AuthAction_AndRule) GetRules() []*AuthAction_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Logic OR that requires at least one rule matches.
type AuthAction_OrRule struct {
	Rules []*AuthAction_Rule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *AuthAction_OrRule) Reset()                    { *m = AuthAction_OrRule{} }
func (m *AuthAction_OrRule) String() string            { return proto.CompactTextString(m) }
func (*AuthAction_OrRule) ProtoMessage()               {}
func (*AuthAction_OrRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *AuthAction_OrRule) GetRules() []*AuthAction_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Check peer identity using X.509 certificate.
type AuthAction_X509Rule struct {
	// How to validate peer certificates.
	ValidationContext *CertificateValidationContext `protobuf:"bytes,3,opt,name=validation_context,json=validationContext" json:"validation_context,omitempty"`
}

func (m *AuthAction_X509Rule) Reset()                    { *m = AuthAction_X509Rule{} }
func (m *AuthAction_X509Rule) String() string            { return proto.CompactTextString(m) }
func (*AuthAction_X509Rule) ProtoMessage()               {}
func (*AuthAction_X509Rule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *AuthAction_X509Rule) GetValidationContext() *CertificateValidationContext {
	if m != nil {
		return m.ValidationContext
	}
	return nil
}

// Element type of AndRule/OrRule, it chooses among different type of rule.
type AuthAction_Rule struct {
	// Types that are valid to be assigned to RuleSpecifier:
	//	*AuthAction_Rule_AndRule
	//	*AuthAction_Rule_OrRule
	//	*AuthAction_Rule_X509Rule
	RuleSpecifier isAuthAction_Rule_RuleSpecifier `protobuf_oneof:"rule_specifier"`
}

func (m *AuthAction_Rule) Reset()                    { *m = AuthAction_Rule{} }
func (m *AuthAction_Rule) String() string            { return proto.CompactTextString(m) }
func (*AuthAction_Rule) ProtoMessage()               {}
func (*AuthAction_Rule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

type isAuthAction_Rule_RuleSpecifier interface {
	isAuthAction_Rule_RuleSpecifier()
}

type AuthAction_Rule_AndRule struct {
	AndRule *AuthAction_AndRule `protobuf:"bytes,1,opt,name=and_rule,json=andRule,oneof"`
}
type AuthAction_Rule_OrRule struct {
	OrRule *AuthAction_OrRule `protobuf:"bytes,2,opt,name=or_rule,json=orRule,oneof"`
}
type AuthAction_Rule_X509Rule struct {
	X509Rule *AuthAction_X509Rule `protobuf:"bytes,3,opt,name=x509_rule,json=x509Rule,oneof"`
}

func (*AuthAction_Rule_AndRule) isAuthAction_Rule_RuleSpecifier()  {}
func (*AuthAction_Rule_OrRule) isAuthAction_Rule_RuleSpecifier()   {}
func (*AuthAction_Rule_X509Rule) isAuthAction_Rule_RuleSpecifier() {}

func (m *AuthAction_Rule) GetRuleSpecifier() isAuthAction_Rule_RuleSpecifier {
	if m != nil {
		return m.RuleSpecifier
	}
	return nil
}

func (m *AuthAction_Rule) GetAndRule() *AuthAction_AndRule {
	if x, ok := m.GetRuleSpecifier().(*AuthAction_Rule_AndRule); ok {
		return x.AndRule
	}
	return nil
}

func (m *AuthAction_Rule) GetOrRule() *AuthAction_OrRule {
	if x, ok := m.GetRuleSpecifier().(*AuthAction_Rule_OrRule); ok {
		return x.OrRule
	}
	return nil
}

func (m *AuthAction_Rule) GetX509Rule() *AuthAction_X509Rule {
	if x, ok := m.GetRuleSpecifier().(*AuthAction_Rule_X509Rule); ok {
		return x.X509Rule
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AuthAction_Rule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AuthAction_Rule_OneofMarshaler, _AuthAction_Rule_OneofUnmarshaler, _AuthAction_Rule_OneofSizer, []interface{}{
		(*AuthAction_Rule_AndRule)(nil),
		(*AuthAction_Rule_OrRule)(nil),
		(*AuthAction_Rule_X509Rule)(nil),
	}
}

func _AuthAction_Rule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AuthAction_Rule)
	// rule_specifier
	switch x := m.RuleSpecifier.(type) {
	case *AuthAction_Rule_AndRule:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AndRule); err != nil {
			return err
		}
	case *AuthAction_Rule_OrRule:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OrRule); err != nil {
			return err
		}
	case *AuthAction_Rule_X509Rule:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.X509Rule); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AuthAction_Rule.RuleSpecifier has unexpected type %T", x)
	}
	return nil
}

func _AuthAction_Rule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AuthAction_Rule)
	switch tag {
	case 1: // rule_specifier.and_rule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AuthAction_AndRule)
		err := b.DecodeMessage(msg)
		m.RuleSpecifier = &AuthAction_Rule_AndRule{msg}
		return true, err
	case 2: // rule_specifier.or_rule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AuthAction_OrRule)
		err := b.DecodeMessage(msg)
		m.RuleSpecifier = &AuthAction_Rule_OrRule{msg}
		return true, err
	case 3: // rule_specifier.x509_rule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AuthAction_X509Rule)
		err := b.DecodeMessage(msg)
		m.RuleSpecifier = &AuthAction_Rule_X509Rule{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AuthAction_Rule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AuthAction_Rule)
	// rule_specifier
	switch x := m.RuleSpecifier.(type) {
	case *AuthAction_Rule_AndRule:
		s := proto.Size(x.AndRule)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthAction_Rule_OrRule:
		s := proto.Size(x.OrRule)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthAction_Rule_X509Rule:
		s := proto.Size(x.X509Rule)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AuthAction)(nil), "envoy.api.v2.auth.AuthAction")
	proto.RegisterType((*AuthAction_AndRule)(nil), "envoy.api.v2.auth.AuthAction.AndRule")
	proto.RegisterType((*AuthAction_OrRule)(nil), "envoy.api.v2.auth.AuthAction.OrRule")
	proto.RegisterType((*AuthAction_X509Rule)(nil), "envoy.api.v2.auth.AuthAction.X509Rule")
	proto.RegisterType((*AuthAction_Rule)(nil), "envoy.api.v2.auth.AuthAction.Rule")
	proto.RegisterEnum("envoy.api.v2.auth.AuthAction_ActionType", AuthAction_ActionType_name, AuthAction_ActionType_value)
}

func init() { proto.RegisterFile("envoy/api/v2/auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xdd, 0x6a, 0xe2, 0x40,
	0x18, 0x86, 0x13, 0x7f, 0xe3, 0x17, 0x10, 0x0d, 0x7b, 0x20, 0x61, 0x59, 0x5c, 0xd9, 0x2d, 0x39,
	0x4a, 0x24, 0x45, 0xa8, 0x47, 0x25, 0xb1, 0x82, 0x05, 0xa9, 0x10, 0x4a, 0xff, 0x0e, 0x2a, 0xd3,
	0x38, 0xea, 0x14, 0xc9, 0x84, 0x71, 0x12, 0xf4, 0x2e, 0x7a, 0x19, 0xbd, 0xae, 0xd2, 0x0b, 0x29,
	0x99, 0xb1, 0x0a, 0xad, 0x58, 0x4a, 0x4f, 0x92, 0x97, 0x6f, 0x78, 0x9f, 0x2f, 0x79, 0x18, 0xf8,
	0x8d, 0xa3, 0x94, 0xae, 0x1d, 0x14, 0x13, 0x27, 0x75, 0x1d, 0x94, 0xf0, 0xb9, 0x78, 0xd8, 0x31,
	0xa3, 0x9c, 0x1a, 0x75, 0x71, 0x6a, 0xa3, 0x98, 0xd8, 0xa9, 0x6b, 0x67, 0x07, 0xe6, 0x9e, 0x42,
	0x88, 0x19, 0x97, 0x05, 0xf3, 0xd7, 0x8c, 0xce, 0xa8, 0x88, 0x4e, 0x96, 0xe4, 0xb4, 0xf5, 0x54,
	0x04, 0xf0, 0x12, 0x3e, 0xf7, 0x42, 0x4e, 0x68, 0x64, 0x9c, 0x83, 0x8e, 0x44, 0x1a, 0xf3, 0x75,
	0x8c, 0x1b, 0x6a, 0x53, 0xb5, 0xaa, 0xae, 0x65, 0x7f, 0xda, 0x65, 0xef, 0x3a, 0xb6, 0x7c, 0x5d,
	0xae, 0x63, 0x1c, 0x00, 0xda, 0x66, 0xe3, 0x04, 0x8a, 0x2c, 0x59, 0xe0, 0x65, 0x23, 0xd7, 0xcc,
	0x5b, 0xba, 0xdb, 0x3a, 0x0c, 0x09, 0x92, 0x05, 0x0e, 0x64, 0xc1, 0xec, 0x41, 0xd9, 0x8b, 0x26,
	0xd9, 0x64, 0x07, 0x51, 0xbf, 0x0b, 0xf1, 0xa1, 0x34, 0x62, 0x3f, 0x64, 0x3c, 0x82, 0x76, 0xd3,
	0x69, 0x77, 0x05, 0xe5, 0x1e, 0x8c, 0x14, 0x2d, 0xc8, 0x04, 0x09, 0x3b, 0x21, 0x8d, 0x38, 0x5e,
	0xf1, 0x46, 0xbe, 0xa9, 0x5a, 0xba, 0xeb, 0xec, 0x41, 0xf6, 0x30, 0xe3, 0x64, 0x4a, 0x42, 0xc4,
	0xf1, 0xd5, 0xb6, 0xd7, 0x93, 0xb5, 0xa0, 0x9e, 0x7e, 0x1c, 0x99, 0xaf, 0x2a, 0x14, 0xc4, 0x22,
	0x1f, 0x34, 0x14, 0x4d, 0xc6, 0xd9, 0x17, 0x08, 0xff, 0xba, 0xfb, 0xff, 0x0b, 0xff, 0xd2, 0xd5,
	0x40, 0x09, 0xca, 0x68, 0xa3, 0xed, 0x14, 0xca, 0x94, 0x49, 0x44, 0x4e, 0x20, 0xfe, 0x1d, 0x46,
	0x48, 0x53, 0x03, 0x25, 0x28, 0x51, 0xe9, 0xac, 0x0f, 0x95, 0x55, 0xa7, 0xdd, 0x95, 0x08, 0xf9,
	0x93, 0x47, 0x87, 0x11, 0xef, 0xa2, 0x06, 0x4a, 0xa0, 0xad, 0x36, 0xd9, 0xaf, 0x41, 0x35, 0x23,
	0x8c, 0x97, 0x31, 0x0e, 0xc9, 0x94, 0x60, 0xd6, 0xfa, 0x0b, 0xb0, 0xbb, 0x2f, 0x46, 0x05, 0x8a,
	0xde, 0x70, 0x38, 0xba, 0xae, 0x29, 0x86, 0x06, 0x85, 0xb3, 0xfe, 0xc5, 0x6d, 0x4d, 0xf5, 0xe1,
	0xf9, 0xe5, 0x8f, 0x7a, 0x57, 0xc8, 0x16, 0x3c, 0x94, 0xc4, 0x2d, 0x3d, 0x7e, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xa7, 0xc4, 0xd9, 0x6e, 0x0c, 0x03, 0x00, 0x00,
}