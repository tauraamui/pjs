package models

import "gorm.io/gorm"

type ProjectWithEntries struct {
	Project Project
	Entries []Entry
}

func (pe *ProjectWithEntries) CreateProjectWithEntries(pKey int, db *gorm.DB) {
	db.Where("id = ?", pKey).First(pe.Project)
	db.Where("project_id = ?", pe.Project.ID).Find(pe.Entries)
}

// leverage Entry functions instead of re-writing
