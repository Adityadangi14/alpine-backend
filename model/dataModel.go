package model

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	NEWSID           string `json:"NEWSID"`
	SCRIPCD          int    `json:"SCRIP_CD"`
	XMLNAME          string `json:"XML_NAME"`
	NEWSSUB          string `json:"NEWSSUB"`
	DTTM             string `json:"DT_TM"`
	NEWSDT           string `json:"NEWS_DT"`
	CRITICALNEWS     int    `json:"CRITICALNEWS"`
	ANNOUNCEMENTTYPE string `json:"ANNOUNCEMENT_TYPE"`
	FILESTATUS       string `json:"FILESTATUS"`
	ATTACHMENTNAME   string `json:"ATTACHMENTNAME"`
	MORE             string `json:"MORE"`
	HEADLINE         string `json:"HEADLINE"`
	CATEGORYNAME     string `json:"CATEGORYNAME"`
	OLD              int    `json:"OLD"`
	RN               int    `json:"RN"`
	PDFFLAG          int    `json:"PDFFLAG"`
	NSURL            string `json:"NSURL"`
	SLONGNAME        string `json:"SLONGNAME"`
	AGENDAID         int    `json:"AGENDA_ID"`
	TotalPageCnt     int    `json:"TotalPageCnt"`
	NewsSubmissionDt string `json:"News_submission_dt"`
	DissemDT         string `json:"DissemDT"`
	TimeDiff         string `json:"TimeDiff"`
	FldAttachsize    int    `json:"Fld_Attachsize"`
	SUBCATNAME       string `json:"SUBCATNAME"`
}
