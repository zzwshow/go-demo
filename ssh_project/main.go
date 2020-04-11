package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

// 需要远程操作的主机
type HostProfile struct {
	Host       string
	Port       int
	AccessType string //接入类型分password 和sshkey ,默认为sshkey
	UserName   string
	Password   string
	KeyPath    string
}

func InitHostStruct(host string, port int, accessType, userName, password, keyPath string) *HostProfile {
	hostP := &HostProfile{
		Host:       host,
		Port:       port,
		AccessType: accessType,
		UserName:   userName,
		Password:   password,
		KeyPath:    keyPath,
	}
	return hostP
}

func main() {
	host := InitHostStruct("10.206.36.31", 5622, "", "tomcat", "", "/Users/zzw/.ssh/id_rsa")
	sshClient, err := NewSshClient(host)
	if err != nil {
		fmt.Println(err)
		log.Println(" 获取ssh client 失败...")
	}
	// 存在返回true,不存在返回false  (-f  标识检查文件，而非目录，目录请使用 -d)
	cmd := `test -f /usr/local/tomcat/startTomcat.sh && echo true || echo false`

	outstr, err := runCommand(sshClient, cmd)
	if err != nil {
		fmt.Println("ssh run command error :", err)
		return
	}
	fmt.Println("result: ", outstr)

}

func NewSshClient(h *HostProfile) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout: time.Second * 5,
		User:    h.UserName,
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if h.AccessType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(h.Password)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(h.KeyPath)}
	}
	addr := fmt.Sprintf("%s:%d", h.Host, h.Port)
	c, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}
	return c, nil
}
func hostKeyCallBackFunc(host string) ssh.HostKeyCallback {
	hostPath, err := homedir.Expand("~/.ssh/known_hosts")
	if err != nil {
		log.Fatal("find known_hosts's home dir failed", err)
	}
	file, err := os.Open(hostPath)
	if err != nil {
		log.Fatal("can't find known_host file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}
	if hostKey == nil {
		log.Fatalf("no hostkey for %s,%v", host, err)
	}
	return ssh.FixedHostKey(hostKey)
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// CreateUserOfRole the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
func runCommand(client *ssh.Client, command string) (stdout string, err error) {
	session, err := client.NewSession()
	if err != nil {
		log.Print(err)
		return
	}
	defer session.Close()

	var buf bytes.Buffer
	session.Stdout = &buf
	err = session.Run(command)
	if err != nil {
		log.Print(err)
		return
	}
	stdout = string(buf.Bytes())

	return
}
