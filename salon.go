package main

import (
	 "fmt"
	 "sync"
 )


type salon struct{
	name string
	num_coiff int
	waiting_line_capacity int
	var wg sync.WaitGroup //file d'attente représentée par un waiting group
}
