package main

import (
	"os"
	"os/signal"

	"github.com/mr7282/testingServer.git/server"
	"go.uber.org/zap"
)



func main() {
	lg := zap.NewExample()
	defer lg.Sync()

	serv := server.NewServer(lg)

	go func() {

		serv.StartServer()
	} ()

	lg.Info("Server is started.",
	zap.String("Port", serv.Port))

	stopSig := make(chan os.Signal, 1)
	signal.Notify(stopSig, os.Interrupt, os.Kill)
	<-stopSig
}

