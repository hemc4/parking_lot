package dao

//InMemoryDB - memory staorge
type InMemoryDB struct {
	maxSlots               int
	slots                  []Slot
	regNoSlotIndexMap      map[string]int
	colourSlotIndexListMap map[string][]int
}

// NewInMemoryDB - constructor
func NewInMemoryDB(maxSlots int) (*InMemoryDB, error) {
	if maxSlots > 0 {
		m := new(InMemoryDB)
		m.maxSlots = maxSlots
		s := make([]Slot, maxSlots)
		m.slots = s
		m.regNoSlotIndexMap = make(map[string]int)
		m.colourSlotIndexListMap = make(map[string][]int)
		return m, nil
	}
	return nil, ErrInvalidMaxSlots
}

// Park - park a car O(n)
func (m *InMemoryDB) Park(v Vehicle) (Slot, error) {
	slot, err := m.GetNextEmptySlot()
	if err != nil {
		return nil, err
	}

	slot.SetVehicle(v)
	i := slot.GetNo() - 1
	m.slots[i] = slot
	m.regNoSlotIndexMap[v.GetRegNo()] = i
	if indexList, ok := m.colourSlotIndexListMap[v.GetColour()]; ok {
		indexList = append(indexList, i)
		delete(m.colourSlotIndexListMap, v.GetColour())
		m.colourSlotIndexListMap[v.GetColour()] = indexList
	} else {
		slotIndexList := []int{i}
		m.colourSlotIndexListMap[v.GetColour()] = slotIndexList
	}
	return slot, err
}

// Leave - O(k) k->no of parked vechicles for a color
func (m *InMemoryDB) Leave(s Slot) error {
	i := s.GetNo() - 1
	slot := m.slots[i]
	if slot != nil {
		slot.SetNo(0)
		slot.SetVehicle(nil)
		if slot.GetVehicle() != nil {
			leavingColour := slot.GetVehicle().GetColour()
			leavingRegNo := slot.GetVehicle().GetColour()

			slotIndexList := m.colourSlotIndexListMap[leavingColour]
			newSlotList := findAndDeleteItem(slotIndexList, i)
			delete(m.colourSlotIndexListMap, leavingColour)
			m.colourSlotIndexListMap[leavingColour] = newSlotList

			if _, ok := m.regNoSlotIndexMap[leavingRegNo]; ok {
				delete(m.regNoSlotIndexMap, leavingRegNo)
			}
		}
		return nil
	}
	return ErrCarNotFound
}

// GetAll - O(1)
func (m *InMemoryDB) GetAll() ([]Slot, error) {
	return m.slots, nil
}

// GetAllSlotsByColour - O(k) - k->no of parked vechicles for a color
func (m *InMemoryDB) GetAllSlotsByColour(colour string) ([]Slot, error) {
	s := make([]Slot, m.maxSlots)
	slotIndexList := m.colourSlotIndexListMap[colour]
	for _, slotIndex := range slotIndexList {
		if m.slots[slotIndex] != nil {
			s = append(s, m.slots[slotIndex])
		}
	}
	return s, nil
}

// GetSlotByRegNo - O(1)
func (m *InMemoryDB) GetSlotByRegNo(regNo string) (Slot, error) {
	if slotindex, ok := m.regNoSlotIndexMap[regNo]; ok {
		if m.slots[slotindex] != nil {
			return m.slots[slotindex], nil
		}
	}
	return nil, ErrCarNotFound
}

// GetNextEmptySlot - O(n)
func (m *InMemoryDB) GetNextEmptySlot() (Slot, error) {
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
