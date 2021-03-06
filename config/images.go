package config

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/simpledb"
)

// ImageClasses is a map of Image classes
type ImageClasses map[string]ImageClass

// ImageClass is a single Image class
type ImageClass struct {
	Instance         string   `json:"instance" awsmClass:"Instance"`
	Rotate           bool     `json:"rotate" awsmClass:"Rotate"`
	Retain           int      `json:"retain" awsmClass:"Retain"`
	Propagate        bool     `json:"propagate" awsmClass:"Propagate"`
	PropagateRegions []string `json:"propagateRegions" awsmClass:"Propagate Regions"`
	Version          int      `json:"version" awsmClass:"Version"`
}

// DefaultImageClasses returns the default Image classes
func DefaultImageClasses() ImageClasses {
	defaultImages := make(ImageClasses)

	defaultImages["awsm-init"] = ImageClass{
		Version:          0,
		Rotate:           true,
		Retain:           5,
		Propagate:        true,
		PropagateRegions: []string{"us-west-2", "us-east-1", "eu-west-1"},
		Instance:         "",
	}

	return defaultImages
}

// SaveImageClass reads unmarshals a byte slice and inserts it into the db
func SaveImageClass(className string, data []byte) (class ImageClass, err error) {
	err = json.Unmarshal(data, &class)
	if err != nil {
		return
	}

	err = Insert("images", ImageClasses{className: class})
	return
}

// LoadImageClass returns a single Image class by its name
func LoadImageClass(name string) (ImageClass, error) {
	cfgs := make(ImageClasses)
	item, err := GetItemByName("images", name)
	if err != nil {
		return cfgs[name], err
	}
	cfgs.Marshal([]*simpledb.Item{item})
	return cfgs[name], nil
}

// LoadAllImageClasses returns all Image classes
func LoadAllImageClasses() (ImageClasses, error) {
	cfgs := make(ImageClasses)
	items, err := GetItemsByType("images")
	if err != nil {
		return cfgs, err
	}

	cfgs.Marshal(items)
	return cfgs, nil
}

// Marshal puts items from SimpleDB into Image Classes
func (c ImageClasses) Marshal(items []*simpledb.Item) {
	for _, item := range items {
		name := strings.Replace(*item.Name, "images/", "", -1)
		cfg := new(ImageClass)

		for _, attribute := range item.Attributes {

			val := *attribute.Value

			switch *attribute.Name {

			case "Version":
				cfg.Version, _ = strconv.Atoi(val)

			case "Propagate":
				cfg.Propagate, _ = strconv.ParseBool(val)

			case "PropagateRegions":
				cfg.PropagateRegions = append(cfg.PropagateRegions, val)

			case "Rotate":
				cfg.Rotate, _ = strconv.ParseBool(val)

			case "Retain":
				cfg.Retain, _ = strconv.Atoi(val)

			case "Instance":
				cfg.Instance = val

			}
		}
		c[name] = *cfg
	}
}

// SetInstance updates the source instance of an Image
func (c *ImageClass) SetInstance(name string, instance string) error {
	c.Instance = instance

	updateCfgs := make(ImageClasses)
	updateCfgs[name] = *c

	return Insert("images", updateCfgs)
}

// SetVersion updates the version of an Image
func (c *ImageClass) SetVersion(name string, version int) error {
	c.Version = version

	updateCfgs := make(ImageClasses)
	updateCfgs[name] = *c

	return Insert("images", updateCfgs)
}

// Increment increments the version of an Image
func (c *ImageClass) Increment(name string) error {
	c.Version++
	return c.SetVersion(name, c.Version)
}

// Decrement decrements the version of an Image
func (c *ImageClass) Decrement(name string) error {
	c.Version--
	return c.SetVersion(name, c.Version)
}
