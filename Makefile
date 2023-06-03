start-server:
	go run ./cmd/...

# Case 1: Send 32 bytes with a limit-rate of 8 bytes per second
# We need 4 seconds to fully pass the request body
# Since we have a 2-second timeout (deadline), the request will fail
read:
	curl --limit-rate 8 -X POST localhost:4000/read -d '{"client_id":9999,"amount":1000}'

# Case 2: Data will be sent without any limit-rate
# However, we encounter a timeout due to the custom and slow reading process of the data
# Note: To reproduce the problem, we need to send a large chunk of data due to the internals of the Go server
custom_read:
	curl -X POST localhost:4000/custom_read -d @data.out

# Case 3: Sending a request to the server to receive a response
# However, an error occurs due to the "long work" that the server needs to perform
# This causes the server's internal write deadline to be exceeded when writing the response
write:
	curl -X GET localhost:4000/write

# Case 4: Timeout due to custom response writing process
# We have a large response data (refer to the code) to avoid buffering (more: https://github.com/golang/go/issues/21389)
custom_write:
	curl -X GET localhost:4000/custom_write

.PHONY: read custom_read write custom_write
