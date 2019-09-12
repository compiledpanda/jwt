#!/usr/bin/env bash

# openssl rsa -in rsa_pkcs1_private.pem -outform PEM -pubout -out rsa_x509_public.pem

# ecdsa
openssl ecparam -name secp521r1 -genkey -noout -outform PEM -out ecdsa_ec_private.pem
openssl ec -in ecdsa_ec_private.pem -outform PEM -out ecdsa_ec_private.der
openssl pkcs8 -topk8 -in ecdsa_ec_private.pem -nocrypt -outform PEM -out ecdsa_pkcs8_private.pem
openssl pkcs8 -topk8 -in ecdsa_ec_private.pem -nocrypt -outform DER -out ecdsa_pkcs8_private.der
openssl ec -in ecdsa_ec_private.pem -outform PEM -pubout -out ecdsa_x509_public.pem
openssl ec -in ecdsa_ec_private.pem -outform DER -pubout -out ecdsa_x509_public.der