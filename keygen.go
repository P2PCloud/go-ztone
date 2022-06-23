package one

import (
	"fmt"
	"os/exec"
	"strings"
)

const idToolExecutablePath = "/usr/sbin/zerotier-idtool"

func GenerateKeys() (string, string, error) {
	stdOut, err := exec.Command(idToolExecutablePath, "generate").Output()
	if err != nil {
		return "", "", err
	}

	identitySecret := string(stdOut)

	parts := strings.Split(identitySecret, ":")
	if len(parts) != 4 {
		return "", "", fmt.Errorf("unexpected output from zerotier-idtool: %s", stdOut)
	}

	identityPublic := strings.Join(parts[0:3], ":")
	return identitySecret, identityPublic, nil
}
