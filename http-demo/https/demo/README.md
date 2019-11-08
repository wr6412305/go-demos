# https

curl --cert client.pem url

生成私钥
openssl genrsa -out server.key 2048

生成证书
openssl req -new -x509 -key server.key -out server.crt -days 3650

不能直接将 .crt 转成 .pem，需要经过 .der　中转
openssl x509 -in client.crt -out client.der -outform der
openssl x509 -in client.der -inform der -outform pem -out client.pem

openssl x509 -in ca.crt -out ca.der -outform der
openssl x509 -in ca.der -inform der -outform pem -out ca_info.pem

将.key 转成 .pem
openssl rsa -in client.key -out client.pem
