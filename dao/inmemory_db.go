package dao

//InMemoryDB - memory staorge
type InMemoryDB struct{
	maxSlots int
	slots []*Slot
}

// NewInMemoryDB - constructor
func NewInMemoryDB(maxSlots int) *InMemoryDB{
	m := new(InMemoryDB);
	m.maxSlots= maxSlots;
	s := make([]*Slot, maxSlots)
	m.slots = s
	return m
}

// Park - park a car
func (m *InMemoryDB) Park(c *Car) (*Slot, error){
	slot, err := m.GetNextEmptySlot();
	if err != nil {
		return nil, err
	}
	
	slot.Car = c
	return slot, err
}

// Leave - 
func (m *InMemoryDB) Leave(s *Slot) (error){
	found := false
	for _,slot := range m.slots {
		if slot.No == s.No {
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

// GetAllCarsByColour - 
func (m *InMemoryDB) GetAllCarsByColour(string) ([]*Car, error){
	return nil,nil
}

// GetAllSlotsByColour - 
func (m *InMemoryDB) GetAllSlotsByColour(string) ([]*Slot, error){
	return m.slots, nil
}

// GetCarByRegNo - 
func (m *InMemoryDB) GetCarByRegNo(regNo string) (*Car, error){
	for _,slot := range m.slots {
		if slot.Car != nil {
			if slot.Car.RegNo == regNo {
				return slot.Car, nil
			}
		}
	}
	return nil, ErrCarNotFound
}

// GetNextEmptySlot -   
func (m *InMemoryDB) GetNextEmptySlot()  (*Slot, error) {
	no := 0;
	for _,slot := range m.slots {
		if slot.No == 0 {
			break
		}else{
			no++;
		}
	}
	if no +1 > m.maxSlots {
		return nil, ErrMaxSlotReached
	}
	s := NewSlot(no+1)
	return s, nil
}