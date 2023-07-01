package mtproto

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRpcError(e *status.Status) (err *TLRpcError) {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(e.Code()),
		ErrorMessage: e.Message(),
	}}
}

func (m *TLRpcError) IsOK() bool {
	if m == nil {
		return true
	}
	return m.GetErrorCode() == int32(codes.OK)
}

func (m *TLRpcError) Error() string {
	return m.DebugString()
}

func (m *TLRpcError) Code() int {
	return int(m.GetErrorCode())
}

func (m *TLRpcError) Message() string {
	return m.GetErrorMessage()
}

func (m *TLRpcError) Details() []interface{} {
	return nil
}
