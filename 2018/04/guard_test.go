package day04

import "testing"

//Date   ID   Minute
//            000000000011111111112222222222333333333344444444445555555555
//            012345678901234567890123456789012345678901234567890123456789
//11-01  #10  .....####################.....#########################.....
//11-02  #99  ........................................##########..........
//11-03  #10  ........................#####...............................
//11-04  #99  ....................................##########..............
//11-05  #99  .............................................##########.....
var guards = GuardMap{
	10: {
		ID: 10,
		Sleeps: []int{0,0,0,0,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,2,1,1,1,1,0,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0,0,0,0,0},
		Slept: 50,
	},
	99: {
		ID: 99,
		Sleeps: []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,1,1,1,2,2,2,2,3,2,2,2,2,1,1,1,1,1,1,0,0,0,0,0},
		Slept: 30,
	},
}

func TestGuard_MostSleptMinute(t *testing.T) {
	g := guards[10]
	m, _ := g.MostSleptMinute()
	if m != 24 {
		t.Errorf("Expected most slept minute at 24, got %d instead", m)
	}
}

func TestGuardMap_MostSleep(t *testing.T) {
	g := guards.MostAsleep()
	if g.ID != 10 {
		t.Errorf("Expected guard 10 to have the most sleep, got guard %d instead", g.ID)
	}
}

func TestGuardMap_MostSleepAtSameTime(t *testing.T) {
	g := guards.MostSleptAtSameTime()
	if g.ID != 99 {
		t.Errorf("Expected guard 99 to have the most frequent sleep, got guard %d instead", g.ID)
	}
}
