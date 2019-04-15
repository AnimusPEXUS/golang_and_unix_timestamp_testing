package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gopherjs/jquery"
	"honnef.co/go/js/dom"
)

func GenPage() dom.Element {

	doc := dom.GetWindow().Document()

	document_div := doc.CreateElement("div")
	ret := document_div

	control_div := doc.CreateElement("div")
	control_rfc_div := doc.CreateElement("div")
	control_unix_div := doc.CreateElement("div")

	use_current_time_button := doc.CreateElement("button")
	use_current_time_button.SetTextContent("Now!")

	use_current_time_utc_button := doc.CreateElement("button")
	use_current_time_utc_button.SetTextContent("Now! UTC")

	use_custom_RFC_time_button := doc.CreateElement("button")
	use_custom_RFC_time_button.SetTextContent("Use This RFC")

	use_custom_Unix_UTC0_time_button := doc.CreateElement("button")
	use_custom_Unix_UTC0_time_button.SetTextContent("Use This Unix UTC0")

	use_custom_Unix_Local_time_button := doc.CreateElement("button")
	use_custom_Unix_Local_time_button.SetTextContent("Use This Unix Local")

	time_shift_div := doc.CreateElement("div")
	time_shift2_div := doc.CreateElement("div")

	// zone_shift_div := doc.CreateElement("div")

	span0 := doc.CreateElement("span")
	span0.SetTextContent("+")
	span3 := doc.CreateElement("span")
	span3.SetTextContent("+")

	span1 := doc.CreateElement("span")
	span1.SetTextContent("Go Minutes")
	span2 := doc.CreateElement("span")
	span2.SetTextContent("Unix Minutes")

	plus_time_input := doc.CreateElement("input")
	plus_time2_input := doc.CreateElement("input")
	rfc_time_input := doc.CreateElement("input")
	unix_time_input := doc.CreateElement("input")

	local_output := doc.CreateElement("div")

	utc_output := doc.CreateElement("div")

	local_output_alt := doc.CreateElement("div")

	utc_output_alt := doc.CreateElement("div")

	time_shift_div.AppendChild(span0)
	time_shift_div.AppendChild(plus_time_input)
	time_shift_div.AppendChild(span1)

	time_shift2_div.AppendChild(span3)
	time_shift2_div.AppendChild(plus_time2_input)
	time_shift2_div.AppendChild(span2)

	// zone_shift_div.AppendChild(zone_shift1_input)

	control_div.AppendChild(use_current_time_button)
	control_div.AppendChild(use_current_time_utc_button)
	// control_div.AppendChild(zone_shift0_input)

	control_rfc_div.AppendChild(use_custom_RFC_time_button)
	control_rfc_div.AppendChild(rfc_time_input)

	control_unix_div.AppendChild(use_custom_Unix_UTC0_time_button)
	control_unix_div.AppendChild(use_custom_Unix_Local_time_button)
	control_unix_div.AppendChild(unix_time_input)

	for _, i := range []dom.Element{
		control_div,
		control_rfc_div,
		control_unix_div,
		time_shift_div,
		time_shift2_div,
		// zone_shift_div,
		local_output,
		utc_output,
		local_output_alt,
		utc_output_alt,
	} {
		document_div.AppendChild(i)
	}

	func0 := func(mode int, already_utc0 bool) {

		var t, tu time.Time

		switch mode {
		case 0:
			t = time.Now()
		case 1:
			t = time.Now().UTC()
		case 2:
			var err = error(nil)
			t, err = time.Parse(time.RFC3339, jquery.NewJQuery(rfc_time_input).Val())
			if err != nil {
				log.Print(err)
				return
			}
		case 3:
			utime, _ := strconv.Atoi(jquery.NewJQuery(unix_time_input).Val())
			t = time.Unix(int64(utime), 0)
		default:
			fmt.Println("error")
		}

		if already_utc0 {
			tu = t
		} else {
			tu = t.UTC()
		}

		f := func(label string, t time.Time) string {
			return fmt.Sprintf("%s: Time: %s, Unix: %v", label, t.Format(time.RFC3339), t.Unix())
		}

		alt_num := 0
		alt_num2 := 0

		{
			alt_val := jquery.NewJQuery(plus_time_input).Val()
			if alt_val != "" {
				alt_num, _ = strconv.Atoi(alt_val)
			}
		}

		{
			alt_val := jquery.NewJQuery(plus_time2_input).Val()
			if alt_val != "" {
				alt_num2, _ = strconv.Atoi(alt_val)
			}
		}

		var t_alt time.Time

		t_alt = t.Add(time.Duration(time.Duration(alt_num) * time.Minute))

		t_alt = time.Unix(t_alt.Unix()+int64(alt_num2*60), 0).In(t.Location())

		tu_alt := t_alt.UTC()

		local_output.SetTextContent(f("local", t))
		utc_output.SetTextContent(f("UTC0 and GMT", tu))
		local_output_alt.SetTextContent(f("local+", t_alt))
		utc_output_alt.SetTextContent(f("UTC0 and GMT+", tu_alt))
		return
	}

	jquery.NewJQuery(use_current_time_button).On(
		jquery.CLICK,
		func() bool {
			func0(0, false)
			return false
		},
	)

	jquery.NewJQuery(use_current_time_utc_button).On(
		jquery.CLICK,
		func() bool {
			func0(1, false)
			return false
		},
	)

	jquery.NewJQuery(use_custom_RFC_time_button).On(
		jquery.CLICK,
		func() bool {
			func0(2, false)
			return false
		},
	)

	jquery.NewJQuery(use_custom_Unix_Local_time_button).On(
		jquery.CLICK,
		func() bool {
			func0(3, false)
			return false
		},
	)

	jquery.NewJQuery(use_custom_Unix_UTC0_time_button).On(
		jquery.CLICK,
		func() bool {
			func0(3, true)
			return false
		},
	)

	return ret

}
