/*
Copyright © 2022 NoStress.dev
*/
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nostressdev/tvx/internal/clicker"
	"github.com/nostressdev/tvx/internal/storages/postgresStorage"
	"github.com/nostressdev/tvx/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
	"net/http"
	"time"
	"regexp"
	"github.com/spf13/viper"
)

var (
	host     = "0.0.0.0"
	port     = "9010"
	httpHost = "0.0.0.0"
	httpPort = "443"
)

func createGrpcNetworkListener() net.Listener {
	addr := fmt.Sprintf("%s:%s", host, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	return listener
}

func CustomHttpHeadersMatcher(key string) (string, bool) {
	switch key {
	case "Authorization":
		return "AccessToken", true
	case "Access-Control-Allow-Origin":
		return "*", true // Allow requests from any domain
	case "Access-Control-Allow-Methods":
		return "GET, POST, OPTIONS", true // Allow GET, POST and OPTIONS requests
	case "Access-Control-Allow-Headers":
		return "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization", true // Allow these headers
	case "Access-Control-Expose-Headers":
		return "Authorization", true // Expose the Authorization header
	default:
		return key, false
	}
}

func allowedOrigin(origin string) bool {
    if viper.GetString("cors") == "*" {
        return true
    }
    if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
        return true
    }
    return false
}

func cors(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if allowedOrigin(r.Header.Get("Origin")) {
            w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
        }
        if r.Method == "OPTIONS" {
            return
        }
        h.ServeHTTP(w, r)
    })
}

func main() {
	time.Sleep(5 * time.Second)
	dsn := "host=postgres user=habrpguser password=pgpwd4habr dbname=habrdb port=5432 sslmode=disable"
	// dsn := "host=localhost user=postgres  dbname=template1 port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	server := grpc.NewServer()
	st, err := postgresStorage.NewPostgresStorage(db)
	if err != nil {
		panic("failed create storage: " + err.Error())
	}
	service := clicker.NewTvxServer(st)
	err = st.GenerateProducts()
	if err != nil {
		panic(err)
	}

	pb.RegisterClickerBackendServer(server, &service)
	reflection.Register(server)

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(CustomHttpHeadersMatcher),
	)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterClickerBackendHandlerFromEndpoint(
		context.Background(),
		mux,
		fmt.Sprintf("%s:%s", host, port),
		opts,
	); err != nil {
		panic(err)
	}

	go func() {
		log.Printf("Start serving http\n")
		// ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", httpHost, httpPort))
		// if err != nil {
		// log.Println(err)
		// return
		// }

		// defer ln.Close()
	   
		
		// if err := http.Serve(ln, cors(mux)); err != nil {
		// panic(err)
		// }


		// prod only
		cer, err := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
		if err != nil {
			log.Println(err)
			return
		}

		config := &tls.Config{Certificates: []tls.Certificate{cer}}
		ln, err := tls.Listen("tcp", fmt.Sprintf("%s:%s", httpHost, httpPort), config)
		if err != nil {
			log.Println(err)
			return
		}
		defer ln.Close()

		if err := http.Serve(ln, cors(mux)); err != nil {
			panic(err)
		}
	}()

	listener := createGrpcNetworkListener()
	log.Printf("Start serving grpc\n")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
