package main

import "fmt"

type Storage struct {
	cleaningAlgo CleaningAlgorithm
}

func (s Storage) CleanStorage() {
	s.cleaningAlgo.clean()
}

type CleaningAlgorithm interface {
	clean()
}

type DefaultAlgo struct{}

func (a DefaultAlgo) clean() {
	fmt.Println("Cleaning storage using default algo")
}

type UpgradedAlgo struct{}

func (a UpgradedAlgo) clean() {
	fmt.Println("Cleaning storage using upgraded algo")
}

func main() {
	defaultAlgo := DefaultAlgo{}
	upgradedAlgo := UpgradedAlgo{}

	storage := &Storage{cleaningAlgo: defaultAlgo}
	storage.CleanStorage()

	storage.cleaningAlgo = upgradedAlgo
	storage.CleanStorage()
}
