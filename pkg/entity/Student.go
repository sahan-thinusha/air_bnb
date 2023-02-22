package entity

type Student struct {
	Id              string `json:"id" bson:"_id"`
	Base            `bson:",inline"`
	StudentUsername string `json:"student_username"`
	Name            struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"name"`
	Demographics struct {
		Gender                  string `json:"gender"`
		BirthDate               string `json:"birth_date"`
		ProjectedGraduationYear int    `json:"projected_graduation_year"`
	} `json:"demographics"`
	Addresses struct {
		Physical struct {
			Street        string `json:"street"`
			City          string `json:"city"`
			StateProvince string `json:"state_province"`
			PostalCode    int    `json:"postal_code"`
		} `json:"physical"`
		Mailing struct {
			Street        string `json:"street"`
			City          string `json:"city"`
			StateProvince string `json:"state_province"`
			PostalCode    int    `json:"postal_code"`
		} `json:"mailing"`
	} `json:"addresses"`
	Alerts struct {
		Legal struct {
			Description string `json:"description"`
			ExpiresDate string `json:"expires_date"`
		} `json:"legal"`
		Discipline struct {
			Description string `json:"description"`
			ExpiresDate string `json:"expires_date"`
		} `json:"discipline"`
		Medical struct {
			Description string `json:"description"`
			ExpiresDate string `json:"expires_date"`
		} `json:"medical"`
		Other struct {
			Description string `json:"description"`
			ExpiresDate string `json:"expires_date"`
		} `json:"other"`
	} `json:"alerts"`
	Phones struct {
		Main struct {
			Number string `json:"number"`
		} `json:"main"`
	} `json:"phones"`
	SchoolEnrollment struct {
		EnrollStatus            string `json:"enroll_status"`
		EnrollStatusDescription string `json:"enroll_status_description"`
		EnrollStatusCode        int    `json:"enroll_status_code"`
		GradeLevel              int    `json:"grade_level"`
		EntryDate               string `json:"entry_date"`
		ExitDate                string `json:"exit_date"`
		SchoolNumber            int    `json:"school_number"`
		SchoolId                int    `json:"school_id"`
		EntryCode               string `json:"entry_code"`
		EntryComment            string `json:"entry_comment"`
		Track                   string `json:"track"`
		FullTimeEquivalency     struct {
			Fteid int    `json:"fteid"`
			Name  string `json:"name"`
		} `json:"full_time_equivalency"`
	} `json:"school_enrollment"`
	EthnicityRace struct {
		SchedulingReportingEthnicity string `json:"scheduling_reporting_ethnicity"`
	} `json:"ethnicity_race"`
	Contact struct {
		EmergencyContactName1 string `json:"emergency_contact_name1"`
		EmergencyPhone1       string `json:"emergency_phone1"`
		EmergencyPhone2       string `json:"emergency_phone2"`
		GuardianEmail         string `json:"guardian_email"`
		Mother                string `json:"mother"`
		Father                string `json:"father"`
		DoctorName            string `json:"doctor_name"`
		DoctorPhone           string `json:"doctor_phone"`
	} `json:"contact"`
	ContactInfo struct {
		Email string `json:"email"`
	} `json:"contact_info"`
}
