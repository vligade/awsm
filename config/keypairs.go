package config

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go/service/simpledb"
)

// KeyPairClasses is a map of Image classes
type KeyPairClasses map[string]KeyPairClass

// KeyPairClass is a single Image class
type KeyPairClass struct {
	Description string `json:"description" awsmClass:"Description"`
	PublicKey   string `json:"publicKey" awsmClass:"Public Key"`
	PrivateKey1 string `json:"-"`
	PrivateKey2 string `json:"-"`
	PrivateKey3 string `json:"-"`
	PrivateKey4 string `json:"-"`
	PrivateKey  string `json:"privateKey" awsm:"ignore"`
}

// DefaultKeyPairClasses returns the default KeyPair classes
func DefaultKeyPairClasses() KeyPairClasses {
	defaultKeyPairs := make(KeyPairClasses)

	defaultKeyPairs["awsm"] = KeyPairClass{
		Description: "Default awsm Key Pair",
	}

	return defaultKeyPairs
}

// SaveKeyPairClass unmarshals a byte slice and inserts it into the db
func SaveKeyPairClass(className string, data []byte) (class KeyPairClass, err error) {
	err = json.Unmarshal(data, &class)
	if err != nil {
		return
	}

	// Generate the keys if needed.
	/*if class.PrivateKey == "" && class.PublicKey == "" {
		var publicKey, privateKey string
		publicKey, privateKey, err = GenerateKeyPair()
		if err != nil {
			// terminal.ErrorLine("Error while generating keypair class: " + className)
			return
		}

		privateKeyLen := len(privateKey) / 4
		class.PublicKey = publicKey
		class.PrivateKey1 = privateKey[:privateKeyLen]
		class.PrivateKey2 = privateKey[privateKeyLen : privateKeyLen*2]
		class.PrivateKey3 = privateKey[privateKeyLen*2 : privateKeyLen*3]
		class.PrivateKey4 = privateKey[privateKeyLen*3:]
	} else {*/
	privateKey := class.PrivateKey
	privateKeyLen := len(privateKey) / 4

	class.PrivateKey1 = privateKey[:privateKeyLen]
	class.PrivateKey2 = privateKey[privateKeyLen : privateKeyLen*2]
	class.PrivateKey3 = privateKey[privateKeyLen*2 : privateKeyLen*3]
	class.PrivateKey4 = privateKey[privateKeyLen*3:]
	/*}*/

	err = Insert("keypairs", KeyPairClasses{className: class})

	if err != nil {
		println(err)
	}
	return
}

// LoadKeyPairClass returns a single KeyPair class by its name
func LoadKeyPairClass(name string) (KeyPairClass, error) {
	cfgs := make(KeyPairClasses)
	item, err := GetItemByName("keypairs", name)
	if err != nil {
		return cfgs[name], err
	}
	cfgs.Marshal([]*simpledb.Item{item})
	return cfgs[name], nil
}

// LoadAllKeyPairClasses returns all Image classes
func LoadAllKeyPairClasses() (KeyPairClasses, error) {
	cfgs := make(KeyPairClasses)
	items, err := GetItemsByType("keypairs")
	if err != nil {
		return cfgs, err
	}

	cfgs.Marshal(items)
	return cfgs, nil
}

// Marshal puts items from SimpleDB into Image Classes
func (c KeyPairClasses) Marshal(items []*simpledb.Item) {
	for _, item := range items {
		name := strings.Replace(*item.Name, "keypairs/", "", -1)
		cfg := new(KeyPairClass)

		for _, attribute := range item.Attributes {

			val := *attribute.Value

			switch *attribute.Name {

			case "Description":
				cfg.Description = val

			case "PublicKey":
				cfg.PublicKey = val

			case "PrivateKey1":
				cfg.PrivateKey1 = val

			case "PrivateKey2":
				cfg.PrivateKey2 = val

			case "PrivateKey3":
				cfg.PrivateKey3 = val

			case "PrivateKey4":
				cfg.PrivateKey4 = val

			}
		}
		c[name] = *cfg
	}
}
