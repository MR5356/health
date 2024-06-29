package host

import (
	"fmt"
	"github.com/MR5356/health"
	"golang.org/x/crypto/ssh"
	"time"
)

type SSHChecker struct {
	hostInfo *HostInfo
	timeout  time.Duration
}

type SSHResult struct {
	Error error `json:"error"`
}

type HostInfo struct {
	PrivateKey string `json:"privateKey"`
	Host       string `json:"host"`
	Port       uint16 `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Passphrase string `json:"passphrase"`
}

func (h *HostInfo) GetAuthMethods() []ssh.AuthMethod {
	authMethods := make([]ssh.AuthMethod, 0)

	if len(h.PrivateKey) > 0 {
		if len(h.Passphrase) > 0 {
			signer, err := ssh.ParsePrivateKeyWithPassphrase([]byte(h.PrivateKey), []byte(h.Passphrase))
			if err == nil {
				authMethods = append(authMethods, ssh.PublicKeys(signer))
			}
		} else {
			signer, err := ssh.ParsePrivateKey([]byte(h.PrivateKey))
			if err == nil {
				authMethods = append(authMethods, ssh.PublicKeys(signer))
			}
		}
	}

	if len(h.Password) > 0 {
		authMethods = append(authMethods, ssh.Password(h.Password))
	}
	return authMethods
}

func NewSSHChecker(hostInfo *HostInfo) *SSHChecker {
	return &SSHChecker{
		hostInfo: hostInfo,
		timeout:  time.Second * 5,
	}
}

func NewSSHCheckerWithTimeout(hostInfo *HostInfo, timeout time.Duration) *SSHChecker {
	return &SSHChecker{
		hostInfo: hostInfo,
		timeout:  timeout,
	}
}

func (sc *SSHChecker) Check() (result *health.Health) {
	result = health.NewHealth()

	start := time.Now()
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sc.hostInfo.Host, sc.hostInfo.Port), &ssh.ClientConfig{
		User:            sc.hostInfo.Username,
		Auth:            sc.hostInfo.GetAuthMethods(),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         sc.timeout,
	})
	rtt := time.Since(start).Milliseconds()
	result.SetRTT(rtt)
	if sshClient != nil {
		defer sshClient.Close()
	}
	if err != nil {
		result.Down()
		result.SetResult(&SSHResult{Error: err})
		return
	}
	result.Up()
	return
}
