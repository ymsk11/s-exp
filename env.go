package main

import "fmt"

// Env is Environment
type Env struct {
	variables map[string]any
}

// CreateEnv create environment
func CreateEnv() *Env {
	return &Env{
		map[string]any{},
	}
}

func (env *Env) copy() *Env {
	vars := map[string]any{}
	for i, e := range env.variables {
		vars[i] = e
	}
	return &Env{
		vars,
	}
}

func (env *Env) set(name string, any any) {
	env.variables[name] = any
}

func (env *Env) get(name string) any {
	return env.variables[name]
}

func (env *Env) print() {
	fmt.Println(env.variables)
}
