# Create a new work queue
curl 127.0.0.1:8000/queue/create -d '{"name": "work"}'

# Add messages to the work queue
curl 127.0.0.1:8000/message/put -d '{"queue": "work", "body": "message 1"}'
curl 127.0.0.1:8000/message/put -d '{"queue": "work", "body": "message 2"}'
curl 127.0.0.1:8000/message/put -d '{"queue": "work", "body": "message 3"}'

# Get the broker stats
curl 127.0.0.1:8000/broker/stats

# Read messages from the work queue.
curl 127.0.0.1:8000/message/get -d '{"queue": "work"}'
curl 127.0.0.1:8000/message/get -d '{"queue": "work"}'
curl 127.0.0.1:8000/message/get -d '{"queue": "work"}'

# Add messages to the work queue and drain them.
curl 127.0.0.1:8000/message/put -d '{"queue": "work", "body": "message 1"}'
curl 127.0.0.1:8000/message/put -d '{"queue": "work", "body": "message 2"}'
curl 127.0.0.1:8000/message/put -d '{"queue": "work", "body": "message 3"}'
curl 127.0.0.1:8000/queue/drain -d '{"name": "work"}'
curl 127.0.0.1:8000/broker/stats

# Delete the work queue.
curl 127.0.0.1:8000/queue/delete -d '{"name": "work"}'
curl 127.0.0.1:8000/broker/stats
