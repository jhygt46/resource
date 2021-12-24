package scp

import (
	"io/ioutil"

	"github.com/povsister/scp"
)

func CopyFile(ip string, localfile string, remotefile string) error {

	privPEM, err := ioutil.ReadFile("/root/.ssh/id_rsa")
	if err != nil {
		return err
	}
	sshConf, err := scp.NewSSHConfigFromPrivateKey("root", privPEM)
	if err != nil {
		return err
	}
	scpClient, err := scp.NewClient(ip+":22", sshConf, &scp.ClientOption{})
	if err != nil {
		return err
	}
	defer scpClient.Close()
	err2 := scpClient.CopyFileFromRemote(localfile, remotefile, &scp.FileTransferOption{})
	if err2 != nil {
		return err2
	} else {
		return nil
	}

}
func CopyFolder(ip string, localfolder string, remotefolder string) error {

	privPEM, err := ioutil.ReadFile("/root/.ssh/id_rsa")
	if err != nil {
		return err
	}
	sshConf, err := scp.NewSSHConfigFromPrivateKey("root", privPEM, "buenanelson")
	if err != nil {
		return err
	}
	scpClient, err := scp.NewClient(ip+":22", sshConf, &scp.ClientOption{})
	if err != nil {
		return err
	}
	defer scpClient.Close()
	err2 := scpClient.CopyDirFromRemote(localfolder, remotefolder, &scp.DirTransferOption{})
	if err2 != nil {
		return err2
	} else {
		return nil
	}

}
