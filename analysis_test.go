package main

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_DoAnalysis(t *testing.T) {
	wd, _ := os.Getwd()
	println(wd, "----")
	Analysis := new(analysis)
	if err := Analysis.DoAnalysis(CallGraphType("pointer"), "", false, []string{filepath.Join(wd, "examples/main")}); err != nil {
		t.Errorf(err.Error())
	}

	
	Analysis.callgraph = 
}

