package main

import (
	"os"
	"os/exec"
	"path/filepath"

	demo "github.com/saschagrunert/demo"
)

func main() {
	d := demo.New()

	d.Add(runPolicyWithKwctl(), "krew-wasm-demo", "WasmDay KubeCon Europe 22 krew-wasm demo")

	d.Run()
}

func runPolicyWithKwctl() *demo.Run {
	r := demo.NewRun(
		"Running kubectl plugins with WebAssembly",
	)

	r.Step(demo.S(
		"List plugins",
	), demo.S(
		"krew-wasm list",
	))

	r.Step(demo.S(
		"Pull kubectl-decoder plugin",
	), demo.S(
		"krew-wasm pull ghcr.io/flavio/krew-wasm-plugins/decoder:latest",
	))

	r.Step(demo.S(
		"List plugins",
	), demo.S(
		"krew-wasm list",
	))

	r.Step(demo.S(
		"Kubernetes is up and running",
	), demo.S(
		"kubectl get nodes -o wide",
	))

	r.Step(demo.S(
		"Create a generic secret",
	), demo.S(
		"kubectl create secret generic db-user-pass",
		"--from-literal=username=devuser",
		"--from-literal=password='very-secret!'",
	))

	r.Step(demo.S(
		"Print raw secret contents",
	), demo.S(
		"kubectl get secret db-user-pass -o json | jq",
	))

	r.Setup(setupKrewWasm)
	r.Cleanup(cleanupKrewWasm)

	return r
}

func setupKrewWasm() error {
	cleanupKrewWasm()
	return nil
}

func cleanupKrewWasm() error {
	exec.Command("kubectl", "delete", "secret", "db-user-pass").Run()
	os.RemoveAll(filepath.Join(os.Getenv("HOME"), ".krew-wasm"))
	os.RemoveAll(filepath.Join(os.Getenv("HOME"), ".cache", "krew-wasm"))
	return nil
}
