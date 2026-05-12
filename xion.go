package main

import "fmt"

type Xion interface {
	GetData() []uint64
}

type Data struct {
	xion          Xion
	Received      uint64
	Sent          uint64
	DownloadSpeed float64
	UploadSpeed   float64
}

func NewData() *Data {
	return &Data{
		xion:          NewXion(),
		Received:      0,
		Sent:          0,
		DownloadSpeed: 0,
		UploadSpeed:   0,
	}
}

func (d *Data) UpdateData() {
	cur := d.xion.GetData()

	d.DownloadSpeed = float64(cur[0] - d.Received)
	d.UploadSpeed = float64(cur[1] - d.Sent)

	d.Received = cur[0]
	d.Sent = cur[1]
}

func (d Data) Print() {
	fmt.Println(d.Received)
	fmt.Println(d.Sent)
	fmt.Println(d.DownloadSpeed)
	fmt.Println(d.UploadSpeed)
}
