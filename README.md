# dist-kv-store

## Put
`curl -X POST -d '{"customer_id":"'"$(uuidgen)"'","line_items":[{"item_id":"'"$(uuidgen)"'","quantity":5,"price":1999}]}' localhost:3000/orders`

## Get
`curl localhost:3000/orders?cursor=2`

## ListAll
`curl localhost:3000/orders?cursor=1717`



## Insert Test Data
```bash
for i in {1..1000}; do   curl -X POST -d '{"customer_id":"'"$(uuidgen)"'","line_items":[{"item_id":"'"$(uuidgen)"'","quantity":5,"price":1999}]}' localhost:3000/orders; done  
```