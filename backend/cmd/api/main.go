package main

import "log"

func main() {
	app := &application{
		cfg: config{
			addr: ":8080",
		},
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
