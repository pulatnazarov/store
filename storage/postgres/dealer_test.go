package postgres

import (
	"context"
	"test/api/models"
	"test/config"
	"test/pkg/logger"
	"testing"
)

func TestDealerRepo_AddSum(t *testing.T) {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)

	pgStore, err := New(context.Background(), cfg, log)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	type testCases struct {
		Name   string        `json:"name"`
		Dealer models.Dealer `json:"dealer"`
		Want   interface{}   `json:"want"`
	}
	dealer := []testCases{
		{
			Name: "success",
			Dealer: models.Dealer{
				ID:   "1cfd84e6-72cb-4135-a802-85d10e4183ea",
				Name: "main dealer",
				Sum:  1000,
			},
			Want: 6000,
		},
		{
			Name: "rows affected",
			Dealer: models.Dealer{
				ID:   "b4a6eca7-5d22-4d93-934d-023f09422a25",
				Name: "some dealer",
				Sum:  5000,
			},
			Want: 0,
		},
		//{
		//  Name: "no update",
		//  Dealer: models.Dealer{
		//    ID:   "1cfd84e6-72cb-4135-a802-85d10e4183ea",
		//    Name: "main dealer",
		//    Sum:  1000,
		//  },
		//  Want: 1000,
		//},

	}

	for _, d := range dealer {
		t.Run(d.Name, func(t *testing.T) {
			err = pgStore.Dealer().AddSum(context.Background(), d.Dealer.Sum, d.Dealer.ID)
			if err != nil {
				t.Errorf("error is while adding sum to dealer error:%v", err)
				return
			}
			dealerData, err := pgStore.Dealer().Get(context.Background(), models.PrimaryKey{ID: d.Dealer.ID})
			if err != nil {
				t.Errorf("error is while getting dealer error:%v", err)
				return
			}
			if dealerData.Sum != d.Want {
				t.Errorf("expected %v but got %v", d.Want, dealerData.Sum)
			}
		})

	}

}

func TestDealerRepo_Get(t *testing.T) {
	cfg := config.Load()

	log := logger.New(cfg.ServiceName)

	pgStore, err := New(context.Background(), cfg, log)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	type testCases struct {
		Name string            `json:"name"`
		IDs  models.PrimaryKey `json:"i_ds"`
		Want interface{}       `json:"want"`
	}
	dealer := []testCases{
		{
			Name: "success",
			IDs: models.PrimaryKey{
				ID: "e0833898-8c9c-44fc-bb94-ba4420b783bb",
			},
			Want: 0,
		},
		{
			Name: "no rows in result set",
			IDs: models.PrimaryKey{
				ID: "fe692e48-2ba0-43d1-9bd1-d5025ef388d0",
			},
			Want: 0,
		},
	}

	for _, d := range dealer {
		t.Run(d.Name, func(t *testing.T) {
			dealerData, err := pgStore.Dealer().Get(context.Background(), models.PrimaryKey{ID: d.IDs.ID})
			if err != nil {
				t.Errorf("error is while getting dealer error: %v", err)
				return
			}
			if d.Want != dealerData.Sum {
				t.Errorf("expected %v but got %v", d.Want, dealerData.Sum)
			}
			if dealerData.ID != "" && dealerData.Name == "" {
				t.Error("expected some name but got nothing")
			}
		})
	}

}
