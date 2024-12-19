//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: contrib/envoy/extensions/filters/http/language/v3alpha/language.proto

package v3alpha

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *Language) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Language) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Language) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.ClearRouteCache {
		i--
		if m.ClearRouteCache {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.SupportedLanguages) > 0 {
		for iNdEx := len(m.SupportedLanguages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SupportedLanguages[iNdEx])
			copy(dAtA[i:], m.SupportedLanguages[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.SupportedLanguages[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.DefaultLanguage) > 0 {
		i -= len(m.DefaultLanguage)
		copy(dAtA[i:], m.DefaultLanguage)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.DefaultLanguage)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Language) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DefaultLanguage)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.SupportedLanguages) > 0 {
		for _, s := range m.SupportedLanguages {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	if m.ClearRouteCache {
		n += 2
	}
	n += len(m.unknownFields)
	return n
}
