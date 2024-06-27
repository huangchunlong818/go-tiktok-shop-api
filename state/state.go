package state

import (
	"crypto/rand"
	"encoding/base64"
)

//随机数，防止跨站攻击

type State struct {
}

var newServer *State

func GetNewService() *State {
	if newServer == nil {
		newServer = &State{}
	}
	return newServer
}

func (s *State) GetState() (string, error) {
	// 生成随机字节序列
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// 转换为 base64 编码的字符串
	stateToken := base64.URLEncoding.EncodeToString(randomBytes)

	return stateToken, nil
}

func (s *State) DecodeState(stateToken string) error {
	// 将 base64 编码的字符串解码为字节序列
	_, err := base64.URLEncoding.DecodeString(stateToken)
	if err != nil {
		return err
	}

	return nil
}
