package models

type ProjectType int

const (
	WebDev ProjectType =  iota
	MailingMarketing
	ADVCampaign
)


type StepStatus int

const(
	TODO StepStatus = iota
	DONE
)

type Steps struct{
	content string
	status StepStatus
}

type ProjectAttribtes struct{
	projectType ProjectType
	description string
	steps []Steps
}