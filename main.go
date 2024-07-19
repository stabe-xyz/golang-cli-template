package main

import "{{ stabe.application_name }}/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
