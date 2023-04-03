CREATE TABLE `job_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `encryptBossId` varchar(255) NOT NULL COMMENT '加密bossId',
  `bossName` varchar(255) NOT NULL COMMENT 'boss name',
  `bossTitle` varchar(255) NOT NULL COMMENT 'boss title',
  `encryptJobId` varchar(255) NOT NULL COMMENT '加密jobId',
  `jobName` varchar(255) NOT NULL COMMENT 'job name',
  `salaryDesc` varchar(255) NOT NULL COMMENT '薪资范围',
  `jobExperience` varchar(255) NOT NULL COMMENT '工作经验范围',
  `jobDegree` varchar(255) NOT NULL COMMENT '学历',
  `cityName` varchar(255) NOT NULL COMMENT '城市名称',
  `areaDistrict` varchar(255) NOT NULL COMMENT '城市区域',
  `businessDistrict` varchar(255) NOT NULL COMMENT '商业区',
  `city` int(11) NOT NULL COMMENT 'city code',
  `encryptBrandId` varchar(255) NOT NULL COMMENT '加密公司ID',
  `brandName` varchar(255) NOT NULL COMMENT '公司名称',
  `brandStageName` varchar(255) NOT NULL COMMENT '融资阶段',
  `brandIndustry` varchar(255) NOT NULL COMMENT '公司所属产业',
  `brandScaleName` varchar(255) NOT NULL COMMENT '公司规模',
  `skills` varchar(255) NOT NULL COMMENT '技能标签',
  `welfare_list` varchar(255) NOT NULL COMMENT '公司福利标签',
  `lid` varchar(255) NOT NULL COMMENT 'lid',
  `security_id` varchar(255) NOT NULL COMMENT 'security_id',
  `create_time` int(10) unsigned DEFAULT '0',
  `update_time` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='job列表';