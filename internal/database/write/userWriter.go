package write

import (
	"backendAleshinaFeklistova/internal/database/read"
	"backendAleshinaFeklistova/internal/models"
)

func (uw *UserWriteRepository) Create() {
	executeSqlUser(1, uw)
}

func (uw *UserWriteRepository) Edit() {
	executeSqlUser(2, uw)
}

func (uw *UserWriteRepository) Delete() {
	executeSqlUser(3, uw)
}

func executeSqlUser(operId int, uw *UserWriteRepository) {
	if operId == 1 {
		read.TestData = append(read.TestData, uw.User)
	}
	if operId == 2 {
		for i, datum := range read.TestData {
			if datum.Id == uw.Id {
				read.TestData[i] = uw.User
			}
		}
	}
	if operId == 3 {
		for i, datum := range read.TestData {
			if datum.Id == uw.Id {
				read.TestData[i] = read.TestData[len(read.TestData)-1]
				read.TestData[len(read.TestData)-1] = models.User{}
				read.TestData = read.TestData[:len(read.TestData)-1]
			}
		}
	}
}
