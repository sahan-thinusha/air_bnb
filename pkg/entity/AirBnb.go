package entity

type AirBnb struct {
	Id                   string `json:"id" bson:"_id"`
	Base                 `bson:",inline"`
	ListingUrl           string `json:"listing_url"`
	Name                 string `json:"name"`
	Summary              string `json:"summary"`
	Space                string `json:"space"`
	Description          string `json:"description"`
	NeighborhoodOverview string `json:"neighborhood_overview"`
	Notes                string `json:"notes"`
	Transit              string `json:"transit"`
	Access               string `json:"access"`
	Interaction          string `json:"interaction"`
	HouseRules           string `json:"house_rules"`
	PropertyType         string `json:"property_type"`
	RoomType             string `json:"room_type"`
	BedType              string `json:"bed_type"`
	MinimumNights        string `json:"minimum_nights"`
	MaximumNights        string `json:"maximum_nights"`
	CancellationPolicy   string `json:"cancellation_policy"`
	LastScraped          struct {
		Date struct {
			NumberLong string `json:"$numberLong"`
		} `json:"$date"`
	} `json:"last_scraped"`
	CalendarLastScraped struct {
		Date struct {
			NumberLong string `json:"$numberLong"`
		} `json:"$date"`
	} `json:"calendar_last_scraped"`
	Accommodates struct {
		NumberInt string `json:"$numberInt"`
	} `json:"accommodates"`
	Bedrooms struct {
		NumberInt string `json:"$numberInt"`
	} `json:"bedrooms"`
	Beds struct {
		NumberInt string `json:"$numberInt"`
	} `json:"beds"`
	NumberOfReviews struct {
		NumberInt string `json:"$numberInt"`
	} `json:"number_of_reviews"`
	Bathrooms struct {
		NumberDecimal string `json:"$numberDecimal"`
	} `json:"bathrooms"`
	Amenities []string `json:"amenities"`
	Price     struct {
		NumberDecimal string `json:"$numberDecimal"`
	} `json:"price"`
	ExtraPeople struct {
		NumberDecimal string `json:"$numberDecimal"`
	} `json:"extra_people"`
	GuestsIncluded struct {
		NumberDecimal string `json:"$numberDecimal"`
	} `json:"guests_included"`
	Images struct {
		ThumbnailUrl string `json:"thumbnail_url"`
		MediumUrl    string `json:"medium_url"`
		PictureUrl   string `json:"picture_url"`
		XlPictureUrl string `json:"xl_picture_url"`
	} `json:"images"`
	Host struct {
		HostId               string `json:"host_id"`
		HostUrl              string `json:"host_url"`
		HostName             string `json:"host_name"`
		HostLocation         string `json:"host_location"`
		HostAbout            string `json:"host_about"`
		HostThumbnailUrl     string `json:"host_thumbnail_url"`
		HostPictureUrl       string `json:"host_picture_url"`
		HostNeighbourhood    string `json:"host_neighbourhood"`
		HostIsSuperhost      bool   `json:"host_is_superhost"`
		HostHasProfilePic    bool   `json:"host_has_profile_pic"`
		HostIdentityVerified bool   `json:"host_identity_verified"`
		HostListingsCount    struct {
			NumberInt string `json:"$numberInt"`
		} `json:"host_listings_count"`
		HostTotalListingsCount struct {
			NumberInt string `json:"$numberInt"`
		} `json:"host_total_listings_count"`
		HostVerifications []string `json:"host_verifications"`
	} `json:"host"`
	Address struct {
		Street         string `json:"street"`
		Suburb         string `json:"suburb"`
		GovernmentArea string `json:"government_area"`
		Market         string `json:"market"`
		Country        string `json:"country"`
		CountryCode    string `json:"country_code"`
		Location       struct {
			Type        string `json:"type"`
			Coordinates []struct {
				NumberDouble string `json:"$numberDouble"`
			} `json:"coordinates"`
			IsLocationExact bool `json:"is_location_exact"`
		} `json:"location"`
	} `json:"address"`
	Availability struct {
		Availability30 struct {
			NumberInt string `json:"$numberInt"`
		} `json:"availability_30"`
		Availability60 struct {
			NumberInt string `json:"$numberInt"`
		} `json:"availability_60"`
		Availability90 struct {
			NumberInt string `json:"$numberInt"`
		} `json:"availability_90"`
		Availability365 struct {
			NumberInt string `json:"$numberInt"`
		} `json:"availability_365"`
	} `json:"availability"`
	ReviewScores struct {
	} `json:"review_scores"`
	Reviews []interface{} `json:"reviews"`
}
