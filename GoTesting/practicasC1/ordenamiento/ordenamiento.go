package ordenamiento

import (
	"errors"
	"sort"
)

func OrdenarSlice(collection []int) ([]int, error) {

	if len(collection) == 0 {
		return nil, errors.New("no es posible ordenar una lista vacia")
		// return nil, errors.New("lista vacia")
	}

	sort.Ints(collection)

	// return nil, nil
	return collection, nil
}
