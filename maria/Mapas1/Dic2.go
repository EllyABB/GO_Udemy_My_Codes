package Mapas

import(
	"fmt"
    "strconv"
) 



func Event_Code(n string, n2 int16)  {

    Event_Code:=map[string][]string{
        "1": []string{"Input 1 Active", "In1 Active"},
        "2": []string{" Input 2 Active", " In2 Active"},
        "3": []string{" Input 3 Active", " In3 Active"},
        "9": []string{" Input 1 Inactive", " In1 Inactive"},
        "10": []string{" Input 2 Inactive", " In2 Inactive"},
        "11": []string{" Input 3 Inactive", " In3 Inactive"},
        "17": []string{" Low Battery", " Low Battery"},
        "18": []string{" Low External Battery", " Low Ext-Battery"},
        "19": []string{" Speeding", "Speeding"},
        "20": []string{" Enter Geo-fence", " Enter Fence N (N means the number of the fence)"},
        "21": []string{" Exit Geo-fence", " Exit Fence N (N means the number of the fence)"},
        "22": []string{" External Battery On", " Ext-Battery On"},
        "23": []string{" External Battery Cut", " Ext-Battery Cut"},
        "24": []string{" GPS Signal Lost", " GPS Signal Lost"},
        "25": []string{" GPS Signal Recovery", " GPS Recovery"},
        "26": []string{" Enter Sleep", " Enter Sleep"},
        "27": []string{" Exit Sleep", " Exit Sleep"},
        "28": []string{" GPS Antenna Cut", " GPS Antenna Cut"},
        "29": []string{" Device Reboot", " Power On"},
        "31": []string{" Heartbeat ", "/"},
        "32": []string{" Cornering ", "Cornering"},
        "33": []string{" Track By Distance ", "Distance"},
        "34": []string{" Reply Current (Passive) ", "Now"},
        "35": []string{" Track By Time Interval", " Interval"},
        "36": []string{" Tow ", "Tow"},
        "37": []string{" iButton/RFID", " (Only for GPRS)"},
        "39": []string{" Photo", " /"},
        "40": []string{" Power Off", " Power Off"},
        "41": []string{" Stop Moving", " Stop moving"},
        "42": []string{" Start Moving", " Start Moving"},
        "44": []string{" GSM Jamming", " GSM Jamming"},
        "50": []string{" Temperature High ", "Temp High"},
        "51": []string{" Temperature Low", " Temp Low"},
        "52": []string{" Full Fuel ", " Full Fuel"},
        "53": []string{" Low Fuel", " Low Fuel"},
        "54": []string{" Fuel Theft ", " Fuel Theft"},
        "63": []string{" No GSM Jamming ", " No GSM Jamming"},
        "70": []string{" Reject Incoming Call", "  /"},
        "71": []string{" Get Location by Call", "  /"},
        "72": []string{" Auto Answer Incoming Call", "  /"},
        "73": []string{" Listen-in (Voice Monitoring) ", " /"},
        "78": []string{" Impact " , " Impact"},
        "82": []string{" Fuel Filling ", " Fuel Filling"},
        "83": []string{" Ult-Sensor Drop ", " Ult-Sensor Drop"},
        "90": []string{" Sharp Turn to Left ", " Harsh Cornering"},
        "91": []string{" Sharp Turn to Right", "  Harsh Cornering"},
        "94": []string{" Output 1 Active ", " Out1 Active"},
        "95": []string{" Output 2 Active ", " Out2 Active"},
        "96": []string{" Output 1 Inactive ", " Out1 Inactive"},
        "97": []string{" Output 2 Inactive", "  Out2 Inactive"},
        "129": []string{" Harsh Braking ", " Harsh Braking"},
        "130": []string{" Harsh Acceleration", "  Fast Accelerate"},
        "133": []string{" Idle Overtime", "  Idle Overtime"},
        "134": []string{" Idle Recovery", "  Idle Recovery"},
        "135": []string{" Fatigue Driving", "  Fatigue Driving"},
        "136": []string{" Enough Rest after Fatigue Driving ", " Enough Rest"},
    }
    fmt.Println(Event_Code[n][n2])  
    //return Event_Code[n] 
} 

func HDOP(Value string){
    vf, _ := strconv.ParseFloat(Value, 8)

    if vf == 0 {
        fmt.Println("\tThe signal is invalid")
    } else if vf > 0.5 && vf < 1.0 {
        fmt.Println("\tThe signal is Perfect")
    } else if vf > 2.0 && vf < 3.0 {
        fmt.Println("\tThe signal is Wonderful")
    } else if vf > 4.0 && vf < 6.0 {
        fmt.Println("\tthe signal is Good")
    } else if vf > 7.0 && vf < 8.0 {
        fmt.Println("\tThe signal is Medium")
    } else if vf > 9.0 && vf < 20.0 {
        fmt.Println("\tThe signal is Below average")
    } else if vf > 21.0 && vf < 99.9 {
        fmt.Println("\tThe signal is Poor")
    } else {
        fmt.Println("\tOkay, Houston, we've had a problem here.")
    }
}
