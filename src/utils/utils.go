/**
		GDG GoLang Backend Daemon
        Copyright (C) 2014+  Gabriele Baldoni
 */
package utils //nome package

import ( //package importati
    //"strconv"
    "os"
    //"time"
    "fmt"
    "strings"
	"bufio"
    "log"
    "bytes"
)

const CONFFILE = "/etc/gdgbackend/backend.conf"
const KEYFILE = "/etc/gdgbackend/keys"



const (
    ASSERT = 0
    DEBUG = 1
    INFO = 2
    WARNING = 3
    ERROR = 4
)

type Configuration struct {
    Address string
    Port string
	Debug bool
    DBHost string
    DBPort string
    DBUser string
    DBPassword string
    DBName string
}

func ( c Configuration) Check() bool {
	if (c.Address!="" && c.Port!="" && c.DBHost!="" && c.DBPort!="" && c.DBPort!="" && c.DBUser!="" && c.DBPassword!="" && c.DBName!=""){
		return true
	}
	return false
}

/*func LoadHostConfiguration() (string,string){
 
    address:=""
    port:=""


    inputFile, err := os.Open(CONFFILE)
    if CheckError(err) {
        return address,port
    }
 
    defer inputFile.Close()
 
    scanner := bufio.NewScanner(inputFile)

    var line string
 
    for scanner.Scan() {
        line=scanner.Text()

        if !strings.HasPrefix(line,"#"){
            strArr:=strings.Split(line," ")

            switch (strArr[0]){

                case "address": address=strArr[1]
                                break;
                case "port" :   port=strArr[1]
                                break;
                default :       Log(ERROR,"Configuration Loader","Error in configuration file")
                                break;
            } 
        }
    }
    return address,port
}*/


func LoadKeys() []string {

    keys := make([]string, 0, 0)




    inputFile, err := os.Open(KEYFILE)
    if err != nil {
        fmt.Printf("%v", err)
        return keys
    }

    defer inputFile.Close()

    scanner := bufio.NewScanner(inputFile)

    var line string

    for scanner.Scan() {
        line = scanner.Text()
        keys = append(keys, line)
    }
    return keys
}


func LoadConfiguration() Configuration {

    var mConf Configuration

   


    inputFile, err := os.Open(CONFFILE)
    if err!=nil {
		fmt.Printf("%v", err)
        return mConf
    }

    defer inputFile.Close()

    scanner := bufio.NewScanner(inputFile)

    var line string

    for scanner.Scan() {
        line=scanner.Text()
		fmt.Printf("%v\n", line)
		
        if !strings.HasPrefix(line,"#"){
            strArr:=strings.Split(line," ")

            switch (strArr[0]){

            case "address": mConf.Address=strArr[1]
                break;
            case "port" :   mConf.Port=strArr[1]
                break;
			case "debug" :   if(strArr[1]=="true"){
								mConf.Debug=true
							} else {
								mConf.Debug = false

							}
				break;
            case "dbhost" : mConf.DBHost=strArr[1]
                break;
            case "dbport" : mConf.DBPort=strArr[1]
                break;
            case "dbuser" : mConf.DBUser=strArr[1]
                break;
            case "dbpassword" : mConf.DBPassword=strArr[1]
                break;
            case "dbname" : mConf.DBName=strArr[1]
                break;
            default :      Log(ERROR,"Configuration Loader","Error in configuration file near:" +  line )
                break;
            }
        }
    }
    return mConf
}

//TODO scrivere log su /var/log/gdgbackend.log
func Log(level int,tag string,txt string){

	logfile:="/tmp/gdgbackend.log"


    //creo la funzione di log in modo tale da lanciarla in una goroutine separata
    logging:= func (level int,tag string,txt string) {

        var buf bytes.Buffer
        logger := log.New(&buf, "logger: ", log.LstdFlags)


		
        f,err:=os.OpenFile(logfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660);

        if err!=nil{
            fmt.Println("Error: " + err.Error())
        }
        defer f.Close()

        switch(level){
            case ASSERT:
                logger.Print("A: (" + tag+") "+txt+"\n")

                f.Write(buf.Bytes())
                //_,err=f.WriteString("A:"+ strconv.FormatInt(time.Now().UnixNano(),10)+" (" + tag+") "+txt+"\n")

            //fmt.Println()
                break
            case DEBUG:
                logger.Print("D: (" + tag+") "+txt+"\n")
                f.Write(buf.Bytes())
                //_,err=f.WriteString("D:"+strconv.FormatInt(time.Now().UnixNano(),10)+" (" + tag+") "+txt+"\n")
            //fmt.Println("D: (" + tag+") "+txt)
                break
            case INFO:
                logger.Print("I: (" + tag+") "+txt+"\n")
                f.Write(buf.Bytes())
                //_,err=f.WriteString("I:"+ strconv.FormatInt(time.Now().UnixNano(),10)+" (" + tag+") "+txt+"\n")
            //fmt.Println("I: (" + tag+") "+txt)
                break
            case WARNING:
                logger.Print("W: (" + tag+") "+txt+"\n")
                f.Write(buf.Bytes())
                //_,err=f.WriteString("W:"+ strconv.FormatInt(time.Now().UnixNano(),10)+" (" + tag+") "+txt+"\n")
            //fmt.Println("W: (" + tag+") "+txt)
                break
            case ERROR:
                logger.Print("E: (" + tag+") "+txt+"\n")
                f.Write(buf.Bytes())
                //_,err=f.WriteString("E:"+ strconv.FormatInt(time.Now().UnixNano(),10)+" (" + tag+") "+txt+"\n")
            //fmt.Println("E: (" + tag+") "+txt)
                break
            default:
                break
        }


    }
    //lancio la goroutine
    go logging(level,tag,txt)
    
     
}


