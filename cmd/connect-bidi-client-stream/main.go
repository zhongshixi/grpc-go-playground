package main

type EventService struct{}

// func main() {
// 	client := serviceconnect.NewEventServiceClient(
// 		&http.Client{
// 			Transport: &http2.Transport{
// 				AllowHTTP: true,
// 				DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
// 					// If you're also using this client for non-h2c traffic, you may want to
// 					// delegate to tls.Dial if the network isn't TCP or the addr isn't in an
// 					// allowlist.
// 					return net.Dial(network, addr)
// 				},
// 			},
// 		},
// 		"http://localhost:7100", // in production example, it is https
// 		connect.WithGRPC(),
// 		// connect.WithSendCompression("gzip"),
// 	)

// 	startTime := time.Now()

// 	defer func() {
// 		log.Println("time elaspsed:", time.Since(startTime).Seconds(), "seconds")
// 	}()

// 	stream := client.HandleBidiStream(context.Background())

// 	count := int64(0)
// 	for {
// 		if count == data.Count {
// 			log.Println("stop")
// 			return
// 		}

// 		count = count + 1
// 		err := stream.Send(&service.EventRequest{
// 			OrderId: count,
// 			Data:    data.Bid,
// 		})

// 		if err != nil {
// 			log.Println("Err:", err)
// 			log.Println("Trailer:", stream.ResponseTrailer())

// 			if errors.Is(err, io.EOF) {
// 				if msg, err := stream.Receive(); err != nil {
// 					log.Println("Msg:", msg)
// 					log.Println("Receive Err:", err)
// 				}
// 			}
// 			return
// 		}

// 		// msg, err := stream.Receive()
// 		// if err != nil {
// 		// 	log.Println("Receive Err:", err)
// 		// }

// 		// log.Println("Resp:", msg)
// 	}
// }
