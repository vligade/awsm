package config

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/service/simpledb"
)

// SnapshotClasses is a map of Snapshot Classes
type SnapshotClasses map[string]SnapshotClass

// SnapshotClass is a single Snapshot Class
type SnapshotClass struct {
	Rotate           bool     `json:"rotate" awsmClass:"Rotate"`
	Retain           int      `json:"retain" awsmClass:"Retain"`
	Propagate        bool     `json:"propagate" awsmClass:"Propagate"`
	PropagateRegions []string `json:"propagateRegions" awsmClass:"Propagate Regions"`
	Description      string   `json:"description" awsmClass:"Description"`
	Volume           string   `json:"volume" awsmClass:"Volume"`
}

// DefaultSnapshotClasses returns the default Snapshot Classes
func DefaultSnapshotClasses() SnapshotClasses {
	defaultSnapshots := make(SnapshotClasses)

	defaultSnapshots["crusher"] = SnapshotClass{
		Description:      "Crusher configuration management",
		Propagate:        true,
		Retain:           5,
		Rotate:           true,
		PropagateRegions: []string{"us-west-2", "us-east-1", "eu-west-1"},
		Volume:           "",
	}

	defaultSnapshots["git"] = SnapshotClass{
		Description:      "Git repositories",
		Propagate:        true,
		Retain:           5,
		Rotate:           true,
		PropagateRegions: []string{"us-west-2", "us-east-1", "eu-west-1"},
		Volume:           "",
	}

	defaultSnapshots["mysql-data"] = SnapshotClass{
		Description:      "MySQL data folder",
		Propagate:        true,
		Retain:           5,
		Rotate:           true,
		PropagateRegions: []string{"us-west-2", "us-east-1", "eu-west-1"},
		Volume:           "",
	}

	return defaultSnapshots
}

// SaveSnapshotClass reads unmarshals a byte slice and inserts it into the db
func SaveSnapshotClass(className string, data []byte) (class SnapshotClass, err error) {
	err = json.Unmarshal(data, &class)
	if err != nil {
		return
	}

	err = Insert("snapshots", SnapshotClasses{className: class})
	return
}

// LoadSnapshotClass loads a Snapshot Class by its name
func LoadSnapshotClass(name string) (SnapshotClass, error) {
	cfgs := make(SnapshotClasses)
	item, err := GetItemByName("snapshots", name)
	if err != nil {
		return cfgs[name], err
	}

	cfgs.Marshal([]*simpledb.Item{item})
	return cfgs[name], nil
}

// LoadAllSnapshotClasses loads all Snapshot Classes
func LoadAllSnapshotClasses() (SnapshotClasses, error) {
	cfgs := make(SnapshotClasses)
	items, err := GetItemsByType("snapshots")
	if err != nil {
		return cfgs, err
	}

	cfgs.Marshal(items)
	return cfgs, nil
}

// Marshal puts items from SimpleDB into a Snapshot Class
func (c SnapshotClasses) Marshal(items []*simpledb.Item) {
	for _, item := range items {
		name := strings.Replace(*item.Name, "snapshots/", "", -1)
		cfg := new(SnapshotClass)
		for _, attribute := range item.Attributes {

			val := *attribute.Value

			switch *attribute.Name {

			case "Description":
				cfg.Description = val

			case "Propagate":
				cfg.Propagate, _ = strconv.ParseBool(val)

			case "PropagateRegions":
				cfg.PropagateRegions = append(cfg.PropagateRegions, val)

			case "Retain":
				cfg.Retain, _ = strconv.Atoi(val)

			case "Rotate":
				cfg.Rotate, _ = strconv.ParseBool(val)

			case "Volume":
				cfg.Volume = val

			}
		}
		c[name] = *cfg
	}
}
