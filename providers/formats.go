package providers

// Deprecated
var (
	Zh_city_format_simple            = "%s市" // city_name city_suffix
	Zh_district_format_simple        = "%s区"
	Zh_building_number_format_simple = "%s座"
	Zh_postcode_format_simple        = "%s"
	Zh_street_name_format_simple     = "%s%s"     // city_name street_suffix
	Zh_street_address_format_simple  = "%s%s"     // street_name building_number
	Zh_address_format_simple         = "%s%s%s%s" // province city district street_address postcode
)

/*
	use {} wrap patterns, "RANDOM" stands for a random interger(string);

	func.`functionname` to call a specific function(without input) and get a return value

	func.`functionname`.[params] to call a specific function(without input) and get a return value

	functionname should be registered in global_function

	eg. func.Test.[name1,name2,name3]  nameN should be registered in global_variants

	"i_???" means `???` is a int

	will be completed :)
*/
var (
	Format_Zh_simple_address          = "{Zh_provinces_simple}{Zh_cities_simple}市{Zh_districts_simple}区{Zh_cities_simple}{Zh_street_suffixes_simple}{RANDOM}座"
	Format_Zh_simple_city             = "{Zh_cities_simple}市"
	Format_Zh_simple_district         = "{Zh_districts_simple}区"
	Format_Zh_simple_province         = "{Zh_provinces_simple}"
	Format_Zh_simple_street_name      = "{Zh_cities_simple}{Zh_street_suffixes_simple}"
	Format_Zh_simple_street_address   = "{Zh_cities_simple}{Zh_street_suffixes_simple}{RANDOM}座"
	Format_Zh_simple_job              = "{Job_zh}"
	Function_Zh_mobile                = `{func.GenerateMobileFromList.[%s]}`
	Format_nickname_male_pre          = "{Prefix_nickname_male}"
	Format_nickname_female_pre        = "{Prefix_nickname_female}"
	Format_nickname_common            = "{Common_nicknames}"
	Format_Zh_simple_male_firstname   = "{First_names_male_zh}"
	Format_Zh_simple_female_firstname = "{First_fenames_male_zh}"
	Format_Zh_simple_lastname         = "{Last_names_zh}"
	Format_Zh_simple_primary_school   = "{Zh_districts_simple}{Primary_school_suffix_zh}"
	Format_Zh_simple_middle_school    = "{Zh_districts_simple}{Middle_school_suffix_zh}"
	Format_Zh_simple_university       = "{Zh_districts_simple}{University_suffix_zh}"
	Format_chrome_user_agent_windows  = `Mozilla/5.0  ({Windows_platform_tokens}) AppleWebKit/{func.Saf} (KHTML, like Gecko) Chrome/{func.Chrome_version.[13,63]}.0.{func.Chrome_build_version.[800,899]}.0 Safari/{func.Saf}`
	Format_chrome_user_agent_linux    = `Mozilla/5.0  (X11; Linux {Linux_processors}) AppleWebKit/{func.Saf} (KHTML, like Gecko) Chrome/{func.Chrome_version.[13,63]}.0.{func.Chrome_build_version.[800,899]}.0 Safari/{func.Saf}`
	Format_chrome_user_agent_mac      = `Mozilla/5.0  (Macintosh; {Mac_processors}) Mac OS X 10_{func.Randn_with_min.[i_5,i_12]}_{func.Randn.[i_9]} AppleWebKit/{func.Saf} (KHTML, like Gecko) Chrome/{func.Chrome_version.[13,63]}.0.{func.Chrome_build_version.[800,899]}.0 Safari/{func.Saf}`
	Format_chrome_user_agent_android  = `Mozilla/5.0  (Linux; Android {Android_versions}) AppleWebKit/{func.Saf} (KHTML, like Gecko) Chrome/{func.Chrome_version.[13,63]}.0.{func.Chrome_build_version.[800,899]}.0 Safari/{func.Saf}`
	Format_chrome_user_agent_ios      = `Mozilla/5.0  ({Ios_versions}) AppleWebKit/{func.Saf} (KHTML, like Gecko)  CriOS/{func.Chrome_version.[13,63]}.0.{func.Chrome_build_version.[800,899]}.0 Safari/{func.Saf}`
	Format_ie_user_agent              = `Mozilla/5.0 (compatible; MSIE {func.Randn_with_min.[i_5,i_9]}.0; {Windows_platform_tokens} Trident/{func.Randn_with_min.[i_3,i_5]}.{func.Randn.[i_2]})`
)
