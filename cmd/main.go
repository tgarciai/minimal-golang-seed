package main

import bootstrap "github.con/tgarcia/seed-golang-server/internal"

func main() {
	bootstrap, err := bootstrap.NewBootstrap()
	if err != nil {
		panic(err)
	}

	if err := bootstrap.Start(); err != nil {
		panic(err)
	}
}
