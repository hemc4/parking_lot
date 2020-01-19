package dao

import (
//"fmt"
)

//TODO  optimize O(n) operations

//InMemoryDB - memory staorge
type InMemoryDB struct {
	maxSlots int
	slots    []*Slot
}

// NewInMemoryDB - constructor
func NewInMemoryDB(maxSlots int) (*InMemoryDB, error) {
	if maxSlots > 0 {
		m := new(InMemoryDB)
		m.maxSlots = maxSlots
		s := make([]*Slot, maxSlots)
		m.slots = s
		return m, nil
	}
	return nil, ErrInvalidMaxSlots
}

// Park - park a car
func (m *InMemoryDB) Park(v Vehicle) (*Slot, error) {
	slot, err := m.GetNextEmptySlot()
	if err != nil {
		return nil, err
	}

	slot.SetVehicle(v)
	i := slot.GetNo() - 1
	m.slots[i] = slot
	return slot, err
}

// Leave -
func (m *InMemoryDB) Leave(s *Slot) error {
	found := false
	for _, slot := range m.slots {
		if slot != nil && slot.GetNo() == s.GetNo() {
			found = true
			slot.SetNo(0)
			slot.SetVehicle(nil)
		}
	}
	if !found {
		return ErrCarNotFound
	}
	return nil
}

// GetAll -
func (m *InMemoryDB) GetAll() ([]*Slot, error) {
	return m.slots, nil
}

// GetAllSlotsByColour -
func (m *InMemoryDB) GetAllSlotsByColour(colour string) ([]*Slot, error) {
	s := make([]*Slot, m.maxSlots)
	for _, slot := range m.slots {
		if slot.GetVehicle() != nil {
			if slot.GetVehicle().GetColour() == colour {
				s = append(s, slot)
			}
		}
	}
	return s, nil
}

// GetSlotByRegNo -
func (m *InMemoryDB) GetSlotByRegNo(regNo string) (*Slot, error) {
	for _, slot := range m.slots {
		if slot.GetVehicle() != nil {
			if slot.GetVehicle().GetRegNo() == regNo {
				return slot, nil
			}
		}
	}
	return nil, ErrCarNotFound
}

// GetNextEmptySlot -
func (m *InMemoryDB) GetNextEmptySlot() (*Slot, error) {
	//fmt.Println(m)
	no := 0
	for _, slot := range m.slots {
		if slot != nil {
			if slot.GetNo() == 0 {
				break
			} else {
				no++
			}
		}
	}
	if no+1 > m.maxSlots {
		return nil, ErrMaxSlotReached
	}
	s := NewSlot(no + 1)
	return s, nil
}
