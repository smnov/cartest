package main

type Database interface {
	Get()
	Delete()
	Edit()
	Add()
}

type MockDB struct{}

func (db MockDB) Get() {
}

func (db MockDB) Delete() {
}

func (db MockDB) Edit() {
}

func (db MockDB) Add() {
}
