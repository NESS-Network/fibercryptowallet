import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12 

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "../Controls" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogSelectAddressByWallet

    property alias model: listView.model
    property alias filterString: textFieldFilterWallet.text
    property string selectedAddress

    padding: 0
    standardButtons: Dialog.Cancel

    closePolicy: Dialog.CloseOnPressOutside
    onAboutToShow: {
        textFieldFilterWallet.forceActiveFocus()
    }

    ColumnLayout {
        anchors.fill: parent

        TextField {
            id: textFieldFilterWallet

            Layout.fillWidth: true
            Layout.leftMargin: 10
            Layout.rightMargin: 10
            placeholderText: qsTr("Filter by wallet name")
            focus: true
            selectByMouse: true
        }

        ListView {
            id: listView

            Layout.fillWidth: true
            Layout.fillHeight: true
            clip: true
            currentIndex: -1

            delegate: MenuItem {
                readonly property string parentWallet: wallet
                readonly property bool matchFilter: !filterString || wallet.toLowerCase().includes(filterString.toLowerCase())

                width: parent.width
                height: matchFilter ? implicitHeight : 0
                Behavior on height { NumberAnimation { easing.type: Easing.OutQuint } }
                focusPolicy: Qt.NoFocus
                text: address
                font.family: "Code New Roman"
                Material.foreground: hovered ? parent.Material.accent : parent.Material.foreground
                highlighted: hovered
                leftPadding: highlighted ? 2*padding : padding // added
                Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added
            
                onHighlightedChanged: {
                    ListView.view.currentIndex = highlighted ? index : -1
                }

                onTriggered: {
                    selectedAddress = address
                    accept()
                }
            }

            section.property: "wallet"
            section.criteria: ViewSection.FullString
            section.delegate: Label {
                readonly property color textColor: (ListView.view.currentItem && ListView.view.currentItem.parentWallet === text) ? Material.accent : Material.foreground
                readonly property bool matchFilter: !filterString || section.toLowerCase().includes(filterString.toLowerCase())

                height: matchFilter ? contentHeight * 1.5 : 0
                Behavior on height { NumberAnimation { easing.type: Easing.OutQuint } }
                verticalAlignment: Text.AlignBottom
                text: section
                clip: true
                leftPadding: 12
                font.bold: true
                font.pointSize: Qt.application.font.pointSize * 1.5
                color: textColor
                Behavior on color { ColorAnimation { duration: 150 } }
            }

            ScrollIndicator.vertical: ScrollIndicator { }
        } // ListView
    } // ColumnLayout
}
