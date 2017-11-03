package main

import (
	"os/exec"
	"testing"
)

func TestRunFaasProc(t *testing.T) {
	obs, ebs, err := runFaasProc(exec.Command("./test-faas-proc.sh"), []byte("input\n"))
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if string(obs) != "input\n" {
		t.Errorf("out: '%s'", obs)
	}
	if string(ebs) != "Err\n" {
		t.Errorf("err: '%s'", ebs)
	}
}
