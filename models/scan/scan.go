package scan

import (
	"gopkg.in/mgo.v2/bson"
	"time"

	"fmt"
	"github.com/bearded-web/bearded/models/plugin"
	"github.com/bearded-web/bearded/pkg/pagination"
)

// useful time points
type Dates struct {
	Created  *time.Time `json:"created,omitempty"`
	Updated  *time.Time `json:"updated,omitempty"`
	Queued   *time.Time `json:"queued,omitempty"`
	Started  *time.Time `json:"started,omitempty"`
	Finished *time.Time `json:"finished,omitempty"`
}

type Session struct {
	Id     bson.ObjectId `json:"id"`
	Plugin bson.ObjectId `json:"plugin"`
	Status ScanStatus    `json:"status" description:"one of [created|queued|working|paused|finished|failed]"`
	Conf   *plugin.Conf  `json:"conf,omitempty"`

	// dates
	Dates `json:",inline"`
}

type ScanConf struct {
	Target string                 `json:"target"`
	Params map[string]interface{} `json:"params"`
}

type Scan struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Status ScanStatus    `json:"status,omitempty" description:"one of [created|queued|working|pause|finished|error]"`
	Conf   ScanConf      `json:"conf,omitempty"`

	//	Report    *report.Report `json:"report,omitempty" form:"-"`
	Sessions []*Session `json:"session,omitempty"`

	Plan    bson.ObjectId `json:"plan"`
	Owner   bson.ObjectId `json:"owner,omitempty"`
	Target  bson.ObjectId `json:"target"`
	Project bson.ObjectId `json:"project"`

	// dates
	Dates `json:",inline"`
}

type ScanList struct {
	pagination.Meta `json:",inline"`
	Results         []*Scan `json:"results"`
}

func (p *Scan) String() string {
	return fmt.Sprintf("%x [%s]", string(p.Id), p.Status)
}