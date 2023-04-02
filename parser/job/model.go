package job

type JobListData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	ZpData  ZpData `json:"zpData"`
}
type JobList struct {
	SecurityID       string        `json:"securityId"`
	BossAvatar       string        `json:"bossAvatar"`
	BossCert         int           `json:"bossCert"`
	EncryptBossID    string        `json:"encryptBossId"`
	BossName         string        `json:"bossName"`
	BossTitle        string        `json:"bossTitle"`
	GoldHunter       int           `json:"goldHunter"`
	BossOnline       bool          `json:"bossOnline"`
	EncryptJobID     string        `json:"encryptJobId"`
	ExpectID         int           `json:"expectId"`
	JobName          string        `json:"jobName"`
	Lid              string        `json:"lid"`
	SalaryDesc       string        `json:"salaryDesc"`
	JobLabels        []string      `json:"jobLabels"`
	JobValidStatus   int           `json:"jobValidStatus"`
	IconWord         string        `json:"iconWord"`
	Skills           []string      `json:"skills"`
	JobExperience    string        `json:"jobExperience"`
	DaysPerWeekDesc  string        `json:"daysPerWeekDesc"`
	LeastMonthDesc   string        `json:"leastMonthDesc"`
	JobDegree        string        `json:"jobDegree"`
	CityName         string        `json:"cityName"`
	AreaDistrict     string        `json:"areaDistrict"`
	BusinessDistrict string        `json:"businessDistrict"`
	JobType          int           `json:"jobType"`
	ProxyJob         int           `json:"proxyJob"`
	ProxyType        int           `json:"proxyType"`
	Anonymous        int           `json:"anonymous"`
	Outland          int           `json:"outland"`
	Optimal          int           `json:"optimal"`
	IconFlagList     []interface{} `json:"iconFlagList"`
	ItemID           int           `json:"itemId"`
	City             int           `json:"city"`
	IsShield         int           `json:"isShield"`
	AtsDirectPost    bool          `json:"atsDirectPost"`
	Gps              interface{}   `json:"gps"`
	EncryptBrandID   string        `json:"encryptBrandId"`
	BrandName        string        `json:"brandName"`
	BrandLogo        string        `json:"brandLogo"`
	BrandStageName   string        `json:"brandStageName"`
	BrandIndustry    string        `json:"brandIndustry"`
	BrandScaleName   string        `json:"brandScaleName"`
	WelfareList      []string      `json:"welfareList"`
	Industry         int           `json:"industry"`
	Contact          bool          `json:"contact"`
}
type ZpData struct {
	ResCount     int         `json:"resCount"`
	FilterString string      `json:"filterString"`
	Lid          string      `json:"lid"`
	HasMore      bool        `json:"hasMore"`
	JobList      []JobList   `json:"jobList"`
	TotalCount   int         `json:"totalCount"`
	BrandCard    interface{} `json:"brandCard"`
}
