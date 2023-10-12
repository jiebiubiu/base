package model

import "time"

type Users struct {
	ID        int64     `gorm:"column:id" json:"id"`
	Uuid      string    `gorm:"column:uuid" json:"uuid"`
	Account   string    `gorm:"column:account" json:"account"`
	Password  string    `gorm:"column:password" json:"password"`
	Salt      string    `gorm:"column:salt" json:"salt"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Users) TableName() string {
	return "users"
}

type Resource struct {
	ID        int64     `gorm:"column:id" json:"id"`
	Path      string    `gorm:"column:path" json:"path"`
	Filename  string    `gorm:"column:filename" json:"filename"`
	Md5       string    `gorm:"column:md5" json:"md5"`
	Key       string    `gorm:"column:key" json:"key"`
	Uuid      string    `gorm:"column:uuid" json:"uuid"`
	Type      string    `gorm:"column:type" json:"type"`
	Ext       string    `gorm:"column:ext" json:"ext"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Resource) TableName() string {
	return "resource"
}

type NoticeTask struct {
	ID          int64     `gorm:"column:id" json:"id"`
	NoticeType  string    `gorm:"column:notice_type" json:"notice_type"` //  email/wechat/feishu
	Status      string    `gorm:"column:status" json:"status"`           //  failed,succ,waiting
	FromAccount string    `gorm:"column:from_account" json:"from_account"`
	ToAccount   string    `gorm:"column:to_account" json:"to_account"`
	OpUid       int64     `gorm:"column:op_uid" json:"op_uid"`
	Msg         string    `gorm:"column:msg" json:"msg"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (NoticeTask) TableName() string {
	return "notice_task"
}

type NoticeRule struct {
	ID           int64     `gorm:"column:id" json:"id"`
	Uid          int64     `gorm:"column:uid" json:"uid"`                       //  创建用户
	NoticeTaskId int64     `gorm:"column:notice_task_id" json:"notice_task_id"` //  通知任务
	CycleType    string    `gorm:"column:cycle_type" json:"cycle_type"`         //  通知模式 only/every_day/every_week/every_year/work_day
	Rule         string    `gorm:"column:rule" json:"rule"`                     //  [{"year":[0],"month":[0],"week":[0],"day":[0],"hour":[0],"minute":[0]}]
	Status       string    `gorm:"column:status" json:"status"`                 //  created/uncreated 是否已经生成了执行计划
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (NoticeRule) TableName() string {
	return "notice_rule"
}

type NoticeRuleTaskRelation struct {
	ID           int64 `gorm:"column:id" json:"id"`
	Uid          int64 `gorm:"column:uid" json:"uid"`                       //  创建用户
	NoticeRuleId int64 `gorm:"column:notice_rule_id" json:"notice_rule_id"` //
	NoticeTaskId int64 `gorm:"column:notice_task_id" json:"notice_task_id"` //
}

func (NoticeRuleTaskRelation) TableName() string {
	return "notice_rule_task_relation"
}

type SendTask struct {
	ID           int64     `gorm:"column:id" json:"id"`
	NoticeTaskId int64     `gorm:"column:notice_task_id" json:"notice_task_id"` //
	SendTime     time.Time `gorm:"column:send_time" json:"send_time"`
	Status       string    `gorm:"column:status" json:"status"` //  waiting/succ/failed
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (SendTask) TableName() string {
	return "send_task"
}
