package cmd

import (
	"bufio"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create a metrics registry.
		reg := prometheus.NewRegistry()
		// Create some standard client metrics.
		grpcMetrics := grpc_prometheus.NewClientMetrics()
		// Register client metrics to registry.
		reg.MustRegister(grpcMetrics)
		// Create a insecure gRPC channel to communicate with the server.
		conn, err := grpc.Dial(
			fmt.Sprintf("localhost:%v", 9093),
			grpc.WithUnaryInterceptor(grpcMetrics.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(grpcMetrics.StreamClientInterceptor()),
			grpc.WithInsecure(),
		)
		if err != nil {
			log.Fatal(err)
		}

		defer conn.Close()

		// Create a HTTP server for prometheus.
		httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9094)}

		// Start your http server for prometheus.
		go func() {
			if err := httpServer.ListenAndServe(); err != nil {
				log.Fatal("Unable to start a http server.")
			}
		}()

		// Create a gRPC server client.
		client := pb.NewDemoServiceClient(conn)
		fmt.Println("Start to call the method called SayHello every 3 seconds")
		go func() {
			for {
				// Call “SayHello” method and wait for response from gRPC Server.
				_, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Test"})
				if err != nil {
					log.Printf("Calling the SayHello method unsuccessfully. ErrorInfo: %+v", err)
					log.Printf("You should to stop the process")
					return
				}
				time.Sleep(3 * time.Second)
			}
		}()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("You can press n or N to stop the process of client")
		for scanner.Scan() {
			if strings.ToLower(scanner.Text()) == "n" {
				os.Exit(0)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

}
