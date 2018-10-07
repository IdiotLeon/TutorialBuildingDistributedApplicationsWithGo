package main

import "github.com/idiotLeon/TutorialBuildingDistributedApplicationsWithGo/mypackage"

func main() {
	println(mypackage.PublicVar)
	mypackage.PublicFunc()
}
