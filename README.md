find helix -type f -name "*" | while read file; do   echo "====== $file ======";   cat "$file";   echo -e "\n"; done


curl "http://localhost:8080/set?key=foo&value=bar"
curl "http://localhost:8080/get?key=foo"
curl "http://localhost:8080/del?key=foo"

go run ./helix/cmd/server  
go build -o helix-server ./helix/cmd/server
