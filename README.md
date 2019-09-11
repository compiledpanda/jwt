# jwt
CLI to encode, decode, and validate JWTs

Uses [github.com/gbrlsnchs/jwt/v3](https://github.com/gbrlsnchs/jwt) and inspired by [jwt-cli](https://github.com/mike-engel/jwt-cli).

## Examples
```bash
$> jwt encode --iss Me -c custom=payload -a RS512 -s @/path/to/private/key
eyJh...sw5c
$> cat ./path/to/jwt | jwt decode - -o "{{.payload.custom}}"
payload
$> jwt validate eyJh...sw5c -a RS512 --iss Me -s @/path/to/public/key
VALID
```

## Docs

### Encode
Create and sign a jwt.

`jwt encode [options]`

#### Header Options
* `--cty` - Set the content type in the header
* `--kid` - Set the key id in the header
#### Payload Options
* `--iss` - Issuer claim
* `--sub` - Subject claim
* `--aud` - Audience claim
* `--exp` - Expiration Time claim
* `--nbf` - Not Before claim
* `--iat` - Issued At claim
* `--jti` - JWT ID claim
* `-c, --claim` - Claim key/value pairs (`a=b` string, `a=-` string from stdin, `a=@file.json` string from file). Will try to parse string as json, and use string as fallback
* `-p, --payload` - The entire payload body in json format (`string`, `@file`, or `-` to read from stdin)
#### Signature Options
* `-a, --algorithm` - (Required) The algorithm to use for signing
    Possible Values: `HS256`, `HS384`,`HS512`,`RS256`,`RS384`,`RS512`,`ES256`,`ES384`,`ES512`,`PS256`,`PS384`,`PS512`,`EdDSA`
* `-s, --secret` - (Required) The secret (`string`, `@file`, or `-` to read from stdin)

### Decode
Decode jwt (string or `-` to read from stdin) and Prettyprint.

`jwt decode [options] <jwt>`

#### Options
`--json` - Output as json
    ```
    {
      "header": {...},
      "payload": {...}
    }
    ```
`-o, --output` - Go template string to format the output

### Validate
Validate the jwt (string or `-` to read from stdin). Will return an error code if JWT is invalid or fails a validation step

`jwt validate [options] <jwt>`

#### Options
* `--iss` - Fails if Issuer claim does not match
* `--sub` - Fails if Subject claim does not match
* `--aud` - Fails if Audience claim does not match
* `--exp` - Fails if Expiration Time claim is before this value
* `--nbf` - Fails of Not Before claim is after this value
* `--iat` - Fails if Issued At claim is after this value
* `--jti` - Fails if JWT ID claim does not match
* `-a, --algorithm` - The algorithm to validate against. Fails on mismatch
* `-s, --secret` - The secret (`string`, `@file`, or `-` to read from stdin). Fails if signature is invalid

## Secrets
Keys can be in PKCS1 or PKCS8 format, DER or PEM encoded.