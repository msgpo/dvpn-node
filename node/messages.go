package node

import (
	"encoding/json"

	"github.com/pkg/errors"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"

	"github.com/ironman0x7b2/vpn-node/types"
)

type MsgBandwidthSignature struct {
	ID                 sdkTypes.ID        `json:"id"`
	Index              uint64             `json:"index"`
	Bandwidth          sdkTypes.Bandwidth `json:"bandwidth"`
	NodeOwnerSignature []byte             `json:"node_owner_signature"`
	ClientSignature    []byte             `json:"client_signature"`
}

func NewMsgBandwidthSignature(id sdkTypes.ID, index uint64, bandwidth sdkTypes.Bandwidth,
	nodeOwnerSignature, clientSignature []byte) *types.Msg {

	msg := MsgBandwidthSignature{
		ID:                 id,
		Index:              index,
		Bandwidth:          bandwidth,
		NodeOwnerSignature: nodeOwnerSignature,
		ClientSignature:    clientSignature,
	}
	data, _ := json.Marshal(msg)

	return &types.Msg{
		Type: msg.Type(),
		Data: data,
	}
}

func (msg *MsgBandwidthSignature) Type() string {
	return "MsgBandwidthSignature"
}

func (msg *MsgBandwidthSignature) Validate() error {
	if msg.ID.String() == "" {
		return errors.Errorf("id is empty")
	}
	if !msg.Bandwidth.AllPositive() {
		return errors.Errorf("bandwidth is not positive")
	}
	if msg.NodeOwnerSignature == nil {
		return errors.Errorf("node_owner_signature is empty")
	}
	if msg.ClientSignature == nil {
		return errors.Errorf("client_signature is empty")
	}

	return nil
}

type MsgError struct {
	Code    int8   `json:"code"`
	Message string `json:"message"`
}

func NewMsgError(code int8, message string) *types.Msg {
	msg := MsgError{
		Code:    code,
		Message: message,
	}
	data, _ := json.Marshal(msg)

	return &types.Msg{
		Type: msg.Type(),
		Data: data,
	}
}

func (msg *MsgError) Type() string {
	return "MsgError"
}

func (msg *MsgError) Validate() error {
	return nil
}
