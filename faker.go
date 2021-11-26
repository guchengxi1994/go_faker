package gofaker

import "github.com/guchengxi1994/go_faker/utils"

/*
	address generator

	company(name) generator

	experience_generator

	ipaddress generator

	job generator

	lorem generator

	mobile generator

	nickname generator

	personname generator

	profile generator

	school generator

	ssn generator

*/
type Faker struct {
	Locale           string
	Args             []float64 // weight
	cachedGenerators map[string]interface{}
	inited           bool
	Gender           bool
}

type Generators struct {
	Gpname   *PersonName
	Gssn     *SSN
	Gmobile  *Mobile
	Gaddress *Address
}

func (f *Faker) Init() {
	f.inited = true
	f.cachedGenerators = make(map[string]interface{})
}

// generate a random name with gender(bool)
func (f *Faker) PersonName() string {
	if !f.inited {
		f.Init()
	}
	value, ok := f.cachedGenerators["personname"]
	if !ok {
		_value := PersonName{
			Locale: f.Locale,
			Gender: f.Gender,
		}
		f.cachedGenerators["personname"] = _value
		_value.Generate(f.Args...)
		return _value.ToString(true)
	} else {
		_value := value.(PersonName)
		return _value.ToString(true)
	}
}

/*
	generate a random profile

	useCache(true) if related generators are used
*/
func (f *Faker) Profile(useCache bool) string {
	if !f.inited {
		f.Init()
	}
	if useCache {
		gs := Generators{}
		value, ok := f.cachedGenerators["personname"]

		if ok {
			_gpname := value.(PersonName)
			gs.Gpname = &_gpname
		}

		value, ok = f.cachedGenerators["ssn"]
		if ok {
			_gssn := value.(SSN)
			gs.Gssn = &_gssn
		}

		value, ok = f.cachedGenerators["mobile"]
		if ok {
			_gmobile := value.(Mobile)
			gs.Gmobile = &_gmobile
		}

		value, ok = f.cachedGenerators["address"]
		if ok {
			_gAddress := value.(Address)
			gs.Gaddress = &_gAddress
		}

		value, ok = f.cachedGenerators["profile"]
		if ok {
			// value.(Profile).Generate(&gs)
			_value := value.(Profile)
			_value.Generate(&gs)
			return value.(Profile).ToString()
		} else {
			_value := Profile{
				Locale: f.Locale,
				Gender: f.Gender,
			}
			f.cachedGenerators["profile"] = _value
			_value.Generate(&gs)
			return _value.ToString()
		}
	} else {
		_value := Profile{
			Locale: f.Locale,
			Gender: f.Gender,
		}
		_value.Generate(nil)
		return _value.ToString()
	}

}

// generate a random company name
func (f *Faker) CompanyName() string {
	if !f.inited {
		f.Init()
	}
	value, ok := f.cachedGenerators["company"]
	if !ok {
		_value := Company{
			Locale: f.Locale,
			Args:   f.Args,
		}
		f.cachedGenerators["company"] = _value
		return _value.Generate()
	} else {
		_value := value.(Company)
		return _value.Generate()
	}
}

// generate a random ip address
func (f *Faker) IpAddr(withport bool) string {
	if !f.inited {
		f.Init()
	}
	if withport {
		return GenerateIpV4Address(false) + ":" + GeneratePort(60000)
	}
	return GenerateIpV4Address(false)
}

// generate a random job
func (f *Faker) Job() string {
	if !f.inited {
		f.Init()
	}
	value, ok := f.cachedGenerators["job"]
	if !ok {
		_value := Job{
			Locale: f.Locale,
		}
		f.cachedGenerators["job"] = _value
		return _value.Generate()
	} else {
		_value := value.(Job)
		return _value.Generate()
	}
}

// generate a random lorem, uclike is just for fun
func (f *Faker) Lorem(uclike bool) string {
	if !f.inited {
		f.Init()
	}
	value, ok := f.cachedGenerators["lorem"]
	if !ok {
		_value := Lorem{
			Locale:    f.Locale,
			Type:      utils.Randn(2),
			LikeUC:    uclike,
			MaxLength: utils.Randn(10) + 10,
		}
		f.cachedGenerators["lorem"] = _value
		_value.Generate()
		return _value.ToString()
	} else {
		_value := value.(Lorem)
		_value.Generate()
		return _value.ToString()
	}
}

// generate a random mobile
func (f *Faker) Mobile() string {
	if !f.inited {
		f.Init()
	}
	value, ok := f.cachedGenerators["mobile"]
	if !ok {
		_value := Mobile{
			Locale: f.Locale,
		}
		f.cachedGenerators["mobile"] = _value
		return _value.Generate()
	} else {
		_value := value.(Mobile)
		return _value.Generate()
	}
}

// generate a random nickname
func (f *Faker) Nickname() string {
	if !f.inited {
		f.Init()
	}
	value, ok := f.cachedGenerators["nickname"]
	if !ok {
		_value := Nickname{
			Gender: f.Gender,
		}
		f.cachedGenerators["nickname"] = _value
		_value.Generate()
		return _value.ToString()
	} else {
		_value := value.(Nickname)
		_value.Generate()
		return _value.ToString()
	}
}

// generate a random school name
func (f *Faker) School() string {
	if !f.inited {
		f.Init()
	}
	value, ok := f.cachedGenerators["school"]
	if !ok {
		_value := School{
			Locale: f.Locale,
			Type:   utils.Randn(4),
		}
		f.cachedGenerators["school"] = _value
		return _value.Generate()
	} else {
		_value := value.(School)
		return _value.Generate()
	}
}

func (f *Faker) SSN() string {
	if !f.inited {
		f.Init()
	}
	value, ok := f.cachedGenerators["ssn"]
	if !ok {
		_value := SSN{
			Locale: f.Locale,
			Gender: f.Gender,
		}
		f.cachedGenerators["ssn"] = _value
		return _value.Generate()
	} else {
		_value := value.(SSN)
		return _value.Generate()
	}
}
