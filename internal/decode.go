package internal

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

// DecodeOptions specifies options for decoding
type DecodeOptions struct {
	JSON   bool
	Output string
}

// Decode jwt
func Decode(token string, opt DecodeOptions) (string, error) {
	// Split and parse JWT
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errors.New("Invalid token: requires 3 parts")
	}
	// Parse header/payload
	header := make(map[string]interface{})
	payload := make(map[string]interface{})
	obj := make(map[string]interface{})
	err := decodePart([]byte(parts[0]), &header)
	if err != nil {
		return "", errors.Wrap(err, "Invalid header")
	}
	err = decodePart([]byte(parts[1]), &payload)
	if err != nil {
		return "", errors.Wrap(err, "Invalid payload")
	}
	obj["header"] = header
	obj["payload"] = payload

	// If json, prettyprint and return
	if opt.JSON {
		str, err := json.MarshalIndent(obj, "", "  ")
		return string(str) + "\n", err
	}

	// If output, parse template and execute
	if len(opt.Output) > 0 {
		t, err := template.New("").Parse(opt.Output)
		if err != nil {
			return "", errors.Wrap(err, "Invalid output")
		}
		var str bytes.Buffer
		err = t.Execute(&str, obj)
		if err != nil {
			return "", errors.Wrap(err, "Invalid output")
		}
		return str.String(), nil
	}

	// Default output
	var str strings.Builder
	str.WriteString("HEADER:\n")
	h, _ := json.MarshalIndent(header, "", "  ")
	str.Write(h)
	str.WriteString("\n\n")
	str.WriteString("PAYLOAD:\n")
	p, _ := json.MarshalIndent(payload, "", "  ")
	str.Write(p)
	str.WriteString("\n")
	return str.String(), nil
}

func decodePart(encoded []byte, obj interface{}) error {
	encoding := base64.RawURLEncoding
	decoded := make([]byte, encoding.DecodedLen(len(encoded)))
	_, err := encoding.Decode(decoded, encoded)
	if err != nil {
		return err
	}
	return json.Unmarshal(decoded, obj)
}
