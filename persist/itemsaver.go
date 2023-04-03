package persist

import (
	"errors"
	"log"
	"strings"

	"crawler_jobs/engine"
	"crawler_jobs/global"
	"crawler_jobs/parser/job"

	"gorm.io/gorm"
)

func SaveJobList() (chan engine.Item, error) {
	out := make(chan engine.Item)
	go func() {
		for item := range out {
			err := saveJobList(item.Data.(job.JobList))
			if err != nil {
				log.Printf("saveJobList err=%v", err)
			}
		}
	}()
	return out, nil
}

func saveJobList(data job.JobList) (err error) {
	jobList := new(JobList)
	result := global.DB.Select("id").Where("encryptJobId = ?", data.EncryptJobID).First(jobList)
	if result.Error != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if result.RowsAffected > 0 {
		log.Printf("saveJobList exist job_id=%v", data.EncryptJobID)
		return nil
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		jobListData := &JobList{
			Encryptbossid:    data.EncryptBossID,
			Bossname:         data.BossName,
			Bosstitle:        data.BossTitle,
			Encryptjobid:     data.EncryptJobID,
			Jobname:          data.JobName,
			Salarydesc:       data.SalaryDesc,
			Jobexperience:    data.JobExperience,
			Jobdegree:        data.JobDegree,
			Cityname:         data.CityName,
			Areadistrict:     data.AreaDistrict,
			Businessdistrict: data.BusinessDistrict,
			City:             data.City,
			Encryptbrandid:   data.EncryptBrandID,
			Brandname:        data.BrandName,
			Brandstagename:   data.BrandStageName,
			Brandindustry:    data.BrandIndustry,
			Brandscalename:   data.BrandScaleName,
			Skills:           strings.Join(data.Skills, ","),
			WelfareList:      strings.Join(data.WelfareList, ","),
			Lid:              data.Lid,
			SecurityID:       data.SecurityID,
		}
		if err := tx.Create(jobListData).Error; err != nil {
			return err
		}
		if len(data.Skills) > 0 {
			skills := make([]JobTags, 0)
			for _, v := range data.Skills {
				skills = append(skills, JobTags{
					Encryptjobid: data.EncryptJobID,
					Type:         "skills",
					Name:         v,
				})
			}
			if err := tx.CreateInBatches(skills, len(skills)).Error; err != nil {
				return err
			}
		}
		if len(data.WelfareList) > 0 {
			welfareList := make([]JobTags, 0)
			for _, v := range data.WelfareList {
				welfareList = append(welfareList, JobTags{
					Encryptjobid: data.EncryptJobID,
					Type:         "welfare",
					Name:         v,
				})
			}
			if err := tx.CreateInBatches(welfareList, len(welfareList)).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
