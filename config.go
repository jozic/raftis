package raftis

import (
	"encoding/json"
	"io"
	"os"
)

type ClusterConfig struct {
	NumSlots uint32
	Whoami   string // flotilla addr
	Shards   []Shard
}

type Shard struct {
	Slots []uint32
	Hosts []Host
}

type Host struct {
	RedisAddr    string // "192.168.0.4:8369"
	FlotillaAddr string // "192.168.0.4:1103"
	Group        string
}

func WriteConfig(c *ClusterConfig, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func WriteConfigFile(c *ClusterConfig, filePath string) error {
	w, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0755))
	if err != nil {
		return err
	}
	defer w.Close()
	return WriteConfig(c, w)
}

func ReadConfig(r io.Reader) (*ClusterConfig, error) {
	d := json.NewDecoder(r)
	ret1 := &ClusterConfig{}
	retErr := d.Decode(ret1)
	return ret1, retErr
}

func ReadConfigFile(filePath string) (*ClusterConfig, error) {
	r, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ReadConfig(r)
}
