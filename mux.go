/*
   Create: 2023/8/11
   Project: Gungnir
   Github: https://github.com/landers1037
   Copyright Renj
*/

package gungnir

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func serveMux(s http.Handler) {
	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%d", Host, Port),
		Handler: s,
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	ctx := context.Background()

	go func() {
		if c := <-sig; c == os.Interrupt {
			srv.Shutdown(ctx)
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("%s stopped: %s\n", APPName, err)
	}
}
