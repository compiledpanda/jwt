#!/usr/bin/env bash

# rsa
openssl genrsa -out rsa_pkcs1_private.pem 2048
openssl rsa -in rsa_pkcs1_private.pem -outform DER -out rsa_pkcs1_private.der
openssl rsa -in rsa_pkcs1_private.pem -outform PEM -RSAPublicKey_out -out rsa_pkcs1_public.pem
openssl rsa -in rsa_pkcs1_private.pem -outform DER -RSAPublicKey_out -out rsa_pkcs1_public.der
openssl pkcs8 -topk8 -in rsa_pkcs1_private.pem -nocrypt -outform PEM -out rsa_pkcs8_private.pem
openssl pkcs8 -topk8 -in rsa_pkcs1_private.pem -nocrypt -outform DER -out rsa_pkcs8_private.der
openssl rsa -in rsa_pkcs1_private.pem -outform PEM -pubout -out rsa_x509_public.pem
openssl rsa -in rsa_pkcs1_private.pem -outform DER -pubout -out rsa_x509_public.der

# ecdsa
openssl ecparam -name secp521r1 -genkey -noout -outform PEM -out ecdsa_ec_private.pem
openssl ec -in ecdsa_ec_private.pem -outform PEM -out ecdsa_ec_private.der
openssl pkcs8 -topk8 -in ecdsa_ec_private.pem -nocrypt -outform PEM -out ecdsa_pkcs8_private.pem
openssl pkcs8 -topk8 -in ecdsa_ec_private.pem -nocrypt -outform DER -out ecdsa_pkcs8_private.der
openssl ec -in ecdsa_ec_private.pem -outform PEM -pubout -out ecdsa_x509_public.pem
openssl ec -in ecdsa_ec_private.pem -outform DER -pubout -out ecdsa_x509_public.der

# openssh
ssh-keygen -t rsa -b 2048 -N "" -C "" -f ./rsa_openssh_private
mv ./rsa_openssh_private ./rsa_openssh_private.pem
mv ./rsa_openssh_private.pub ./rsa_openssh_public.pub
ssh-keygen -t ecdsa -b 521 -N "" -C "" -f ./ecdsa_openssh_private
mv ./ecdsa_openssh_private ./ecdsa_openssh_private.pem
mv ./ecdsa_openssh_private.pub ./ecdsa_openssh_public.pub