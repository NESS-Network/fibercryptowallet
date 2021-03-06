package data

import (
	"encoding/json"
	"github.com/fibercrypto/fibercryptowallet/src/core"
	"github.com/fibercrypto/fibercryptowallet/src/util"
	"strings"
)

// Contact is a contact of the addrsBook
type Contact struct {
	ID      uint64
	Address []Address
	Name    []byte
}

// Address is the relation of an address and his coin type.
type Address struct {
	Value []byte
	Coin  []byte
}

// MarshalBinary encodes a user to binary format.
func (c *Contact) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalBinary decodes a user from binary data.
func (c *Contact) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &c)
}

// GetID get ID of current contact.
func (c *Contact) GetID() uint64 {
	return c.ID
}

// SetID set an ID to current contact.
func (c *Contact) SetID(id uint64) {
	c.ID = id
}

// GetAddresses get address list of current contact.
func (c *Contact) GetAddresses() []core.StringAddress {
	var addresses []core.StringAddress
	for e := range c.Address {
		addresses = append(addresses, &c.Address[e])
	}
	return addresses
}

// SetAddresses set an address list to the current contact.
func (c *Contact) SetAddresses(addrs []core.StringAddress) {
	for e := range addrs {
		if v, ok := addrs[e].(*Address); ok {
			c.Address = append(c.Address, *v)
		} else {
			panic("Error in SetAddress: addrs cannot parse to type data.Address")

		}
	}
}

// GetName return contact name
func (c *Contact) GetName() string {
	return string(c.Name)
}

// SetName set a name to the contact
func (c *Contact) SetName(newName string) {
	c.Name = []byte(newName)
}

func (c *Contact) IsValid() bool {
	if strings.ReplaceAll(c.GetName(), " ", "") == "" {
		return false
	}

	for e := range c.GetAddresses() {
		if !c.GetAddresses()[e].IsValid() {
			return false
		}
	}
	return true
}

// .....Address

// GetValue get address string.
func (ad *Address) GetValue() []byte {
	return ad.Value
}

// SetValue set an address string.
func (ad *Address) SetValue(val []byte) {
	ad.Value = val
}

// GetCoinType get coin type of an address.
func (ad *Address) GetCoinType() []byte {
	return ad.Coin
}

// SetCoinType set the coin type to current address.
func (ad *Address) SetCoinType(coinType []byte) {
	ad.Coin = coinType
}

func (ad *Address) IsValid() bool {
	if _, err := util.AddressFromString(string(ad.GetValue()),
		string(ad.GetCoinType())); err != nil {
		return false
	}
	return true
}
