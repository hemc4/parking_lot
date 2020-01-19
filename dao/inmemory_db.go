package dao

import (
	//"fmt"
)
//TODO  optimize O(n) operations

//InMemoryDB - memory staorge
type InMemoryDB struct{
	maxSlots int
	slots []*Slot
}

// NewInMemoryDB - constructor
func NewInMemoryDB(maxSlots int) (*InMemoryDB, error){
	if maxSlots > 0 {
		m := new(InMemoryDB);
		m.maxSlots= maxSlots;
		s := make([]*Slot, maxSlots)
		m.slots = s
		return m, nil
	}
	return nil, ErrInvalidMaxSlots
}

// Park - park a car
func (m *InMemoryDB) Park(c *Car) (*Slot, error){
	slot, err := m.GetNextEmptySlot();
	if err != nil {
		return nil, err
	}
	
	slot.Car = c
	i :=  slot.No - 1
	m.slots[i] = slot
	return slot, err
}

// Leave - 
func (m *InMemoryDB) Leave(s *Slot) (error){
	found := false
	for _,slot := range m.slots {
		if slot != nil && slot.No == s.No {
			found = true
			slot.No = 0
			slot.Car = nil
		}
	}
	if !found {
		return ErrCarNotFound
	}
	return nil
}
	
// GetAll - 
func (m *InMemoryDB) GetAll() ([]*Slot, error){
	return m.slots, nil
}

// GetAllSlotsByColour - 
func (m *InMemoryDB) GetAllSlotsByColour(colour string) ([]*Slot, error){
	s := make([]*Slot, m.maxSlots)
	for _,slot := range m.slots {
		if slot.Car != nil {
			if slot.Car.Colour == colour {
				s =append(s, slot)
			}
		}
	}
	return s, nil
}

// GetSlotByRegNo - 
func (m *InMemoryDB) GetSlotByRegNo(regNo string) (*Slot, error){
	for _,slot := range m.slots {
		if slot.Car != nil {
			if slot.Car.RegNo == regNo {
				return slot, nil
			}
		}
	}
	return nil, ErrCarNotFound
}

// GetNextEmptySlot -   
func (m *InMemoryDB) GetNextEmptySlot()  (*Slot, error) {
	//fmt.Println(m)
	no := 0;
	for _,slot := range m.slots {
		if slot != nil {
			if slot.No == 0 {
				break
			}else{
				no++;
			}
		}
	}
	if no +1 > m.maxSlots {
		return nil, ErrMaxSlotReached
	}
	s := NewSlot(no+1)
	return s, nil
}