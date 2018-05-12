// Copyright (c) 2018 Zededa, Inc.
// All rights reserved.

package diskmetrics

import (
	"encoding/json"
	"os"
	"os/exec"
)

// Matches the json output of qemu-img info
type ImgInfo struct {
	VirtualSize uint64 `json:"virtual-size"`
	Filename    string `json:"filename"`
	ClusterSize uint64 `json:"cluster-size"`
	Format      string `json:"format"`
	ActualSize  uint64 `json:"actual-size"`
	DirtyFlag   bool   `json:"dirty-flag"`
}

func GetImgInfo(diskfile string) (*ImgInfo, error) {
	var imgInfo ImgInfo

	if _, err := os.Stat(diskfile); err != nil {
		return nil, err
	}
	output, err := exec.Command("/usr/lib/xen/bin/qemu-img",
		"info", "--output=json", diskfile).Output()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(output, &imgInfo); err != nil {
		return nil, err
	}
	return &imgInfo, nil
}
