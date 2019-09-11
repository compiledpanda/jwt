package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/compiledpanda/jwt/internal"
	"github.com/spf13/cobra"
)

var encodeCty string
var encodeKid string
var encodeIss string
var encodeSub string
var encodeAud []string
var encodeExp string
var encodeNbf string
var encodeIat string
var encodeJti string
var encodeClaims []string
var encodePayload string
var encodeAlgorithm string
var encodeSecret string

var encode = &cobra.Command{
	Use:   "encode",
	Short: "Encode a JWT",
	Run: func(cmd *cobra.Command, args []string) {
		options := internal.EncodeOptions{
			Header:  make(map[string]string),
			Payload: make(map[string]interface{}),
		}

		// Header
		if encodeCty != "" {
			options.Header["cty"] = encodeCty
		}
		if encodeKid != "" {
			options.Header["kid"] = encodeKid
		}

		// Payload
		if len(encodePayload) > 0 {
			var payload []byte
			if encodePayload == "-" {
				payload = readStdIn()
			} else if strings.HasPrefix(encodePayload, "@") {
				payload = readFile(encodePayload[1:])
			} else {
				payload = []byte(encodePayload)
			}
			err := json.Unmarshal(payload, &options.Payload)
			if err != nil {
				fmt.Println("Could not unmarshal json", err)
				os.Exit(1)
			}
		} else {
			if encodeIss != "" {
				options.Payload["iss"] = encodeIss
			}
			if encodeSub != "" {
				options.Payload["sub"] = encodeSub
			}
			if len(encodeAud) > 0 {
				options.Payload["aud"] = encodeAud
			}
			if encodeExp != "" {
				// TODO validate
				options.Payload["exp"] = encodeExp
			}
			if encodeNbf != "" {
				// TODO validate
				options.Payload["nbf"] = encodeNbf
			}
			if encodeIat != "" {
				// TODO validate
				options.Payload["iat"] = encodeIat
			}
			if encodeJti != "" {
				options.Payload["jti"] = encodeJti
			}
			for _, claim := range encodeClaims {
				parts := strings.Split(claim, "=")
				if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
					fmt.Println("claim must be in the format of a=b")
					os.Exit(1)
				}
				var value []byte
				if parts[1] == "-" {
					value = readStdIn()
				} else if strings.HasPrefix(parts[1], "@") {
					value = readFile(parts[1][1:])
				} else {
					value = []byte(parts[1])
				}
				fmt.Println(string(value))
				var obj interface{}
				err := json.Unmarshal(value, &obj)
				// If we can't parse as json, set as string
				if err != nil {
					fmt.Println(string(value))
					options.Payload[parts[0]] = string(value)
				} else {
					options.Payload[parts[0]] = obj
				}
			}
		}

		// Algorithm
		if !internal.ValidAlgorithm(encodeAlgorithm) {
			fmt.Println("algorithm is invalid. Must be one of ", strings.Join(internal.Algorithms, ", "))
			os.Exit(1)
		}
		options.Algorithm = encodeAlgorithm

		// Secret
		if encodeSecret == "-" {
			options.Secret = readStdIn()
		} else if strings.HasPrefix(encodeSecret, "@") {
			options.Secret = readFile(encodeSecret[1:])
		} else {
			options.Secret = []byte(encodeSecret)
		}

		// Encode and print
		str, err := internal.Encode(options)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(str)
	},
}

func readStdIn() []byte {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Could not read from stdin", err)
		os.Exit(1)
	}
	return b
}

func readFile(file string) []byte {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Could not open file", file, err)
		os.Exit(1)
	}
	return b
}

func init() {
	encode.Flags().StringVar(&encodeCty, "cty", "", "Set the content type in the header")
	encode.Flags().StringVar(&encodeKid, "kid", "", "Set the key id in the header")
	encode.Flags().StringVar(&encodeIss, "iss", "", "Issuer claim")
	encode.Flags().StringVar(&encodeSub, "sub", "", "Subject claim")
	encode.Flags().StringSliceVar(&encodeAud, "aud", nil, "Audience claim")
	encode.Flags().StringVar(&encodeExp, "exp", "", "Expiration Time claim")
	encode.Flags().StringVar(&encodeNbf, "nbf", "", "Not Before claim")
	encode.Flags().StringVar(&encodeIat, "iat", "", "Issued At claim")
	encode.Flags().StringVar(&encodeJti, "jti", "", "JWT ID claim")
	encode.Flags().StringSliceVarP(&encodeClaims, "claim", "c", nil, "Claim key/value pairs (`a=b` string, `a=-` string from stdin, `a=@file.json` string from file). Will try to parse string as json, and use string as fallback")
	encode.Flags().StringVarP(&encodePayload, "payload", "p", "", "The entire payload body in json format (string, @file, or '-' to read from stdin)")
	encode.Flags().StringVarP(&encodeAlgorithm, "algorithm", "a", "", "(Required) The algorithm to use for signing. Must be one of "+strings.Join(internal.Algorithms, ", "))
	encode.MarkFlagRequired("algorithm")
	encode.Flags().StringVarP(&encodeSecret, "secret", "s", "", "(Required) The secret (string, @file, or '-' to read from stdin)")
	encode.MarkFlagRequired("secret")
	root.AddCommand(encode)
}
