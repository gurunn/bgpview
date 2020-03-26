package bgpview

const URL = "https://api.bgpview.io/"

type Status struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Meta          Meta   `json:"@meta"`
}

type Meta struct {
	TimeZone      string `json:"time_zone"`
	APIVersion    int    `json:"api_version"`
	ExecutionTime string `json:"execution_time"`
}

// ASN
type ASN struct {
	Data DataASN `json:"data"`
	Status
}

type RirAllocation struct {
	RirName          string `json:"rir_name"`
	CountryCode      string `json:"country_code"`
	DateAllocated    string `json:"date_allocated"`
	AllocationStatus string `json:"allocation_status"`
}

type IanaAssignment struct {
	AssignmentStatus string      `json:"assignment_status"`
	Description      string      `json:"description"`
	WhoisServer      string      `json:"whois_server"`
	DateAssigned     interface{} `json:"date_assigned"`
}

type DataASN struct {
	Asn               int            `json:"asn"`
	Name              string         `json:"name"`
	DescriptionShort  string         `json:"description_short"`
	DescriptionFull   []string       `json:"description_full"`
	CountryCode       string         `json:"country_code"`
	Website           string         `json:"website"`
	EmailContacts     []string       `json:"email_contacts"`
	AbuseContacts     []string       `json:"abuse_contacts"`
	LookingGlass      string         `json:"looking_glass"`
	TrafficEstimation string         `json:"traffic_estimation"`
	TrafficRatio      string         `json:"traffic_ratio"`
	OwnerAddress      []string       `json:"owner_address"`
	RirAllocation     RirAllocation  `json:"rir_allocation"`
	IanaAssignment    IanaAssignment `json:"iana_assignment"`
	DateUpdated       string         `json:"date_updated"`
}

// ASNPrefixes
type ASNPrefixes struct {
	Data DataASNPrefixes `json:"data"`
	Status
}

type Parent struct {
	Prefix           string `json:"prefix"`
	IP               string `json:"ip"`
	Cidr             int    `json:"cidr"`
	RirName          string `json:"rir_name"`
	AllocationStatus string `json:"allocation_status"`
}

type Ipv4Prefixes struct {
	Prefix      string `json:"prefix"`
	IP          string `json:"ip"`
	Cidr        int    `json:"cidr"`
	RoaStatus   string `json:"roa_status"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
	Parent      Parent `json:"parent"`
}

type Ipv6Prefixes struct {
	Prefix      string `json:"prefix"`
	IP          string `json:"ip"`
	Cidr        int    `json:"cidr"`
	RoaStatus   string `json:"roa_status"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
	Parent      Parent `json:"parent"`
}

type DataASNPrefixes struct {
	Ipv4Prefixes []Ipv4Prefixes `json:"ipv4_prefixes"`
	Ipv6Prefixes []Ipv6Prefixes `json:"ipv6_prefixes"`
}

// ASNPeers
type ASNPeers struct {
	Data DataASNPeers `json:"data"`
	Status
}

type Ipv4Peers struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

type Ipv6Peers struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

type DataASNPeers struct {
	Ipv4Peers []Ipv4Peers `json:"ipv4_peers"`
	Ipv6Peers []Ipv6Peers `json:"ipv6_peers"`
}

// ASNUpstreams
type ASNUpstreams struct {
	Data DataASNUpstreams `json:"data"`
	Status
}

type Ipv4Upstreams struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

type Ipv6Upstreams struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

type DataASNUpstreams struct {
	Ipv4Upstreams []Ipv4Upstreams `json:"ipv4_upstreams"`
	Ipv6Upstreams []Ipv6Upstreams `json:"ipv6_upstreams"`
	Ipv4Graph     string          `json:"ipv4_graph"`
	Ipv6Graph     string          `json:"ipv6_graph"`
	CombinedGraph string          `json:"combined_graph"`
}

// ASNDownstreams
type ASNDownstreams struct {
	Data DataASNDownstreams `json:"data"`
	Status
}

type Ipv4Downstreams struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

type DataASNDownstreams struct {
	Ipv4Downstreams []Ipv4Downstreams `json:"ipv4_downstreams"`
	Ipv6Downstreams []interface{}     `json:"ipv6_downstreams"`
}

// ASNIXs
type ASNIXs struct {
	Data []DataASNIXs `json:"data"`
	Status
}

type DataASNIXs struct {
	IxID        int    `json:"ix_id"`
	Name        string `json:"name"`
	NameFull    string `json:"name_full"`
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
	Ipv4Address string `json:"ipv4_address"`
	Ipv6Address string `json:"ipv6_address"`
	Speed       int    `json:"speed"`
}

// Prefix
type Prefix struct {
	Data DataPrefix `json:"data"`
	Status
}

type PrefixUpstreams struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

type Asns struct {
	Asn             int               `json:"asn"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	CountryCode     string            `json:"country_code"`
	PrefixUpstreams []PrefixUpstreams `json:"prefix_upstreams"`
}

type CountryCodes struct {
	WhoisCountryCode         string `json:"whois_country_code"`
	RirAllocationCountryCode string `json:"rir_allocation_country_code"`
	MaxmindCountryCode       string `json:"maxmind_country_code"`
}

type RirAllocationPrefix struct {
	RirName          string `json:"rir_name"`
	CountryCode      string `json:"country_code"`
	IP               string `json:"ip"`
	Cidr             int    `json:"cidr"`
	Prefix           string `json:"prefix"`
	DateAllocated    string `json:"date_allocated"`
	AllocationStatus string `json:"allocation_status"`
}

type IanaAssignmentPrefix struct {
	AssignmentStatus string      `json:"assignment_status"`
	Description      string      `json:"description"`
	WhoisServer      string      `json:"whois_server"`
	DateAssigned     interface{} `json:"date_assigned"`
}

type Maxmind struct {
	CountryCode string `json:"country_code"`
	City        string `json:"city"`
}

type DataPrefix struct {
	Prefix           string               `json:"prefix"`
	IP               string               `json:"ip"`
	Cidr             int                  `json:"cidr"`
	Asns             []Asns               `json:"asns"`
	Name             string               `json:"name"`
	DescriptionShort string               `json:"description_short"`
	DescriptionFull  []string             `json:"description_full"`
	EmailContacts    []string             `json:"email_contacts"`
	AbuseContacts    []string             `json:"abuse_contacts"`
	OwnerAddress     []string             `json:"owner_address"`
	CountryCodes     CountryCodes         `json:"country_codes"`
	RirAllocation    RirAllocationPrefix  `json:"rir_allocation"`
	IanaAssignment   IanaAssignmentPrefix `json:"iana_assignment"`
	Maxmind          Maxmind              `json:"maxmind"`
	RelatedPrefixes  []interface{}        `json:"related_prefixes"`
	DateUpdated      string               `json:"date_updated"`
}

// IP
type IP struct {
	Data DataIP `json:"data"`
	Status
}

type Asn struct {
	Asn         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

type Prefixes struct {
	Prefix      string `json:"prefix"`
	IP          string `json:"ip"`
	Cidr        int    `json:"cidr"`
	Asn         Asn    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

type RirAllocationIP struct {
	RirName          string `json:"rir_name"`
	CountryCode      string `json:"country_code"`
	IP               string `json:"ip"`
	Cidr             string `json:"cidr"`
	Prefix           string `json:"prefix"`
	DateAllocated    string `json:"date_allocated"`
	AllocationStatus string `json:"allocation_status"`
}

type IanaAssignmentIP struct {
	AssignmentStatus string      `json:"assignment_status"`
	Description      string      `json:"description"`
	WhoisServer      string      `json:"whois_server"`
	DateAssigned     interface{} `json:"date_assigned"`
}

type MaxmindIP struct {
	CountryCode string      `json:"country_code"`
	City        interface{} `json:"city"`
}

type DataIP struct {
	IP             string           `json:"ip"`
	PtrRecord      interface{}      `json:"ptr_record"`
	Prefixes       []Prefixes       `json:"prefixes"`
	RirAllocation  RirAllocationIP  `json:"rir_allocation"`
	IanaAssignment IanaAssignmentIP `json:"iana_assignment"`
	Maxmind        MaxmindIP        `json:"maxmind"`
}

// IX
type IX struct {
	Data DataIX `json:"data"`
	Status
}

type DataIX struct {
	Name         string        `json:"name"`
	NameFull     string        `json:"name_full"`
	Website      string        `json:"website"`
	TechEmail    interface{}   `json:"tech_email"`
	TechPhone    interface{}   `json:"tech_phone"`
	PolicyEmail  interface{}   `json:"policy_email"`
	PolicyPhone  interface{}   `json:"policy_phone"`
	City         string        `json:"city"`
	CountryCode  string        `json:"country_code"`
	URLStats     interface{}   `json:"url_stats"`
	MembersCount int           `json:"members_count"`
	Members      []interface{} `json:"members"`
}

// Search
type Search struct {
	Data DataSearch `json:"data"`
	Status
}

type AsnsSearch struct {
	Asn           int      `json:"asn"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	CountryCode   string   `json:"country_code"`
	EmailContacts []string `json:"email_contacts"`
	AbuseContacts []string `json:"abuse_contacts"`
	RirName       string   `json:"rir_name"`
}

type Ipv4PrefixesSearch struct {
	Prefix        string   `json:"prefix"`
	IP            string   `json:"ip"`
	Cidr          int      `json:"cidr"`
	Name          string   `json:"name"`
	CountryCode   string   `json:"country_code"`
	Description   string   `json:"description"`
	EmailContacts []string `json:"email_contacts"`
	AbuseContacts []string `json:"abuse_contacts"`
	RirName       string   `json:"rir_name"`
	ParentPrefix  string   `json:"parent_prefix"`
	ParentIP      string   `json:"parent_ip"`
	ParentCidr    int      `json:"parent_cidr"`
}

type Ipv6PrefixesSearch struct {
	Prefix        string   `json:"prefix"`
	IP            string   `json:"ip"`
	Cidr          int      `json:"cidr"`
	Name          string   `json:"name"`
	CountryCode   string   `json:"country_code"`
	Description   string   `json:"description"`
	EmailContacts []string `json:"email_contacts"`
	AbuseContacts []string `json:"abuse_contacts"`
	RirName       string   `json:"rir_name"`
	ParentPrefix  string   `json:"parent_prefix"`
	ParentIP      string   `json:"parent_ip"`
	ParentCidr    int      `json:"parent_cidr"`
}

type DataSearch struct {
	Asns              []AsnsSearch         `json:"asns"`
	Ipv4Prefixes      []Ipv4PrefixesSearch `json:"ipv4_prefixes"`
	Ipv6Prefixes      []Ipv6PrefixesSearch `json:"ipv6_prefixes"`
	InternetExchanges []interface{}        `json:"internet_exchanges"`
}
