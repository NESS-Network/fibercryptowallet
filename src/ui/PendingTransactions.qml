import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import PendingModels 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    property bool showOnlyMine: false

    GroupBox {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        label: CheckDelegate {
            text: qsTr("Show only mine")
            checked: showOnlyMine
            onCheckedChanged: {
                showOnlyMine = checked
                if(checked){
                    modelPendingTransactions.getMine()
                }
                else{
                    modelPendingTransactions.getAll()
                }
            }
        }

        ColumnLayout {
            id: columnLayoutFrame
            anchors.fill: parent

            ColumnLayout {
                id: columnLayoutHeader

                RowLayout {
                    Layout.topMargin: 30

                    Label {
                        text: qsTr("Transaction ID")
                        font.pointSize: 9
                        Layout.leftMargin: 10 // Empirically obtained
                        Layout.fillWidth: true
                    }
                    Label {
                        text: qsTr("Sky")
                        font.pointSize: 9
                        Layout.rightMargin: 92
                    }
                    Label {
                        text: qsTr("Coin hours")
                        font.pointSize: 9
                        Layout.rightMargin: 33
                    }
                    Label {
                        text: qsTr("Timestamp")
                        font.pointSize: 9
                        Layout.rightMargin: 98
                    }
                } // RowLayout

                Rectangle {
                    Layout.fillWidth: true
                    height: 1
                    color: "#DDDDDD"
                }
            } // ColumnLayout (header)

            ScrollView {
                id: scrollItem
                Layout.fillWidth: true
                Layout.fillHeight: true
                ScrollBar.horizontal.policy: ScrollBar.AlwaysOff

                ListView {
                    id: listPendingTransactions
                    anchors.fill: parent
                    clip: true

                    model: modelPendingTransactions.transactions
                    delegate: PendingTransactionsDelegate {
                        property bool hide: false

                        width: parent.width

                        modelMine: modelData.mine
                        modelTransactionID: modelData.transactionID
                        modelSky: modelData.sky
                        modelCoinHours: modelData.coinHours
                        modelTimestamp: modelData.timeStamp

                        height: hide ? 0 : implicitHeight
                        Behavior on height { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }
                        opacity: hide ? 0 : 1
                        Behavior on opacity { NumberAnimation { duration: 100 } }

                        clip: true
                    }
                } // ListView
            } // ScrollView
        } // ColumnLayout (frame)
    } // Frame

    QPendingList{
        id: modelPendingTransactions
    }
}