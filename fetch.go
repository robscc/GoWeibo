package main

import(
	"fmt"
	"flag"
	"exec"
	"json"
)

type STWeetsError struct{
	Request string
	Error_code string
	Error string
}

type STWeetsUser struct{
	Id int64
	Screen_name string
	Name string
	Province string
	City string
	Location string
	Description string
	Url string
	Profile_image_url string
	Domain string
	Gender string
	Followers_count uint
	Friends_count uint
	Statuses_count uint
	Favourites_count uint
	Created_at string
	Following bool
	Allow_all_act_msg bool
	Geo_enabled bool
	Verified bool
}

type STWeetsRetweet struct{
	created_at string
	Id int64
	Text string
	Source string
	Favorited bool
	Truncated bool
	In_reply_to_status_id string
	In_reply_to_user_id string
	In_reply_to_screen_name string
	Thumbnail_pic string
	Bmiddle_pic string
	Original_pic string
	Geo string
	Mid string
	User STWeetsUser
}

type STWeets struct{
	Created_at string
	Id int64
	Text string
	Source string
	Favorited bool
	Truncated bool
	In_reply_to_status_id string
	In_reply_to_user_id string
	In_reply_to_screen_name string
	Geo string
	Mid string
	User STWeetsUser
	Retweeted_status STWeetsRetweet 
	
}

type STweetsFields struct {
	Name string
	Text string
	Date string
	Picture string
}
func ParseTweets(b []byte) (string,bool,[]STWeets){
	var tFilter []STWeets
	err:=json.Unmarshal(b,&tFilter);

	if err == nil {
		fmt.Println("Parse Success");
	} else {
		fmt.Println(err.String());
	}

	for index,value := range tFilter {
		
		fmt.Printf("%d|%s|%s|%s\n",index,value.User.Name,value.Text,value.Created_at);	
	}	
	return "",true,tFilter;
	
}

func ParseError(b []byte) bool {
	var tErrorFilter STWeetsError;
	err:=json.Unmarshal(b,&tErrorFilter);

    if err == nil {
		fmt.Println("");
		fmt.Println(tErrorFilter.Request);
		return true;	
	} else {
		fmt.Println(err.String());
		return false;
	}
	
	return false;
}

func main(){
	var username string;
	var password string;
	var action string;
	var source string = "3215873400";

	var base_url string ="http://api.t.sina.com.cn/";
	var api_map  = map[string]string{"public":"statuses/public_timeline","friends":"statuses/friends_timeline","user":"user_timeline","mention":"statuses/mentions","respost":"statuses/repost","update":"statuses/update","comments":"comments/show"} 

	flag.StringVar(&username,"u","","Set Your weibo Account");
	flag.StringVar(&password,"p","","Set Your weibo Password");
	flag.StringVar(&action,"a","","Set Your weibo Action");
	flag.Parse();

	if username == "" || password=="" || action == "" {
		fmt.Println("Do not empty the Username or Password or Action");
	}

	var target_api string;
	var bTest bool;

	target_api,bTest=api_map[action];
	if bTest == false {
		fmt.Println("Use the wrong action");
		return;
	}

	pcmd := exec.Command("curl","-u",username+":"+password,base_url+target_api+".json?source="+source);
	output,err:= pcmd.Output();

	if err != nil {
		fmt.Println("Get tweets Error!");
		return;
	}


	switch action {
		case "public","friends":
			if ParseError(output) {
				fmt.Println("Wrong account info");
				return;
			}
			ParseTweets(output);
		default :
			fmt.Println("To be compeleted");
	}

//	outstr:=string(output);	
//	fmt.Println(outstr);
}

