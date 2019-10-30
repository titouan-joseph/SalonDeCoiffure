package salon

import "sync"

type Salon struct{
	name                  string
	num_coiff             int
	Waiting_line_capacity int
	Wg                    sync.WaitGroup //file d'attente représentée par un waiting group
}
