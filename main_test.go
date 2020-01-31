package sliceable

import (
    "testing"
)

func TestMapSlice(t *testing.T) {
    type Officer struct {
        ID   int
        Name string
    }

    // What you have
    var officers = []Officer{
        {ID: 20, Name: "Captain Piett"},
        {ID: 24, Name: "General Veers"},
        {ID: 56, Name: "Admiral Ozzel"},
        {ID: 88, Name: "Commander Jerjerrod"},
    }

    var ids []int

    if err := MapSlice(officers, func(thing interface{}) interface{} {
       return thing.(Officer).ID
    }, &ids); err != nil {
        t.Errorf("Received error %+v", err)
    }

    expectedTotal := 4

    if len(ids) != expectedTotal {
        t.Errorf("Received %d, Expected %d", len(ids), expectedTotal)
    }
}

func TestReduce(t *testing.T) {
    type Pilot struct {
        ID    int
        Name  string
        Years int
    }

    // What you have
    var pilots = []Pilot{
        {
            ID:    10,
            Name:  "Poe Dameron",
            Years: 14,
        },
        {
            ID:    2,
            Name:  "Temmin 'Snap' Wexley",
            Years: 30,
        },
        {
            ID:    41,
            Name:  "Tallissan Lintra",
            Years: 16,
        },
        {
            ID:    99,
            Name:  "Ello Asty",
            Years: 22,
        },
    }

    var totalYears int

    if err := Reduce(pilots, func(thing interface{}) int {
        return thing.(Pilot).Years
    }, &totalYears); err != nil {
        t.Errorf("Received error %+v", err)
    }

    expectedTotal := 82
    if totalYears != expectedTotal {
        t.Errorf("Received %d, Expected %d", totalYears, expectedTotal)
    }
}

func TestFilter(t *testing.T) {
    type Pilot struct {
        ID      int
        Name    string
        Faction string
    }

    // What you have
    var pilots = []Pilot{
        {
            ID:      2,
            Name:    "Wedge Antilles",
            Faction: "Rebels",
        },
        {
            ID:      8,
            Name:    "Ciena Ree",
            Faction: "Empire",
        },
        {
            ID:      40,
            Name:    "Iden Versio",
            Faction: "Empire",
        },
        {
            ID:      66,
            Name:    "Thane Kyrell",
            Faction: "Rebels",
        },
    }

    var filteredPilots []Pilot

    if err := Filter(pilots, func(thing interface{}) bool {
        return thing.(Pilot).Faction  == "Empire"
    }, &filteredPilots); err != nil {
        t.Errorf("Received error %+v", err)
    }

    expectedTotal := 2
    if len(filteredPilots) != expectedTotal {
        t.Errorf("Received %d, Expected %d", len(filteredPilots), expectedTotal)
    }
}
