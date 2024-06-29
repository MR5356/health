package host

import (
	"github.com/MR5356/health"
	"github.com/gliderlabs/ssh"
	"net"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

var hostInfo = &HostInfo{
	PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4v1dRTNkTJdR/hlew+JjSyOywhEzyJQZoMImh9rxw8YHv60A
Pya8CbWf0Fn2v11ZQmkspPLqqjKqSg97HFIJiFA/A1yvjo/kMxEm58JTzZ4ONjub
1pAwwllxq9qqPHsSIspYndSGzKqZeFUYjeJi9LG06UaxLwG6036pUKmpmyScOIVq
CeE8uuudzbOa4qL4AxExx2ziuMa2b4BLBFH2bcsiz3xnpM4D4i7wubrC2ZXu9t9C
3VtZRC1QRRH/+k5Jpevd46XiBCO2U+yZlXefYYN0hZs/4o9XhGMkk3xQQaAh0cmJ
2bKsODqGnVSYll4zmDBFqBCWrYC/DrkYu8LGuQIDAQABAoIBAAyJujEsWYnxgdHr
TNS9GIb5/dHCaX3W0GOU1dQDP/90XNE1mSHj3dcbdSxpC7weR+PnM1QZJuUnz0gv
+rjhvy0MYa6p2/if8hbwB5odnctpw5czS7RaWpchoanAdk7I7YOGccamCpwbgEap
TKXpr1Wcy0VnZjoWel9cS7Xs6TVsgZ9lAYZwbxr17cy02GYn8ZmEnJIRkd2bqLlf
9beHibzkZzTvEyVmNWn9attLlAo4uqnusVBTIM4WkFeeb6/zvuEQVamOcV/cVcTS
yzIrbu+D0MVjbJ4wQGaQAI3SvHCOKraiilro7cflZhRmZABzMsAywxNlDw1KF1bZ
rgTq6EECgYEA+SXCBGWNtuIkef0i42S3E7n8ztzBsrzWisHKQ4oE5CEiZ9bGmHjp
V8IuXuREC8803o/nDBSw20TvLPovgRJlHMBzLJ/20faPMFDnb34gZXioZiO36c7g
XqGMEDX/fIZQVCk1s0NLoN6KdP/UFTghDCJO+4kJ9hGFdAbzCTRkWxUCgYEA6TuY
DWy+n1khCit2nUFcxnr95EhuX9kJdtmCC3ldGwm/6z+MyXsAelNI7N5X7cplPprs
GpDopPKO3h3+jJkHhoKFgwW6K5yXzX8pXuuXpKCHTjn9N/pyQRpBOCk00TxmfCJo
NK5/UpNzrVncs4nbnpe3wjCN81Mz0r4nlx/FlhUCgYEAvoHgmPwOIIlK2vx0cOvS
EYNli8fBVKqQYglMX9hpZQbxB/VyZaQOyMvyKRzh6yXFh2kBgntPwFo1iG7FZCxs
pE+DwduPH30ogAlc7iPDIdPg8DjfqChH6BQexUaE1PLe+XuJVElgyuFffZcg1BjX
LDAPtMZUl5NOOCoYLDGjiwECgYAVBKvVOeGL1qplkjkkPsmvkVHkCqr3tEVoEn3n
rs43K4/CEX4Mgisu2uaNghQGd+Db3XY9hqWQh9NcLPPNk9TbyFNj1VZLq9b6S+Vc
inql+Vl4MT2il81IFKef+gaqfHj34tnlNXx/4o3gJ2L+QwQprJ6Av4NrVCotabmD
ovdZaQKBgQC2ikNy+o3vId7IhgI/PBHJrH6YlgeznqJQiCXHLmcjxSw4Pek9wQIl
rv5/8rsSIlm1fN1j02oiolJkV8ev+015g68Vhd1mit/3/zNBH4+M7Afq72VWdqvQ
pMpriqHVXt5l9++XHgGDh7anXUlF/Ee+RaBxze4GjI4qRsLR0IpJkg==
-----END RSA PRIVATE KEY-----`,
	Host:     "host",
	Port:     22,
	Username: "username",
	Password: "password",
}

func TestNewSSHChecker(t *testing.T) {
	sshChecker := NewSSHChecker(hostInfo)
	if !reflect.DeepEqual(sshChecker.hostInfo, hostInfo) {
		t.Errorf("Expected hostInfo %v, got %v", hostInfo, sshChecker.hostInfo)
	}
	if sshChecker.timeout != time.Second*5 {
		t.Errorf("Expected timeout %v, got %v", time.Second*5, sshChecker.timeout)
	}
}

func TestNewSSHCheckerWithTimeout(t *testing.T) {
	sshChecker := NewSSHCheckerWithTimeout(hostInfo, time.Second*10)
	if !reflect.DeepEqual(sshChecker.hostInfo, hostInfo) {
		t.Errorf("Expected hostInfo %v, got %v", hostInfo, sshChecker.hostInfo)
	}
	if sshChecker.timeout != time.Second*10 {
		t.Errorf("Expected timeout %v, got %v", time.Second*10, sshChecker.timeout)
	}
}

func TestSSHChecker_Check_Success(t *testing.T) {
	s := &ssh.Server{
		Handler: func(session ssh.Session) {
			session.Write([]byte("Hello from SSH server"))
			session.User()
		},
		PasswordHandler: func(ctx ssh.Context, password string) bool {
			if ctx.User() == "username" && password == "password" {
				return true
			}
			return false
		},
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	go s.Serve(listener)
	defer s.Close()

	port, _ := strconv.Atoi(strings.Split(listener.Addr().String(), ":")[1])

	sshChecker := NewSSHChecker(&HostInfo{
		Host:     strings.Split(listener.Addr().String(), ":")[0],
		Port:     uint16(port),
		Username: "username",
		Password: "password",
	})

	result := sshChecker.Check()
	if result.Status != health.StatusUp {
		t.Errorf("Expected status %s, got %s", health.StatusUp, result.Status)
	}
}

func TestSSHChecker_Check_Failure(t *testing.T) {
	sshChecker := NewSSHChecker(&HostInfo{
		Host:     "aasd",
		Port:     22,
		Username: "username",
		Password: "password",
	})

	result := sshChecker.Check()
	if result.Status != health.StatusDown {
		t.Errorf("Expected status %s, got %s", health.StatusUp, result.Status)
	}
}

func TestSSHChecker_Check_Failure_PrivateKey(t *testing.T) {
	sshChecker := NewSSHChecker(hostInfo)

	result := sshChecker.Check()
	if result.Status != health.StatusDown {
		t.Errorf("Expected status %s, got %s", health.StatusUp, result.Status)
	}
}
