package entity

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
)

type KeePassTree struct {
	*widgets.QTreeWidget
	KeePassTable *KeePassTable
}

func NewKeePassTree(tableWidget *KeePassTable) *KeePassTree {
	treeWidget := &KeePassTree{
		QTreeWidget:  widgets.NewQTreeWidget(nil),
		KeePassTable: tableWidget,
	}
	//treeWidget.SetHeaderLabels([]string{"yangwl"})
	file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("111111")
	_ = gokeepasslib.NewDecoder(file).Decode(db)
	treeWidget.SetRootIsDecorated(false) // Hide root item's expand arrow

	db.UnlockProtectedEntries()
	rootGroups := db.Content.Root.Groups
	for i, rootGroup := range rootGroups {
		fmt.Println(i, "rootGroup:", rootGroup.Name)
		// Create the root item
		rootItem := widgets.NewQTreeWidgetItem4(treeWidget, []string{rootGroup.Name, "1.1"}, 0)
		rootItem.SetExpanded(true) // Set the root item initially expanded
		icon := gui.NewQIcon5("Ext/Images_App_HighRes/Nuvola/B48x48_KGPG_Key1.png")
		rootItem.SetIcon(0, icon)

		groups := rootGroup.Groups
		buildGroupTree(rootItem, groups)

		treeWidget.InsertTopLevelItem(i, rootItem)
	}
	// Create the root item

	// Set the root item as the top-level item of the tree widget

	treeWidget.SetHeaderHidden(true)

	// Connect the itemClicked signal of the tree widget
	treeWidget.TreeItemClicked(tableWidget, rootGroups)
	return treeWidget
}

func (treeWidget *KeePassTree) LoadKeePassTree(dbPath string, keePassTable *KeePassTable) {

	//treeWidget.SetHeaderLabels([]string{"yangwl"})
	file, _ := os.Open(dbPath)

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("111111")
	_ = gokeepasslib.NewDecoder(file).Decode(db)

	db.UnlockProtectedEntries()
	rootGroups := db.Content.Root.Groups
	for i, rootGroup := range rootGroups {
		fmt.Println(i, "rootGroup:", rootGroup.Name)
		// Create the root item
		rootItem := widgets.NewQTreeWidgetItem4(treeWidget, []string{rootGroup.Name, "1.1"}, 0)
		rootItem.SetExpanded(true) // Set the root item initially expanded
		groups := rootGroup.Groups
		buildGroupTree(rootItem, groups)

		treeWidget.InsertTopLevelItem(i, rootItem)
	}
	// Create the root item

	// Set the root item as the top-level item of the tree widget

	treeWidget.SetHeaderHidden(true)

	// Connect the itemClicked signal of the tree widget
	treeWidget.TreeItemClicked(keePassTable, rootGroups)
}

func buildGroupTree(parent *widgets.QTreeWidgetItem, groups []gokeepasslib.Group) {
	for _, group := range groups {
		txt, _ := group.UUID.MarshalText()

		fmt.Println("group.UUID -----------:", group.Name, ":", string(txt))
		treeItem := widgets.NewQTreeWidgetItem2([]string{group.Name}, 0)
		//treeItem.SetData(0, core.Qt__UserRole, core.NewQVariant1(group.UUID.String()))

		treeItem.SetData(1, 0, core.NewQVariant1(string(txt)))
		icon := gui.NewQIcon5("Ext/Images_Client_HighRes/C48_Folder.png")
		treeItem.SetIcon(0, icon)
		parent.AddChild(treeItem)
		buildGroupTree(treeItem, group.Groups)
	}
}

func (treeWidget *KeePassTree) TreeItemClicked(tableWidget *KeePassTable, rootGroups []gokeepasslib.Group) {

	treeWidget.ConnectItemClicked(func(item *widgets.QTreeWidgetItem, column int) {

		groupUUID := item.Data(1, 0).ToString()
		fmt.Println(item.Text(0), "点击了", groupUUID)
		tableWidget.SetObjectName(groupUUID)
		group := FindGroupByUUID(rootGroups, groupUUID)

		if group != nil && group.Entries != nil {
			tableWidget.SetTableItems(group)
		} else {
			tableWidget.SetRowCount(0)
		}
	})

}
