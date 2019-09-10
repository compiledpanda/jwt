package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var encodeCty string
var encodeKid string
var encodeIss string
var encodeSub string
var encodeAud string
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
		fmt.Println("Encoding...")
	},
}

func init() {
	encode.Flags().StringVar(&encodeCty, "cty", "", "Set the content type in the header")
	encode.Flags().StringVar(&encodeKid, "kid", "", "Set the key id in the header")
	encode.Flags().StringVar(&encodeIss, "iss", "", "Issuer claim")
	encode.Flags().StringVar(&encodeSub, "sub", "", "Subject claim")
	encode.Flags().StringVar(&encodeAud, "aud", "", "Audience claim")
	encode.Flags().StringVar(&encodeExp, "exp", "", "Expiration Time claim")
	encode.Flags().StringVar(&encodeNbf, "nbf", "", "Not Before claim")
	encode.Flags().StringVar(&encodeIat, "iat", "", "Issued At claim")
	encode.Flags().StringVar(&encodeJti, "jti", "", "JWT ID claim")
	encode.Flags().StringSliceVarP(&encodeClaims, "claim", "c", nil, "custom claim key/value pairs (`a=b` string, `a='true'` raw json, `a=@file.txt` text from file, `a=@file.json` json from file)")
	encode.Flags().StringVarP(&encodePayload, "payload", "p", "", "The entire payload body in json format (string, @file, or '-' to read from stdin)")
	encode.Flags().StringVarP(&encodeAlgorithm, "algorithm", "a", "", "The algorithm to use for signing")
	encode.Flags().StringVarP(&encodeSecret, "secret", "s", "", "The secret (string, @file, or '-' to read from stdin)")
	root.AddCommand(encode)
}
