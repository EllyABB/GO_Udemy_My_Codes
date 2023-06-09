package main


import(
	"fmt"
	"strconv"
	"strings"
	"time"
    "Mapas1"
) 

//1912
func main() {
	//printData("code","fecha","hora","imei","altitud","latitud","longitud","velocidad","direccion","panico","valid")

	Line := "H157,868998030916866,AAA,135,19.521011,-99.211266,230419001605,A,5,25,0,288,2.3,2242,6413,376944,334|50|75F4|00BE2931,0400,0000|0000|0000|0185|04B7,,,1,0000*7C"
	decodeProcess(Line)
}




//  This Function olny prints 
func printData(s ...string) {
    // line: $$H157,868998030916866,AAA,135,19.521011,-99.211266,230419001605,A,5,25,0,288,2.3,2242,6413,376944,334|50|75F4|00BE2931,0400,0000|0000|0000|0185|04B7,,,1,0000*7C
    
    code := strings.Split(s[0], s[0][3:3])

    //fmt.Println("Aca está s0",s[0][0:1])

    // Type: string. AAA. Hexadecimal
    command_type := s[2]

    if command_type == "AAA"{
            //imei,eventcode,latitude,longitude,gps_utc,valid,speed,heading,mileage,analog:=s[1],s[3],s[4],s[5],s[6],s[7],s[10],s[11],s[14],s[18]
            //imei, eventcode, latitude, longitude, gps_utc, valid, speed, heading := s[1], s[3], s[4], s[5], s[6], s[7], s[10], s[11]
            imei,eventcode,latitude := s[1],s[3],s[4]
            longitude,gps_utc,valid := s[5],s[6],s[7]
            speed,heading,altitude  := s[10],s[11],s[13]
            mileage,analog := s[14],s[18]
            
            fmt.Println("\nEvent Code:")
            Dic2.Event_Code(eventcode,0)
            Dic2.Event_Code(eventcode,1)
            fmt.Println(" ")

            var km float64
            //km,_= strconv.ParseFloat(mileage,64)
            km = km / 1000

            

            year := "20" + gps_utc[0:2]
            month := gps_utc[2:4]
            day := gps_utc[4:6]
            hour := gps_utc[6:8]
            minute := gps_utc[8:10]
            seconds := gps_utc[10:12]
            //generationDate := day + "-" + month+ "-"+ year  + " " + hour + ":"+minute +":"+ seconds
            generationDate := year + "-" + month + "-" + day + "T" + hour + ":" + minute + ":" + seconds
            //"2020-02-20T13:15:22"
            now := time.Now().UTC()
            generationTime := fmt.Sprintf("%02d-%02d-%04d %02d:%02d:%02d",
                now.Day(), now.Month(), now.Year(),
                now.Hour(), now.Minute(), now.Second())
            ny := strconv.Itoa(now.Year())

            if year == ny {
                fmt.Println("Header:",code[0]) // @@ $$

                //  Contains 1 byte. The type is the ASCII, and its value ranges from 0x41 to 0x7A.
                fmt.Println("Data identifier:",code[0]) // Se manda en split. cambiarlo

                //  Indicates the length of characters from the first comma (,) to \r\n. Decimal.
                fmt.Println("Data length:",code[1:]) 

                // Indicates the tracker's IMEI number.
                fmt.Println("IMEI:", imei) 

                //  AAA
                fmt.Println("Command type:",command_type) 

                //  Decimal    
                fmt.Println("Event code**:", eventcode)

                //  Decimal
                fmt.Println("Latitude:", latitude)

                //  Decimal
                fmt.Println("Longitude:", longitude) 

                //  yymmddHHMMSS
                fmt.Println("\nDate and time:") 
                fmt.Println("Date:", generationDate,", Time:", generationTime)

                //  Indicates the GPS signal status. A = Valid || V = Invalid
                fmt.Println("Positioning status:") 
                if valid == "A" {
                    valid = "true"
                    fmt.Println("\tThe GPS is valid.")
                    fmt.Println(" ")
                } else {
                    valid = "false"
                    fmt.Println("\tThe GPS is invalid.")
                    fmt.Println(" ")
                }


                //  Indicates the number of received GPS satellites.
                fmt.Println("Number of satellites:",s[8]) 

                // Value: 0–31 || The signal strength is 12.
                fmt.Println("GSM signal strength:",s[9],"/31") 

                //  Decimal Unit: km/h
                fmt.Println("Speed:", speed) 

                //  Decimal 0 to 359
                fmt.Println("Direction:", heading," (from north)") 

                //  The value ranges from 0.5 to 99.9.
                fmt.Println("HDOP:",s[12]) 
                Dic2.HDOP(s[12])
                fmt.Println("") 


                //  Decimal meter
                fmt.Println("Altitude:", altitude) 
    
                //  Decimal meter Indicates the total mileage
                fmt.Println("Mileage:",mileage) 
    
                //  Decimal meter Indicates the total time.
                fmt.Println("Run time:",s[15]) 
    
                //  The base station information includes: MCC|MNC|LAC|CI
                fmt.Println("Base station info**:",s[16]) 
    
                //  Status values of eight input ports and eight output ports
                fmt.Println("I/O port status",s[17]) 
    
                //  Hexadecimal Eight analog input values are separated by
                fmt.Println("Analog input value",analog)  
            }
        }else{
            fmt.Println("Header:",code[0]) // @@ $$
            fmt.Println("Data identifier:",s[0][3:3]) // Contains 1 byte. The type is the ASCII, and its value ranges from 0x41 to 0x7A.
            fmt.Println("Data length:",code[1]) // Indicates the length of characters from the first comma (,) to \r\n. Decimal.
            fmt.Println("IMEI:",s[1]) // Indicates the tracker's IMEI number.
            fmt.Println("Command type:",s[2]) // AAA
            fmt.Println("HEX Data",s[3]) // AAA
        }


}






func FloatToString(input_num float64, presicion int) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', presicion, 64)
}




// This function decode the information
func decodeProcess(line string) {
		// line: $$H157,868998030916866,AAA,135,19.521011,-99.211266,230419001605,A,5,25,0,288,2.3,2242,6413,376944,334|50|75F4|00BE2931,0400,0000|0000|0000|0185|04B7,,,1,0000*7C
		
		// Type s []string split by ,
		s := strings.Split(line, ",")

        fmt.Println("Command Description")
        Dic1.Command_type_map(s[0])
        fmt.Println(" ")

        printData(s...)

}


