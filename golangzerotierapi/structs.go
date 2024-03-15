package golangzerotierapi

type Network struct {
	ID                    string         `json:"id"`
	Type                  string         `json:"type"`
	Clock                 int64          `json:"clock"`
	Config                *NetworkConfig `json:"config"`
	Description           string         `json:"description"`
	RulesSource           string         `json:"rulesSource"`
	OwnerID               string         `json:"ownerId"`
	OnlineMemberCount     int            `json:"onlineMemberCount"`
	AuthorizedMemberCount int            `json:"authorizedMemberCount"`
	TotalMemberCount      int            `json:"totalMemberCount"`
	CapabilitiesByName    struct {
	} `json:"capabilitiesByName"`
	TagsByName struct {
		Server struct {
			Default int `json:"default"`
			Enums   struct {
				False int `json:"false"`
				True  int `json:"true"`
			} `json:"enums"`
			Flags struct {
			} `json:"flags"`
			ID int `json:"id"`
		} `json:"server"`
	} `json:"tagsByName"`
	UI struct {
		MembersHelpCollapsed  bool `json:"membersHelpCollapsed"`
		RulesHelpCollapsed    bool `json:"rulesHelpCollapsed"`
		SettingsHelpCollapsed bool `json:"settingsHelpCollapsed"`
		V4EasyMode            bool `json:"v4EasyMode"`
	} `json:"ui"`
}

type UpdateNetwork struct {
	Config      NetworkConfig `json:"config"`
	Description string        `json:"description"`
	RulesSource string        `json:"rulesSource"`
	Permissions struct {
		Zero0000000000000000000000000000000 struct {
			A bool `json:"a"`
			D bool `json:"d"`
			M bool `json:"m"`
			R bool `json:"r"`
		} `json:"00000000-0000-0000-0000-000000000000"`
	} `json:"permissions"`
	OwnerID            string `json:"ownerId"`
	CapabilitiesByName struct {
	} `json:"capabilitiesByName"`
	TagsByName struct {
	} `json:"tagsByName"`
}

type NetworkConfig struct {
	AuthTokens        string   `json:"authTokens"`
	CreationTime      int64    `json:"creationTime"`
	Capabilities      []string `json:"capabilities"`
	EnableBroadcast   bool     `json:"enableBroadcast"`
	ID                string   `json:"id"`
	IPAssignmentPools []struct {
		IPRangeStart string `json:"ipRangeStart"`
		IPRangeEnd   string `json:"ipRangeEnd"`
	} `json:"ipAssignmentPools"`
	LastModified      int64       `json:"lastModified"`
	Mtu               int         `json:"mtu"`
	MulticastLimit    int         `json:"multicastLimit"`
	Name              string      `json:"name"`
	Private           bool        `json:"private"`
	RemoteTraceLevel  int         `json:"remoteTraceLevel"`
	RemoteTraceTarget interface{} `json:"remoteTraceTarget"`
	Routes            []struct {
		Target string `json:"target"`
		Via    string `json:"via,omitempty"`
	} `json:"routes"`
	Rules []struct {
		EtherType int    `json:"etherType,omitempty"`
		Not       bool   `json:"not,omitempty"`
		Or        bool   `json:"or,omitempty"`
		Type      string `json:"type"`
		ID        int    `json:"id,omitempty"`
		Value     int    `json:"value,omitempty"`
	} `json:"rules"`
	Tags []struct {
		Default int `json:"default"`
		ID      int `json:"id"`
	} `json:"tags"`
	V4AssignMode struct {
		Zt bool `json:"zt"`
	} `json:"v4AssignMode"`
	V6AssignMode struct {
		SixPlane bool `json:"6plane"`
		Rfc4193  bool `json:"rfc4193"`
		Zt       bool `json:"zt"`
	} `json:"v6AssignMode"`
	DNS struct {
		Domain  string      `json:"domain"`
		Servers interface{} `json:"servers"`
	} `json:"dns"`
	SsoConfig struct {
		Enabled bool   `json:"enabled"`
		Mode    string `json:"mode"`
	} `json:"ssoConfig"`
}

type NetworkMember struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Clock        int64  `json:"clock"`
	NetworkID    string `json:"networkId"`
	NodeID       string `json:"nodeId"`
	ControllerID string `json:"controllerId"`
	Hidden       bool   `json:"hidden"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Config       struct {
		ActiveBridge         bool     `json:"activeBridge"`
		Address              string   `json:"address"`
		Authorized           bool     `json:"authorized"`
		Capabilities         []int    `json:"capabilities"`
		CreationTime         int64    `json:"creationTime"`
		ID                   string   `json:"id"`
		Identity             string   `json:"identity"`
		IPAssignments        []string `json:"ipAssignments"`
		LastAuthorizedTime   int64    `json:"lastAuthorizedTime"`
		LastDeauthorizedTime int      `json:"lastDeauthorizedTime"`
		NoAutoAssignIps      bool     `json:"noAutoAssignIps"`
		Nwid                 string   `json:"nwid"`
		Objtype              string   `json:"objtype"`
		RemoteTraceLevel     int      `json:"remoteTraceLevel"`
		RemoteTraceTarget    string   `json:"remoteTraceTarget"`
		Revision             int      `json:"revision"`
		Tags                 [][]int  `json:"tags"`
		VMajor               int      `json:"vMajor"`
		VMinor               int      `json:"vMinor"`
		VRev                 int      `json:"vRev"`
		VProto               int      `json:"vProto"`
		SsoExempt            bool     `json:"ssoExempt"`
	} `json:"config"`
	LastOnline          int64  `json:"lastOnline"`
	LastSeen            int64  `json:"lastSeen"`
	PhysicalAddress     string `json:"physicalAddress"`
	PhysicalLocation    string `json:"physicalLocation"`
	ClientVersion       string `json:"clientVersion"`
	ProtocolVersion     int    `json:"protocolVersion"`
	SupportsRulesEngine bool   `json:"supportsRulesEngine"`
}

type UpdateNetworkMember struct {
	Hidden      bool   `json:"hidden"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Config      struct {
		ActiveBridge    bool     `json:"activeBridge"`
		Authorized      bool     `json:"authorized"`
		Capabilities    []int    `json:"capabilities"`
		IPAssignments   []string `json:"ipAssignments"`
		NoAutoAssignIps bool     `json:"noAutoAssignIps"`
		Tags            [][]int  `json:"tags"`
	} `json:"config"`
}
