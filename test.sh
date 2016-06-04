curl -X POST 127.0.0.1:8000/queues/work
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 1"
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 2"
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 3"
curl 127.0.0.1:8000/stats
curl 127.0.0.1:8000/queues/work/messages
curl 127.0.0.1:8000/queues/work/messages
curl 127.0.0.1:8000/queues/work/messages
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 1"
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 2"
curl -X POST 127.0.0.1:8000/queues/work/messages -d "message 3"
curl -X POST 127.0.0.1:8000/queues/work/drain
curl 127.0.0.1:8000/stats
curl -X DELETE 127.0.0.1:8000/queues/work
