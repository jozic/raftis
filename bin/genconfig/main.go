package main

import (
	"flag"
	"fmt"
	"github.com/jbooth/raftis/config"
	"log"
	"os"
	"strconv"
	"strings"
)

// first arg is MODE, either "singlenode" or "cluster"
// if singlenode, 2nd arg is output directory, 3rd arg is number of shards.  we'll generate 3 datacenters "dc1,dc2,dc3" and a node in each for each shard
// if cluster, 2nd arg is output directory, 3rd arg is a TSV file denoting datacenter,host.
func main() {
	flag.Parse()
	args := flag.Args()
	fmt.Printf("Args : %+v\n", args)
	if len(args) < 3 {
		usage(args)
		return
	}
	mode := args[0]
	outDir := args[1]

	fmt.Printf("mode %s\n", mode)

	if strings.ToLower(mode) == "singlenode" {
		numShards, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}
		cfgs := config.Singlenode(numShards, 3)
		writeConfigs(cfgs, outDir)
	} else if strings.ToLower(mode) == "cluster" {

	} else {
		usage(args)
		return
	}

}

func usage(args []string) {
	log.Printf("First arg should either be 'singlenode' or 'cluster'.  Args provided : %+v", args)
	log.Printf(`Usage: \
		first arg is MODE, either "singlenode" or "cluster" \
		if singlenode, 2nd arg is output directory, 3rd arg is number of shards.  we'll generate 3 datacenters "dc1,dc2,dc3" and a node in each for each shard \
		if cluster, 2nd arg is output directory, 3rd arg is a TSV file denoting datacenter,host.`)

}

func writeConfigs(cfgs []config.ClusterConfig, outDir string) error {
	err := os.MkdirAll(outDir, 0777)
	if err != nil {
		return err
	}
	for _, cfg := range cfgs {
		outPath := outDir + "/" + cfg.Me.RedisAddr + ".conf"
		err = config.WriteConfigFile(&cfg, outPath)
		if err != nil {
			return err
		}
	}
	return nil
}