```sh
curl \
    -X GET \
    --location \
    'http://127.0.0.1:8086/getBook?id=1'





curl \
    -H "Content-Type: application/json" \
    -X POST \
    --data-raw '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
    --location \
    8.122.177.188:8000
```