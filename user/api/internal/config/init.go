package config

import (
	"io"
	"os"
)

func Init(c *Config) {
	c.rsaInit()
}

func (c *Config) rsaInit() {
	public, err := os.Open(c.Rsa.PublicFile)
	if err != nil {
		panic("初始化rsa-public :" + err.Error())
	}

	private, err := os.Open(c.Rsa.PrivateFile)
	if err != nil {
		panic("初始化rsa-private :" + err.Error())
	}

	defer private.Close()
	defer public.Close()

	pb, _ := io.ReadAll(public)
	rb, _ := io.ReadAll(private)
	c.Rsa.PublicKey = string(pb)
	c.Rsa.PrivateKey = string(rb)

}
