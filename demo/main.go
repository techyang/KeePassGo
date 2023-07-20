package main

import (
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"log"
)

type Group struct {
	Name   string
	Groups []Group
}

func main() {
	widgets.NewQApplication(0, nil)

	// Open the KeePass database
	db := openDatabase("path/to/your/keepass.kdbx", "password")
	if db == nil {
		return
	}

	// Create the main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("KeePass Group Tree")

	// Create the tree view
	treeView := widgets.NewQTreeWidget(nil)
	treeView.SetHeaderLabels([]string{"Group"})

	// Build the group tree and add it to the tree view
	groups := buildGroupTree(db.Content.Root.Groups)
	addGroupsToTreeView(treeView, groups)

	// Handle itemClicked signal
	treeView.ConnectItemClicked(func(item *widgets.QTreeWidgetItem, column int) {
		// Get the selected group name
		groupName := item.Text(0)

		// Find the group in the KeePass database
		selectedGroup := findGroupByName(db.Content.Root.Groups, groupName)
		if selectedGroup == nil {
			return
		}

		// Display the passwords for the selected group
		displayPasswords(selectedGroup.Entries)
	})

	// Set the tree view as the central widget of the main window
	window.SetCentralWidget(treeView)

	// Show the main window
	window.Show()

	widgets.QApplication_Exec()
}

// Function to open the KeePass database
func openDatabase(filePath, password string) *gokeepasslib.Database {
	db, err := gokeepasslib.OpenDatabase(filePath, []byte(password))
	if err != nil {
		log.Println("Error opening KeePass database:", err)
		return nil
	}
	return db
}

// Function to build the group tree from the KeePass database
func buildGroupTree(groups []*gokeepasslib.Group) []Group {
	var result []Group
	for _, group := range groups {
		groupItem := Group{
			Name: group.Name,
		}
		groupItem.Groups = buildGroupTree(group.Groups)
		result = append(result, groupItem)
	}
	return result
}

// Function to add groups to the tree view
func addGroupsToTreeView(treeView *widgets.QTreeWidget, groups []Group) {
	for _, group := range groups {
		treeItem := widgets.NewQTreeWidgetItem2([]string{group.Name}, 0)
		addGroupsToTreeView(treeItem, group.Groups)
		treeView.AddTopLevelItem(treeItem)
	}
}

// Function to find a group in the group tree by its name
func findGroupByName(groups []*gokeepasslib.Group, name string) *gokeepasslib.Group {
	for _, group := range groups {
		if group.Name == name {
			return group
		}
		if foundGroup := findGroupByName(group.Groups, name); foundGroup != nil {
			return foundGroup
		}
	}
	return nil
}

// Function to display the passwords in a group
func displayPasswords(entries []*gokeepasslib.Entry) {
	for _, entry := range entries {
		// Here, you can access the information in each entry
		log.Println("Title:", entry.GetTitle())
		log.Println("Username:", entry.GetUsername())
		log.Println("Password:", entry.GetPassword())
		log.Println("URL:", entry.GetURL())
		log.Println("Notes:", entry.GetNotes())
		log.Println("----------------------")
	}
}
