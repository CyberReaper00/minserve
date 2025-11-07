package humain

import (
    "os"
    "fmt"
    "log"
	"math"
    "bufio"
    "strings"
    "strconv"
	"reflect"
)

// Print accepts a string and prints the specified string without a newline
func Print(msg ...any) {
	if len(msg) == 0 { msg = append(msg, "") }
	for _, item := range msg { fmt.Printf("%v\n", item) }
}

// Print_List accepts a []string and prints all items present in an ascii box,
// seperated by bars
func Print_List(msg string, list []string) {
	Print("", msg)

	var list_len int
	for _, item := range list { list_len += len(item)+3 }
	fmt.Printf("\t┏%s┓\n", strings.Repeat("━", list_len-1)); fmt.Print("\t┃ ")
	for _, item := range list { fmt.Print(item + " ┃ ") }
	Print(); fmt.Printf("\t┗%s┛\n", strings.Repeat("━", list_len-1))
	Print()
}

// GetType takes in a value as an any type and returns its actual type as a string
func GetType(value any) string {
	switch value.(type) {
		case int:	  return "int"
		case bool:	  return "bool"
		case float32: return "float32"
		case float64: return "float64"
		default: return "string"
	}
}

// CheckType checks the given value for the specified type and returns a bool
func CheckType(value any, check_type string) bool {
	switch check_type {
		case "int":		return reflect.TypeOf(value).Kind() == reflect.Int
		case "float32":	return reflect.TypeOf(value).Kind() == reflect.Float32
		case "float64":	return reflect.TypeOf(value).Kind() == reflect.Float64
		case "string":	return reflect.TypeOf(value).Kind() == reflect.String
	}

	// unreachable
	return false
}

// TrueType takes in any type value and converts it into its proper type and returns
// the converted value as an any type and the type itself as a string value which
// can be used for comparisons
func TrueType(data any) (any, string) {

	// [EDGE CASE] the value entered might be string and has no need to be checked
	// seperately so it is returned immediately if string, otherwise it moves onto be
	// parsed for a different type
	str, ok := data.(string);
	if ok { return str, "string" }

	// [EDGE CASE] t and f are considered as bool by ParseBool and the user might
	// want the value t or f instead of getting a bool, so they are returned immediately
	// as strings
	if strings.ToLower(str) == "f" || strings.ToLower(str) == "t" { return data, "string" }

	var converted_value any
	switch data.(type) {
		case int:  converted_value = data
		case bool: converted_value = data

		case float64:
			if data.(float64) - math.Trunc(data.(float64)) == 0 {
				converted_value = int(data.(float64))
			} else { converted_value = data }

		case float32:
			f64 := data.(float64)
			if f64 - math.Trunc(f64) == 0 { converted_value = int(f64)
			} else { converted_value = data }
	}

	got_type := GetType(converted_value)
	return converted_value, got_type
}

// Input is used to take in a value from the user on the commandline and then
// converts the string value into its actual type and returns that value as an
// any type
func Input(msg string, val ...any) any {
	fmt.Printf(msg + ": ", val...)

	reader 	 := bufio.NewReader(os.Stdin)
	data, _	 := reader.ReadString('\n')
	data1	 := strings.TrimSpace(data)
	input, _ := TrueType(data1)
	return input
}

const top = "╭──────────────────────────────────────────────────╮"
const mid = "├──────────────────────────────────────────────────┤"
const bot = "╰──────────────────────────────────────────────────╯"
// InputMenu takes in an int slice and an arbitrary amount of messages that 
func InputMenu(lc_msgs []int, msgs ...string) []any {

	if len(msgs) < 1 { log.Fatalln("Error: No arguments were provided for the menu") }

	fmt.Println("\033[H\033[2J")
	final_slice := make( []any, len(msgs) )

	for i, msg := range msgs {

		// Input handling
		if i == 0 { fmt.Println(top)
		} else if i > 0 && i < len(msgs) { fmt.Println(mid) }
		user_input := Input("│ %s", msg)

		// Output handling
		if IntSliceContains(lc_msgs, i + 1, "exact") || IntSliceContains(lc_msgs, -1, "exact") &&
			reflect.TypeOf(user_input).Kind() == reflect.String {

			str_inp := user_input.(string)
			final_slice[i] = strings.ToLower(str_inp)

		} else if IntSliceContains(lc_msgs, 0, "exact") { final_slice[i] = user_input }
	}

	fmt.Println(bot)
	return final_slice
}

func PauseExit() {
    xyz := Input("\n\033[1;42m Press Enter to exit... \033[0m\n")
    if xyz != "" { PauseExit() }
    return
}

func Err(msg string, err error, val ...any) {
    if err != nil { log.Fatalf( fmt.Sprintf(msg, val...) )}
}

func PrettyErr(msg string, err error, fg bool, val ...any) {
    var scheme int
    if fg { scheme = 3
    } else { scheme = 4 }

    if err != nil {
		log.Fatalf(fmt.Sprintf("\033[1;%d1m", scheme) +
				fmt.Sprintf(msg, val...) +
				"\033[0m\n")
	}
}

func PrettyMsg(msg string, color string, fg bool, val ...any) {
    var scheme int
    var color_code string

    if fg { scheme = 3
    } else { scheme = 4 }

    if !strings.Contains(color, "#") {
		switch color {
			case "black": 	color_code = fmt.Sprintf("\033[1;%d0m", scheme)
			case "red": 	color_code = fmt.Sprintf("\033[1;%d1m", scheme)
			case "green": 	color_code = fmt.Sprintf("\033[1;%d2m", scheme)
			case "orange": 	color_code = fmt.Sprintf("\033[1;%d3m", scheme)
			case "blue": 	color_code = fmt.Sprintf("\033[1;%d4m", scheme)
			case "purple": 	color_code = fmt.Sprintf("\033[1;%d5m", scheme)
			case "teal": 	color_code = fmt.Sprintf("\033[1;%d6m", scheme)
			case "beige": 	color_code = fmt.Sprintf("\033[1;%d7m", scheme)
			default: log.Fatalln(fmt.Sprintf("\033[1;41 Invalid color '%s' given\033[0m", color))
		}

    } else {
		parts := strings.Split(color, "#")
		if len(parts) < 2 { log.Fatalln("Invalid input was given") }
		hex := parts[len(parts) - 1]

		if len(hex) != 6 { log.Fatalln("Invalid hex code was entered") }

		r64, errR := strconv.ParseInt(hex[0:2], 16, 0)
		g64, errG := strconv.ParseInt(hex[2:4], 16, 0)
		b64, errB := strconv.ParseInt(hex[4:6], 16, 0)

		if errR != nil || errG != nil || errB != nil {
			log.Fatalln("" +
			"\033[1;32mParsing Error:\n" +
			"Unknown value was encountered when parsing hex code\033[0m")
		}

		r := int(r64)
		g := int(g64)
		b := int(b64)

		color_code = fmt.Sprintf("\033[%d8;2;%d;%d;%dm", scheme, r, g, b)
    }

    fmt.Printf(color_code+msg+"\033[0m\n", val...)
}

// TODO: complete Int_Contains function
func Int_Contains(num int) {
	e := 1

	for counter := 0; counter < 2; {
		num2 := num%int(math.Pow10(e))
		if num2 == num { counter++ }
		e++
	}

	e = e-2
	e2 := e
	num_list := make([]int, e)
	for i := 0; e > 0; i++ {
		num_list[i] = num%int(math.Pow10(e))
		e--
	}

	subt_list := make([]int, e2)
	slli := len(subt_list) - 1
	nlli := len(num_list) - 1
	for i := 0; i < nlli; i++ {
		if i == 0 { subt_list[slli] = num_list[nlli] }
		subt_list[i] = num_list[i] - num_list[i+1]
	}


	ind_num_list := make([]int, e2)
	for i := 0; i < e2; i++ {
		ind_num_list[i] = subt_list[i]/int(math.Pow10(e2 - (i+1)))
	}

	e3 := e2-1
	i := 0
	k := 0
	j := e3
	final_list := make([]int, e3)
	for ; i < e3; i++ {
		for l := 0; l < e3; l++ {
			final_list[l] = ind_num_list[l]*int(math.Pow10(k+1)) + ind_num_list[l+1]
		}
		k++ 
		j--
	}
}

func IntSliceContains(slice []int, target int, method string) bool {
	if method == "exact" {
		for _, num := range slice {
			if num == target { return true }
		}

	} else if method == "fuzzy" {
		// TODO: add Int_Contains
		for _, num := range slice {
			str := strconv.Itoa(num)
			target_str := strconv.Itoa(target)

			if strings.Contains(str, target_str) { return true }
		}
	} else { panic("Invalid method was provided in IntSliceContains") }

    return false
}

func StrSliceContains(slice []string, target string, method string) bool {
	if method == "exact" {
		for _, str := range slice {
			if str == target { return true }
		}

	} else if method == "fuzzy" {
		for _, str := range slice {
			if strings.Contains(str, target) { return true }
		}

	} else { panic("Invalid method was provided in StrSliceContains") }

    return false
}
