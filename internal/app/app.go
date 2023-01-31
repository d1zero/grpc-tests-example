package app

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"go-grpc-tests/config"
	"go-grpc-tests/internal/controller/grpc"
	"go-grpc-tests/internal/domain/service"
	"go-grpc-tests/internal/repository"
	"go-grpc-tests/pkg/govalidator"
	pb "go-grpc-tests/pkg/proto/bank/account"
	gogrpc "google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	validator := govalidator.New()

	log := zerolog.New(os.Stdout)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = validator.Validate(context.Background(), cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	log = log.Level(zerolog.Level(*cfg.Logger.Level)).With().Timestamp().Logger()

	defer log.Info().Msg("Application has been shut down")

	log.Debug().Msg("Loaded configuration")

	db, err := sqlx.Open("sqlite3", "db/account.db")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	accountRepo := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepo)
	accountController := grpc.NewAccountContoller(accountService)

	grpcServer := gogrpc.NewServer()
	defer grpcServer.GracefulStop()

	pb.RegisterDepositServiceServer(grpcServer, accountController)

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port))
		if err != nil {
			log.Fatal().Msgf("tcp sock: %s", err.Error())
		}
		defer lis.Close()

		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal().Msgf("GRPC server: %s", err.Error())
		}
	}()

	log.Debug().Msg("Started GRPC server")

	log.Info().Msg("Application has started")

	exit := make(chan os.Signal)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit
}
