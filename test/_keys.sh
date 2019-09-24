#!/usr/bin/env bash

# hmac jwt
go run ../jwt.go encode -a HS256 -s secret > hmac_hs256.jwt
go run ../jwt.go encode -a HS384 -s secret > hmac_hs384.jwt
go run ../jwt.go encode -a HS512 -s secret > hmac_hs512.jwt

# rsa
openssl genrsa -out rsa_pkcs1_private.pem 2048
openssl rsa -in rsa_pkcs1_private.pem -outform DER -out rsa_pkcs1_private.der
openssl rsa -in rsa_pkcs1_private.pem -outform PEM -RSAPublicKey_out -out rsa_pkcs1_public.pem
openssl rsa -in rsa_pkcs1_private.pem -outform DER -RSAPublicKey_out -out rsa_pkcs1_public.der
openssl pkcs8 -topk8 -in rsa_pkcs1_private.pem -nocrypt -outform PEM -out rsa_pkcs8_private.pem
openssl pkcs8 -topk8 -in rsa_pkcs1_private.pem -nocrypt -outform DER -out rsa_pkcs8_private.der
openssl rsa -in rsa_pkcs1_private.pem -outform PEM -pubout -out rsa_x509_public.pem
openssl rsa -in rsa_pkcs1_private.pem -outform DER -pubout -out rsa_x509_public.der

# rsa jwt
go run ../jwt.go encode -a RS256 -s @./rsa_pkcs1_private.pem > rsa_rs256.jwt
go run ../jwt.go encode -a RS384 -s @./rsa_pkcs1_private.pem > rsa_rs384.jwt
go run ../jwt.go encode -a RS512 -s @./rsa_pkcs1_private.pem > rsa_rs512.jwt
go run ../jwt.go encode -a PS256 -s @./rsa_pkcs1_private.pem > rsa_ps256.jwt
go run ../jwt.go encode -a PS384 -s @./rsa_pkcs1_private.pem > rsa_ps384.jwt
go run ../jwt.go encode -a PS512 -s @./rsa_pkcs1_private.pem > rsa_ps512.jwt

# ecdsa
openssl ecparam -name secp521r1 -genkey -noout -outform PEM -out ecdsa_ec_private.pem
openssl ec -in ecdsa_ec_private.pem -outform PEM -out ecdsa_ec_private.der
openssl pkcs8 -topk8 -in ecdsa_ec_private.pem -nocrypt -outform PEM -out ecdsa_pkcs8_private.pem
openssl pkcs8 -topk8 -in ecdsa_ec_private.pem -nocrypt -outform DER -out ecdsa_pkcs8_private.der
openssl ec -in ecdsa_ec_private.pem -outform PEM -pubout -out ecdsa_x509_public.pem
openssl ec -in ecdsa_ec_private.pem -outform DER -pubout -out ecdsa_x509_public.der

# ecdsa jwt
go run ../jwt.go encode -a ES256 -s @./ecdsa_ec_private.pem > ecdsa_es256.jwt
go run ../jwt.go encode -a ES384 -s @./ecdsa_ec_private.pem > ecdsa_es384.jwt
go run ../jwt.go encode -a ES512 -s @./ecdsa_ec_private.pem > ecdsa_es512.jwt

# openssh
ssh-keygen -t rsa -b 2048 -N "" -C "" -f ./rsa_openssh_private
mv ./rsa_openssh_private ./rsa_openssh_private.pem
mv ./rsa_openssh_private.pub ./rsa_openssh_public.pub
ssh-keygen -t ecdsa -b 521 -N "" -C "" -f ./ecdsa_openssh_private
mv ./ecdsa_openssh_private ./ecdsa_openssh_private.pem
mv ./ecdsa_openssh_private.pub ./ecdsa_openssh_public.pub

# openssh jwt
go run ../jwt.go encode -a RS256 -s @./rsa_openssh_private.pem > rsa_openssh_rs256.jwt