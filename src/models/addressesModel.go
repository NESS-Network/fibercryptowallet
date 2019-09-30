package models

import (
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

const (
	Address    = int(core.Qt__UserRole) + 1
	ASky       = int(core.Qt__UserRole) + 2
	ACoinHours = int(core.Qt__UserRole) + 3
	AMarked    = int(core.Qt__UserRole) + 4
)

type AddressesModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*QAddress              `property:"addresses"`

	_ func(*QAddress)                        `slot:"addAddress"`
	_ func(int)                              `slot:"removeAddress"`
	_ func(int, string, uint64, uint64, int) `slot:"editAddress"`
	_ func([]*QAddress)                      `slot:"loadModel"`
	_ int                                    `property:"count"`
}

type QAddress struct {
	core.QObject

	_ string `property:"address"`
	_ string `property:"addressSky"`
	_ string `property:"addressCoinHours"`
	_ int    `property:"marked"`
}

func (m *AddressesModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		Address:    core.NewQByteArray2("address", -1),
		ASky:       core.NewQByteArray2("addressSky", -1),
		ACoinHours: core.NewQByteArray2("addressCoinHours", -1),
		AMarked:    core.NewQByteArray2("marked", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(m, qml.QQmlEngine__CppOwnership)
	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)

	m.ConnectAddAddress(m.addAddress)
	m.ConnectEditAddress(m.editAddress)
	m.ConnectRemoveAddress(m.removeAddress)
	m.ConnectLoadModel(m.loadModel)
	m.SetCount(0)

}

func (m *AddressesModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Addresses()) {
		return core.NewQVariant()
	}

	var a = m.Addresses()[index.Row()]

	switch role {

	case Address:
		{
			return core.NewQVariant1(a.Address())
		}
	case ASky:
		{
			return core.NewQVariant1(a.AddressSky())
		}
	case ACoinHours:
		{
			return core.NewQVariant1(a.AddressCoinHours())
		}
	case AMarked:
		{
			return core.NewQVariant1(a.Marked())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *AddressesModel) rowCount(parent *core.QModelIndex) int {
	return len(m.Addresses())
}

func (m *AddressesModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *AddressesModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *AddressesModel) addAddress(address *QAddress) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Addresses()), len(m.Addresses()))
	qml.QQmlEngine_SetObjectOwnership(address, qml.QQmlEngine__CppOwnership)
	m.SetAddresses(append(m.Addresses(), address))
	m.EndInsertRows()
	m.SetCount(m.Count() + 1)
}

func (m *AddressesModel) removeAddress(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetAddresses(append(m.Addresses()[:row], m.Addresses()[row+1:]...))
	m.EndRemoveRows()
	m.SetCount(m.Count() - 1)
}

func (m *AddressesModel) editAddress(row int, address string, sky, coinHours uint64, marked int) {
	a := m.Addresses()[row]
	a.SetAddress(address)
	//TODO: report possible error
	accuracy, _ := util.AltcoinQuotient("SKY")
	a.SetAddressSky(util.FormatCoins(sky, accuracy))
	//TODO: report possible error
	accuracy, _ = util.AltcoinQuotient("SKYCH")
	a.SetAddressCoinHours(util.FormatCoins(coinHours, accuracy))
	changeMarked := true
	if marked == a.Marked() {
		changeMarked = false
	}
	a.SetMarked(marked)
	pIndex := m.Index(row, 0, core.NewQModelIndex())
	if changeMarked {
		m.DataChanged(pIndex, pIndex, []int{Address, ASky, ACoinHours, AMarked})
	} else {
		m.DataChanged(pIndex, pIndex, []int{Address, ASky, ACoinHours})
	}

}

func (m *AddressesModel) loadModel(Qaddresses []*QAddress) {
	for _, addr := range Qaddresses {
		qml.QQmlEngine_SetObjectOwnership(addr, qml.QQmlEngine__CppOwnership)
	}
	addresses := make([]*QAddress, 0)
	address := NewQAddress(nil)
	address.SetAddress("--------------------------")
	address.SetAddressSky("0")
	address.SetAddressCoinHours("0")
	address.SetMarked(0)
	qml.QQmlEngine_SetObjectOwnership(address, qml.QQmlEngine__CppOwnership)
	addresses = append(addresses, address)
	addresses = append(addresses, Qaddresses...)

	m.BeginResetModel()
	m.SetAddresses(addresses)
	m.SetCount(len(addresses))

	m.EndResetModel()

}