//go:build linux

package main

import (
	"bytes"
	"os"
	"strconv"
)

type LinuxXion struct {
	RECEIVE_PATH string
	SEND_PATH    string
}

func NewXion() Xion {
	return &LinuxXion{
		RECEIVE_PATH: `/sys/class/net/eth0/statistics/rx_bytes`,
		SEND_PATH:    `/sys/class/net/eth0/statistics/tx_bytes`,
	}
}

func (l *LinuxXion) GetData() []uint64 {
	rec, _ := os.ReadFile(l.RECEIVE_PATH)
	sen, _ := os.ReadFile(l.SEND_PATH)
	rec = bytes.TrimSpace(rec)
	sen = bytes.TrimSpace(sen)

	r, err := strconv.ParseUint(string(rec), 10, 64)
	if err != nil {
		return []uint64{0, 0}
	}

	s, err := strconv.ParseUint(string(sen), 10, 64)
	if err != nil {
		return []uint64{r, 0}
	}

	return []uint64{r, s}
}
