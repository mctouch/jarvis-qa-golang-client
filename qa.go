package main

import (
    "context"
    "flag"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "jarvis-qa/jarvis_speech"
)

func main() {
    serverFlag := flag.String("server", "localhost:50051", "Jarvis server to connect to")
    queryFlag := flag.String("query", "", "Question to ask")
    contextFlag := flag.String("context", "", "Context to search for answer")

    flag.Parse()

    if *queryFlag == "" || *contextFlag == "" {
        log.Fatal("-query and -context must both be set.")
    }

    // Set up a connection to the server.
    conn, err := grpc.Dial(*serverFlag, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
    if err != nil {
        log.Fatalf("Unable to connect to server: %v", err)
    }
    defer conn.Close()
    c := pb.NewJarvisNLPClient(conn)

    // Contact the server and print out its response.
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.NaturalQuery(ctx, &pb.NaturalQueryRequest{
        Query:   *queryFlag,
        Context: *contextFlag,
        TopN:    1,
    })

    if err != nil {
        log.Fatalf("Issue encountered with Jarvis Request: %v", err)
    }

    log.Printf("Question: %s", *queryFlag)
    log.Printf("Answer:   %s", r.GetResults()[0].GetAnswer())
}
