package domain

const (
	queryInsertSection    = `insert into sections (title, desc, parent, access_level, created_at, updated_at) values ($1,$2,$3,$4,$5,$6) returning id`
	querySelectAllSection = `SELECT r.id, r.first_name, r.last_name, r.email, r.phone, r.room_id, r.status, r.start_date, r.end_date, r.created_at, r.updated_at, rt.id, rt.title, r.processed
										FROM reservations r 
										left join room_types rt 
										on (r.room_id = rt.id) 
										order by r.start_date desc`
	queryGetSectionById = `SELECT * FROM sections WHERE id = $1`
	// querySearchAvailability        = `SELECT count(id) FROM room_allotments WHERE room_no_id = $1 AND $2 < end_date AND $3 > start_date`
	// querySearchAvailabilityAllRoom = `SELECT r.id, r.roomtype_id, r.room_no FROM rooms r WHERE r.id not in (SELECT ra.room_no_id FROM room_allotments ra WHERE $1 < ra.end_date AND $2 > ra.start_date)`
)

var SectionService sectionDomainInterface = &Category{}

type Section section
type sectionDomainInterface interface {
	Create() (int, error)
}

func (s *Section) Create() (int, error) {
	return 0, nil
}
