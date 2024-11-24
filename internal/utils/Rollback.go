package utils

import (
	"os"
)

func BackupFile() error {
	filePaths := []string{
		"/home/enigma/GolandProjects/mnc_test/data/customer.json",
		"/home/enigma/GolandProjects/mnc_test/data/transaction.json",
		"/home/enigma/GolandProjects/mnc_test/data/merchant.json",
	}
	for i := range filePaths {
		backupPath := filePaths[i] + ".backup"
		input, err := os.ReadFile(filePaths[i])
		if err != nil {
			return err
		}

		err = os.WriteFile(backupPath, input, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func RollbackFile() error {
	filePaths := []string{
		"/home/enigma/GolandProjects/mnc_test/data/customer.json",
		"/home/enigma/GolandProjects/mnc_test/data/transaction.json",
		"/home/enigma/GolandProjects/mnc_test/data/merchant.json",
	}
	for i := range filePaths {
		input, err := os.ReadFile(filePaths[i] + ".backup")
		if err != nil {
			return err
		}

		err = os.WriteFile(filePaths[i], input, 0644)
		if err != nil {
			return err
		}

		return os.Remove(filePaths[i] + ".backup")
	}
	return nil
}
