package persist

type JobList struct {
	Id               int    `gorm:"id" json:"id"`                             // 自增ID
	Encryptbossid    string `gorm:"encryptBossId" json:"encryptBossId"`       // 加密bossId
	Bossname         string `gorm:"bossName" json:"bossName"`                 // boss name
	Bosstitle        string `gorm:"bossTitle" json:"bossTitle"`               // boss title
	Encryptjobid     string `gorm:"encryptJobId" json:"encryptJobId"`         // 加密jobId
	Jobname          string `gorm:"jobName" json:"jobName"`                   // job name
	Salarydesc       string `gorm:"salaryDesc" json:"salaryDesc"`             // 薪资范围
	Jobexperience    string `gorm:"jobExperience" json:"jobExperience"`       // 工作经验范围
	Jobdegree        string `gorm:"jobDegree" json:"jobDegree"`               // 学历
	Cityname         string `gorm:"cityName" json:"cityName"`                 // 城市名称
	Areadistrict     string `gorm:"areaDistrict" json:"areaDistrict"`         // 城市区域
	Businessdistrict string `gorm:"businessDistrict" json:"businessDistrict"` // 商业区
	City             int    `gorm:"city" json:"city"`                         // city code
	Encryptbrandid   string `gorm:"encryptBrandId" json:"encryptBrandId"`     // 加密公司ID
	Brandname        string `gorm:"brandName" json:"brandName"`               // 公司名称
	Brandstagename   string `gorm:"brandStageName" json:"brandStageName"`     // 融资阶段
	Brandindustry    string `gorm:"brandIndustry" json:"brandIndustry"`       // 公司所属产业
	Brandscalename   string `gorm:"brandScaleName" json:"brandScaleName"`     // 公司规模
	Skills           string `gorm:"skills" json:"skills"`                     // 技能标签
	WelfareList      string `gorm:"welfare_list" json:"welfare_list"`         // 公司福利标签
	CreateTime       uint   `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime       uint   `gorm:"autoCreateTime" json:"update_time"`
}

type JobTags struct {
	Id           int    `gorm:"id" json:"id"`                     // 自增ID
	Encryptjobid string `gorm:"encryptJobId" json:"encryptJobId"` // 加密jobId
	Type         string `gorm:"type" json:"type"`                 // 标签类型 skills experience welfare
	Name         string `gorm:"name" json:"name"`                 // 标签名称
	CreateTime   uint   `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime   uint   `gorm:"autoCreateTime" json:"update_time"`
}
