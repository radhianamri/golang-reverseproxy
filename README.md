# golang-reverseproxy
Reverse Proxy that supports layer 4 &amp; 7 load-balancing using Golang


Run the following command to create an SSL certificate to invoke HTTP2 protocol

openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key

Using the private key, generate a cerificate and fillout the common name baased on your corresponding domain:
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
