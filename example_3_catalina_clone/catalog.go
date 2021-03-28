package catalina

import (
	"encoding/json"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rezaindrag/golang-grpc/example_3_catalina_clone/model"
)

type Catalog struct {
	ID          string    `json:"id"`
	Provider    Provider  `json:"provider"`
	Items       []Promo   `json:"items"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedTime time.Time `json:"-"`
	UpdatedTime time.Time `json:"-"`
	Image       Image     `json:"image"`
	Cover       Image     `json:"cover"`
}

func (c Catalog) MarshalJSON() ([]byte, error) {
	type Alias Catalog
	alias := (Alias)(c)

	if len(c.Items) == 0 {
		alias.Items = make([]Promo, 0)
	}

	return json.Marshal(alias)
}

func (c Catalog) ConvertToOriginalStruct(catalog *model.Catalog) Catalog {
	return Catalog{
		ID: catalog.Id,
	}
}

func (c Catalog) ConvertToGRPCStruct() *model.Catalog {
	items := make([]*model.Promo, len(c.Items))
	for index, promoItem := range c.Items {
		items[index] = &model.Promo{
			Id:    promoItem.ID,
			Title: promoItem.Title,
		}
	}

	return &model.Catalog{
		Id:        c.ID,
		Items:     items,
		StartTime: timestamppb.New(c.StartTime),
		EndTime:   timestamppb.New(c.EndTime),
	}
}
