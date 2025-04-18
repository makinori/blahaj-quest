package blahaj

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"

	"github.com/makinori/blahaj-quest/common"

	"github.com/charmbracelet/log"
)

const maxConcurrency int = 12 // requests at a time

// for api

type BlahajStore struct {
	Quantity     int    `json:"quantity"`
	Name         string `json:"name"`
	Lat          string `json:"lat"`
	Lng          string `json:"lng"`
	CountryCode  string `json:"countryCode"`
	LanguageCode string `json:"languageCode"`
	ItemCode     string `json:"itemCode"`
}

type BlahajData struct {
	Updated time.Time     `json:"updated"`
	Data    []BlahajStore `json:"data"`
}

// from api

type IkeaStore struct {
	ID   string
	Name string
	Lat  string
	Lng  string
}

// cache

func getFromAPI[T any](t *T, url string, headers map[string]string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := http.Client{}
	client.Do(req)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, t)
}

func getStockForCountry(country BlahajDbCountry) ([]BlahajStore, error) {
	// get stores information

	var stores []IkeaStore

	err := getFromAPI(
		&stores,
		strings.Join([]string{
			"https://www.ikea.com/", country.CountryCode,
			"/", country.LanguageCode,
			"/meta-data/informera/stores-detailed.json",
		}, ""),
		map[string]string{},
	)

	if err != nil {
		return []BlahajStore{}, err
	}

	// add additional and deduplicate

	if len(country.AdditionalStores) > 0 {
		for _, additionalStore := range country.AdditionalStores {
			foundIndex := slices.IndexFunc(stores, func(needle IkeaStore) bool {
				return additionalStore.ID == needle.ID
			})
			if foundIndex > -1 {
				continue
			}
			stores = append(stores, additionalStore)
		}
	}

	// get stock from all stores

	var stock struct {
		Availabilities []struct {
			BuyingOption struct {
				CashCarry struct {
					Availability struct {
						Quantity int `json:"quantity,omitempty"`
					} `json:"availability"`
				} `json:"cashCarry"`
			} `json:"buyingOption"`
			ClassUnitKey struct {
				ClassUnitCode string `json:"classUnitCode"`
			} `json:"classUnitKey"`
		} `json:"availabilities"`
	}

	stockUrl, err := url.Parse(
		"https://api.ingka.ikea.com/cia/availabilities/ru/" + country.CountryCode,
	)

	if err != nil {
		return []BlahajStore{}, err
	}

	stockUrlQuery := url.Values{
		"itemNos": []string{country.ItemCode},
		// "expand": []string{"StoresList,Restocks,SalesLocations"},
		"expand": []string{"StoresList"},
	}

	stockUrl.RawQuery = stockUrlQuery.Encode()

	err = getFromAPI(
		&stock,
		stockUrl.String(),
		map[string]string{
			"Accept":  "application/json;version=2",
			"Referer": "https://www.ikea.com/",
			// "X-Client-Id": "b6c117e5-ae61-4ef5-b4cc-e0b1e37f0631",
			"X-Client-Id": "ef382663-a2a5-40d4-8afe-f0634821c0ed",
		},
	)

	if err != nil {
		return []BlahajStore{}, err
	}

	// map stock with stores

	var blahajStores []BlahajStore

	for _, storeAvail := range stock.Availabilities {
		quantity := storeAvail.BuyingOption.CashCarry.Availability.Quantity
		storeID := storeAvail.ClassUnitKey.ClassUnitCode

		storeIndex := slices.IndexFunc(stores, func(needle IkeaStore) bool {
			return needle.ID == storeID
		})
		if storeIndex == -1 {
			continue
		}

		store := stores[storeIndex]

		blahajStores = append(blahajStores, BlahajStore{
			Quantity:     quantity,
			Name:         store.Name,
			Lat:          store.Lat,
			Lng:          store.Lng,
			CountryCode:  country.CountryCode,
			LanguageCode: country.LanguageCode,
			ItemCode:     country.ItemCode,
		})
	}

	return blahajStores, nil
}

func GetBlahajData() BlahajData {
	var data BlahajData

	err := common.GetCache("blahajData", &data)
	if err == nil {
		return data
	}

	// get fresh

	var dataMutex sync.Mutex

	var sem = semaphore.NewWeighted(int64(maxConcurrency))
	ctx := context.Background()

	for _, country := range BlahajDb {
		sem.Acquire(ctx, 1)
		go func() {
			defer sem.Release(1)

			newStores, err := getStockForCountry(country)
			if err != nil {
				log.Warn(
					"failed to get stock for country: "+
						country.CountryCode+"/"+country.LanguageCode,
					"err", err,
				)
				return
			}

			dataMutex.Lock()
			data.Data = append(data.Data, newStores...)
			dataMutex.Unlock()
		}()
	}

	sem.Acquire(ctx, int64(maxConcurrency))

	log.Infof("fetched %d stores", len(data.Data))

	data.Updated = time.Now()

	err = common.SetCache("blahajData", data)
	if err != nil {
		log.Error(err)
	}

	return data
}

func GetBlahajDataJSON() ([]byte, error) {
	data := GetBlahajData()

	json, err := json.Marshal(data)
	if err != nil {
		return []byte{}, err
	}

	return json, nil
}
