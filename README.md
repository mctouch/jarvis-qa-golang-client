# jarvis-qa-golang-client
A Question and Answer Client written in Go using gRPC protocol to a Nvidia Jarvis Triton service

go mod tidy

./qa -server <your triton container hostmame>:50051 -query "Which founder of NVIDIA previously worked at Sun?"      -context "Nvidia was founded on April 5, 1993, by Jensen Huang (CEO as of
          2020), a Taiwanese American, previously director of CoreWare at LSI
          Logic and a microprocessor designer at Advanced Micro Devices (AMD),
          Chris Malachowsky, an electrical engineer who worked at Sun
          Microsystems, and Curtis Priem, previously a senior staff engineer
          and graphics chip designer at Sun Microsystems."
