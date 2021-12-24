# Calculate hashes.
curl http://127.0.0.1:8080/send \
  --request POST \
  --header "Content-Type:application/json" \
  --data '["1","2"]' \
  --verbose

# Check hashes.
curl http://127.0.0.1:8080/check?ids=1,2,3,10 \
  --request GET \
  --verbose